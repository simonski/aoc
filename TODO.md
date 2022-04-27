# AOC TODO

## DOING

- writeup github action build for aoc
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
