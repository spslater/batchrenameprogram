package brp

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type FileInfo struct {
	dir  string
	name string
	ext  string
}

func NewFileInfo(name string) FileInfo {
	var dir string = filepath.Dir(name)
	var ext string = filepath.Ext(name)
	var base string = strings.TrimSuffix(filepath.Base(name), ext)

	return FileInfo{dir: dir, name: base, ext: ext}
}

func (info FileInfo) Fullname() string {
	if info.ext == "" {
		return fmt.Sprintf("%v/%v", info.dir, info.name)
	}
	return fmt.Sprintf("%v/%v%v", info.dir, info.name, info.ext)
}

type FileHistory struct {
	original FileInfo   // Name of file when started program
	current  FileInfo   // Name of file on the file system right now (after saving before quitting)
	previous FileInfo   // Name in program before most recent change
	rename   FileInfo   // What will be renamed to when saving
	namelist []FileInfo // List of all previous rename steps (including saves)
}

func (file FileHistory) Output() string {
	return fmt.Sprintf("{\n %v\n %v\n %v\n %v\n %v\n}", file.original, file.current, file.previous, file.rename, file.namelist)
}

func NewFileHistory(name string) *FileHistory {
	var orig FileInfo = NewFileInfo(name)
	var curr FileInfo = orig
	var prev FileInfo = orig
	var rename FileInfo = orig
	return &FileHistory{original: orig, current: curr, previous: prev, rename: rename}
}

func (file FileHistory) Display() string {
	return fmt.Sprintf("%v\n%v\n", file.current.Fullname(), file.rename.Fullname())
}

func (file FileHistory) PeakName() string {
	return file.rename.name
}

func (file *FileHistory) pushName(base string) {
	file.pushInfo(base, "")
}

func (file *FileHistory) pushExt(ext string) {
	file.pushInfo("", ext)
}

func (file *FileHistory) pushInfo(base string, ext string) {
	if base == "" {
		base = file.rename.name
	}
	if ext == "" {
		ext = file.rename.ext
	}

	var new_name FileInfo = FileInfo{dir: file.rename.dir, name: base, ext: ext}

	file.namelist = append(file.namelist, file.previous)
	file.previous = file.rename
	file.rename = new_name
}

func (file *FileHistory) popInfo() bool {
	var length int = len(file.namelist)
	if length == 0 {
		return false
	}

	file.rename = file.previous
	file.previous = file.namelist[length-1]
	file.namelist = file.namelist[:length-1]

	return true
}

func (file *FileHistory) ChangeExt(new_ext string, pattern string, match_ext bool) {
	if new_ext[0:1] != "." {
		new_ext = string("." + new_ext)
	}
	if pattern == "" {
		file.pushExt(new_ext)
	} else {
		var check string = file.rename.name
		if match_ext {
			check = file.rename.ext
		}
		re, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Printf("regex err: %v\n", err)
			file.Noop()
		} else if re.MatchString(check) {
			file.pushExt(new_ext)
		} else {
			file.Noop()
		}
	}
}

func (file *FileHistory) Replace(find string, repl string) {
	findp, err := regexp.Compile(find)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	var new_name string = findp.ReplaceAllString(file.rename.name, repl)
	file.pushName(new_name)
}

func insertinto(orig string, idx int, ins string) string {
	var len = len(orig)
	var into int
	if idx > len { // idx is positive and past end of name
		into = len
	} else if idx >= 0 { // idx is positive
		into = idx
	} else if len*-1 > idx { // idx is negative and past beginnig of name
		into = 0
	} else { // idx is negative (-5, word is length 10)
		into = len + idx
	}
	return orig[:into] + ins + orig[into:]
}

func (file *FileHistory) Insert(idx int, insert string) {
	var tmp string = insertinto(file.rename.name, idx, insert)
	file.pushName(tmp)
}

func (file *FileHistory) Noop() {
	file.namelist = append(file.namelist, file.previous)
	file.previous = file.rename
}

func (file *FileHistory) Undo() bool {
	return file.popInfo()
}

func (file *FileHistory) Reset() {
	// TODO: Bug? If reset current, can't save changes
	file.current = file.original
	file.namelist = make([]FileInfo, 1)
	file.previous = file.original
	file.rename = file.previous
}

func (file *FileHistory) Save() {
	file.move()
	file.current = file.rename
}

func (file *FileHistory) move() {
	os.Rename(file.current.Fullname(), file.rename.Fullname())
}

func (file FileHistory) History() string {
	var sb strings.Builder
	sb.WriteString(file.Display())
	var num int = len(file.namelist)
	if file.previous != file.original {
		num += 1
	}
	if num == 0 {
		sb.WriteString("  NA\n")
		return sb.String()
	}
	var pad int = len(strconv.Itoa(num))
	for _, name := range file.namelist {
		sb.WriteString(fmt.Sprintf("%*d %v\n", pad, num, name.Fullname()))
		num--
	}
	sb.WriteString(fmt.Sprintf("%*d %v\n", pad, num, file.previous.Fullname()))
	sb.WriteString(strings.Repeat("~", 20))
	return sb.String()
}

func (file *FileHistory) ChangeCase(cases []CaseId) {
	var name string = file.rename.name
	for _, c := range cases {
		name = ChangeCase(name, c)
	}
	file.pushName(name)
}

func (file FileHistory) MatchName(re *regexp.Regexp) bool {
	return re.MatchString(file.previous.name)
}

func (file FileHistory) MatchExt(re *regexp.Regexp) bool {
	return re.MatchString(file.previous.ext)
}
