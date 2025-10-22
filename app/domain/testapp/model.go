package testapp

import "encoding/json"

type status struct {
	Status string
}

func (app status) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}
