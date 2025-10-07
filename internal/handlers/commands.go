package handlers

import (
	"fmt"
	"racer/form/internal/models"
)

func (handler *Handler) startPersonalForm(chatID uint64) {

	//ToDo Realize me!

	personalForm := models.LoadFormQuestions("personal")
	for _, q := range personalForm {
		fmt.Println(q)

	}
	_, err := handler.tlgService.SendMessage(chatID, "Здесь будет заполнение личной заявки.")
	if err != nil {
		return // ToDo Обработать ошибку???
	}
}

//func (handler *Handler) startTeamForm(chatID uint64) {
//	//ToDo Realize me!
//
//	teamForm := models.LoadFormQuestions("team")
//	for _, q := range teamForm {
//		fmt.Println(q)
//	}
//
//	_, err := handler.tlgService.SendMessage(chatID, "Здесь будет заполнение командной заявки.")
//	if err != nil {
//		return // ToDo Обработать ошибку???
//	}
//}

func (handler *Handler) sendCompetitors(chatID uint64) {
	//ToDo Realize me!
	_, err := handler.tlgService.SendMessage(chatID, "Здесь будет реализован вывод списка заявленных пилотов.")
	if err != nil {
		return // ToDo Обработать ошибку???
	}
}
func (handler *Handler) sendDefault(chatID uint64) {
	//ToDo Realize me!
	_, err := handler.tlgService.SendMessage(chatID, "Здесь будет реализован ответ в случае получения неизвестной команды.")
	if err != nil {
		return // ToDo Обработать ошибку???
	}
}

func (handler *Handler) mailList(chatID uint64) { // ToDo Нужно разделять роли либо как-то иначе организовывать получение списков организаторами
	//ToDo Realize me!
	_, err := handler.tlgService.SendMessage(chatID, "Здесь должна быть реализована отправка файла со списками.")
	if err != nil {
		return // ToDo Обработать ошибку???
	}
}
