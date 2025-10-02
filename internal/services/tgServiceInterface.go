package services

import "racer/form/internal/models"

type TelegramService interface {
	GetUpdates(offset int) ([]models.Update, error)
	SendMessage(chatID int64, text string) (int, error)
}
