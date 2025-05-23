package models

import "time"

type CreateTaskRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ListTaskResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type DeleteTaskRequest struct {
	Id int `json:"id"`
}

type PutTaskRequest struct {
	Status bool `json:"status"`
	Id     int  `json:"id"`
}
