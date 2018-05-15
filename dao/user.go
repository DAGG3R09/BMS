package dao

import (
	"bookmyshow/utils"
	"database/sql"
	"fmt"
)

//User : The User Table schema
type User struct {
	ID       int
	Email    string
	Name     string
	Phone    string
	Password string
}

//GetAllUsers return all users from users table
func GetAllUsers() []User {
	var users []User

	rows, err := db.Query("Select * from users")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Phone); err != nil {
			panic(err)
			// log.Fatal(err)
		}
		users = append(users, user)

	}
	return users
}

// GetUserByEmail : returns user by email
func GetUserByEmail(email string) User {
	var user User

	err := db.QueryRow("Select * from users where email=$1", email).Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Phone)

	if err == sql.ErrNoRows {
		return User{}
	} else if err != nil {
		fmt.Println("GetUserByEmail >>", err)
		return User{}
	}

	return user
}

//AuthenticateUser : Compares the entered password with pass of user
func AuthenticateUser(email string, password string) bool {
	user := GetUserByEmail(email)
	pass := utils.GetMD5Hash(password)

	if pass == user.Password {
		return true
	}
	return false
}

// CreateUser : Creates a user in database
func CreateUser(user User) bool {

	u := GetUserByEmail(user.Email)

	if (u != User{}) {
		return false
	}

	pass := utils.GetMD5Hash(user.Password)
	db.QueryRow("Insert into users (email, name, pwd, phone) values ($1, $2, $3, $4)", user.Email, user.Name, pass, user.Phone)

	// if err != nil {
	// log.Fatal("Error in insert", err)
	// 	return "Internal Server Error: Error in User Insert"
	// }
	return true
}
