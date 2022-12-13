Hill Climbing


2022-12-13 

I have this thing about perhaps I need a tree for the paths investigated, where node could have a visited 
so that I can discount it from future endevours.   Honestly, I dunnno.  The visualisation ehlps thioguht.

I'm guessing, so I think I should go read about hill climbing again.


Ok, so it's Dijkstra. Except, not quite.  The visualisation is a bit of a pain so while I'm
tempted to go write a GUI (fyne?) I know deep down it's all just about the algo.  Really the terminal
algo is fine as long as I mod 100000 or so.   I could really do with some sort of... keyboard input to
fast-forward a number of steps.  Silly, I'll end up writing a game.

2022-12-12 Phew! - made it.  LCM fixed day 11, which got me to this horror show.

Ok, so it's Dijkstra. Except, maybe not quite.

- I've gone done a weird rabbit hole.  I made a grid, then I started my walking using Dijkstra and decided
this was not correct.

- Visualisation of an (checks notes.. X.X grid) is <b>hard</b> in the terminal.  Combine that with the
fact that I (incorrectly) am running to millions of iterations now and something is clearly incorrect.  This makes me think I need some sort of GUI to visualise where I'm going wrong.

- Exccept really I just need a goood algo.

- I think that some manual tweaking could get me there, for example, preferring to "select the neighbour in the direction nearest the target", or "select the same direction as last time" if possible *might* be a speedier way (given the map we have, that is).

I am so tempted to manually walk it on paper.  What a nonsense!

- MAybe I could go backwards? but does that actually make a difference?

