package handlers

import (
	"fmt"
	"racer/form/internal/models"
	"racer/form/internal/repositories"
	"racer/form/internal/services"
	"strings"
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

func (handler *Handler) handleCallbackQuery(callbackQuery *models.CallbackQuery) {
	//TODO Realize me!
	chatID := uint64(callbackQuery.Message.Chat.ID)
	//userID := callbackQuery.From.ID
	fmt.Printf("ChatID with buttons: %d\n", chatID) // TODO Delete!

	parts := strings.SplitN(callbackQuery.Data, ":", 2)
	if len(parts) != 2 {
		handler.tlgService.SendMessage(chatID, "Ошибка: неверный формат ответа") // TODO Такая ошибка возможна?
		return
	}

	if parts[0] == "team" {
		handler.handleTeamCallbacks(chatID, parts[1])
		return
	}

}
