package servicehandlers

import (
	"bookmyshow/dao"
	"encoding/json"
	"net/http"
)

// AuthenticateHandler : Restful Authenticate Handler
type AuthenticateHandler struct {
}

func (p AuthenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

// Get : Restful Get Method
func (p AuthenticateHandler) Get(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Put : Restful PUT Method
func (p AuthenticateHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Post : Restful Post Method
func (p AuthenticateHandler) Post(r *http.Request) SrvcRes {

	type authenticateBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var payload authenticateBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
		panic(err)
	}

	user := dao.GetUserByEmail(payload.Email)
	if (user == dao.User{}) {
		return SimpleBadRequest("User doesn't Exist")
	}

	if !dao.AuthenticateUser(payload.Email, payload.Password) {
		return SimpleBadRequest("Invalid Email or Password")
	}

	sesion := dao.CreateSession(payload.Email)

	return Response200OK(sesion)
}
