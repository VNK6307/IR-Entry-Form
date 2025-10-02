package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (tgSvc *telegramService) makeRequest(method string, payload interface{}) ([]byte, error) {
	data, _ := json.Marshal(payload)
	resp, err := tgSvc.client.Post(fmt.Sprintf("%s/%s", tgSvc.baseURL, method), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("telegram error: %s", string(body))
	}
	return body, nil
}
