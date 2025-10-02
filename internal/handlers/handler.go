package handlers

import (
	"fmt"
	"racer/form/internal/models"
	"racer/form/internal/services"
)

type Handler struct {
	tlgService services.TelegramService
}

func NewHandler(telegram services.TelegramService) *Handler {
	return &Handler{telegram}
}

func (handler *Handler) HandleUpdate(upd models.Update) {
	if upd.Message != nil {
		handler.handleMessage(upd.Message)
	} else if upd.CallbackQuery != nil {
		handler.handleCallbackQuery(upd.CallbackQuery)
	}
}

func (handler *Handler) handleMessage(message *models.Message) {
	chatID, text, msgID := message.Chat.ID, message.Text, message.MessageID

	fmt.Println("ChatID:", chatID, "Text:", text, "MessageID:", msgID)

	switch text {
	case "/personal":
		{
			handler.startPersonalForm(chatID)
			return // ToDo Это для чего? Что будет при отсутствии?
		}
	case "/team":
		handler.startTeamForm(chatID)
	case "/list":
		handler.sendCompetitors(chatID)
	case "/send":
		handler.mailList(chatID)
	default:
		handler.sendDefault(chatID)
	}
}

func (handler *Handler) handleCallbackQuery(query any) {
	//ToDo Realize me!
}
