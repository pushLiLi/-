package main

import (
	"awesomeProject4/api/dbops"
	"awesomeProject4/api/defs"
	"awesomeProject4/api/session"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	requestBody, err := io.ReadAll(r.Body) //获取body
	//解析body
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(requestBody))
	}
	userBody := &defs.UserCredential{}

	if err := json.Unmarshal(requestBody, userBody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	//添加到数据库中
	err = dbops.AddUserCredential(userBody.Username, userBody.Password)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBFailed)
		return
	}
	//产生新的sessionId
	sessionId := session.GenerateNewSessionId(userBody.Username)
	signedUp := &defs.SignedUp{Success: true, SessionId: sessionId}
	//session包装成json，输出
	if resp, err := json.Marshal(signedUp); err != nil {
		sendErrorResponse(w, defs.ErrorInternalError)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userName := p.ByName("user_name") //路径中含有参数
	fmt.Println("Login function called")
	io.WriteString(w, userName)
}
