package day7

import (
	"advent-of-code-go/util"
	"strconv"
	"strings"
)

const day = "day7"

var inputFile string

func Solve(easy bool) (name string, res string, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

	res, err = partOne(lines)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

const command = '$'
const cd, ls, dir, root, out = "cd", "ls", "dir", "/", ".."

type fileSystemObject interface {
	getName() string
	getSize() int64
}

type file struct {
	size   int64
	name   string
	parent *directory
}

func (f *file) getName() string {
	return f.name
}

func (f *file) getSize() int64 {
	return f.size
}

type directory struct {
	size   int64
	name   string
	files  []fileSystemObject
	parent *directory
}

func (d *directory) getName() string {
	return d.name
}

func (d *directory) getSize() int64 {
	return d.size
}

func (d *directory) findChildDirectory(s string) *directory {
	for _, child := range d.files {
		if c, ok := child.(*directory); ok && c.name == s {
			return c
		}
	}
	return nil
}

func (d *directory) addChild(fs fileSystemObject) {
	for _, child := range d.files {
		if child.getName() == fs.getName() { // already have this child
			return
		}
	}
	d.files = append(d.files, fs)
}

func parseFileSystemObject(p *directory, s string) fileSystemObject {
	res := strings.Split(s, " ")

	if res[0] == dir {
		return &directory{
			name:   res[1],
			parent: p,
		}
	}

	size, _ := strconv.ParseInt(res[0], 10, 64)
	return &file{
		name:   res[1],
		size:   size,
		parent: p,
	}
}

func partOne(lines []string) (string, error) {
	rootDir := buildFileSystem(lines)

	dirs := findDirectories(rootDir, 100000)

	sum := int64(0)
	for _, d := range *dirs {
		sum += d.size
	}
	return strconv.FormatInt(sum, 10), nil
}

func buildFileSystem(lines []string) *directory {
	rootDir := &directory{
		name:   root,
		parent: nil,
	}

	var currentDir *directory
	for _, line := range lines {
		res := strings.Split(line, " ")
		if line[0] == command {
			if res[1] == ls {
				continue
			}
			to := res[2]
			switch to {
			case root:
				currentDir = rootDir
			case out:
				currentDir = currentDir.parent
			default:
				currentDir = currentDir.findChildDirectory(to)

			}
		} else {
			fs := parseFileSystemObject(currentDir, line)
			currentDir.addChild(fs)
		}
		if currentDir == nil {
			break
		}
	}
	return rootDir
}

func findDirectories(rootDir *directory, upTo int64) *[]*directory {
	res := []*directory{}
	calculateSizes(rootDir)
	find(rootDir, upTo, &res)
	return &res
}

func calculateSizes(dir *directory) {
	for _, object := range dir.files {
		if o, ok := object.(*file); ok {
			o.parent.size += o.size
		} else {
			d := object.(*directory)
			calculateSizes(d)
			d.parent.size += d.size
		}
	}
}

func find(dir *directory, upTo int64, res *[]*directory) {
	for _, object := range dir.files {
		if o, ok := object.(*directory); ok {
			if o.size <= upTo {
				*res = append(*res, o) // TODO might be problematic and a set needed
			}
			find(o, upTo, res)
		}
	}
}
