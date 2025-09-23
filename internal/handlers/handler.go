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
	//ToDo Realize me!
	chatID, text, msgID := message.Chat.ID, message.Text, message.MessageID
	// ToDo Change to switch
	if text == "/personal" {
		println("Здесь надо начинать обработку сообщений бота")
		fmt.Printf("ChatID: %d \n MessageId: %d \n", chatID, msgID)
	}

}

func (handler *Handler) handleCallbackQuery(query any) {
	//ToDo Realize me!
}
