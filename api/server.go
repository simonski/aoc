package api

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/simonski/goutils"

	_ "embed"
)

//go:embed css js index.html
var staticFiles embed.FS

// https://tutorialedge.net/golang/creating-restful-api-with-golang/
type Server struct {
	cli *goutils.CLI
}

func NewServer(c *goutils.CLI) *Server {
	server := &Server{cli: c}
	return server
}

func (server *Server) Run() {
	port := server.cli.GetIntOrDefault("-p", 8000)
	portstr := fmt.Sprintf(":%v", port)

	var staticFS = http.FS(staticFiles)
	fs := http.FileServer(staticFS)

	http.Handle("/", fs)

	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", htmlFunc)
	// myRouter.HandleFunc("/style.css", cssFunc)
	// myRouter.HandleFunc("/code.js", jsFunc)
	// myRouter.HandleFunc("/tachyons.css", tachyonsFunc)

	fmt.Printf("AOC Server listening on %v\n", portstr)
	log.Fatal(http.ListenAndServe(portstr, nil))
}
