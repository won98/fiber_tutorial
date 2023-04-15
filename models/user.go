package models

import (
	"database/sql"
	"fmt"

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
	return &UserModel{
		db: db,
	}
}

func (um *UserModel) Insert(user *User) error {
	insertQuery := "insert into users (name, nick) values(?,?)"
	_, err := um.db.Exec(insertQuery, user.Name, user.Nick)
	return err
}

// 배열에 싸는 방법
func (um *UserModel) Select(name string) ([]User, error) {
	selectQuery := "select no,nick,name from golang.users where name = ?"
	rows, err := um.db.Query(selectQuery, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	users := make([]User, 0)

	for rows.Next() {
		var user User
		err := rows.Scan(&user.No, &user.Nick, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

//
// 배열 없이하는 방법
// func (um *UserModel) Select(name string) (User, error) {
// 	selectQuery := "select no,nick,name from golang.users where name = ?"
// 	rows, err := um.db.Query(selectQuery, name)
// 	if err != nil {
// 		return User{}, err
// 	}
// 	defer rows.Close()

// 	var users User

//		for rows.Next() {
//			err := rows.Scan(&users.No, &users.Nick, &users.Name)
//			if err != nil {
//				return User{}, err
//			}
//		}
//		if err := rows.Err(); err != nil {
//			return User{}, err
//		}
//		return users, nil
//	}

func (um *UserModel) Delete(no int) error {
	deleteQuery := "delete from users where no = ?"
	_, err := um.db.Exec(deleteQuery, no)
	return err
}

func (um *UserModel) Update(user *User) error {
	updateQuery := "update users set nick =?  where name = ?"
	_, err := um.db.Exec(updateQuery, user.Nick, user.Name)
	return err
}
