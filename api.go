package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type HealthResponse struct {
	Database     bool `json:"database"`
	Chromedriver bool `json:"chromedriver"`
}

func getEcoindexHealth(url string) error {
	var healthResponse HealthResponse

	res, err := http.Get(url + "/health")
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("Status code is " + res.Status)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("response body can not be read")
	}

	if err := json.Unmarshal(resBody, &healthResponse); err != nil {
		return errors.New("response body is not valid json")
	}

	if !healthResponse.Chromedriver || !healthResponse.Database {
		return errors.New("Ecoindex is KO: " + string(resBody))
	}

	return nil
}
