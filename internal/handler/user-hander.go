package handler

import (
	"encoding/json"
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
		formatter.ErrorResponse(500, "internal server error", err.Error())
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	formatter.SuccessResponse(200, "user data succesfully fetched", users)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
