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
		metrics.CounterAdd(login, 1, map[string]string{label: fmt.Sprintf("%s login failed , %[1]s repeat login", u.UserName)})
		ctx.Write([]byte(u.UserName + "  aAlready landing , do not login agent"))
		return
	}
	userLogin[u.UserName] = u
	metrics.CounterAdd(login, 1, map[string]string{label: fmt.Sprintf("%s login", u.UserName)})
	b, _ := json.Marshal(u)
	ctx.Write(b)
}

func (*User) SignOut(ctx *restful.Context) {
	user := ctx.ReadQueryParameter("user_name")
	if _, ok := userLogin[user]; !ok {
		metrics.CounterAdd(signOut, 1, map[string]string{label: fmt.Sprintf("%s sign out failed , %[1]s not login", user)})
		ctx.Write([]byte(user + " not login , did not need sign out"))
		return
	}
	delete(userLogin, user)
	metrics.CounterAdd(signOut, 1, map[string]string{label: fmt.Sprintf("%s sign out", user)})
	ctx.Write([]byte(user + "  sing out success"))
}

//URLPatterns helps to respond for corresponding API calls
func (*User) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodPost, Path: "/login", ResourceFuncName: "Login"},
		{Method: http.MethodGet, Path: "/sign_out", ResourceFuncName: "SignOut"},
	}
}

const (
	login   = "login_total"
	signOut = "sign_out_total"
	label   = "username"
)

func init() {
	metrics.Init()
	metrics.CreateCounter(metrics.CounterOpts{
		Help:   "count user login",
		Name:   login,
		Labels: []string{label}})
	metrics.CreateCounter(metrics.CounterOpts{
		Help:   "user sign out",
		Name:   signOut,
		Labels: []string{label}})
}
