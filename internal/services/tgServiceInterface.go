package services

import "racer/form/internal/models"

type TelegramService interface {
	GetUpdates(offset int) ([]models.Update, error)
	SendMessage(chatID uint64, text string) (int, error)
}
