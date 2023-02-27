package models

type Order struct {
	Id int `json:"id" db:"id"`
	UserId int
	TaskId int
	Cost int `json:"cost" db:"cost"`
}

func Create(userId, taskId, cost int) {

}