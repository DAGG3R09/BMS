package servicehandlers

import (
	"net/http"
)

// TicketHandler : Restful TicketHandler Method
type TicketHandler struct {
}

func (p TicketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : Restful Get Method
func (p TicketHandler) Get(r *http.Request) (string, int) {
	return "GET Called", 200
}

// Put : Restful Put Method
func (p TicketHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : Restful Post Method
func (p TicketHandler) Post(r *http.Request) (string, int) {
	return "POST Called", 200
}
