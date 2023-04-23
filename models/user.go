package models

import (
	"fmt"
	"gotest/utils"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	No       int    `json:"no"`
	Id       string `json:"id"`
	Password string `json:"pass"`
}

type UserModel struct {
	//db *sql.DB
	db *gorm.DB
}

// row query
//
//	func NewUserModel(db *sql.DB) *UserModel {
//		return &UserModel{
//			db: db,
//		}
//	}
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (um *UserModel) Insert(user *User, userid, name, nick, pass string) error {
	u := new(User)
	hashPassword, err := utils.Hashing(pass)
	if err != nil {
		return err
	}
	//pass = string(hashPassword)
	//insertQuery := "insert into users (name, nick, id, password) values(?,?,?,?)"
	u.Id = userid
	u.Name = name
	u.Nick = nick
	u.Password = string(hashPassword)
	err = um.db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

// 배열에 속에 방법
func (um *UserModel) Select(name string) ([]User, error) {
	//selectQuery := "select no,nick,name from golang.users where name = ?"
	var users []User
	err := um.db.Where("name = ?", name).Find(&users).Error
	if err != nil {

		fmt.Println(err)
		return nil, err
	}
	//defer rows.Close()
	//users := make([]User, 0)

	// for rows.Next() {
	// 	var user User
	// 	err := rows.Scan(&user.No, &user.Nick, &user.Name)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, user)
	// 	fmt.Println(users)
	// }
	// if err := rows.Err(); err != nil {
	// 	return nil, err
	// }
	return users, nil
}

/* 배열 없이하는 방법
func (um *UserModel) Select(name string) (User, error) {
	selectQuery := "select no,nick,name from golang.users where name = ?"
	rows, err := um.db.Query(selectQuery, name)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	var users User

		for rows.Next() {
			err := rows.Scan(&users.No, &users.Nick, &users.Name)
			if err != nil {
				return User{}, err
			}
		}
		if err := rows.Err(); err != nil {
			return User{}, err
		}
		return users, nil
	}*/

func (um *UserModel) Delete(no int) error {
	//deleteQuery := "delete from users where no = ?"
	err := um.db.Where("no = ?", no).Delete(&User{}).Error
	return err
}

func (um *UserModel) Update(user *User) error {
	//updateQuery := "update golang.users set nick =?  where id = ?"
	err := um.db.Model(&user).Updates(User{Nick: user.Nick}).Error
	fmt.Println(err)
	return err
}

func (um *UserModel) Login(id string, pass string) (string, string, error) {
	//loginQuery := "SELECT no, nick, name, password FROM golang.users WHERE id = ?"
	var users []User
	err := um.db.Where("id = ?", id).First(&users).Error
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	// defer rows.Close()
	// users := make([]User, 0)

	// for rows.Next() {
	// 	var user User
	// 	err := rows.Scan(&user.No, &user.Nick, &user.Name, &user.Pass)
	// 	if err != nil {
	// 		return "", "", err
	// 	}
	// 	users = append(users, user)
	// 	fmt.Println(users)
	// }

	// if len(users) == 0 {
	// 	return "", "", fmt.Errorf("invalid username or password")
	// }

	// if compare := utils.CompareHash(users.Pass, pass); compare != nil {
	// 	return "", "", compare
	// }

	token, err := utils.CreateToken(id, int(users[0].No))
	if err != nil {
		return "", "", err
	}

	refreshtoken, err := utils.CreateRefreshToken(id)
	if err != nil {
		return "", "", err
	}

	return token, refreshtoken, nil
}
