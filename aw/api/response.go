package main

import (
	"awesomeProject4/api/defs"
	"encoding/json"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errRes defs.ErrorResponse) {
	w.WriteHeader(errRes.HttpStatusCode) //输出状态码，这里拿到的实际是err的一个指针

	resStr, _ := json.Marshal(errRes)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, res string, statusCode int) {
	w.WriteHeader(statusCode)
	io.WriteString(w, res)
}
