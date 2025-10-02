package handlers

import "fmt"

func (handler *Handler) startPersonalForm(chatID int64) {
	fmt.Println("Здесь будет заполнение личной заявки.") //ToDo Realize me!
}

func (handler *Handler) startTeamForm(chatID int64) {
	fmt.Println("Здесь будет заполнение командной заявки.") //ToDo Realize me!
}

func (handler *Handler) sendCompetitors(chatID int64) {
	fmt.Println("Здесь будет реализован вывод списка заявленных пилотов.") //ToDo Realize me!
}
func (handler *Handler) sendDefault(chatID int64) {
	fmt.Println("Здесь будет реализован ответ в случае получения неизвестной команды.") //ToDo Realize me!
}
