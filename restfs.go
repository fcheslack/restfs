package main

import (
	"github.com/gorilla/mux"
	"github.com/toqueteos/webbrowser"
	"io/ioutil"
	"log"
	"net/http"
	//	"time"
)

type RestfsServer struct {
	BaseFileSystemPath string
}

func main() {
	var _ = mux.NewRouter()
	var _ = webbrowser.Open

	rfs := &RestfsServer{"./fs/"}

	//Launch browser in goroutine that will get executed once
	//the server starts listening and yields
	log.Println("Launching server")
	go func() {
		log.Println("launching browser")
		webbrowser.Open("http://localhost:8080/static/")
	}()

	rfs.ListenAndServe()
}

//set base directory
func (rfs *RestfsServer) SetBaseDirectory(s string) {
	rfs.BaseFileSystemPath = s
}

//get base directory
func (rfs *RestfsServer) GetBaseDirectory() string {
	return rfs.BaseFileSystemPath
}

//GET to read file
func (rfs *RestfsServer) ReadFileHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("ReadFileHandler")
	vars := mux.Vars(request)
	fullFilename := rfs.BaseFileSystemPath + vars["filename"]
	log.Println(fullFilename)
	http.ServeFile(response, request, fullFilename)
}
func ReadFile(filename string) ([]byte, error) {
	return nil, nil
}

//POST to create file
//PUT to overwrite file
func (rfs *RestfsServer) WriteFileHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("WriteFileHandler")
	vars := mux.Vars(request)
	fullFilename := rfs.BaseFileSystemPath + vars["filename"]
	log.Println(fullFilename)
	body := request.Body
	bodyBytes, _ := ioutil.ReadAll(body)
	ioutil.WriteFile(fullFilename, bodyBytes, 0770)

	//http.ServeFile(response, request, filename)
}
func WriteFile(filename string, data []byte) error {
	return nil
}

//DELETE to delete file
func DeleteFile(filename string) error {
	return nil
}

func (rfs *RestfsServer) ListenAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/fs/{filename}", rfs.ReadFileHandler).Methods("GET")
	r.HandleFunc("/fs/{filename}", rfs.WriteFileHandler).Methods("POST", "PUT")
	http.Handle("/", r)

	//http.Handle("/", http.FileServer(http.Dir("/home/fcheslack/Dev/go/src/github.com/fcheslack/restfs")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", nil)
	log.Println("Listening")
}
