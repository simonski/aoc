# AOC

This houses my Advent of Code [https://adventofcode.com/](https://adventofcode.com/) attempts, in go.

## Prepare

Every year I forget to remember and prepare.   This year will be no different, however if you want to go fast, read slow:

- https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
- https://en.wikipedia.org/wiki/Memoization
- https://en.wikipedia.org/wiki/Recursion
- https://en.wikipedia.org/wiki/Recursion
- https://en.wikipedia.org/wiki/Shoelace_formula
- https://en.wikipedia.org/wiki/Pick%27s_theorem
- https://en.wikipedia.org/wiki/Chinese_remainder_theorem    
- https://en.wikipedia.org/wiki/Flood_fill
- https://en.wikipedia.org/wiki/Breadth-first_search
- https://en.wikipedia.org/wiki/Depth-first_search

## Install

```bash
git clone github.com/simonski/aoc.git
cd aoc
make setup
make
```

This will build `aoc` which you can then type `./aoc` on and play with.

## Run

Type `./aoc`

A calendar will be printed with the stars. Each day can be run using `aoc run <year> <day>`.   Every year I break this, so for 2023 I think I will refactor the lot to the `Puzzle` interface I dreamed up.

### Webserver time

The whole thing is wrapped as a Rube Goldberg toy webserver with webassembly and what-not. Run the server and see some visualisations for the problems I liked.  Note that I use `Caddy` for the tls. 

I use a complicated arrangement of javascript talking back to the server for a `d3.js` visualisation. Good luck with that part next year :).  There's a bunch of dead and unused code in there that I'll tidy up.

```bash
cd $CODE/aoc
./aoc server
```

## Extend

Let's say you work on a new day, for example 1st December 2023.  I use the script `start.sh` which token-switches some template code:

```bash
./start.sh
YEAR: 2022
DAY: 1
```

You can then run it as an application (<b>edit</b> not true: run via tests until I integrate.)

```bash
make
aoc run 2022 01
```

Eventually I'll be finished and can merge back to `main`, at which point if I addded a visualisation or blogged it at all, an image can be built and redeployed to [https://aoc.simonski.com](https://aoc.simonski.com).  I do this by running `make docker publish`.

## Deploy

Build and push the AOC image to `ghcr.io`, which is then picked up by my server via a `watchtowerr`.

```bash
make docker publish
```
