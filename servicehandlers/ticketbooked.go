package servicehandlers

import (
	"bookmyshow/dao"
	"bookmyshow/utils"
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
	token := r.Header.Get("token")
	session := dao.GetSessionByToken(token)

	if (session == dao.Session{}) {
		return UnauthorizedAccess("Unauthorized Access")
	}

	tickets := dao.GetBookedTicketByUserID(session.UserID)

	return Response200OK(tickets)
}

// Put : Restful Put Method
func (p TicketBookedHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Post : Restful Post Method
func (p TicketBookedHandler) Post(r *http.Request) SrvcRes {

	token := r.Header.Get("token")

	session := dao.GetSessionByToken(token)

	if (session == dao.Session{}) {
		return UnauthorizedAccess("Unauthorized Access")
	}

	q := r.URL.Query()
	tickets := q.Get("tickets")
	movieName := q.Get("movie_name")
	ntickets := utils.StringToInteger(tickets)

	for ntickets != 0 {
		dao.CreateBookedTicket(0, 1, session.UserID, movieName)
		ntickets--
	}

	return Response200OK("Tickets Booked")
}
