package route

import (
	"github.com/gorilla/mux"
	"go-clean-restAPI/internal/handler"
)

func InitRoute(h *handler.Handlers) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", h.UserHandler.GetAllUser).Methods("GET")

	return r
}
