Puzzle: Pyroclastic Flow
Year: 2022
Day: 17

2022-12-23 130am.  Ah, the relief!

Just finished D16 - that was HARD.  This is nicer.

This I think I can do reasonably quickly - Eric will have thrown a dragon in here for Pt2 - I suspect he will say rocks break apart on landing and fall into individual pieces, but we'll see.

I'm going for a Chamber, Rock and a RockFragment which is basically a rock but ready to decompose.

A rock will have a Bounds() which is its bounding box then inside that each square is identified as a Piece of solid rock or not ( a pice of air ).  The rock itself can be moved around, the x,y of the pieces becomes relative to the position we are in.

The chamber will then have an active piece of rock and a Tick(input) where it will move them down.

The chamber will have a debug, once it works we can then generate the chamber knowing the chamber bounds and that will give us a visualisation opportunity.

2022-12-23 0830 - putting the model together

Rock, Piece, Chamber.  I'll cool my jets a little and put some TDD Around these as there are some assumptiosn I'd like to get correct first time if possible. 

2022-12-23 1518 - all together and then I go and mistake the leftmost/rightmost and bottom-most entries for the "funny" shapes.  So I incur a 5m penalty.  Let's see, I *should* get it right next time...




2022-12-26 

Got up at 6am very tired and unwell with my cold, but I got P1 done. I messed about and got completely tangled in figuring out the position of pieces.  In the end I rewrote a piececache and found a - instead of + bug.  Ho-hum. Should have taken 3 hours, took day.s  but that's okay!

P2 - cycle detection I think. My approach is: 

- Find a repeating pattern - a key - by grabbing a few rows and looking for them repeating. 
- Then take the difference (between the repeats) and verify *it* repeats.
- Then look for the first occurance of that to get the row.  
- Then figure out the rock number and what that pattern looks like.

I think the "find a pattern" happens *after* some preamble of rows.  Additionally, a single repeating row won't be 
enough - I'll need a few, or quite a few (I don't know).  So it needs to be something I can vary.   This part I think
will be a bit of stdout etc; semi-automated via tests etc.

Tuesday 27th 0800

Am thinking: a dequeue or buffer with a moving head and tail; 
    1. commence the buffer after some period of time (or from row 0 and let it just fill up a bit but I thinka preamble isnecessary)
    2. fill the buffer
    3. assume it is ain a cycle; detect by taking FINAL N rows as a key and walking FORWARD to find the key.
    4. vary the key size until we see what we see

