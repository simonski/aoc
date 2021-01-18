package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/simonski/aoc/utils"
	goutils "github.com/simonski/goutils"
)

const DAY_1_DATA = `1749
1897
881
1736
1161
1720
1676
305
264
1904
1880
1173
483
1978
1428
1635
1386
1858
1602
1916
1906
1212
1730
1777
1698
1845
1812
1922
1729
1803
1761
1901
1748
1188
1964
1935
1919
1810
1567
1849
1417
1452
54
1722
1784
1261
1744
1594
1526
1771
1762
1894
1717
1716
51
1955
1143
1741
1999
1775
1944
1983
1962
1198
1553
1835
1867
1662
1461
1811
1764
1726
1927
1179
1468
1948
1813
1213
1905
1371
1751
1215
1392
1798
1823
1815
1923
1942
1987
1887
1838
1395
2007
1479
1752
1945
1621
1538
1937
565
1969
1493
1291
1438
1578
1770
2005
1703
1712
1943
2003
1499
1903
1760
1950
1990
1185
1809
1337
1358
1743
1707
1671
1788
1785
1972
1863
1690
1512
1963
1825
1460
1828
1902
1874
1755
1951
1830
1767
1787
1373
1709
1514
1807
1791
1724
1859
1590
1976
1572
1947
1913
1995
1728
1624
1731
1706
1782
1994
1851
1843
1773
1982
1685
2001
1346
1200
1746
1520
972
1834
1909
2008
1733
1960
1280
1879
1203
1979
1133
1647
1282
1684
860
1444
1780
1989
1795
1819
1797
1842
1796
1457
1839
1853
1711
1883
1146
1734
1389`

func (app *Application) Y2020D01P1() {
	AOC_2020_01(&app.CLI)
}

func (app *Application) Y2020D01P2() {
	AOC_2020_02(&app.CLI)
}

// AOC_2020_01 is the entrypoint to the various attempts for day one
func AOC_2020_01(cli *goutils.CLI) {

	logger := utils.NewLogger("Day 01-1")
	logger.ShowTime = false
	logger.ShowLevel = false
	AOC_2020_01_part1_attempt1(cli, logger)
	AOC_2020_01_part1_attempt2(cli, logger)
	logger = utils.NewLogger("Day 01-2")
	logger.ShowTime = false
	logger.ShowLevel = false
	AOC_2020_01_part2_attempt1(cli, logger)

}

func day_1_load_data(cli *goutils.CLI) []int {
	data := make([]int, 0)
	if cli.IndexOf("-input") == -1 {
		data = make([]int, 0)
		for _, value := range strings.Split(DAY_1_DATA, "\n") {
			ival, _ := strconv.Atoi(value)
			data = append(data, ival)
		}

	} else {
		filename := cli.GetStringOrDie("-input")
		data = utils.Load_file_to_ints(filename)
	}
	return data
}

// AOC_2020_01_part2_attempt1 the second part of day 1, attempt1
// this is a brute-force attempt that gets over the line
// in the spirit of make it work, make it fast, this is make it work
// so this is 3 inner loops, giving o(n^3) performance I believe - it works, it is not fast.
func AOC_2020_01_part2_attempt1(cli *goutils.CLI, logger *utils.Logger) {

	// now we need to find 3 numbers that meet our total

	// so that is a case of

	// 1... 2020
	// a * b * c ! > maximum

	// make it work
	// make it right
	// make it fast

	// find two entries in input that sum to 2020
	// find combination that yields highest product

	// left to right
	// attempt 1: brute force attempt first entry and walk up to find entry that totals
	// retain maximum

	// then come back and do it properly
	logger.Debug(fmt.Sprintf("Part2:"))
	data := day_1_load_data(cli)
	totalRequired := 2020
	maxSoFar := 0
	maxValue1 := 0
	maxValue2 := 0
	maxValue3 := 0
	oCount := 0
	for index1 := 0; index1 < len(data); index1++ {
		value1 := data[index1]
		for index2 := 0; index2 < len(data); index2++ {
			if index2 == index1 {
				continue
			}
			value2 := data[index2]
			if value1+value2 >= totalRequired {
				continue
			}
			for index3 := 0; index3 < len(data); index3++ {
				oCount++
				value3 := data[index3]
				if value1+value2+value3 == totalRequired {
					product := value1 * value2 * value3
					if product > maxSoFar {
						maxSoFar = product
						maxValue1 = value1
						maxValue2 = value2
						maxValue3 = value3
						logger.Debug(fmt.Sprintf("Part2: New maximum: %v+%v+%v=%v, %v*%v*%v=%v", value1, value2, value3, value1+value2+value3, value1, value2, value3, value1*value2*value3))
					}
				}
			}

		}
	}

	logger.Debug(fmt.Sprintf("Part2: Maximum: %v, (%v * %v * %v )", maxSoFar, maxValue1, maxValue2, maxValue3))
	logger.Debug(fmt.Sprintf("Part2: o(n) is o(%v)=%v", len(data), oCount))

}

// AOC_2020_01_part1_attempt1 this is part 1 of day 1, attempt 1
// a brute-force attempt which as the volume is small works fine
// we have an inner loop giving o(n^2) which again works but is not fast
func AOC_2020_01_part1_attempt1(cli *goutils.CLI, logger *utils.Logger) {

	// make it work
	// make it right
	// make it fast

	// find two entries in input that sum to 2020
	// find combination that yields highest product

	// left to right
	// attempt 1: brute force attempt first entry and walk up to find entry that totals
	// retain maximum

	// then come back and do it properly

	data := day_1_load_data(cli)

	totalRequired := 2020
	maxSoFar := 0
	maxValue1 := 0
	maxValue2 := 0
	oCount := 0
	for index1 := 0; index1 < len(data); index1++ {
		value1 := data[index1]
		for index2 := 0; index2 < len(data); index2++ {
			oCount++
			if index2 == index1 {
				continue
			}

			value2 := data[index2]

			if value1+value2 == totalRequired {
				if value1*value2 > maxSoFar {
					maxSoFar = value1 * value2
					maxValue1 = value1
					maxValue2 = value2
					logger.Debug(fmt.Sprintf("Part1: New maximum: %v+%v=%v, %v * %v=%v", value1, value2, value1+value2, value1, value2, value1*value2))
				}
			}

		}
	}

	logger.Debug(fmt.Sprintf("Part1: Maximum: %v, (%v x %v)", maxSoFar, maxValue1, maxValue2))
	logger.Debug(fmt.Sprintf("Part1: o(n^2)=%v", oCount))

}

// AOC_2020_01_part1_attempt2
// in this atempt I preload an "inty" map affording go's own probably binsearch by keying on the int value
// itself
// this uses more memory (the inty map in addition to the list) but avoids an initial sort and binsearch
// I intend to do my own sort and binsearch as an attempt3
func AOC_2020_01_part1_attempt2(cli *goutils.CLI, logger *utils.Logger) {

	// make it fast
	// so now I think sorting the numbers and doing a binary chop will give me o(log n) performance
	// as the first impl gave me my inner loop with is o(n^2) as I have to search everything; this way
	// I'll reduce my search space down somewhat

	// I think I'll start again in a loop but for each I'll workout my maximum value I can multiply with by
	// list = sorted(list)
	// for index in list:
	// 	entry = list[index]
	//  2020 / entry = ?   if value is integer, binsearch, else discard
	//	if found, retain if > max

	logger.Debug(fmt.Sprintf("\nPart1:\n"))
	data := day_1_load_data(cli)

	mapints := utils.Make_map_of_inty_list(data)

	// don't need to binsearch if use an inbuild map

	totalRequired := 2020
	maxSoFar := 0
	maxValue1 := 0
	maxValue2 := 0
	oCount := 0
	for index := 0; index < len(data); index++ {
		// check - will there be an int availble?
		oCount++
		value := data[index]

		// we want searchFor exactly
		searchFor := totalRequired - value

		// otherwise it's an int. Do we have it?
		_, exists := mapints[searchFor]
		if exists {
			// yes, it exists.  These sum to our max
			product := value * searchFor
			if product > maxSoFar {
				maxSoFar = product
				maxValue1 = value
				maxValue2 = searchFor
				logger.Debug(fmt.Sprintf("Part1: New maximum: %v+%v=%v, %v * %v=%v", maxValue1, maxValue2, maxValue1+maxValue2, maxValue1, maxValue2, maxValue1*maxValue2))
			}
		}
	}

	logger.Debug(fmt.Sprintf("Part1: Maximum: %v, (%v x %v)", maxSoFar, maxValue1, maxValue2))
	logger.Debug(fmt.Sprintf("Part1: o(n) is o(n log n)=%v", oCount))

}
