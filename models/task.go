package models

type Task struct {
	Id         int    `json:"id,omitempty"`
	Title      string `json:"title"`
	Desce      string `json:"desce"`
	Completed  bool   `json:"completed"`
	Created_at string `json:"created_at,omitempty"`
}
