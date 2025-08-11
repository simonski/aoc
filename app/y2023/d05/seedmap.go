package d5

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SeedMap struct {
	Seeds    []int
	Mappings map[string]*Mapping
}

func NewSeedMap(input string) *SeedMap {
	lines := strings.Split(input, "\n")
	seeds := readLineToint(strings.Split(lines[0], ":")[1])
	mappings := make(map[string]*Mapping)
	var mapping *Mapping
	var mapName string
	for _, line := range lines[1:] {
		if line == "" {
			continue
		} else if strings.Contains(line, ":") {
			// it's a new map
			mapName = strings.Split(line, " ")[0]
			mapping = NewMapping(mapName)
			mappings[mapName] = mapping
		} else {
			mapping.Add(NewEntry(line))
		}
	}

	sm := SeedMap{}
	sm.Seeds = seeds
	sm.Mappings = mappings
	return &sm
}

type Mapping struct {
	Name    string
	Entries []*Entry
}

func NewMapping(name string) *Mapping {
	return &Mapping{Name: name, Entries: make([]*Entry, 0)}
}

func (m *Mapping) Size() int {
	return len(m.Entries)
}

func (m *Mapping) SourceToDestination(in int, verbose bool) int {
	entry := m.FindEntryForSeed(in)
	if entry == nil {
		if verbose {
			fmt.Printf("SourceToDestination(in=%v), entry=nil\n", in)
		}
		return in
	}
	if verbose {
		fmt.Printf("SourceToDestination(in=%v), entry=%v\n", in, entry.Debug())
	}
	candidate := entry.SourceToDestination(in, verbose)
	if candidate != in {
		return candidate
	}
	return in
}

func (m *Mapping) FindEntryForSeed(in int) *Entry {
	for _, entry := range m.Entries {
		if in >= entry.SourceMin && in <= entry.SourceMax {
			return entry
		}
	}
	return nil
}

func (m *Mapping) Add(e *Entry) {
	m.Entries = append(m.Entries, e)
}

type Entry struct {
	Line                  string
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
	SourceMin             int
	SourceMax             int
	DestinationMin        int
	DestinationMax        int
}

func NewEntry(line string) *Entry {
	e := Entry{}
	e.Line = line
	ints := readLineToint(line)
	e.DestinationRangeStart = ints[0]
	e.SourceRangeStart = ints[1]
	e.RangeLength = ints[2]

	e.SourceMin = e.SourceRangeStart
	e.SourceMax = e.SourceRangeStart + e.RangeLength - 1

	e.DestinationMin = e.DestinationRangeStart
	e.DestinationMax = e.DestinationRangeStart + e.RangeLength - 1

	return &e
}

func (e *Entry) Debug() string {
	return fmt.Sprintf("Entry: (%v) - dest=%v source=%v size=%v (%v<->%v becomes %v <-> %v)", e.Line, e.DestinationRangeStart, e.SourceRangeStart, e.RangeLength, e.SourceMin, e.SourceMax, e.DestinationMin, e.DestinationMax)
}

func (e *Entry) SourceToDestination(source int, verbose bool) int {

	smin := e.SourceMin
	smax := e.SourceMax
	smin_dest := e.DestinationMin
	if source < smin {
		if verbose {
			fmt.Printf("SourceToDestination(%v) (DestStart: %v, SourceStart: %v, RangeLength: %v) min-max (%v/%v/range), result=%v\n", source, e.DestinationRangeStart, e.SourceRangeStart, e.RangeLength, smin, smax, source)
		}
		return source
	} else if source > smax {
		if verbose {
			fmt.Printf("SourceToDestination(%v) (DestStart: %v, SourceStart: %v, RangeLength: %v) min-max (%v/%v/range), result=%v\n", source, e.DestinationRangeStart, e.SourceRangeStart, e.RangeLength, smin, smax, source)
		}
		return source
	} else {
		index := source - smin
		result := smin_dest + index
		if verbose {
			fmt.Printf("SourceToDestination(%v) (DestStart: %v, SourceStart: %v, RangeLength: %v) min-max (%v/%v/range), result=%v\n", source, e.DestinationRangeStart, e.SourceRangeStart, e.RangeLength, smin, smax, result)
		}
		return result
	}
}

func (sm *SeedMap) GetLocationFromSeed(seed int) int {
	v := sm.Mappings["seed-to-soil"].SourceToDestination(seed, false)
	v = sm.Mappings["soil-to-fertilizer"].SourceToDestination(v, false)
	v = sm.Mappings["fertilizer-to-water"].SourceToDestination(v, false)
	v = sm.Mappings["water-to-light"].SourceToDestination(v, false)
	v = sm.Mappings["light-to-temperature"].SourceToDestination(v, false)
	v = sm.Mappings["temperature-to-humidity"].SourceToDestination(v, false)
	v = sm.Mappings["humidity-to-location"].SourceToDestination(v, false)
	return v
}

func (sm *SeedMap) GetSeedRanges() [][]int {
	seedranges := make([][]int, 0)
	for index := 0; index < len(sm.Seeds); index++ {
		index_start := sm.Seeds[index]
		length := sm.Seeds[index+1]
		index_end := index_start + length - 1
		seedrange := []int{index_start, index_end, length}
		seedranges = append(seedranges, seedrange)
		index += 1
	}
	return seedranges
}

func (sm *SeedMap) GetMinSeedAndLocation() (int, int) {
	return sm.GetMinSeedAndLocationA()
}

func (sm *SeedMap) GetMinSeedAndLocationA() (int, int) {

	seedranges1 := sm.GetSeedRanges()
	fmt.Printf("There are %v ranges before collapsing.\n", len(seedranges1))
	for index, srange := range seedranges1 {
		fmt.Printf("[%v] %v <-> %v\n", index, srange[0], srange[1])
	}

	// seedranges2 := sm.GetSeedRangesCollapsed()
	// fmt.Printf("\nThere are %v ranges after collapsing.\n", len(seedranges2))
	// for index, srange := range seedranges2 {
	// 	fmt.Printf("[%v] %v <-> %v\n", index, srange[0], srange[1])
	// }

	seedranges := seedranges1
	min_location := math.MaxInt
	min_seed := math.MaxInt

	total_seeds := 0
	for _, sr := range seedranges {
		total_seeds += sr[2]
	}
	fmt.Printf("There are %v seed ranges, totalling %v seeds\n", len(seedranges), total_seeds)

	// collapse the seed ranges down to contiguous blocks

	counter := 0
	for seedrangeIndex, sr := range seedranges {
		lowseed := sr[0]
		highseed := sr[1]
		// cache := make(map[int]int)
		fmt.Printf("SeedRange[%v] (%v-%v), %v seeds:\n", seedrangeIndex, lowseed, highseed, highseed-lowseed+1)
		for seed := lowseed; seed <= highseed; seed++ {
			counter++
			location := sm.GetLocationFromSeed(int(seed))
			if location < min_location {
				min_location = location
				min_seed = int(seed)
				fmt.Printf("Seed: %v, Location=%v, CacheSize=%v\n", min_seed, min_location, 0)
			}
			if counter%10000000 == 0 {
				fmt.Printf("(%v/%v)\n", counter, total_seeds)
			}
		}
	}
	return min_seed, min_location
}

func readLineToint(line string) []int {
	results := make([]int, 0)
	splits := strings.Split(line, " ")
	for _, split := range splits {
		v, e := strconv.Atoi(split)
		if e == nil {
			results = append(results, v)
		}
	}
	return results
}
