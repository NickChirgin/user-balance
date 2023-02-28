package models

import (
	"errors"

	"github.com/nickchirgin/user-balance/internal/helpers"
	"github.com/nickchirgin/user-balance/internal/storage"
)
  
type User struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Balance int `json:"balance" db:"balance"`
	Reserve int `json:"reserve" db:"reserve"`
}

func CreateUser(name string, balance int, reserve int) {
	stmt, err := storage.Client.Prepare("INSERT INTO users (name, balance, reserve) VALUES ($1, $2, $3) RETURNING id;")
	helpers.CheckErr(err)
	defer stmt.Close()
	stmt.QueryRow(name, balance, reserve)
}

func FindUser(id int) *User{
	stmt, err := storage.Client.Prepare("SELECT id, name, balance, reserve FROM users WHERE id=$1;")
	helpers.CheckErr(err)
	user := User{}
	defer stmt.Close()
	result := stmt.QueryRow(id)	
	e := result.Scan(&user.Id, &user.Name, &user.Balance, &user.Reserve)
	helpers.CheckErr(e)
	return &user
}


func (u *User) AddBalance(amount int) {
	stmt, err := storage.Client.Prepare("UPDATE users SET balance=$1 WHERE id=$2;")
	helpers.CheckErr(err)
	defer stmt.Close()
	updErr := stmt.QueryRow(u.Balance + amount, u.Id).Scan(&u.Id, &u.Name, u.Balance, u.Reserve)
	helpers.CheckErr(updErr)
}

func (u *User) ReserveBalance(amount int) error {
	stmt, err := storage.Client.Prepare("UPDATE users SET balance=$1, reserve=$2 WHERE id=$3;")
	helpers.CheckErr(err)
	defer stmt.Close()
	if u.Balance - amount >= 0 {
		updErr := stmt.QueryRow(u.Balance - amount, u.Reserve + amount, u.Id).Scan(&u.Id, &u.Name, u.Balance, u.Reserve)
		helpers.CheckErr(updErr)
		return nil
	}
	return errors.New("Reserve amount greater than balance")
}

func (u *User) SendMoney(amount int) error {
	stmt, err := storage.Client.Prepare("UPDATE users SET reserve=$1 WHERE id=$2;")
	helpers.CheckErr(err)
	defer stmt.Close()
	if u.Reserve - amount >= 0 {
		updErr := stmt.QueryRow(u.Reserve - amount, u.Id).Scan(&u.Id, &u.Name, u.Balance, u.Reserve)
		helpers.CheckErr(updErr)
		return nil
	}
	return errors.New("Amount greater than reserve")
}