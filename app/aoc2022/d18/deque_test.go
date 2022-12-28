package d18

import "testing"

func Test_Deque(t *testing.T) {
	q := NewQ()
	if q.Size() != 0 {
		t.Fatalf("Q shoudl be 0")
	}
	q.Pushleft(NewCube("1,1,1"))
	if q.Size() != 1 {
		t.Fatalf("Q shoudl be 1")
	}
	q.Pushleft(NewCube("1,1,1"))
	if q.Size() != 2 {
		t.Fatalf("Q shoudl be 2")
	}
	c := q.Popleft()
	if q.Size() != 1 {
		t.Fatalf("Q shoudl be 1")
	}
	if c == nil {
		t.Fatalf("C shoul dnot be nil")
	}

	c = q.Popleft()
	if q.Size() != 0 {
		t.Fatalf("Q shoudl be 0")
	}
	if c == nil {
		t.Fatalf("C shoul dnot be nil")
	}

	c = q.Popleft()
	if q.Size() != 0 {
		t.Fatalf("Q shoudl be 0")
	}
	if c != nil {
		t.Fatalf("C should be nil")
	}

}
