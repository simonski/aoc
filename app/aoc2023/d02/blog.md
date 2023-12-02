Puzzle: Cube Conundrum
Year: 2023
Day: 02

04:56.  Woke up 4am on unrelated issue (heating) - it's -5 outside and boiler not doing its one job.  Pip was cold so woke me.   Reviewed and considered writing an API to fetch puzzle stuff but decided a bit of elbow work is fine.

I want to refactor the whole thing down to a standard `Puzzle` | `Summary` set of structs so that I can streamline the
server, webassembly, visualisations etc. 

The main `app.go` has a nasty `years * days` import and invoke cycle commented out as I can't figure a nicer way of 
instantiating each puzzle.  I can hide it by year but it's the same import/invoke.   I'll take a look later.

05:41.  Solved.  Ok, part two day one was an aberration :) - today was a nice gentle one.  I tried going TDD and probably went TD-ish-D on it.  Pythons string is a winner, go's is verbose.  But then come back later, perhaps gos rigidity is a strength.  strings - the eternal golden struggle.