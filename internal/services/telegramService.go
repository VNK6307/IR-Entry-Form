package services

import (
	"encoding/json"
	"net/http"
	"racer/form/config"
	"racer/form/internal/models"
	"time"
)

type TelegramService interface {
	GetUpdates(offset int) ([]models.Update, error)
}

type telegramService struct {
	baseURL string
	client  *http.Client
}

func NewTelegramService(cfg *config.Config) TelegramService {
	return &telegramService{
		baseURL: cfg.BotURL,
		client:  &http.Client{},
	}
}

func (tgSvc *telegramService) GetUpdates(offset int) ([]models.Update, error) {

	result, err := tgSvc.makeRequest("getUpdates", map[string]interface{}{
		"offset":  offset,
		"timeout": 30 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	var rslt struct {
		Result []models.Update `json:"result"`
	}

	err = json.Unmarshal(result, &rslt)
	if err != nil {
		return nil, err
	}
	return rslt.Result, nil
}
