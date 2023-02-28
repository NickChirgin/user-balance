package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nickchirgin/user-balance/internal/helpers"
	"github.com/nickchirgin/user-balance/internal/models"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", initial).Methods(http.MethodGet)
	router.HandleFunc("/{id}", balance).Methods(http.MethodGet)
	router.HandleFunc("/{id}/add", addMoney).Methods(http.MethodPost)
	router.HandleFunc("/{id}/reserve", reserve).Methods(http.MethodPost)
	router.HandleFunc("/{id}/send", send).Methods(http.MethodPost)
}

func initial(w http.ResponseWriter, r *http.Request) {
	msg, _ := json.Marshal("yo")
	w.Write(msg)
}

func addMoney(w http.ResponseWriter, r *http.Request) { // user id, amount of money
	vars := mux.Vars(r)	
	w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(vars["id"])
	body, err := ioutil.ReadAll(r.Body)
	helpers.CheckErr(err)
	defer r.Body.Close()
	money, e := strconv.Atoi(string(body))
	helpers.CheckErr(e)
	user := models.FindUser(id)
	user.AddBalance(money)
}

func reserve(w http.ResponseWriter, r *http.Request) { // userid, taksid, orderid, cost 
	vars := mux.Vars(r)	
	w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(vars["id"])
	order := models.Order{UserId: id}
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := dec.Decode(&order)
	helpers.CheckErr(err)
	user := models.FindUser(id)
	e := user.ReserveBalance(order.Cost)
	helpers.CheckErr(e)	
	models.CreateOrder(user.Id, order.TaskId, order.Cost)
}

func send(w http.ResponseWriter, r *http.Request) { // userid, taksid, orderid, cost 
	vars := mux.Vars(r)	
	w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(vars["id"])
	user := models.FindUser(id)
	models.FindUser(id)
	order := models.Order{UserId: id}
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := dec.Decode(&order)
	helpers.CheckErr(err)
	e := user.SendMoney(order.Cost)
	helpers.CheckErr(e)
	models.AddToFinance(user.Id, order.TaskId, order.Cost)
}

func balance(w http.ResponseWriter, r *http.Request) { //user id
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(vars["id"])
	user := models.FindUser(id)
	balance, _ := json.Marshal(user.Balance)
	w.Write(balance)
}