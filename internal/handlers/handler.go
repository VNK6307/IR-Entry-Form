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

	fmt.Println("ChatID:", chatID, "Text:", text, "MessageID:", msgID)

	switch text {
	case "/personal":
		fmt.Println("Здесь будет заполнение личной заявки.") //ToDo Realize me!
	case "/team":
		fmt.Println("Здесь будет заполнение командной заявки.") //ToDo Realize me!
	case "/list":
		fmt.Println("Здесь будет реализован вывод списка заявленных пилотов.") //ToDo Realize me!
	default:
		fmt.Println("Здесь будет реализован ответ в случае получения неизвестной команды.") //ToDo Realize me!

	}

}

func (handler *Handler) handleCallbackQuery(query any) {
	//ToDo Realize me!
}
