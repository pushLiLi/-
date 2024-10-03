package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func RegisterHandler() *httprouter.Router {
	r := httprouter.New() //创建新的router

	r.GET("/videos/:vid-id", streamHandler)

	r.POST("/update/:vid-id", updateHandler)

	r.GET("/testPage", testPageHandler)

	return r

}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() { //没有接收到tocken
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}
	m.r.ServeHTTP(w, r)     //启用中间件服务
	defer m.l.ReleaseConn() //释放channel
}

func main() {
	r := RegisterHandler()
	mh := NewMiddleWareHandler(r, 10)
	http.ListenAndServe(":9000", mh)
}
