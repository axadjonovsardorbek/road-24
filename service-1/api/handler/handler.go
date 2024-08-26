package handler

import (
	"auth/grpc/clients"
	"auth/service"
)

type Handler struct {
	User *service.UsersService
	srvs *clients.GrpcClients
}

func NewHandler(us *service.UsersService, srvs *clients.GrpcClients) *Handler {
	return &Handler{User: us, srvs: srvs}
}
