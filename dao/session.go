package dao

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Session : Structure of a session in SessionDB
type Session struct {
	UserID    int
	Token     string
	LastLogin time.Time
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
func CreateSession(userID string) string {

	token := createUID()

	err := db.QueryRow("Insert into session (user_id, token) values ($1, $2)", userID, token)

	if err != nil {
		panic(err)
	}
	// return token
	return ""
}

func createUID() uuid.UUID {
	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return uid
	}
	return uid
}

// AuthenticateSession : Authenticates the token
func AuthenticateSession(token string) bool {
	var us Session
	err := db.QueryRow("Select * from Session where token=$1", token).Scan(&us.UserID, &us.Token)

	if err != nil {
		// panic(err)
		return false
	}
	return true
}
