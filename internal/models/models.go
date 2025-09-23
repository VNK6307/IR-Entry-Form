package models

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
	MessageID int `json:"message_id"`
	Chat      struct {
		ID int64 `json:"id"`
	}
	Text string `json:"text,omitempty"`
}

type CallbackQuery struct {
	ID      string
	From    struct{ ID int64 }
	Message struct{ Chat struct{ ID int64 } }
	Data    string
}
