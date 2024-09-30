package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-clean-restAPI/internal/handler"
	"net/http"
)

func InitRoute(h *handler.Handlers) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcome")
	}).Methods("GET")

	r.HandleFunc("/users", h.UserHandler.GetAllUser).Methods("GET")
	r.HandleFunc("/users", h.UserHandler.CreateUser).Methods("POST")

	//r.HandleFunc("/login", h.AuthHandler.Login).Methods("POST")
	//r.HandleFunc("/register", h.AuthHandler.Register).Methods("POST")

	return r
}
