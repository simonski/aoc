package d07

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 05:  ---

*/

type MODE string

const MODE_CREATE MODE = "Create"
const MODE_WALK MODE = "Walk"

type FileSystem struct {
	input   []string
	Root    *Entry
	Current *Entry
	Mode    MODE
}

func NewFileSystem(input string) *FileSystem {
	splits := strings.Split(input, "\n")
	fs := FileSystem{input: splits}
	fs.Mode = MODE_CREATE
	fs.Root = NewEntry("dir /")
	fs.Current = fs.Root

	for _, line := range splits {
		if line[0:1] == "$" {
			// instruction
			if line[0:4] == "$ cd" {
				// change dir and add (or view things)
				fs.CD(line[5:])
			} else if line[0:4] == "$ ls" {
				// learn what exists here (create some entries)
				// nothing to do - as the add will occur in the other part
			}
		} else {
			// entry
			e := NewEntry(line)
			fs.Current.Add(e)
		}

	}
	fs.Mode = MODE_WALK

	return &fs
}

func (fs *FileSystem) CD(path string) *Entry {
	if path == ".." {
		fs.Current = fs.Current.Parent
		return fs.Current
	} else if path == "/" {
		fs.Current = fs.Root
		return fs.Current
	}

	var current *Entry
	if path[0:1] == "/" {
		current = fs.Root
		fmt.Printf("Fs.cd('%v'), current is root\n", path)
	} else {
		current = fs.Current
		fmt.Printf("Fs.cd('%v') current is '%v'\n", path, current.Path())
	}
	splits := strings.Split(path, "/")
	for _, subpath := range splits {
		if subpath == "" {
			continue
		}
		fmt.Printf("cd '%v', subpath is '%v'\n", path, subpath)
		if subpath == ".." {
			current = current.Parent
		} else {
			current = current.GetChild(subpath)
		}
	}
	fs.Current = current
	return fs.Current
}

func (fs *FileSystem) PWD() string {
	return fs.Current.Path()
}

func (fs *FileSystem) Get(path string) *Entry {
	if path == "/" {
		return fs.Root
	}
	original := fs.Current
	var entry *Entry
	lastIndex := strings.LastIndex(path, "/")
	dir := path[0 : lastIndex+1]
	filename := path[lastIndex+1:]
	fmt.Printf("path='%v', dir='%v', filename='%v'\n", path, dir, filename)
	fs.CD(dir)
	entry = fs.Current.GetChild(filename)
	fs.Current = original
	return entry
}

type FILETYPE string

const DIR FILETYPE = "DIR"
const FILE FILETYPE = "FILE"

type Entry struct {
	Type     FILETYPE
	Name     string
	Size     int
	Parent   *Entry
	Children []*Entry
}

func NewEntry(input string) *Entry {
	e := Entry{}
	e.Children = make([]*Entry, 0)
	splits := strings.Split(input, " ")
	size, err := strconv.Atoi(splits[0])
	if err != nil {
		// then it's a dir
		e.Name = splits[1]
		e.Type = DIR
		e.Size = 0
		fmt.Printf("input=%v, size=%v, err=%v, '%v' is a '%v'\n", input, size, err, e.Name, e.Type)
	} else {
		// then it's a file
		e.Name = splits[1]
		e.Type = FILE
		e.Size = size
		fmt.Printf("input=%v, size=%v, err=%v, '%v' is a '%v'\n", input, size, err, e.Name, e.Type)
	}
	return &e
}

func (e *Entry) TotalSize() int {
	total := 0
	for _, child := range e.Children {
		if child.Type == FILE {
			total += child.Size
		} else {
			total += child.TotalSize()
		}
	}
	return total
}

func (e *Entry) GetChild(name string) *Entry {
	for _, e := range e.Children {
		if e.Name == name {
			return e
		}
	}
	return nil
}

func (e *Entry) BuildPathToRoot() []*Entry {
	current := e
	entries := make([]*Entry, 0)
	entries = append(entries, current)
	for {
		current = current.Parent
		if current == nil {
			break
		}
		entries = append(entries, current)
	}
	return entries
}

func (e *Entry) Path() string {
	entries := e.BuildPathToRoot()
	path := ""
	for _, entry := range entries {
		if path == "" {
			path = fmt.Sprintf("%v", entry.Name)
		} else {
			path = fmt.Sprintf("%v/%v", entry.Name, path)
		}
	}
	// TODO so nasty
	path = strings.ReplaceAll(path, "//", "/")
	return path
}

func (e *Entry) Add(child *Entry) {
	e.Children = append(e.Children, child)
	child.Parent = e
}
