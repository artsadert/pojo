package models

type Task struct {
	ID        uint64 `json:"id"`
	Value     string `json:"value"`
	Completed bool   `json:"completed"`
}
