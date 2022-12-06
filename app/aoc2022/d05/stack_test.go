package d05

import (
	"testing"
)

func Test_Stack1(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Fatalf("Stack size should be 0, was %v\n", s.Size())
	}
	if s.Peek() != "" {
		t.Fatalf("Stack peek on size 0 should be empty strign, was %v\n", s.Peek())
	}
	s.Push("F")
	if s.Size() != 1 {
		t.Fatalf("Stack size should be 1, was %v\n", s.Size())
	}
	if s.Peek() != "F" {
		t.Fatalf("Stack peek on size 0 should be F, was %v\n", s.Peek())
	}

	s.Push("AF")
	if s.Size() != 2 {
		t.Fatalf("Stack size should be 2, was %v\n", s.Size())
	}
	if s.Peek() != "AF" {
		t.Fatalf("Stack peek on size 0 should be AF, was %v\n", s.Peek())
	}

	v := s.Pop()
	if v != "AF" {
		t.Fatalf("Stack pop should be AF, was %v\n", v)
	}
	if s.Size() != 1 {
		t.Fatalf("Popped Stack size should be 1, was %v\n", s.Size())
	}

}
