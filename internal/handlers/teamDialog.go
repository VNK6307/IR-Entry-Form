package handlers

import (
	"fmt"
	"racer/form/internal/repositories"
)

const (
	StateNone                  = 0
	WaitingTeamNameState       = 1
	WaitingNextCompetitorState = 2
	WaitingChoiceState         = 3
)

func (handler *Handler) startTeamForm(chatID uint64) {
	//ToDo Realize me!

	State[chatID] = WaitingTeamNameState // ToDo add mutex

	_, err := handler.tlgService.SendMessage(chatID, "Введите название вашей команды.")
	if err != nil {
		return // ToDo Обработать ошибку???
	}
}

func (handler *Handler) saveTeamName(chatID uint64, text string) {
	if text == "" {
		_, err := handler.tlgService.SendMessage(chatID, "Введите название вашей команды.")
		if err != nil {
			repositories.SaveTeamName(text)
		}
	}
	fmt.Printf("Название команды %s\n", text)
	State[chatID] = WaitingNextCompetitorState

	_, err := handler.tlgService.SendMessage(chatID, "Введите Фамилию и имя участника")
	if err != nil {
		return
	}
	// TODO Stopped here 0710
}
