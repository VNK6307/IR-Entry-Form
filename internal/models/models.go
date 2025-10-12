package models

type UserState struct {
	UserID        int64
	ChatID        int64
	Step          int
	Data          map[string]string
	LastBotMsgID  int
	LastUserMsgID int
}

type InlineButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type Update struct {
	UpdateID      int            `json:"update_id"`
	Message       *Message       `json:"message,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

type Message struct {
	MessageID int    `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text,omitempty"`
}

type Chat struct {
	ID uint64 `json:"id"`
}

type CallbackQuery struct {
	ID      string
	From    struct{ Chat}
	//From    struct{ ID uint64 }
	Message struct{ Chat }
	//Message struct{ Chat struct{ ID uint64 } }
	Data string
}
