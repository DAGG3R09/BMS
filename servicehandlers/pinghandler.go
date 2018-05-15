package servicehandlers

import (
	"net/http"
)

// PingHandler : Restful PingHandler Method
type PingHandler struct {
}

func (p PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : Restful Get Method
func (p PingHandler) Get(r *http.Request) (string, int) {
	return "GET Called", 200
}

// Put : Restful Put Method
func (p PingHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : Restful Post Method
func (p PingHandler) Post(r *http.Request) (string, int) {
	return "POST Called", 200
}
