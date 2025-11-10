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
	// TODO Realize keyboard
	case WaitingUserChoiceState:
		fmt.Println("Waiting user's choice.") // TODO Delete before finish
		handler.askTeamChoice(chatID)
		// TODO Realize case

	case WaitingNextCompetitorState:

	default:
		handler.SendDefault(chatID, defaultText)
	}
}
