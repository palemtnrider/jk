// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package lib

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
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 12, 2, 11, 26, 20, 149482710, time.UTC),
		},
		"/std.js": &vfsgen۰CompressedFileInfo{
			name:             "std.js",
			modTime:          time.Date(2018, 12, 2, 19, 11, 11, 819151706, time.UTC),
			uncompressedSize: 26723,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7c\x5d\x73\xdb\x36\xd6\xf0\xfd\xfb\x2b\x68\x5f\x78\xc8\x92\xd6\x88\xb4\xa2\xba\xa2\xa0\x4c\xd2\xc6\xbb\xde\xc9\x26\x9d\xc4\xed\xce\xae\xc7\xa3\xa1\x45\xd0\xc6\x5b\x9a\x74\x41\x30\x5e\xaf\xac\xe7\xb7\x3f\x03\x80\x00\x01\x10\xa4\x24\xc7\x4d\xf7\xe2\xb9\x49\x2c\x12\xe7\x03\x07\xe7\x1c\x9c\x0f\x80\x5f\x12\xec\x10\x40\x9e\x9e\xd6\x9b\x98\x8c\xce\x4a\x7c\x97\x10\xb0\x7e\x53\x93\x72\x36\x0e\xfe\xf6\xf9\xe3\x87\x59\x18\xfc\xf3\xcd\xdf\xdf\xcf\xa2\xe0\x53\xf2\x30\x3b\xa1\xa3\xfe\x81\x11\x81\x6f\xf0\x4d\x05\xb2\xba\x58\x11\x54\x16\xae\xb7\x26\xb7\xa8\x1a\x5d\x5f\x83\xa2\xce\xf3\xb8\xf9\xb1\xbc\x2f\x2b\x30\xd6\x40\x46\xf7\xb8\x24\x25\x79\xbc\x87\xa3\xe5\x12\x15\x88\xb4\x38\x48\x00\x25\x1a\x06\x49\x04\x1e\x00\x63\x0c\x49\x8d\x0b\x87\x3e\xd0\xf1\xdd\x40\xf2\xa9\x2c\xc9\x9b\xca\xc2\x15\x0c\xb0\xb7\xe6\x90\x2e\x7e\x7a\x2a\xe0\x83\xa3\x80\x7a\x0d\x07\x2e\x1c\x61\x98\xa4\xe7\x05\x39\x89\x5c\x38\xba\x2f\x2b\xc4\xe7\xe4\xf9\xea\xaf\x00\x7a\x7d\x33\xb9\x4f\xc8\xad\x32\x0f\x6f\x4d\xa5\x0a\x41\xc3\xfd\x68\xb9\x2c\xb3\xac\x82\xc4\x55\x26\x17\x4c\x3c\x31\x27\xf8\xba\x1d\x58\x11\x8c\x8a\x1b\x75\xa0\x0f\x03\xe2\xcd\xa8\x54\xfb\xa8\x7f\x49\xf2\x1a\xee\x4b\x7e\xfa\x62\xe4\xe9\x3f\xaa\x22\xec\x40\xfc\xd4\x42\xbc\x59\x82\x53\x9d\xb8\x37\x13\x4a\x39\xa2\x3a\xd9\xc7\x03\x2a\x52\x58\x10\x93\x0b\x32\xcc\x45\x38\x96\x6c\x90\x61\x36\x88\x37\x33\xb4\xb8\x22\x09\x26\x16\x8d\x23\xde\x9a\xf0\x97\x1f\xaf\xff\x3f\x5c\x11\x77\x62\x28\x4d\x92\xa6\x3f\xeb\xda\xc2\xb4\x9e\x3e\x3f\x43\x30\x4f\x3f\x72\x3e\xc7\x01\x0c\xc6\x5d\xd0\x5f\x8d\xa5\xb6\xc2\x86\x76\xd8\x0b\x6d\x9d\x98\x69\x40\x09\xca\xa6\x1c\x05\x38\xd0\xc4\xdd\x45\x72\x6e\x08\xda\xe0\x80\xa1\x39\xb1\xd0\x87\x45\x6a\x97\x56\xa3\x2d\x74\x40\x23\xb1\x56\x37\x36\x31\x7f\x0b\x99\x77\x82\x23\x06\xbc\xfe\xf0\xf1\xc3\xbb\xd9\x38\x90\xe8\x66\x61\xf0\x09\x26\x29\xfb\x33\x0a\x7e\x4c\x8a\x15\xcc\xd9\x8f\x93\xe0\x0c\xe5\xf0\xbc\xc8\x4a\xf6\x73\x12\xbc\x47\x15\x61\x7f\xbe\xa2\xd8\xfe\x0e\xab\x2a\xb9\x81\xbb\xfa\x30\x09\xf0\x22\x1e\xac\xc5\x26\xfd\x57\x87\x1f\xd2\xf5\x5e\x12\x4c\xfa\x2e\xa2\xf8\x2e\xa2\xf9\x2e\xf5\x57\x40\x3c\xfb\x0c\x12\x7c\x53\x5d\x58\x0c\x78\x8b\xe9\x4c\x7a\x2c\xe7\x17\x64\x31\x1d\xbe\x70\x23\xba\x6e\xfd\x4c\xbc\x8c\xff\xaa\x0b\x2e\x38\xc3\x87\x70\xf7\xd5\x92\x66\x06\xda\x95\xb7\x69\xbb\x91\x26\xb4\x24\x65\x3a\xa6\x4b\x8b\xad\x91\xa1\xff\xe3\x00\x07\xca\x9c\x6d\x48\x76\x35\xe1\x16\x10\x16\xa9\x8d\xe1\xed\xe6\x83\xc1\x7a\x13\xe3\xd1\xe7\xf3\x7f\xbd\xfb\x78\xb6\xfc\xfc\xd7\x8f\x9f\x2e\x40\xd4\x3e\x38\xff\x70\x01\x26\x31\x1e\x9d\x9d\xbf\x7f\xb7\x3c\xff\xe9\xdd\x87\x8b\xf3\xb3\xf3\x77\x9f\x96\xef\xdf\x7d\xf8\xcb\xc5\x5f\xd9\xab\x77\xc5\xaa\x4c\x51\x71\x03\xd6\xbf\x5c\x9c\x9d\x2e\xdf\xfe\xf3\xe2\xdd\xe7\x59\x18\xfc\x72\x71\x16\x4e\x97\x9f\x2f\x3e\x9d\x7f\xf8\xcb\x2c\xa2\x34\x10\xd5\x42\x40\xf5\x94\xe9\xe3\x1b\x8c\x93\x47\x37\xf2\x62\x3c\xca\xf2\x32\x11\xef\xce\xf8\xdf\xfc\x6d\x03\x34\xba\xae\xb3\x0c\x62\x39\x74\x3a\x69\x87\x4e\x27\x7d\x43\x51\xf5\x1e\x11\x92\xc3\x77\x45\x8a\x92\x82\x41\x50\x05\x0c\xa7\x1c\x40\xfc\x3e\xe5\x3f\x2f\xc3\x60\x7c\xe5\x09\xf0\xcb\xf1\x15\x00\x20\x8c\xf1\xe8\x7d\x59\xdc\xd8\x4c\x38\x2f\x1f\x00\x79\x1a\x73\x0b\xbe\x45\x37\xb7\x00\x3e\x8d\x37\x0d\xc0\x68\x85\x61\x42\x4c\x57\x2c\x6c\x02\x80\xf1\xd1\x11\x04\x60\xfc\xba\x19\xfd\xaf\x77\x9f\x3e\xce\x28\x3f\xfc\x37\x1b\x2d\x51\x29\x5b\x69\xd9\xcc\x58\x35\xc7\xc6\x03\x08\x9e\x16\x8b\xc5\xd8\xf3\x25\x53\xdf\x4d\xa2\x1f\x26\x3f\x4c\xbf\x8f\x7e\x98\x5a\x10\xc2\xdf\xeb\x24\xd7\x6d\x4b\x71\x43\x6c\x8a\x80\xd0\xff\x8e\x8e\xda\x69\x02\xc2\xfe\x97\xe8\x28\xf3\x40\x61\x7e\x1c\xd0\x8d\x73\xf4\xb6\x46\x79\x0a\xb1\x86\x1c\x65\xee\x81\xd4\xca\x70\x1c\x4d\x36\x30\xaf\xa0\xd0\xd2\x8d\x70\x86\x78\xf4\xf6\x91\xc0\xb7\x6c\x25\x46\x49\x9e\x97\xab\x84\x40\x17\x7a\x5c\xd6\xd5\x7d\xb2\x82\x00\xf2\x1f\x77\xa8\x48\x72\x74\x53\x80\x90\xff\xfe\x42\x92\xeb\x1c\x2a\x1e\x9a\x3f\x58\xa2\x62\x59\x57\x10\x34\xcb\x85\xaa\x0f\xb0\x22\x30\x05\x59\x92\x57\x0d\xa6\x92\xd9\xc7\x92\x19\xb8\x18\xc7\x81\x2b\x70\x79\xd5\xfc\x86\x2b\x52\xe2\x65\x51\xdf\x2d\x61\x0e\xef\x2a\x31\x2e\x2b\xf1\x0a\x2e\x53\x98\x25\x75\x4e\x2a\x8e\x75\xd3\xca\x40\x91\x38\x1b\xf9\x93\x1c\xa8\xfa\x16\x0b\x22\x62\x47\x92\x26\x24\xe1\xd2\xe9\x2a\x82\xd3\x08\xd1\x0e\x99\x54\xad\xc6\xf7\xc3\x8e\xae\x1f\x09\xac\x5c\x6f\x54\xd5\xd7\x09\x33\x0e\xf1\x42\xdd\x32\x3a\x8f\xb8\xd6\x35\xce\xd8\xf3\xfa\x38\x60\xc4\xb7\xb1\x2f\x59\xe0\xf6\x38\xaa\x72\xb4\x82\x2f\xc8\xc6\x3d\x86\xf7\x86\x75\xa2\xcc\x25\x0b\x4d\xa9\x9a\x45\x91\x3a\x46\x36\x54\x53\x11\xf8\x1f\xc9\xc8\x2a\xb9\x4f\x56\x88\x3c\xba\xde\x71\xab\x9b\x3e\xf4\xfc\xf0\x88\x1c\x87\xf1\xc3\x2d\xca\x1b\xae\xd9\x9b\x39\xf2\x89\x0f\xb9\x01\x54\xa0\x8b\x24\x56\x2c\xa0\x61\xf9\x06\x97\x0f\xad\x35\x08\xc2\xaa\x29\xf8\x16\x44\xc7\x15\x37\xa6\xfb\x24\x75\x51\x9f\x08\x92\x54\x53\xc0\xac\xc4\x2e\xb7\xc4\x71\x0c\xe7\x24\x86\xbe\x2f\xa3\x95\xd1\x03\x0d\xa7\xd8\xf6\x75\xac\xcc\x94\x6e\x42\x76\xe4\x72\x7c\x57\xc7\x35\x6c\x2d\xae\x63\x10\xb2\x30\x64\x08\x5b\x38\x1d\x44\x17\x4e\x35\x7c\xd1\x56\x7c\x27\xd1\x20\x3e\x1a\x33\x29\xf8\x26\x5b\xf1\xa9\x9e\xd9\x82\x6f\x3a\xd1\xf0\x9d\x0e\xe3\x6b\x36\xc2\x7e\x8c\xcd\x80\x7d\x78\xec\xec\x1f\x56\x9c\x3b\xf3\xc9\x22\x7e\xdb\x1a\x53\xf3\x72\x43\xba\x11\xb0\x9f\xca\x7a\x0f\x62\xb2\xad\x2f\x43\x15\x75\x50\xd1\xb5\x1e\xc4\x65\x93\x1c\xc3\x35\xe9\xe0\xa2\x32\x1c\xc4\x65\x93\x18\xc3\x75\xda\xc1\x45\x65\xd7\x8f\xab\x77\x4d\x6d\x9c\xc9\xf5\xdd\x82\x6f\x57\xee\xe4\xda\x0e\xe0\x13\x51\xaa\xe6\x1b\x69\x10\x4b\xbd\x63\x77\x87\x7a\x7a\x82\x07\x00\x37\x34\x1b\x6d\x68\x37\xea\xbc\x24\x94\xd8\x56\x6a\xda\xb2\xef\x49\x2e\x9c\xee\x4f\x4f\x5b\x80\x3d\xe9\x9d\x44\xfb\xd3\xd3\x16\x68\x98\xde\x81\x08\xca\x5c\xec\x69\x74\xa7\x93\x3d\xe9\x76\x55\x6d\x9f\x99\x0a\xdd\x7b\x06\xcd\x3d\x66\x6b\xa1\xb9\xf7\x3c\x79\x4a\xf4\x4c\x92\x4d\x3e\xb5\x1f\xc5\xcf\x04\xd7\x2b\x2b\x45\x05\x79\xc1\x02\xcc\x1d\x31\x17\x4d\x34\xaa\x87\xca\xe4\x00\xe8\xa1\xcc\x9a\xdc\xe2\xf2\xc1\xa1\xb1\xf6\x3b\x8c\x4b\xec\x1e\x9e\xe5\x09\xe1\x91\x41\x35\x73\x2a\xc6\x97\x73\x57\x57\xc4\xb9\x86\x4e\x05\x31\x4a\x72\xf4\x1f\x98\x3a\xa8\xc8\x51\x01\x47\x87\xbd\xf4\x4b\xf2\xc1\x60\xa1\x95\xa0\x08\x96\xb7\x91\xe7\x01\xb4\x24\x9b\x50\x34\x9c\x99\xa2\x64\x0c\xf1\x59\xf6\x73\x41\x65\xd4\xf5\x64\x3c\x06\xbf\x24\x57\xba\x30\xec\x28\x4a\x43\x19\x3a\x61\xa5\x35\x60\x53\x71\xe9\xf1\x96\x35\x99\x56\xe2\x35\xba\xe6\x47\x27\x51\x14\x46\xd1\xab\xc9\xf7\xd1\x36\x11\xad\x92\x82\xca\x82\xd2\x70\x78\x60\xeb\x5c\xc3\xc7\xb2\x48\x9d\xc8\xb9\x41\x37\x09\x0b\x7a\xa9\x80\x78\x9c\x09\xe7\xf3\x30\xe6\xb1\xa2\x3d\x25\x42\x5e\x5c\x8d\x2a\x48\x7e\x16\xc1\x2f\x3a\x86\xf4\x91\x8c\xdf\x21\x71\x89\xf8\x15\xb0\x97\x8d\x3c\xaa\x5e\x15\xef\xd8\x93\xba\xa3\xa8\x75\x01\xeb\x36\xaa\x2e\xd1\x31\xf1\xd5\xf1\x3d\x4b\xa6\x54\x54\xba\x44\xa5\x62\x72\x59\xab\x29\x1e\xcb\xf1\x34\x1d\x01\x97\x57\x1b\x4b\xce\x47\xe2\xfe\xc8\xb6\x51\x2e\x78\x05\xc6\x1b\x3d\x33\x24\xb8\xb6\x25\x86\x3b\xe8\xa0\xac\xb5\xd8\x8c\x49\x65\xfe\xe9\xe9\x60\x2f\x03\x93\x88\x9d\x55\x92\xe7\x30\x75\x1e\x10\xb9\x2d\x6b\xe2\x28\x22\x3c\xf4\x36\xfa\x66\x35\xf6\x62\xa5\x60\x27\x38\x8f\x95\x12\x9a\x26\xad\xe3\x90\x49\x2b\x86\x0b\x30\x6e\x92\xfe\x56\x46\x54\x7c\xc7\xc7\xde\x5a\xa8\xa7\xaf\x0c\xe6\x6f\xf4\x8d\x59\x07\x3f\x00\xe3\xd7\xe4\x58\x7f\x36\x1b\x73\x5d\xaf\x40\x14\x1b\xb0\xc7\x1d\xe1\x73\xae\x0b\xe0\x22\xbf\xf2\xbe\xd3\x6b\x56\x06\x74\xc1\xc7\x96\x60\xcc\xfe\x4f\x40\x6b\xec\x31\x99\x51\xa6\x1b\x5d\x50\xb2\xfc\x51\x0e\x8b\x1b\x72\xcb\xb5\x83\x42\x65\xb6\xc4\x4a\x85\xb8\x84\x57\x54\x2b\x0b\x00\xd4\x22\x27\xe7\x20\xf3\xda\x8c\xaa\x06\x06\xb3\xf5\xbc\x88\x6b\xdf\x78\xda\x2a\x89\x86\x28\xf1\x6b\xef\xc0\x46\xc0\xaf\x3d\x6f\xbd\x2a\x0b\x82\x8a\x1a\x3a\x64\xb3\x29\x81\xc9\xdc\x35\x86\xc9\x6f\x9b\x0d\xca\xdc\xb2\x59\x1b\x5e\x3c\xb1\xcd\x2b\x1e\xcc\x7e\x82\xf2\x98\x78\xbc\x5a\xa3\xc9\xec\xbe\xae\x6e\x75\xab\xf7\x7a\x11\x19\x04\x03\xc3\x59\x78\x1b\x5b\x6d\x46\xf8\xef\x9e\x72\x0a\x2a\x50\x65\x76\x4b\xa8\x53\xe6\x2b\x88\x44\x99\x88\x39\x2f\x2d\x99\x0f\x54\xd7\xe4\xf7\x15\x38\x99\xd7\x41\x8d\x66\x1c\x80\xde\x61\x5b\x2c\x37\x43\x39\x74\x50\x0a\x0b\x82\x32\x04\xb1\xdc\xa2\x39\x5e\xe7\xb0\x9f\xfe\x46\xe8\x50\xd5\x4b\xfc\x38\x8c\x2b\x6a\x80\x95\x34\xc0\x36\xdd\x42\xa3\xd5\x6d\x82\x7f\x2c\x53\xf8\x86\xb8\x95\xe7\x6d\x36\xbb\x88\xc3\x8b\x8d\xf8\x88\xb4\x8b\xaa\xee\x36\xad\x7a\xf4\xb8\x42\x0c\x7f\xaf\x11\x86\x3c\x6c\x32\x56\x89\x17\x9f\xad\x9a\xc8\x97\x0e\x1f\x1b\x4a\x7f\x12\xb9\xd8\x8b\xf5\xda\x49\x6b\x0f\xc8\x87\xde\x01\x18\xd3\x05\x3b\xa8\xb6\x2f\x08\xcc\x53\xe7\xd0\x87\xfe\xa1\x12\x2f\x91\xfe\xd0\x84\xfa\x9f\x5f\x59\xf5\x4f\x8f\xfb\x90\x65\xa3\xe2\x06\x62\x96\x0a\x55\x4d\xd4\xb6\x51\xf2\x9d\x08\x12\xd9\x3b\xc4\x1e\xf4\x6e\x2d\x26\x13\xc6\x92\x0b\x5b\x33\xc9\x7b\x6a\x7f\x69\xcb\x06\xc6\x6b\xd7\x9f\x59\xaf\xb7\x13\x93\x3a\xa8\xa8\x48\x52\xac\x60\x99\x29\xd5\x73\x19\x1c\xa9\x05\xdd\xcb\xab\xa6\xc3\x30\x6e\x0a\x61\x78\x4e\x1a\x53\x6a\xcc\x53\x2c\xa6\xaa\xa5\xd8\xf7\x99\xd5\x55\xf3\x57\xaf\xa2\x1f\xa6\x4f\x4f\xd5\x02\xbc\x9a\x9e\x44\x63\x6f\x8d\x40\xd5\xe2\x2f\x6c\x60\xc0\xad\xe6\xf3\x70\xec\xf9\x85\xef\x4e\x5f\xbd\x3a\x99\x1e\xbb\x0c\x0b\x7b\x78\xcc\xd1\x50\x77\x88\xe6\x61\x74\xea\xad\x21\xf7\x5f\xa8\xf1\x6c\xec\x45\x34\x9e\x28\x6f\x16\x8b\xe9\xd1\x49\xf8\x14\xfe\x10\xa9\x63\x18\x6a\x75\x50\x18\x1d\x85\xaf\x9e\xa2\x68\xd2\x8c\x52\xde\x9c\x1e\x7d\xff\x14\x4d\xc6\x01\x1f\x35\x3d\x79\xa2\x94\x37\x2a\x7e\xf3\x99\x7c\xb0\xd9\x68\x79\xb5\x88\xb7\x14\x65\x74\xc3\x00\x36\x22\x0d\xc2\x6d\x46\x7a\x0c\xc4\x58\x4f\x86\x44\x18\x8c\x83\x52\xd9\x1d\x83\x44\x9a\x56\x13\x36\xc6\x78\x2e\xc0\x62\x2a\xe4\x75\x72\x59\xfa\xfe\x15\x80\x97\xf8\x6a\xa3\x2a\x95\x54\xce\x61\xbd\xb2\xf4\x53\x1a\x2c\x5a\xeb\x44\xf6\x40\x7a\x82\x71\xce\x25\x65\x71\x29\xda\xa8\xa2\x04\xbc\x04\x63\x1d\x50\x06\xcc\xb6\x4e\x07\xef\x5c\x28\x25\x56\xa3\x2f\x44\x3c\x83\x0d\x65\x46\x8c\x7e\x6f\xb6\xc1\x98\xeb\x85\x15\xdc\xf6\x81\xcb\xd9\xf4\x62\x50\x56\xd9\x12\xae\x4b\x61\x90\x5e\x04\xc2\xe9\x0e\xcf\xa0\x59\xfc\x5e\x2c\xe2\x48\x44\x6f\x1b\x49\xe9\xfc\x7a\xf3\x79\x34\x59\x2c\xa2\xc9\x20\x36\x36\xb8\x17\x1d\xe7\xea\x92\x5c\x6d\xe3\xc8\xa8\x26\xda\x58\x62\x55\xc4\xf9\x3c\x9c\x2e\x16\xe1\x74\x2b\x4f\x03\x08\x25\x53\x4f\xda\x4f\x3f\xbc\x9a\xcf\x4f\xb7\x31\x6a\x14\x04\x77\xc6\x6b\x3c\x8b\xae\xe8\x4c\x8c\x87\x27\x57\x54\xe4\x5b\xa7\x36\xc0\x82\x72\x80\xc0\x5b\x2c\x16\xa6\x6d\x75\x66\x63\x94\x23\x35\x33\xe3\xdd\x4d\x13\x6b\x60\x3e\xf1\x27\x03\x46\x27\x58\xde\x9d\x0e\x9f\xa0\x46\x48\x3c\xda\x4a\xc9\x56\xb1\x6d\x5a\xcd\x97\xe3\xa6\x26\xa1\xce\x25\x96\xce\xac\x69\x6a\x5f\x8e\x87\x15\xd5\x56\xc2\x15\x04\xcc\x1e\xf6\xeb\xf1\x2c\xb4\xd2\xec\x1b\x1f\xce\x2c\x3c\xfa\x13\x93\xcb\xe9\x64\x88\x4b\x5b\x2f\xa7\x3d\xc9\x22\x14\x14\xc0\x61\x04\xa6\x51\xef\x8f\xa1\x5b\x29\xee\x62\x88\x4d\x33\x01\x70\xb1\xe8\x37\x40\xc9\xd9\x1f\x82\xb8\x5b\x6a\xde\x15\x6f\x6c\x5a\x36\x7d\x1a\x4e\x63\xd3\xb6\xe9\xe3\x01\xeb\x96\xd3\xfb\x2f\x60\xa3\x5b\x07\xb7\x84\xae\x34\x9a\x29\x1f\x2c\x75\x25\x7f\x12\x40\x76\xec\xa0\xdf\x5a\xe5\x64\x07\xc8\x08\xb3\xef\xd2\x69\x1d\xc2\x6e\x84\x6c\xe5\x75\x6a\xb8\xad\xd5\x0b\xa1\x6a\xf3\x6b\x5d\xc7\x2e\xf8\x3b\x33\x51\xed\x75\x10\xbf\xd5\x73\xd8\xe5\x3a\xe8\x3c\xfa\xb9\xbc\x81\x4d\x3a\x75\x2e\xf3\x5a\x5b\xc1\x4b\x8b\x21\xe6\x7a\x68\xe2\xef\x96\x83\x6f\xc9\xe5\x5a\xee\x1c\x54\x39\xa4\x2c\x9d\xea\xb6\xc4\xc4\x21\xa5\xb3\x2a\x0b\x92\xa0\xc2\x49\x0a\x25\xfb\x16\xf5\x55\x02\x0e\x0f\x8d\xd2\x60\x1f\x0b\x4d\xc5\xd0\x07\x3c\x29\x1a\x65\xb8\xbc\xfb\xb1\xc9\x3d\xb4\xed\xac\x69\x68\xf7\xcc\x10\x7a\xde\x46\xab\x67\xd8\xe4\x2a\x8e\xb3\xd9\xd3\xe5\xe3\xde\x0d\xa7\x29\x66\xb5\x09\x31\xf6\x5e\x9b\x4f\x7c\xc8\x4e\xaa\xf6\x52\x66\xe7\xe3\x3a\x87\xcf\x9a\xc3\x8a\xd0\x37\x68\xd3\xb4\x75\x74\x7d\xcd\xf6\x97\x78\x87\x79\x55\x66\x46\xc9\xd0\xfb\x96\xfd\x8c\x17\x00\xec\xcf\x2b\xba\x6c\x3c\x01\x1c\xc7\x44\x29\xa2\x9d\x7f\xb8\x60\xf5\x78\x00\x40\x7b\x28\x6d\xd4\x9e\x49\xb3\xc6\xb6\xed\x09\x97\x80\xf8\xc8\xdb\xf0\x44\xb5\x98\x23\x2e\xf1\x52\x2d\x1f\x2a\x71\xac\x5f\x34\x59\x6a\x32\xa7\xa9\xe1\xba\x04\x49\x9b\x99\x66\x43\xc3\x69\x8e\xb8\x2e\x81\x9b\x1c\x9d\x84\xde\x7c\x3e\x7d\xca\x8e\xa6\x27\x2d\x6c\x3d\x08\x3b\x19\x37\xb0\xe1\x2b\x1a\xb5\x46\x4f\x2e\x85\x66\x68\x6a\x0d\x4d\xda\x83\x86\x01\x7f\x4f\x61\x4f\x25\x2c\x45\x53\x0b\x34\x29\x45\xc3\x2b\x86\x22\xcb\xad\xec\x6a\x5f\x36\xa9\x6e\x79\x0c\xd8\xc0\xb8\x67\x9c\x5b\x2e\x16\x34\x23\x67\x79\x78\xe0\x96\x47\x6e\xc8\xb3\xf1\xd0\xf3\x9b\x84\x7c\xa3\x35\x25\xec\xca\x83\x8a\x14\x61\xb3\x47\x20\x16\xd4\xd4\x4c\xd2\xef\xb6\x96\xcb\x2f\x66\x19\x67\x08\x8f\x66\xc0\x5b\x91\x2e\x73\x58\xec\x12\x47\x77\xe9\x0c\x31\x7c\x9b\x54\x4b\x64\x71\xb2\x4d\x45\xe6\xbf\xa1\x3a\xb9\x8b\xfb\x64\xbc\x2a\x15\x1b\x28\xca\xda\xbb\xfa\x4e\x21\x4d\x7e\xbc\x4e\x7a\x52\x5c\xf7\xc7\x8d\xcf\x29\x39\xac\xca\xa2\x22\x0e\x02\xbc\xa1\x32\xca\x30\x84\xff\x81\xae\x38\xf8\xee\xc5\x02\x91\x53\xb9\x55\x50\x80\xc3\xc3\x60\x9d\xb1\x57\xb3\x12\x20\x76\x2e\x3e\xe0\xf7\x0d\x66\x09\x88\x36\x60\xbd\x61\x55\xfa\x8a\x38\x59\x73\x4a\xb2\xa9\x8c\xb8\xe1\x38\x9a\x78\x0d\xb9\x1a\x94\x00\xa0\xd1\xa7\xe4\xe1\x75\x35\x22\x25\x37\x24\xd7\x9b\xfd\xed\xf3\xc7\x0f\x23\xee\x38\x51\xf6\xe8\x56\x62\x7c\x0a\x32\xad\x4e\xe7\xd6\xe2\xcd\xb5\xf9\xa6\xf0\x06\xae\x26\xb8\x99\x67\xbd\x44\xe0\x66\x41\xda\x79\xf3\x73\x42\x6e\xdd\x2c\xb8\xee\xbc\xb8\x78\xbc\xa7\x10\x65\xe7\x05\xbf\x0f\xe0\x66\x41\x22\xd8\xbb\x05\x7d\xe7\xfe\x29\x2b\xf6\x33\xd8\xfa\x1b\xe5\xa0\xb5\x9b\x89\x93\xd4\xed\xad\x9d\xee\x48\x37\x0b\x6e\x05\xf9\x7b\x60\x3b\x33\x4d\x09\x64\x4d\x33\xc1\xbd\xf7\xe2\x5f\x4f\xff\x51\xe2\xdf\x20\x8e\x46\x15\x2c\x52\x37\xd3\x4f\x2d\xba\x9e\xb7\x91\x4a\xd0\xe8\x53\xe5\x92\xe0\xf0\x90\xea\x0f\xef\x3b\x95\xec\x4e\x42\x39\xfa\x09\x66\x10\x63\x98\x7e\x82\xe4\x4b\x92\xcb\xdb\x09\xe2\xf1\x2c\x0c\x98\x45\xce\x22\x3a\xf8\xac\xce\x33\x94\xdf\xc1\x82\xf0\x8b\x1c\x72\x74\x42\x12\x3a\xb2\x48\x3f\x66\x9f\x09\x86\xc9\xdd\x2c\x6a\xe0\x4e\x28\x5c\x7b\xa9\x61\xd7\xdb\x0a\x2a\xcc\x8b\x5c\x58\xd0\x10\xca\x3b\x0b\x36\xc6\x14\xf3\x73\x21\xbf\xb6\xa0\x02\x3f\xe7\xe6\x42\xcf\x6c\xf8\x49\x84\x17\xbc\xba\x20\xce\xc5\xb5\x77\x17\x64\xfb\x42\xfa\x19\x76\xfc\xd9\x60\x89\xe9\xb2\x55\x16\xe6\x95\x82\xd0\x04\x4d\xd2\xf4\xb3\x31\x8d\xee\x9d\x9a\xe9\x84\x5d\x08\x22\x26\x1b\x26\x32\x58\xa4\x3d\x5c\x0c\xdf\x13\x68\xd5\x78\x77\xfd\x12\x10\x2f\xa4\x5d\x12\x9d\xd4\xad\x2e\x4b\x56\xcd\x12\xc3\x9e\xa7\x57\x96\x59\xfc\xb9\x5a\x25\x19\x62\x5a\x63\x91\x81\x5d\xa3\x24\xd8\xd7\xeb\x93\x44\x05\x8b\xd4\x4a\x7f\x57\x5d\xfa\x04\xab\xfb\xb2\xa8\x76\xbe\x61\xd5\x85\x7c\x61\xdd\x92\x68\x3b\x3a\xd6\x65\x75\x50\xd7\xc4\xf0\xaf\xd3\x39\xcb\x2c\x31\xdb\x45\xfe\xe8\x0b\x59\xe6\xae\xd5\x5c\xcd\xda\x81\xb3\x6f\x70\x4b\xcb\xc2\x85\x66\x0a\x96\xa5\xb2\xdd\xdb\xb2\xa0\x49\xd2\x66\xbe\xe6\x15\xae\xee\x15\x46\x6a\x1d\x56\x29\x6d\xc1\xbc\xeb\xbd\x2e\x0b\x0e\xc5\xda\xac\x53\xdc\x6a\x75\x09\x49\xf6\xb0\xb4\x84\x24\x2f\x65\x5d\x14\x55\x6b\x51\x1a\x1b\x76\x2b\x4a\x48\xf2\x4c\xcb\xd1\xb9\x36\xba\x73\x7f\xdc\xbd\xeb\x86\x32\xd7\x43\x7d\x82\x3d\xee\x98\x0e\x4f\xd2\xf4\xad\xc1\xe1\xe0\x5d\xdf\x06\x8c\x2a\x82\x49\x63\xdb\xe2\x2b\x91\xe3\xee\x3a\xa0\x00\xbd\x90\x2a\xa8\x18\xa5\x46\x58\x79\xb3\x2a\x86\x32\xf2\x79\xfa\xa1\xd2\x67\x8b\x62\xa7\x6d\xae\xd9\xb8\x03\x0c\x8b\xb4\x0f\x74\xdb\x52\xb4\x31\xfe\xee\x2b\xd1\xc2\xbc\xd0\x42\x28\x08\xe5\x3a\xd8\x18\xb3\x2e\x43\x3b\xf0\x79\xab\x60\x9d\xcd\x9f\x1b\x55\x59\x59\x62\xdf\x4d\x78\xc6\x5e\x3b\xdd\x6f\xaf\x35\x92\x3e\xb9\xd9\xf6\xf3\xb4\xaf\x4b\xb3\x7d\x4e\x61\xeb\x2e\xab\x90\x67\xb6\x60\x55\x0f\xd3\x52\x4e\x4c\x69\x7e\x7d\xb8\xa9\x23\xfb\xb5\xbb\x26\xb6\xfd\x39\x64\xfb\xb3\x4d\xb2\x7d\x38\xb7\x39\xe1\x48\x3a\x61\x05\x18\x16\x69\x8f\x58\x86\xbd\x00\x2f\x45\x24\x60\xbd\x69\x8b\x49\x19\x05\xe3\x2f\x20\xd8\xe1\x4c\x8a\x2c\x53\x6d\xb3\x66\x17\x8a\xb1\x15\x40\x8d\x9d\xb9\x5e\x7b\x97\xd8\xf5\xe2\x1c\x12\xa7\x60\xff\x66\x71\xf5\x80\xc8\xea\xd6\x45\xad\xfa\xbb\x9e\xb7\x5e\x25\x15\x74\xba\x12\x65\x85\x89\xb5\xbb\x4e\xe9\xff\xc5\x06\x24\x97\xd5\x95\xa0\x46\x40\xbb\x9d\xc7\x0d\x3a\x97\x78\x71\x06\xe4\x89\xf4\xe6\x7c\xaa\x82\xbc\x46\x2a\x7a\x5e\xe2\x58\xbb\x6b\xc8\xfe\xe8\x21\xc0\x46\xe9\x14\xda\x4a\x27\x19\x31\x58\xd7\xb3\x10\xd3\x48\x29\xb5\x15\x77\x0d\x8b\xb4\xa5\xc6\xe0\xe2\xe6\x06\xc9\xac\x53\x4a\xfd\xa5\xf8\xad\x28\x1f\x0a\xe7\x8e\x97\x92\x1c\x0c\x57\x08\x7e\x81\xa9\x93\xe1\xf2\xce\xc1\x75\x41\xd0\x1d\x3c\x64\xa7\xce\x0a\x00\x40\x5d\xa4\x30\x43\x05\x4c\x85\x63\xdd\x14\x6e\xa6\xd4\x93\x6a\x7e\x9d\x24\x40\xde\x3a\xb9\x24\x57\x80\xcb\x16\x06\x5c\x04\x38\xa0\xac\xa1\x4d\x3b\x3e\x6d\xcb\xcc\xae\x07\x16\x9d\x52\x2f\xf1\x94\xc1\xd7\xaa\x92\x51\x9d\x14\x4a\xb4\x4d\xe1\xa0\xd7\x2a\xd1\x3e\x09\x13\xbb\xa1\xc0\x35\xaa\x52\xb2\x17\x45\xa5\x8c\x20\xba\x59\x72\xdb\x0a\x0b\x04\x4a\xcf\xeb\x67\x5c\xde\xa1\x8a\x66\x1f\xcc\xc6\xd4\x65\x6f\x96\xc3\xf5\x3c\x6f\x63\x27\x25\x6b\x70\x06\x35\xf1\x5c\x23\x28\x85\x66\xb7\x20\xe5\xf0\x49\xc3\x93\xcb\xbe\xfd\x00\x16\x6b\x29\x7c\x44\x85\x9f\xc2\x1c\x12\xe8\x24\x97\xf0\x4a\xdd\x94\xe9\x32\xd5\x2e\x0c\xd0\xe8\x1a\x15\x74\x49\xc5\x5f\x58\xfe\x95\x52\x55\x83\xff\xbe\x87\x2b\x02\x53\x47\xd1\x57\x27\x2b\xb1\x73\xcf\xa8\xa2\x0c\xc1\xd4\x49\x9b\x09\x1c\xd2\xb9\x7b\x1b\xa1\xb9\xdb\x64\x76\x78\x96\xa0\x1c\xa6\x0e\x29\x9d\x14\xae\xca\x94\xaa\x32\x5f\xc5\x46\x95\xe1\xef\x35\xac\xc8\xa1\xe7\x6d\x36\x6d\x81\x14\xc3\xd5\x17\x37\xe3\xad\xb9\x5b\x70\xcb\xca\x9e\xb7\x23\xf1\xc5\x95\x5d\x63\x9c\x16\xe2\x45\x22\x1c\x05\x9d\x54\xcd\x2e\x4b\x96\xe8\xa6\x05\x7c\x4e\x6c\x63\x9d\x45\x8d\xf7\x4e\x8b\x9f\x93\x83\x58\x69\x53\xd7\x53\xd6\xfb\x7e\xf8\x68\x28\x80\x69\xcf\xf2\x2b\x5f\x3e\x52\x68\xb3\x68\xc0\x22\x6a\x5b\x0e\xae\x80\x25\x69\xfa\x0b\xde\x9a\x1e\x8b\x4c\x48\x07\xbc\x30\x27\xd9\x8d\x08\x4e\x22\x99\x5a\x2b\xb0\xb0\x48\xad\x9c\x0e\x6f\xde\xd2\x9e\xef\x5b\x67\x8a\x8c\xce\xce\xab\x30\x52\x37\x5d\xad\x1d\x43\xbc\x5e\x79\x51\x6f\xd9\x11\x8a\x8b\x02\xd9\xf5\x29\x80\x9d\x7d\x0a\xd8\xd3\x3a\x41\x7d\xad\x13\x24\x5a\x27\x52\xe3\x2d\x9d\x13\x14\x14\x82\x76\x69\xef\x9c\x20\x2f\x46\xa2\x73\x52\x4a\x39\x5d\xbb\x74\x33\x32\xda\x28\xa8\xd3\x46\xe1\xdd\x92\x1c\xe4\xcc\x6d\xe4\xa3\x33\x94\xc3\xcf\x8f\x15\x81\x77\x46\xbf\x44\x7c\xa9\x69\x16\x06\x3f\xb1\x36\x70\x89\x1f\xd5\x1e\x08\x07\xa5\x23\x76\xf5\x38\x2d\xc4\x8b\x78\x1c\x05\x5d\x1b\x83\x75\x58\xb2\x78\x9c\x16\xf0\x39\x1e\xc7\x3a\x8b\xeb\xa4\x82\x45\x72\xb7\x77\x9e\xf0\x1c\xb7\x63\x65\xe0\x39\xdf\xbb\x7b\xce\x07\xe7\xac\xc4\x51\x95\x22\xbc\xa7\xc7\x3b\x55\x3c\xde\xc1\x81\xea\xf3\x6c\xdf\x7a\x6b\x3e\xef\xa2\x50\xe7\x19\x52\x77\xb9\x6d\xf9\x91\x02\x96\xa4\xe9\xdb\xee\x5a\x0d\xba\x3e\x1d\x7a\x97\x4f\xc5\x85\x56\xc8\x73\x5d\x4a\xb6\x24\x2a\x0a\x7c\x18\xf8\x6c\xb2\x3a\x38\x4d\x7b\x6c\x73\x1d\xf6\x9b\xf9\x48\x5a\xee\xee\x36\x2a\x41\x5e\xc8\x48\x5b\x7c\x6d\xc8\xda\xe5\xca\x6a\xa6\x72\xdc\xf3\xec\xd4\x36\x93\x6f\x6b\xa8\x36\x0e\xbe\x9d\xa5\xda\xa8\x67\x28\xef\x94\x3f\xf5\x6b\x61\x5b\xac\x15\xbf\x1e\x72\xa3\x12\x87\x38\x3b\xa4\x3c\xe2\xe7\x75\x34\x76\xb1\xe7\x93\xef\x26\xf2\x8b\x42\x3b\xb0\xfe\x9e\x1d\x90\xf9\x0a\x5f\x63\xf2\xb3\xcc\x61\x61\x89\xaf\x54\x16\x78\x91\xd9\xa2\xb3\x76\x5f\xd3\x02\x3e\xcb\xd9\x68\xe0\xfb\x79\x1b\x0d\xf4\xcc\xb2\xd2\xbd\x35\x16\x15\x96\x87\x4e\x0c\xdc\x72\xfd\x4e\xce\xb9\xb9\x78\x34\x69\x6f\x43\x4d\xd4\x3b\x4e\xe2\xe9\x71\x18\xe3\x05\x18\xc7\x98\x5d\x95\x54\x3f\x00\x71\x89\xaf\xda\xb3\x99\xfa\x55\xa6\x8e\xf4\xf7\xe0\x26\x98\x18\x08\x60\x91\xda\x17\x6f\x9b\xf3\x54\xbf\x55\xb9\x7f\x8c\xf3\x62\x99\x95\x81\xb2\x13\xeb\x6c\xcd\xb0\x74\x04\x5f\x17\xf3\xfc\x49\x9f\xd9\x35\x18\xd0\x36\xff\x1d\xce\x77\x18\xe0\x7b\x7e\xf1\xd5\x80\x56\xf6\xe2\x3d\xf3\x98\x7c\x24\xbe\x77\xba\xbb\x3a\x09\x88\x17\x52\x25\x89\x4e\xaa\x51\x97\x25\xab\x0a\x89\x61\xcf\x53\x1f\xcb\x2c\xbe\x9d\xea\x48\xe2\x4c\x2d\x2c\xf3\xb5\xab\x8c\x04\xdb\x5b\x5d\x24\x24\x2c\x52\x2b\xb9\x5d\x3c\x8f\x48\xca\xf6\x3b\xb1\x61\x83\x7d\x41\x2f\x64\x20\xd6\x7c\x51\x1f\xc3\xbd\x1e\x49\x07\x78\xbe\x5f\xea\x9d\xed\xb7\x39\xbb\xd1\xcd\xa1\x9b\x86\xd2\x4e\xdc\x7d\x83\xf3\x1b\x56\x3e\xa4\x0b\xed\x5d\x36\x5b\xfd\xc8\x8a\x6a\xcf\x53\x1c\x3d\xf2\xda\x8a\x7d\xbf\xac\xcb\xc0\xd2\x38\xed\x81\xc9\x6e\xff\x6e\xef\x0a\xac\x58\xd1\x64\xc5\xab\xe3\xbb\x9a\x64\x33\xfc\x45\xac\x50\xe0\x6a\x5b\xf9\x3a\x27\x16\x5b\x6b\x40\x9e\x63\x5e\x5d\xce\xef\xfa\x3f\x78\xfc\xd2\x7e\x5b\x10\xe7\xa7\x06\x8c\x69\xda\x3c\xb6\x00\x48\x52\xcb\x77\x99\x07\x3d\xb6\x80\x84\x45\xda\x25\xb4\xa5\xaf\x98\x27\x55\xe5\x2c\x79\x51\x12\xd7\x3c\xc3\x91\xcb\xc9\xf6\x38\x22\x3e\x80\x4b\xd3\x7f\xb8\xd9\x70\x90\xc7\x1d\x40\x78\xb2\x06\x95\x6e\xd2\xcd\xb3\x0b\xa0\xdb\x22\x28\x17\x75\xc6\x88\xe3\xe7\x5a\x35\x74\x38\x18\xfa\xaa\xaa\xa8\x16\xa5\xbe\x44\x65\x54\x34\x7d\xb7\x16\x45\x63\xe3\xb6\x40\x6f\x53\x2e\xf1\xda\xdb\x03\xfb\xee\x88\x6e\x26\x1b\x73\xb5\xad\x31\x67\xf1\x8b\xb2\x06\xab\xb6\xcb\xda\xc7\x71\xdd\xed\xcf\xd1\x11\x4b\x6a\xd8\x74\xe5\xbc\x80\x70\xbd\x63\x73\xec\x21\xc2\x2b\xba\x6a\x1f\xbc\xb1\x87\x16\x3d\xfb\x5a\x8a\xde\xe4\x84\x4a\xbb\x6f\xa8\x53\x2b\xdb\x67\xb2\xb3\x45\x4a\x76\xfd\x05\x15\x59\x79\xa8\x76\x4a\xcf\xbf\x42\xb7\xed\x21\x1e\xd7\x69\x33\x8e\x33\xf5\xd9\x16\xad\x7d\x95\x1e\xcb\x50\xf9\xff\x74\x38\x51\x12\x71\x43\x89\xe5\x73\x4d\x8b\x55\x35\x6c\x0e\x5e\xa8\x85\x1f\x3a\x4f\x51\x63\x48\xc1\x38\x4e\xdb\x0f\xa5\xa4\xbe\x2f\xd4\x07\x83\x06\xc8\x4d\xbd\x18\x5e\xa6\x57\x80\x5b\x05\x16\x56\x81\xa5\x55\x6c\x14\xb3\x79\x6c\xcd\x06\xee\x66\x2e\xd8\x6e\x2e\xb8\x6b\x2e\xf8\xa5\xcc\x85\xce\xfc\x0c\xac\xf3\xf2\x66\x56\x04\xfc\x7e\xd4\x0c\x05\xec\x7e\xf1\xac\x0a\xe8\xd6\x3e\xbb\x0f\xe8\xe0\xd9\x4d\x90\x22\x3c\x3b\xdf\xc4\xf0\xdf\xf7\x25\x26\x4e\x43\xd5\x39\x8b\xff\xdf\xff\x06\x00\x00\xff\xff\x0b\xe2\x5f\x6e\x63\x68\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/std.js"].(os.FileInfo),
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
