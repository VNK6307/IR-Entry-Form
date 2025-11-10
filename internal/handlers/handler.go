package handlers

import (
	"fmt"
	"racer/form/internal/models"
	"racer/form/internal/repositories"
	"racer/form/internal/services"
)

var State = make(map[uint64]uint8)
var defaultText = "Неизвестная команда. Выберите нужную из меню."

type Handler struct {
	tlgService services.TelegramService
	teamRepo   map[uint64]*repositories.Team
}

func NewHandler(telegram services.TelegramService) *Handler {
	return &Handler{
		telegram,
		make(map[uint64]*repositories.Team),
	}
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
	fmt.Printf("State - %+v\n", State[chatID]) // TODO Delete before completion

	switch text {
	case "/personal":
		handler.startPersonalForm(chatID)
	case "/team":
		handler.startTeamForm(chatID)
	case "/list":
		handler.sendCompetitors(chatID)
	case "/send":
		handler.mailList(chatID)
	default:
		handler.checkState(chatID, text)
	}
}

func (handler *Handler) handleCallbackQuery(query any) {
	//TODO Realize me!
}
