# AOC

This houses my Advent of Code [https://adventofcode.com/](https://adventofcode.com/) attempts, in go.

You can install, run and extend it

## Install

If you just want to run it, do this

    go get github.com/simonski/aoc

Or build it yourself

    git clone github.com/simonski/aoc.git
    cd aoc
    make

This will build `$GOBIN/aoc` which you can then type `aoc` on and play with.

## Extend

When you work on a new day, for example 1st December 2022.  Each day goes in its own `feature/YYYY_DD` branch taken from `develop`:

    ./start_problem.sh 2021 01

You can then run it as an application

    ./aoc run 2021 01

Eventually I'll be finished and can merge back to `develop` and finally `main`, at which point an image is built and redeployed to [https://aoc.simonski.com](https://aoc.simonski.com).

## Runnning

Type

    aoc

A completed year/day combo will show up as a green star.  If the puzzle has a visualisation it will show as a yellow star.

Run 2020, day 1

    aoc run 2020 01

Let's render day 24, 2020 as an animation

    aoc render 2020 24

Let's see what I thought of day 14, 2019 

    aoc summary 2019 14

Run the server and see some visualisations

    caddy start
    ./aoc server


## Deploy

    make docker publish

Will build and push the AOC image, which is then picked up by the server via a `watchtowerr` setup elsewhere.



// AOC_2015_07 is the entrypoint
func (app *Application) Y2015D08(cli *cli.CLI) {
	app.Y2015D08P1(cli)
	app.Y2015D08P2_inprogress(cli)
}

Where _inprogress is an in progress development entrypoint
where removing the _inprogress means it is complete.
