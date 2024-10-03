package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	//这是一个interface{}
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	//上面的interface{}的一个实现
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//表明实现的是middleWareHandler接口，
	//check session
	verifySession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()             //申请了一个http的server，包括下面对http的一些配置
	router.POST("/user", CreateUser)       //用户注册
	router.POST("/user/:user_name", Login) //用户登录

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe("localhost:8080", mh)
}
