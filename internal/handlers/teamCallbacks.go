package handlers

import "fmt"

func (handler *Handler) handleTeamCallbacks(chatID uint64, command string) {

	var err error

	switch command {
	case "checkTeam":
		answer := fmt.Sprintf("Команда: %s\n", handler.teamRepo[chatID].TeamName)
		for i := range handler.teamRepo[chatID].TeamMember {
			answer += fmt.Sprintf("Участник № %d %s\n", i+1, handler.teamRepo[chatID].TeamMember[i])
		}

		_, err = handler.tlgService.SendMessage(chatID, answer)
		if err != nil {
			fmt.Println(err)
			return
		}

	case "nextTeamMember":
		_, err = handler.tlgService.SendMessage(chatID, "Введите Имя и Фамилию следующего участника:\n")
		if err != nil {
			return
		}
		State[chatID] = WaitingCompetitorState

	case "sendForm":
		State[chatID] = StateNone
		err = handler.teamRepo[chatID].SaveTeam("")
		if err != nil {
			return
		}
	}

}
