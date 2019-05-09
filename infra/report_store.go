package infra

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type reportStoreImpl struct {
	webHookURL string
	channel    string
	botName    string
	botIcon    string

	baseDir string
}

func NewReportStore(webHookURL, channel, botName, botIcon, baseDir string) repository.ReportStore {
	return &reportStoreImpl{
		webHookURL: webHookURL,
		channel:    channel,
		botName:    botName,
		botIcon:    botIcon,
		baseDir:    baseDir,
	}
}

type slackData struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}

func (s *reportStoreImpl) Notify(ctx context.Context, slackId, message string) error {
	message = fmt.Sprintf("<@%s> %s", slackId, message)
	data := slackData{
		Text:      message,
		Username:  s.botName,
		IconEmoji: s.botIcon,
		Channel:   s.channel,
	}
	jsonParams, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = http.PostForm(s.webHookURL, url.Values{"payload": {string(jsonParams)}})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *reportStoreImpl) Save(ctx context.Context) error {
	return nil
}

func (s *reportStoreImpl) ReportToFile(ctx context.Context, username, projectName string, repos []*model.Report) error {
	// create report file
	t := time.Now()
	filename := fmt.Sprintf("%d-%d-%d_%d-%d-%d_report.csv", t.Year(), t.Month(), t.Day(), t.Hour(), t.Hour(), t.Minute())
	filename = filepath.Join(s.baseDir, username, projectName, filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			grpclog.Errorln(errors.WithStack(err))
		}
	}()

	w := csv.NewWriter(file)
	for _, r := range repos {
		s := []string{
			r.GitRepo.Owner,
			r.GitRepo.Repo,
			r.GitRepo.Tag,
			r.Code.ToString(),
			r.Message,
		}
		err := w.Write(s)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
