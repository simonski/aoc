package api

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	app "github.com/simonski/aoc/app"
	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
)

//go:embed css visualisations index.html
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
	http.HandleFunc("/rss", rssFunc)
	http.HandleFunc("/blog", blogFunc)
	http.HandleFunc("/attempts", attemptsFunc)

	http.HandleFunc("/api/2021/05", api202105)
	http.HandleFunc("/api/2021/09", api202109)
	http.HandleFunc("/api/2021/11", api202111)

	// http.HandleFunc("/api/solutions", apiSolutionsFunc)
	http.HandleFunc("/api/summary", apiSummaryFunc)

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

// func apiSolutionsFunc(w http.ResponseWriter, r *http.Request) {

// 	progress := Progress{YearStart: 2015, YearEnd: 2021}
// 	solutions := make(map[string]*Solution)

// 	a := SERVER.application

// 	for year := constants.MIN_YEAR; year <= constants.MAX_YEAR; year++ {
// 		appLogic := a.GetAppLogic(year)
// 		for day := 1; day <= 25; day++ {
// 			methodNamePart1 := fmt.Sprintf("Y%vD%02dP1", year, day)
// 			methodNamePart2 := fmt.Sprintf("Y%vD%02dP2", year, day)
// 			methodNamePart1Api := fmt.Sprintf("Y%vD%02dP1Api", year, day)
// 			methodNamePart2Api := fmt.Sprintf("Y%vD%02dP2Api", year, day)

// 			_, _, m1exists := appLogic.GetMethod(methodNamePart1)
// 			_, _, m2exists := appLogic.GetMethod(methodNamePart2)
// 			_, _, m1existsApi := appLogic.GetMethod(methodNamePart1Api)
// 			_, _, m2existsApi := appLogic.GetMethod(methodNamePart2Api)

// 			s := &Solution{Part1Solution: m1exists, Part2Solution: m2exists, Part1Api: m1existsApi, Part2Api: m2existsApi}
// 			key := fmt.Sprintf("%v.%v", year, day)
// 			solutions[key] = s
// 		}
// 	}

// 	progress.Solutions = solutions
// 	msgb, _ := json.Marshal(progress)
// 	msg := string(msgb)
// 	length_str := fmt.Sprintf("%v", len(msg))
// 	w.Header().Set("Content-Type", "application/json") // this
// 	w.Header().Set("Content-Length", length_str)       // this
// 	fmt.Fprint(w, msg)
// }

func apiSummaryFunc(w http.ResponseWriter, r *http.Request) {

	a := SERVER.application
	// m := make([]*utils.Summary, 0)
	// for year := constants.MIN_YEAR; year <= constants.MAX_YEAR; year++ {

	// 	for day := 1; day <= 25; day++ {

	// 		// splits := strings.Split(r.URL.Path, "/")
	// 		// year := splits[len(splits)-2]
	// 		// day := splits[len(splits)-3]
	// 		// iyear, _ := strconv.Atoi(year)
	// 		// iday, _ := strconv.Atoi(day)
	// 		summary := a.GetSummary(year, day)
	// 		m = append(m, summary)
	// 		// if summary != nil {

	// 		// 	msgb, _ := json.Marshal(summary)
	// 		// 	// msgb, _ := json.MarshalIndent(summary, "", "   ")
	// 		// 	msg := string(msgb) + "\n"
	// 		// 	// fmt.Println(msg)

	// 		// 	key := fmt.Sprintf("%v_%v", year, day)
	// 		// 	value := summary

	// 		// 	m[key] = value
	// 		// 	fmt.Printf("key=%v\n", key)
	// 		// }
	// 	}
	// }
	// fmt.Println(m)
	// type F struct {
	// 	A string
	// 	B string
	// }
	// f1 := &F{A: "hi", B: "foo"}
	// f2 := &F{A: "hi", B: "foo"}
	// var s utils.Summary
	arr := make([]*utils.Summary, 0)
	for _, s := range a.GetSummaries() {
		arr = append(arr, s)
	}
	// arr = append(arr, f1)
	// arr = append(arr, f2)

	msgb, _ := json.MarshalIndent(arr, "", "   ")
	msg := string(msgb) + "\n"
	fmt.Println(msg)
	length_str := fmt.Sprintf("%v", len(msg))
	w.Header().Set("Content-Type", "application/json") // this
	w.Header().Set("Content-Length", length_str)       // this
	fmt.Fprint(w, msg)

}

func filterBlogEntries(summaries []*utils.Summary) []*utils.Entry {
	results := make([]*utils.Entry, 0)
	for _, s := range summaries {
		if s != nil && s.Entries != nil && len(s.Entries) > 0 {
			for _, e := range s.Entries {
				results = append(results, e)
			}
		}
	}

	sort.Slice(results, func(i int, j int) bool {
		e1 := results[i]
		e2 := results[j]
		return e1.Date > e2.Date
	})
	return results
}

func blogFunc(w http.ResponseWriter, r *http.Request) {
	// msg := "<!DOCTYPE html>\n<!--\nHi!\n\nThis is my Rube Goldberg Advent of Code visualisations attempt.\n\nThere isn't anything to see here yet - but there is an api at /api/solutions \n-->\n<html>\n\t<head>\n\t\t<title>AOC</title>\n\t</head>\n\t<body>AOC {YEAR} <a href='/api/solutions'>[solutions]</a></body>\n</html>"
	a := SERVER.application
	// sort by lastmodified, then by year/day in reverse
	summaries := a.GetSummaries()
	blogEntries := filterBlogEntries(summaries)

	msg := ""
	msg += "<DOCTYPE html>"
	msg += "<html>"
	msg += "<head>"
	msg += "<title>AOC Blog</title>"
	msg += "<meta charset=\"UTF-8\">"
	msg += "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">"
	msg += "<link rel=\"stylesheet\" href=\"/css/style.css\" />"
	msg += "<link rel=\"stylesheet\" href=\"/css/blog.css\" />"
	msg += "</head>"

	msg += "<body>"
	msg += "AOC 2022"
	msg += "<p>This is my set of solutions for <a href=\"https://www.adventofcode.com\">Advent Of Code</a>.  The code is available on <a href=\"https://github.com/simonski/aoc\">github.com/simonski/aoc</a></p>"

	msg += "<table>"
	for _, entry := range blogEntries {
		msg += "<div class='entry'>"
		msg += "<span class='entry_header'>"
		msg += fmt.Sprintf("<span class='entry_date'>%v</span>", entry.Date)
		msg += fmt.Sprintf("<span class='entry_title'>%v</span>", entry.Summary.Name)
		msg += "</span>"
		msg += "<p>"
		msg += fmt.Sprintf("<span class='entry_notes'>%v</span>", entry.Notes)
		msg += "</p>"
		msg += "</div>"
		msg += "<br/>"
	}

	msg += "</table>"
	msg += "</body>"
	msg += "</html>"

	// msg = strings.ReplaceAll(msg, "{YEAR}", fmt.Sprintf("%v", time.Now().Year()))
	w.Header().Set("Content-Type", "text/html")

	length_str := fmt.Sprintf("%v", len(msg))
	w.Header().Set("Content-Type", "text/html")  // this
	w.Header().Set("Content-Length", length_str) // this

	fmt.Fprint(w, msg)
}

/*
<rss version="2.0">
<channel>
<title>Liftoff News</title>
<link>http://liftoff.msfc.nasa.gov/</link>
<description>Liftoff to Space Exploration.</description>
<language>en-us</language>
<pubDate>Tue, 10 Jun 2003 04:00:00 GMT</pubDate>
<lastBuildDate>Tue, 10 Jun 2003 09:41:01 GMT</lastBuildDate>
<docs>http://blogs.law.harvard.edu/tech/rss</docs>
<generator>Weblog Editor 2.0</generator>
<managingEditor>editor@example.com</managingEditor>
<webMaster>webmaster@example.com</webMaster>
<item>
<title>Star City</title>
<link>http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp</link>
<description>How do Americans get ready to work with Russians aboard the International Space Station? They take a crash course in culture, language and protocol at Russia's <a href="http://howe.iki.rssi.ru/GCTC/gctc_e.htm">Star City</a>.</description>
<pubDate>Tue, 03 Jun 2003 09:39:21 GMT</pubDate>
<guid>http://liftoff.msfc.nasa.gov/2003/06/03.html#item573</guid>
</item>
<item>
<description>Sky watchers in Europe, Asia, and parts of Alaska and Canada will experience a <a href="http://science.nasa.gov/headlines/y2003/30may_solareclipse.htm">partial eclipse of the Sun</a> on Saturday, May 31st.</description>
<pubDate>Fri, 30 May 2003 11:06:42 GMT</pubDate>
<guid>http://liftoff.msfc.nasa.gov/2003/05/30.html#item572</guid>
</item>
<item>
<title>The Engine That Does More</title>
<link>http://liftoff.msfc.nasa.gov/news/2003/news-VASIMR.asp</link>
<description>Before man travels to Mars, NASA hopes to design new engines that will let us fly through the Solar System more quickly. The proposed VASIMR engine would do that.</description>
<pubDate>Tue, 27 May 2003 08:37:32 GMT</pubDate>
<guid>http://liftoff.msfc.nasa.gov/2003/05/27.html#item571</guid>
</item>
<item>
<title>Astronauts' Dirty Laundry</title>
<link>http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp</link>
<description>Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them. Instead, astronauts have other options.</description>
<pubDate>Tue, 20 May 2003 08:56:02 GMT</pubDate>
<guid>http://liftoff.msfc.nasa.gov/2003/05/20.html#item570</guid>
</item>
</channel>
</rss>
*/
func rssFunc(w http.ResponseWriter, r *http.Request) {
	msg := "<!DOCTYPE html>"
	msg += "<head>"
	msg += "\n\t\t<title>AOC</title>"
	msg += "<!--\nHi!\n\nThis is my Rube Goldberg Advent of Code visualisations attempt.\n\nThere isn't anything to see here yet - but there is an api at /api/solutions \n-->\n"
	msg += "</head>"
	msg += "\n\t<body>AOC {YEAR} <a href='/api/solutions'>[solutions]</a>"
	msg += "</body>"
	msg += "\n</html>"

	msg = strings.ReplaceAll(msg, "{YEAR}", fmt.Sprintf("%v", time.Now().Year()))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, msg)

}

func sortSummaries(summaries []*utils.Summary) []*utils.Summary {
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
	return summaries
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

	msg := ""
	msg += "<DOCTYPE html>"
	msg += "<html>"
	msg += "<head>"
	msg += "<title>AOC Blog</title>"
	msg += "<meta charset=\"UTF-8\">"
	msg += "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">"
	msg += "<link rel=\"stylesheet\" href=\"/css/style.css\" />"
	msg += "<link rel=\"stylesheet\" href=\"/css/blog.css\" />"
	msg += "</head>"

	msg += "<body>"
	msg += "AOC 2022"
	msg += "<p>This is my set of solutions for <a href=\"https://www.adventofcode.com\">Advent Of Code</a>.  The code is available on <a href=\"https://github.com/simonski/aoc\">github.com/simonski/aoc</a></p>"

	msg += "<table>"

	summaries = sortSummaries(summaries)
	for _, v := range summaries {
		if v == nil {
			continue
		}
		ticks := ""
		if v.ProgressP1 == utils.Completed {
			ticks += "&#9733;"
		} else if v.ProgressP1 == utils.Started {
			ticks += "&#9775;"
		}

		if v.ProgressP2 == utils.Completed {
			ticks += "&#9733;"
		} else if v.ProgressP2 == utils.Started {
			ticks += "&#9775;"
		}

		href := fmt.Sprintf("https://adventofcode.com/%v/day/%v", v.Year, v.Day)
		line := "<tr>"
		line += fmt.Sprintf("<td>%v<td>", ticks)
		line += fmt.Sprintf("<td>%v %02d</td><td><a href='%v'>%v</a></td>\n", v.Year, v.Day, href, v.Name)
		line += "</tr>\n"

		msg += line
	}
	msg += "</table>"
	msg += "</body>"
	msg += "</html>"

	// msg = strings.ReplaceAll(msg, "{YEAR}", fmt.Sprintf("%v", time.Now().Year()))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, msg)
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
