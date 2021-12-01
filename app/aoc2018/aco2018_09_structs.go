package aoc2018

import (
	"fmt"
	"strconv"
	"strings"
)

type Route struct {
	OriginalLine string
	Source       string
	Target       string
	Distance     int
	Key          string
}

func (route *Route) Debug() string {
	return route.OriginalLine
}

func NewRoute(line string) *Route {
	//Tristram to Tambi = 63
	splits := strings.Split(line, " ")
	source := splits[0]
	target := splits[2]
	distance, _ := strconv.Atoi(splits[4])
	key := fmt.Sprintf("%v->%v", source, target)
	return &Route{OriginalLine: line, Source: source, Target: target, Distance: distance, Key: key}
}

// the path a->b
type Path struct {
	Routes   []*Route
	Distance int
}

type Location struct {
	Name             string
	SourceCount      int // number of ties this location is a source
	DestinationCount int // number of times it is a destination
}

func (path *Path) AddRoute(route *Route) {
	path.Routes = append(path.Routes, route)
	path.Distance += route.Distance
}

func (path *Path) GetStart() *Route {
	return path.Routes[0]
}

func (path *Path) GetEnd() *Route {
	index := len(path.Routes) - 1
	return path.Routes[index]
}

type RouteLogic struct {
	Routes    []*Route
	Locations map[string]*Location
}

func NewRouteLogic(data string) *RouteLogic {
	var locations = make(map[string]*Location)

	routes := make([]*Route, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		route := NewRoute(line)
		routes = append(routes, route)

		location := locations[route.Source]
		if location == nil {
			location = &Location{Name: route.Source}
			locations[route.Source] = location
		}
		location.SourceCount += 1
		location.DestinationCount += 1

		location = locations[route.Target]
		if location == nil {
			location = &Location{Name: route.Source}
			locations[route.Source] = location
		}
		location.DestinationCount += 1

	}
	return &RouteLogic{Routes: routes, Locations: locations}
}

func (logic *RouteLogic) GetLocation(name string) *Location {
	return logic.Locations[name]
}

func (logic *RouteLogic) Analsyse() {
	// Tristram to AlphaCentauri = 34
	// Tristram to Snowdin = 100
	// Tristram to Tambi = 63
	// Tristram to Faerun = 108
	// Tristram to Norrath = 111
	// Tristram to Straylight = 89
	// Tristram to Arbre = 132
	// AlphaCentauri to Snowdin = 4
	// AlphaCentauri to Tambi = 79
	// AlphaCentauri to Faerun = 44

}
func (logic *RouteLogic) FindFirstAndLastLocations() (*Location, *Location) {
	var first *Location
	var last *Location

	for _, route := range logic.Routes {
		location := logic.GetLocation(route.Source)
		if location.SourceCount == 0 {
			first = location
			if last != nil {
				break
			}
		}
		if location.DestinationCount == 0 {
			last = location
			if first != nil {
				break
			}
		}
	}
	return first, last
}
