package handlers

import (
	"fmt"
	"log"
	"racer/form/internal/models"
	"racer/form/internal/repositories"
)

var teamChoiceQuestion = "<b>Выберите дальнейшее действие:</b>\n"

func (handler *Handler) saveTeamName(chatID uint64, text string) {
	if text == "" {
		_, err := handler.tlgService.SendMessage(chatID, "Это поле обязательно к заполнению.\n Введите название вашей команды.")
		if err != nil {
			log.Printf("SendMessage mistake: %v", err)
			return
		}
	}

	handler.teamRepo[chatID] = repositories.NewTeamRepository() //TODO Здесь создан репозиторий

	handler.teamRepo[chatID].TeamName = text

	State[chatID] = WaitingCompetitorState

	_, err := handler.tlgService.SendMessage(chatID, "Введите Фамилию и Имя участника\nЭто поле обязательно к заполнению.")
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}
}

func (handler *Handler) saveTeamMember(chatID uint64, text string) {

	State[chatID] = WaitingUserChoiceState

	fmt.Printf("State = %+v\n", State[chatID]) // TODO Delete before finish

	handler.teamRepo[chatID].TeamMember = append(handler.teamRepo[chatID].TeamMember, text)

	fmt.Printf("Команда -  %+v\n", handler.teamRepo[chatID]) // TODO Delete before finish

	_, err := handler.tlgService.SendMessage(chatID, teamChoiceQuestion)
	if err != nil {
		return
	}

	handler.askTeamChoice(chatID)

}
func (handler *Handler) askTeamChoice(chatID uint64) {
	row1 := []models.InlineButton{
		{Text: "Следующий участник", CallbackData: "team:nextTeamMember"},
	}

	row2 := []models.InlineButton{
		{Text: "Проверить ввод", CallbackData: "team:checkTeam"},
		{Text: "Отправить заявку", CallbackData: "team:sendForm"},
	}

	teamButtons := [][]models.InlineButton{
		row1,
		row2,
	}

	_, err := handler.tlgService.SendMessageWithKeyboard(chatID, "Варианты:\n", teamButtons)
	if err != nil {
		return
	}
}
