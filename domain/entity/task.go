package entity

import "encoding/json"

type Task struct {
	Title       string
	Description string
	Date        string
	Status      bool
}

type Tasks []Task

func (ts *Tasks) ToJson() ([]byte, error) {
	return json.Marshal(ts)
}
