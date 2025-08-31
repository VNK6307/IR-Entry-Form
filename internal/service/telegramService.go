package service

import "net/http"

type telegramService struct {
	baseUrl string
	client  *http.Client
}

//type TgService interface {
//	Updater
//	MessageSender
//}

type Updater interface {
	GetUpdates(offset int64) []Update
}

type MessageSender interface {
	SendMessage(chatId int64, message string) error
}
