package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", initial).Methods(http.MethodGet)
	router.HandleFunc("/:id", balance).Methods(http.MethodGet)
	router.HandleFunc("/:id/add", addMoney).Methods(http.MethodPost)
	router.HandleFunc("/:id/reserve", reserve).Methods(http.MethodPost)
	router.HandleFunc("/:id/send", send).Methods(http.MethodPost)
}

func initial(w http.ResponseWriter, r *http.Request) {
	msg, _ := json.Marshal("yo")
	w.Write(msg)
}

func addMoney(w http.ResponseWriter, r *http.Request) { // user id, amount of money

}

func reserve(w http.ResponseWriter, r *http.Request) { // userid, taksid, orderid, cost 

}

func send(w http.ResponseWriter, r *http.Request) { // userid, taksid, orderid, cost 

}

func balance(w http.ResponseWriter, r *http.Request) { //user id

}