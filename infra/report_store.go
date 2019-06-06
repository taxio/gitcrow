package infra

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
	"google.golang.org/grpc/grpclog"
)

type reportStoreImpl struct {
	Url     string
	channel string
	botName string
	botIcon string

	baseDir string
}

func NewReportStore(url, channel, botName, botIcon, baseDir string) repository.ReportStore {
	return &reportStoreImpl{
		Url:     url,
		channel: channel,
		botName: botName,
		botIcon: botIcon,
		baseDir: baseDir,
	}
}

type xhookRequest struct {
	Channel  string   `json:"channel"`
	Mentions []string `json:"mentions"`
	Message  string   `json:"message"`
	BotName  string   `json:"bot_name"`
	BotIcon  string   `json:"bot_icon"`
}

func (s *reportStoreImpl) Notify(ctx context.Context, username, message string) error {
	data := xhookRequest{
		Channel:  s.channel,
		Mentions: []string{username},
		Message:  message,
		BotName:  s.botName,
		BotIcon:  s.botIcon,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}

	res, err := http.Post(s.Url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.WithStack(err)
	}
	if res.StatusCode >= http.StatusInternalServerError {
		return errors.New("Server Internal Error")
	} else if res.StatusCode >= http.StatusBadRequest {
		return errors.New("Bad Request")
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
			grpclog.Errorf("%+v\n", errors.WithStack(err))
		}
	}()

	w := csv.NewWriter(file)
	for _, r := range repos {
		var scss string
		if r.Success {
			scss = "success"
		} else {
			scss = "failed"
		}

		s := []string{
			r.GitRepo.Owner,
			r.GitRepo.Repo,
			r.GitRepo.Tag,
			scss,
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
