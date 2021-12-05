package api

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	app "github.com/simonski/aoc/app"
	"github.com/simonski/goutils"

	_ "embed"
)

//go:embed css js index.html
var staticFiles embed.FS
var SERVER *Server

// https://tutorialedge.net/golang/creating-restful-api-with-golang/
type Server struct {
	cli         *goutils.CLI
	application *app.AOCApplication
}

func NewServer(c *goutils.CLI, application *app.AOCApplication) *Server {
	server := &Server{cli: c, application: application}
	SERVER = server
	return server
}

func (server *Server) Run() {
	port := server.cli.GetIntOrDefault("-p", 8000)
	portstr := fmt.Sprintf(":%v", port)

	var staticFS = http.FS(staticFiles)
	fs := http.FileServer(staticFS)

	http.Handle("/", fs)

	// myRouter := mux.NewRouter().StrictSlash(true)
	http.HandleFunc("/api/solutions", apiSolutionsFunc)
	// myRouter.HandleFunc("/style.css", cssFunc)
	// myRouter.HandleFunc("/code.js", jsFunc)
	// myRouter.HandleFunc("/tachyons.css", tachyonsFunc)

	fmt.Printf("AOC Server listening on %v\n", portstr)
	log.Fatal(http.ListenAndServe(portstr, nil))
}

// returns a list of days with solutions so we can then render what we want
// to look at

type Solution struct {
	Year          int  `json:"year"`
	Day           int  `json:"day"`
	Part1Solution bool `json:"part1Solution"`
	Part2Solution bool `json:"part2Solution"`
	Part1Api      bool `json:"part1Api"`
	Part2Api      bool `json:"part2Api"`
}

func apiSolutionsFunc(w http.ResponseWriter, r *http.Request) {
	solutions := make([]*Solution, 0)

	a := SERVER.application

	for year := 2015; year <= 2021; year++ {
		appLogic := a.GetAppLogic(year)
		for day := 1; day <= 25; day++ {
			methodNamePart1 := fmt.Sprintf("Y%vD%02dP1", year, day)
			methodNamePart2 := fmt.Sprintf("Y%vD%02dP2", year, day)
			methodNamePart1Api := fmt.Sprintf("Y%vD%02dP1Api", year, day)
			methodNamePart2Api := fmt.Sprintf("Y%vD%02dP2Api", year, day)

			_, _, m1exists := appLogic.GetMethod(methodNamePart1)
			_, _, m2exists := appLogic.GetMethod(methodNamePart2)
			_, _, m1existsApi := appLogic.GetMethod(methodNamePart1Api)
			_, _, m2existsApi := appLogic.GetMethod(methodNamePart2Api)

			s := &Solution{Year: year, Day: day, Part1Solution: m1exists, Part2Solution: m2exists, Part1Api: m1existsApi, Part2Api: m2existsApi}
			solutions = append(solutions, s)
		}
	}
	msgb, _ := json.Marshal(solutions)
	msg := string(msgb)
	length_str := fmt.Sprintf("%v", len(msg))
	w.Header().Set("Content-Type", "application/json") // this
	w.Header().Set("Content-Length", length_str)       // this
	fmt.Fprint(w, msg)
}
