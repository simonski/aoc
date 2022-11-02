package api

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	app "github.com/simonski/aoc/app"
	"github.com/simonski/aoc/app/constants"
	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
)

//go:embed css visualisations api index.html
var staticFiles embed.FS
var SERVER *Server

// https://tutorialedge.net/golang/creating-restful-api-with-golang/
type Server struct {
	cli         *cli.CLI
	application *app.AOC
}

func NewServer(c *cli.CLI, application *app.AOC) *Server {
	server := &Server{cli: c, application: application}
	SERVER = server
	return server
}

func (server *Server) Run() {
	port := server.cli.GetIntOrDefault("-p", 8000)
	portstr := fmt.Sprintf(":%v", port)

	var staticFS http.FileSystem
	if server.cli.GetStringOrDefault("-fs", "") == "" {
		staticFS = http.FS(staticFiles)
		fmt.Print("Serving files from memory.\n")
	} else {
		root := server.cli.GetStringOrDefault("-fs", "")
		staticFS = http.Dir(root)
		fmt.Printf("Serving files from filesystem '%v'.\n", root)
	}
	fs := http.FileServer(staticFS)

	http.Handle("/", fs)

	// myRouter := mux.NewRouter().StrictSlash(true)
	// http.HandleFunc("/", indexFunc)
	http.HandleFunc("/blog", blogFunc)
	http.HandleFunc("/attempts", attemptsFunc)
	// http.HandleFunc("/api/solutions", apiSolutionsFunc)
	// http.HandleFunc("/api/summary", apiSummaryFunc)
	http.HandleFunc("/api/2021/05", api202105)
	http.HandleFunc("/api/2021/09", api202109)
	http.HandleFunc("/api/2021/11", api202111)
	// myRouter.HandleFunc("/style.css", cssFunc)
	// myRouter.HandleFunc("/code.js", jsFunc)
	// myRouter.HandleFunc("/tachyons.css", tachyonsFunc)

	fmt.Printf("AOC Server listening on %v\n", portstr)
	log.Fatal(http.ListenAndServe(portstr, nil))
}

// returns a list of days with solutions so we can then render what we want
// to look at

type Solution struct {
	Part1Solution bool `json:"c1"`
	Part2Solution bool `json:"c2"`
	Part1Api      bool `json:"a1"`
	Part2Api      bool `json:"a2"`
}

type Progress struct {
	Solutions map[string]*Solution `json:"solutions"`
	YearStart int                  `json:"start"`
	YearEnd   int                  `json:"end"`
}

func apiSolutionsFunc(w http.ResponseWriter, r *http.Request) {

	progress := Progress{YearStart: 2015, YearEnd: 2021}
	solutions := make(map[string]*Solution)

	a := SERVER.application

	for year := constants.MIN_YEAR; year <= constants.MAX_YEAR; year++ {
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

			s := &Solution{Part1Solution: m1exists, Part2Solution: m2exists, Part1Api: m1existsApi, Part2Api: m2existsApi}
			key := fmt.Sprintf("%v.%v", year, day)
			solutions[key] = s
		}
	}

	progress.Solutions = solutions
	msgb, _ := json.Marshal(progress)
	msg := string(msgb)
	length_str := fmt.Sprintf("%v", len(msg))
	w.Header().Set("Content-Type", "application/json") // this
	w.Header().Set("Content-Length", length_str)       // this
	fmt.Fprint(w, msg)
}

func apiSummaryFunc(w http.ResponseWriter, r *http.Request) {

	year := r.URL.Query().Get("year")
	day := r.URL.Query().Get("day")
	a := SERVER.application

	// splits := strings.Split(r.URL.Path, "/")
	// year := splits[len(splits)-2]
	// day := splits[len(splits)-3]
	iyear, _ := strconv.Atoi(year)
	iday, _ := strconv.Atoi(day)
	summary := a.GetSummary(iyear, iday)

	msgb, _ := json.MarshalIndent(summary, "", "   ")
	msg := string(msgb) + "\n"
	length_str := fmt.Sprintf("%v", len(msg))
	w.Header().Set("Content-Type", "application/json") // this
	w.Header().Set("Content-Length", length_str)       // this
	fmt.Fprint(w, msg)
}

func blogFunc(w http.ResponseWriter, r *http.Request) {
	// msg := "<!DOCTYPE html>\n<!--\nHi!\n\nThis is my Rube Goldberg Advent of Code visualisations attempt.\n\nThere isn't anything to see here yet - but there is an api at /api/solutions \n-->\n<html>\n\t<head>\n\t\t<title>AOC</title>\n\t</head>\n\t<body>AOC {YEAR} <a href='/api/solutions'>[solutions]</a></body>\n</html>"
	a := SERVER.application
	// sort by lastmodified, then by year/day in reverse
	summaries := a.GetSummaries()
	// s := make([]*utils.Summary, 0)
	// for _, v := range summaries {
	// 	s = append(s, v)
	// }

	sort.Slice(summaries, func(i1 int, i2 int) bool {
		s1 := summaries[i1]
		s2 := summaries[i2]
		if s1 == nil {
			return false
		}
		if s2 == nil {
			return true
		}
		// fmt.Printf("i1=%v, i2=%v, s1=%v, s2=%v\n", i1, i2, s1, s2)
		if s1.LastModified > s2.LastModified {
			return true
		}
		s1Key := fmt.Sprintf("%v%02d", s1.Year, s1.Day)
		s2Key := fmt.Sprintf("%v%02d", s2.Year, s2.Day)
		return s1Key > s2Key
	})

	lines := ""
	for _, v := range summaries {
		if v == nil {
			continue
		}
		ticks := ""
		if v.ProgressP1 == utils.Completed && v.ProgressP2 == utils.Completed {
			ticks = "&#9733"
		} else if v.ProgressP1 == utils.Completed || v.ProgressP2 == utils.Completed {
			ticks = "&#9734;"
		} else if v.ProgressP1 == utils.Started || v.ProgressP2 == utils.Started {
			ticks = "&#9775;"
		} else {
			ticks = "&#9776;"
		}
		href := fmt.Sprintf("https://adventofcode.com/%v/day/%v", v.Year, v.Day)
		line := fmt.Sprintf("<li>%v %v %02d <a href='%v'>%v</a></li>\n", ticks, v.Year, v.Day, href, v.Name)
		lines += line
	}

	// msg = strings.ReplaceAll(msg, "{YEAR}", fmt.Sprintf("%v", time.Now().Year()))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, lines)
}

func attemptsFunc(w http.ResponseWriter, r *http.Request) {
	// msg := "<!DOCTYPE html>\n<!--\nHi!\n\nThis is my Rube Goldberg Advent of Code visualisations attempt.\n\nThere isn't anything to see here yet - but there is an api at /api/solutions \n-->\n<html>\n\t<head>\n\t\t<title>AOC</title>\n\t</head>\n\t<body>AOC {YEAR} <a href='/api/solutions'>[solutions]</a></body>\n</html>"
	a := SERVER.application
	// sort by lastmodified, then by year/day in reverse
	summaries := a.GetSummaries()
	// s := make([]*utils.Summary, 0)
	// for _, v := range summaries {
	// 	s = append(s, v)
	// }

	sort.Slice(summaries, func(i1 int, i2 int) bool {
		s1 := summaries[i1]
		s2 := summaries[i2]
		if s1 == nil {
			return false
		}
		if s2 == nil {
			return true
		}
		// fmt.Printf("i1=%v, i2=%v, s1=%v, s2=%v\n", i1, i2, s1, s2)
		if s1.LastModified > s2.LastModified {
			return true
		}
		s1Key := fmt.Sprintf("%v%02d", s1.Year, s1.Day)
		s2Key := fmt.Sprintf("%v%02d", s2.Year, s2.Day)
		return s1Key > s2Key
	})

	lines := ""
	for _, v := range summaries {
		if v == nil {
			continue
		}
		ticks := ""
		if v.ProgressP1 == utils.Completed && v.ProgressP2 == utils.Completed {
			ticks = "&#9733"
		} else if v.ProgressP1 == utils.Completed || v.ProgressP2 == utils.Completed {
			ticks = "&#9734;"
		} else if v.ProgressP1 == utils.Started || v.ProgressP2 == utils.Started {
			ticks = "&#9775;"
		} else {
			ticks = "&#9776;"
		}
		href := fmt.Sprintf("https://adventofcode.com/%v/day/%v", v.Year, v.Day)
		line := fmt.Sprintf("<li>%v %v %02d <a href='%v'>%v</a></li>\n", ticks, v.Year, v.Day, href, v.Name)
		lines += line
	}

	// msg = strings.ReplaceAll(msg, "{YEAR}", fmt.Sprintf("%v", time.Now().Year()))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, lines)
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	msg := "<!DOCTYPE html>\n<!--\nHi!\n\nThis is my Rube Goldberg Advent of Code visualisations attempt.\n\nThere isn't anything to see here yet - but there is an api at /api/solutions \n-->\n<html>\n\t<head>\n\t\t<title>AOC</title>\n\t</head>\n\t<body>AOC {YEAR} <a href='/api/solutions'>[solutions]</a></body>\n</html>"

	msg = strings.ReplaceAll(msg, "{YEAR}", fmt.Sprintf("%v", time.Now().Year()))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, msg)
}

func api202105(w http.ResponseWriter, r *http.Request) {
	a := SERVER.application
	appLogic := a.GetAppLogic(2021)
	response := appLogic.Api(5)
	// msgb, _ := json.Marshal(response)
	// msg := string(msgb)
	length_str := fmt.Sprintf("%v", len(response))
	w.Header().Set("Content-Type", "application/json") // this
	w.Header().Set("Content-Length", length_str)       // this
	fmt.Fprint(w, response)
}

func api202109(w http.ResponseWriter, r *http.Request) {
	a := SERVER.application
	appLogic := a.GetAppLogic(2021)
	response := appLogic.Api(9)
	// msgb, _ := json.Marshal(response)
	// msg := string(msgb)
	length_str := fmt.Sprintf("%v", len(response))
	w.Header().Set("Content-Type", "application/json") // this
	w.Header().Set("Content-Length", length_str)       // this
	fmt.Fprint(w, response)
}

func api202111(w http.ResponseWriter, r *http.Request) {
	a := SERVER.application
	appLogic := a.GetAppLogic(2021)
	response := appLogic.Api(11)
	// msgb, _ := json.Marshal(response)
	// msg := string(msgb)
	length_str := fmt.Sprintf("%v", len(response))
	w.Header().Set("Content-Type", "application/json") // this
	w.Header().Set("Content-Length", length_str)       // this
	fmt.Fprint(w, response)
}
