package handlers

import (
	"fmt"
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
		return // ToDo Обработать ошибку??? Add to logger
	}
}

func (handler *Handler) saveTeamName(chatID uint64, text string) {
	if text == "" {
		_, err := handler.tlgService.SendMessage(chatID, "Введите название вашей команды.")
		if err != nil {
			return // ToDo Обработать ошибку??? Add to logger
		}
	}

	handler.teamRepo[chatID] = repositories.NewTeamRepository() //TODO Здесь создан репозиторий

	handler.teamRepo[chatID].TeamName = text

	fmt.Printf("Название команды %+v\n", handler.teamRepo[chatID]) // TODO Delete before end

	State[chatID] = WaitingFirstCompetitorState

	_, err := handler.tlgService.SendMessage(chatID, "Введите Фамилию и Имя участника")
	if err != nil {
		return
	}
}

func (handler *Handler) saveTeamMember(chatID uint64, text string) {
	if text == "" {
		if text == "" {
			_, err := handler.tlgService.SendMessage(chatID, "Введите Фамилию и Имя участника")
			if err != nil {
				return // ToDo Обработать ошибку??? Add to logger
			}
		}
	}

	handler.teamRepo[chatID].TeamMember = append(handler.teamRepo[chatID].TeamMember, text)
	fmt.Printf("Команды %+v\n", handler.teamRepo[chatID]) // TODO Delete before end

	//State[chatID] = WaitingNextCompetitorState

	_, err := handler.tlgService.SendMessage(chatID, "Введите Фамилию и Имя следующего участника")
	if err != nil {
		return
	}

	// TODO Send keyboard
	// TODO Кнопки - следующий, посмотреть список, отправить заявку

	// TODO 1510 Продолжить с этого места
	//handler.teamRepo[chatID].SaveTeam("")
}
