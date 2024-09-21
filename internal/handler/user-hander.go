package handler

import (
	"encoding/json"
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
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
