Puzzle: Day 3: Gear Ratios
Year: 2023
Day: 3

I "saw" this one, so the solution was fixed from the start.  In this case I believe the experience of "keep that for later" that Eric has so kindly taught me over the years helped me out.  I thought the symbol locations would be useful.


Part1:

1. detect a number
2. check if it has a symbol adjacent


for row in rows
1. for col in cols
    digit = GRID(col,row)
    if isDigit: (5)
        for col in cols
            if next is digit (5)
                append to number = 55
            else:
                break returning final index
    now we have a start index, the final number and a final index
    walk adjacent around the number (row, startcol, endcol)
    if adjacent, include
    else col = endcol+1

2. 
