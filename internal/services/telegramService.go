package services

import (
	"fmt"
	"net/http"
	"racer/form/config"
	"racer/form/internal/models"
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
	//TODO implement me
	//panic("implement me")
	fmt.Println("Here will be updates")
	return nil, nil
}
