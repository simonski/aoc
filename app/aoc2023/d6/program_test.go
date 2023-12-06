package d6

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	p := NewPuzzleWithData(TEST_DATA)
	fmt.Printf("There are %v lines.\n", len(p.lines))
}

func Test_2(t *testing.T) {

	expect := func(racetime int, downtime int, expected int) {
		actual := distance(downtime, racetime)
		if actual != expected {
			t.Fatalf("For race time '%v',  downtime '%v' expects '%v' but was '%v'\n", racetime, downtime, expected, actual)
		}

	}
	expect(7, 0, 0)
	expect(7, 1, 6)
	expect(7, 2, 10)
	expect(7, 3, 12)
	expect(7, 4, 12)
	expect(7, 5, 10)
	expect(7, 6, 6)
	expect(7, 7, 0)

}

func Test_3(t *testing.T) {
	times1 := get_best_times(7, 9)
	times2 := get_best_times(15, 40)
	times3 := get_best_times(30, 200)
	result := len(times1) * len(times2) * len(times3)
	if result != 288 {
		t.Fatalf("test3: %v\n", result)
	}
}

func Test_4(t *testing.T) {
	times1 := get_best_times(71530, 940200)
	t.Fatalf("test4:\n%v\nways = %v\n", times1, len(times1))
}

func Test_5(t *testing.T) {
	times1 := get_best_times(35696887, 213116810861248)
	t.Fatalf("test5:\n%v\nways = %v\n", times1, len(times1))
}
