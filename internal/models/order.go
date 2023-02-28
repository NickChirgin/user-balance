package models

import (
	"github.com/nickchirgin/user-balance/internal/helpers"
	"github.com/nickchirgin/user-balance/internal/storage"
)

type Order struct {
	Id int `json:"id" db:"id"`
	UserId int
	TaskId int
	Cost int `json:"cost" db:"cost"`
}

func CreateOrder(userId, taskId, cost int) {
	stmt, err := storage.Client.Prepare("INSERT INTO orders (userid, taskid, cost) VALUE ($1, $2, $3) RETURNING id;")
	helpers.CheckErr(err)
	stmt.QueryRow(userId, taskId, cost)
}

func AddToFinance(userId, taskId, cost int) {
	stmt, err := storage.Client.Prepare("INSERT INTO finances (userid, taskid, cost) VALUE ($1, $2, $3) RETURNING id;")
	helpers.CheckErr(err)
	stmt.QueryRow(userId, taskId, cost)
}