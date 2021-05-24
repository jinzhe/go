package file

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// List of files
func List(params ...string) (files []string, e error) {
	files = make([]string, 0, 30)
	suffix := ""
	if len(params) > 1 {
		suffix = strings.ToUpper(params[1])
	}
	e = filepath.Walk(params[0], func(path string, file os.FileInfo, e error) error {
		if file.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(file.Name()), suffix) {
			files = append(files, path)
		}
		return nil
	})
	return files, e
}

// ModTime
func ModTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// GetFileSize get file size as how many bytes
func Size(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// Suffix
func Suffix(name string) string {
	fileName := strings.Split(name, ".")
	return fileName[len(fileName)-1]
}

// Remove
func Remove(file string) error {
	return os.Remove(file)
}

// Rename
func Rename(file string, to string) error {
	return os.Rename(file, to)
}

// Copy
func Copy(from, to string) error {
	si, e := os.Lstat(from)
	if e != nil {
		return e
	}
	if si.Mode()&os.ModeSymlink != 0 {
		target, e := os.Readlink(from)
		if e != nil {
			return e
		}
		return os.Symlink(target, to)
	}

	sr, e := os.Open(from)
	if e != nil {
		return e
	}
	defer sr.Close()

	dw, e := os.Create(to)
	if e != nil {
		return e
	}
	defer dw.Close()

	if _, e = io.Copy(dw, sr); e != nil {
		return e
	}
	if e = os.Chtimes(to, si.ModTime(), si.ModTime()); e != nil {
		return e
	}
	return os.Chmod(to, si.Mode())
}

// Create
func Create(path string, content string) (int, error) {
	f, e := os.Create(path)
	if e != nil {
		return 0, e
	}
	defer f.Close()
	return f.WriteString(content)
}

// Read
func Read(path string) (string, error) {
	if IsDir(path) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(path)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, e := os.Stat(path)
	return e == nil || os.IsExist(e)
}

// IsDir
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

// CreateDir
func CreateDir(path string) bool {
	if IsDir(path) {
		return true
	}
	e := os.MkdirAll(path, 0777)
	return e == nil
}
