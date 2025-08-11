package aoc2021

/*
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg


acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf

 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc

acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf

1 ab
4 eafb
7 dab
8 acedgfb

> this means we know ab is the right digits
D=TOP > this means 7 as dab means d is the top

So now we can say, okay, difference between 9 and 0 is
the bottom left is empty in 9 and the middle is occupied

So, now, we have 0, 2, 3, 5, 6, 9

0 has 6
* 1 has 2
2 has 5
3 has 5
* 4 has 4
5 has 5
6 has 6
* 7 has 3
* 8 has 7
9 has 6



  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg


*a
*b
*c
*d
*e
*f
*g


1. find 1, 7, 4, 8
2. calculate TOP (a)
2. find the size 6 numbers
3. find the size 5 numbers

there are three sixes (0, 6, 9)
there are three fives (2, 3, 5)

4. finding e and 9
	subtract 4 from one of the unknown sets to find the difference
	9 - 4 = (abcdfg) - (a) - (bcdf) = g
	6 - 4 = (abdefg) - (a) - (bcdf) = eg
	0 - 4 = (abcefg) - (a) - (bcdf) = eg
		the smaller is g
		the diffference (e) is now found, too
		also we now know 9 (the smaller)

5. if we know 9 is one of the size six entries, we know
	one of the others is a six, so the other must be the 0

	6 - 7 = (abdefg) - (acf) = bdeg
	0 - 7 = (abcefg) - (acf) = beg

	So the 6 contains is size 4
	So teh 0 is size 0
	The difference bwteeen them is the d in 6 now we have middle

	now we know
	d
	6
	9
	0

6. now we can calculate 4 - 6  leaving top-right, c
	now we know c

7. look at 1 again, now we know c, we know f

8. now we know f, the last remaining is b











 8 (a,b,c,d,e,f,g)
-7 (a c, f)
= b,d,e,g
-4 (b,c,d,f)
= e,g
-1 (c,f)
= e,g

1, 4, 7, 8



e/g is bottom left / bottom




*TOP
MIDDLE
BOTTOM
TOPLEFT
BOTTOMLEFT
TOPRIGHT
BOTTOMRIGHT


-----
    |
    |
    |
-----
|   |
|   |
|   |


1. find the 1
2. FOUND TOP find the top (difference between 1 and 7)
2. FOUND BOTTOM LEFT compare 8 with 9 - the bottom left is then identified
3. FOUND TOP RIGHT compare 8 with 6 - the top right is identified
4. FOUND BOTTOM RIGHT now we know top right, we cna work out bottom right
5. FOUND MIDDLE COMPARE 8 with 0, we have MIDDLE
6. FOUND TOPLEFT number 4 contains TOP-LEFT - whats missing against our comparison
	4 is TOPLEFT TOPRIGHT MIDDDLE BOTTOMRIGHT
	SO WE CAN NOW IDENTIFY TOPLEFT
7. FOUND BOTTOM - REMAINDER IS BOTTOM

*/
