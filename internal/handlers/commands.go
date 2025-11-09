package handlers

import (
	"log"
)

func (handler *Handler) startPersonalForm(chatID uint64) {
	//TODO Realize me!

	_, err := handler.tlgService.SendMessage(chatID, "Здесь будет заполнение личной заявки.")
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}
}

func (handler *Handler) sendCompetitors(chatID uint64) {
	//TODO Realize me!
	_, err := handler.tlgService.SendMessage(chatID, "Здесь будет реализован вывод списка заявленных пилотов.")
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}
}
func (handler *Handler) SendDefault(chatID uint64, text string) {
	_, err := handler.tlgService.SendMessage(chatID, text)
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}
}

func (handler *Handler) mailList(chatID uint64) { // TODO Нужно разделять роли либо как-то иначе организовывать получение списков организаторами
	//TODO Realize me!
	_, err := handler.tlgService.SendMessage(chatID, "Здесь должна быть реализована отправка файла со списками.")
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}
}
