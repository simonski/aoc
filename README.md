# AOC

This houses my Advent of Code [https://adventofcode.com/](https://adventofcode.com/) attempts, in go.

You can Install, Run and Extend it

## Install

If you just want to run it, do this

    go get github.com/simonski/aoc

Or build it yourself

    git clone github.com/simonski/aoc.git
    cd aoc
    make

This will build `$GOBIN/aoc` which you can then type `aoc` on and play with.

## Runnning

Type

    aoc

A completed year/day combo will show up as a green star.  If the puzzle has a visualisation it will show as a yellow star.

Run 2020, day 1

    aoc run 2020 01

Let's render day 24, 2020 as an animation

    aoc render 2020 24

Run the server and see some visualisations

    caddy start
    ./aoc server

## Extend

When you work on a new day, for example 1st December 2022.  Each day goes in its own `feature/YYYY_DD` branch taken from `develop`:

    ./start_problem.sh 2021 01

You can then run it as an application

    ./aoc run 2021 01

Eventually I'll be finished and can merge back to develop and finally main.

## Deploy

    make docker publish

Will build and push the AOC image, which is then picked up by the server via a `watchtowerr` setup elsewhere.
