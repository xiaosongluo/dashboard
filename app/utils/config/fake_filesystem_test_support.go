package config

import (
	"os"
	"path"
	"strings"
	"syscall"
)

// FakeFile representss a fake instance of the File interface.
//
// Methods from FakeFile act like methods in os.File, but instead of working in
// a real file, them work in an internal string.
//
// An instance of FakeFile is returned by RecordingFs.Open method.
type FakeFile struct {
	content string
	current int64
	r       *strings.Reader
	f       *os.File
}

func (f *FakeFile) reader() *strings.Reader {
	if f.r == nil {
		f.r = strings.NewReader(f.content)
	}
	return f.r
}

func (f *FakeFile) Close() error {
	f.current = 0
	if f.f != nil {
		f.f.Close()
		f.f = nil
	}
	return nil
}

func (f *FakeFile) Read(p []byte) (n int, err error) {
	n, err = f.reader().Read(p)
	f.current += int64(n)
	return
}

func (f *FakeFile) ReadAt(p []byte, off int64) (n int, err error) {
	n, err = f.reader().ReadAt(p, off)
	f.current += off + int64(n)
	return
}

func (f *FakeFile) Seek(offset int64, whence int) (int64, error) {
	var err error
	f.current, err = f.reader().Seek(offset, whence)
	return f.current, err
}

func (f *FakeFile) Fd() uintptr {
	if f.f == nil {
		var err error
		p := path.Join(os.TempDir(), "testing-fs-file.txt")
		f.f, err = os.Create(p)
		if err != nil {
			panic(err)
		}
	}
	return f.f.Fd()
}

func (f *FakeFile) Stat() (fi os.FileInfo, err error) {
	return
}

func (f *FakeFile) Write(p []byte) (n int, err error) {
	n = len(p)
	diff := f.current - int64(len(f.content))
	if diff > 0 {
		f.content += strings.Repeat("\x00", int(diff)) + string(p)
	} else {
		f.content = f.content[:f.current] + string(p)
	}
	return
}

func (f *FakeFile) WriteString(s string) (ret int, err error) {
	return f.Write([]byte(s))
}

func (f *FakeFile) Truncate(size int64) error {
	f.content = f.content[:size]
	return nil
}

// RecordingFs implements the Fs interface providing a "recording" file system.
//
// A recording file system is a file system that does not execute any action,
// just record them.
//
// All methods from RecordingFs never return errors.
type FakeFs struct {
	actions []string
	files   map[string]*FakeFile

	// FileContent is used to provide content for files opened using
	// RecordingFs.
	FileContent string
}

// HasAction checks if a given action was executed in the filesystem.
//
// For example, when you call the Open method with the "/tmp/file.txt"
// argument, RecordingFs will store locally the action "open /tmp/file.txt" and
// you can check it calling HasAction:
//
//     rfs.Open("/tmp/file.txt")
//     rfs.HasAction("open /tmp/file.txt") // true
func (r *FakeFs) HasAction(action string) bool {
	for _, a := range r.actions {
		if action == a {
			return true
		}
	}
	return false
}

func (r *FakeFs) open(name string, read bool) (*os.File, error) {
	if r.files == nil {
		r.files = make(map[string]*FakeFile)
		if r.FileContent == "" && read {
			return nil, syscall.ENOENT
		}
	} else if f, ok := r.files[name]; ok {
		f.r = nil
		return f.f, nil
	} else if r.FileContent == "" && read {
		return nil, syscall.ENOENT
	}
	fil := &FakeFile{content: r.FileContent}
	r.files[name] = fil
	return fil.f, nil
}

// Open records the action "open <name>" and returns an instance of FakeFile
// and nil error.
func (r *FakeFs) Open(name string) (*os.File, error) {
	r.actions = append(r.actions, "open "+name)
	return r.open(name, true)
}
