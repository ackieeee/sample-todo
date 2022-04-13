package entity

import "encoding/json"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Status      bool   `json:"status"`
}

type Tasks []Task

func (ts *Tasks) ToJson() ([]byte, error) {
	return json.Marshal(ts)
}
