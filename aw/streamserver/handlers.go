package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 解析模板文件
	t, err := template.ParseFiles("./streamserver/videos/a.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 执行模板并渲染到响应中
	if err := t.Execute(w, nil); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error when trying to get current directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)

	//// 获取当前文件的绝对路径，找到 "main.go" 所在的目录
	//_, filePath, _, _ := runtime.Caller(0)
	//baseDir := filepath.Dir(filePath) //获取main.go所属的文件夹
	//// 拼接出视频的绝对路径，加上扩展名 ".mp4"
	//vl := filepath.Join(baseDir, "videos", vid+".mp4")
	vl := VIDEO_DIR + vid + ".mp4"

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when trying to open file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func updateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	//defer r.Body.Close()
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too large")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil { //尝试打开文件
		log.Printf("Error when trying to open file: %v", err)
		sendErrorResponse(w, http.StatusBadRequest, "File is missing")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error when trying to read file: %v", err)
		sendErrorResponse(w, http.StatusBadRequest, "File is missing")
		return
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Error when trying to write file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")

}
