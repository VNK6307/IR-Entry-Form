package handlers

import (
	"fmt"
	"log"
	"racer/form/internal/repositories"
)

const (
	StateNone                   = 0
	WaitingTeamNameState        = 1
	WaitingFirstCompetitorState = 2
	WaitingNextCompetitorState  = 3
	WaitingChoiceState          = 4
)

func (handler *Handler) startTeamForm(chatID uint64) {
	//ToDo Realize me!

	State[chatID] = WaitingTeamNameState // ToDo add mutex

	_, err := handler.tlgService.SendMessage(chatID, "Введите название вашей команды.")
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}
}

func (handler *Handler) saveTeamName(chatID uint64, text string) {
	if text == "" {
		_, err := handler.tlgService.SendMessage(chatID, "Это поле обязательно к заполнению. Введите название вашей команды.")
		if err != nil {
			log.Printf("SendMessage mistake: %v", err)
			return
		}
	}

	handler.teamRepo[chatID] = repositories.NewTeamRepository() //TODO Здесь создан репозиторий

	handler.teamRepo[chatID].TeamName = text

	State[chatID] = WaitingFirstCompetitorState

	_, err := handler.tlgService.SendMessage(chatID, "Это поле обязательно к заполнению. Введите Фамилию и Имя участника")
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}
}

func (handler *Handler) saveTeamMember(chatID uint64, text string) {
	if text == "" {
		if text == "" {
			_, err := handler.tlgService.SendMessage(chatID, "Это поле обязательно к заполнению. Введите Фамилию и Имя участника")
			if err != nil {
				log.Printf("SendMessage mistake: %v", err)
				return
			}
		}
	}

	handler.teamRepo[chatID].TeamMember = append(handler.teamRepo[chatID].TeamMember, text)

	fmt.Printf("Команды %+v\n", handler.teamRepo[chatID]) // TODO Delete before completion

	//State[chatID] = WaitingNextCompetitorState

	_, err := handler.tlgService.SendMessage(chatID, "Введите Фамилию и Имя следующего участника")
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}

	// TODO Send keyboard
	// TODO Кнопки - следующий, посмотреть список, отправить заявку

	// TODO 1510 Продолжить с этого места
	//handler.teamRepo[chatID].SaveTeam("")
}
