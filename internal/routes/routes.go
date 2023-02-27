package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", initial).Methods(http.MethodGet)
}

func initial(w http.ResponseWriter, r *http.Request) {
	msg, _ := json.Marshal("yo")
	w.Write(msg)
}