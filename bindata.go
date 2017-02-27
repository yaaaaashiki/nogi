// Code generated by go-bindata.
// sources:
// data/members.toml
// data/songs.toml
// DO NOT EDIT!

package main

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataMembersToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x99\xdf\x52\x1a\x49\x14\xc6\xef\x79\x8a\x2e\x9f\x00\x50\x41\x2e\xf6\x62\xf7\x05\xf6\x01\x52\xa9\x2d\x93\x58\xd1\xaa\xa8\x5b\x89\x37\xb9\xb3\x1b\x44\x1c\x90\xc1\x3f\x08\x44\x0c\xf8\x07\x35\x24\x6a\x44\xdc\x00\x11\x78\x98\x33\xdd\x03\x6f\xb1\xd5\x3d\xc0\x9a\x99\xee\x4e\x4d\xd6\x2a\xaa\x72\x21\x4e\x7e\x7e\x7d\xce\x77\xbe\xd3\xf3\xec\xd9\xf2\xc2\xf2\x8b\x85\xb7\xef\x9e\x3f\x0f\xac\xcc\x2f\x2f\xa0\xdf\xd0\x94\x7d\x91\xa6\x1b\x71\x64\x97\x2b\xf4\xcc\x9c\x0a\x2c\x2e\xbd\x9d\x7f\x3d\xbf\x32\xcf\x7f\x04\x98\x00\xde\x06\x42\x00\x5f\x22\xc0\x3d\xc0\x75\xc0\x67\x53\x81\x17\x4b\x6f\xd7\x16\x5f\xcd\xbf\xe7\xdf\x09\xc5\x62\xd3\xb4\xdd\x9c\x63\xe5\x54\x38\xc8\x0a\xb5\xa9\xc0\x8b\x37\xab\xab\xaf\xfe\x5a\x7b\xff\xb7\x78\xfc\x1f\xf4\x63\x7a\x2a\xf0\x72\x75\xe5\xdd\xda\xc2\x9b\x37\xf3\x6b\x4b\xab\x2b\xce\xa3\x0b\x80\x0b\xb4\x73\x31\x15\x58\x5c\x58\x7a\xbd\xb8\x26\x9e\x35\x1b\x79\xb9\x3c\x15\x08\x48\x31\xf7\x2b\xf6\xfe\x57\x64\xdf\xdf\xb3\x93\xcb\x81\x71\xeb\x26\x4d\x00\x36\x01\x57\x10\xe0\x14\x10\x03\x70\xda\x83\x19\xa5\xed\x66\x88\x63\x86\x25\x98\xbf\xab\x30\xfb\x80\x3f\x00\xce\x00\xc1\x6e\xd8\x48\x50\x07\x3b\xbc\xd8\x45\xc3\xcd\x0c\xad\xa5\x24\xa4\x7b\x80\x7b\x48\x60\xd6\x3d\x98\xb3\x1c\x33\xcc\x39\x63\x32\x4e\x95\x9e\x24\x01\x38\xeb\xd1\x73\x5a\x89\x68\x3d\x18\x83\xc2\x19\x02\x9c\xe6\x1c\x64\x4f\x42\x79\x09\x38\xf9\xf8\x0b\x92\x63\x9f\xe5\x9c\x11\x09\xe7\x9f\x2a\x3d\xef\x01\x57\x00\xef\x7b\x48\x67\x7e\x46\x6a\x37\xf7\xe5\x62\x8e\x30\x8b\x40\x36\x80\xec\xc9\x24\x9d\xe3\x92\xf2\xa3\x9f\x96\x55\xa8\xf2\xe8\x13\x80\xcf\x3d\x87\x3e\xfb\x33\x4e\xab\xb5\x69\xe7\x92\x03\xf3\x46\x8d\xda\x53\x54\x68\x84\xb6\x9b\x61\x55\x23\x29\x15\x4d\x02\x36\x3c\x72\x6a\x30\x3b\x79\xab\x65\x20\xfa\xd5\xb4\x4b\x3d\x9a\x93\x29\x7a\x2d\x1e\x9a\x42\x80\xf3\x40\x92\x40\x0c\x0f\xe9\xcc\xb8\x48\x43\x33\x7e\x7a\x5e\xa6\xa8\xe6\xe4\x07\xc7\x87\xe2\xe4\x7b\x59\xda\xfd\xe4\xe6\x4c\x4d\xe4\xec\x03\xce\x4b\xab\x93\x1f\xb9\x8c\xcf\x5f\x13\x45\xc2\x4a\x40\xfa\xed\x88\xf6\x32\x68\x58\xea\x0e\xcc\xb2\x1b\x30\x0d\xc4\x04\xbc\x8f\x00\x37\x54\x45\x39\xad\xb2\x23\xe5\x61\x1b\xe2\x69\x67\x80\x8b\x1e\x1d\xe7\x74\x98\xec\x68\x87\x5b\x3c\x6b\xa6\xa5\x98\x64\x1d\xc8\x96\x63\xf1\x0d\x20\xdb\x52\x53\x8a\x72\xd8\x69\x3f\x0d\xe4\xd3\xe2\x69\xa6\x30\xdc\xcc\x22\x56\xa8\xd1\x5a\x8a\x5e\xe5\xdc\xa4\xdb\xc2\x3a\xae\x47\x82\x02\xf6\x5a\x52\x64\x84\x19\x8a\xfa\xd1\x34\x0d\xf8\xb3\x07\x33\xaa\xc4\x1c\x76\xd3\xbc\x2a\x87\xa7\x87\xc3\xbb\x9a\x9b\x31\xff\xb8\xcf\x09\xe0\x92\xac\xcf\xe7\x46\x03\x33\xe4\xab\xcf\xa5\x6a\xaa\x9b\x87\x1d\x64\x1d\x83\x3f\x06\x7c\x03\xc4\x73\xee\x3f\x90\x4e\xbe\x23\x99\x99\x61\x95\xa0\xbf\x32\x33\xd5\xa7\xcf\x0e\xb6\x38\x2f\x4d\xd4\x87\x9b\x19\x1d\x2c\x49\x72\x7b\x92\x98\xd2\xf4\xb8\x48\x7d\xc9\x2a\x3d\x7d\x4d\x0e\x39\x6c\xb1\x5a\x17\x59\x0f\x86\xdd\xde\x97\x60\x66\x9c\x46\x4a\xf0\x56\x95\x30\x46\xc7\xa3\xc8\xdf\xd4\x94\x8e\x22\xcd\xd9\x1f\x97\xad\x4e\x1e\xd9\xd9\xc6\xf0\xbc\x24\x81\x34\x81\x6c\x01\x4e\x20\x20\x19\xf1\x6c\xef\xc1\xcf\x8c\x86\x7b\x48\x86\xa9\x3c\x78\x43\x0c\x0e\x3f\x3d\x6f\x75\x73\xb0\xbe\xc1\xca\x97\xc8\xce\x35\x25\x3d\x9f\x17\x9f\x6d\x24\x32\xd3\xa5\xac\xe7\xc7\xcd\x14\x9e\xf3\x49\x7a\x29\xab\xcf\x90\xfa\xe0\x4b\x5d\xbb\x72\x87\x86\x9d\xef\x83\xe3\x53\x37\x66\xc1\x11\x14\x70\xc1\x31\xd1\x84\x07\x33\xac\x0b\xc9\xbe\x1c\x54\xc3\xc8\x0e\xbe\xd2\xe4\x06\xb2\xcb\x1f\xe5\x8c\x7b\xe2\x7c\x8e\x75\x8c\xca\x84\xac\x1c\xea\x9a\x6e\x57\x57\xe8\x30\xd5\x14\x87\x7e\x7f\x32\xac\xdc\xb9\x49\x4b\xe2\x89\xdb\xc2\x41\xf9\x48\xbe\xf2\x90\xc6\x46\xc3\x73\xf6\x7f\x07\x25\x8d\xcf\x7f\x2e\xd2\xdb\x5b\x64\xb5\xd6\xe9\x75\xc5\x8d\x58\x11\x83\x33\x21\x72\x3c\xb7\x90\x0f\x80\xfb\xd2\x26\xe2\xee\xe9\xab\x32\x35\x72\xaa\x0f\x9e\xde\x74\xf8\x76\x34\x28\x5e\x59\xad\x03\x37\xeb\xb9\x28\xce\x2a\xe2\xff\xf0\x12\x28\x4b\x7b\x28\xa6\x1a\xef\x56\x6b\x9b\x15\xb3\x52\xd6\x73\xf1\xbc\x3b\x20\x7b\x7e\x02\x68\xeb\x4a\xb0\x1a\xb7\x92\x3c\x5f\x17\x6a\x56\x1d\x51\xbd\xa1\x69\x66\xd4\x45\xbe\x5c\x49\x3a\x38\xd5\xe7\x6e\xb5\xae\xf8\x42\xcc\x0a\xb5\x81\xd1\x95\xf8\xa7\x40\x1c\xef\xc4\x0d\x20\x58\x15\xe5\x67\xb8\x7f\xca\x14\x7d\xda\x74\x37\xa8\xf5\x79\x6c\xb2\x5a\x71\x7b\xfd\x8b\x1b\xf6\xb3\xf8\xf3\x79\x66\xaa\x8b\xcf\xa1\xd2\xe9\xc3\xbe\x7a\x49\xb5\xc6\xc5\xd4\x98\xf1\x2e\x3b\x39\x42\x2c\x71\xc8\xca\x75\x37\xa6\xb3\x71\x14\x81\x6c\x4e\x92\x93\xbc\x9d\x42\x41\x2e\xaa\xaf\x86\x52\x17\x69\x44\xbd\x1e\xb3\x4f\x69\x56\xfe\x82\x68\x2d\x25\xc6\x93\x07\xf8\x46\xf8\xe9\xa8\x08\x1c\x69\xbd\xc0\xd3\xba\x7d\x4e\xe9\xa7\x32\x9b\xd2\x91\x5e\xa6\xa9\x79\x3a\x4a\xcd\x6e\xcc\x06\xe0\x9c\x63\xf8\xca\x1d\x44\xec\xc5\xa1\xa7\x71\x28\x4d\x57\xb1\x62\x85\x57\xa9\xf8\x6d\x47\xac\x1f\x49\xbf\x4f\x0a\xb5\xaf\x50\x73\x6e\xa4\xa6\xac\xf7\x7f\x61\x38\xcd\x6a\xd6\xba\xea\x3a\x62\xe5\x3a\x3d\xbb\x96\x48\xfa\x1d\x88\xe1\x50\x4a\xaf\x6f\x22\x93\x22\xf5\xd5\x4f\x9a\x22\x55\x5f\x33\xb1\xa3\x1e\xdf\xeb\x58\xa3\x44\xcd\xb4\x9d\x4b\xba\x51\x7b\xdc\x4b\xc6\xab\x9d\x6a\x95\x9f\x04\x13\x59\xc4\x57\xea\xaa\xca\x4f\x1a\x51\x6f\x6f\x69\x33\x8b\xd8\x7a\xd9\x2b\xaa\x18\xa4\x80\x0f\xc4\xd0\xe7\x99\xd4\xab\x6b\x74\x6c\x53\xb2\x52\x55\x62\x2a\x6c\x4a\x13\x4c\x06\xe9\x1a\x2b\xa7\x90\xd5\xdd\xb1\x7b\x59\x37\xa6\x29\xc6\xd2\x99\x83\x99\x54\x79\x54\x44\xa5\xe6\x53\x6d\xa0\xac\x75\x3c\xe8\x75\x44\x11\x12\x83\x25\x0e\x25\x98\x15\xd1\x44\xed\xd1\x77\x00\x13\x69\x84\x12\x9b\x88\x4c\x50\x25\x69\x1e\xf0\x11\x10\x6f\x8c\xd2\x44\x93\xdd\x0c\x1f\xf7\xe2\x80\x79\x9c\x93\xc0\x56\x1f\xff\x54\xd2\xf5\x33\xaa\x68\xa2\x1b\xa4\xd2\x45\x44\x33\xf5\x47\xb7\x8c\xb9\x24\x37\xfc\x23\xcf\x4d\xf8\x7f\xeb\xa7\x21\x13\x34\x1c\x0c\x86\xc7\x9d\x2f\x9b\x4e\xca\x0a\xd5\x44\x28\x75\x91\xd2\xc6\x27\x3e\x9d\x06\x7b\xd7\xd4\xa8\x7a\x49\x89\x39\x1e\x4d\x24\x23\x1e\x9f\xf6\xc0\x8e\xa3\xe9\x13\x25\x7d\x4d\xdc\x63\x27\x1b\xac\x27\x2e\xf1\xd8\xdd\x89\x9b\x35\x29\xe2\xd3\x01\x10\x53\x63\xfc\xb1\xd1\x94\xf2\x13\xfa\xa4\x37\x79\x51\xb5\x95\xd2\xb3\x0b\x5a\xde\x45\xec\x38\x2e\x59\x42\x0d\xf1\xf9\x28\xc6\x13\x21\x42\x5b\xef\x1e\x1a\x1b\x65\x68\xdf\x89\x4f\xe6\xa3\xba\x8d\xb9\xbd\x61\xf5\x8f\x10\x35\x1b\xf6\xb7\x5d\xc9\x55\x89\x09\xf8\x01\x89\x61\xea\xbd\x80\x08\x07\x83\xa1\xf1\x15\x99\xec\x62\xd4\x97\x41\x69\x06\xd3\xb0\x58\xe7\x99\xc4\x36\xab\x5e\x13\xe5\x16\x92\x9e\xc4\x92\x8a\xe8\xfc\xbe\x14\xd3\x71\x27\x5f\xfb\xb2\xc2\x9e\x34\x9b\x93\xd5\xcd\xf1\xbe\x67\x35\xcf\x1b\x90\xfc\xa3\x37\x20\x38\x05\xf8\x42\x7d\x85\xeb\xeb\x56\x54\x1d\xf2\x35\x8a\x5a\xad\x2b\x3e\xea\x87\x9d\x82\xd5\x8e\xcb\x37\x12\x67\xce\xf3\xe9\x79\x2d\x95\x33\xa6\x1a\x4b\xbf\xb4\xe0\x69\x5e\x84\xd0\xdc\x8e\xd5\xc9\xa3\xc1\xce\x16\x2b\x7b\x26\xfd\xba\x90\x33\x81\x44\x84\xae\x01\xf6\xde\x35\xc7\xc6\x81\xc4\xaf\xaa\xd2\x46\xd2\x06\x12\xab\x95\x16\xc6\xe4\xc5\x74\x02\x49\xc1\x79\xa3\xd8\xd7\x90\x46\x7d\xdf\xe5\x49\xd7\x51\xf5\xde\x44\x73\x5b\xe2\xc5\xe7\x3f\x3d\xab\x1d\x47\x40\x6e\x20\x5e\x07\xd2\x86\x78\x12\xc8\x29\xc4\x1f\xdc\xe4\x29\xf1\x3f\x54\x27\x97\x27\x3c\x52\x9b\x62\x64\x95\xf8\x79\xe2\xb8\xf8\x1d\xc9\xfd\x7e\xcc\xef\x5e\xed\xff\x0e\xcd\x6a\x65\xc5\xdf\x52\xcb\xd1\x96\xdb\xb8\x38\x78\x75\x7c\xbb\xeb\x16\x3b\x1c\x0c\x06\x47\x01\xd0\x57\xac\x56\xe5\x00\x67\x9d\xfa\x37\x00\x00\xff\xff\x99\x59\x31\x54\x0f\x1f\x00\x00")

func dataMembersTomlBytes() ([]byte, error) {
	return bindataRead(
		_dataMembersToml,
		"data/members.toml",
	)
}

func dataMembersToml() (*asset, error) {
	bytes, err := dataMembersTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/members.toml", size: 7951, mode: os.FileMode(420), modTime: time.Unix(1488114913, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSongsToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x92\xcd\x8e\xd2\x50\x18\x86\xf7\xbd\x8a\x66\x16\xae\x1d\x7f\x46\x5d\xcc\xc6\x44\x6f\xc0\xe5\x64\x16\x84\x34\xc6\x44\x20\x81\xb8\x60\xe7\x77\xbe\x46\x85\x16\x5a\x48\x0c\x31\xa0\x98\x80\x22\x3f\x16\xb4\x01\x09\x16\xbd\x98\x97\xfe\xb0\xe2\x16\x0c\xcd\xa9\x6c\xce\x6c\x4e\xce\xe2\x79\x4e\xde\xf3\xbd\xdf\xd5\x55\xa5\x54\x7c\x5e\xb9\xbe\xd6\x8a\xb9\x82\xa1\x5f\xea\x67\x20\x17\xc2\x92\xa7\x98\x82\x03\xf0\x1b\xb0\x7f\xa6\x15\x4b\xfa\xa5\x7e\xae\x69\x2a\xa7\x0e\x32\x41\x23\x88\x5f\xe0\x01\xd8\x07\x77\xc0\x81\x74\xee\xa8\x9c\x64\xb9\x80\xb0\x0f\x01\x3d\x7e\x91\xaf\xe6\x5f\x1a\x92\xbd\xab\x62\xc3\x77\xab\xa8\xd7\x00\x79\xe0\x4f\xe0\x06\xc4\xf7\x53\xa0\x7b\x4a\xc1\xed\x82\xbc\xd0\x6d\x80\xe6\xe1\xda\x8e\x7a\x5d\x49\xdf\x57\xc6\x17\xb3\xf4\x9b\x53\x88\xcd\xf1\x4c\xef\x52\xb8\x50\x0a\xec\x82\x67\x60\x86\xf8\x2b\xb9\x07\x2a\x2e\x5a\x74\x40\x5f\xd2\xd1\xf4\x21\x6a\x71\xed\x6d\xc4\x3e\xc8\x94\xce\x43\x65\xf4\xa1\x03\xf2\x9e\x96\x0d\xe3\xd6\x93\x5c\xa5\x2a\xd1\x47\x2a\x74\xb7\x7d\x1f\x6e\xbe\xc6\x5d\x0f\xe4\xed\x3f\xb6\xe3\xf1\x06\x64\x1d\x82\x7e\x56\xd5\x6d\xe5\xfb\xad\x2d\x68\x1e\xff\x69\x82\x3a\xa7\x28\xe7\xca\x5e\xc3\xe1\x64\xff\x61\x0b\x6e\xa6\x5f\x9d\x67\xac\xb2\xcf\xdd\xef\x3a\x5e\x53\x32\xfe\x91\xbe\xdb\x07\x99\xc9\x64\x01\xb2\x40\xf6\x71\x00\xc2\xca\x6c\x65\xc3\x60\x27\x9d\xff\x1a\x62\x72\x5c\x1f\xb2\xc3\xf6\x4f\x90\xb3\xff\xcc\x99\xa7\x2c\x3a\x19\xac\x93\x95\x0f\x1a\x3d\x7b\x55\x28\x18\xe5\x8c\xbd\xa1\xe6\x25\xf8\x1b\xb8\x0e\x1e\x83\xbc\xc8\x74\xc2\xd6\xff\xb5\xbe\xd0\xfe\x05\x00\x00\xff\xff\x32\x71\x29\xc3\x0b\x03\x00\x00")

func dataSongsTomlBytes() ([]byte, error) {
	return bindataRead(
		_dataSongsToml,
		"data/songs.toml",
	)
}

func dataSongsToml() (*asset, error) {
	bytes, err := dataSongsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/songs.toml", size: 779, mode: os.FileMode(420), modTime: time.Unix(1488116701, 0)}
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
	"data/members.toml": dataMembersToml,
	"data/songs.toml":   dataSongsToml,
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
	"data": &bintree{nil, map[string]*bintree{
		"members.toml": &bintree{dataMembersToml, map[string]*bintree{}},
		"songs.toml":   &bintree{dataSongsToml, map[string]*bintree{}},
	}},
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
