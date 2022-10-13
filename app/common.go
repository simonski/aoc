package app

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
