package scheduler

import (
	"io"
	"net/http"
)

func sendResponse(sc int, w http.ResponseWriter, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
