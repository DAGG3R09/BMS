package dao

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Session : Structure of a session in SessionDB
type Session struct {
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	LastLogin time.Time `json:"login_time"`
}

// GetAllSession : Returns all the User Sessions
func GetAllSession() []Session {
	var Sessions []Session

	err := db.QueryRow("Select * from Session")

	if err != nil {
		panic(err)
	}

	return Sessions
}

//CreateSession : Creates a new session for a user
func CreateSession(email string) Session {

	token := createUID()
	us := GetSessionByEmail(email)

	var err error

	if (us == Session{}) {
		usr := GetUserByEmail(email)
		_, err = db.Exec("Insert into session (user_id, token, login_time) values ($1, $2, $3)", usr.ID, token, time.Now())
	} else {
		_, err = db.Exec("Update session set token=$1, login_time=$2 where user_id=$3", token, time.Now(), us.UserID)
	}

	if err != nil {
		log.Fatal("Error in CreateSession", err)
	}

	us = GetSessionByEmail(email)
	return us
}

func createUID() uuid.UUID {
	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return uid
	}
	return uid
}

// GetSessionByToken : Returns Session corresponding to token
func GetSessionByToken(token string) Session {
	var us Session
	err := db.QueryRow("Select * from Session where token=$1", token).Scan(&us.UserID, &us.Token, &us.LastLogin)

	if err == sql.ErrNoRows {
		return (Session{})
	} else if err != nil {
		log.Fatal("Error in SessionDB", err)
	}
	return us
}

// GetSessionByEmail : Returns Session corresponding to token
func GetSessionByEmail(email string) Session {

	var us Session

	query := `Select * from Session where user_id =
				(Select user_id from users where email=$1)`

	err := db.QueryRow(query, email).Scan(&us.UserID, &us.Token, &us.LastLogin)

	if err == sql.ErrNoRows {
		return (Session{})
	} else if err != nil {
		log.Fatal("Error in SessionDB", err)
	}
	return us
}

// DeleteSession : Deletes the Session
func DeleteSession(userID int) {
	query := "Delete from SESSION where user_id=$1"

	_, err := db.Exec(query, userID)

	if err != nil {
		log.Fatal("Error in DeleteSession")
	}
}

// AuthenticateSession : Authenticates the token
func AuthenticateSession(token string) bool {

	session := GetSessionByToken(token)
	if (session == Session{}) {
		return false
	}
	return true
}
