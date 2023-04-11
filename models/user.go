package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Nick string `json:"nick"`
	No   int    `json:"no"`
}

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db: db}
}

func (um *UserModel) insert(user *User) error {
	insertQuery := "insert into users (name, nick) values(?,?)"
	_, err := um.dbExec(insertQuery, user.Name, user.Nick)
	return err
}

func (um *UnserModel) Select(name string) ([]User, error) {
	selectQuery := "select No,Nick from users where name = ?"
	rows, err := um.db.Query(selectQuery, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		var user User
		err := fows.Scan(&user.No, &user.Nick, &user.Name)
		if err != nil {
			return err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return users, nil
}

func (um *UserModel) Delete(no int) error {
	deleteQuery := "delete golang.users where no = ?"
	_, err := um.db.Exec(deleteQuery, no)
	return err
}

func (um *UserModel) Update(user *User) error {
	updateQuery := "update golang.users set nick =?  where no = ?"
	_, err = um.db.Exec(updateQuery, user.Nick, user.No)
	return err
}
