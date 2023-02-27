package models

import (
	"database/sql"
	"errors"

	"github.com/nickchirgin/user-balance/internal/helpers"
)
  
type User struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Balance int `json:"balance" db:"balance"`
	Reserve int `json:"reserve" db:"reserve"`
	Db *sql.DB
}

func Find(id int, db *sql.DB) *User{
	stmt, err := db.Prepare("SELECT id, name, balance, reserve FROM users WHERE id=$1;")
	helpers.CheckErr(err)
	user := User{Db: db}
	defer stmt.Close()
	result := stmt.QueryRow(id)	
	e := result.Scan(&user.Id, &user.Reserve, &user.Balance, &user.Name)
	helpers.CheckErr(e)
	return &user
}


func (u *User) AddBalance(amount int) {
	stmt, err := u.Db.Prepare("UPDATE users SET balance=$1 WHERE id=$2;")
	helpers.CheckErr(err)
	defer stmt.Close()
	updErr := stmt.QueryRow(amount, u.Id).Scan(&u.Id, &u.Reserve, &u.Balance, &u.Name)
	helpers.CheckErr(updErr)
}

func (u *User) ReserveBalance(amount int) error {
	stmt, err := u.Db.Prepare("UPDATE users SET balance=$1, reserve=$2 WHERE id=$3;")
	helpers.CheckErr(err)
	defer stmt.Close()
	if u.Balance - amount >= 0 {
		updErr := stmt.QueryRow(u.Balance - amount, amount, u.Id).Scan(&u.Id, &u.Reserve, &u.Balance, &u.Name)
		helpers.CheckErr(updErr)
		return nil
	}
	return errors.New("Reserve amount greater than balance")
}

func (u *User) SendMoney(amount int) error {
	stmt, err := u.Db.Prepare("UPDATE users SET reserve=$1 WHERE id=$2;")
	helpers.CheckErr(err)
	defer stmt.Close()
	if u.Reserve - amount >= 0 {
		updErr := stmt.QueryRow(u.Reserve - amount, u.Id).Scan(&u.Id, &u.Reserve, &u.Balance, &u.Name)
		helpers.CheckErr(updErr)
		return nil
	}
	return errors.New("Amount greater than reserve")
}