package handlers

import "auth-service/service"

type Handler struct {
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{}
}
