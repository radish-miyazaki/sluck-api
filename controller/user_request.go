package controller

import (
	"time"

	"github.com/radish-miyazaki/sluck/model"
)

type UserRequest struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func toModel(req UserRequest) *model.User {
	now := time.Now()
	return &model.User{
		Name:      req.Name,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
