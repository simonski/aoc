package d5

import (
	"fmt"
	"log"
	"testing"
)

func Test_1(t *testing.T) {

	sm := NewSeedMap(TEST_DATA)
	if len(sm.Seeds) != 4 {
		t.Fatalf("Seeds should be 5., was %v\n", len(sm.Seeds))
	}

	if len(sm.Mappings) != 7 {
		t.Fatalf("Seeds should be 7., was %v\n", len(sm.Mappings))
	}

	checkSize := func(name string, sm *SeedMap, expected int) {
		mapping := sm.Mappings[name]
		if mapping.Size() != expected {
			actual := mapping.Size()
			t.Fatalf("%v should be %v, was %v\n", name, expected, actual)
		}
	}

	checkSize("soil-to-fertilizer", sm, 3)
	checkSize("fertilizer-to-water", sm, 4)
	checkSize("water-to-light", sm, 2)
	checkSize("light-to-temperature", sm, 3)
	checkSize("temperature-to-humidity", sm, 2)
	checkSize("humidity-to-location", sm, 2)

}

func Test_2(t *testing.T) {

	sm := NewSeedMap(TEST_DATA)
	expect := func(name string, sm *SeedMap, source int, expected int) {
		mapping := sm.Mappings[name]
		actual := mapping.SourceToDestination(source, false)
		if actual != expected {
			log.Fatalf("%v %v -> %v (!= expected %v)\n", name, source, actual, expected)
		}
	}

	expect("seed-to-soil", sm, 0, 0)
	expect("seed-to-soil", sm, 1, 1)
	expect("seed-to-soil", sm, 48, 48)
	expect("seed-to-soil", sm, 49, 49)
	expect("seed-to-soil", sm, 50, 52)
	expect("seed-to-soil", sm, 51, 53)
	expect("seed-to-soil", sm, 96, 98)
	expect("seed-to-soil", sm, 97, 99)
	expect("seed-to-soil", sm, 98, 50)
	expect("seed-to-soil", sm, 99, 51)

	expect("seed-to-soil", sm, 79, 81)
	expect("seed-to-soil", sm, 14, 14)
	expect("seed-to-soil", sm, 55, 57)
	expect("seed-to-soil", sm, 13, 13)

}

func Test_3(t *testing.T) {

	sm := NewSeedMap(TEST_DATA)
	expect := func(name string, sm *SeedMap, source int, expected int) {
		mapping := sm.Mappings[name]
		actual := mapping.SourceToDestination(source, false)
		if actual != expected {
			log.Fatalf("%v %v -> %v (!= expected %v)\n", name, source, actual, expected)
		}
	}

	expect("seed-to-soil", sm, 79, 81)
	expect("soil-to-fertilizer", sm, 81, 81)
	expect("fertilizer-to-water", sm, 81, 81)
	expect("water-to-light", sm, 81, 74)
	expect("light-to-temperature", sm, 74, 78)
	expect("temperature-to-humidity", sm, 78, 78)
	expect("humidity-to-location", sm, 78, 82)

	expect("seed-to-soil", sm, 14, 14)
	expect("soil-to-fertilizer", sm, 14, 53)
	expect("fertilizer-to-water", sm, 53, 49)
	expect("water-to-light", sm, 49, 42)
	expect("light-to-temperature", sm, 42, 42)
	expect("temperature-to-humidity", sm, 42, 43)
	expect("humidity-to-location", sm, 43, 43)

	expect("seed-to-soil", sm, 55, 57)
	expect("soil-to-fertilizer", sm, 57, 57)
	expect("fertilizer-to-water", sm, 57, 53)
	expect("water-to-light", sm, 53, 46)
	expect("light-to-temperature", sm, 46, 82)
	expect("temperature-to-humidity", sm, 82, 82)
	expect("humidity-to-location", sm, 82, 86)

	expect("seed-to-soil", sm, 13, 13)
	expect("soil-to-fertilizer", sm, 13, 52)
	expect("fertilizer-to-water", sm, 52, 41)
	expect("water-to-light", sm, 41, 34)
	expect("light-to-temperature", sm, 34, 34)
	expect("temperature-to-humidity", sm, 34, 35)
	expect("humidity-to-location", sm, 35, 35)

}

func Test_4(t *testing.T) {
	expect2 := func(sm *SeedMap, source int, expected int) {

		actual := sm.GetLocationFromSeed(source)
		if actual != expected {
			log.Fatalf("%v -> %v (!= expected %v)\n", source, actual, expected)
		}

	}

	sm := NewSeedMap(TEST_DATA)

	expect2(sm, 79, 82)
	expect2(sm, 14, 43)
	expect2(sm, 55, 86)
	expect2(sm, 13, 35)

}

func Test_5(t *testing.T) {

	sm := NewSeedMap(TEST_DATA)
	seedranges := sm.GetSeedRanges()
	var total_seeds int
	total_seeds = 0
	for _, sr := range seedranges {
		total_seeds += sr[2]
	}
	if total_seeds != 27 {
		log.Fatalf("Expected 27 seeds, got %v\n", total_seeds)
	}

	min_seed, min_location := sm.GetMinSeedAndLocation()
	fmt.Printf("MinSeed: %v, Minlocation: %v\n", min_seed, min_location)
}
