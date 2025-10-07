package handlers

import "fmt"

func (handler *Handler) startTeamForm(chatID uint64) {
	//ToDo Realize me!

	State[chatID] = StateAwaitingTeamName // ToDo add mutex

	_, err := handler.tlgService.SendMessage(chatID, "Введите название вашей команды.")
	if err != nil {
		return // ToDo Обработать ошибку???
	}
}

func (handler *Handler) saveTeamName(chatID uint64, text string) {
	if text == "" {
		_, err := handler.tlgService.SendMessage(chatID, "Введите название вашей команды.")
		if err != nil {
			return
		}
	}
	fmt.Printf("Название команды %s\n", text)
	State[chatID] = StateAwaitingNextCompetitor

	_, err := handler.tlgService.SendMessage(chatID, "Введите Фамилию и имя участника")
	if err != nil {
		return
	}

}
