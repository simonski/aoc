Puzzle: Day 5: If You Give A Seed A Fertilizer
Year: 2023
Day: 5

08:00

It's all about the data structure.  Get that right, it should be okay.   It's *also* all about testing.  Check as I got, did someone say "TDD"? :)

My impression is this is the enjoyable not-hard but finicky problem we come for and end up staying for the horror. 

Step 1.
    Figure some tests against imaginary strutures
    Figure some structure
    Ensure I create exactly what he has made via testing


Part1: OK

Part2: I *swear* there is a way of reducing the 2b search space but I couldn't fathom it.  I could have run a goroutine per seedset but I left it to crank for 5m or so.  Turns out embedded values in structs over functions to compute is actually faster - no compute, just memory!
