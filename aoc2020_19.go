package main

/*

--- Day 19: Monster Messages ---
You land in an airport surrounded by dense forest. As you walk to your high-speed train, the Elves at the Mythical Information Bureau contact you again. They think their satellite has collected an image of a sea monster! Unfortunately, the connection to the satellite is having problems, and many of the messages sent back from the satellite have been corrupted.

They sent you a list of the rules valid messages should obey and a list of received messages they've collected so far (your puzzle input).

The rules for valid messages (the top part of your puzzle input) are numbered and build upon each other. For example:

0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"
Some rules, like 3: "b", simply match a single character (in this case, b).

The remaining rules list the sub-rules that must be followed; for example, the rule 0: 1 2 means that to match rule 0, the text being checked must match rule 1, and the text after the part that matched rule 1 must then match rule 2.

Some of the rules have multiple lists of sub-rules separated by a pipe (|). This means that at least one list of sub-rules must match. (The ones that match might be different each time the rule is encountered.) For example, the rule 2: 1 3 | 3 1 means that to match rule 2, the text being checked must match rule 1 followed by rule 3 or it must match rule 3 followed by rule 1.

Fortunately, there are no loops in the rules, so the list of possible matches will be finite. Since rule 1 matches a and rule 3 matches b, rule 2 matches either ab or ba. Therefore, rule 0 matches aab or aba.

Here's a more interesting example:

0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
Here, because rule 4 matches a and rule 5 matches b, rule 2 matches two letters that are the same (aa or bb), and rule 3 matches two letters that are different (ab or ba).

Since rule 1 matches rules 2 and 3 once each in either order, it must match two pairs of letters, one pair with matching letters and one pair with different letters. This leaves eight possibilities: aaab, aaba, bbab, bbba, abaa, abbb, baaa, or babb.

Rule 0, therefore, matches a (rule 4), then any of the eight options from rule 1, then b (rule 5): aaaabb, aaabab, abbabb, abbbab, aabaab, aabbbb, abaaab, or ababbb.

The received messages (the bottom part of your puzzle input) need to be checked against the rules so you can determine which are valid and which are corrupted. Including the rules and the messages together, this might look like:

0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
Your goal is to determine the number of messages that completely match rule 0. In the above example, ababbb and abbbab match, but bababa, aaabbb, and aaaabbb do not, producing the answer 2. The whole message must match all of rule 0; there can't be extra unmatched characters in the message. (For example, aaaabbb might appear to match rule 0 above, but it has an extra unmatched b on the end.)

How many messages completely match rule 0?

*/

import (
	"fmt"
	"regexp"
	"strings"

	goutils "github.com/simonski/goutils"
)

const DAY_19_INPUT = `3: 7 45 | 10 39
120: 109 45 | 16 39
84: 96 39 | 104 45
6: 120 39 | 113 45
111: 45 93 | 39 45
13: 17 45 | 96 39
74: 122 45 | 17 39
94: 66 45 | 119 39
127: 39 84 | 45 132
129: 45 128 | 39 35
112: 39 35 | 45 58
24: 45 76 | 39 112
43: 39 17 | 45 96
2: 45 5 | 39 77
71: 100 45
11: 42 31
51: 77 45
4: 124 39 | 85 45
45: "a"
78: 111 39 | 128 45
8: 42
104: 45 39 | 39 93
29: 122 39 | 66 45
42: 63 45 | 20 39
41: 73 45 | 19 39
110: 39 98 | 45 114
55: 45 104 | 39 122
0: 8 11
53: 39 34 | 45 89
39: "b"
61: 77 45 | 104 39
121: 45 65 | 39 1
105: 45 44 | 39 99
113: 9 39 | 103 45
117: 96 93
125: 39 108 | 45 43
69: 45 39
56: 50 45 | 12 39
73: 39 35 | 45 100
87: 39 100 | 45 111
10: 45 13 | 39 73
19: 39 69 | 45 58
100: 39 45
66: 93 93
46: 45 82 | 39 74
76: 66 45 | 111 39
103: 45 52 | 39 115
77: 45 39 | 45 45
52: 82 39 | 112 45
15: 45 47 | 39 132
68: 39 77 | 45 17
1: 45 47 | 39 61
65: 37 45 | 51 39
98: 45 2 | 39 80
35: 39 45 | 45 45
93: 39 | 45
126: 111 39 | 77 45
32: 45 36 | 39 37
37: 45 69 | 39 119
90: 39 17 | 45 69
21: 39 66 | 45 58
22: 39 56 | 45 64
7: 123 39 | 48 45
60: 45 102 | 39 26
107: 45 29 | 39 71
58: 39 45 | 39 39
70: 128 39 | 111 45
81: 27 45 | 129 39
67: 5 39 | 100 45
96: 45 45 | 39 39
116: 39 87 | 45 55
106: 39 51 | 45 92
14: 45 128 | 39 58
48: 39 104 | 45 5
72: 45 35 | 39 111
130: 118 45 | 28 39
115: 45 91 | 39 87
31: 39 6 | 45 22
30: 79 45 | 57 39
9: 125 39 | 49 45
122: 39 39 | 45 93
23: 101 45 | 78 39
47: 39 100 | 45 58
28: 45 111 | 39 122
101: 45 77 | 39 66
33: 39 5 | 45 111
95: 39 5
27: 58 45 | 17 39
16: 15 39 | 116 45
80: 45 119 | 39 66
92: 45 111 | 39 58
57: 39 73 | 45 86
123: 45 58 | 39 77
5: 39 39
128: 45 45
124: 101 39 | 126 45
108: 45 122 | 39 119
119: 45 39 | 39 39
50: 81 45 | 106 39
99: 130 45 | 46 39
132: 17 45 | 119 39
49: 70 45 | 117 39
63: 131 39 | 83 45
85: 39 33 | 45 97
18: 45 60 | 39 25
83: 39 54 | 45 4
38: 62 45 | 21 39
64: 39 30 | 45 121
118: 45 17
91: 39 17 | 45 100
82: 39 100 | 45 119
86: 119 39 | 111 45
89: 39 119 | 45 77
44: 39 127 | 45 107
88: 45 95 | 39 89
17: 39 45 | 45 39
131: 39 3 | 45 110
12: 39 24 | 45 23
26: 39 108 | 45 90
36: 58 45 | 100 39
97: 45 111
25: 39 88 | 45 32
62: 96 39 | 111 45
59: 39 119 | 45 96
34: 45 35 | 39 100
79: 68 39 | 94 45
40: 39 72 | 45 14
20: 39 105 | 45 18
75: 111 45 | 69 39
114: 67 45 | 59 39
54: 38 39 | 53 45
102: 39 75 | 45 80
109: 39 40 | 45 41

aaaabbaabbaaabaaabbaaaaa
aaababaabaaaabaaabababbabbbbaabbabbbbaababbaaabaababababababbabbabbaaabb
bbbbbbbaabbaabaaaaabbaababbbaaabbababbabaababbaa
bbbbbbabbabaababbababababbbabbabbbabaabaabbaaaba
abbbbbaabaabbaabbabaaaba
aaaabbaaabbbbbabaaabaabaaaaaaaab
bbabbababaababaabababaaaaabbbbbabbbbabbbabbbaabbbbabbbbb
aaabaaaaaaaabbaaabaaabbb
abbabaaabbaaabbabbbabbba
bababbaaaaaabbababbabaaa
aaaaabaaaababbbabbaabaaa
babaabbaaababbbabbaaaaaaaabbbabbbabbbaaabbabaabbaababbaabaabbbbb
baaaabbbbaaabaababaaaaaabbabaaaabbbaabbabbabaabbabbbbbabbaabbbabaaababba
aaaaabbbaabbbbbbbbabaaaababaaaab
abaabaaabaabbbabaaabbbaaabaabbbabbaaaaaabaaabababbaaabbabbbbabaa
abbbbbbabbaabbababaabaabbabaaabaabbbaaaa
babbbaaabaaaababababbbba
babbbabbbaaaababaaabaababbaababbababbaabaabaaababaabbbaaaabbababaaabaabbaababbaabbabbaab
bbaabaabaabababbbaabaabb
baaaaabbabbbabbabbbababbbabababb
baaabbbbaababbabbababababbbbbaabbabbabba
aaaaabbbaabaaabbbbbbabbbbbaaaaaaabbaaaab
bbaababababbbbbaabaaabbb
abaabbababaabaaabaababab
abbabbabababbababaababaababbababbbababaaabbbbbbaaaababaa
bbaaaabababbbbaabaaaaaabbabbabbbbabbabbbbaabbaab
aabaaabbbabaabbabbabbbbb
baaabbabaabbabbbbaaababb
ababaababbaabbbbbbbaaabaaaabbbbb
aaaaabbaaaababbbabaaabab
abbbaaabbaaabbbaababbabbbbbababbbbbbabba
aaaabaaabbbbbbabbaaabaaababbbabb
bbaabbabbababaabaaaabbaaaaaababbabbbaabaaabbaabaabaaaabaaaabbbbbbaabbabb
bbaabbaabbabaaabababbbaaaababbababbbabbbbaaaabbabaabbbaabaabbabbaabbabba
aaabbbbabbbbbbaabaabbbbb
ababbbaaabaaaaaaabaaaaab
aabaaabbbaaabbabbbbbbabababaaabbbabbabbaaaabbbbb
abababbaaaabaaaaabbaaaab
bbbaabaaaabbbbabbbaaaabababbbaaaaaaaabbbbabbbbbaabaaabbaabbababbaaaaababbbababbbaabaabba
bbbbbbabaabaaababaaabbabbbaabbbbababbaaaaaabbaaabbabaabb
aaaaaabaaaaabababaaababa
abaaabaabbaaaaabaaabbbabaabbbaabbaababbbbbbbaaab
aabbbbbabbbbbabbabbbbbab
bbabaabaaababbabbaababab
aabbaaababaaaaaabbabbabababbabbb
bbabbabbaaaaabbaabbbaaaa
aabababaaaabbaabbbbbbabb
aababbabaaabbaababaaaaba
aabbaaababbbaabaaabaabbb
bbbbaaaaabaababbabbaabaa
abaabbabbbbbabbababbbabb
baaabbbabaaabaaaabaaaaab
aaababbbaabbabaababbaabaabaabababaabbabbababbaaaaaaaababbaabbaba
aaaaaabbaababbbabbababbb
aabaabaaaaaaaabbbbbaababaabbabab
abbbbbbbaabbbababaaaaabbbababbba
babbbbaaabbbbabbaabbbbaa
baaabbabbbaabaabaabbabab
aaaabbabaabababbbbabbaab
babaabbbbbbbaabaababaabaabaaaaba
babbbaabbbababbabaabbbbb
bbaabababaaabbbbaabbbabbbbbaabbbbbbbbaba
bbbbabbbbaaaababbbabbbba
aabaaabbabbbaaababbaabab
abaababbbbbaabbaaaaabbaabbbaabaababaaabb
bbaabbbbaaabaaaaabaaaaba
aaaabaabbbaaababaababaaa
abaabaaaaabbaabbaababaab
bbaaabaabbabbbaaabaaabba
aabbbbbbabaabbabbbaababb
bbbbaaaaabbbbabaababbbab
babbbaaaaaaaaabbbabbaaab
aabbbabbaaaabbabbaabbaaa
abbbaabaaabaaabbababbaba
baaabbbaabbbbabbbabaabbbbbbbaabaabbabbbb
babbbbabbbaabbaaabaabaaabaaaababaababaab
abbbaabbbabbbabaababbaba
babbbbabbbbbabaabababbba
aaababbbababaaaaaababaaa
baaaabaababaabaaabaababbbbaaaaaabbbbbbaaaabaabba
bbababbaaababababababbbabbbbbbabbbbbaaaabaaaaabaaaabaabaababaaba
abbbbabaabaabbabbbbaaababaaabaaaabbbaabb
bababbbbbabbbababbbbbbba
baaaaaaabbaabababbaababb
bbabbaaabbabbbbaabbaaaab
bbaabbabbbbbbabbaaaaaaba
abbbbaaaabbbaaabbabbabbb
bbbabbababbbbaaababaaaab
aabbbbbbababbbbbaabbbaaa
babaabbaabaababbbbaababb
ababaaaabbaababaaabbbbab
bbbabbaabaaaabbbaabbabab
aabaaabbaaaabbabbababaababaaabbbabababab
baaabaaaabbbbaaaaabbabbbbaaaabba
bbbaabbaabbabbbabbabbbba
babaababbaaabbabbbbaabbb
babbbaabbbabbababbbbaabb
babbbabababbbbbabbaaaaab
baaabaaabaaaabaabbabaabaababaabbaaababbbabbabbabbabaaaaaabbaaaabbbaabaaa
bbabbaabbbabbabaaaababababbabbbaaaababaaaabbbbaaababaabb
baababaabbbbbbabababbaab
aaabbaababaabbabaaaaabab
aaabbbaabbbbbabbaaaaabaabbaaaabb
abbbbabaaaaababbaaabbaaa
abbbbbbaabbbbbaabaaabbbbaababaaa
baabbaabbbbabbabbabbabaa
ababbbaaaabbaaabbaabbbba
babbbbbaabbbababbbbabababbaababbababababababbbba
aabbbbbbbbabaaaababababb
aababbbabbbbabbbbbaabbba
aababbabbbbababbbabbaaab
abbabbaabaaaaaabbaabaaab
bababaaabbabababaabbabaaabbabbbbbaabaaaa
bababababbaabbaabababbba
aaaabbababbbaaaaabbaaabaababbabb
aaaabbbbaaaaaabbbaabbbbb
aaabbbbabbbaaabaaabaabba
babaababbbbbabaaabaaaaba
bbbaaabaabbaabbaabbbbbab
aaaabbababbbbaaaabbbbabababbbbbbababbaaabaabaaaa
aaabbaabbabbbbababaaaaba
abbabbaabbabbbaaaaaaaaba
abaabbaaaaaaabaaababaababaabbaaa
babbbaabbbbaaabbaaaaaaab
aaaabbbbaaaabbbbbaabaaba
ababbbaaaaaabaaaaabaabaa
abababaaaabbbbbbbaaaaaba
babbbaabaaaabaababaaaaba
bbaaabaaaaabbaababaaaaab
babbbaaabbbaaabaaababbbabbbbabbbbabaaabbabababab
bbbaaabaababbabbaabaabbb
bababaaaaabbbabaababbaba
baabababbbbaabbbabaabbbbbabbbaababbabbbaabbaabbbbbbbabbabbbbaaabbababbba
abbbabbabaaabaabbaaaabaaabbabaabbaabaabb
bbbabbbbbaaaaabbabbbabaa
bbaaabbbbbbbabaaabababab
bbababbabbbabbabbbaaaaab
aababbbabbbbabbbbbbbabab
bbbbabbabbbbbabbaabbaaba
aabbbbbaaabbabaaaaaabbaaababbbaaabbbaabbbaabbbbabbaababb
aabbabbbaabbbbabbbbaabbbabaabaabbaaaaaaaaaaababbabbaaaabaabbaabb
aabaaaaaaabaabbabababbbabaababbb
bbbbbbbbbabbababbababaabbaabababababbbbbaaababaa
aaabaaabbaaaabababbabbaaabbababaababbaab
bababaaaabbabbaabbbabaab
ababbabbbbabaabaabaababa
aabbbabbaabaaabbabbaaaab
bbaaabaaaabbaabbbbbabaaa
aabababbbaaabbabbbaabaabaaaaaaaaaaaabaabbaabaaabaaabaaba
aabaaabaaaaababbabaaabab
bbabbababbaabaabbaabaaba
aaababababababbababababb
aaaaabbbaaaabbabaabbabab
babaababbbbbaabababababb
babbbbaabbaabbbabaabaabb
aaaabbaaababbbaabbabbbba
bbbbaabbbababbbbaaabaaabbbbbaaababaaabbababaabbbbaabaabbbabbaaab
abbbbbaabaaabbbbabababaabbabaaaaababbaba
bbbbabbabbabaabaabbbabab
baaaababbbbbabbabaabaaab
aabbbbbabaaaabbbbbabbaab
bbabbabbaaaabbbbbbbabbbbabababab
aaaababbaabababbbbaabbbbbaaaaaaaabaaaaba
aaabaaaababaabbababbbbaabbababaa
baaabbabbaababbaabbbbaab
babbbbaaababaabaabbababb
baabbabaabaaaabaaaabbaaabaabaaaaaaababbaabbaabaaaabbbaaaaaaababbaaaabaabaaaabbaaabbabbbaaabbaaaa
aaaaaabbbababbbbaaaaabab
abbbabbabbaaabbbababbbbbbbbbaaabbaabbbba
aaababaabbbbabaaaabababa
babbabababaaaaaababbbbaabbabbabbbbabbbbb
bbabaaababaabaabaaaababa
aabbaabbbbbbabbababbababbabbaabbabaaaabbbaaabbaa
baaabaaaaaaaabbababbbabababbbabb
aabbbabbabbaabbaaabaaaaa
bbaaabaaabbbabbabbbbabab
abaaaaaaaabbbbbbbbbaabbb
babbbbabbabbaababaabaaab
ababaabbaabaaababaaaaaba
baababaabababbbbbbbbaabb
bbbbbaabbaababaaaaabbaaaabaaaabaaaaabbbaaaaababbbaabbbba
bababbbbaabbbabababbbbaabaabbbbbababbaaaabbabababaabaabb
abababaaaabbbababaaaabba
bbaaaaaaabaaaaaabbbbbbbb
abaabbbbaabbaaabbbbbaaabaabbabbaaabbbababbbbbbaaaabaaaaabbbbaaba
abbbabbabaaaaaabaabaabba
aabababbbaababaaabaabbbb
baababbababbbbbabbabbabbbbbbaabb
bababbaaaabbbabbaaaabbaaabaababaabbabbab
abbbbabbabbbbaaaabaaabaa
abababbaaaabaaababaaabaa
babbababbbabaaaaabbbaabbaaaaaaab
abbbbaaaabbbbbbaaabaaaab
baaaabaaaabaabaaabbaaabb
bbbbbbabbbaaabababbaabbaaabbbbba
aaaabbaaaabaabaabbaabbaaabbbabaaaabaabab
babaabababbbbababbababaa
baabbbabbaabbbaaaababbaa
bbbabbbbabbbaabaabaabbabbaaabbaa
bbaaabbbaabbbababaaaabaabaababab
babbaabaaaaaaaaaaaabaaba
bbabbabaaabaaabbabbaabab
babaababaaabbbbaabbaaabb
aaaabbabaabbbbbabbbabbba
bbbbbbabaaabbaababbaababbaabbbbbabbbabbb
aaabbbbabbbbbabbaabbbaab
bbbbaabaaaaabaababbabaab
abbbaabbbbbbabaaaabaabba
aabbbbaaabababbbaaaabababbaabbbaaabaabbb
aaaaabbaaaaabbababbbaaaa
aaababbbabaaaabbbbabbaaaababaaabaabaaaab
aaabaaabbbaaabbbbbbaaaab
abaabaabbbbaababbbbaabaa
aaaabaaaaababbababababab
aabababbbabbbbbbabbbbabb
abbabbaababaabbbabbaaaaa
babbbaabababaabbbbbaababbaaaaabbbbbbbbbaaabbaaaa
baaabbbbaaabababaaabaabb
baaabbbabbbaaabaaabaaaaa
ababaababbbbabaaabaababa
ababbbaaaaabbbbabbababbb
abbbbabaabbbbbaabbaaaaab
baaabbbbababbabbaaabbaaa
babaabaaaaabbbbabbbbbbaaabaabbabbbbbabbaabaabbbaabaaabbb
bbbababbabbbbbbabbbbabaaaabaaaabbaababbb
bbabbabbbbbaaabaaabbabaabaaaabba
baaabaaabbaaabbbaaaaaababbabbbab
aaaabbabaaaabaababaaabbb
abbbbbbaabbaaabbabbbaaaaabbbabab
abaabaaabbbbbbaaaababaaa
aaabbaabbabbbaaaabbaaaba
baaaaaabaaababaaababaaaabbbbaabb
baaabaabbbaabbaaabbbabababbbbaab
aaabbbaabbbbabbababbbabb
baaaababbaababaaaaaababa
babbbbbababbbaabbabbaaab
babbaabbbbabbaaabbbaabababaababbbabbbbbbbbbabaabbaaaaabaababaaba
ababbaaaabbabaabbabbabbbbbbaaaab
bbaaaaaabbbabbaaaaaaaaba
bababaaabaaabbbabbbbbbbb
bbabaaaabbbababbbabbaaaa
bbaabbbbbaaabaaaababaaab
babbabababaababbbbbbbaaa
aabbaabbabbabbbaaaaaabaaabbaabbaaabaaabbbbabbababaabaabb
aabababbbaaaabbbaaaaabab
aabbbbbbabbbbabbaababbbababaabbbabbabbbbbaabbabaabababbb
bbbbaaaaabaabbaaaababbaa
baaabbababbbaaabaaabbaba
aaababbbbabbbbbaaaabaaba
babaabaaabababaaaaaababbabbbabab
bbaaaaaaaabbabbbaaababba
bbbabbbbbaaaaaaabababbaaabababbaaabbbbbabaababaabbaaaaab
bbabbbaaababaaaaabaaabbb
ababbbaaababbbbbaabaaaaa
abaababbbabbbaaaababbaaa
bbbbabbbaaaaaaaababaaaba
bbbbaabaabababaabbbbaabb
bbaabaababbbbbabaaaaabababbabbababababbbabbababbbbaababb
aaaaabbaabababbabaaaaaba
aabbbabbbbabbabaaaabbabb
babbbaabaaabaaabbbaabbabbaaabbbbbbbbbabbbabbbabbbaabaababbabbaaa
babbbaabbaaaabbbbaabbaba
bbbababbbaaabbbbabbbbbaaababbbbb
bbbbabbbbbabbabaaaaaaaaaabaaabaaaababaaa
ababaabbaaaaabbaabbababb
bababaaaabababaaababbaaa
bbaabababaaabbbbabbaaabb
babbbbaabaaabbbaababaaababaaababbbbbbbbbbaaababb
baaaabbbaaabbaabbabbbabb
ababbbbbbbbbaaaaaabaabab
abaabbaabaaaababbaabbabb
aaaabbababaabaaababaaaab
baaabaabaaaabaababbbbbbbbabbaaaa
bbabaaaabaaabbbbaababbbb
ababbbaaaabbaabbababbbbbaababbaaaaababba
abbaabbaaaaaaabbbabaaaaa
aaabababababbabbbabaabbabbabbaababbbbabbbabaaabbaaababab
bbaababaaaaabaaaabaaaaab
bbbaababbabbbbbabbbbaabb
ababaaaabbbbbabbaaabaabb
babaabbabbbbabaabbbaaaaa
bbbbbabbabaabaaaabbbabaa
baaaaaaaaabababbbbbaaabaaabbbabbabbababa
baaabbbabbaabaabbabbababbaaaabbbbbbbbbbabababbab
abbbabbabababababaabaaab
babaababbbaabababbabaabb
bbabababaabababbbaabbabb
abaabaabbabaabaaabbaabaa
aaaaabaaabbaabaaaabbaaababaaabbababbbabbabaabbbbbabbbbbabbaaabababbabbaa
bbabbbabaababababbbbabbbabaaaabbabbbabbaaabaaaaaababbaba
babbbbaababbbbaaabbbaaabababbbababaaaabb
bbbabbbabbabbbababbabaab
aababbabaaabbbbaaabababa
abbbbababbabababbababbab
aaabbaabaabaabaaaabaabab
bbabbaabbabbbbbbabaaaaaa
bbabbbaabaabbaabbaabbbbb
bbbbbabaaabaaabaababbbbababbbbaabaaaabbaababaaaabbababaabbaaabbababbbbbb
bbaabaaaababbabbababbbaaaaabbabbaabbbabaaaabaaabbabbbbbabaababbbabbbbbbbbabbbbba
aaabbbbaabaabbaaabbaabbb
ababbbaaaabbbaaabbbbbaaabbbbabbbbabbabbaababbaabaaababba
bbaaaaaaababaaaabbaaaaaabbaaabaabbbabaaaababbabaaabaabba
abbabbaaaaaaaaaaaaaababa
abbbbabbbaaabbabbbaabaaa
bbbbaaaabaaabaaabbbabbbaabbabaab
bbaaababbbbabbbbbbbabaab
abbaaabbbbaaaaababaaaabaabbaaaaabaabaaab
aaabaaaababbbbaaaababbaa
aabbabbbbbbaabbabbabbbab
aabaabbbabaababaaabaaaaa
abbbbbbbaaabaaabaaabbaab
babbaabababbbaababababab
babbbaaabbababbabaabbbaa
abbaabbaabbbaabbaababaaa
ababbbbbbbbbbbaabaaababb
ababaaaabbbbabaabbaaaaaabbbaabaababababb
baababaaaaababbbbaaaabbbbabbababbaabbbaaaabbaaaa
baaaababbbbbbabbbabaaaab
bbabbabaabaababbababbbab
abbbbabaabaabbbaabbabbbbbaabbbab
ababaaaabbaabbbbbabaaaba
aabaaababaababbaabbaabab
babaabaaabaababbabbaabbb
aabbbbbaababaabaabbaabaa
baaaaabbbababbaaabbbbbbbbbabbaabaaabaabbaaaaaaabaaabbabb
babababaaaaaaaabbbabaabaaaabbaaaaabaaabbbbababbbababaabbaaaabbbababbbaab
babbaababbbbaababaabbaba
baabbabaaaabbbbbbababbabbaabaaab
abaabaaabaaabbbbbabbabba
ababbbaaaabaaabbabbbbbab
aaaaabbbabbaabbaaabaaaab
aaaaaaaaabbbbbabbaaabbabbaaabbaabaaabbabaaabaaaa
aaaabbbbbbaabaabbababbab
babababaabbbbaaaaaaaabbabbbbbababababbab
aaaabbaaabbbaaabbaabaaab
aabaabaaabbabbbbaababbaa
abaabbabbabbbababbbbaabaabbbbbaabaaabbbbabbaaaaa
bababbbbbbbbabbbbbbaababbabbbbaaaabbababaababbbb
aaabbbaabbaabbbbabaaabbb
abbbbbbbbbbaaabbababbabbbaababbb
aaaabbbbaabbabaaaaabbabb
aaaaaaaababbbaaaabbaaabb
babbaabababbaababbbabbbbbbbbbbababbbababbaabbbab
bbbbabbbbbbabbaabaabbbaa
aaaabbbbbaaaaaaabaaaaabbabaababaabbbbbab
aabaaaabbbabbbaaabbabbaaabaabbbaaabbbbbaababbabbabaaabbabbabaaab
aaaaabbbbbabaaaaabaabbbb
abaabbaabbbbbbaabaaabbaaababaaababbaabababbbbbbbbaabbbbbbabbaaabbabbbbaaaaaaaaaaaaaabbaa
abaabaabbbaaabbbabbaabaa
babbaababaabbbaabbabbaaa
ababaaabbababbaabaaaaababaabbababbabaabbabbbbaaa
aaababaabbbbbabbbbabaaaaaabaaaaaaabaabab
aaaababbbbbbaababababaabaababaaaababbaba
bbabababababbbaaaababaaa
bbbabbabaabbbabbabbaabaa
bbaaababbaaabaaabaaaaababbabaabb
abbbbbbabaababaaaabbabab
aaaaaaaaaaaabbaabbabbaab
bbaaabaaabbabbbaaababaab
aabaabbbbbbbbbbbbbbabbabaabaabbbabbbbbbbbbbbababbbbbabbbaabbbbbabaaaabba
abbabbbababbbbbaaabbbaaa
baaaaaaabbabaaabbaaabaabbbabaabbabbbaabbabbababbaaabababbabaabbaaabbbbabbababbab
bbbaabaaabaabbabaabaabbabbbbaaaa
baaabaaabaababaaabbbbaab
bbbabbaababbbbababaaaaab
aabbabbbaaabaaaabbbaaaab
aabbabbbaabbaabbabaabaaabaaabbbb
bababbbbaaaaabbbaabababbabaabbaababbbaabbabbababaaaaaaabbaabbabbbaabaabb
bababbaaababbabbbbbbbbba
aabababbaabaabaabbabbbbb
bbbbaaaaabbbaababbbbbbba
bababbbbaabbbbbbabbabaaa
aaabbbaaaaaaaaaabaabbbab
baaaabaaaaababbabbbbababbbbaaaab
abbbaabbbabbbaaabbbabbba
baaaabbbbabababaabbbabab
aaaabbabbbababbababbbbaa
aabbabaabababaabaababbbb
bbaabbbbbaabbaababbabbbaaaaabbabbbaabbba
abbbaaabbaaaabbbaabbbaab
abaaababbaabbbabababaaabbbbaabbbbbaaaaabaaabaaabbabbbabaabaaababbabaabbbaaaabaaa
aaaaabbbbababaababbaaabb
bbbbabbbbbabaabaaabaabab
aababbbababaabbbbbaabbabaabaaabbbabbaaabaaababba
abaaababababbbbbbbababbabbbabbbbbbaabbbabaaabbbbabaababb
babbbbbababaabaabbbaabaa
bbbbabaaaabbbabaaaabaaaaabaaaaaabbababaa
baaaabaabaaaabbbbabaabbaaabaaababbababbabaababab
babbbbababbbbaaaabaaaabb
abaabaaaabbaabbabaaaabbbbaaaaaababbaaabb
abbaabbabababbaaaaababba
abbbbbbababaabbabbbbbaab
bbaabaabaababbbabaaabbabbbaaaaaababaabbbbbbbabbaababbababaabbaaa
bbaaababbbbababbbbbbaaaaaabaabbb
bbbbbbaaabbbaabaabbbabbabbbbbbaaabbababbabbbabbbbbabaabb
aabbaababbaababbbbbababbaabaababbbaaabbaababbbbbbbbaaaaabbabaaaababbabbaaaababbbaababbbb
babbbaabbbabbabbbbbabbaaabbaaaab
abbbaabbbbabababaaaababa
ababbbbbbbaaabaabbbabaaa
ababaaaabaaaabababbaaaba
ababbbbbbababababbbbbabbbbbbbbba
bbbbabaabaaabaaabaabaaaa
baaabaaaabbbbabbababaababbbbbbbbababbbababaaaaabbbbbabababbbbbaaaabbbbababbaaaaaaabaaababaaaabbb
bbabababaaaabaaaababbbba
bbaaababaaaabaabaabababa
bbbaababaaabbaabbbaabbba
aabbabbabaaaabbabbaaaaabaabaabaababbbbabbaabaaba
abbabaabbbbabbbabbaabbaabbabbaabbababbbb
abbbabbaabbbbabbbbbbbaab
aabababbabaababbbaaabbbbaabaaababbaabaaa
bbbbaababbbabbababbbabbb
aaabaaabbbaabbbbbababbba
bbbbaabaaaabbabababaaaaabaaaaabaabaaabab
baabbaabababbbaababbaaaa
aababbabaaababbbbaaabbaa
babbbaaaababbbbbabababbb
baaaaaaabbbaababababbbba
bbabaaabaaaababbbaabaabb
abbbbabaabbbabbabbaaaabb
bbaabbbbbabbbaabbaaabbaa
aaabbbaaabbababbbbbbabab
babbbaaaaabbabbbabbaababbbaaababaabababbabbababaaaabbbbbaaabaabb
bbbababbbbbbbbbaabbabbababaabbbb
baaaabaabaaabbabbbbbaaab
babababaabbbbaaabbbababbaaabaabb
babaaaabbabbabbbbbbabbaaababbaaabaabaaaabaaaabaabbabbbaabaabaaabababbbbb
bbabbbaababbbbbaaaaaaaaaaabbabab
bbaabaaaaababaabaabbababaaaabbababaabbbabbbaaababbababbb
abaabaabbbbbbbabaabbabab
aabbbababbaaababababbbab
bbbabbaabbbababbbaabaabb
babbbabaaabbbbbaaaaaabbaababbbbbabaaabbb
aaaabaaabbaaabbbababbabbbaaabaaabbbaababaabbbbaa
abbabbbbbaaaababbaabbaba
bbaaabbbbbbabbabababaabaabbbbbbbababbbba
bbbbabbaaabaaabbaaabaaabaabaaaab
babaababbabaabbbaaababba
ababaababaaaaaabbbbabbabaabbbaaa
babaababbbbbabababaaabbaabaaaaaaabbabbababbabaaaabbbabbbaaababaabbbbabaaababbbba
abbabbaaaaaaabbabbbbaaaaaaaaabbb
aaabbbbabbaaabbbaababbabbabababb
bbaabbbbababbaaababaabaabbbbabaaaaabaabbaaabbbaaaaaabbba
ababaabaaabbababbbbabbabaabaaabbabaabbab
baaabbbabbbbabbababbbbabaaabaabaaaaabbab
baaabbbbaabababbbaaaaaba
abbabbbbbaaaaabbaabaaaaa
bbabbabbaaaaabbabbabbbaabaabbbaa
bbbaaabbaaaababbbabbbbaababaabbbabbbaaababababbb
babbaabbbbbbaaaaababbaba
abbaabbaabababbaaababbbb
abbbbaaabaabaabbaabbabbabaaabaabbbabbaabaaabababaabbaabababaaaaabbbbabbaaabbbbaaaaaaabba
abbbaabaabbbbbbbbabaaabb
baaaabaabaaaaaabbaaaababaaabbbaabbbabbabbaabbaaa
baaabbabbbabaaaaaabaabba
babbaabaaaababbbbaaaabaabaababab
bbbaabbababaabbbbaaaaaba
aaaabbbbaabbabaaabaaaabb
aabbaaabbbbaabbaaaabbaab
babbbaaaababaaaabbababbb
abbabbaabbabaaabbbaaaabb
aaaabaaababbbbbbaababbbb
bbbababbbbaaaaaaabbbbbab
bbbabababaabaaabbaaabaaabbabaababbabbaabbbaababbaaaabaababbbabab
aaaababbbaaabbababababbb
abaabbabbbaaabaabbaaaaabbbaabbbbbbabaaaabbaaaaabaaaabbabbbababaa
bbaabbaaaabaabaabbaababb
aaababbbbaaaaabbbbaaababaabaaaaaabbabaab
bbbabbbbbaaaaaababaabbbaaaababbbbaabbbaaabbbabaa
babbbbbbaaababbbbaaabbaa
bbbbabbaabaaaabbaabbabaababbbbaabbabaabbbbbabbbbabbabaab
baaaabbbaaaaabbbaababbaa`

// AOC_2020_19 is the entrypoint
func AOC_2020_19(cli *goutils.CLI) {
	AOC_2020_19_part1_attempt1(cli)
}

func AOC_2020_19_part1_attempt1(cli *goutils.CLI) {
	rre := NewRegexRuleEngine(DAY_19_INPUT)
	rre.ParseRules()
	total := rre.Apply("0")

	fmt.Printf("133 rules, 471 messages? : %v, %v\n", len(rre.Rules), len(rre.Messages))
	// rule := rre.Rules["0"]
	// for _, message := range rre.Messages {
	// 	if rule.IsMessageValid(message) {
	// 		total++
	// 	}
	// }
	fmt.Printf("%v messages pass.\n", total)

}

type RegexRuleEngine struct {
	Rules    map[string]*RegexRule
	Messages []string
}

// Init parses all Rules to create RegexRules
func (r *RegexRuleEngine) ParseRules() {
	// first find our literal rules
	for _, rule := range r.Rules {
		rule.Value = strings.ReplaceAll(rule.Value, "\"", "")
		if rule.Value == "\"a\"" || rule.Value == "\"b\"" {
			regex := strings.ReplaceAll(rule.Value, "\"", "")
			rule.Regex = regex
			rule.Evaluated = true
		}
	}

	fmt.Printf(" >>>>>>> RegexRule: before >>>>> \n\n")
	r.Debug()
	fmt.Printf("\n <<<<<<< RegexRule: before <<<<< \n\n")

	// 	0: 4 1 5
	// 1: 2 3 | 3 2
	// 2: 4 4 | 5 5
	// 3: 4 5 | 5 4
	// 4: "a"
	// 5: "b"

	// now go over each rule until we have a fully evaluated sequence
	totalToEvaluate := len(r.Rules)
	for {
		// loop over until we evaluate everything
		totalEvaluated := 0
		for _, rule := range r.Rules {
			if rule.Evaluated {
				totalEvaluated++
			} else {
				// this requires evaluation
				// split this rule to its sub-rules
				// for any sub-rule that is evaluated, replace the literal with the sub rule regex
				// newValue is the value that contains e.g. 4 5 | 45 45
				// for each evaluated subrule, replace it with the evaluated content
				// newValue := rule.Value

				// the subrules are an array of [ 4, 5, |, 45 ,45 ]
				subRules := strings.Split(rule.Value, " ")
				fmt.Printf("starting value is : '%v'\n", rule.Value)
				fmt.Printf("subrules are      : '%v'\n", subRules)

				// we will go over all subrules and replace any that are evaluated
				// then we will rebuild the rule based on the subrules
				allSubRulesEvaluated := true
				changes := 0
				for subRuleIndex, subRuleKey := range subRules {
					if subRuleKey == "|" {
						continue
					} else if subRuleKey == " " || subRuleKey == "" {
						continue
					} else if subRuleKey == "a" || subRuleKey == "b" {
						continue
					} else {
						// fmt.Printf("rule[%v], SubRules %v SubRule %v\n", rule.Key, subRules, subRuleKey)
						subRule, exists := r.Rules[subRuleKey]
						if !exists {
							fmt.Printf("subRule '%v', exists=%v\n", subRuleKey, exists)
							continue
						}
						fmt.Printf("subRule '%v', exists=%v, subruleValue=%v\n", subRuleKey, exists, subRule.Value)
						if subRule.Evaluated {
							// then this sub-rule has been fully evalulated; we can replace
							// the reference to it with the evaluated content
							// if the ruleKey was 4 and I had others like 4 44 444
							// then I need to e "careful" about how I replace using this key
							// splits := strings.Split(newValue, " ")
							// for _, splitValue := range splits {
							// 	if splitValue == subRuleKey {
							// 		splitValue = subRule.Value
							// 	}
							// 	newValue += " "
							// 	newValue += splitValue
							// }
							// subRuleKey = subRule.Value
							subRules[subRuleIndex] = subRule.Value
							changes++

							// newValue = strings.ReplaceAll(newValue, subRuleKey, subRule.Value)
							// rule.Value = newValue
						} else {
							allSubRulesEvaluated = false
						}
					}
				}

				if changes > 0 {
					// make a string from the subRules as they were changed
					newValue := ""
					for _, value := range subRules {
						newValue += " "
						newValue += value
					}
					rule.Value = newValue
				}

				// now rebuild the subrules to a single value for the rule

				if allSubRulesEvaluated {

					// there seems to be a bug where I am including a number... I don't know why
					testValue := rule.Value
					testValue = strings.ReplaceAll(testValue, "(", "")
					testValue = strings.ReplaceAll(testValue, ")", "")
					testValue = strings.ReplaceAll(testValue, "a", "")
					testValue = strings.ReplaceAll(testValue, "b", "")
					testValue = strings.ReplaceAll(testValue, "|", "")
					testValue = strings.ReplaceAll(testValue, " ", "")
					if testValue != "" {
						fmt.Printf(">>>>>>>>>> INVALID RULE %v >>>>>>>\n", testValue)
						fmt.Printf("%v\n", rule.Line)
						fmt.Printf("%v\n", rule.Value)
						fmt.Printf("<<<<<<<<<< INVALID RULE %v <<<<<<<\n", testValue)

						// fmt.Printf("%vINVALID RULE \n", rule.Debug())
						// fmt.Printf("<<<<<<<<<< INVALID RULE <<<<<<<<\n")

					}

					rule.Evaluated = true
					rule.Regex = strings.ReplaceAll(rule.Value, " ", "")
					if len(rule.Regex) > 1 {
						rule.Regex = "(" + rule.Regex + ")"
					}
					rule.Value = rule.Regex
				}

			}
		}

		if totalEvaluated == totalToEvaluate {
			break
		}

		fmt.Printf(" >>>>>>> RegexRule: during >>>>> \n\n")
		r.Debug()
		fmt.Printf("\n\n <<<<<<< RegexRule: during <<<<< \n\n")

	}

	fmt.Printf(" >>>>>>> RegexRule: after >>>>> \n\n")
	r.Debug()
	fmt.Printf("\n\n <<<<<<< RegexRule: after <<<<< \n\n")

}

func (r *RegexRuleEngine) Debug() {
	for index := 0; index < 10000; index++ {
		sindex := fmt.Sprintf("%v", index)
		rule, exists := r.Rules[sindex]
		if !exists {
			break
		}
		line := rule.Debug()
		fmt.Printf("%v\n", line)
	}
}

func (r *RegexRuleEngine) Apply(ruleId string) int {
	rule := r.Rules[ruleId]
	total := 0
	for _, message := range r.Messages {
		if rule.IsMessageValid(message) {
			total++
			fmt.Printf("REGEX PASS '%v'\n", message)
		} else {
			fmt.Printf("REGEX FAIL '%v'\n", message)

		}
	}
	fmt.Printf("REGEX VALUE %v\n", rule.Regex)
	return total
}

func NewRegexRuleEngine(input string) *RegexRuleEngine {
	splits := strings.Split(input, "\n")
	messages := make([]string, 0)
	ruleMap := make(map[string]*RegexRule)
	useMessages := false
	for _, line := range splits {
		if strings.TrimSpace(line) == "" {
			useMessages = true
			continue
		}
		if useMessages {
			messages = append(messages, line)
		} else {
			rule := NewRegexRule(line)
			ruleMap[rule.Key] = rule
		}
	}
	return &RegexRuleEngine{Rules: ruleMap, Messages: messages}
}

/*
const DAY_19_TEST_INPUT = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
`
*/

type RegexRule struct {
	Line      string // the original line    1: 2 3 | 3 2
	Key       string // the rule id 1
	Value     string // the 2 3 | 3 2
	Evaluated bool   // indicates if the regex has been evaluated
	Regex     string // the regex for this rule
	FullRegex string // the "full" regex composing all regexes (after evaluation)
}

func (rr *RegexRule) Debug() string {
	line := fmt.Sprintf("%v : %v : value=%v, regex=%v", rr.Evaluated, rr.Line, rr.Value, rr.Regex)
	return line
}

func NewRegexRule(line string) *RegexRule {
	splits := strings.Split(line, ":")
	key := strings.TrimSpace(splits[0])
	value := strings.TrimSpace(splits[1])
	value = strings.ReplaceAll(value, "\"", "")
	rr := RegexRule{Key: key, Value: value, Line: line, Evaluated: false}
	return &rr
}

func (rr *RegexRule) IsMessageValid(message string) bool {
	expr, _ := regexp.Compile("^" + rr.Regex + "$")
	return expr.MatchString(message)
}

/*
ParseRules
	for each rule, create a regex
	scan list and find literals and create a rule for them
		"a" becomes a



	for each ruleId
		if isLiteral? ("a", "b")
			rule_regex = "a"
		else:
			0: 4 1 5
			1: 2 3 | 3 2
			2: 4 4 | 5 5
			3: 4 5 | 5 4
			4: "a"
			5: "b"

			total_regex = '^regex$

	find literal rules first


	expect 1:1 Rule
	a
	Rule.IsValid(message)
		rule has a regex
		regex.Match(message)

*/
