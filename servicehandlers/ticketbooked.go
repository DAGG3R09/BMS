package servicehandlers

import (
	"net/http"
)

// TicketBookedHandler : Restful TicketBookedHandler Method
type TicketBookedHandler struct {
}

func (p TicketBookedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : Restful Get Method
func (p TicketBookedHandler) Get(r *http.Request) (string, int) {
	return "GET Called", 200
}

// Put : Restful Put Method
func (p TicketBookedHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : Restful Post Method
func (p TicketBookedHandler) Post(r *http.Request) (string, int) {
	return "POST Called", 200
}
