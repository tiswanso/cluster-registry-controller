// Code generated by vfsgen; DO NOT EDIT.

package cluster_registry

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

// Chart statically implements the virtual filesystem provided to vfsgen.
var Chart = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/.helmignore": &vfsgen۰CompressedFileInfo{
			name:             ".helmignore",
			modTime:          time.Time{},
			uncompressedSize: 349,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\x8e\x41\x6e\xe3\x30\x0c\x45\xf7\x3c\xc5\x1f\x78\x33\x63\x0c\xe4\x43\x24\xb3\x98\x55\x0b\xa4\xc8\xb6\x90\x6d\x46\x62\x22\x8b\x82\x44\x27\x6d\x17\x3d\x7b\x91\x04\x41\xbb\x79\x20\x3f\xc8\x8f\xd7\xe1\xd9\x9b\x71\xcd\x0d\xa6\x90\x90\xb5\x32\x2e\x91\x33\xc6\x55\xd2\x2c\x39\xa0\xf8\xe9\xe4\x03\x37\x47\x1d\x5e\xa2\x34\xb4\xb5\x14\xad\xd6\xd0\x22\xa7\x84\x90\x74\xc4\xe2\x6d\x8a\x92\xc3\x5f\x54\x4e\xde\xe4\xcc\x28\xde\xe2\x8f\xdc\xe7\x99\x3a\x64\x0e\xde\x44\x33\x7e\x97\xca\x07\x79\xe3\x19\x17\xb1\x88\x5f\x7f\x1c\x9e\x72\x7a\x87\xe6\xdb\xe7\x55\x09\x85\x2b\x92\x64\x76\xe4\xb6\xbb\xd7\x9d\x69\x65\xea\xb0\xd1\x65\xd1\x8c\xfd\x66\x87\x59\x6a\x23\x17\xc4\x86\x1b\xef\xfa\xe4\xc6\x8f\x3a\xdc\xf8\x08\x62\x18\xae\x78\xac\xed\x9c\x87\xef\xa2\xd1\x4f\xa7\xb5\xe0\x20\x89\x1b\xf5\xae\x5d\x0a\xf5\x6e\xf4\x27\xea\x9d\x2d\xd7\x59\xab\x04\xea\x3f\xa9\xc3\xde\x57\xd1\xb5\xe1\xff\xf6\x5f\x23\x57\xaa\x1e\x79\x32\x72\x32\xb3\x1f\xee\xe7\x55\x8f\xe4\xce\x6d\xd2\x99\x07\xfa\x0a\x00\x00\xff\xff\x16\xec\x32\x27\x5d\x01\x00\x00"),
		},
		"/Chart.yaml": &vfsgen۰CompressedFileInfo{
			name:             "Chart.yaml",
			modTime:          time.Time{},
			uncompressedSize: 260,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x5c\x8f\x3d\x6e\x85\x40\x0c\x84\x7b\x9f\xc2\x17\x78\xc0\xeb\xa2\xad\xa2\x50\xa6\x4f\x6f\x16\x03\x56\xd8\x1f\xd9\xbb\x48\xe4\xf4\x11\x10\x94\x28\xad\xbf\xf9\x34\x1e\xca\xf2\xc1\x6a\x92\xa2\xc3\xed\x09\x91\x02\x3b\xf4\x6b\xb5\xc2\xfa\x50\x9e\xc5\x8a\xee\x30\xb2\x79\x95\x5c\xce\x58\x7f\x51\xbc\x29\x4e\x49\xf1\xfd\xc5\x60\x49\x87\xbc\x94\x92\xcd\xb5\xed\x40\xf1\x8b\xc4\xaf\xa9\x8e\x8d\x4f\x01\x20\x90\xc4\x42\x12\x59\xcd\xc1\x03\x39\x90\xac\x0e\x25\x4e\xe9\xf5\x7f\x16\xf1\x7a\xe4\xed\xbc\x63\x7f\x00\xb0\x54\xd5\xf3\xe9\xde\x1d\xb3\x94\xa5\x0e\x87\xf2\xb7\xae\x1d\xc8\x7f\xee\xa4\xa3\x01\x6c\xf7\xb8\xae\xe9\x9a\x67\x07\x94\xf3\xef\xde\x9f\xdb\x77\x00\x00\x00\xff\xff\x3c\x72\xe8\x6c\x04\x01\x00\x00"),
		},
		"/crds": &vfsgen۰DirInfo{
			name:    "crds",
			modTime: time.Time{},
		},
		"/crds/crds.yaml": &vfsgen۰CompressedFileInfo{
			name:             "crds.yaml",
			modTime:          time.Time{},
			uncompressedSize: 16609,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x3b\x5b\x6f\xdb\xc8\xd5\xef\xfa\x15\x07\xfe\x3e\x20\x49\x61\xd2\x09\xf6\xa5\x15\xb0\x48\x0d\x3b\x8b\xaa\xc9\x26\x86\xed\xcd\x43\x83\xb4\x18\x91\x47\xd2\xd4\xc3\x19\x66\x66\xa8\x84\x59\xec\x7f\x2f\xe6\x42\x8a\xa4\x78\x19\xc9\x4e\xba\x6d\xad\x97\x98\x73\x39\xf7\x3b\x99\x59\x14\x45\x33\x92\xd3\xf7\x28\x15\x15\x7c\x0e\x24\xa7\xf8\x45\x23\x37\x4f\x2a\xbe\xfb\xa3\x8a\xa9\x38\xdb\xbe\x58\xa2\x26\x2f\x66\x77\x94\xa7\x73\xb8\x28\x94\x16\xd9\x35\x2a\x51\xc8\x04\x2f\x71\x45\x39\xd5\x54\xf0\x59\x86\x9a\xa4\x44\x93\xf9\x0c\x80\x70\x2e\x34\x31\xcb\xca\x3c\x02\x24\x82\x6b\x29\x18\x43\x19\xad\x91\xc7\x77\xc5\x12\x97\x05\x65\x29\x4a\x8b\xa1\xc2\xbf\x7d\x1e\xff\x10\x3f\x9f\x01\x24\x12\xed\xf5\x5b\x9a\xa1\xd2\x24\xcb\xe7\xc0\x0b\xc6\x66\x00\x9c\x64\x38\x87\x84\x15\x4a\xa3\x54\xb1\xff\x43\xe2\x9a\x2a\x2d\x4b\x4b\x73\x42\x55\x22\xe2\x44\x64\x33\x95\x63\x62\xe9\x49\x53\x4b\x24\x61\x57\x92\x72\x8d\xf2\x42\xb0\x22\x73\xc4\x45\xf0\xd7\x9b\x77\x6f\xaf\x88\xde\xcc\x21\x36\x17\x2a\x98\x8b\x4b\x4b\xba\x43\xe8\x1f\x72\x49\x85\xa4\xba\x9c\xc3\x0b\xfb\xac\xcb\x1c\xe7\xa0\xb4\xa4\x7c\xbd\x07\x4b\x13\x5d\x28\xfb\x0f\x36\x20\xdd\xd8\xe5\xd0\xdb\xe6\x48\xe3\xf2\x6d\xf5\x18\x70\x35\x11\xdc\x71\xad\x3e\xbc\x7c\xfa\x67\x0b\xe9\xc7\x1f\x4f\x2e\xbc\xe4\x6e\x4a\x9e\x60\x7a\xf2\xec\xa3\x3f\xde\xa4\xd0\x6e\x85\xa2\xf1\xca\x6b\xdc\x7f\xdf\x58\x39\x42\x60\xb9\x14\x5b\x9a\xa2\x6c\x40\xbc\x6a\x2e\x1d\x01\x32\x35\xc6\x41\x97\x85\x6e\x13\x7a\xd9\x5d\x3e\x02\x34\x13\x09\x61\x54\x97\xb1\x31\xc1\x16\xf4\xeb\xdd\xc2\x11\x70\x33\x54\x8a\xac\xf7\x0d\x07\x7e\x6e\x6c\x1c\x01\x37\xd4\x2a\x7a\xf0\x97\x3c\x39\x04\xfb\x5a\x8a\x22\xaf\x5d\x75\xc0\x43\x1d\x70\x1f\x25\x7c\x84\x71\x17\xec\x0a\xa3\x4a\xbf\x6e\xae\xbe\xa1\x4a\x3b\xe4\xac\x90\x84\xed\x22\x81\x5d\x54\x94\xaf\x0b\x46\x64\xbd\x3c\x33\x54\xa2\x42\xb9\xc5\x5f\xf8\x1d\x17\x9f\xf9\x4f\x14\x59\xaa\xe6\xb0\x22\x4c\x19\x36\x54\x22\x0c\xd5\x3b\xa4\xaa\x58\x4a\x1f\xdd\x3c\x59\x4e\x72\x73\xf8\xf5\xb7\x19\xc0\x96\x30\x9a\xda\xd8\xe4\x36\x45\x8e\xfc\xfc\x6a\xf1\xfe\x87\x9b\x64\x83\x19\x71\x8b\x00\x29\xaa\x44\xd2\xdc\x9e\xab\x80\x03\x55\xa0\x37\x08\xee\x24\xac\x84\xb4\x8f\x15\x07\x70\x7e\xb5\xf0\xb7\x73\x29\x72\x94\x9a\x56\x14\x98\x5f\x23\x4e\xd7\x6b\x1d\x3c\x4f\x0c\x21\xee\x0c\xa4\x26\x32\xa3\x43\xe8\x5d\x14\x53\x50\x0e\xb5\x58\x81\xde\x50\x05\x12\xad\x74\xb8\x8b\xd5\x0d\xb0\x60\x8e\x10\x0e\x62\xf9\x4f\x4c\x74\x0c\x37\x46\x82\x52\x81\xda\x88\x82\xa5\x26\x9c\x6f\x51\x6a\x90\x98\x88\x35\xa7\x5f\x6b\xc8\x0a\xb4\xb0\x28\x19\xd1\xe8\x35\x55\xfd\x6c\xe0\xe5\x84\x19\x11\x16\x78\x0a\x84\xa7\x90\x91\x12\x24\x1a\x1c\x50\xf0\x06\x34\x7b\x44\xc5\xf0\xb3\x90\x08\x94\xaf\xc4\x1c\x36\x5a\xe7\x6a\x7e\x76\xb6\xa6\xba\xca\x4c\x89\xc8\xb2\x82\x53\x5d\x9e\xd9\xfc\x62\x3c\x59\x48\x75\x96\xe2\x16\xd9\x99\xa2\xeb\x88\xc8\x64\x43\x35\x26\xba\x90\x78\x46\x72\x1a\x59\xc2\xb9\x35\xff\x38\x4b\xff\xaf\x56\xf4\x93\x06\xa5\x1d\x33\x76\x3f\x6b\x9a\x83\x72\x37\x26\x6a\xb4\x4b\xfc\x35\x47\xff\x4e\xbc\x66\xc9\x48\xe5\xfa\xd5\xcd\x2d\x54\x48\xad\x0a\xda\x32\xb7\xd2\xde\x5d\x53\x3b\xc1\x1b\x41\x51\xbe\x42\xe9\x14\xb7\x92\x22\xb3\x10\x91\xa7\xb9\xa0\x5c\x7b\x4b\xa2\xc8\xdb\x42\x57\xc5\x32\xa3\xda\x68\xfa\x53\x81\x4a\x1b\xfd\xc4\x70\x61\xf3\x33\x2c\x11\x8a\x3c\x25\x1a\xd3\x18\x16\x1c\x2e\x48\x86\xec\x82\x28\xfc\xe6\x62\x37\x12\x56\x91\x11\xe9\xb4\xe0\x9b\x65\x45\xfb\xa0\x93\x56\xbd\x5c\xe5\xfa\x5e\x0d\x79\x0f\xbc\xc9\x31\x69\x79\x46\x8a\x8a\x4a\x63\xbd\x26\x49\x1b\x9b\x6f\x06\x9f\x61\x5f\xb4\xfe\x58\xe8\xcd\xc2\x88\xa8\xb5\xda\xc1\x7b\xee\x0f\xc1\x46\xb0\x54\x59\x91\xca\xcc\x3a\x1b\xe8\x0d\xd1\xfe\xf0\x12\x15\x6c\xc4\x67\x20\x7d\x1a\xb4\xe5\x13\xe1\xb0\x46\x6d\xca\xa2\xd4\xc8\x91\x30\xeb\x68\x24\x49\x50\xa9\x66\x10\x89\x3b\x57\x87\x88\xb7\x02\xc3\x44\xa2\xbe\xc6\xd5\xfe\x56\x87\x8b\x57\x9f\x0a\xba\x25\x0c\xb9\xb6\x91\xc3\x68\x2f\x7e\x6b\xc2\x76\x4e\x12\x4c\xcd\x5f\xf0\x99\xea\x8d\xcd\x36\xa0\xc9\x5a\xf5\x00\x1c\xa3\xa4\xce\x31\xbd\x3b\x03\x86\xd1\xbd\x6c\x89\x39\x12\x42\xaf\x41\x4d\x6c\xd5\x55\xe2\xa8\x01\xfc\xb2\xb8\x74\xd1\x16\xc1\x14\xbd\x91\x2a\x95\xc6\x6c\x47\xf0\x2c\x88\x4e\xe3\xbd\xc6\x4e\x9b\xa8\x22\x68\xd7\xa9\x63\xae\xe1\x12\xd8\x94\x73\xb8\xfa\xa2\xe9\x1e\x62\x69\x53\xe7\x51\xfe\xb1\x2b\x34\x46\x05\x74\x51\x1f\xb3\x4d\x02\xa1\xdc\x7b\x26\x5d\xad\x50\x1a\x93\xab\x01\x79\x3e\x50\x99\xd4\xb9\xa7\x45\x1b\x17\x07\xdc\x80\x6a\xcc\x7a\xec\xae\x4f\x0a\x35\x3d\x3b\x72\x76\x04\x34\xfd\xd7\xa4\x6f\xd2\x63\x4b\x03\x24\x4c\x79\x00\x23\x4a\xff\x05\x89\xd4\x4b\x24\xda\xb4\x3d\xfd\xa6\xdc\x22\xf9\x4d\xf7\x4e\x55\x61\x18\x60\xa0\xcd\x82\x93\x4a\xc5\xc0\x80\x77\x7c\x26\xaa\xce\x04\xbd\x47\x1c\xd7\x73\x30\x47\x22\x03\x77\x76\x84\x93\x19\xa2\x6e\x25\xe1\x8a\x56\x9d\x5d\x20\x8b\xed\x4b\x7d\x3c\xe2\x24\x8b\xc9\x86\xf0\x35\xa6\x2e\x71\x0a\x8e\xde\x96\x6c\x14\xe5\x42\x6f\xfa\x14\xf6\x60\x9c\xfb\x5a\x3a\x80\x5d\x5f\x5e\xbb\x5a\x62\x53\x64\x84\x47\x12\x49\x4a\x96\x0c\x2b\x28\x40\x79\x4a\x13\x62\x6b\x8a\x14\x35\xa1\x4c\x0d\xf0\x4c\x96\xa2\xd0\x3b\x59\x79\x8e\x9d\x24\xe2\x63\xf8\x90\x48\x54\xbb\xfc\x1c\x60\xe3\xda\x1e\x74\x5c\x3c\x5d\x4a\x8a\xab\x67\xfe\xf2\xae\xea\xad\x14\xf6\x44\x59\xf2\x06\x78\xb8\x3f\xd1\xfb\xc1\x6f\x80\x68\x1f\xff\xbc\x79\x79\xc4\x3e\x76\xd7\xd4\xc6\xf0\x8e\xdb\x40\x78\x2b\x0b\x3c\x1d\x20\xfa\x27\xd3\x5b\x9c\x82\xef\x38\x8e\xa2\xda\x6e\x4f\xd3\x7c\x5b\xe6\xb5\x43\x98\x2b\x35\xbd\xbe\xe3\xd8\xd1\x7d\x38\x11\x7d\x49\xa7\x4a\x3d\x8d\x99\x41\x7b\xa3\x9e\x57\x1c\x94\x5e\x89\x94\xa4\x6c\xed\x34\x3b\xf6\xf9\x2c\x90\x6a\x93\x61\xaf\xa4\xf8\x52\xfa\x26\x68\x4f\xeb\x03\x79\x60\x44\x0c\x43\xf4\x19\x54\x0c\xf5\xb7\x47\xc4\x90\xa4\x28\xfb\x45\xb0\x14\x82\x21\x69\x87\xbc\x6a\x1e\x31\x3f\xa0\x12\x74\xa3\x8b\xf9\xec\x60\x03\x59\xf7\xf1\x3e\xc2\x7f\x70\x29\xb6\x2f\x07\xf3\xfb\x2a\x38\x7e\x37\x74\x83\x46\x3b\x10\xcd\x07\x11\x55\xf3\xac\xe0\x0b\xb6\xd8\x0a\x3e\xdd\x17\x29\x06\x0f\x6f\xf7\x27\x08\x23\xe7\x7b\x44\xd0\x59\xda\x4d\x6e\x5f\x10\x96\x6f\xc8\x8b\xdd\x9a\x1f\xae\xba\xd9\x51\x63\xdb\x34\x1e\xa6\xa6\x9c\x83\x96\x05\xfa\x01\x8b\x90\x46\xa2\x6e\x65\x17\xb0\x4d\x7f\x93\x6b\xd7\x61\xb4\x46\x44\x27\x27\xad\x19\x90\x7d\x6c\xd4\x9b\xf0\xe1\xe3\xcc\x41\xc5\xb4\xf6\x50\xb3\xf8\x9f\x3b\xf0\xae\x27\x15\x25\x4f\x64\xc1\x30\x74\xf2\x7d\xec\x10\xae\xe2\xf7\xa6\xe4\xc9\x75\xc1\xb0\x33\x8d\xeb\x6e\xef\x8d\xe5\xf6\xe8\xed\xcc\xe7\xba\xfb\xbf\x8f\x41\x5d\x97\xad\x81\x89\x5d\x3d\xc1\x31\xd4\x83\x21\xff\x71\x76\xf7\x38\xbb\x7b\x9c\xdd\x3d\xd8\xec\x6e\xa8\x58\xb1\xd1\xcc\x7b\xc9\xeb\x8e\x56\xa7\x8a\x1c\x17\x09\x0f\xad\x71\xee\x7a\xb0\x4c\x5e\x1a\xc8\xb2\x41\x15\x60\x4f\xc1\x61\xe3\x67\x60\x85\x39\x3e\x64\xc8\x88\x4e\x36\xfd\x75\xd2\x48\x09\x35\x3d\xbc\xb3\x31\xae\x9b\x02\xfb\x7f\xa3\x78\xc2\xb1\x35\xf8\x39\x0f\x43\xec\xa9\x6c\xbc\x09\x0e\xc5\x12\x34\x82\x0c\xec\x77\x7a\x39\x78\xf5\xc5\x04\x0d\x15\xc6\x41\x80\x00\xfb\x06\xd1\x0d\xfd\x80\x42\x86\x89\x16\xb2\x6a\xf1\x32\xe4\x1a\xa8\x0a\x80\x09\x26\x4c\x56\xb7\xed\x04\xbb\x9e\x92\xb9\xb0\x7f\x0a\x04\xee\xb0\x74\x19\x82\xf0\x20\x90\x46\x0f\xa4\x06\x28\xd1\xe6\x1f\x37\x32\xc5\xd2\x02\xf2\x29\x25\x00\x5a\x7e\x90\x56\xc1\x60\x08\x3b\xd8\x91\xa7\xa1\xac\x1e\x44\x2d\x91\xd9\x05\x4b\xbf\x1d\x1e\x78\x11\x05\x42\x36\x05\x42\xce\x28\xda\x08\x1f\x78\xe7\x20\x83\x6c\x4a\xf9\x28\x76\x6b\x15\xed\xf2\x9b\x53\xf4\x13\xe5\x14\x66\x6c\x77\x43\xf3\x60\x86\xb5\xb0\x96\x64\x5f\x28\x54\x05\xc3\x7b\x53\xb1\xd5\xa8\x14\x10\x89\xb0\xe0\xa7\xc1\x30\xdf\x0a\xbd\xe0\xa7\xf0\xea\x0b\x35\xc9\xd2\xd8\xcd\xa5\x40\xf5\x56\x68\xbb\xf2\xcd\x04\xeb\xc8\x3f\x4a\xac\xee\xaa\xad\x3e\xb8\x6b\x42\x8d\x3c\x9a\x75\x88\x8a\x83\xd9\x5f\xb8\xd1\x4f\xad\x2a\xaa\x4c\x65\x20\x64\x25\x17\x5b\x4d\x5a\x98\xe1\x66\x69\x49\xca\x0a\x65\x0b\x0e\x2e\x78\x84\x59\xae\xcb\xb8\x07\x57\x30\x4c\xaf\x1e\x21\x5b\xda\x69\x92\xd7\x40\x1b\x0c\x75\x89\xe0\x49\xbb\x35\x35\x96\x83\xe0\xaa\x64\x46\x12\x4c\x21\x2d\xac\x50\x49\x30\x44\xa5\x25\xd1\xb8\xa6\x09\x64\x28\xd7\x08\xb9\x89\xd4\xa1\xda\x08\x0e\xd2\x3e\x6a\x11\x6d\xea\xeb\x39\xfc\xfd\xe9\xd3\x0f\xe7\xd1\xdf\x48\xf4\xf5\x79\xf4\xa7\x8f\x1f\xa2\xfa\xef\x7f\xc4\x1f\xff\xf0\xec\x65\x63\xef\xd9\xcb\xff\x0f\x77\xb6\x43\x4d\x7a\x7c\x0c\xd3\x37\x18\x1e\x1a\x15\xee\xff\x22\x13\x36\x82\xce\x55\xd6\x15\x9c\x9e\x83\x32\x6e\x38\x6f\x81\x40\x43\xc0\x99\x2c\x89\x5c\x7f\xb7\xc2\x28\x28\xad\x1d\x60\x16\xd6\x31\x03\xaa\x2b\x5e\xbe\x5b\x4d\x1f\x8b\x3c\x6a\xd3\x55\xae\x51\x06\x9f\x0f\xb4\xe0\x2f\xd1\x5d\xb1\x44\xc9\x51\xa3\x8a\x28\xd7\x91\x90\x91\xbb\xda\x98\x39\xdd\xcf\x94\xa7\x8d\x38\x72\x32\xfb\x5e\x06\x66\x8b\x90\xfb\x16\xde\x9d\x8a\xd1\x15\x36\x75\xb9\x67\x7b\x64\xb7\xf6\xa9\x40\x59\x82\xd8\xa2\x9c\x0c\xa8\x3e\xc7\xd7\xdd\xba\x09\xd0\x76\x82\x52\x30\xbb\x61\x0b\xe0\x37\x96\x7a\x3f\x55\x68\x17\xc4\xb3\xa9\x0c\x85\x70\xfe\xf6\xd2\xb4\xc1\xe7\xdc\xa5\x80\x2e\xdd\x16\xa2\xc9\x2a\x8c\x79\x59\x4f\x26\xd5\x73\x3b\x7c\x1b\x02\xc4\x45\x18\x9c\x03\xfb\x98\x83\xba\x80\x96\xaa\xba\xd7\xbd\xaa\xa8\xb2\x12\x6e\x73\x11\x1e\xcd\x33\x37\xc5\x70\xea\xda\xad\x34\x44\xfe\xcd\x7a\x95\x8e\xe0\xdb\x6d\x4a\xa3\x05\x09\x4a\x66\x01\x6d\x4a\xbb\x05\x09\x82\xfa\xd8\xa6\x3c\xb6\x29\x8f\x6d\xca\x63\x9b\xf2\xbf\xd4\xa6\x3c\xf6\x11\x47\xf2\xd6\x28\x72\x7e\x27\x23\xca\xfd\xfa\xc1\xd7\x60\x36\xbf\x66\x24\x37\x1e\xfe\xab\x49\x91\xd6\xd8\x7f\x83\x9c\x50\x19\xe4\xe5\xe7\xf6\x2d\x1f\xc3\xd6\x6d\xca\xad\xe3\x34\x11\x19\x1c\x54\x01\xd6\xdf\x96\xce\xc2\xc2\x31\x07\x64\xae\x14\xa8\xaa\xc7\x46\xe5\x73\x0a\x9f\x37\x42\xb9\x8c\xbc\xa2\xc8\xd2\x00\xa0\x54\xc1\xc9\x1d\x96\x27\xa7\x7b\x71\xe9\x64\xc1\x4f\x5c\x89\xd0\xf5\xfa\x00\xb0\x75\xc5\x21\x38\x2b\xe1\xc4\xde\x3e\xb9\x5f\x39\x15\x6c\x9d\x0f\xd8\x58\xd4\x5f\xaa\xde\xb7\xb9\x08\x34\xcf\x10\x9a\x1c\x63\xaf\xc7\x6b\xa2\xd0\x12\x6b\xec\xc3\xe3\x83\x5d\x6b\xf2\x43\xe4\xa3\xa4\x31\xaa\xc8\xa9\x98\x19\x35\x54\x38\x3b\x12\xcb\xb8\x52\xb2\x62\xf4\xf5\xcb\xb4\x22\x82\x5e\x1d\x85\xe9\x93\xa4\x13\xc9\xe3\x98\xe8\x1a\xac\xfe\x40\xbf\x93\x98\x89\xed\x84\x89\x04\xa5\xe6\x03\x09\x1b\xf7\xaa\x00\xe2\xa7\xde\xbf\x1e\xaa\xad\xc1\x77\xb2\x07\xb3\x78\x37\x4a\xcd\x01\x80\x46\xde\xdd\x1e\x08\x2b\x40\x9e\x53\x53\x9b\x47\x9b\xff\xf7\xdb\xbc\xd8\xa2\x94\x34\x1d\x93\xd8\x24\xdd\xa1\xb9\x28\x27\x52\xe1\xfb\x90\x59\xeb\xf0\x07\xa0\x3d\xef\x16\x36\x0f\x2b\xd3\x87\x02\xb6\x0d\x67\xf4\x21\x12\x65\x88\x45\xa8\x92\x27\x37\x23\x1f\x8c\x87\x89\x7e\x94\x98\x83\x3f\x86\xee\xff\xcf\x3f\xdd\x50\xdc\xda\xdc\x7d\x71\x37\x82\x73\xff\xc3\xf8\xff\xba\xef\x3d\xff\x15\x00\x00\xff\xff\xe4\x96\xac\x23\xe1\x40\x00\x00"),
		},
		"/examples": &vfsgen۰DirInfo{
			name:    "examples",
			modTime: time.Time{},
		},
		"/examples/test.yaml": &vfsgen۰CompressedFileInfo{
			name:             "test.yaml",
			modTime:          time.Time{},
			uncompressedSize: 153,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x3c\xcc\x4d\x8e\xc2\x30\x0c\x40\xe1\x7d\x4e\x91\x0b\xb8\x23\x0f\x69\xeb\x66\x0b\x1b\x2e\xc0\xde\x75\x2c\x88\xfa\xab\x38\x20\x71\x7b\x84\x40\x6c\x3f\x3d\x3d\x00\x70\xbc\xe7\x8b\x16\xcb\xdb\x1a\xbd\xcc\x77\xab\x5a\x8a\x5e\xb3\xd5\xf2\x6c\x26\xb2\x46\xb2\xc9\xd6\xc8\xb6\xfc\x3d\x90\xe7\xfd\xc6\xe8\xa6\xbc\xa6\xe8\x8f\x9f\xd8\x2d\x5a\x39\x71\xe5\xe8\xbc\x5f\x79\xd1\xdf\x06\xd0\xd9\xae\xf2\xf6\xaf\x9c\x4f\xd1\xa7\x16\x03\x92\xf6\xd0\x86\x5e\x21\x74\x34\x02\xa7\x51\x60\x90\xff\xa1\xa3\x03\x0d\x48\xc1\xbd\x02\x00\x00\xff\xff\x73\x74\xf2\x63\x99\x00\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Time{},
		},
		"/templates/_helpers.tpl": &vfsgen۰CompressedFileInfo{
			name:             "_helpers.tpl",
			modTime:          time.Time{},
			uncompressedSize: 2035,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x54\x41\x6b\x1b\x3d\x10\xbd\xfb\x57\x0c\xcb\x17\xf8\x9a\xb2\xf2\xa1\xd0\x83\x21\x87\x92\xf6\x50\x5a\x52\x68\x20\x3d\x16\x59\x3b\x8a\x87\x6a\xb5\xaa\x46\x72\x63\x12\xff\xf7\x22\x69\xbd\x59\x27\x8e\xb3\x29\xb9\x89\xd5\x9b\x37\x33\xef\x3d\xed\xed\xed\xfc\x14\xd6\xd4\x2e\x80\x31\x80\x26\x83\x61\xe3\xf0\xac\x8d\x1c\xa4\x5a\xe1\x02\x4e\xe7\xdb\xed\x2c\xa1\x66\x9f\x6e\x9c\xb4\x0d\x84\x15\x82\x95\x2d\x42\xa7\xf3\x59\xad\xa4\x0f\x62\xd6\xe3\x6a\x68\x50\x93\x45\xa8\x94\x89\x1c\xd0\xd7\x1e\xaf\x89\x83\xdf\x88\x54\x54\x41\x7d\x0f\x93\xd1\x04\x10\xe7\xb9\xfe\x22\x31\x8a\x2b\x69\x22\x72\x46\x7e\x5b\xa3\xf7\xd4\x20\xdc\x41\xf0\xd1\x2a\x78\xff\x2e\x1f\xa9\xbd\x8c\x5a\xd3\x0d\x54\x75\x05\x3d\x17\xda\x26\x1d\xcb\x98\xe7\x1e\x65\x40\x90\x43\x07\x1d\x8d\xd9\xc0\xef\x28\x0d\x69\xc2\x06\xa4\x73\x79\x01\x31\xfb\x81\x85\x3b\xe3\x43\xea\x90\x96\x61\x58\xa2\x92\x91\x11\xb8\x6b\x11\xbe\xc4\x25\x7a\x8b\x01\xb9\xac\xad\x09\x4d\xc3\x20\x3d\x82\xa1\x96\x02\x36\x10\x3a\x08\x2b\x62\xf8\x7f\xb9\xc9\x92\x7c\xbc\xb8\x4c\x58\xb2\xd7\xc0\x0e\xd5\x1b\x31\xfb\xac\xc1\xa3\x41\xc9\xbd\x76\xaa\xb3\x41\x92\xe5\xa2\x5e\xf9\x46\x01\xfe\x90\x31\xb0\x44\x88\x9c\xe6\x64\x90\x79\xf8\x7e\xda\xe7\x15\x4e\xe0\x7d\x95\x49\x0f\xa2\xee\x2e\x07\x61\x7b\xc8\x93\xf7\x53\x84\x37\x3c\xf0\xfc\x97\x97\x58\x9c\x4d\x77\xf6\x7e\xc6\x41\x8e\x42\x22\xbe\x17\xad\x4a\xed\x6e\xce\xbd\x8f\x2f\x1c\xce\x79\xb2\x41\x43\x75\xc2\xf5\x09\x57\x0f\xb8\x4a\xd3\xe9\x39\x3b\x7c\xdc\x4b\xdf\xc8\xd6\xf4\x66\xd6\xe8\x99\x3a\x9b\x2c\xcd\xd6\xf6\x39\x29\x28\x23\x97\x68\xa6\xd8\x9b\xe1\xf7\xde\x3e\xdc\x69\x2c\x77\x39\x5f\xf5\x6d\xef\xc0\xa3\x33\x52\x21\x54\x6f\x2b\xa8\x7e\x56\x2f\x79\x54\x47\x67\xaa\x93\x73\xbe\x33\x06\xfd\xa3\xf4\x01\x59\x65\x62\x73\x3c\xa8\x02\xb6\xdb\x11\xc9\x03\x41\xa7\x35\x9e\xd8\xf4\xf5\x1a\x66\xc7\x38\x2b\x25\x9d\x5b\xc0\xb1\xb6\x87\x15\x12\x7d\xad\xf8\x35\xfc\x5d\x04\x75\xf3\x74\x39\x9d\x6e\x44\xb5\x42\xd3\x0a\x5e\xcd\x73\x44\x8e\x33\xec\x62\xf4\xc4\x08\xad\xb4\xf2\x1a\x9b\x7a\xb9\xc9\x34\xc3\x4b\xb9\x44\xbf\x26\x85\x87\x8b\xc8\x72\x90\x56\xe1\x7e\xc9\xee\xf5\x3e\xc6\xf7\xef\xa1\xc0\x4b\x56\x3f\x38\xf7\x74\x5c\x0f\x92\xa8\xae\x75\x9d\x45\x1b\x16\x70\x44\xa5\x03\x85\x4e\xfa\x50\x77\xfa\x19\x99\x46\xea\xfe\x4b\x46\x18\x0d\xaa\xd0\xf9\xaf\x7d\x56\xea\xd7\x35\xfc\xa5\x1e\x8c\x56\xf8\x1b\x00\x00\xff\xff\x0e\x0d\x26\xe8\xf3\x07\x00\x00"),
		},
		"/templates/deployment.yaml": &vfsgen۰CompressedFileInfo{
			name:             "deployment.yaml",
			modTime:          time.Time{},
			uncompressedSize: 2959,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xbc\x96\xdf\x8f\xe2\x36\x10\xc7\xdf\xf7\xaf\xb0\x78\xdf\xac\xae\xaa\xaa\x2a\x6f\xb9\x90\x3b\xad\xc4\x02\x0a\x74\xa5\x7b\x42\x83\x33\x80\x55\xc7\xb6\xec\x49\x5a\xba\xc7\xff\x5e\x99\x84\x90\x84\xc0\x1d\xdd\x55\x79\x42\x33\xf1\xf7\xf3\x1d\x8f\x7f\x81\x11\xaf\x68\x9d\xd0\x2a\x64\x60\x8c\x7b\x2a\x3f\x3d\xfc\x29\x54\x16\xb2\x31\x1a\xa9\xf7\x39\x2a\x7a\xc8\x91\x20\x03\x82\xf0\x81\x31\x05\x39\x86\xec\xed\x8d\x09\xc5\x65\x91\x21\x1b\x71\x59\x38\x42\xfb\x68\x71\x2b\x1c\xd9\xfd\x23\xd7\x8a\xac\x96\x12\x6d\xb0\x29\xa4\xf4\x23\x46\x2c\x60\x87\x43\x3d\xdc\x19\xe0\x95\x46\x90\xa2\x44\x70\x18\x4c\x4f\xe1\xea\x2b\x09\x6b\x94\xce\xe3\x18\x7b\x7b\x7b\xfc\x39\x56\x35\xc8\x93\xbe\x33\x25\x54\x86\x8a\xd8\xaf\x5e\xcf\x19\xe4\x5e\xcb\xa2\x91\x82\x83\xab\xd0\xaf\x20\x0b\x74\xc1\x29\x58\x81\x1d\x4a\xe4\xa4\x6d\x85\xce\x81\xf8\x6e\xd2\xf2\x72\x87\x9b\x93\xd2\x64\xc0\xd5\x6f\x15\x8c\x30\x37\x12\x08\x6b\x58\x6b\x92\xfd\x0f\x94\xd2\x04\x24\xb4\x6a\xe0\x35\x7e\xd3\x98\x17\x8e\x84\x0e\x2c\x96\xc2\x77\xb0\x52\xad\x7e\x55\x66\x0d\xea\x1f\x10\x5c\xea\x22\x0b\x84\x7e\xb2\x58\x76\x6a\xbf\x3e\xdc\x73\x50\x65\xfd\xd0\x5f\x82\x76\xcd\x68\xa3\xb3\xe8\xec\xb1\xff\x29\xe9\x6f\x90\xcb\x4e\xd9\xbf\xdf\x24\xc8\xce\x3c\xbf\xb3\xef\x0d\xea\xd4\xfb\xe3\x7f\xe4\x85\x15\xb4\x8f\xb5\x22\xfc\x9b\xc2\x41\xbf\xe7\xe2\x16\xdd\xcf\x87\x0b\x71\x68\x4b\xc1\x31\xe2\x5c\x17\x8a\xa6\xef\xdb\x1c\xfe\xe7\xf3\x20\x14\xda\xd6\x54\x3c\xd6\xbb\x2e\x07\x05\x5b\xb4\x4d\xfc\x46\x49\x57\xca\x72\x57\x6b\xfa\xf4\x4b\xbb\x3b\x8c\x89\x1c\xb6\x18\xb2\x51\x7b\xb9\xf8\x90\xdf\x30\xda\x09\xd2\x76\xcf\x0e\x87\xf0\x22\x4d\xb0\x65\xdf\x59\x86\x1b\x28\x24\xb1\x20\xde\x81\xa5\x20\x32\xa6\x3e\x65\xd8\xe1\x30\xea\x53\xe6\x85\x94\x73\x2d\x05\xdf\x77\x57\xe7\x51\xcf\x34\xc9\xae\x3f\xa3\x2d\xb9\x6e\xbd\xcd\x34\x21\x59\xc1\x5d\x27\xd7\x9a\xd8\xb9\xb6\xd4\x01\xd5\x3d\x0c\xbc\x64\x17\x72\x04\x59\x4d\x9a\x6b\x19\xb2\x65\x3c\x6f\xe5\xa4\x28\x51\xa1\x73\x73\xab\xd7\xd8\x35\xb2\x23\x32\x5f\xb1\xd7\x0d\xc6\x0c\xd0\x2e\x64\x4f\xc3\xf6\xcc\xd1\xd5\x65\xce\x22\x64\xe2\x7f\xe1\x38\x5d\x58\x8e\xee\x87\x6b\xa8\xf9\xf2\xc6\xea\x41\x55\x0e\xf7\xe6\x25\x59\xa6\xcf\xf1\x62\x15\x8d\xc7\x69\xcf\x59\xe9\xe5\x43\x36\x0a\xaf\xb7\x66\x34\xa8\x39\x49\xa2\x71\x92\xae\x92\x49\x12\x2f\x9f\x67\xd3\x55\x32\x8d\x3e\x4f\x92\xf1\x15\xf9\x96\x7a\xfb\x10\x41\xc8\xd0\x26\xfe\xcc\x16\x5a\x05\xa8\x60\x2d\x31\xfb\x69\xe4\x34\x7a\x49\xde\xc3\xf3\xaa\x77\xc1\x16\xf3\x28\xbe\x45\x1c\xba\x58\xaf\xa8\xcf\xbe\xae\xbe\xcc\xd2\x97\x68\x79\x5f\x01\x7a\x1b\x6c\xb4\xcd\x81\x6e\x2a\xbf\x26\xe9\xe7\xd9\xe2\x79\xf9\xed\x6e\xf1\x12\xed\xda\x1f\x34\xfb\x0b\x7d\x7f\xbe\x9e\xef\x40\xa9\x39\xc8\xb8\x3a\x66\x4f\xf3\x38\x68\x67\x9e\xce\x5e\x9f\x17\x7e\x02\x27\xb3\x38\x9a\xac\xe2\xc9\x1f\x8b\x65\x92\xfe\xd8\xd8\x10\xe1\xc2\x51\xef\xb6\x6c\x83\x6f\xf6\xeb\x8b\xd5\x79\x7f\x03\x33\xb6\x11\x28\xb3\x14\x37\x97\x99\x3a\x37\x3f\xee\xf1\xd3\x9b\x21\x68\x9e\x55\x0f\x57\x2e\x6b\xa5\x33\x5c\xd4\x4f\x92\xb3\xcf\x76\x34\xbc\xe3\xfa\xbe\xb8\xbc\x2f\x78\xb0\xd9\x08\x55\x35\xef\xf4\x9e\xa9\x23\x1f\xcb\x21\x2d\xd1\xf6\x5f\x20\xad\xe0\xc7\xd2\x9a\xeb\x6a\x81\xdc\x22\xb5\x90\xfd\xcc\x7f\xe4\xfe\x1b\x00\x00\xff\xff\x90\xbe\x63\xc0\x8f\x0b\x00\x00"),
		},
		"/templates/poddistruptionbudget.yaml": &vfsgen۰CompressedFileInfo{
			name:             "poddistruptionbudget.yaml",
			modTime:          time.Time{},
			uncompressedSize: 434,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x91\xc1\x4a\x03\x31\x10\x40\xef\xfb\x15\x43\xef\x8d\x2c\x88\x87\xbd\x29\x1e\x8b\x88\x87\xde\x67\x93\x69\x1d\x9c\x9d\x84\x64\x52\x28\x6b\xff\x5d\xd2\x50\x10\xf4\xa0\xd7\x90\xf7\xde\x24\xb3\xae\x5b\xe0\x03\xb8\x3d\x4a\xa5\xe2\x52\x0c\xcf\x5c\x72\x4d\xc6\x51\x9f\x6a\x38\x92\x39\x52\x9c\x85\x02\x5c\x2e\x03\x26\xde\x53\x2e\x1c\x75\x82\x14\x85\xfd\xf9\xee\x34\xce\x64\x38\x0e\x1f\xac\x61\x82\xd7\x9f\xfc\xb0\x90\x61\x40\xc3\x69\x00\x50\x5c\x68\x82\x75\x05\x56\x2f\x35\x10\x6c\xbc\xd4\x62\x94\xb7\x99\x8e\x5c\x2c\x9f\xb7\x3e\xaa\xe5\x28\x42\xd9\x1d\xaa\x48\x23\x36\xe0\x5a\xbd\xe3\x25\xa1\xef\x0e\xf7\x46\x42\x58\xc8\xbd\xdc\x8e\xfb\x2d\xc1\x99\xa4\xb4\x1c\xc0\xf5\x7d\x7f\x69\x75\xa8\x95\x3e\x41\x59\x03\xa9\xc1\x7d\xf3\x95\x44\xbe\xb9\x16\xd6\xc7\x13\xb2\xb4\xcf\x98\x60\x1c\x00\x0a\x09\x79\x8b\xb9\x97\x16\x34\xff\xbe\xfb\x96\xfe\x47\xfc\x66\xda\xfd\x32\xc4\x43\x1b\xa2\x99\x48\xaf\x3b\xf8\x0a\x00\x00\xff\xff\xa2\xa6\x23\xc5\xb2\x01\x00\x00"),
		},
		"/templates/rbac-leader-election.yaml": &vfsgen۰CompressedFileInfo{
			name:             "rbac-leader-election.yaml",
			modTime:          time.Time{},
			uncompressedSize: 1093,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xdc\x93\xbb\xae\x14\x31\x0c\x86\xfb\x3c\x85\xb5\x7d\xe6\x08\x89\x02\x4d\x07\x0d\x1d\xc5\x22\xd1\x7b\x93\x7f\xf7\x84\xe3\x4d\x46\x8e\x33\x08\x86\x79\x77\x34\x97\x45\x08\x04\x5a\x2e\x15\xd5\x58\x1e\xf9\xfb\xac\x5f\x0e\x0f\xe9\x1d\xb4\xa6\x92\x7b\xd2\x13\x87\x8e\x9b\x3d\x16\x4d\x9f\xd8\x52\xc9\xdd\xd3\x8b\xda\xa5\xf2\x30\x3e\x73\x4f\x29\xc7\x9e\x8e\x45\xe0\xae\x30\x8e\x6c\xdc\x3b\xa2\xcc\x57\xf4\x34\x4d\x94\x72\x90\x16\x41\x87\x20\xad\x1a\xd4\x2b\x2e\xa9\x9a\x7e\xf4\xa1\x64\xd3\x22\x02\xed\xce\x4d\x64\x99\x38\x50\x47\xf3\xec\x05\x1c\xa1\x1e\x82\xb0\xd8\x76\x5c\x1d\x38\x6c\xcc\xee\x08\x01\x57\x74\x6f\x6e\x6d\x9a\x67\x47\x24\x7c\x82\xd4\x45\x4f\x34\x4d\xfe\x3e\xf7\x36\xb4\x98\x3f\x53\x4e\x39\x22\x1b\x3d\x5f\x78\xda\x04\xb5\x77\x9e\x78\x48\xaf\xb5\xb4\x61\x25\x7b\x3a\x1c\x1c\x91\xa2\x96\xa6\x01\x7b\x2f\x94\x7c\x4e\x97\x2b\x0f\xd5\x11\x8d\xd0\xd3\xde\xbf\xc0\xd6\xaf\xa4\xba\x15\x1f\xd8\xc2\xe3\x36\xa2\x60\xc3\x5a\xb6\x21\xde\xca\xe1\xeb\xff\x08\x81\xe1\x77\xf5\x0f\xd5\xd8\xda\x4f\xb6\xf8\xc1\x73\x17\x1c\x23\xb2\x7d\x47\xdc\x97\xf7\xde\xbb\x3f\xb8\x94\x57\x29\xc7\x94\x2f\xff\xdd\xc1\x14\xc1\x11\xe7\x05\x77\x8b\xf5\x17\x91\x38\xa2\x6f\xde\xce\x3f\x0e\xa0\xb6\xd3\x7b\x04\x5b\xcf\x77\xb3\xbc\x85\x8e\x29\xe0\x65\x08\xa5\x65\xfb\x4b\xdf\x9d\x01\x7f\x09\x00\x00\xff\xff\x4f\x3b\x02\xe0\x45\x04\x00\x00"),
		},
		"/templates/rbac.yaml": &vfsgen۰CompressedFileInfo{
			name:             "rbac.yaml",
			modTime:          time.Time{},
			uncompressedSize: 2970,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x55\xb1\x6e\xdb\x30\x10\xdd\xf5\x15\x07\x8d\x05\xc4\xa0\x40\x87\x42\x5b\xdb\xa1\x4b\xd1\xc1\x01\x02\x14\x45\x06\x9a\x7c\x91\xd9\xd0\xa4\x70\x24\x1d\xb4\x6e\xfe\xbd\x10\x65\x3b\x6e\x64\xa7\x4a\xe3\x28\x99\x74\x22\x78\xef\xdd\x3d\x92\xef\x64\x6b\x2e\xc0\xc1\x78\x57\xd3\xea\x6d\x71\x6d\x9c\xae\xe9\x1c\xbc\x32\x0a\x1f\x94\xf2\xc9\xc5\x62\x89\x28\xb5\x8c\xb2\x2e\x88\x9c\x5c\xa2\xa6\xf5\x9a\x8c\x53\x36\x69\x50\xa9\x6c\x0a\x11\x5c\x31\x1a\x13\x22\xff\xac\x94\x77\x91\xbd\xb5\x60\x71\x95\xac\xed\x32\x4a\x12\x74\x7b\xbb\x49\x0f\xad\x54\x3d\x86\x98\xc1\x42\x06\x88\xaf\xdb\xe5\x7e\x97\x95\x73\xd8\xd0\xd1\x11\xad\xd7\xd5\x38\xae\x3e\xa9\x63\xfa\x4d\xce\x38\x0d\x17\xe9\x5d\x8f\xd7\x61\xdc\x98\xb8\x20\x71\x21\x6d\x42\x10\xe1\xaf\x06\x85\x74\xce\x47\x19\x8d\x77\xa1\x4f\xd8\x5b\xb8\xab\x22\xfa\x6f\x72\x69\x8f\xe1\xc3\xe9\xee\xa7\xaa\xaa\x62\x5f\x53\x9e\x4b\x25\x64\x8a\x0b\xcf\xe6\x57\x46\x14\xd7\xef\x83\x30\xfe\x6c\xa7\xf6\xa7\xbe\xa9\x99\xb7\x38\x20\xf5\x03\x2d\x57\xb2\x69\x18\x8d\x8c\xd0\xc5\x36\x34\xde\xcd\x92\x45\x97\xaf\xee\x70\xcf\x61\xa1\xa2\xe7\xdc\x4d\x45\x4b\x19\xd5\xe2\xcb\x9e\xc8\x34\xa0\xc9\x55\x2a\x13\x94\x17\xca\x2f\xcf\x0e\x92\xd6\x54\x46\x4e\x28\x0b\x4e\x16\xa1\xa6\xef\x97\xa7\x6f\xff\xbf\x6f\xda\xa9\xef\xd0\x49\x24\x2a\x2a\x92\xad\xf9\xcc\x3e\xb5\x9d\x5e\xdb\x5a\x0e\x23\x96\x97\x05\x11\x23\xf8\xc4\x2a\xcb\x5b\xbe\xc9\x4b\x2b\xf0\x7c\x73\x90\x0d\x62\xfe\x5a\x13\xfa\xe0\xa6\x3b\xd9\x1c\x29\x86\x8c\xc8\x61\x6a\xf5\x36\xd4\xb0\xd8\x84\x6d\xde\x7a\xaf\xa0\x7b\x9c\x79\x23\x56\x70\x31\xbc\x00\xf1\xce\x2c\x42\xff\xeb\xf5\x26\x0a\x50\x8c\x91\x25\x8d\xe0\x39\x08\x37\xb6\x8d\x27\x5d\xf8\x8f\xc6\x69\xe3\x9a\xd7\x7b\xef\xd9\x5b\xcc\x70\xd5\xc1\x0d\x5f\xeb\x23\x2d\x8a\x76\x27\xf1\x80\x4c\x45\x48\xf3\x1f\x50\x31\xbf\x95\x83\xd3\x68\x92\x19\x54\x0c\x0e\x76\xfc\x70\x1c\x14\xc3\x90\x1a\x3c\x92\xf9\xf9\x07\x48\x5f\xce\x84\xc3\x63\x40\x38\xc5\xe0\x28\x8f\xf4\x5d\xbe\xc4\x6c\xf8\x97\x02\x13\xcd\x85\xc9\x3c\xf7\x99\x4c\xf1\xd8\xcb\xda\x73\xa9\x31\x16\xf3\x28\x2f\x1b\xbe\x96\xd1\x0e\xf5\x34\x23\xf8\x13\x00\x00\xff\xff\xce\x17\xa9\x76\x9a\x0b\x00\x00"),
		},
		"/templates/service.yaml": &vfsgen۰CompressedFileInfo{
			name:             "service.yaml",
			modTime:          time.Time{},
			uncompressedSize: 470,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x90\xcd\x6a\xeb\x30\x10\x85\xf7\x7e\x8a\x21\x7b\x0b\x2e\xdc\x95\xb7\xdd\x96\x12\xda\x92\xbd\x2a\x9f\xa6\xa2\x63\x49\xcc\x8c\x03\xc1\xcd\xbb\x17\x45\x0a\xdd\xb4\xd0\x7a\xe7\x39\x3f\x1f\x47\xbe\xc4\x03\x44\x63\x4e\x13\x9d\xfe\x0d\xef\x31\xcd\x13\x3d\x41\x4e\x31\x60\x58\x60\x7e\xf6\xe6\xa7\x81\x28\xf9\x05\x13\x6d\x1b\xc5\x14\x78\x9d\x41\xbb\xc0\xab\x1a\x64\x14\x1c\xa3\x9a\x9c\xc7\x90\x93\x49\x66\x86\xb8\xd7\x95\xb9\x26\x76\xe4\xe8\x72\xe9\x71\x2d\x3e\xb4\x0e\xf7\x08\x86\x57\xb8\x87\xdb\xb9\xb9\xd8\xbf\x80\xb5\xe2\x88\xb6\x6d\xfc\x1d\xab\x85\x2a\xe9\x83\x52\x4c\x33\x92\xd1\xff\xda\xa7\x05\xa1\x76\xd9\xb9\x74\xec\xc1\xf3\x0a\x75\xda\xf6\xb9\x2a\x34\x70\xc9\x62\x9d\x3b\x5e\x7f\xbe\xf5\x57\xa1\xf9\xeb\x67\x5e\x8e\xb0\xfd\xd5\xbc\xc0\x24\x06\xed\x4a\x91\x6c\x39\x64\x9e\xe8\xf9\x6e\xdf\x6f\xed\xfd\xde\xcc\xca\xf8\x65\x56\x30\x82\x65\xf9\xe3\xe2\x5b\xec\xfe\xa7\xe5\x9f\x01\x00\x00\xff\xff\xb2\xbe\xa1\xef\xd6\x01\x00\x00"),
		},
		"/values.yaml": &vfsgen۰CompressedFileInfo{
			name:             "values.yaml",
			modTime:          time.Time{},
			uncompressedSize: 989,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\x53\x41\x4f\xdb\x4c\x10\xbd\xef\xaf\x18\x85\x73\x8c\x21\x04\xf8\x7c\xcb\x07\xe8\x13\xd2\x47\x15\x89\xf6\x50\x55\x3d\x4c\xd6\xe3\xb0\xb0\xde\x71\x67\x66\x43\x0d\xe2\xbf\x57\xb6\x49\x8b\x40\xea\x2d\xfb\xf2\xe6\xe9\xcd\xbc\xe7\x03\xb8\xa4\x06\x73\x34\xd8\x61\xcc\xa4\xd0\xb0\x80\x8f\x59\x8d\x64\x2e\xb4\x0d\x6a\xd2\xcf\x3d\x27\x13\x8e\x91\xc4\x1d\xc0\xe7\xbb\xa0\x10\x14\x10\xbe\xae\x6e\xfe\x9f\x37\x2c\x2d\x9a\x51\x0d\x4d\x88\x54\xb8\x41\xd1\x47\x14\x82\x1d\x4a\xc0\x4d\x24\x05\x63\xd8\x10\x74\xa8\x4a\x35\x84\x64\x0c\x3d\x67\x01\xa3\xb6\x8b\x68\xa4\x85\x73\x42\x5d\x0c\x1e\xb5\x82\x23\xe7\x22\x7b\x8c\x17\x93\x8b\xca\x01\x1c\x80\x76\xe4\x43\xd3\x0f\x4a\x98\x8d\x5b\xb4\xe0\x31\xc6\x1e\x3a\xe1\x5d\xd0\xc0\x09\xec\x8e\xf6\xce\x81\x37\xf7\xe4\x0d\x72\xc7\x09\x9a\x20\x6a\xa0\x86\x62\x0e\x20\x61\x4b\x15\xcc\x66\xce\x05\xb5\xc0\x83\xba\xd0\xa4\x30\xc1\x1d\xd7\xab\x94\xd8\xd0\x02\x27\xad\xe0\xdb\x77\x17\x5a\xdc\xd2\x3a\xc7\x78\x4b\x5e\xc8\x26\x70\x20\xde\x92\xcf\x12\xac\xbf\xe0\x64\xf4\xd3\x46\xb1\x9c\x56\xfa\x45\x49\x2a\x38\x5d\x2e\x17\x27\x7b\xe8\x3f\xe1\xdc\xed\x31\xfd\x38\x87\x31\xf2\xe3\x5a\xc2\x2e\x44\xda\xd2\x95\x7a\x8c\xa3\x83\x0a\x1a\x8c\x4a\x93\x87\xc9\x6d\xc7\x1a\x8c\xa5\xaf\xa0\x5c\x2c\x4e\xfe\x39\x3f\x5d\x9e\x2d\x97\x67\x45\xfd\x20\x05\x79\x29\xb2\xce\x09\xd5\xe6\xc7\x05\xb6\xf8\xc4\x09\x1f\xb5\xf0\xdc\x1e\x6e\x30\x3d\x61\xf0\x91\x73\x7d\xf8\xb7\x84\x01\x0c\xb7\x15\xec\xca\xa2\x2c\x8e\x4a\x07\xd0\xe5\x18\xd7\x1c\x83\xef\x2b\xb8\x6e\x3e\xb1\xad\x85\x94\x92\x39\x97\xb8\xa6\x5b\x8a\xe4\x8d\xa5\x82\xe7\x17\x87\x4d\x13\x52\xb0\x7e\x7c\x18\x47\x92\x37\x67\x14\x52\xce\xe2\x49\xa7\x35\x7e\x64\x52\x1b\x7f\x03\xb4\xd4\x8e\x0b\xcd\x8e\xca\xf2\x26\xcc\x46\xcc\x77\x79\x02\xda\xe1\x1d\x43\x1b\x3e\xb0\x8f\xdf\xb3\x17\x23\xdb\x29\xc9\x2e\xf8\xf1\x5c\xd6\x77\x54\xc1\x6b\x97\xae\xd7\xc3\x3a\x2c\x56\xc1\x79\x79\x5e\xfe\x26\xae\xbc\xe7\x9c\xa6\x1c\xde\x66\xff\xfc\x32\xc6\x7c\x19\x54\x72\x37\x60\xff\xe6\x7a\x4b\x23\x8f\xd2\x50\xed\x7a\x1f\x8f\xfb\x73\xc0\xe1\xdf\x48\x58\x93\x5c\x0d\x97\x19\x32\x1c\x1d\xbe\x9b\x18\xa0\xd7\x32\x7e\x48\x63\x1a\x9f\xd3\xeb\xfc\xb8\x3f\x6f\x27\x99\xe9\x7b\xab\xe0\x5e\x39\x8d\xc0\x8e\x64\x33\x34\x62\xe8\x83\x03\x78\x64\x79\x20\xd1\x0a\x8e\xdd\xaf\x00\x00\x00\xff\xff\x4f\x74\x8c\x3b\xdd\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/.helmignore"].(os.FileInfo),
		fs["/Chart.yaml"].(os.FileInfo),
		fs["/crds"].(os.FileInfo),
		fs["/examples"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
		fs["/values.yaml"].(os.FileInfo),
	}
	fs["/crds"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/crds/crds.yaml"].(os.FileInfo),
	}
	fs["/examples"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/examples/test.yaml"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/_helpers.tpl"].(os.FileInfo),
		fs["/templates/deployment.yaml"].(os.FileInfo),
		fs["/templates/poddistruptionbudget.yaml"].(os.FileInfo),
		fs["/templates/rbac-leader-election.yaml"].(os.FileInfo),
		fs["/templates/rbac.yaml"].(os.FileInfo),
		fs["/templates/service.yaml"].(os.FileInfo),
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
