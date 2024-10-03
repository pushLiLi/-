package scheduler

import (
	"awesomeProject4/scheduler/dbops"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		sendResponse(400, w, "video id should not be null")
		return
	}

	if err := dbops.AddVideoDeletionRecord(vid); err != nil {
		sendResponse(500, w, "Internal server error")
		return
	}

	sendResponse(200, w, "ok")
	return
}
