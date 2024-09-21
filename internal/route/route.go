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

	return r
}
