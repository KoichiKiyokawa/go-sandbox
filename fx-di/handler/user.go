package handler

import "log"

type UserHandler struct {
	logger *log.Logger
}

func NewUserHandler(logger *log.Logger) *UserHandler {
	return &UserHandler{logger}
}

func (h *UserHandler) FindOne(id int) {
	h.logger.Printf("Executing FindOne. id: %d", id)

}
