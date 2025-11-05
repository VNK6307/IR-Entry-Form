package handlers

import "fmt"

const (
	StateNone                   = 0
	WaitingTeamNameState        = 1
	WaitingFirstCompetitorState = 2
	WaitingUserChoiceState      = 3
	WaitingNextCompetitorState  = 4
)

func (handler *Handler) checkState(chatID uint64, text string) {

	switch State[chatID] {
	case WaitingTeamNameState:
		handler.saveTeamName(chatID, text)
	case WaitingFirstCompetitorState:
		handler.saveTeamMember(chatID, text)
	case WaitingNextCompetitorState:
	// TODO Realize keyboard
	case WaitingUserChoiceState:
		// TODO Realize case
		fmt.Println("Waiting user's choice.")
	default:
		handler.SendDefault(chatID, defaultText)
	}
}
