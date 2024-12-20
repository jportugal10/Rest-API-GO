package service

import (
	"encoding/json"
	"log"
	"os"
)

type Data struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Payload struct {
	Data []Data
}

func raw() ([]Data, error) {
	r, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var payload Payload
	err = json.Unmarshal(r, &payload.Data)
	if err != nil {
		log.Printf("Unable to unmarshal JSON due to %s", err)
		return nil, err
	}
	return payload.Data, nil
}

func GetAll() ([]Data, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetById(idx int) (any, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	if idx > len(data) {
		res := make([]string, 0)
		return res, nil
	}
	return data[idx], nil
}
