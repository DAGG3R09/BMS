package main

import (
	"bookmyshow/dao"
	"bookmyshow/servicehandlers"
	"log"
	"net/http"
)

func main() {

	// Serves the html pages
	http.Handle("/", http.FileServer(http.Dir("./html")))

	p := servicehandlers.PingHandler{}
	u := servicehandlers.UserHandler{}
	m := servicehandlers.MovieHandler{}
	a := servicehandlers.AuthenticateHandler{}

	dao.InitiateDatabaseConnection()
	defer dao.CloseDatabaseConnection()

	// users := dao.GetAllUsers()
	// fmt.Println(users)
	// user := dao.GetUserByEmail("sufiyana@gmail.com")
	// fmt.Println(user)
	// fmt.Println(dao.AuthenticateUser("sufiyan@gmail.com", "pass"))

	http.Handle("/bms/ping/", p)
	http.Handle("/bms/users/", u)
	http.Handle("/bms/authenticate/", a)
	http.Handle("/bms/movies/", m)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
