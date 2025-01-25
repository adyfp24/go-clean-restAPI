package handler

import (
	"encoding/json"
	"go-clean-restAPI/internal/dto"
	"go-clean-restAPI/internal/formatter"
	"go-clean-restAPI/internal/usecase"
	"net/http"
)

type UserHandler interface {
	GetAllUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
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

	if len(users) == 0 {
		formatter.NotFoundResponse(w, "user data not found")
		return
	}

	formatter.SuccessResponse(w, "user data succesfully fetched", map[string]interface{}{"users": users})
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequestDTO dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequestDTO)
	if err != nil {
		formatter.BadRequestResponse(w, "bad request", err.Error())
		return
	}

	if userRequestDTO.Name == "" || userRequestDTO.Address == "" || userRequestDTO.Age <= 0 || userRequestDTO.Role == "" {
		formatter.BadRequestResponse(w, "bad request", "name or address or age or role is invalid")
		return
	}

	newUser, err := h.userUC.CreateUser(userRequestDTO)
	if err != nil {
		formatter.ErrorResponse(w, 500, "error when create user data", err.Error())
		return
	}

	formatter.CreatedResponse(w, "success store new user data", newUser)
}
