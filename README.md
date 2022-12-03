# AOC

This houses my Advent of Code [https://adventofcode.com/](https://adventofcode.com/) attempts, in go.

## Install

```bash
git clone github.com/simonski/aoc.git
cd aoc
make
```

This will build `$GOBIN/aoc` which you can then type `aoc` on and play with.

### Webserver time

The whole thing is wrapped as a toy server.  Run the server and see some visualisations for the problems I liked.  Note that I use `Caddy` for the tls.

I use a complicated arrangement of javascript talking back to the server for a `d3.js` visualisation. Good luck with that part next year :).  There's a bunch of dead and unused code in there that I'll tidy up.

```bash
cd $CODE/aoc
caddy start
make
./aoc server
```



## Run

Type `aoc`

A calendar will be printed with the completion stars. Each day can be run using `aoc run <year> <day>`

## Extend

Let's say you work on a new day, for example 1st December 2022.  Each day goes in its own `feature/YYYY_DD` branch taken from `develop`.  I do this with a script, `start_problem.sh` which token-switches some template code:

```bash
./start_problem.sh 2022 01
```

You can then run it as an application

```bash
make
aoc run 2022 01
```

Eventually I'll be finished and can merge back to `develop` and finally `main`, at which point an image is built and redeployed to [https://aoc.simonski.com](https://aoc.simonski.com).  I do this by running `make docker publish`.

## Deploy

Build and push the AOC image to `ghcr.io`, which is then picked up by my server via a `watchtowerr`.

```bash
make docker publish
```