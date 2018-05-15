package servicehandlers

import (
	"net/http"
)

// TicketHandler : Restful TicketHandler Method
type TicketHandler struct {
}

func (p TicketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

// Get : Restful Get Method
func (p TicketHandler) Get(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Put : Restful Put Method
func (p TicketHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Post : Restful Post Method
func (p TicketHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
