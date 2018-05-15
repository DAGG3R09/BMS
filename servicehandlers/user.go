package servicehandlers

import (
	"bookmyshow/dao"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UserHandler : Restful UserHandler Method
type UserHandler struct {
}

func (p UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

// Get : Restful Get Method
func (p UserHandler) Get(r *http.Request) SrvcRes {
	users := dao.GetUserByEmail("sufiyan@gmail.com")
	fmt.Println(users)
	return Response200OK("Users: ok")
}

// Put : Restful Put Method
func (p UserHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Post : Restful Post Method - User Signup
func (p UserHandler) Post(r *http.Request) SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var cred dao.User
	err = json.Unmarshal(body, &cred)
	res := dao.CreateUser(cred)
	if !res {
		return SimpleBadRequest("User Already Present!")
	} else {
		return Response200OK(res)
	}
}
