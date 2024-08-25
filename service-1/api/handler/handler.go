package handler

import (
	"auth/service"
)

type Handler struct {
	User     *service.UsersService
}

func NewHandler(us *service.UsersService) *Handler {
	return &Handler{User: us}
}
