# AOC

This houses my Advent of Code [https://adventofcode.com/](https://adventofcode.com/) attempts, in go.

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

## Extending

When you work on a new day, e.g. 2021-01-01

    - copy the `aoc20XX_XX.go -> aoc2021_01.go`
    - copy the test `aoc_20XX_XX_test.go -> aoc_2021_01_test.go`
    - copy the data `aoc_20XX_XX_data.go -> aoc_2021_01_data.go`

You can then run it

    ./aoc run 2021 01
