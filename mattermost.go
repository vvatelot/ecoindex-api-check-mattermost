package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendMessage(message string, url string, name string) error {
	text := fmt.Sprintf("The Ecoindex API %s is in error. Here is the detail: \\n```%s```", name, message)

	values := map[string]string{"text": text}
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
