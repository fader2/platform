// Code generated by vfsgen; DO NOT EDIT.

// +build deploy_build

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	mustUnmarshalTextTime := func(text string) time.Time {
		var t time.Time
		err := t.UnmarshalText([]byte(text))
		if err != nil {
			panic(err)
		}
		return t
	}

	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: mustUnmarshalTextTime("2017-06-25T12:57:59Z"),
		},
		"/boltdb": &vfsgen۰DirInfo{
			name:    "boltdb",
			modTime: mustUnmarshalTextTime("2017-06-25T13:00:06Z"),
		},
		"/boltdb/bootstrap.lua": &vfsgen۰FileInfo{
			name:    "bootstrap.lua",
			modTime: mustUnmarshalTextTime("2017-07-09T08:40:56Z"),
			content: []byte("\x6c\x6f\x63\x61\x6c\x20\x62\x6f\x6c\x74\x64\x62\x20\x3d\x20\x72\x65\x71\x75\x69\x72\x65\x28\x22\x62\x6f\x6c\x74\x64\x62\x22\x29\x0a\x62\x6f\x6c\x74\x64\x62\x2e\x6f\x70\x65\x6e\x73\x28\x7b\x54\x50\x4c\x5f\x46\x52\x41\x47\x4d\x45\x4e\x54\x53\x5f\x42\x55\x43\x4b\x45\x54\x5f\x4e\x41\x4d\x45\x7d\x29\x0a\x0a\x70\x72\x69\x6e\x74\x28\x22\x62\x6f\x6f\x74\x73\x74\x72\x61\x70\x20\x62\x6f\x6c\x74\x64\x62\x22\x29\x0a\x63\x66\x67\x3a\x73\x65\x74\x28\x22\x62\x6f\x6c\x74\x64\x62\x22\x2c\x20\x74\x72\x75\x65\x29"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/boltdb"].(os.FileInfo),
	}
	fs["/boltdb"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/boltdb/bootstrap.lua"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// We already imported "compress/gzip" and "io/ioutil", but ended up not using them. Avoid unused import error.
var _ = gzip.Reader{}
var _ = ioutil.Discard

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}