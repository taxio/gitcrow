package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/domain/repository"
	"net/http"
	"net/url"
)

type reportStoreImpl struct {
	webHookURL string
	channel    string
	botName    string
	botIcon    string

	saveDir string
}

func NewReportStore(webHookURL, channel, botName, botIcon, saveDir string) repository.ReportStore {
	return &reportStoreImpl{
		webHookURL: webHookURL,
		channel:    channel,
		botName:    botName,
		botIcon:    botIcon,
		saveDir:    saveDir,
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
