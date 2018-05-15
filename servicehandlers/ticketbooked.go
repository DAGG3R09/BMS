package servicehandlers

import (
	"net/http"
)

// TicketBookedHandler : Restful TicketBookedHandler Method
type TicketBookedHandler struct {
}

func (p TicketBookedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

// Get : Restful Get Method
func (p TicketBookedHandler) Get(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Put : Restful Put Method
func (p TicketBookedHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Post : Restful Post Method
func (p TicketBookedHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
