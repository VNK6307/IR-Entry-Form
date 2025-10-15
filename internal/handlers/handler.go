package handlers

import (
	"fmt"
	"racer/form/internal/models"
	"racer/form/internal/services"
)

var State = make(map[uint64]uint8)
var defaultText = "Неизвестная команда. Выберите нужную из меню."

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
	fmt.Printf("%+v\n", State[chatID])

	switch text {
	case "/personal":
		{
			handler.startPersonalForm(chatID)
			return // TODO Это для чего? Что будет при отсутствии?
		}
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

func (handler *Handler) checkState(chatID uint64, text string) {

	switch State[chatID] {
	case WaitingTeamNameState:
		handler.saveTeamName(chatID, text)
	case WaitingNextCompetitorState:
	// TODO Realize keyboard
	case WaitingChoiceState:
		// TODO Realize case
		fmt.Println("Waiting user's choice.")
	default:
		handler.SendDefault(chatID, defaultText)
	}
}
