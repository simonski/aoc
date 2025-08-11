Puzzle: Boiling Boulders
Year: 2022
Day: 18

Tuesday day 27th - Part1 solved in minutes.  I anticipated a complex p2 so I made a cube, point3d, grid classes ready for it.

Tuesday night before bed - I read and didn't understand the wording at first, then I looked at the example and couldn't
visualise (in my head) what he was talking about until I re-read it afew times and decided ok, it's about working out
the external surface area.  Maybe it straight up says that but I'm addled and dumb.  

1. So, I had ideas about planes collapsing inwards and "scoring" nodes until they meet their other plane, doing that on
all three axes (six collapsing planes) would give me all cube faces detected.

2. Then I thought about the internal air pockets as "not-knowing" vs. "knowing".  Then I had an idea about the complicated part; cubes "inside"; if they have a way out this is going to affect the external area; so if they don't, I can discount them from the surface area.  So then I go, find a way to.. djikstra/bfs.. bfs!.  So I think I need to

- find air pockets
- subtract whatever they gave me from P1

I don't understand his "diagonal" reference though.

	