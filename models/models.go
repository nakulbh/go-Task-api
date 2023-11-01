package models

type Task struct {
	TaskID      int64  `json:"taskid"`
	Name        string `json:"name"`
	DateCreated string `json:"date"`
	DueDate     string `json:"DueDate"`
}
