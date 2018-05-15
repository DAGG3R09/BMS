package servicehandlers

import (
	"net/http"
)

// PingHandler : Restful PingHandler Method
type PingHandler struct {
}

func (p PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

// Get : Restful Get Method
func (p PingHandler) Get(r *http.Request) SrvcRes {
	return Response200OK("GET CALLED!")
}

// Put : Restful Put Method
func (p PingHandler) Put(r *http.Request) SrvcRes {
	return Response200OK("PUT Called")
}

// Post : Restful Post Method
func (p PingHandler) Post(r *http.Request) SrvcRes {
	return Response200OK("POST Called")
}
