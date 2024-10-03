package main

import (
	"awesomeProject4/api/defs"
	"awesomeProject4/api/session"
	"net/http"
)

var HEADER_FILE_SESSION = "X-Session-Id"
var HEADER_FILE_USERNAME = "X-Session_Id"

// 身份验证
func verifySession(r *http.Request) bool {
	sessionId := r.Header.Get(HEADER_FILE_SESSION)
	if len(sessionId) == 0 {
		return false
	}
	userName, ok := session.IsSessionExpired(sessionId)
	if !ok {
		//不过期
		r.Header.Add(HEADER_FILE_USERNAME, userName)
		return true
	} else {
		return false
	}
}

func verifyUser(w http.ResponseWriter, r *http.Request) bool {
	userName := r.Header.Get(HEADER_FILE_USERNAME)
	if len(userName) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
