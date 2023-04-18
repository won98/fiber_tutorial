package models

import (
	"database/sql"
	"fmt"
	"gotest/utils"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Nick string `json:"nick"`
	No   int    `json:"no"`
	Id   string `json:"id"`
	Pass string `json:"pass"`
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
	hashPassword, err := utils.Hashing(user.Pass)
	if err != nil {
		return err
	}
	insertQuery := "insert into users (name, nick, id, password) values(?,?,?,?)"
	_, err2 := um.db.Exec(insertQuery, user.Name, user.Nick, user.Id, hashPassword)
	return err2
}

// 배열에 속에 방법
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
		fmt.Println(users)
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
	updateQuery := "update golang.users set nick =?  where id = ?"
	_, err := um.db.Exec(updateQuery, user.Nick, user.Id)
	return err
}

// func (um *UserModel) Login(id string, pass string) (string, error) {
// 	loginQuery := "select no,nick,name,password from golang.users where id = ? and password = ?"
// 	rows, err := um.db.Query(loginQuery, id, pass)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	defer rows.Close()
// 	users := make([]User, 0)

//		for rows.Next() {
//			var user User
//			err := rows.Scan(&user.No, &user.Nick, &user.Name, &user.Pass)
//			if err != nil {
//				return "", err
//			}
//			users = append(users, user)
//			fmt.Println(users)
//		}
//		if len(users) == 0 {
//			return "", fmt.Errorf("invalid username or password")
//		}
//		compare := utils.CompareHash(users[0].Pass, pass)
//		if compare != nil {
//			return "", compare
//		}
//		if err := rows.Err(); err != nil {
//			return "", err
//		}
//		token, err := utils.CreateToken(id)
//		if err != nil {
//			return "", err
//		}
//		return token, nil
//	}
func (um *UserModel) Login(id string, pass string) (string, string, error) {
	loginQuery := "SELECT no, nick, name, password FROM golang.users WHERE id = ?"
	rows, err := um.db.Query(loginQuery, id)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	defer rows.Close()
	users := make([]User, 0)

	for rows.Next() {
		var user User
		err := rows.Scan(&user.No, &user.Nick, &user.Name, &user.Pass)
		if err != nil {
			return "", "", err
		}
		users = append(users, user)
		fmt.Println(users)
	}

	if len(users) == 0 {
		return "", "", fmt.Errorf("invalid username or password")
	}

	if compare := utils.CompareHash(users[0].Pass, pass); err != nil {
		return "", "", compare
	}

	token, err := utils.CreateToken(id, users[0].No)
	if err != nil {
		return "", "", err
	}

	refreshtoken, err := utils.CreateRefreshToken(id)
	if err != nil {
		return "", "", err
	}

	return token, refreshtoken, nil
}
