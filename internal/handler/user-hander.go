package handler

import (
	"go-clean-restAPI/internal/formatter"
	"go-clean-restAPI/internal/usecase"
	"net/http"
)

type UserHandler interface {
	GetAllUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUC usecase.UserUsecase
}

func NewUserHandler(userUC usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUC: userUC,
	}
}

func (h *userHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUC.GetAllUser()

	if err != nil {
		formatter.ErrorResponse(w, 500, "internal server error", err.Error())
		return
	}

	formatter.SuccessResponse(w, "user data succesfully fetched", users)
}
