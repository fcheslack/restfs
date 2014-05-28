package main

import (
	"github.com/fcheslack/restfs"
	"github.com/toqueteos/webbrowser"
	"log"
)

func main() {
	var _ = webbrowser.Open

	rfs := &restfs.RestfsServer{}
	rfs.SetBaseDirectory("./fs/")

	//Launch browser in goroutine that will get executed once
	//the server starts listening and yields
	log.Println("Launching server")
	go func() {
		log.Println("launching browser")
		webbrowser.Open("http://localhost:8080/static/")
	}()

	rfs.ListenAndServe()
}
