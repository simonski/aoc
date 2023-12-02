package d7

import (
	"fmt"
	"testing"
)

func require(fs *FileSystem, path string, fsType FILETYPE, size int, t *testing.T) {
	entry := fs.Get(path)
	if entry == nil {
		t.Fatalf("%v must exist.\n", path)
	}

	fmt.Printf("Get('%v'), found absolute '%v'\n", path, entry.Path())

	if entry.Type != fsType {
		t.Fatalf("%v must be type %v, was %v\n", path, fsType, entry.Type)
	}
	if entry.Size != size {
		t.Fatalf("%v size must be %v, was %v\n", path, size, entry.Size)
	}
}

func Test_NewFileSystem(t *testing.T) {
	fs := NewFileSystem(TEST_DATA)
	require(fs, "/a", DIR, 0, t)
	require(fs, "/b.txt", FILE, 14848514, t)
	require(fs, "/c.dat", FILE, 8504156, t)
	require(fs, "/d", DIR, 0, t)
	require(fs, "/a/e/i", FILE, 584, t)
	require(fs, "/a/f", FILE, 29116, t)
	require(fs, "/a/g", FILE, 2557, t)
	require(fs, "/a/h.lst", FILE, 62596, t)
	require(fs, "/b.txt", FILE, 14848514, t)
	require(fs, "/c.dat", FILE, 8504156, t)
}

func Test_TotalSize(t *testing.T) {
	fs := NewFileSystem(TEST_DATA)
	entry := fs.Get("/a/e")
	total := entry.TotalSize()
	if total != 584 {
		t.Fatalf("/a/e should be total 584, was %v\n", total)
	}

	entry = fs.Get("/a")
	total = entry.TotalSize()
	if total != 94853 {
		t.Fatalf("/a shoudl be total 94853, was %v\n", total)
	}

	entry = fs.Get("/d")
	total = entry.TotalSize()
	if total != 24933642 {
		t.Fatalf("/d should be total 24933642, was %v\n", total)
	}

	entry = fs.Get("/")
	total = entry.TotalSize()
	if total != 48381165 {
		t.Fatalf("/ should be total 48381165, was %v\n", total)
	}
}

func Test_CD(t *testing.T) {
	fs := NewFileSystem(TEST_DATA)
	fs.CD("/")
	pwd := fs.PWD()
	if pwd != "/" {
		t.Fatalf("cd / should give pwd of /, gave %v\n", pwd)
	}

	fs.CD("/a/e")
	pwd = fs.PWD()
	if pwd != "/a/e" {
		t.Fatalf("cd /a/e should give pwd of /a/e , gave %v\n", pwd)
	}

}

func Test_NewEntry(t *testing.T) {
	e1 := NewEntry("dir fo")
	if e1.Type != DIR {
		t.Fatal("e1 should be DIR")
	}
	if e1.Size != 0 {
		t.Fatal("e1 size should be 0")
	}
	if e1.Name != "fo" {
		t.Fatal("e1 name should be fo")
	}

	e2 := NewEntry("34324 blah.h")
	if e2.Type != FILE {
		t.Fatal("e2 should be FILE")
	}
	if e2.Size != 34324 {
		t.Fatal("e2 should be 34324")
	}
	if e2.Name != "blah.h" {
		t.Fatal("e2 should be blah.h")
	}
}

func Test_TestData_DirsOver(t *testing.T) {
	fs := NewFileSystem(TEST_DATA)
	threshold := 100000
	current := fs.Root
	biggies := make([]*Entry, 0)
	biggies = walkLTE(current, biggies, threshold)

	total := 0
	for _, entry := range biggies {
		fmt.Printf("BigDir %v (%v) size %v\n", entry.Name, entry.Path(), entry.TotalSize())
		total += entry.TotalSize()
	}
	fmt.Printf("Total big dir size is %v\n", total)
	t.Fatalf("a")
}

func walkLTE(current *Entry, biggies []*Entry, threshold int) []*Entry {
	for _, child := range current.Children {
		if child.Type == DIR {
			totalSize := child.TotalSize()
			fmt.Printf("walk(%v), child=%v, size=%v\n", current.Name, child.Name, totalSize)
			if totalSize <= threshold {
				biggies = append(biggies, child)
			}
			biggies = walkLTE(child, biggies, threshold)
		}
	}
	return biggies
}

func walkGTE(current *Entry, biggies []*Entry, threshold int) []*Entry {
	for _, child := range current.Children {
		if child.Type == DIR {
			totalSize := child.TotalSize()
			fmt.Printf("walk(%v), child=%v, size=%v\n", current.Name, child.Name, totalSize)
			if totalSize >= threshold {
				biggies = append(biggies, child)
			}
			biggies = walkGTE(child, biggies, threshold)
		}
	}
	return biggies
}

func Test_Real(t *testing.T) {
	fs := NewFileSystem(REAL_DATA)
	fs.Get("/").TotalSize()
}

func Test_RealData_DirsOver(t *testing.T) {
	fs := NewFileSystem(REAL_DATA)
	threshold := 100000
	current := fs.Root
	biggies := make([]*Entry, 0)
	biggies = walkLTE(current, biggies, threshold)

	total := 0
	for _, entry := range biggies {
		fmt.Printf("BigDir %v (%v) size %v\n", entry.Name, entry.Path(), entry.TotalSize())
		total += entry.TotalSize()
	}
	fmt.Printf("Total big dir size is %v\n", total)
	t.Fatalf("a")
}

func Test_TestData_DirsOverX(t *testing.T) {
	fs := NewFileSystem(TEST_DATA)
	total_available := 70000000
	total_required := 30000000
	total_used := fs.Get("/").TotalSize()
	total_unused := total_available - total_used

	total_to_find := total_required - total_unused
	fmt.Printf("total_available: %v\n", total_available)
	fmt.Printf("total_required: %v\n", total_required)
	fmt.Printf("total_used: %v\n", total_used)
	fmt.Printf("total_unused: %v\n", total_unused)
	fmt.Printf("total_to_find: %v\n", total_to_find)

	threshold := total_to_find
	current := fs.Root
	biggies := make([]*Entry, 0)
	biggies = walkGTE(current, biggies, threshold)

	for _, entry := range biggies {
		fmt.Printf("Findit %v (%v) size %v\n", entry.Name, entry.Path(), entry.TotalSize())
	}
	t.Fatalf("a")
}

func Test_TestData_DirsOverX_REAL(t *testing.T) {
	fs := NewFileSystem(REAL_DATA)
	total_available := 70000000
	total_required := 30000000
	total_used := fs.Get("/").TotalSize()
	total_unused := total_available - total_used

	total_to_find := total_required - total_unused
	fmt.Printf("total_available: %v\n", total_available)
	fmt.Printf("total_required: %v\n", total_required)
	fmt.Printf("total_used: %v\n", total_used)
	fmt.Printf("total_unused: %v\n", total_unused)
	fmt.Printf("total_to_find: %v\n", total_to_find)

	threshold := total_to_find
	current := fs.Root
	biggies := make([]*Entry, 0)
	biggies = walkGTE(current, biggies, threshold)

	smallest_entry := biggies[0]
	for _, entry := range biggies {
		if entry.TotalSize() < smallest_entry.TotalSize() {
			smallest_entry = entry
			fmt.Printf("Smallest is now %v, size %v\n", entry.Name, entry.TotalSize())
		}
	}
	fmt.Printf("Finditreal %v (%v) size %v\n", smallest_entry.Name, smallest_entry.Path(), smallest_entry.TotalSize())
	t.Fatalf("a")
}
