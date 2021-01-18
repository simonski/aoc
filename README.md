# AOC (go)

This houses my Advent of Code [https://adventofcode.com/2020](https://adventofcode.com/2020) attempts, in go.

## Install

If you just want to run it, do this

    go get github.com/simonski/aoc

## Building

If you want to build it yourself, try this

    git clone github.com/simonski/aoc.git
    cd aoc
    make all

This will build `$GOBIN/aoc` which you can then type `aoc` on and play with.

## Runnning

Once you've got it, run it like so

    aoc

After that, have a look at the puzzles

    aoc list

A completed year/day combo will show up as a green star.  If the puzzle has a visualisation it will show as a yellow star.

Let's run 2020, day 1

    aoc run 2020 01

Let's render day 24, 2020 as an animation

    aoc render 2020 24


    