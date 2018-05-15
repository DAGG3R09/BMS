package servicehandlers

import (
	"bookmyshow/dao"
	"fmt"
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

// Post : Restful Post Method
func (p UserHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
