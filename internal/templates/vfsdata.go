// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

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

// FileSystem contains project templates.
var FileSystem = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 6, 17, 18, 6, 45, 882059488, time.UTC),
		},
		"/error.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "error.tmpl",
			modTime:          time.Date(2020, 6, 17, 18, 6, 45, 882059488, time.UTC),
			uncompressedSize: 826,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xd1\x6e\xd3\x4c\x10\x85\xef\xf3\x14\xfb\xfb\xe2\xbf\x81\xec\x26\x81\xd2\x02\x5e\xa3\x02\xa5\xa5\xa8\x6a\x49\xda\xd2\x5e\x4e\xd7\x13\x7b\xd3\xf5\xae\xb5\x33\x71\xea\x58\x91\x78\x07\xde\x90\x27\x41\x4e\x55\x10\x37\x95\x40\x5c\xad\xce\x59\xcd\x99\x6f\xa4\x93\xfe\x97\x07\xc3\x6d\x8d\xa2\xe4\xca\x65\x83\xb4\x7f\x84\x03\x5f\x68\xf4\xd9\x20\xad\x90\x41\x98\x12\x22\x21\xeb\x25\xcf\x87\x7b\x0f\xa6\x87\x0a\x75\x63\x71\x55\x87\xc8\xc2\x04\xcf\xe8\x59\x27\x2b\x9b\x73\xa9\x73\x6c\xac\xc1\xe1\x56\x3c\xb5\xde\xb2\x05\x37\x24\x03\x0e\xf5\x38\xc9\x06\x29\x5b\x76\x98\x75\x9d\x9c\x31\xf0\x92\xde\x85\x1c\x37\x1b\xd1\x75\x25\x73\x7d\x6f\x9d\xe3\x1d\x8b\xdf\xbf\xbf\x7f\xfd\x26\xac\x27\x06\xc9\x15\x97\xb1\xc8\x65\x8e\x4d\xaa\xee\xb3\x06\xa9\xb3\xfe\x56\x44\x74\x9a\xb8\x75\x48\x25\x22\x8b\x32\xe2\x5c\xf7\xa1\xf4\x4a\x29\x93\xfb\x05\x49\xe3\xc2\x32\x9f\x3b\x88\x28\x4d\xa8\x14\x2c\xe0\x4e\x39\x7b\x43\xca\x87\x58\x81\xb3\x6b\x54\x7b\x72\x24\xc7\xbf\xb4\xac\xac\x97\x86\x48\x58\xcf\x58\x44\xcb\xad\x4e\xa8\x84\xc9\xce\x8b\xa1\xdb\xdb\x39\xad\xce\x4e\x17\x8d\x75\x6a\x76\xda\x5c\xf2\xb3\x23\x3f\x9b\x2d\xd6\x1f\xc6\xe7\x17\x27\xed\xf9\x4b\xbc\x1c\x99\xc9\xdb\xf5\xc1\xe1\xfa\x42\x27\xc2\xc4\x40\x14\xa2\x2d\xac\xd7\xe0\x83\x6f\xab\xb0\xa4\x7f\x40\x4e\xb7\xe8\x90\x83\x57\x13\x39\x92\xcf\x7f\xca\x47\xb8\x27\xd7\x9f\xa7\xc7\x27\x57\xef\x77\xeb\x8f\xfb\x67\x47\xf6\x2a\x8e\xe8\x49\xe3\xa6\x5f\xf6\x77\x0f\xaf\x8f\x0f\x3e\x8d\xf6\xa7\x9e\x76\x6f\x27\x74\x73\x74\xfd\x37\xdc\xc9\x03\xf8\x3c\x78\x26\x59\x84\x50\x38\x84\xda\xd2\x16\xdc\x10\xbd\x99\x43\x65\x5d\xab\xa7\xe0\x70\x05\xed\xff\x50\xd5\xaf\x73\x4b\xb5\x83\x56\xd3\x0a\xea\xe4\xb1\xf4\xae\x03\x22\xe4\x33\xe0\x52\x24\x0a\x63\x0c\xb1\x3f\x33\xd9\x6c\xfa\x6e\x82\xf5\xc2\x38\x20\xd2\x7d\x27\xc1\x7a\x8c\x7d\xad\xc7\x7f\xd8\xb6\x54\x95\xe3\x6c\x90\xd6\xfd\xd8\x09\x12\x41\xb1\x35\xeb\x6c\x90\xaa\x7e\x47\xf6\x23\x00\x00\xff\xff\x1c\xcb\xd2\xd5\x3a\x03\x00\x00"),
		},
		"/index.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "index.tmpl",
			modTime:          time.Date(2019, 6, 6, 4, 45, 14, 653991180, time.UTC),
			uncompressedSize: 1096,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x61\x6f\x1c\x35\x10\xfd\x7e\xbf\xc2\x1c\x12\x4a\x08\x6b\x5f\x0e\x4a\xd2\xb0\xde\x53\x50\x4b\xd3\x42\x95\x70\x69\x4a\x03\x87\xd0\xc4\x9e\xbd\x9d\x8b\xd7\x5e\xec\xd9\xdd\xde\x95\xfe\x77\xb4\x97\x4b\x10\x12\xca\x07\xd4\x2f\x6b\xcd\xbe\x67\xfb\xf9\xcd\x9b\xfc\x33\x1b\x0c\xaf\x1b\x14\x15\xd7\xae\x18\xe5\xc3\x22\x1c\xf8\xa5\x46\x5f\x8c\xf2\x1a\x19\x84\xa9\x20\x26\x64\xdd\x72\x99\x1d\xdf\xff\xf4\x50\xa3\xee\x08\xfb\x26\x44\x16\x26\x78\x46\xcf\x7a\xdc\x93\xe5\x4a\x5b\xec\xc8\x60\xb6\x2d\xbe\x22\x4f\x4c\xe0\xb2\x64\xc0\xa1\x3e\x1c\x17\xa3\x9c\x89\x1d\x16\xe4\x13\x83\xe4\x9a\xab\xb8\xb4\xd2\x62\x97\xab\x3b\x60\x94\x3b\xf2\xb7\x22\xa2\xd3\x89\xd7\x0e\x53\x85\xc8\xa2\x8a\x58\xea\x8a\xb9\x49\x27\x4a\x19\xeb\x57\x49\x1a\x17\x5a\x5b\x3a\x88\x28\x4d\xa8\x15\xac\xe0\xbd\x72\x74\x93\x94\x0f\xb1\x06\x47\x1b\x54\xc7\x72\x22\x0f\xff\xa9\x65\x4d\x5e\x9a\x94\x04\x79\xc6\x65\x24\x5e\xeb\x71\xaa\x60\xfa\xe4\xdb\xcc\x1d\x3f\x39\xaf\x2f\xce\x57\x1d\x39\x75\x79\xde\xbd\xe5\xaf\xcf\xfc\xe5\xe5\x6a\xf3\xc3\xe1\x9b\xab\xd7\xeb\x37\x4f\xf1\xed\xc4\x4c\xbf\xdf\x3c\x7f\xb1\xb9\xd2\x63\x61\x62\x48\x29\x44\x5a\x92\xd7\xe0\x83\x5f\xd7\xa1\x4d\x9f\x40\x79\xba\x45\x87\x1c\xbc\x9a\xca\x89\xfc\xe6\xa1\x7c\x44\xf7\xf4\xfa\xe7\xf9\xab\xd7\xef\x9e\x1d\x35\x2f\x4f\x2f\xce\xe8\x5d\x9c\xa4\x83\xce\xcd\x7f\x39\x3d\x7a\x71\xfd\xea\xf9\x8f\x93\xd3\xb9\x4f\x47\xb7\xd3\x74\x73\x76\xfd\x7f\x74\x8f\xef\x85\x97\xc1\x73\x92\xcb\x10\x96\x0e\xa1\xa1\xb4\x15\x6e\x52\x9a\x95\x50\x93\x5b\xeb\x39\x38\xec\x61\xfd\x05\xd4\xcd\x77\x96\x52\xe3\x60\xad\x53\x0f\xcd\xf8\xb1\xd3\x3f\x7c\x80\x94\x90\x2f\x80\x2b\x31\x56\xe4\x2d\xbe\x1f\x9e\x39\xfe\xf8\x71\x08\x1a\x90\x17\xc6\x41\x4a\x7a\x08\x18\x90\xc7\x58\x8c\xf2\x32\xc4\x5a\x80\x61\x0a\x5e\xab\x46\xd4\xc8\x55\xb0\x7a\x89\xbc\xe3\xda\xd0\x7b\x17\xc0\x66\x03\xb1\x18\xe5\x96\xba\x1d\x12\x43\x3f\x88\x81\x1b\x74\xa2\x0c\x51\xb7\xd1\x15\x2f\x87\x14\x2e\x23\xd4\xe2\x6a\xfe\xd3\x89\xc8\xd5\x16\x2e\x46\x39\xf9\xa6\x65\x31\x0c\xc7\xc0\xdb\x9d\xd0\x66\x65\xeb\xdc\x5d\xb2\xef\x66\x60\xc0\xc8\x6e\x97\x88\x7f\xb6\x14\xd1\x8a\xc6\x81\xc1\x2a\x38\x8b\xf1\xa1\xef\x7d\xdf\x4b\xba\xbf\x6b\xeb\x5d\xa3\xa4\x94\xa2\x01\x66\x8c\x7e\xe7\xf3\xec\x44\xa9\xbd\xbe\xef\x17\xf2\xaf\x7a\x21\xf7\x67\x0f\x3b\x16\xdb\x2d\x7b\xf2\x40\xed\xcf\x1a\xf5\xdb\x69\xf6\x2b\x64\x9b\x49\xf6\x74\x91\xfd\xf1\xfb\x81\x9a\xed\x2d\x66\xf2\xcb\xfd\xd9\xde\xe7\xc3\x77\x70\x5c\x59\xea\xfe\xfd\x88\xd4\xde\xd4\xc4\xa2\x03\xd7\xa2\x7e\xb6\xf3\x68\x60\xee\x6c\x52\x83\xdd\xc5\x28\x4f\x26\x52\xc3\xc2\x62\x89\x51\xa4\x68\xfe\xb3\x47\xab\x6d\x8b\x72\x75\x47\x2e\xfe\x0e\x00\x00\xff\xff\xa9\xdd\x99\x0f\x48\x04\x00\x00"),
		},
		"/post.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "post.tmpl",
			modTime:          time.Date(2019, 6, 6, 4, 45, 54, 24178631, time.UTC),
			uncompressedSize: 992,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xd2\xdd\x6e\xd3\x4a\x10\x07\xf0\xfb\x3c\xc5\x1e\xab\x17\xe7\xe8\xd4\x76\x12\x28\x2d\x95\xd7\x52\xa0\xa5\x1f\x50\x35\xa4\x1f\xb4\x97\xdb\xf5\xc4\x9e\x74\xbd\x6b\x76\x26\x69\x1d\xcb\x12\xef\xc0\x1b\xf2\x24\xc8\xa4\x05\xaa\x8a\x22\x21\xae\xac\xff\x7a\xc6\xf3\xb3\x76\x92\x7f\x32\xa7\xb9\xae\x40\x14\x5c\x9a\xb4\x97\x74\x0f\x61\x94\xcd\x25\xd8\xb4\x97\x94\xc0\x4a\xe8\x42\x79\x02\x96\x73\x9e\x86\x5b\xf7\x87\x56\x95\x20\x17\x08\x37\x95\xf3\x2c\xb4\xb3\x0c\x96\x65\x70\x83\x19\x17\x32\x83\x05\x6a\x08\xbf\x85\x75\xb4\xc8\xa8\x4c\x48\x5a\x19\x90\x83\x20\xed\x25\x8c\x6c\x20\x6d\x9a\x68\xec\x88\x0f\x76\xda\x56\x7c\xf9\xf4\x59\xa0\x25\x56\x11\x97\x5c\xf8\x3c\x8b\x32\x58\x24\xf1\xaa\xb0\x97\x18\xb4\xd7\xc2\x83\x91\xc4\xb5\x01\x2a\x00\x58\x14\x1e\xa6\xb2\x60\xae\x68\x3b\x8e\x75\x66\x67\x14\x69\xe3\xe6\xd9\xd4\x28\x0f\x91\x76\x65\xac\x66\xea\x36\x36\x78\x45\xb1\x75\xbe\x54\x06\x97\x10\x6f\x45\xfd\x68\xf0\x23\x47\x25\xda\x48\x13\x09\xb4\x0c\xb9\x47\xae\x65\x40\x85\x1a\x6e\xbc\x08\xcd\xd6\xc6\x71\x39\x3e\x9e\x2d\xd0\xc4\x27\xc7\x8b\x73\x7e\xb6\x6f\x4f\x4e\x66\xcb\x37\x83\xd3\xb3\xa3\xfa\xf4\x25\x9c\xf7\xf5\xf0\xd5\x72\x77\x6f\x79\x26\x03\xa1\xbd\x23\x72\x1e\x73\xb4\x52\x59\x67\xeb\xd2\xcd\xe9\x2f\xc8\xe9\x1a\x0c\xb0\xb3\xf1\x30\xea\x47\xcf\xbf\xc7\x27\xdc\xc3\xcb\xf7\x93\xc3\xa3\x8b\x9d\xcd\xea\x60\x34\xde\xc7\x0b\xdf\xa7\xff\x17\x66\xf2\x61\xb4\xb9\x77\x79\xb8\xfb\xb6\x3f\x9a\x58\xda\xbc\x1e\xd2\xd5\xfe\xe5\x9f\xb8\x9b\x46\x11\x01\x8f\x15\x17\x22\x88\x2b\x47\xdc\x39\x82\xb6\xed\x36\x43\xa1\x15\xda\x28\x22\xd9\x6d\x84\x42\x0b\x3e\xed\x35\x4d\x28\x70\x2a\xe0\xa3\xf8\xd7\x80\x15\xd1\x0e\x52\x65\x54\x7d\x36\x79\x47\xff\x89\x41\xdb\xf6\x12\x2c\xf3\xbb\xb6\xee\x7b\x21\x96\x2a\x07\x41\x5e\xcb\xa0\x69\xd0\x66\x70\xfb\xa0\x49\xf4\xdb\x36\x10\xca\xf0\x4f\xaf\x47\x5a\x03\x11\x5e\xa1\x41\xae\x5f\xab\x8a\xd1\xd9\xbb\x42\x0f\x53\xf0\x1e\x7c\xe5\x0c\xea\x5a\x5a\x17\xde\x9f\xac\x68\x60\x08\x3a\xc4\xdc\x3c\x32\xd0\xaa\xc2\x2b\x9b\x83\x58\xc3\xec\x76\x5d\xac\xad\x70\xdb\xf2\x01\xa9\xeb\x37\x98\x3e\xf9\x23\xd1\x23\xf5\xda\x2f\xd8\xdd\xa4\xdf\xc8\x93\xd8\xe0\x1d\xdf\x66\xdd\xf4\x78\x6e\x1e\xe6\xee\x2e\xd2\xaf\x01\x00\x00\xff\xff\x51\x20\x29\x01\xe0\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/error.tmpl"].(os.FileInfo),
		fs["/index.tmpl"].(os.FileInfo),
		fs["/post.tmpl"].(os.FileInfo),
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
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
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

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
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
