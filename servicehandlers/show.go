package servicehandlers

import (
	"net/http"
)

// ShowHandler : Restful ShowHandler Method
type ShowHandler struct {
}

func (p ShowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

// Get : Restful Get Method
func (p ShowHandler) Get(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Put : Restful Put Method
func (p ShowHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Post : Restful Post Handler
func (p ShowHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
