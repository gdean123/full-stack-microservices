package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type StaticFileHandler struct {
	staticFilePath string
}

func NewStaticFileHandler(staticFilePath string) StaticFileHandler {
	return StaticFileHandler{
		staticFilePath: staticFilePath,
	}
}

func (s StaticFileHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := ioutil.ReadFile(filepath.Join(s.staticFilePath, request.URL.String()))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Write(data)
}
