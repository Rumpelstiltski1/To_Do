package models

type CreateTaskRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
