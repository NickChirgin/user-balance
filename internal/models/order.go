package models

import (
	"github.com/nickchirgin/user-balance/internal/helpers"
	"github.com/nickchirgin/user-balance/internal/storage"
)

type Order struct {
	Id int `db:"id"`
	UserId int
	TaskId int `json:"taskid"`
	Cost int `json:"cost" db:"cost"`
}

func CreateOrder(userId, taskId, cost int) {
	stmt, err := storage.Client.Prepare("INSERT INTO orders (userid, taksid, cost) VALUES ($1, $2, $3) RETURNING id;")
	helpers.CheckErr(err)
	stmt.QueryRow(userId, taskId, cost)
}

func AddToFinance(userId, taskId, cost int) {
	stmt, err := storage.Client.Prepare("INSERT INTO finances (userid, taksid, cost) VALUES ($1, $2, $3) RETURNING id;")
	helpers.CheckErr(err)
	stmt.QueryRow(userId, taskId, cost)
}