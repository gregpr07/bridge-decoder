// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets.go
// call_tracer.js
package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func assetsGoBytes() ([]byte, error) {
	return bindataRead(
		_assetsGo,
		"assets.go",
	)
}

func assetsGo() (*asset, error) {
	bytes, err := assetsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1633419907, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _call_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\xdd\x73\xdb\x36\x12\x7f\xf7\x5f\xb1\xc9\x43\x2d\x4d\x14\x49\x4e\x7a\xbd\x19\xb9\xea\x8d\xea\x28\xa9\x67\xdc\x38\x63\x2b\xcd\x64\x32\x79\x80\xc8\xa5\x84\x1a\x02\x58\x00\x94\xcc\x6b\xfd\xbf\xdf\x2c\x08\x52\xfc\x12\xad\xb4\x73\xbd\x6b\xf5\x24\x92\xbb\x0b\x60\x3f\x7e\xfb\x41\x8e\x46\x70\xa1\xe2\x54\xf3\xd5\xda\xc2\x8b\xf1\xd9\x3f\x61\xb1\x46\x58\xa9\xe7\x68\xd7\xa8\x31\xd9\xc0\x2c\xb1\x6b\xa5\xcd\xc9\x68\x04\x8b\x35\x37\x10\x71\x81\xc0\x0d\xc4\x4c\x5b\x50\x11\xd8\x1a\xbd\xe0\x4b\xcd\x74\x3a\x3c\x19\x8d\x32\x9e\xd6\xc7\x24\x21\xd2\x88\x60\x54\x64\x77\x4c\xe3\x04\x52\x95\x40\xc0\x24\x68\x0c\xb9\xb1\x9a\x2f\x13\x8b\xc0\x2d\x30\x19\x8e\x94\x86\x8d\x0a\x79\x94\x92\x48\x6e\x21\x91\x21\x6a\xb7\xb4\x45\xbd\x31\xf9\x3e\xde\xbc\x7d\x0f\x57\x68\x0c\x6a\x78\x83\x12\x35\x13\xf0\x2e\x59\x0a\x1e\xc0\x15\x0f\x50\x1a\x04\x66\x20\xa6\x3b\x66\x8d\x21\x2c\x9d\x38\x62\x7c\x4d\x5b\xb9\xf5\x5b\x81\xd7\x2a\x91\x21\xb3\x5c\xc9\x01\x20\xa7\x9d\xc3\x16\xb5\xe1\x4a\xc2\xcb\x7c\x29\x2f\x70\x00\x4a\x93\x90\x1e\xb3\x74\x00\x0d\x2a\x26\xbe\x3e\x30\x99\x82\x60\x76\xcf\x7a\x84\x42\xf6\xe7\x0e\x81\x4b\xb7\xcc\x5a\xc5\x08\x76\xcd\x2c\x9d\x7a\xc7\x85\x80\x25\x42\x62\x30\x4a\xc4\x80\xa4\x2d\x13\x0b\x1f\x2e\x17\x3f\x5c\xbf\x5f\xc0\xec\xed\x47\xf8\x30\xbb\xb9\x99\xbd\x5d\x7c\x3c\x87\x1d\xb7\x6b\x95\x58\xc0\x2d\x66\xa2\xf8\x26\x16\x1c\x43\xd8\x31\xad\x99\xb4\x29\xa8\x88\x24\xfc\x38\xbf\xb9\xf8\x61\xf6\x76\x31\xfb\xfe\xf2\xea\x72\xf1\x11\x94\x86\xd7\x97\x8b\xb7\xf3\xdb\x5b\x78\x7d\x7d\x03\x33\x78\x37\xbb\x59\x5c\x5e\xbc\xbf\x9a\xdd\xc0\xbb\xf7\x37\xef\xae\x6f\xe7\x43\xb8\x45\xda\x15\x12\xff\xe3\x3a\x8f\x9c\xf5\x34\x42\x88\x96\x71\x61\x72\x4d\x7c\x54\x09\x98\xb5\x4a\x44\x08\x6b\xb6\x45\xd0\x18\x20\xdf\x62\x08\x0c\x02\x15\xa7\x47\x1b\x95\x64\x31\xa1\xe4\xca\x9d\xf9\xa0\x43\xc2\x65\x04\x52\xd9\x01\x18\x44\xf8\x76\x6d\x6d\x3c\x19\x8d\x76\xbb\xdd\x70\x25\x93\xa1\xd2\xab\x91\xc8\xc4\x99\xd1\x77\xc3\x13\x92\x19\x30\x21\x16\x9a\x05\xa8\xc9\x38\x0c\xa2\x84\xd4\x2f\xd4\x4e\x82\xd5\x4c\x1a\x16\x90\xa9\xe9\x7f\xe0\x9c\x91\x59\xc0\x7b\xba\xb2\x86\x9c\x16\x34\xc6\x4a\xd3\x7f\x21\x72\x3f\xe3\xd2\xa2\x96\x4c\x38\xd9\x06\x36\x2c\x44\x58\xa6\xc0\xca\x02\x07\xe5\xc3\x90\x1b\x65\xe6\x06\x2e\x23\xa5\x37\xce\x2d\x87\x27\xbf\x9e\x00\x00\xf8\x4d\x1a\xcb\x82\x3b\xda\x23\x2d\x11\x24\x5a\xa3\xb4\xa4\xcd\x44\x1b\xbe\x45\x47\x02\x19\x8d\x57\xe9\xfc\xa7\x1f\x01\xef\x31\x48\x32\x61\x24\xaa\x90\x33\x81\x4f\xbf\x3e\x7c\x1e\x9c\xe4\x0b\x84\x68\x02\x94\x21\x86\xee\xa0\x77\x06\x76\x6b\xa7\x5a\xd8\xe1\xe9\x16\xe1\xe7\xc4\xd8\x12\x4d\xa4\xd5\x06\x98\x04\x95\x90\xeb\x97\xd5\xc4\xa5\x55\xb9\x4c\x46\x97\x12\xb5\x5b\x35\x5b\xbf\x10\x31\x81\x88\x09\x83\xfb\x0d\x18\x8b\x31\x1d\x8e\xcb\xad\xba\xa3\x25\x94\x26\xa7\xd6\x29\xa8\x38\x50\xa1\x0f\x0f\x3a\x56\x71\x2a\x34\x99\x50\x62\x9d\x40\x94\x48\xb7\x85\x9e\x50\xab\x01\x84\xcb\x3e\x64\xda\xf3\xf2\x2f\x58\x6c\x13\x8d\x4e\xd5\xa8\xb5\xd2\x06\xf8\x66\x83\x21\x67\x16\x45\x5a\x50\x6e\x99\xce\x1e\xc3\x14\x84\x5a\x0d\x57\x68\xe7\x74\xd9\xeb\x9f\x17\x34\x3c\x82\x5e\x46\xf3\x64\x3a\x75\x20\x15\x71\x89\x61\x79\x41\xfa\xd9\x35\x37\xc3\x88\x25\xc2\x16\x5b\x3a\xaf\x10\x68\xb4\x89\x96\xfb\x7b\x0f\xe5\xfd\x7e\x40\x50\x52\xa4\x10\x10\x54\xb1\x25\xc5\xb8\x49\x8d\xc5\x8d\xd7\x87\x19\x40\xc4\x0c\xa9\x9f\x47\xb0\x43\x88\x35\x3e\x0f\xd6\x48\xd6\x97\x01\x56\xce\x63\x52\xe3\x9c\x63\x0a\xb4\x91\xa1\x8a\x87\x56\xbd\x4d\x36\x4b\xd4\xbd\x3e\x7c\x05\xe3\xfb\x68\xdc\x87\xe9\xd4\xfd\xa9\x9e\xd2\x73\xd6\x4f\x46\x42\x55\xec\x35\xe4\xc4\xdd\x5a\xcd\xe5\xaa\xac\xa4\xca\x61\x2e\x23\x60\x20\x71\x07\x81\x92\x2e\x74\xc8\xd2\x4b\xe4\x72\x05\x81\x46\x66\x31\x1c\x00\x0b\x43\xb0\x2a\x73\xee\xc2\x95\xdb\x76\x03\x5f\x7d\x05\x3d\x5a\x7e\x0a\xa7\x17\x37\xf3\xd9\x62\x7e\x0a\xbf\xfd\x06\xd9\x9d\xa7\xd9\x9d\x17\x4f\xfb\x6d\x9b\xe6\xf2\x3a\x8a\xfc\xbe\x9d\xfc\x61\x8c\x78\xd7\x3b\xeb\x0f\xb7\x4c\x24\x78\x1d\xf5\x6a\x26\xca\x98\xe6\x32\x84\xa9\x67\x7e\x56\x67\x7e\x51\x61\xae\x70\x8f\x46\x30\x33\x06\x37\x4b\x81\x4d\x5c\xf0\xc0\xe1\x30\xc4\x58\x02\x4e\x72\xf9\x40\x6d\x62\x81\xe4\xc7\x8d\x7d\x78\x1b\x56\x4f\xe5\x1c\x2d\x8d\x71\x42\x7f\x54\x3c\x68\x3c\xa4\x50\x75\x0f\xad\xfa\x01\xef\x9d\x03\xe4\x46\x20\xe7\x9e\x85\xa1\x46\x63\x7a\xfd\x7e\x93\x95\xcb\x38\xb1\x93\x0a\xeb\x06\x37\x4a\xa7\x43\x43\x20\xda\x73\x0a\x19\x64\xfa\x69\xe3\x5f\x31\x73\x29\x89\xdf\x07\xd2\x1b\x66\x7a\xed\x64\x17\xca\xd8\x49\x4e\x46\x17\x6d\x74\x4e\xcb\x24\xee\x74\x7c\x7f\xda\xb4\xc3\xb8\xbf\xf7\xc3\xb3\x6f\xfa\x15\xf6\x87\xf3\x66\x64\x16\x50\x38\x8c\x13\xb3\xee\x39\x3f\x6f\x21\xdb\x83\xde\x14\xac\x4e\xf0\xf8\x08\x76\x4e\xdf\x74\x78\x83\x22\x22\x1c\xb4\x3a\x09\x9c\xe3\xaf\x98\x03\x5a\x87\x6f\x8c\x32\x90\x49\x96\xce\xd6\x56\xa9\x43\xfe\xef\xdd\xff\x76\x7e\xf5\xfa\xd5\xfc\x76\x71\xf3\xfe\x62\x71\xda\xe6\xf0\x02\x23\x4b\x1b\xaf\x1e\x58\xa0\x5c\xd9\x75\xf5\xb0\xb4\x40\x95\xec\x13\x31\x3f\x3f\xfb\x9c\xdd\x81\x69\x17\xd6\x35\xb5\x5a\xe7\x86\x4f\x9f\xab\x0b\x3e\x74\x98\xa4\xca\x9c\x19\xe8\xcf\xf3\x7b\xab\x26\xf9\x1f\x62\xb5\x2a\x27\xee\xf6\xb8\xff\x41\x08\x84\x4b\xa2\xfe\x9e\x09\x26\x03\xec\x38\x63\x47\x64\xb4\xe6\xa4\x47\x50\x7c\x83\x76\xad\x42\x97\xaa\x03\x96\xa5\xfd\xdc\xbb\x43\x25\xf1\xf7\x62\xf9\xec\xea\xaa\x84\xe4\xee\xfa\xe2\xfa\x55\x19\xdd\x4f\x5f\xcd\xaf\xe6\x6f\x66\x8b\x79\x9d\xf6\x76\x31\x5b\x5c\x5e\xb8\xbb\x0d\xe0\x1f\x8d\xe0\xf6\x8e\xc7\x2e\xf3\xbb\x2c\xa9\x36\xb1\xeb\x6e\x8a\xed\x9b\x01\xd8\xb5\xa2\xbe\x41\xfb\x52\x27\x62\x32\xc8\x2b\x0f\xd3\x88\x2a\xab\x28\xa6\x0e\xb9\xc5\x59\xcd\x2d\x9a\x71\xc6\xcd\x3b\x8d\x7e\x1b\x61\xcf\xaa\xc6\x96\x5b\x2c\x01\x8d\x90\x71\x59\xd8\xa5\xb3\xde\xf1\xda\x81\x7f\xc1\x18\x26\x70\x56\x4f\x55\x1d\xd9\xf1\x05\x3c\xa3\x75\xfe\x48\x8e\x7c\xd9\x22\xe2\xef\x94\x29\x1b\x88\xf1\xff\x9f\x4d\x55\x62\xaf\xa3\x68\x02\x75\x53\x7d\xdd\x30\x55\x2b\xef\x15\xca\x26\xef\x3f\x1a\xbc\x5d\x59\x98\x22\x41\xc5\xf0\xa4\xe1\xb8\x59\x7e\x7b\x52\x0b\xeb\x16\x73\xba\xbe\xc2\xad\x05\xd3\x03\x55\xc1\x8b\x6a\x30\x1e\x9f\x82\xfe\x8b\x55\x41\x6b\x4b\x45\x8d\x53\xb5\x63\x1a\x90\x20\xcd\x71\x8b\xc0\xed\xa9\x71\xeb\x50\x97\xa9\x76\x84\xf6\x43\xf8\x80\x65\xb9\x12\xd1\x61\xae\xef\x4d\x49\xbb\xae\x4b\xa3\xfe\xd2\x4f\x19\x5c\x58\x30\xd7\x3f\x6a\x84\x0d\x4b\x61\x89\xd4\x36\xdd\xa5\xe4\x35\x10\xa6\x92\x6d\x78\x60\xca\x52\x5d\x77\xaa\x71\xc5\xb4\x13\xae\xf1\x97\x04\x8d\xa5\xe6\x5d\x86\xc0\x02\x9b\x30\x21\x52\x58\xf1\x2d\x4a\x27\xa3\xf7\xe2\xe5\x78\x0c\xc6\xf2\x18\x65\x38\x80\x6f\x5e\x8e\xbe\xf9\x1a\x74\x22\xb0\x3f\xac\xc0\x7f\x55\x89\x75\xdb\x12\x85\x77\xdf\x57\x18\xdb\x75\xaf\x0f\xdf\x1d\x28\x62\x8e\x28\x44\x5a\xf9\xe0\x39\x9c\x7d\x1e\xd2\x96\xa7\x95\x80\xaa\x39\x08\xa0\x30\xd8\xb2\xc4\x68\x04\x8b\xeb\x57\xd7\xbd\x3b\xa6\x99\x60\x4b\xec\x4f\xdc\xc4\xc7\x69\x78\xc7\x7c\xcb\x4f\x06\x85\x58\x30\x2e\x81\x05\x81\x4a\xa4\x25\xa3\xe5\xad\xbb\x48\x29\x59\x9e\xda\x36\xd9\x6e\x50\xc2\x82\x00\x8d\xc9\xf3\xa8\xb3\x3e\xed\x97\x6d\x48\x12\x70\x69\x78\x88\x25\xbb\x12\x38\x2a\x97\xe4\x3c\xc5\x8e\x0b\xd1\x26\x7c\xa3\x0c\x2d\xbe\x44\xd8\x69\x45\x05\x29\x97\x81\x1b\xc1\x85\x48\x56\x33\xa0\x24\x30\x10\xca\xcd\xfd\x1c\x6a\x01\xd3\x2b\x33\xcc\xb2\x28\x6d\x81\xe0\x57\xaa\xdd\xf0\xb1\x60\x2a\x47\x89\x6b\xf7\x3b\x2a\x65\x09\x78\xcf\x8d\x75\x2d\x21\x9d\x87\x1b\x1f\x47\x5c\xae\x06\x10\xab\xd8\xe5\xba\xe3\x2a\x0a\x9f\xf6\x6e\xe6\x3f\xcd\x6f\x9a\x75\xf1\xf1\xee\x91\x0f\x02\x9e\x16\x53\x14\xd0\xb8\x45\x6d\x31\x7c\x7a\x6c\x2b\xdf\xe2\xcb\xd3\x03\xbe\x4c\x8b\xb6\x14\x2f\xef\x4a\x67\x17\xcc\xd8\xbd\xbd\x57\x98\x4d\x44\xca\xdb\x33\x89\xb0\xcd\x8a\xc5\x67\xc4\x3a\xc4\xa9\xb8\x91\x89\x69\xbf\x0e\x57\x29\x57\xd6\x7b\xec\xca\x83\x7d\xab\xdd\x1e\x20\x97\x25\x73\xed\x5c\x63\x93\x31\x94\x70\xcd\x3d\xcf\x3b\x24\x96\x25\x58\x77\x32\x95\x58\xf2\x3b\x2a\xbf\xda\x51\x7f\xc5\xcc\x7b\xe3\x1c\xcb\xe3\xfe\x92\xaf\x2e\xa5\xed\xe5\x0f\x2f\x25\x3c\x2f\x28\x29\x07\xc2\xf3\x4a\x9c\x77\xa5\x05\x70\x83\x2a\x81\x16\x61\x2f\xee\xbc\x7e\x8b\x84\xd6\x74\x97\x6b\x5b\xa3\x6d\x96\x51\xe3\x96\x55\x48\xd9\x4f\x34\xda\x21\xfe\x92\x30\x61\x7a\xe3\xd6\x4a\xb0\x38\xb4\x55\xee\xff\xb4\xd1\x98\x90\x88\x6a\x2b\xd2\x5c\xab\x90\xe2\x75\x9b\x4b\xc9\x7a\x88\x0b\x15\x62\xa7\xc0\x16\x89\x1e\x1d\x0b\x8f\xf1\xe1\xf2\x58\x9f\x58\x6c\xa4\x08\xaf\xa2\xda\x8b\x18\x17\x89\xc6\xa7\xe7\x6d\x00\x6b\x12\x1d\xb1\xc0\x39\x8d\x41\x70\x83\x35\x03\x46\x6d\x70\xad\x76\xcd\xbd\x1d\x8d\xe3\x4d\x37\x2d\x3c\xb2\x96\x85\xdd\xc8\x99\x19\x48\x0c\x5b\x61\xc9\x4d\x5b\xcd\x9a\xbb\x49\xf7\x8c\xb0\xa2\x8f\x2f\xf7\xe9\x67\xc5\xe5\x17\xba\xf7\xc3\x9f\xea\xb7\x35\x8f\x6b\x94\xbe\x39\x91\x2b\x80\x4b\x17\xf9\xf9\xb2\x7a\xf3\xaf\xed\x82\xf0\x05\xa0\xd2\xc5\x97\x69\xe6\xbc\x7e\xef\x0a\x65\x57\x6d\xfb\x65\x1e\x59\x50\x1e\xf2\xc4\xe3\x8b\x6a\x0a\x30\xf9\x33\x06\x76\x1f\x64\xae\xd2\xa5\xab\x58\xe3\x96\xab\x84\x0a\x0e\xfc\x3b\x4f\xae\x6a\x4d\x44\xc6\xfe\xb0\x7f\xf5\xe1\xfc\xa8\xfc\xee\x63\xb7\xf6\x2f\xf3\xb2\x0a\xbb\x94\xdf\x95\xab\x94\xfc\x1b\x91\x28\x7b\xcd\x46\x52\x9c\x88\x47\xdf\x81\x78\xa8\xb3\x2a\xa6\x1a\xd0\xd7\x11\x42\x23\x0b\xd3\xa2\xb0\x19\x64\x55\x29\xac\x99\x0c\x7d\x33\xce\xc2\x90\x93\x54\x17\x1d\xb4\x55\xb6\x62\x5c\x36\xeb\xf9\x2f\xa8\xa9\x3a\x5d\xb0\xb3\x7f\x2a\x57\x43\x7e\x54\xc3\x05\x86\xee\x30\x27\x65\x07\xea\x2c\x79\x2a\x9e\x7e\xe8\x7d\x4f\x79\xd9\x0b\x25\x4d\xb2\x71\xfd\x17\xb0\x2d\xe3\x82\x2d\x85\xaf\xc7\x65\x08\x81\x40\x26\xb3\x97\xc1\x18\x59\xb5\x45\x6d\x2a\xea\x39\x2e\xf4\xfe\x50\xd8\xd5\x32\x48\x7e\xd9\xa2\xc4\x2f\x45\xa1\x63\xd1\xa7\xac\xae\xd7\x82\x59\xeb\x9d\xb8\x64\xa0\x2c\xf6\xb9\x75\xdf\x15\xa0\xb4\x15\x83\x1d\x17\xf1\xae\x9e\x26\xca\xef\x60\xdc\xd6\x37\xfe\x95\x90\xe0\x51\x67\xbf\x2a\xca\x7d\xaf\x44\xab\xd4\x00\x04\x32\x37\x13\xc8\xbf\x1e\xa8\xf6\x42\x8f\x8c\x31\x4a\xb8\x93\x35\x0b\x0d\xe0\x71\xaf\x1e\xd6\x98\x4f\x3f\xb3\x66\x74\x89\x28\x81\x5b\xd4\x8c\xba\x7f\xf2\x70\xff\xf2\x9b\x76\x6f\x72\x89\xce\xdc\x9c\x80\xc2\xcb\xf6\x6f\xa2\xa9\x9c\xe2\x72\x95\x61\x55\xf6\xa8\x04\x56\x81\xbd\xaf\x83\x55\x56\x8d\x38\x11\xf5\x39\x5e\x31\xc3\xdb\xbb\xbe\xbd\x77\x4d\x49\x75\x50\x55\x8c\xf3\xf6\x9c\xae\xfa\x20\x6a\x7a\x56\x9b\x6b\x15\x23\xbc\x16\xf2\xc6\x3c\xaf\x98\xc3\x17\x3f\x1f\xb0\x44\xed\x1e\x56\x22\xb5\xca\xbc\x62\xa6\xb6\x54\x2d\xda\xed\x7d\x33\xd8\x1b\x22\x28\xce\x27\x8f\x88\x20\x9a\x2e\x31\xc5\x48\xb2\xe5\xcc\xee\x59\x8d\x3e\xab\xe3\x26\xad\xf4\xd9\xb3\xba\x52\xf9\xa6\xcd\x56\x7c\x53\xb2\xd5\xc3\x79\x57\x26\x19\xe7\xc1\xf7\x48\xc6\x20\x57\x29\xc2\xf4\x80\x8c\x43\xed\x79\x93\xfc\xa8\x24\xe5\x96\xcc\x73\xc7\x01\x19\xa5\x25\x4b\x95\xaa\xbd\xff\x1d\x2b\x14\x5c\x87\x8e\x51\xa1\xaf\x48\x76\x6f\x59\x1a\x8f\xdb\xa6\x1a\xd4\xe4\x7b\xc2\xbc\x68\x9f\x4e\x9f\x8e\xef\x9b\xaf\xd3\x7d\x06\xa8\x10\xb7\xed\x2c\x43\x08\xff\x31\x04\x41\x03\xff\x37\xfa\xbd\x34\x11\x29\x27\x00\x8d\xd9\x77\x01\xae\x1d\x23\x40\x52\x4b\x57\x44\x26\x86\xcb\x55\x09\x66\x42\x34\x5c\x63\x08\x11\x47\x11\x82\x0a\x51\xbb\xe1\xd4\xcf\xc6\xbf\x15\x18\x8d\xc0\xa0\xe6\x24\x34\xfb\xa4\x26\xfb\xba\xcd\x7d\xe8\x23\x79\x80\x36\x85\x08\x99\xfb\x30\xc4\x2a\x88\x99\x31\xb0\x41\x26\xb9\x5c\x45\x89\x10\x29\x28\x1d\x22\xc9\x2f\xcf\x55\x08\xe7\x14\x24\x06\xb5\x81\xdd\x5a\xf9\x7a\xc9\x35\x10\x31\x75\x50\xdc\x0e\xfc\x70\x97\x9b\x58\xb0\x14\xb8\xf5\x45\x9a\x3f\x5d\x19\xfa\x6a\x5f\x57\xb8\xcf\x35\x14\x99\xa2\x1b\xf7\x8a\x41\x4c\x07\xe6\x39\x1a\xba\xd3\x01\x74\x7e\xac\xd0\x81\x6e\xfb\x19\x7b\x07\x8a\xe5\x85\x43\x07\x4a\x95\xeb\x93\x0e\x14\x72\x64\xee\x4e\x07\xf4\x94\x9a\xca\x2a\x95\x73\xee\x8a\x28\x77\xa7\x03\x94\x32\x05\x54\x00\x29\xbf\x6d\x2a\x82\xdc\x9d\x41\x8b\x83\x93\xbf\xf5\xc8\x6e\x77\x98\x52\x46\xce\xcc\xd7\x56\x9a\x64\x4f\x3e\xdd\x61\xfa\xf9\xf1\x4a\xc4\xc7\x57\x89\xe7\x50\xe9\x51\x45\x81\x8c\xe1\x18\xd0\x2c\x36\xce\xa7\xe3\x73\xe0\xdf\x96\x39\xf3\xb2\x0b\xf8\xb3\x67\x6d\xbb\x2b\xd3\x7e\xe2\x9f\x73\x00\x2c\x02\xbc\xf6\xfc\x60\x8b\xd8\xc0\x89\x8c\xd1\x03\xc3\xc9\xc3\x7f\x02\x00\x00\xff\xff\x72\xce\x1a\x33\xb9\x2a\x00\x00")

func call_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_call_tracerJs,
		"call_tracer.js",
	)
}

func call_tracerJs() (*asset, error) {
	bytes, err := call_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "call_tracer.js", size: 10937, mode: os.FileMode(420), modTime: time.Unix(1633419875, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets.go":      assetsGo,
	"call_tracer.js": call_tracerJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"assets.go":      &bintree{assetsGo, map[string]*bintree{}},
	"call_tracer.js": &bintree{call_tracerJs, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}