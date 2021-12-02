package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/simonski/goutils"
)

// https://tutorialedge.net/golang/creating-restful-api-with-golang/
type Server struct {
	cli *goutils.CLI
}

func NewServer(c *goutils.CLI) *Server {
	server := &Server{cli: c}
	return server
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AOC 2021 Time!")
	// fmt.Println("Endpoint hit: homepage")

	// where Articles is a []Article
	// json.NewEncoder(w).Encode(Articles)

}

func (server *Server) Run() {
	port := server.cli.GetIntOrDefault("-p", 8000)
	portstr := fmt.Sprintf(":%v", port)
	http.HandleFunc("/", homePage)
	fmt.Printf("AOC Server listening on %v\n", portstr)
	log.Fatal(http.ListenAndServe(portstr, nil))
}
