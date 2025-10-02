package handlers

import (
	"fmt"
	"racer/form/internal/models"
	"racer/form/internal/services"
)

type Handler struct {
	tlg services.TelegramService
}

func NewHandler(tlg services.TelegramService) *Handler {
	return &Handler{tlg}
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
		handler.startPersonalForm(chatID)
	case "/team":
		handler.startTeamForm(chatID)
	case "/list":
		handler.sendCompetitors(chatID)
	default:
		handler.sendDefault(chatID)
	}
}

func (handler *Handler) handleCallbackQuery(query any) {
	//ToDo Realize me!
}
