package servicehandlers

import (
	"net/http"
)

// MultiplexHandler : Restful Multiplex Handler
type MultiplexHandler struct {
}

func (p MultiplexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

// Get : Restful Get Method
func (p MultiplexHandler) Get(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Put : Restful Put Handler
func (p MultiplexHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Post : Restful Post Handler
func (p MultiplexHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
