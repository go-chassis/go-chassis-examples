package schema

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chassis/go-chassis/pkg/metrics"
	"github.com/go-chassis/go-chassis/server/restful"
)

type User struct {
	UserName string `json:"user_name"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Pwd      string `json:"-"`
}

var userLogin = make(map[string]*User)

func (u *User) Login(ctx *restful.Context) {
	ctx.ReadEntity(u)
	if _, ok := userLogin[u.UserName]; ok {
		metrics.CounterAdd(Login, 1, map[string]string{Label: fmt.Sprintf("%s Login failed , %[1]s repeat Login", u.UserName)})
		ctx.Write([]byte(u.UserName + "  aAlready landing , do not Login agent"))
		return
	}
	userLogin[u.UserName] = u
	metrics.CounterAdd(Login, 1, map[string]string{Label: fmt.Sprintf("%s Login", u.UserName)})
	b, _ := json.Marshal(u)
	ctx.Write(b)
}

func (*User) SignOut(ctx *restful.Context) {
	user := ctx.ReadQueryParameter("user_name")
	if _, ok := userLogin[user]; !ok {
		metrics.CounterAdd(SignOut, 1, map[string]string{Label: fmt.Sprintf("%s sign out failed , %[1]s not Login", user)})
		ctx.Write([]byte(user + " not Login , did not need sign out"))
		return
	}
	delete(userLogin, user)
	metrics.CounterAdd(SignOut, 1, map[string]string{Label: fmt.Sprintf("%s sign out", user)})
	ctx.Write([]byte(user + "  sing out success"))
}

//URLPatterns helps to respond for corresponding API calls
func (*User) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodPost, Path: "/Login", ResourceFuncName: "Login"},
		{Method: http.MethodGet, Path: "/sign_out", ResourceFuncName: "SignOut"},
	}
}

const (
	Login   = "login_total"
	SignOut = "sign_out_total"
	Label   = "username"
)

func init() {

}
