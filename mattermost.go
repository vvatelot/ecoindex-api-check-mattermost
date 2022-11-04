package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func sendMessage(message string, url string) error {
	values := map[string]string{"text": message}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}

	println(resp.Status)

	return nil
}
