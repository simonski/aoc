package aoc2018

/*
Describe the problem
*/

const DAY_2018_10_DATA_TEST = `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>`

const DAY_2018_10_DATA = `position=<-41214, -10223> velocity=< 4,  1>
position=< 41635, -10215> velocity=<-4,  1>
position=<-51569,  20845> velocity=< 5, -2>
position=<-30826,  20850> velocity=< 3, -2>
position=< 31249, -30930> velocity=<-3,  3>
position=<-30862, -20578> velocity=< 3,  2>
position=<-20528, -30930> velocity=< 2,  3>
position=< 52013, -51642> velocity=<-5,  5>
position=<-20491,  41552> velocity=< 2, -4>
position=<-41218,  31200> velocity=< 4, -3>
position=< 31271, -10216> velocity=<-3,  1>
position=< 52010,  51915> velocity=<-5, -5>
position=< 41634, -30933> velocity=<-4,  3>
position=<-30826,  41551> velocity=< 3, -4>
position=<-20499,  31199> velocity=< 2, -3>
position=< 20953,  41559> velocity=<-2, -4>
position=<-51576, -41285> velocity=< 5,  4>
position=<-20523,  51906> velocity=< 2, -5>
position=< 10593,  51914> velocity=<-1, -5>
position=<-51593,  31204> velocity=< 5, -3>
position=< 20893, -30930> velocity=<-2,  3>
position=< 20929,  20843> velocity=<-2, -2>
position=< 51985,  20845> velocity=<-5, -2>
position=<-51534, -10215> velocity=< 5,  1>
position=<-41222,  51908> velocity=< 4, -5>
position=<-30859,  51914> velocity=< 3, -5>
position=<-30874,  51906> velocity=< 3, -5>
position=<-10129, -10220> velocity=< 1,  1>
position=<-41198,  31203> velocity=< 4, -3>
position=<-41238,  20847> velocity=< 4, -2>
position=<-10141,  20850> velocity=< 1, -2>
position=< 41661,  51915> velocity=<-4, -5>
position=<-41230,  31202> velocity=< 4, -3>
position=< 10561, -20578> velocity=<-1,  2>
position=<-30835,  20848> velocity=< 3, -2>
position=<-20467, -30926> velocity=< 2,  3>
position=<-51561,  31202> velocity=< 5, -3>
position=< 20900, -41280> velocity=<-2,  4>
position=< 10588,  31198> velocity=<-1, -3>
position=<-30854,  41551> velocity=< 3, -4>
position=< 31263,  10488> velocity=<-3, -1>
position=<-10117,  10490> velocity=< 1, -1>
position=<-51569,  20843> velocity=< 5, -2>
position=< 20928, -51640> velocity=<-2,  5>
position=< 10537, -10221> velocity=<-1,  1>
position=<-41201, -41286> velocity=< 4,  4>
position=<-10141, -20572> velocity=< 1,  2>
position=< 20908, -30929> velocity=<-2,  3>
position=<-20512,  51911> velocity=< 2, -5>
position=<-30835, -10224> velocity=< 3,  1>
position=< 41610, -10219> velocity=<-4,  1>
position=< 10537,  20849> velocity=<-1, -2>
position=< 41644, -20579> velocity=<-4,  2>
position=< 51990, -41289> velocity=<-5,  4>
position=< 10569, -20571> velocity=<-1,  2>
position=< 20905,  41559> velocity=<-2, -4>
position=< 20940,  20843> velocity=<-2, -2>
position=<-20512,  20843> velocity=< 2, -2>
position=<-41202,  51915> velocity=< 4, -5>
position=<-10173,  51911> velocity=< 1, -5>
position=<-41230, -10220> velocity=< 4,  1>
position=< 31279,  31205> velocity=<-3, -3>
position=<-41226,  10492> velocity=< 4, -1>
position=<-51577, -41280> velocity=< 5,  4>
position=<-10125,  20846> velocity=< 1, -2>
position=< 41658,  20845> velocity=<-4, -2>
position=< 31274,  41555> velocity=<-3, -4>
position=<-10117,  51913> velocity=< 1, -5>
position=< 31287,  51910> velocity=<-3, -5>
position=< 10545, -20576> velocity=<-1,  2>
position=<-20520,  51913> velocity=< 2, -5>
position=<-20528, -10216> velocity=< 2,  1>
position=<-41180, -20579> velocity=< 4,  2>
position=< 10561, -41283> velocity=<-1,  4>
position=< 31252,  51909> velocity=<-3, -5>
position=<-30822, -30927> velocity=< 3,  3>
position=< 20929, -41281> velocity=<-2,  4>
position=<-51553,  10489> velocity=< 5, -1>
position=<-30883, -41287> velocity=< 3,  4>
position=<-30846,  10487> velocity=< 3, -1>
position=< 10594,  10495> velocity=<-1, -1>
position=< 10547,  31200> velocity=<-1, -3>
position=< 41658, -10216> velocity=<-4,  1>
position=<-51585,  41553> velocity=< 5, -4>
position=<-41198, -20579> velocity=< 4,  2>
position=<-10168, -10223> velocity=< 1,  1>
position=<-30851, -20575> velocity=< 3,  2>
position=< 10574, -41283> velocity=<-1,  4>
position=< 31260,  20849> velocity=<-3, -2>
position=< 41607, -51636> velocity=<-4,  5>
position=< 20944,  10494> velocity=<-2, -1>
position=< 31255, -10223> velocity=<-3,  1>
position=<-20510,  51906> velocity=< 2, -5>
position=<-30851, -41286> velocity=< 3,  4>
position=<-41198, -51636> velocity=< 4,  5>
position=<-41190,  41553> velocity=< 4, -4>
position=<-10170, -30930> velocity=< 1,  3>
position=<-41233,  20848> velocity=< 4, -2>
position=< 31255,  31197> velocity=<-3, -3>
position=<-30843,  10488> velocity=< 3, -1>
position=< 20948,  10492> velocity=<-2, -1>
position=<-41206, -51639> velocity=< 4,  5>
position=<-10161, -30928> velocity=< 1,  3>
position=< 31308,  20848> velocity=<-3, -2>
position=<-10149, -30934> velocity=< 1,  3>
position=<-10165,  20841> velocity=< 1, -2>
position=< 20948, -51642> velocity=<-2,  5>
position=<-41222, -51637> velocity=< 4,  5>
position=<-30823,  51914> velocity=< 3, -5>
position=< 41642, -20574> velocity=<-4,  2>
position=<-20523, -41280> velocity=< 2,  4>
position=< 10589,  41559> velocity=<-1, -4>
position=< 10577, -41283> velocity=<-1,  4>
position=< 20900, -30934> velocity=<-2,  3>
position=<-30851, -51635> velocity=< 3,  5>
position=< 31249,  10490> velocity=<-3, -1>
position=< 51992, -41280> velocity=<-5,  4>
position=<-30833, -20576> velocity=< 3,  2>
position=<-20509, -41285> velocity=< 2,  4>
position=< 41607,  51907> velocity=<-4, -5>
position=<-20485, -51639> velocity=< 2,  5>
position=<-20523,  31196> velocity=< 2, -3>
position=< 10585,  41558> velocity=<-1, -4>
position=< 41653, -30927> velocity=<-4,  3>
position=< 20900, -41285> velocity=<-2,  4>
position=< 31276, -20573> velocity=<-3,  2>
position=<-10117,  41555> velocity=< 1, -4>
position=<-10133,  51910> velocity=< 1, -5>
position=< 31249,  10490> velocity=<-3, -1>
position=< 31295,  10492> velocity=<-3, -1>
position=<-30843, -20570> velocity=< 3,  2>
position=<-30827, -20572> velocity=< 3,  2>
position=< 10586, -51639> velocity=<-1,  5>
position=< 10589, -10223> velocity=<-1,  1>
position=< 20943,  20848> velocity=<-2, -2>
position=<-10157, -10218> velocity=< 1,  1>
position=< 10598, -41284> velocity=<-1,  4>
position=< 41660, -30934> velocity=<-4,  3>
position=< 31263,  20849> velocity=<-3, -2>
position=< 31276,  20846> velocity=<-3, -2>
position=<-41238,  20846> velocity=< 4, -2>
position=<-10132, -41285> velocity=< 1,  4>
position=< 20924,  51906> velocity=<-2, -5>
position=<-20504,  51908> velocity=< 2, -5>
position=< 52007, -20576> velocity=<-5,  2>
position=< 51957,  41555> velocity=<-5, -4>
position=< 31291, -10224> velocity=<-3,  1>
position=<-10161, -10217> velocity=< 1,  1>
position=<-30847, -20575> velocity=< 3,  2>
position=< 41615, -41287> velocity=<-4,  4>
position=< 41626,  20850> velocity=<-4, -2>
position=< 31307, -10224> velocity=<-3,  1>
position=<-20470,  51915> velocity=< 2, -5>
position=<-51593,  31198> velocity=< 5, -3>
position=<-30873,  31196> velocity=< 3, -3>
position=<-20504, -10220> velocity=< 2,  1>
position=<-10157, -20571> velocity=< 1,  2>
position=<-10156,  41555> velocity=< 1, -4>
position=< 51970, -30933> velocity=<-5,  3>
position=< 41626,  41557> velocity=<-4, -4>
position=< 52009, -20578> velocity=<-5,  2>
position=< 51994, -20577> velocity=<-5,  2>
position=< 41626,  10491> velocity=<-4, -1>
position=<-41193,  10494> velocity=< 4, -1>
position=<-51545, -51635> velocity=< 5,  5>
position=< 31279, -20571> velocity=<-3,  2>
position=< 41611,  31200> velocity=<-4, -3>
position=<-41182,  41552> velocity=< 4, -4>
position=< 52008,  10493> velocity=<-5, -1>
position=< 31273,  51910> velocity=<-3, -5>
position=<-20512, -10215> velocity=< 2,  1>
position=<-30843, -51643> velocity=< 3,  5>
position=<-30849,  10486> velocity=< 3, -1>
position=<-20502,  10490> velocity=< 2, -1>
position=<-20496, -51640> velocity=< 2,  5>
position=< 41615,  10489> velocity=<-4, -1>
position=<-20485, -41284> velocity=< 2,  4>
position=<-20520,  20849> velocity=< 2, -2>
position=< 51986, -30929> velocity=<-5,  3>
position=<-51564,  31198> velocity=< 5, -3>
position=<-30841,  10490> velocity=< 3, -1>
position=<-51573,  20841> velocity=< 5, -2>
position=< 10589, -41281> velocity=<-1,  4>
position=< 20900,  31199> velocity=<-2, -3>
position=<-41225,  20850> velocity=< 4, -2>
position=< 20905, -30932> velocity=<-2,  3>
position=< 41663, -30933> velocity=<-4,  3>
position=<-30833, -51641> velocity=< 3,  5>
position=<-41190, -30934> velocity=< 4,  3>
position=< 31291, -51637> velocity=<-3,  5>
position=<-10132, -10220> velocity=< 1,  1>
position=< 41638, -10220> velocity=<-4,  1>
position=<-10165,  10486> velocity=< 1, -1>
position=<-20504,  41552> velocity=< 2, -4>
position=< 20929,  41559> velocity=<-2, -4>
position=<-10124,  31200> velocity=< 1, -3>
position=< 31279,  41556> velocity=<-3, -4>
position=< 20908, -51637> velocity=<-2,  5>
position=<-30867,  31199> velocity=< 3, -3>
position=< 41628, -51640> velocity=<-4,  5>
position=< 52001,  41555> velocity=<-5, -4>
position=<-10139, -10215> velocity=< 1,  1>
position=< 51957, -10224> velocity=<-5,  1>
position=< 52018,  51912> velocity=<-5, -5>
position=< 20892, -20572> velocity=<-2,  2>
position=<-20520,  31204> velocity=< 2, -3>
position=<-51585,  10488> velocity=< 5, -1>
position=< 31251,  51910> velocity=<-3, -5>
position=<-41204,  10490> velocity=< 4, -1>
position=<-51566,  20845> velocity=< 5, -2>
position=< 41623,  41553> velocity=<-4, -4>
position=<-41233, -51641> velocity=< 4,  5>
position=<-41194,  31203> velocity=< 4, -3>
position=<-10112, -20570> velocity=< 1,  2>
position=< 10574,  10488> velocity=<-1, -1>
position=< 41646, -20573> velocity=<-4,  2>
position=<-51560,  31200> velocity=< 5, -3>
position=<-41214, -41281> velocity=< 4,  4>
position=< 31308, -51635> velocity=<-3,  5>
position=<-10130,  10486> velocity=< 1, -1>
position=<-51588, -41285> velocity=< 5,  4>
position=< 41643,  10486> velocity=<-4, -1>
position=<-51588,  31197> velocity=< 5, -3>
position=< 31284, -10218> velocity=<-3,  1>
position=< 31295,  31199> velocity=<-3, -3>
position=<-20509,  41555> velocity=< 2, -4>
position=<-51588,  41558> velocity=< 5, -4>
position=<-20486,  31196> velocity=< 2, -3>
position=< 10569, -41288> velocity=<-1,  4>
position=< 10593,  51907> velocity=<-1, -5>
position=< 20925,  31200> velocity=<-2, -3>
position=<-20480, -30929> velocity=< 2,  3>
position=<-30870,  20842> velocity=< 3, -2>
position=<-30843,  31202> velocity=< 3, -3>
position=< 41631, -41288> velocity=<-4,  4>
position=<-30854,  41559> velocity=< 3, -4>
position=<-30867, -20575> velocity=< 3,  2>
position=<-30859,  20844> velocity=< 3, -2>
position=< 41661, -51639> velocity=<-4,  5>
position=< 31259, -20575> velocity=<-3,  2>
position=< 31276,  51906> velocity=<-3, -5>
position=< 41615,  10494> velocity=<-4, -1>
position=<-41226, -51644> velocity=< 4,  5>
position=< 20932,  31204> velocity=<-2, -3>
position=< 31271,  51910> velocity=<-3, -5>
position=< 10561,  51906> velocity=<-1, -5>
position=< 41650, -10216> velocity=<-4,  1>
position=<-51582, -20574> velocity=< 5,  2>
position=< 41647,  20843> velocity=<-4, -2>
position=<-41227,  31200> velocity=< 4, -3>
position=< 20926, -41289> velocity=<-2,  4>
position=< 51986,  51915> velocity=<-5, -5>
position=< 41642, -51644> velocity=<-4,  5>
position=< 20908, -30933> velocity=<-2,  3>
position=< 10569,  51906> velocity=<-1, -5>
position=<-10162,  41555> velocity=< 1, -4>
position=< 31276,  51912> velocity=<-3, -5>
position=< 41636,  20850> velocity=<-4, -2>
position=< 51973, -51641> velocity=<-5,  5>
position=<-51588, -10219> velocity=< 5,  1>
position=< 20948, -41283> velocity=<-2,  4>
position=<-30822, -30933> velocity=< 3,  3>
position=<-30823,  51911> velocity=< 3, -5>
position=<-30875,  51908> velocity=< 3, -5>
position=< 31287, -10222> velocity=<-3,  1>
position=< 31247, -10222> velocity=<-3,  1>
position=<-10116, -41289> velocity=< 1,  4>
position=< 41643,  10490> velocity=<-4, -1>
position=< 20925,  31200> velocity=<-2, -3>
position=<-30854,  20848> velocity=< 3, -2>
position=<-30872,  20841> velocity=< 3, -2>
position=< 41610,  41552> velocity=<-4, -4>
position=<-41228,  20845> velocity=< 4, -2>
position=< 31283,  10486> velocity=<-3, -1>
position=<-20512,  41558> velocity=< 2, -4>
position=< 41623,  41554> velocity=<-4, -4>
position=< 51973, -30933> velocity=<-5,  3>
position=< 31296,  20845> velocity=<-3, -2>
position=< 41642, -10215> velocity=<-4,  1>
position=<-20517, -41285> velocity=< 2,  4>
position=<-51548,  31197> velocity=< 5, -3>
position=< 20897, -51642> velocity=<-2,  5>
position=< 41642,  51915> velocity=<-4, -5>
position=< 20935,  41555> velocity=<-2, -4>
position=< 20932,  41555> velocity=<-2, -4>
position=< 41634, -20577> velocity=<-4,  2>
position=<-51574,  10486> velocity=< 5, -1>
position=< 52002,  10494> velocity=<-5, -1>
position=< 51997,  41559> velocity=<-5, -4>
position=< 41610, -10221> velocity=<-4,  1>
position=< 20909, -41289> velocity=<-2,  4>
position=< 31268,  10489> velocity=<-3, -1>
position=<-51577,  10492> velocity=< 5, -1>
position=<-51574, -20579> velocity=< 5,  2>
position=< 31263, -51640> velocity=<-3,  5>
position=<-51564,  10494> velocity=< 5, -1>
position=< 41607,  20846> velocity=<-4, -2>
position=<-10114, -51644> velocity=< 1,  5>
position=< 31295,  20842> velocity=<-3, -2>
position=< 31247,  51907> velocity=<-3, -5>
position=<-10164,  41555> velocity=< 1, -4>
position=<-30827, -51639> velocity=< 3,  5>
position=<-41233,  10488> velocity=< 4, -1>
position=< 52001,  41558> velocity=<-5, -4>
position=< 51981,  41551> velocity=<-5, -4>
position=<-20469,  10486> velocity=< 2, -1>
position=< 41607,  10490> velocity=<-4, -1>
position=< 20953, -20574> velocity=<-2,  2>
position=< 41631,  41556> velocity=<-4, -4>
position=< 20929, -41286> velocity=<-2,  4>
position=<-51574, -41289> velocity=< 5,  4>
position=<-30851,  20848> velocity=< 3, -2>
position=< 51997, -41287> velocity=<-5,  4>
position=<-41206,  31200> velocity=< 4, -3>
position=<-10141, -51641> velocity=< 1,  5>
position=< 31271,  10488> velocity=<-3, -1>
position=<-20491, -30929> velocity=< 2,  3>
position=< 10574,  51912> velocity=<-1, -5>
position=< 51986, -30930> velocity=<-5,  3>
position=<-51585,  20847> velocity=< 5, -2>
position=< 31303, -10216> velocity=<-3,  1>
position=< 10574, -20572> velocity=<-1,  2>
position=< 41622,  31196> velocity=<-4, -3>
position=< 52005, -51635> velocity=<-5,  5>
position=< 52007,  20847> velocity=<-5, -2>
position=< 31300,  10486> velocity=<-3, -1>
position=<-41203,  51906> velocity=< 4, -5>
position=< 31303, -30929> velocity=<-3,  3>
position=<-30882,  10490> velocity=< 3, -1>
position=<-10128,  41559> velocity=< 1, -4>
position=< 31271, -10215> velocity=<-3,  1>
position=<-20468,  20846> velocity=< 2, -2>
position=<-20500, -20575> velocity=< 2,  2>
position=< 31308, -20578> velocity=<-3,  2>
position=<-10125, -20575> velocity=< 1,  2>
position=< 10594,  31196> velocity=<-1, -3>
position=< 41629,  31200> velocity=<-4, -3>
position=<-10161, -10217> velocity=< 1,  1>
position=< 31295,  41559> velocity=<-3, -4>
position=<-51537, -30931> velocity=< 5,  3>
position=<-51585,  10491> velocity=< 5, -1>
position=<-41182,  31198> velocity=< 4, -3>
position=<-41238,  20850> velocity=< 4, -2>
position=<-51556,  10491> velocity=< 5, -1>
position=< 51970,  31199> velocity=<-5, -3>
position=< 41607,  31204> velocity=<-4, -3>
position=< 10569,  31198> velocity=<-1, -3>
position=<-20491,  31197> velocity=< 2, -3>
position=<-30848,  31200> velocity=< 3, -3>
position=<-20494, -10224> velocity=< 2,  1>
position=<-10131,  20841> velocity=< 1, -2>
position=<-10130, -51640> velocity=< 1,  5>
position=<-10141,  31202> velocity=< 1, -3>
position=<-41188, -51638> velocity=< 4,  5>
position=< 20904, -30928> velocity=<-2,  3>
position=< 20929, -51641> velocity=<-2,  5>
position=<-10114, -20570> velocity=< 1,  2>
position=<-30883, -10219> velocity=< 3,  1>
position=< 41607,  10495> velocity=<-4, -1>
position=<-30851,  51912> velocity=< 3, -5>
position=< 20937,  20850> velocity=<-2, -2>
position=<-51575, -30930> velocity=< 5,  3>
position=< 10566, -30925> velocity=<-1,  3>
position=< 51986, -20572> velocity=<-5,  2>
position=<-51569, -51637> velocity=< 5,  5>
position=<-10130, -20574> velocity=< 1,  2>
position=<-10157, -41281> velocity=< 1,  4>
position=<-41217,  31197> velocity=< 4, -3>
position=<-30835,  41556> velocity=< 3, -4>
position=<-10128, -20576> velocity=< 1,  2>
position=< 20940,  10487> velocity=<-2, -1>
position=< 52013,  41556> velocity=<-5, -4>
position=<-51588, -20570> velocity=< 5,  2>
position=<-10157, -51641> velocity=< 1,  5>
position=<-10112,  20847> velocity=< 1, -2>
position=<-51573, -10220> velocity=< 5,  1>
position=< 20917,  51910> velocity=<-2, -5>
position=< 51997, -30933> velocity=<-5,  3>
position=< 41618,  20845> velocity=<-4, -2>
position=<-41205, -10224> velocity=< 4,  1>
position=<-30875,  41555> velocity=< 3, -4>
position=<-51559,  20845> velocity=< 5, -2>
position=< 31256,  51906> velocity=<-3, -5>
position=< 10590,  41560> velocity=<-1, -4>
position=< 41610, -41284> velocity=<-4,  4>
position=<-30842,  31196> velocity=< 3, -3>
position=< 10542, -10218> velocity=<-1,  1>
position=<-41227, -41289> velocity=< 4,  4>
position=< 10566,  10492> velocity=<-1, -1>
position=< 10546, -20579> velocity=<-1,  2>
position=< 10537, -30927> velocity=<-1,  3>
position=<-51564,  51910> velocity=< 5, -5>
position=<-41222, -41289> velocity=< 4,  4>
position=< 10593, -10217> velocity=<-1,  1>
position=<-20528, -20575> velocity=< 2,  2>`
