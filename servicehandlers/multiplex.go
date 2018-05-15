package servicehandlers

import (
	"net/http"
)

// MultiplexHandler : Restful Multiplex Handler
type MultiplexHandler struct {
}

func (p MultiplexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : Restful Get Method
func (p MultiplexHandler) Get(r *http.Request) (string, int) {
	return "GET Called", 200
}

// Put : Restful Put Handler
func (p MultiplexHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : Restful Post Handler
func (p MultiplexHandler) Post(r *http.Request) (string, int) {
	return "POST Called", 200
}
