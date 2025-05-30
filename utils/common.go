package utils

// Each solution as I complete it should return the problems so I can do something (Websitey) with them for
// future reference
type Summary struct {
	Name            string   // headline
	URL             string   // the problem on AOC website
	Year            int      // YYYY
	Day             int      // 1..25
	API             bool     // is there an API I can call to get the solution and visualise it over the web?
	StdOut          bool     // does it print the solution to stdout?
	VisualisationP1 bool     // can I visualise i over the web?
	VisualisationP2 bool     // can I visualise i over the web?
	ProgressP1      Progress // has it defeated me
	ProgressP2      Progress // has it defeated me
	Attempts        int      // how many times have I tried this
	ComplexityP1    int      // 0..N easy..complex
	ComplexityP2    int      // 0..N easy..complex
	SatisfactionP1  int      // 0..N  0=made me insane, 100 = so happy
	SatisfactionP2  int      // 0..N  0=made me insane, 100 = so happy
	EffortP1        int      // how hard did I have to work (so, it might have been easy but took a lot because of data wrangling)
	EffortP2        int      // how hard did I have to work (so, it might have been easy but took a lot because of data wrangling)
	DateStarted     string   // date time
	DateCompleted   string   // date time
	LastModified    string   // date time
	Concepts        string   // what CS concepts, algos, data structures did I use
	Entries         []*Entry
}

func NewSummary(year int, day int) *Summary {
	s := Summary{Year: year, Day: day}
	s.Entries = make([]*Entry, 0)
	s.ProgressP1 = Unknown
	s.ProgressP2 = Unknown
	s.Name = "Never attempted"
	return &s
}

type Entry struct {
	Date    string
	Title   string
	Notes   string
	Summary *Summary
}

func NewEntry(date string, title string, notes string, s *Summary) *Entry {
	return &Entry{Date: date, Title: title, Notes: notes, Summary: s}
}

type Progress int

const (
	Unknown Progress = iota
	NotStarted
	Started
	Failed
	Paused
	Completed
)

func (p Progress) String() string {
	switch p {
	case Unknown:
		return "Unknown"
	case NotStarted:
		return "Not started"
	case Started:
		return "Started"
	case Failed:
		return "Failed"
	case Paused:
		return "Paused"
	case Completed:
		return "Completed"
	}
	return "unknown"
}

// Each solution as I complete it should return the problems so I can do something (Websitey) with them for
// future reference
type Problem struct {
	Name          string // headline
	URL           string // the problem on AOC website
	StdOut        bool   // does it print the solution to STDout?
	API           bool   // is there an API I can call to get the solution and visualise it over the web?
	Visualisation bool   // can I visualise i over the web?
	Year          int    // YYYY
	Day           int    // 1..25
	Part          int    // 1 or 2
	GaveUp        bool   // has it defeated me
	InProgress    bool   // is it currently in progress
	Paused        bool   // have I stopped working on it for non giveup reasons
	Attempts      int    // how many times have I tried this
	Solved        bool   // did I manage it
	Complexity    int    // 0..N easy..complex
	Satisfaction  int    // 0..N  0=made me insane, 100 = so happy
	Effort        int    // how hard did I have to work (so, it might have been easy but took a lot because of data wrangling)
	Started       string // date time
	Completed     string // date time
	Notes         string // what did I think
	Concepts      string // what CS concepts, algos, data structures did I use
}

type AOCDay interface {
	Part1() Problem
	Part2() Problem
}
