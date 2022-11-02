# AOC TODO

## DOING

-- Verify an example "new day" call for a new problem
-- Implement the now-implemented Summary() calls acrosss the problems
-- Refactor landing/default page to show a table of progress
-- API for summaries
-- GUI on website for summaries for each day.
-- embed the whole app so I can see the source online also... urgh

branch feature/website and any open feature/202x_xx branches are incomplete.

 - figure out MVP for the common javascript (if any?) and then render the page.	


    WEB LANDING PAGE FOR PROBLEMS
        https://localhost/visualisations and
        https://localhost/visualisations/home
        
     /api/visualisations/home/index.html and sketch.js to create a heatmap of 
     each solution I've worked on by year.   

     Each block should be able to render the completion status, noetes etc.   

    fetch results from /api/solutions  (returns json)
    for each result
     for year
        start row
        for day 01..24
            start col
            render star
            end col
        end row

        contains current refactored layout
        intention is to create a github action to deploy
        AND then create a heatmap for progress
            WHICH will involve refactoring for every problem to indicate
                Problem {
                    Year
                    Day
                    State: solved, unsolved, inprogress
                    Complexity: N
                    Started: time
                    Completed: time
                    Notes: string
                }

- writeup github action build for aoc
- staticcheck for compliance and go idiomsS
- refactor goutils/cli to cli/cli
- fix all tests to compile and... well, work
  - fix tests to timeout properly
  - fix tests to run on specific years
  - look for "if (true) { return }" and fix the
- standardise the webassemly visualisation
- writeup and get the docker build working for the aoc build
- move to ghcr.io via a github action
- modify the webwite to pull from ghcr if possible?
- move all years work down to their own package
    doing 2015
    2020
    2021

## TODO

- Can't work out how to run the tests.
- remove statik -> replace with embed
- review the module usage
- finish the template off - code, test, data files.
- update all code for completeness so `./aoc list` looks nice.
- update readme to figure out how to reflect on the app where we can drop the number of files down
- I think I can make an application struct per year, then reflect directly on that perhaps?

This is my TODO list for AOC 2020.

- all days: code style, comments.
- day1 part1, sort and binchop
- day1 part2, go look at algos book again

## 2022-11-02

-- DONE Include a function that returns the summary for a day YyyyyDdd_Summary() *utils.Summary

-- DONE  - refactor the API so I can determine the state of problem, its complexity and notes, if it has a visualisation.
	struct Problem
		day
		part
		solved
		effort
		notes
		labels

