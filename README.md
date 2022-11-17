# AOC

This houses my Advent of Code [https://adventofcode.com/](https://adventofcode.com/) attempts, in go.

## Install

    git clone github.com/simonski/aoc.git
    cd aoc
    make

This will build `$GOBIN/aoc` which you can then type `aoc` on and play with.

## Extend

When you work on a new day, for example 1st December 2022.  Each day goes in its own `feature/YYYY_DD` branch taken from `develop`:

    ./start_problem.sh 2022 01

You can then run it as an application

    ./aoc run 2022 01

Eventually I'll be finished and can merge back to `develop` and finally `main`, at which point an image is built and redeployed to [https://aoc.simonski.com](https://aoc.simonski.com).  I do this by running `make docker publish`.

## Runnning

Type `aoc`

A calendar will be printed with the completion stars.

When working on a given problem, I use a unit test for that day

```bash
aoc run 2020 01
```

Run 2020, day 1

```bash
aoc run 2020 01
```

Let's render day 24, 2020 as an animation

```bash
aoc render 2020 24
```

Let's see what I thought of day 14, 2019 

```bash
aoc summary 2019 14
```

Run the server and see some visualisations

```bash
caddy start
./aoc server
```

## Deploy

```bash
make docker publish
```

Will build and push the AOC image, which is then picked up by the server via a `watchtowerr` setup elsewhere.


```go
// AOC_2015_07 is the entrypoint
func (app *Application) Y2015D08(cli *cli.CLI) {
	app.Y2015D08P1(cli)
	app.Y2015D08P2_inprogress(cli)
}
```

