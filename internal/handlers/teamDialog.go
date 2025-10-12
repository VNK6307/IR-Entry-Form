package handlers

import "fmt"

const (
	NoneState                  = 0
	WaitingTeamNameState       = 1
	WaitingNextCompetitorState = 2
	WaitingChoiceTeamState     = 3
)

func (handler *Handler) startTeamForm(chatID uint64) {
	//TODO Realize me!

	State[chatID] = WaitingTeamNameState // TODO add mutex

	_, err := handler.tlgService.SendMessage(chatID, "Введите название вашей команды.")
	if err != nil {
		return // TODO Обработать ошибку???
	}
}

func (handler *Handler) saveTeamName(chatID uint64, text string) {
	if text == "" { // TODO Обработать этот вариант. Из телеграм не отправляются пустые сообщения.
		_, err := handler.tlgService.SendMessage(chatID, "Введите название вашей команды.")
		if err != nil {
			return
		}
	}
	fmt.Printf("Название команды %s\n", text) // TODO Delete before finish

	// TODO Create and save team

	_, err := handler.tlgService.SendMessage(chatID, "Введите Фамилию и имя участника")
	if err != nil {
		return
	}
	State[chatID] = WaitingNextCompetitorState

	// TODO Stopped here 0710
}

func (handler *Handler) saveTeamMember(chatID uint64, text string) {

}
