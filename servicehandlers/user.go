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
	methodRouter(p, w, r)
}

// Get : Restful Get Method
func (p UserHandler) Get(r *http.Request) (string, int) {
	users := dao.GetUserByEmail("sufiyan@gmail.com")
	fmt.Println(users)
	return "GET Called", 200
}

// Put : Restful Put Method
func (p UserHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : Restful Post Method
func (p UserHandler) Post(r *http.Request) (string, int) {
	return "POST Called", 200
}
