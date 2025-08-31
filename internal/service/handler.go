package service

type Update struct {
	UpdateID int64 `json:"update_id"`
	Message  struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		}
	} `json:"message"`
}
