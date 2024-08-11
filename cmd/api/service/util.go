package service

import (
	"encoding/json"
	"log"
	"os"
	//"strconv"
)

type Data struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Payload struct {
	Data []Data
}

// type StringInt int

// func (st *StringInt) UnmarshalJSON(b []byte) error {
// 	var item interface{}
// 	if err := json.Unmarshal(b, &item); err != nil {
// 		return err
// 	}
// 	switch v := item.(type) {
// 	case int:
// 		*st = StringInt(v)
// 	case float64:
// 		*st = StringInt(int(v))
// 	case string:
// 		i, err := strconv.Atoi(v)
// 		if err != nil {
// 			return err
// 		}
// 		*st = StringInt(i)
// 	}

// 	return nil
// }

// type Item struct {
// 	Name  string    `json:"name"`
// 	Phone StringInt `json:"phone"`
// }

func raw() ([]Data, error) {
	r, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var payload Payload
	err = json.Unmarshal(r, &payload.Data)
	if err != nil {
		log.Fatalf("Unable to unmarshal JSON due to %s", err)
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
