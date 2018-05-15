package dao

import (
	"bookmyshow/utils"
	"fmt"
	"log"
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

	if err != nil {
		fmt.Println(">>", err)
		return User{}
	}

	return user
}

//AuthenticateUser : Compares the entered password with pass of user
func AuthenticateUser(email string, password string) bool {
	user := GetUserByEmail(email)

	if utils.GetMD5Hash(password) == user.Password {
		return true
	}
	return false
}

// CreateUser : Creates a user in database
func CreateUser(email string, name string, password string, phone string) {

	pass := utils.GetMD5Hash(password)
	err := db.QueryRow("Insert into users (email, name, pwd) values ($1, $2, $3, $4)", email, name, pass, phone)

	if err != nil {
		log.Fatal("Error in insert", err)
	}
}
