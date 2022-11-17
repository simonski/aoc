package constants

const VERSION = "0.0.1"

const MIN_YEAR = 2015
const MAX_YEAR = 2023

const USAGE_OVERALL = `aoc is my Advent Of Code set of attempts.

Usage:
    aoc <command> [arguments]
	
LIST
The commands are:

    run    (year) (day)      run a puzzle 
    render (year) (day)      render a puzzle to an animated gif
    summary (year) (day)     prints a summary of a given problem
    
    info                     prints build & environment information
    server                   runs the webserver   (-p port [-fs path])
    version                  prints aoc version
`
