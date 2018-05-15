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
	s := servicehandlers.ShowHandler{}
	b := servicehandlers.TicketBookedHandler{}

	dao.InitiateDatabaseConnection()
	defer dao.CloseDatabaseConnection()

	http.Handle("/bms/ping/", p)
	http.Handle("/bms/users/", u)
	http.Handle("/bms/authenticate/", a)
	http.Handle("/bms/movies/", m)
	http.Handle("/bms/shows/", s)
	http.Handle("/bms/booktickets/", b)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
