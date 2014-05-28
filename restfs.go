package restfs

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type RestfsServer struct {
	baseFileSystemPath string
}

//set base directory
func (rfs *RestfsServer) SetBaseDirectory(s string) {
	var err error
	rfs.baseFileSystemPath, err = filepath.Abs(s)
	if err != nil {
		log.Fatal("Error getting absolute path for base fs directory")
	}
}

//get base directory
func (rfs *RestfsServer) GetBaseDirectory() string {
	return rfs.baseFileSystemPath
}

//get full file path given fs path
func (rfs *RestfsServer) FullFilepath(filename string) string {
	fullPath := filepath.Clean(filepath.Join(rfs.baseFileSystemPath, filename))
	return fullPath
}

//GET to read file
func (rfs *RestfsServer) ReadFileHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("ReadFileHandler")
	vars := mux.Vars(request)
	fullFilename := rfs.FullFilepath(vars["filename"])
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
	fullFilename := rfs.FullFilepath(vars["filename"])
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
func (rfs *RestfsServer) DeleteFileHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("DeleteFileHandler")
	vars := mux.Vars(request)
	fullFilename := rfs.FullFilepath(vars["filename"])
	log.Println(fullFilename)
	log.Printf("Delete %s\n", fullFilename)
	//os.Remove(fullFilename)
}
func DeleteFile(filename string) error {
	return nil
}

func (rfs *RestfsServer) ListenAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/fs/{filename:.*}", rfs.ReadFileHandler).Methods("GET")
	r.HandleFunc("/fs/{filename:.*}", rfs.WriteFileHandler).Methods("POST", "PUT")
	r.HandleFunc("/fs/{filename:.*}", rfs.DeleteFileHandler).Methods("DELETE")
	http.Handle("/", r)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", nil)
	log.Println("Listening")
}
