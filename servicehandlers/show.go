package servicehandlers

import (
	"net/http"
)

// ShowHandler : Restful ShowHandler Method
type ShowHandler struct {
}

func (p ShowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : Restful Get Method
func (p ShowHandler) Get(r *http.Request) (string, int) {
	return "GET Called", 200
}

// Put : Restful Put Method
func (p ShowHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : Restful Post Handler
func (p ShowHandler) Post(r *http.Request) (string, int) {
	return "POST Called", 200
}
