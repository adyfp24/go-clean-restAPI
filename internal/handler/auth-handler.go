package handler

import (
	"go-clean-restAPI/internal/usecase"
	"net/http"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authUC usecase.AuthUsecase
}

func NewAuthHandler(authUC usecase.AuthUsecase) AuthHandler {
	return &authHandler{
		authUC: authUC,
	}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {

}
