package main

import (
	"fmt"
	"time"
)

type StopwatchEntry struct {
	Moment  time.Time
	Message string
}

type Stopwatch struct {
	Prefix string
	Splits []*StopwatchEntry
}

func (s *Stopwatch) Start() {
	s.Clear()
	s.Split("Start")
}

func (s *Stopwatch) End() {
	s.Clear()
	s.Split("End")
}

func (s *Stopwatch) Clear() {
	s.Splits = make([]*StopwatchEntry, 0)
}

func (s *Stopwatch) First() *StopwatchEntry {
	if s.Size() > 0 {
		return s.Splits[0]
	} else {
		return nil
	}
}

func (s *Stopwatch) Last() *StopwatchEntry {
	index := s.Size() - 1
	if index > -1 {
		return s.Splits[index]
	} else {
		return nil
	}
}

func (s *Stopwatch) Split(msg string) *StopwatchEntry {
	entry := &StopwatchEntry{Moment: time.Now(), Message: msg}
	s.Splits = append(s.Splits, entry)
	return entry
}

func (s *Stopwatch) Size() int {
	return len(s.Splits)
}

func NewStopwatch(prefix string) *Stopwatch {
	sw := &Stopwatch{Prefix: prefix}
	sw.Splits = make([]*StopwatchEntry, 0)
	return sw
}

func (s *Stopwatch) DebugMilliseconds() string {
	line := ""
	startTime := s.Splits[0].Moment
	lastTime := startTime
	// endTime := s.Splits[len(s.Splits)-1].Moment
	for _, entry := range s.Splits {
		now := entry.Moment
		parsedTime := now.Format(time.RFC3339)
		msSinceStart := now.Sub(startTime).Milliseconds()
		msSinceLast := now.Sub(lastTime).Milliseconds()
		lastTime = now
		line += fmt.Sprintf("%v %v %v %v\n", parsedTime, msSinceStart, msSinceLast, entry.Message)
	}
	return line
}

func (s *Stopwatch) DebugNanoseconds() string {
	line := ""
	startTime := s.Splits[0].Moment
	lastTime := startTime
	// endTime := s.Splits[len(s.Splits)-1].Moment
	for _, entry := range s.Splits {
		now := entry.Moment
		parsedTime := now.Format(time.RFC3339)
		msSinceStart := now.Sub(startTime).Nanoseconds()
		msSinceLast := now.Sub(lastTime).Nanoseconds()
		lastTime = now
		line += fmt.Sprintf("%v %v %v %v\n", parsedTime, msSinceStart, msSinceLast, entry.Message)
	}
	return line
}
