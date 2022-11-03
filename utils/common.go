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
