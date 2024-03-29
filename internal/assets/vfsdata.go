// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package assets

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

// FileSystem contains project assets.
var FileSystem = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 10, 21, 2, 5, 28, 283002013, time.UTC),
		},
		"/.vfshash-assets.json": &vfsgen۰CompressedFileInfo{
			name:             ".vfshash-assets.json",
			modTime:          time.Time{},
			uncompressedSize: 216,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\x56\xd2\x4f\x2d\x2a\xca\x2f\xd2\x4b\x2e\x2e\x56\xb2\x82\x72\x74\x23\x12\x8d\xca\xbc\xcc\x03\x8a\xbd\xf2\x13\xdd\x75\x0b\x03\x32\xc1\xb2\x3a\x4a\xfa\x69\x89\x65\x99\xc9\xf9\x79\x7a\x99\xc9\xf9\x20\xc5\x50\xae\xae\x57\x59\x66\x55\x6a\x58\x7e\x62\x52\x6a\x89\x69\x61\x48\x69\x32\x58\x5e\x47\x49\x3f\x33\x2f\x25\xb5\x02\x66\x32\x98\x83\xd3\x64\x88\xd2\x2c\x24\x95\x29\xa9\xfe\x96\x21\x11\xae\xe9\x21\x15\xae\xae\x49\xc9\xae\x51\x20\x49\x1d\x25\xfd\x82\xfc\xe2\x12\x98\x91\x20\xb6\x6e\x60\x55\x98\x53\x40\x81\xa1\x45\x61\xbe\x79\x66\xa6\x65\xba\x37\x58\xb2\x16\x10\x00\x00\xff\xff\xc7\x8b\xf4\x50\xd8\x00\x00\x00"),
		},
		"/error-Xa2vJ7PsJoaG-qPi.css": &vfsgen۰FileInfo{
			name:    "error-Xa2vJ7PsJoaG-qPi.css",
			modTime: time.Date(2019, 5, 9, 13, 1, 15, 900399520, time.UTC),
			content: []byte("\x68\x74\x6d\x6c\x2c\x20\x62\x6f\x64\x79\x20\x7b\x0a\x09\x6d\x69\x6e\x2d\x68\x65\x69\x67\x68\x74\x3a\x20\x31\x30\x30\x76\x68\x3b\x0a\x7d\x0a\x0a\x62\x6f\x64\x79\x20\x7b\x0a\x09\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x66\x6c\x65\x78\x3b\x0a\x09\x66\x6c\x65\x78\x2d\x64\x69\x72\x65\x63\x74\x69\x6f\x6e\x3a\x20\x63\x6f\x6c\x75\x6d\x6e\x3b\x0a\x09\x6a\x75\x73\x74\x69\x66\x79\x2d\x63\x6f\x6e\x74\x65\x6e\x74\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x7d"),
		},
		"/favicon-JvizeVoabet5qTuc.ico": &vfsgen۰CompressedFileInfo{
			name:             "favicon-JvizeVoabet5qTuc.ico",
			modTime:          time.Date(2019, 5, 9, 5, 44, 40, 794039265, time.UTC),
			uncompressedSize: 3438,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\xd6\x6b\x50\x53\x77\xde\x07\xf0\x24\x84\x4b\xb8\x0b\x01\x11\x2c\x37\xe5\x2a\x28\xf7\x20\xa4\x90\x40\x12\x42\x08\xe7\x7f\x11\x51\xe4\x22\x08\x92\x80\x06\x21\x80\x02\x2a\x04\x12\x90\x10\x14\x28\xa0\x54\x91\x50\x84\x27\x4a\x3b\x88\x97\x16\xb5\x7d\x9e\x7a\xc1\x0e\xda\xa7\x93\x96\xe9\x76\xb4\xab\x06\xb9\xb4\xda\x75\x67\x30\xba\xb0\xbb\xb3\xd5\x9d\xc3\xa5\xdb\x17\x3b\xfb\xda\xfd\xbe\xfa\xce\x2f\xff\xdf\x9b\x93\xcf\xfc\xcf\x49\x4a\x16\x44\x6f\xcd\xb3\xa7\xd8\x53\x16\x20\xc5\x84\xe9\x6d\xb1\xce\xdf\xc2\x35\x80\xe1\x16\x68\xe9\x1d\x66\xe3\x15\x66\xeb\x13\xe5\xb4\x81\xe5\x1c\xc8\x72\x0d\x8a\xb5\x0d\x4e\xb4\x09\xe3\x59\x45\xa5\x5a\x47\x88\x9c\x42\x93\x1c\xa2\xc5\x36\xf1\x79\x5b\xd8\x49\x56\xdc\x7c\xcb\x64\x89\xa5\x50\xea\x9e\xb2\x6f\x1d\x2e\xf1\xe5\x65\x44\x26\x02\x96\x00\xfb\x10\x15\x5e\x19\xf2\xb0\x54\x59\x68\x8e\x2c\x36\xad\x80\x43\xe4\xa4\x64\xee\xc1\xd9\x92\xf4\x3d\xb2\x8c\xdc\xc2\xcc\xfc\xe2\x6c\x69\x65\x9e\x44\xb6\xf7\x50\x8d\x35\xa1\xb6\x4f\xaf\x5f\x9b\xad\x74\xdb\xa3\x72\x95\xf6\x04\x97\x69\xc2\x73\x15\x58\xae\xdc\x56\xda\x5c\x50\xa3\x2c\x3b\xac\xac\xac\x6b\xde\xa2\xba\x58\xd9\x74\xb5\xf6\xc4\x75\xad\x56\xab\xd3\xe9\xce\x7f\x3c\x72\xe1\x93\x8b\x6d\x67\x6f\xb5\x9d\xbd\xad\x1b\xbe\x7b\xe1\xca\xb7\x7d\x5f\xdf\xbc\x71\xe7\x47\xbd\x5e\x7f\xe7\xfe\xa3\x3b\xf7\x9f\x4c\xfe\xf0\xd0\x60\x30\x7c\x3d\x39\xf3\x64\xe6\xa7\xa9\xe9\x99\xa9\x99\xd9\xb9\x67\x2f\xe7\x5f\x1a\x17\x17\x17\x17\x16\xff\xfe\xf6\xed\x3f\xde\xbe\x7d\x4b\x79\xc7\xe2\xf5\x37\x3a\x83\x49\x81\x14\xaf\xb7\x56\x04\x3f\x6d\x5b\x22\x17\xf2\xc3\x36\x87\x98\x50\x29\x14\x4a\x10\xf9\x3b\xf9\x37\x51\xcc\xdf\x28\xe0\x72\x52\x53\xe1\x6a\xa2\xa1\x00\x42\x00\x20\x14\xc0\x68\x08\x21\xdb\x82\x0f\xa1\x08\x00\x11\x84\x7c\x73\x36\x84\xfc\x35\x70\x83\x10\x00\x00\x92\x37\x42\x4f\x3e\x84\x2e\x5e\x90\xe2\x0b\x38\x1c\xb0\x9e\x0a\x7d\xd6\x41\xb8\xd5\x1b\x9a\x91\xdb\x00\x50\xa1\x77\x0c\x84\x6c\x7b\x68\x0a\xe2\x58\xac\x38\x40\x85\x21\x6c\x08\x21\xd7\x2e\x1c\xb0\xc4\x29\x2c\x10\xc6\xe0\x42\xd8\xcf\xa9\x37\xf2\x94\x4a\xf7\xa8\xfd\x69\xd7\x46\xa7\x18\x4c\x7f\xd8\xcf\xa1\x1a\x55\x13\x9d\xb6\x96\x16\xd7\xaa\xe4\xb3\x0c\x7b\x72\x52\x6d\xdc\xf8\x15\x9d\x11\xe8\x1f\x91\x5b\x3a\x7d\xcf\xd1\x1f\xf6\xbb\x54\x1b\xe9\x2d\xf6\x56\xc1\xe9\x5f\xb8\xd0\xe6\xee\x39\x6e\x80\xfd\xec\xea\xb0\x82\x89\x76\xbb\x80\xa8\x2a\x97\xb0\x82\x89\x4e\x3f\xd8\xc7\xae\x36\xce\xde\xeb\xf4\x8b\x87\x2b\x05\x54\xcd\x4f\xdd\x52\x9a\xdb\xb9\x07\x49\x57\x8a\x13\xbb\x9f\x53\x4f\x33\xb5\x64\xfa\x83\xbe\x95\x02\x69\xd4\x95\x67\x07\xfe\xf5\xa0\x9a\x5c\x36\x84\xfb\x44\xc4\xdb\x84\x0a\xde\xdb\xc4\xf5\x08\xe6\x3a\x85\x27\xb9\xb3\x51\xc0\xfb\x7c\xdf\xb8\xcc\xb0\xf7\xd3\xdd\xb9\xe9\xd6\x29\x25\x6e\xe9\x87\x58\xe2\xac\x2d\xa9\x85\xc1\x58\x1e\xbd\xad\x30\x25\xbb\x38\x6d\xf7\xc1\x1c\x49\x09\x92\x1c\x49\xdf\x5f\xb3\xab\xb4\x2e\x3f\xaf\x42\x52\x54\xb9\x67\x5f\x4d\x5e\x51\x5d\x8c\xac\x39\x46\xa6\x8e\x2d\xd5\x26\x96\x74\xf2\x65\xed\x82\x92\x0e\xd1\x81\x6e\x41\xf9\xa9\x2c\xe9\xa9\xac\xca\x36\x7c\xf8\x6c\xda\xe1\x33\x39\x07\x4f\xed\x92\xf7\x64\x1d\x19\x2a\x38\xd2\x5b\xd3\x50\x97\xdd\xd2\x2f\x6f\x6e\xa9\x69\x3c\x51\x77\xac\x53\xd2\x36\x58\x7c\xac\xbf\x44\x71\xa1\xb4\xa9\xbf\x42\x3d\x70\xa4\xe9\xa3\x86\xb6\x9e\x9e\x9e\x1e\xad\x56\xdb\xa5\x1d\x6b\x1a\x1a\xeb\xfb\xa8\x5f\xa7\xfb\x1f\x9d\x4e\x37\x3c\x7a\xf5\xd2\xd5\xab\xdd\xa3\x37\x07\xbf\xb8\x77\xe5\xea\xf8\x97\x9f\xe9\xf5\x7a\xfd\x1f\x7e\x7c\x6c\x30\x18\xbe\x7b\x3c\x3b\x3d\xfb\xf3\xf3\x17\x2f\x5e\x1a\x5f\xcd\xbf\x5e\x58\x5c\x98\x7f\x37\x29\x53\x96\xa0\x82\xe5\xa4\xa4\x80\xd5\xb0\x01\x17\x00\x82\x00\x80\x0b\xd8\x00\x00\x37\x8a\x00\x80\x24\x82\x48\x02\x40\x40\x73\x03\x40\xb0\x09\xb8\x08\x09\x82\x20\x92\x1d\x81\xbd\x00\x00\xa7\x58\x60\x1e\x4a\xc4\xc4\x10\x81\x54\xe0\xbc\x06\x80\xf8\x60\x40\x27\xb7\x09\xc2\x12\xac\x8b\x03\xc0\xcd\x1d\x98\x10\x21\xd1\xd1\x21\x84\x29\x70\x77\x03\x00\xf0\xcc\xb6\x10\xd1\x62\x51\x34\x11\xe5\xc7\x03\xa0\x8f\x5d\x3d\x3d\x4b\x6d\xb0\xb5\x76\x99\x28\xf5\xbc\x45\x72\xea\x63\x7f\xf7\x74\xb0\xb9\xd3\x81\xe1\xaa\x2a\xc5\xe3\xea\x13\xe4\xa4\x6c\x7a\xb6\xa9\xd9\xce\x43\xd8\x9a\xfb\xf4\xb6\x9a\x3c\x43\x2f\x9b\x9e\x65\xd8\x33\x7d\x0a\x21\xc5\x38\xae\xee\xf4\x04\x7d\xd1\x65\x1f\x98\x58\xd8\x3b\x04\x14\x67\x51\x68\xf4\x66\x47\x3f\xd0\x1b\x5d\x36\x3d\xae\xee\xf4\xdd\x22\x5f\x2a\x8e\x7e\x44\xd4\xfe\x07\xa3\x79\x77\x9a\x6c\x12\x55\x2b\xc5\xc9\x6d\x55\x33\xd1\xbb\xaa\xf9\xdf\x42\x75\xf6\x8d\x79\xcf\x9f\x84\xba\x3e\x0e\x07\x04\xf1\x22\x23\x20\x3b\x36\x8d\x15\x2b\x61\xc7\xe5\xc7\xec\x2c\xe2\xf3\x77\x11\x29\xe9\x59\xe2\x9d\x1c\x8e\x24\x49\x54\x98\x2a\xc8\x4b\x21\x8a\x71\x9a\x64\xd7\x2e\xf9\xee\x1c\x12\xe7\x81\x03\xb5\xe5\xf2\xa3\xc9\xe2\x8e\xb4\xf4\x63\x38\x4b\xb9\x73\x7b\x4b\x4e\x86\x72\x4f\xa6\x26\x6f\xf7\xf5\xc2\x5c\xa5\x4c\xd6\x58\x5c\xde\x2a\xad\xd5\x94\xca\x1b\xab\xaa\x9a\xeb\x8e\x9e\x3d\x22\x1f\x56\xed\x3d\xdc\x78\x54\xa5\x3e\x7c\xf1\xd2\x91\xe3\x6a\xf5\xc9\xce\x96\x36\xad\x56\xdb\xd4\x73\xee\xcc\xe9\xfe\xee\xc1\xe1\xee\xa1\x91\x81\xf6\x4b\x17\x3a\xae\xeb\x74\xba\x0b\x03\x9f\x0e\x5e\xbe\xf1\xe9\xf9\x2b\x97\x2e\x8d\x5d\xb9\x32\x36\xd2\x73\xfb\x7c\xef\xc4\xe8\xe9\x3b\x13\x8a\xae\xbb\x1f\x0e\x8e\x9f\x19\x9e\x3b\x3f\x76\xf3\xb3\x6f\xc7\x6f\x4c\xea\xf5\xfa\x87\xb7\x1f\x4e\x3e\x78\xf4\xe0\xc1\x23\x83\xc1\x30\x35\x33\x3b\xf3\x74\xf6\xc5\xf3\x3f\xcd\x1b\x5f\x19\x8d\xaf\x16\x17\x17\xff\x0b\xa0\x8a\x44\xbf\x41\x0d\x07\x91\xcb\x50\x23\x41\x38\x00\x80\x41\x42\xe5\x11\x04\x0f\x00\x01\x95\x01\x80\xc0\x1e\x38\x27\x93\x50\x93\x98\xc0\x56\x00\x80\xa3\x37\x30\x8b\x22\x82\x83\x09\x17\x3a\x70\x5a\x03\x00\xcb\x03\x58\x2d\x41\x4d\xb5\x04\x4e\x2c\x00\x18\x11\xc0\x9a\xf0\x0a\x0a\xf2\x4a\x35\x05\xfe\x0c\x00\x40\x82\xdb\x66\x22\x48\x2c\x0a\x4a\x0d\x34\x4f\x58\x86\x4a\x57\xa9\x5d\x4c\xcc\x2a\x38\x37\x3f\x5a\x86\x5a\xf4\x94\xae\x76\x74\xb7\x73\x28\x95\xe2\xc1\x64\xcd\x32\xd4\x82\xc6\x36\x66\xbc\x28\xeb\x66\xc8\x32\xd4\xef\x4b\xa7\x67\xef\x69\xdc\x13\x15\xd2\xaa\x97\xe3\xea\x4e\x57\x12\x2a\xcd\xc4\xac\xd5\x71\x2b\x14\x56\xbd\x9c\x6e\xfc\x0d\xaa\xa3\x5f\x6c\xc1\x2a\x54\x87\x8d\xc2\x76\x4f\x6e\x03\x9d\x53\xb1\x52\x9c\x18\xff\x01\x2a\xfa\x1d\xd4\x58\xef\xd0\x65\xa5\x31\x3b\x8b\x38\x1e\xa1\x7c\xbf\x98\xd4\x20\xf6\xfe\xcd\xf1\xe9\x31\xbc\x1c\x56\xb2\x74\x33\xaf\x30\x0c\x1c\x64\x65\x55\xb2\x33\xf7\x0a\x33\x25\xc4\x2e\x49\x6a\x56\x5d\x66\x52\x0d\x2e\x2c\xcf\xdf\x57\x51\x28\x3f\x58\x24\xcf\xa9\x51\x49\x6b\x35\x0d\x31\xf9\x8a\x1d\x09\x0d\x09\x45\xc7\x53\xcb\xdb\x52\x2b\x1b\xb6\x49\x54\xdb\x25\x0a\x2c\x3b\xb6\x4d\xd6\x0a\xcb\xdb\x76\x1f\xea\xda\x56\xd5\xb1\xf3\xe8\x99\x8c\xba\xb6\xfd\xc5\x8d\xf2\xba\xf6\xc2\x8a\xde\x12\xf9\x50\xc5\xa1\xde\x74\xe5\xa9\xaa\x63\xe7\xf3\x5a\x3f\x96\x9e\x18\x94\x2b\xbf\xc8\xd3\x7c\x5e\xd2\xff\x4d\x65\xa3\x5a\xd5\xd4\x59\xaf\x59\x36\xdc\x3d\x34\xd2\xdf\x7a\xf2\x5c\xe7\xd9\xff\x3b\xd9\xa7\xd3\xe9\x06\x2f\xdf\xf8\x72\xf4\xe2\xe7\xd7\x6f\xdc\xd5\x5d\xbf\x7d\xf9\x96\x5e\xaf\xff\x56\xff\xcd\xe4\x83\x47\x0f\x7f\x78\x68\x78\x48\x42\xfd\x79\x66\xee\xc5\xdc\xb3\x79\xe3\xab\xd7\xc6\x57\xaf\x5f\x2d\x90\x50\xdf\xfc\x75\xc9\xea\x9b\x77\x8d\xab\xf9\x1b\x05\x5a\x4e\x72\x32\x5a\x4d\x2c\x62\x23\x04\x21\x42\x6c\x14\x8b\x10\x32\xb1\x14\x20\x94\x08\x61\x22\x42\x02\x33\x13\x84\x04\xfe\xc8\x57\x48\x7e\x22\x24\x39\x21\x5b\x01\x42\xee\x01\x68\x5d\x18\x8c\x8a\x82\x5b\x29\x88\xe9\x8c\xd0\xfb\x81\x68\x0d\xb9\x0d\x21\x1d\x39\xc6\x21\x64\x12\x82\xcc\xe1\xda\xc8\xc8\xb5\xd0\x14\x79\x98\x20\x84\x12\x36\x44\xc3\x48\xb1\x28\x12\x46\x58\x70\x11\x1a\xe0\xa9\x69\x73\x7c\x13\x4f\x1a\x7d\x5c\x38\x4f\x11\x75\xf9\xa3\x01\xde\x25\xea\x90\x05\xd3\xda\xca\xb9\x4b\x91\x40\x6f\x39\x4e\x4e\xea\x69\x8d\x7c\x8d\xed\xfa\x0a\x95\x7d\x19\xcd\x92\x3c\x33\xaf\xa0\xf9\xdc\xb7\x31\x73\xe5\x08\xaf\xad\x31\xb5\x64\xba\xa3\x01\x4e\xbd\x7c\xe6\xab\xe3\xeb\xdc\x6a\x39\xd7\xca\x66\xcc\x98\xfe\x68\xf5\xad\x9e\x22\x5c\x7d\xab\x9b\x59\xbb\x84\xe4\xde\x54\x50\x37\x65\xad\x14\x27\x93\x01\x9e\xda\xc1\x47\xd4\xe5\x0f\x4f\xaf\x14\xf4\x1b\x54\xfc\x3b\xa8\x4b\x4a\x11\x09\x75\x87\x24\x66\x67\x91\x2c\x5e\x5c\x24\x4a\x2f\x4e\xde\xbe\x4f\xb4\xbd\x58\xb4\xe3\x80\x38\x7d\x1f\x91\x21\x13\x67\x55\xc0\xec\xaa\x6c\xa9\xb4\x56\x73\x2e\x96\xab\xd9\xbe\xa3\x55\xbc\xb7\x57\x9c\xdb\x0b\x0a\x54\xf9\x32\x45\x91\xbc\xb9\xb8\xfc\x78\x91\xbc\xad\xf4\x68\x73\xf5\xd1\xf6\x82\xd2\x0e\x49\x79\x57\x51\xf5\x27\x20\x63\xa0\xe8\xd0\x8c\x2c\x63\xee\x40\x45\x5b\x6d\xd3\xa9\xc3\xca\xb3\x35\xca\xc7\x9a\x8a\x6e\xd5\x09\xad\x52\x73\xea\x83\xae\xd3\xed\x9d\x2b\x0e\x97\x2e\xd2\xe1\xe6\xae\x91\x96\xee\xb1\xf6\x0f\xcf\x9f\x39\x37\x72\xba\xff\xca\xc9\xa1\x65\x93\x9f\x5f\x1c\xfd\xdf\xb1\x6b\x4f\x3e\xd0\x4e\x5c\xb8\x3a\xad\x3e\xf3\x5c\xfb\xc9\x8c\xf6\xf2\xc2\xf0\xe9\x9f\x74\xd7\xa6\x46\x6e\x4c\x2d\x41\x9d\x24\x2f\xd2\x27\x06\x83\xe1\xd9\xed\xff\x9f\x7d\x6c\x78\xad\xff\xe3\x9c\x61\xea\xc5\xb3\x5f\xfe\xfc\xfc\x97\x5f\xa7\x49\xb1\x0b\xaf\xff\xb2\x7a\xaf\xfe\xfa\x4e\x59\x35\x7f\xa3\xc0\xcb\x11\x8b\xf1\x6a\x22\x70\x24\xc6\x08\x61\x1c\x89\x23\x30\xc6\x74\x06\x0f\xe3\x04\x84\x12\x30\xe6\x59\xd2\x31\xe6\x39\x63\x57\x11\x09\x5a\xe4\x8d\x5d\x78\x18\xaf\xf5\xc5\x56\x5b\xd1\xa6\x4d\x68\x2b\x03\x6f\x70\xc2\x98\xe5\x8f\x6d\xc9\x6d\x84\x6c\xf0\xc6\x28\x8c\xe9\xe1\x98\x81\xec\x02\x03\xed\x10\x03\x87\xd0\x31\xc6\x5c\xbf\x50\x14\x48\x10\x81\x28\xcc\x83\x8b\xf1\xa0\xb0\xc3\x38\xce\x37\xf7\xa4\xd2\x82\xee\xcd\x7b\xb7\x30\x03\xf0\xa0\x30\x63\xcd\x5d\x61\x97\xf7\xfa\xa8\xf9\xe6\x3d\x1b\xef\x5b\x92\x13\xb5\x83\xcf\xdd\x0e\x8f\x98\x9a\x17\xf3\x1d\xde\x22\xf2\xcc\xf7\xcd\x0e\x7c\x91\xc6\x21\x3a\x81\x93\xeb\xe5\x23\xea\x72\xc5\x83\x3c\xf5\xe4\x00\x5f\x63\xc1\xe4\x84\x48\xbf\x1b\xe0\x77\xf9\xe3\x55\x83\xee\x1b\x57\x0d\x5a\xd8\x7b\x72\x14\x94\x7a\xaa\xa9\xaf\x35\x59\xa8\x34\x27\xfa\xa0\xb0\xc3\x9b\xdf\xc2\x0c\x40\x03\x2b\x05\xd3\xa8\x94\xd8\x7f\x06\x00\x00\xff\xff\x75\x56\x57\x52\x6e\x0d\x00\x00"),
		},
		"/index-Xa2vJ7PsJoaG-qPi.css": &vfsgen۰FileInfo{
			name:    "index-Xa2vJ7PsJoaG-qPi.css",
			modTime: time.Date(2019, 5, 10, 3, 32, 24, 199955553, time.UTC),
			content: []byte("\x68\x74\x6d\x6c\x2c\x20\x62\x6f\x64\x79\x20\x7b\x0a\x09\x6d\x69\x6e\x2d\x68\x65\x69\x67\x68\x74\x3a\x20\x31\x30\x30\x76\x68\x3b\x0a\x7d\x0a\x0a\x62\x6f\x64\x79\x20\x7b\x0a\x09\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x66\x6c\x65\x78\x3b\x0a\x09\x66\x6c\x65\x78\x2d\x64\x69\x72\x65\x63\x74\x69\x6f\x6e\x3a\x20\x63\x6f\x6c\x75\x6d\x6e\x3b\x0a\x09\x6a\x75\x73\x74\x69\x66\x79\x2d\x63\x6f\x6e\x74\x65\x6e\x74\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x7d"),
		},
		"/index-deO9TXEgTxEEbcEZ.js": &vfsgen۰CompressedFileInfo{
			name:             "index-deO9TXEgTxEEbcEZ.js",
			modTime:          time.Date(2019, 5, 13, 3, 54, 16, 326669837, time.UTC),
			uncompressedSize: 266,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\x4d\x4b\xc4\x30\x10\x86\xcf\xcd\xaf\x18\xf0\x30\x09\x6a\x7a\xb7\x54\x2f\x7a\x5b\x10\x14\xf1\x1c\x9b\xb7\x6e\x20\x9d\xd4\x74\xb2\xab\x88\xff\x5d\xf6\x22\x0b\xde\xde\x8f\x07\x9e\x58\xa6\xb6\x40\xd4\x7f\x34\xd4\xaf\x67\x64\x4c\x5a\xaa\x65\x1f\xcb\x51\x72\x09\xf1\x7a\x2e\x75\x61\xe7\x43\x8c\x0f\x07\x88\xee\xd2\xa6\x10\x54\xcb\x5b\x7b\x5b\x92\xf2\x15\x81\xc6\x5b\xfa\x36\x1d\xfc\x5a\x71\x62\xee\x31\x87\x96\xd5\xba\xc1\x98\x6e\x2a\xb2\x29\x35\x1a\x49\x70\xa4\x97\xa7\x9d\x85\xd7\x50\xdf\xf1\xcf\x79\xd1\x6a\x66\xe7\x0f\x21\x37\xb8\xc1\x74\xb9\x4c\x41\x53\x11\xbf\x06\xdd\x4b\x58\x40\x23\xb5\xbf\xe2\xb7\x9c\x26\xd8\xb3\x21\x49\xc4\xe7\xe3\x6c\xb9\x5f\x7b\x76\xce\x74\xdd\x25\x9d\xff\x90\xb8\xbd\x26\xdd\x5b\xee\xd9\xd1\x1d\x31\xd3\x0d\x9d\xf2\x60\x7e\xdc\xf0\x1b\x00\x00\xff\xff\x03\xa1\x31\x20\x0a\x01\x00\x00"),
		},
		"/post-QzVBPp18qo7ii9gK.css": &vfsgen۰CompressedFileInfo{
			name:             "post-QzVBPp18qo7ii9gK.css",
			modTime:          time.Date(2019, 5, 9, 15, 3, 3, 179756685, time.UTC),
			uncompressedSize: 592,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xcb\x6a\xeb\x30\x10\x86\xd7\xd6\x53\x0c\x1c\x0e\x9c\x03\x55\x71\x7a\xd9\xc8\x9b\xbe\x8a\x2c\x4d\xe4\x69\x75\x43\x9a\xb4\x31\x25\xef\x5e\xe4\xc4\x69\x93\x42\xbb\x10\xc3\x5c\x34\xff\xf7\x4b\x13\x07\x7f\x03\x63\xb2\x33\xbc\x8b\x2e\x50\x94\x13\x92\x9b\x58\xc1\xa6\xef\x5f\xa7\x41\x1c\x84\x58\xbb\x96\x6a\xf6\x7a\x56\xb0\xf5\xb8\x1f\x44\xd7\x82\xb4\x54\xd0\x30\xa5\xa8\xc0\x24\xbf\x0b\x71\x10\xdd\xf3\xae\x32\x6d\x67\x69\x52\x64\x8c\xac\xc0\x60\x64\x2c\xcb\xb2\xdb\x56\xd4\x14\xb1\x2c\x82\xba\x38\x8a\x92\x53\x56\x70\x5f\x30\x0c\xe7\xd2\x98\x98\x53\x58\xab\xa2\x63\xdc\xb3\xd4\x9e\x5c\xbc\x5c\x97\x53\x65\x49\x41\x3b\xac\x17\x8c\xae\x90\x1d\x44\xd7\x82\x64\x0c\xd9\x6b\x46\x79\x24\xac\x0a\x36\xdb\xd2\xce\x3a\xe0\x74\x56\x70\x77\x52\xca\xda\x5a\x8a\x4e\x41\xdf\x32\x4f\x95\x65\xe5\xd9\xa3\xe4\x39\xa3\x82\x98\x22\x2e\xd2\x4f\x01\x2d\x69\xf8\x17\xf4\x5e\xbe\x91\xe5\x49\xc1\xc3\x63\xc1\xf0\xbf\x71\x5c\x73\xfd\x00\x32\x88\xee\xf0\xcd\x8a\xa7\xcf\xe7\x59\x48\xae\x07\xbe\x24\x17\xbe\x47\x9f\xcc\x4b\x03\x5f\xff\x51\xef\x38\x0d\xa2\x3b\x11\x6e\xfa\xfe\x6f\xeb\x8e\xa9\x58\x2c\x0a\x78\xa2\x08\x35\x79\xb2\xf0\xc7\x18\x73\xad\x73\x84\x38\xfb\xfb\xfd\xf6\x47\x00\x00\x00\xff\xff\x6d\xa9\x94\xba\x50\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/.vfshash-assets.json"].(os.FileInfo),
		fs["/error-Xa2vJ7PsJoaG-qPi.css"].(os.FileInfo),
		fs["/favicon-JvizeVoabet5qTuc.ico"].(os.FileInfo),
		fs["/index-Xa2vJ7PsJoaG-qPi.css"].(os.FileInfo),
		fs["/index-deO9TXEgTxEEbcEZ.js"].(os.FileInfo),
		fs["/post-QzVBPp18qo7ii9gK.css"].(os.FileInfo),
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
