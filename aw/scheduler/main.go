package scheduler

import (
	"awesomeProject4/scheduler/taskrunner"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}

func main() {
	go taskrunner.Start()
	router := RegisterHandlers()
	http.ListenAndServe(":9001", router)
}
