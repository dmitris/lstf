// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// CREDITS
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

var _credits = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7c\xff\x72\xdb\x38\x92\xff\xff\xaa\xd2\x3b\xf4\xd7\x55\xdf\x3a\xfb\x8a\x91\x1d\xff\x88\x9d\x99\x9d\xad\x55\x2c\x26\xe1\xad\x23\xf9\x24\x79\xbc\xae\xad\xad\x2b\x90\x04\x25\x6c\x48\x82\x0b\x90\x76\x74\x7f\xcd\x83\xec\x3e\xc0\xbd\xc6\x3d\xca\x3c\xc9\x55\x37\x00\xfe\x90\x64\xc7\x71\x32\xbb\x33\xa9\xe4\x8f\x19\x4b\x22\x80\xee\x46\xf7\xa7\x7f\x01\xfc\xe1\x17\xfb\xd7\xef\xfd\x00\xa9\x2e\x13\x48\x45\xc4\x73\xcd\x63\xa8\xf2\x98\xab\xef\xe0\x87\x7e\x6f\x59\x96\x85\xfe\x6e\x7f\x7f\x21\xca\x65\x15\x0e\x22\x99\xed\xaf\xaa\xea\xbd\xd8\xc7\x01\xfb\xfd\x5e\xbf\x77\x2e\x8b\x95\x12\x8b\x65\x09\xbb\xd1\x1e\x1c\x1e\x3c\x3f\x83\xd5\x7f\xe1\x23\xf8\xe3\xbb\x60\x0e\x17\x66\x56\xfc\x78\xc9\x55\x26\xb4\x16\x32\x07\xa1\x61\xc9\x15\x0f\x57\xb0\x50\x2c\x2f\x79\xec\x41\xa2\x38\x07\x99\x40\xb4\x64\x6a\xc1\x3d\x28\x25\xb0\x7c\x05\x05\x57\x5a\xe6\x20\xc3\x92\x89\x5c\xe4\x8b\x7e\x8f\x41\x24\x8b\x15\x3e\x5a\x2e\x85\x06\x2d\x93\xf2\x8e\x29\x0e\x2c\x8f\x81\x69\x2d\x23\xc1\x4a\x1e\x43\x2c\xa3\x2a\xe3\x79\xc9\x4a\x5c\x30\x11\x29\xd7\xb0\x5b\x2e\x79\xbf\xb7\x33\xb3\x43\x76\xf6\x68\x99\x98\xb3\x14\x44\x0e\xe5\x92\x83\xfb\x09\xee\x44\xb9\x94\x55\x09\x8a\xeb\x52\x89\x08\x27\xf1\x40\xe4\x51\x5a\xc5\x44\x85\xfb\x3d\x15\x99\xb0\x6b\xe0\x78\x92\x85\xc6\x59\x2b\xcd\x3d\xa2\xd4\x83\x4c\xc6\x22\xc1\xff\x73\xe2\xac\xa8\xc2\x54\xe8\xa5\xd7\xef\xc5\x02\x27\x0f\xab\x92\x7b\xa0\xf1\x5b\x92\x95\x87\xac\xec\x4b\x05\x9a\xa7\x29\x4e\x21\xb8\x36\xec\x36\xf4\xd1\x33\x50\xca\x7e\xaf\x40\xa9\x96\x56\x4e\xb4\xf2\xdd\x52\x66\x5d\x66\x84\x86\xa4\x52\xb9\xd0\x4b\x1e\x13\xc7\x12\xb4\xa4\x25\xff\xca\xa3\x92\xa6\xc1\xe7\x13\x99\xa6\xf2\x4e\xe4\x0b\x88\x64\x1e\x0b\x64\x4a\x7f\x87\x3b\x37\x5f\x72\x60\xa1\xbc\xe5\xc4\x90\xd9\xef\x5c\x96\x22\x32\x52\xa7\x7d\x28\x9a\xdd\xb5\x3f\xe9\x25\x4b\x53\x08\x79\xbf\x67\xe4\xc6\x63\x94\x32\x6b\xf1\xa4\x90\x04\x5d\xb2\xbc\x14\x2c\x85\x42\x2a\x5a\x72\x9d\xd7\x01\x91\xf0\xd6\x87\xd9\xe4\xf5\xfc\x7a\x38\xf5\x21\x98\xc1\xe5\x74\xf2\x63\x30\xf2\x47\xb0\x33\x9c\x41\x30\xdb\xf1\xe0\x3a\x98\xbf\x9d\x5c\xcd\xe1\x7a\x38\x9d\x0e\xc7\xf3\x1b\x98\xbc\x86\xe1\xf8\x06\xfe\x18\x8c\x47\x5e\xbf\xe7\xff\xe9\x72\xea\xcf\x66\x30\x99\x42\xf0\xee\xf2\x22\xf0\x47\x1e\x04\xe3\xf3\x8b\xab\x51\x30\x7e\x03\xaf\xae\xe6\x30\x9e\xcc\xe1\x22\x78\x17\xcc\xfd\x11\xcc\x27\x80\x2b\xda\xb9\x02\x7f\x06\x93\xd7\xfd\xde\x3b\x7f\x7a\xfe\x76\x38\x9e\x0f\x5f\x05\x17\xc1\xfc\xc6\x83\xd7\xc1\x7c\x8c\x93\xbe\x9e\x4c\x61\x08\x97\xc3\xe9\x3c\x38\xbf\xba\x18\x4e\xe1\xf2\x6a\x7a\x39\x99\xf9\x30\x1c\x8f\xfa\xbd\xf1\x64\x1c\x8c\x5f\x4f\x83\xf1\x1b\xff\x9d\x3f\x9e\x0f\x20\x18\xc3\x78\x02\xfe\x8f\xfe\x78\x0e\xb3\xb7\xc3\x8b\x0b\x5a\x6c\x78\x35\x7f\x3b\x99\x12\x85\xe7\x93\xcb\x9b\x69\xf0\xe6\xed\x1c\xde\x4e\x2e\x46\xfe\x74\x06\xaf\xfc\x7e\xef\x22\x18\xbe\xba\xf0\xcd\x62\xe3\x1b\x38\xbf\x18\x06\xef\x3c\x18\x0d\xdf\x0d\xdf\xf8\x34\x6c\x32\x7f\xeb\x4f\x01\x1f\xb3\xf4\x5d\xbf\xf5\xe9\xab\x60\x0c\xc3\x31\x0c\xcf\xe7\xc1\x64\xdc\xef\x4d\x5e\xc3\xf9\x64\x3c\x9f\x0e\xcf\xe7\x1e\xcc\x27\xd3\x79\x3d\xf6\x3a\x98\xf9\x1e\x0c\xa7\xc1\x0c\x85\xf2\x7a\x3a\x79\xe7\x01\xca\x74\xf2\x9a\xe4\x36\xc6\x71\x63\xdf\x4e\x83\x02\x87\xce\xbe\x4c\xa6\xf4\xf9\x6a\xe6\x37\xd4\x8c\xfc\xe1\x45\x30\x7e\x33\xc3\xd1\xed\x87\x69\x57\x7f\x80\x8b\xe0\xdc\x1f\xcf\x7c\x38\x89\xa2\xe4\xe4\xec\xc5\xe9\xc9\x31\x3b\x3d\x0d\x0f\xc3\x93\x97\x61\x72\xfa\xfc\x34\x3e\x0d\xcf\x4e\xcf\x4e\xa2\xb3\x7e\xef\x97\x84\x40\xa4\xe4\x8d\x24\x7c\x00\xd4\xc7\x98\xa9\x18\x52\x11\x2a\xa6\x56\x7b\x0f\x22\xa3\x4c\x59\xbe\x18\x48\xb5\xd8\x0a\x87\x07\x2f\x01\x6d\xe7\x8d\x84\x61\x55\x2e\xa5\xd2\x03\x18\xa6\xa9\x43\x09\xc5\x35\x57\xb7\x3c\x26\x51\x4c\x79\x8d\x06\x68\x43\x68\x57\x95\xe6\x68\x31\x5a\x56\xca\x5a\x5a\x28\x72\xa6\x56\x90\x48\x95\x69\x8f\x30\x0a\x8d\xc8\x62\x51\xbf\x47\x40\x23\x22\x66\xd0\x0a\x2d\xdf\xa0\x03\xe2\x61\xa1\xe4\xad\x40\x1b\x2c\x97\xac\x84\xfb\xcc\x1d\x07\xf5\x7b\x19\x2f\xc9\xee\x01\xe0\xdf\xa1\x4b\x18\x59\xa7\xa5\x28\x92\x31\x87\xac\xd2\x88\x92\x08\xcf\x34\xeb\x1a\x4e\xf4\x7b\x06\x0d\x3c\x83\x12\xa9\xd0\x25\xa1\x7c\x6b\x45\x42\x90\x36\x39\xb1\xd0\x51\xca\x44\xc6\xd5\xe0\x1e\x1a\x44\xde\x16\x85\xa3\xa1\x50\x32\xae\x22\xde\x90\xd1\xef\xad\xe3\xd5\xd3\xc8\x40\x04\xa3\x1f\xbb\x4e\xc5\xc2\xb4\x2c\x97\x5c\x41\xc6\x4a\xae\x04\x4b\x75\x23\x69\xda\x1f\x72\x39\x6d\xea\x1d\x4f\x63\x2e\x68\x20\xce\x9b\xb3\x8c\x9c\xdf\x1b\x29\x17\x29\x87\x20\x8f\x06\x90\xcb\xe6\x37\x92\xba\x28\x35\x32\x94\x9b\x99\xa4\xd2\x90\xb1\x15\x84\x1c\xf5\x84\x60\x9d\xe7\xb1\x54\x9a\xa3\x4a\x14\x4a\x66\xb2\xe4\x60\x44\x52\x6a\x88\xb9\x12\xb7\x3c\x86\x44\xc9\x0c\xd1\xbe\xed\x39\x9d\x2f\xd3\x05\x8f\x50\x7f\xa0\x50\x02\xd5\x4a\xa1\xe6\xe4\x2d\x64\xb7\x48\x1c\xcc\xb6\x43\xf1\xab\x1b\xb2\xf0\x4d\xf0\x1a\x8e\x47\x06\x71\x82\x57\x57\xf3\xc9\x74\xd6\xef\x59\xd4\xa6\x5f\x10\xcb\x36\xd1\xb9\x85\xbd\x2d\xa0\xf6\x1c\x52\x23\x1a\x3a\xa8\xf6\x68\xd9\xcd\x71\x08\x5b\x6b\x90\x4d\x0b\xb6\x50\xbb\xdf\xdb\x0e\xdb\x53\x1f\x46\xc1\x8c\x10\xd6\x1f\xdd\x07\xd8\x35\xa3\xfd\xde\xe4\x7a\xec\x4f\x0d\x72\x37\x6c\xc2\x2b\x1f\xd6\x30\x7b\x14\x4c\x7d\x44\xdd\x60\xdc\xfc\x75\x1e\x8c\xfc\xf1\x7c\x78\xe1\xf5\x7b\xb3\x4b\xff\x3c\x18\x5e\x78\xe0\xff\xc9\x7f\x77\x79\x31\x9c\xde\x78\x76\xd2\x99\xff\x9f\x57\xfe\x78\x1e\x0c\x2f\x6a\xc0\xdf\xfd\x98\x54\x2e\xa7\x93\xf3\xab\x29\xf9\x1c\x14\xc5\xec\xea\xd5\x6c\x1e\xcc\xaf\xe6\x3e\xbc\x99\x4c\x46\x24\xec\x99\x3f\xfd\x31\x38\xf7\x67\xdf\xc3\xc5\x64\x46\x02\xbb\x9a\xf9\x5e\xbf\x37\x1a\xce\x87\xb4\xf4\xe5\x74\xf2\x3a\x98\xcf\xbe\xc7\xbf\x5f\x5d\xcd\x02\x12\x5c\x30\x9e\xfb\xd3\xe9\xd5\x25\xfa\x81\x3d\x78\x3b\xb9\xf6\x7f\xf4\xa7\x70\x3e\xbc\x9a\xf9\x23\x92\xf0\x04\xfd\xcd\x0d\x79\xed\xc9\x94\x3c\xf1\x76\xa7\xd4\xb8\xa1\xd9\x7c\x1a\x9c\xcf\xdb\x8f\xa1\x33\x99\x4c\xe7\xfd\x5e\xc3\x27\x8c\xfd\x37\x17\xc1\x1b\x7f\x7c\xee\x77\x5c\xd6\x5e\xed\xb2\xc8\xd1\xdd\xc0\xf5\xf0\xc6\xf9\x2d\xeb\x90\xc8\xeb\x75\x54\xd7\xa3\xfd\x84\xe0\x35\x0c\x47\x3f\x06\x48\xb9\x7d\xfa\x72\x32\x9b\x05\x56\x5d\x48\x6c\xe7\x6f\xad\xcc\x8d\xcf\xfa\x45\xfd\xd0\x2f\x3b\xff\x0f\x70\x4b\x20\xd1\x8e\xe8\x67\x25\x8b\xde\xfb\x1f\xa2\x25\xcb\x17\x7c\xff\x2e\x13\x8f\xcc\x04\x36\xc6\xed\xbb\x50\xb1\x15\xf7\xc3\xee\xbb\x60\xbe\xb7\x35\x53\x38\x02\x9a\x01\xdc\x14\x5f\x2c\x45\x80\x3a\x43\x58\x07\xba\x47\xa6\x08\x70\x4f\x86\x60\xa2\xe4\x47\xa6\x08\xf0\xb1\x0c\xa1\xdf\xfb\x78\x8a\x00\x9f\x94\x21\x74\xe9\x73\x29\x02\x3c\x3d\x43\xe8\xf7\x9a\x14\xe1\xde\x90\xe1\x0b\x64\x08\xb0\x96\x20\x90\xdf\xfe\x27\x67\x08\x2d\x17\xd4\xef\x3d\x31\x43\x58\xf7\x36\x75\x82\xd0\xef\x3d\x94\x21\xc0\x27\x27\x08\xfd\xde\xb6\x0c\x61\xdd\xd9\x3c\x3e\x41\xe8\xf7\xda\x19\x02\x3c\x3d\x41\x40\xc2\x5c\x86\x00\x5f\x24\x41\x78\x04\x60\xed\xbb\x24\xe2\x94\xbf\x38\x0e\x8f\x0e\xd9\xcb\x23\x1e\x26\xcf\xe3\xc3\xb3\x83\xd3\xe7\x61\x12\x1f\x1d\x1c\x1e\xbc\x3c\x7b\x7e\x72\xf8\xf5\x81\x37\x4f\x99\x2e\x45\xb4\xbf\x90\x5a\x2c\x98\x7a\x24\x74\xaf\x8d\xda\xb7\xb1\xfe\x03\xff\x86\x05\x8b\x96\xbc\xa9\xe6\xdc\xf7\xdc\x8f\x5c\x91\x81\x1f\x0e\x0e\x3c\xf8\x0f\x96\x57\x18\xa3\x1f\x1e\x1c\x1c\xdf\x33\x04\x89\xfb\x6e\x7f\xff\xee\xee\x6e\xc0\x68\x09\x4a\xa5\x2c\x17\xda\x78\x14\x7f\xfa\xae\x0e\x21\x47\x01\x6a\x96\x49\xb8\x31\x52\x81\xa9\x7f\x39\x9d\x8c\xae\x48\xe1\x3c\x7a\x6a\x14\xcc\x4c\x04\x46\x49\x6a\xbf\xf7\x7c\x00\x23\x9e\x88\xdc\xa0\xd5\xc0\xf2\xba\x63\x59\xd9\xb1\x20\x94\x71\x66\xf0\xb9\xe4\x2a\x33\xd9\x40\x2b\x39\x48\xa4\x32\x05\x1d\x97\x64\x10\xd8\xd3\x44\xf8\x64\x37\x73\xc3\x60\x3b\x11\x39\x8f\x21\x5c\xc1\x8c\x47\x66\x8a\xe7\x50\x2e\x95\xac\x16\x4b\x78\x59\xd7\xae\x9c\x03\x5a\x23\x4a\xaa\x0d\xaa\x1a\x54\x95\x77\x39\x57\x08\x8c\x3c\x2f\x45\xb9\x02\x46\x89\xa5\xf8\x6f\x5a\x8e\x66\xd9\xf6\x3c\x25\x7d\x42\x1b\x5f\x8a\xf8\x5d\x36\xbb\xd9\xac\xce\x17\x2c\x05\x9f\xe6\xdd\xa0\xa0\xca\x91\x39\x0b\xbc\x2c\xa2\x49\x1c\x09\xe8\x56\x11\xb5\x01\x6c\x32\x44\x3f\x20\x80\xd3\xba\x94\xb3\xc8\xd4\x64\xa4\xf6\x43\x4a\xf4\x7a\xc8\x08\x7e\x4b\xea\x0a\x91\xcc\x32\x99\xd3\x3c\xf6\x31\x97\x43\xb1\xd2\x2e\x36\x80\xd7\x36\x2f\x2a\x2a\x55\x48\xed\x4a\x63\xc2\x4a\x5d\x34\x3b\xb3\x63\xe7\xd8\x21\x26\x34\xec\x8a\x3d\x33\x50\xde\x71\x85\xae\x55\xa1\x67\x93\x0a\x44\x6e\xfe\x26\x5f\x1f\x31\x4c\xbb\x29\x6b\x03\xb0\x0f\x11\xe7\x98\xe2\xe5\x6c\xc1\x71\xc3\x28\x09\xae\xa2\xa5\x25\xca\x83\xbb\x25\x27\xc6\xc3\x95\xa1\x9c\xd1\xcc\x8d\x44\xee\x04\x6a\x8f\x54\xb0\x2b\xc4\x9e\xd9\x13\xbd\x14\x05\xce\x93\x88\xa4\xa4\x20\x26\xc2\x89\x77\x4f\x0e\xfe\xff\x1e\x2d\x26\x15\xb7\xe2\x36\xd3\x54\x25\x55\x26\x50\xee\x7a\xc9\x14\xd7\x6e\x3e\xb1\x07\x21\xcf\x79\x22\x22\xf4\x91\x9d\xb9\x5b\x34\xd6\xbb\x7c\x23\xab\x1d\xd8\x95\x8a\xfe\x52\x3b\x7b\xed\x8d\x66\x39\x09\xe3\x56\xc4\x15\x4e\xa5\xa0\xad\x12\x34\x9c\x7f\xe0\x2a\x12\x1a\xa9\x68\xfc\xb9\x76\x31\x1a\xf2\x4f\x7b\xb1\xae\x5a\x33\xaa\x18\xec\x98\x74\x7d\x4d\xb3\x0a\xc5\x13\xae\x14\x26\xa7\xf8\x6b\x42\x92\x7e\x8f\x2b\xb4\xeb\x19\xda\x6c\x6a\x13\x5f\x85\x15\x85\x18\x26\xbe\x32\x51\x4b\x1d\xe8\xb5\x0a\x14\x5e\x37\xcc\xa3\x49\xcc\xcf\x9e\xb3\xf2\x44\x2c\x2a\xd5\x0a\x02\x6b\xaa\x27\x14\xfc\x6c\x52\x8d\x51\x27\x7d\xa7\xb8\xae\x52\x32\x05\xcc\xab\x21\xe3\xe8\x94\x44\xc4\x8c\x2d\x94\x8a\xe5\x1a\x9f\x63\x4e\x83\xe8\x9b\xd4\x7e\x4c\x80\x81\x91\x0b\x4d\xe6\x75\x79\xa3\x19\xd6\xf8\x8b\x64\x56\x08\xb4\x1c\x69\xa2\x32\xc3\xdf\x82\xe7\x5c\x6d\x06\xb4\x0d\x3c\x45\x32\xbf\x35\xb8\x4c\xd1\x9f\xad\x57\xf0\x58\x30\x28\x57\x45\x8b\xdf\x6b\xa9\xde\x6f\x18\xfe\x9d\x54\xef\x89\x58\x53\xc1\x5a\x8a\xa2\x51\x77\x91\x3b\x0e\xac\xb2\x1b\x89\x59\x7e\x32\x16\x73\x60\xb7\x4c\xa4\x2c\x4c\x9d\x8d\xb7\x80\xc7\x43\xac\x44\x7d\x8b\x98\xd5\x1d\x66\x6d\x7f\x2d\x84\x74\xe8\xd5\x0e\x13\x11\x38\xca\x12\x5d\x46\xec\xa2\x53\xa4\x94\x26\xd8\x65\x39\xf0\x0f\x2c\x2b\x52\x8a\x6c\xeb\x82\x8c\xad\xe2\x0c\x8b\x82\xe7\xb1\xf8\x00\x21\x4f\xe5\xdd\x5e\xcd\xfe\x88\x2b\x71\xcb\x4a\x71\xcb\x01\x25\xa1\x77\xd6\xf7\x1c\x17\xd8\xce\xbc\x65\x9c\xe6\x31\xcc\x3b\x9a\x43\x86\xee\x58\xe6\x64\x72\xed\x2a\x8c\x01\x23\x5c\x88\x36\x09\xd5\xfe\x6e\x29\xa2\x65\x6d\xf1\x3c\x16\xa5\x54\x68\xd3\x8a\xdf\x0a\xda\x3e\x54\xd9\x5c\x96\xd6\x20\x80\xa7\x2c\x94\xca\x7d\x6a\x4a\x51\x6d\xb3\xa1\xa9\xd0\x6d\x71\xcd\xf3\x92\x44\xce\x30\x01\x48\x49\xff\x41\x2a\xb1\x10\x39\x4b\xb7\x6c\xf3\x26\xd4\x1a\x20\x4a\x3a\x26\xee\xc1\xba\xd8\xac\xd4\x50\x77\xed\x86\xd1\xe4\xd6\x19\x28\x9e\x31\x61\xcd\x90\x17\x4c\x91\x6a\xa0\x3c\x88\x81\x8c\x2b\x9e\xae\x20\x15\xf9\x7b\x12\x58\x28\x72\x52\x8c\x9c\x65\x7c\xcf\xed\xb3\xc8\x4b\xae\x12\x16\x11\xf6\x7b\xb5\xcf\xab\x45\xb9\x41\x10\x4a\x85\xcb\xa4\xde\xe8\x73\x57\x44\x13\x32\xdf\xba\xc9\xeb\xfa\xde\x6a\xfc\xd8\xc5\x6a\xb9\x59\xcb\x72\x9e\xb1\x26\x02\xa7\xea\x6c\x04\x69\x6c\x6c\x23\x0a\x33\x8f\x34\x22\xa1\x31\x52\xdd\x4b\xb7\xd7\x32\x80\x12\xd1\x5c\xe6\x2c\x4d\x0d\x1e\xeb\x2a\xb4\x05\xde\x52\x82\x0b\x1f\x48\x99\x88\x66\x93\x4a\xe7\x0d\x69\x04\xd0\x1b\xf1\x81\xd9\x58\x72\x5f\x0f\xba\x80\x76\xb4\x81\x78\x4b\x8b\xa3\x72\x87\x7c\xc9\xd2\x84\x72\xd0\xed\x11\xc8\xe3\xfc\x36\xec\xd4\xfc\xec\xd0\x4c\xc6\x73\xd7\x78\x2b\x13\xe0\x29\x8f\x4a\x25\x73\x11\x79\x28\xfb\x90\xa5\xa4\x38\xae\x58\x89\x21\x44\x95\x5b\x99\x03\xaa\x7c\x23\x6a\xde\x08\x08\xe5\x43\x15\x78\x6b\x17\x24\x75\xed\x3d\xe8\x5e\x2c\x36\xb5\xe7\x97\x79\x8b\x1e\xc8\x98\x48\x71\x68\x2a\x74\xa9\xbd\x4e\x95\xdc\x85\x33\x7a\xa5\x4b\x9e\xe9\x06\x9b\x85\xd6\x15\x47\xbf\x10\x91\xc7\xb3\xbf\x9b\x0d\x47\x4f\x66\x62\x8e\x3a\x56\x6a\x8b\xda\xab\x91\xa2\xb3\xef\x2d\x19\xa3\xbc\x62\xa1\xa3\x4a\x93\xc3\xa6\xf5\x32\x42\x43\x1b\x01\x5e\x13\xa0\x39\x6f\xc3\x3f\x38\xe6\xbb\x5c\x3a\xed\x8b\x64\xae\x0b\x11\x55\xb2\xd2\xe9\x0a\x32\xa6\xde\x23\xb2\xa9\x26\xc2\x31\x21\x13\xd7\x62\x91\x13\xa2\x8b\x9c\xf6\x85\xc4\xb9\x55\xef\x10\x8d\x76\xc6\xb2\x04\x06\x6d\x9b\x1c\xec\x6c\x58\xea\x5a\x4c\x5c\x73\xec\x4c\xed\x23\x61\x4b\x5b\x6e\xa6\xfc\xd1\x5d\x11\x96\x4c\x43\xc8\x79\x0e\x8a\x47\x9c\x30\x3a\x5c\x75\x56\x71\xd6\xa6\xf9\xdf\x2a\x9e\x97\x29\x2e\x19\x49\x55\x48\xe3\x7a\x31\x4e\x6d\xd9\x19\x41\xcd\xe1\x00\xde\x60\x5c\x84\x6b\x36\xe5\x2f\x17\x1a\xc1\xac\x5b\x5b\xd9\x9a\x78\xd4\x06\xd5\x06\x5c\xce\xa2\x25\xb4\x24\xd3\x29\x94\x91\x8b\xbf\x91\x15\x30\x0c\xd0\x0a\x5e\x56\x2c\x35\xea\x76\x27\x55\x1a\xdf\x09\x0c\x19\x72\x99\x3f\xa3\xdd\xd6\xe2\x96\x3e\x3e\x73\x35\x35\x25\x57\x2c\x2d\x57\xcf\x12\xc5\xb9\x07\x42\x29\x7e\x2b\x23\xc4\xe8\x35\xcf\x6c\x53\x34\x5c\xac\x6e\xbc\x78\x18\xca\x15\xa8\xb5\x1b\x48\xe6\x70\x9a\x0a\x5b\x51\xba\x42\xb5\x2c\x52\xb6\xf2\x9a\x6f\x0a\xae\x8c\xe7\x5c\xab\x73\xb5\x6a\x60\xb5\xc2\xd7\x28\x4b\x01\xee\xc6\x6a\x5b\x3c\x33\xe1\x07\xed\xca\x51\x6b\x57\x2e\x19\xe2\xe9\x6f\x7a\x4b\x76\xf9\x87\x88\x17\x25\x5a\x92\x2e\x9d\xd5\x99\xb2\xa7\xc9\x5c\xf6\xa0\x30\x5c\xb6\xb6\x2c\x63\xef\xb9\x07\x4b\x76\xcb\x29\x48\x33\xc4\x50\x82\x2b\x93\x04\x83\x34\x49\xa5\x45\xcf\xfe\x57\x64\x85\x54\xa5\xd9\x8d\xda\xd8\x6d\x78\x6b\x43\x3a\x42\x12\xc3\x13\xb2\x6e\x36\xc6\xad\xc8\x8a\x22\xa5\x72\x5e\x9e\xae\x8c\x6c\x11\x9c\x2c\x59\xd4\x63\xd3\xf6\xd9\x9a\xad\x70\x65\xa6\x68\xcb\xb4\x06\xc5\x9c\x47\x5c\x6b\xa6\x04\x59\x61\xa2\x44\xbe\x70\xb9\x07\x17\xc6\x99\xb5\x8d\x7b\x57\xef\x01\x4b\x65\xce\xad\x8b\x8b\x64\x16\x8a\xbc\x8e\xc2\x69\xd0\xfa\x00\xc3\x8a\x6d\xe0\x19\x7d\xa3\xc2\x29\x06\x68\x5d\xc2\xec\x02\x77\xb8\x01\xce\x79\x0d\x20\x48\x70\xc7\x6d\xca\xa2\x4b\x51\xa2\xfa\xd6\x1b\x51\x8a\x85\x6d\x21\x2e\x18\xfe\x4c\x10\x66\xb3\xe9\xdd\xc6\x07\xd9\x70\x58\x49\xad\x9f\x91\x98\x90\x81\x48\x56\x18\xff\x98\xcf\x22\x07\x06\x29\xbb\xd3\x95\x28\x91\xc9\x94\x2f\x0c\xb6\xdb\x1e\xef\xb5\x0b\x89\x11\xc4\xba\x88\xf7\x10\x7c\x11\xd4\x1b\xa2\xb5\xcd\x81\xdd\x2c\xad\x06\xe4\xca\x31\xe4\xf6\x20\xa3\x08\xb3\x5c\x72\x13\x49\x75\xf5\xce\xc4\x3c\x2e\x51\xb4\x16\xe1\x72\x82\xc6\x96\xac\x0f\x73\x61\x91\xc1\x7c\x34\x44\xdc\x31\xa3\x1b\xac\x2e\x00\xc7\xac\xac\x55\xad\x96\xa9\xd0\x94\xc6\x99\xb6\xfa\xf1\x60\xad\x79\x3c\xa0\x75\x33\xb6\x6a\x35\x8c\xd7\x50\xa6\x73\xe4\xa6\xc1\x9b\x07\x22\x34\xda\x07\x0c\xf8\x78\x2c\xaa\x6c\x4b\x67\x1e\x43\x99\x4e\x32\x6b\x9c\xf1\x3d\x38\xe5\xad\x75\xeb\x9d\x26\x65\x9c\xdf\xdf\xba\x77\x1d\xfb\x5d\xb6\x67\x58\xac\x74\x09\x0b\x24\x15\x29\x33\x99\x81\xe2\x91\x28\x04\x47\x48\x6a\xc7\xab\x36\x79\xc3\x7f\x1b\x1c\xae\x1d\xb7\xb2\xbb\xf4\xbd\x71\x89\x66\xc1\xb0\xb5\xa0\x29\x9f\x34\xc1\x2f\x26\x3b\xd4\x58\xa1\xd2\x8a\x42\x9d\x51\x32\x13\x39\x2a\x86\xc9\xee\x74\xbd\x36\xe2\x57\xad\xbd\x38\xa3\x29\xf1\xda\xbe\x3c\xce\xd2\x59\x36\x6a\x2d\x6b\x0e\x20\x78\xcd\x31\xae\x3a\xad\xa6\x50\x3e\x5f\x6d\x30\x56\xaf\x5a\xaf\xd6\xee\xb5\xd8\x53\x4a\xc6\xd9\x79\x56\x93\x3d\x84\xbc\x98\x63\xe0\xe3\xd5\x21\x01\xfe\x63\x65\x63\x56\x96\x29\x53\x10\xd8\x42\x4b\x17\x2c\xa1\x13\x76\x19\x5c\x74\x33\x10\x61\xb1\xa4\x28\xb4\xe0\xca\x9c\xb0\xb0\x2d\x2f\xa6\x4a\xe7\x88\xc0\x06\xdc\xeb\x0c\x76\x64\x15\xef\x21\x22\xd5\x1b\x6e\xf3\x32\xdc\xdd\x9d\xf1\x64\x1e\x9c\xfb\x3b\x50\xf2\x0f\x25\x49\x19\xed\xcb\x2e\x60\xce\x1c\xd8\x45\xda\x56\xd4\xb2\xf2\x2d\x36\xb1\x21\x50\xda\xa4\x7a\x22\x97\x15\x32\x50\x9c\xc5\x94\x00\x36\x3a\xc6\xb7\x4a\x13\x51\x87\x89\x9c\x37\x32\xb7\x88\x45\xc6\x6f\x58\x20\xe2\xbd\xc7\x88\xb3\x9e\x64\xbb\x58\xb7\x8a\x93\x74\x8b\x95\x90\x72\xa6\x31\xe5\x69\x8a\xe0\x76\x40\x63\x93\x45\x8a\xb9\xe9\x77\x8e\x44\xe6\xe8\x6b\x24\xdc\x48\xa6\xa5\x44\xfa\xc1\xf5\xbf\x6f\x63\x74\x47\xa7\x1a\xe3\xed\xd6\x7f\x40\x24\x0d\x8e\xa0\xff\x5b\x34\x0e\x6d\x73\x76\xa9\xbc\x75\xd9\x32\x17\xa5\xb5\x4a\x4c\x36\x8e\xdf\x22\x9d\xa4\x63\x12\x14\x05\xdc\x72\x65\x36\xa8\x5c\x0a\x15\x3f\x43\xf6\x56\xf5\x7e\xe4\x52\x65\x98\xc6\x62\x74\xc0\x99\x1a\xd0\x49\x29\xdc\x67\x84\xa7\x75\xe1\xb6\x76\x98\x22\x00\x93\xe0\xd6\xb5\x35\x96\xb6\xd2\x4a\x0c\x32\xda\xa4\x58\x13\x32\x2d\xd6\x4e\xe1\xbb\x76\x05\x2c\x8e\xf1\x6f\x85\x59\x49\x5b\xff\xea\x39\x1c\xd1\x56\x32\x8f\xd1\x79\xcf\xc8\x5c\x8b\xb8\xa5\x2a\x94\xf1\xb0\x1c\x17\xe4\x79\x5c\x65\x2e\xd2\xec\x68\x88\x43\x0e\x93\x9d\xb9\x2d\xec\x02\x16\x89\xd5\x95\x13\x58\xba\xdd\x68\xa8\x58\x04\x21\x37\xee\x5c\x55\x5d\x6d\x33\x02\xb9\xa7\x19\xb0\x55\x32\x4d\xf8\x4f\xa1\x26\x55\xc2\x8d\x23\x5f\x2b\x3a\xd5\xf2\xc7\x29\x2c\x03\x6d\x6a\xa5\x82\x58\x60\xa4\xd9\x89\x4b\xb7\x44\xdb\xae\xa0\xb6\xa5\xfb\x62\x26\x69\x35\x5e\x64\xb2\x85\x12\xcf\xd9\x47\x42\x69\xdc\xea\x9e\x8c\xa1\x5d\x15\xab\x6d\x86\x66\xc3\x85\xeb\x1a\x5a\xb3\xf8\x46\xdb\xa7\xe3\x4f\xeb\x18\x39\x92\x99\x09\x7d\x51\x71\x5a\xb5\x91\x3a\x9b\x58\x8b\xd9\xdb\xbb\x70\x42\xf9\x88\x6b\x99\x53\x0e\xd9\x84\x6f\x7a\x00\x57\x79\xca\xb5\xa6\x9d\xe2\x1f\x8a\x54\x44\x02\xb3\x52\x9a\xb0\xd5\x74\xb0\x95\x86\xd5\x7a\xf0\xd7\x2a\x24\xb5\x8a\x48\xf7\x16\x8e\x5c\x50\x8e\xab\xad\x17\x53\xea\x36\x7e\x53\xe1\x7d\x7c\xe6\xe4\x8e\x49\x20\x89\x2d\x0d\x31\x13\x98\x78\x33\xae\x5b\x77\x00\x30\x96\x25\x0e\xa9\x7b\x21\xf5\x61\x3f\xcc\x99\xd0\x38\x17\x94\x7b\xa1\x6b\x20\xb2\x74\x55\x70\xa5\x79\xcc\x4d\x53\x05\x55\xbe\xde\x07\xbb\x88\x09\x10\x4c\x35\xb2\xe4\x4d\xd6\xb2\x50\xdc\xe8\xf8\xca\x1a\x03\x25\x4c\xfc\x03\x8f\x6a\xe0\x26\x40\xad\x05\xa1\xf8\x82\x29\xd3\xa1\x59\x4f\x11\x4c\x91\xfd\xc5\x00\xe6\x2e\x84\xd0\x88\x78\xad\xb8\x37\x96\x04\x8a\xa5\x09\x91\xdb\xa7\x25\xcc\x69\x72\x43\x31\x8e\x36\xad\x01\x96\x71\xdd\x8a\x48\x34\x26\x6b\xea\x56\x44\x1c\xec\x47\x73\x0c\x10\x35\xb6\x39\x43\xd8\xde\x37\xcf\x15\x7d\x6c\xfa\xa8\xf8\xdf\x2a\x61\xdb\x30\xe8\x99\xb5\xcc\xc9\x37\xd3\x2e\x56\xba\x94\x19\x53\x2b\x77\x28\x35\xe6\x3a\x52\x22\xb4\x3b\x60\x93\x03\xb1\x10\x9b\x85\x50\x67\x37\x6e\xb3\x2c\xc4\x6f\x41\x76\x12\xd1\xe9\x00\x46\xf5\xc1\x4b\x7c\xe8\x9a\x29\x14\xc8\xaa\x56\xf8\x9a\xce\x70\x65\x32\x4b\x4a\x86\x31\x0b\x72\xb6\x4e\x3b\x47\x39\x46\x53\x81\xf2\x9a\x6d\xb2\x06\xae\x1b\x3a\x77\x91\x50\xcc\xe1\xbb\xb9\x63\xfb\x59\x51\xea\xee\x86\xee\x01\x1d\x01\x75\xc7\x4b\xe0\xd5\x70\x16\xcc\x8c\x4c\xd7\x0e\x9a\x04\xbe\x3d\xd4\x5d\xb7\xaf\x3b\x07\x4f\xec\x39\x50\xfe\xa1\x50\xc8\x9e\xe5\x41\x10\x70\xc4\xad\x92\xa4\xb7\xe5\x48\x91\x67\x4a\xd6\x46\x44\xf6\xc4\xcc\x1a\x7a\xca\x04\xe6\xc1\xfc\xc2\xf7\x60\x3c\x19\x3f\x6b\x9f\x33\xf1\x36\xcf\xab\x48\xd5\x3d\xd3\x4e\x33\x6c\x9e\x5a\x31\x9e\xd3\x74\xda\x52\x9e\x62\x2e\xa5\x0b\x99\x6b\x41\xc5\x7c\x6a\x74\x98\x94\xad\xad\x1f\xac\x28\x94\x2c\x94\xc0\x58\x9a\x18\x4d\xa0\xa2\xba\x24\xa9\x5b\x03\xa6\xad\xda\xa4\x3b\xb1\x55\x65\x94\x53\x18\x1c\x16\x9a\x00\xbb\x3e\xc6\x45\x36\x48\x58\x6d\x9b\x93\x54\xf7\x6c\x77\x27\x37\xb3\x4c\x52\xb5\xb3\x01\x5c\x34\x87\xb3\x64\x02\x17\x82\x85\x22\xa5\x2e\x73\x80\x6e\x14\xf8\x2d\x6a\x2a\x9d\xc5\xa6\x29\x72\x09\x29\x95\x16\xcb\x25\x97\x6a\x55\x17\x3c\x5c\x3f\xa8\x94\xaa\x6c\x27\xf0\x39\x5f\xa4\x62\xc1\xf3\x88\xef\x79\x75\x63\xd8\xeb\xd4\x4c\x6d\xed\xe5\xa3\xaa\xbd\x6b\x3c\xbe\x86\x98\xa7\x22\xa4\x68\x8c\x08\x5b\x28\xa9\xb5\x6d\x06\xb8\xe5\x4a\x60\x51\xa9\xa9\x8d\xbc\xdd\x14\x0c\x34\x76\xbc\x82\x54\x74\x55\x03\x00\x52\x41\x8b\xda\x0c\x9d\x76\x93\x65\x6c\xd1\x2d\x8f\xe3\x58\xd7\x35\x6f\xfa\xe7\x74\x82\xd8\x15\xb7\x44\x1e\x89\x18\xa3\x51\x53\xa1\xc7\x18\xc4\xd4\x4e\x05\x4b\xdd\x94\x0e\x7c\xa3\x25\x43\xd1\x70\x05\x4c\x99\xee\x32\x7a\x64\xeb\x77\x75\x95\x96\xeb\x09\x28\xc9\xb0\xaa\x51\xa4\x32\xdf\x88\xdc\x6e\x60\x0b\x33\x9b\x04\x7e\xf7\xc1\xe6\xb1\xa3\x08\x19\x4e\xa5\xd1\xcf\x85\x94\xf1\x9d\x48\x9b\x6a\xdd\x7b\xd0\xa5\x2c\x0a\xb6\xa0\x93\x7b\x59\x51\x21\xc9\x09\x13\x69\xa5\x8c\x83\x61\x69\x52\xe5\x4d\x80\x42\x5e\x6d\xe3\x78\x44\x24\xb3\x0c\x35\xb5\x2d\x07\xb3\x28\xd7\x7b\x1e\x69\x1d\xc6\xd3\xeb\x05\x30\x9a\xa1\x2e\x54\xb3\xf8\x56\x50\x6b\x31\xb1\xa7\x1a\xb4\x16\x96\x79\xd7\xf9\xb7\x93\x93\xb2\xbf\x1c\xc0\x30\x42\xa8\x47\xf6\x1d\xa6\xe2\xb2\xc3\xc6\xe9\xb6\xf4\xff\x7a\x89\xa1\x76\xd7\x2c\xbb\x9d\xb6\x07\x1b\x56\x2e\x7c\x8c\x96\x52\x9a\x92\x23\x15\x16\x5b\x3d\x69\x2a\x6e\x02\x83\x84\x13\x62\x78\xc0\x88\x3a\x96\x47\xdc\xd0\x5f\x98\x9a\xa3\xc5\xb6\x15\xa9\x19\xcf\x72\x51\x5a\xb3\xab\xdb\x9d\xa9\xa3\x1a\x64\x98\xda\x22\x90\x76\x27\x24\xed\x71\x4b\x54\x3e\xa1\xc9\xef\xd8\x1c\x48\xe8\x56\xeb\x84\x0f\xe0\xad\xbc\xc3\x7c\xc5\xa4\x79\xb5\xa0\x48\x8c\xad\x69\x1b\xce\xe8\x84\x47\x9e\xd6\xed\x85\x3a\x48\xb6\x7d\x06\xaa\x95\xda\xaf\x11\x24\x1b\x88\x24\x5a\x29\x5a\x69\xda\x12\x0e\xa9\x9b\x52\x4d\x6b\xdf\x6d\xe1\x15\xf3\x1a\x91\x18\xe4\x45\xab\x36\x46\x4d\x32\x49\xac\x4c\x62\x9e\xf0\x3c\x36\xcf\x2f\x65\x1a\x6f\xa9\x4b\x33\x95\x11\xd4\xb8\x78\xb8\x96\x9e\xb3\xd9\x4a\xa9\xa6\xdf\x64\x8b\xb3\x4c\x6b\xae\xd0\x4e\x6c\xbd\xd2\xdb\x2c\xce\x86\x2b\x1b\x33\x38\x56\x56\xc8\x79\x23\xc9\x3a\xf6\xbe\x6b\xe9\x5e\x2b\xe0\xab\xe9\x20\x6d\xf5\xc7\xe6\x38\xf6\x96\xb3\x5f\xf8\xf3\xf0\xf2\xd2\x1f\x8f\x82\x3f\x7d\x87\xdb\x46\xd9\x7b\x51\xa4\x2b\xdb\xdc\x6f\x9f\x54\xc3\xdf\x88\x90\x3b\xd7\x94\x01\x80\xf9\x23\x9f\xf7\xec\x09\x83\x6e\x7a\x6f\x22\x61\x29\x52\xae\x8a\x14\x51\xd8\xdd\x2c\xa9\x93\xeb\x44\xf0\x34\xd6\xc0\xf3\x28\x95\xda\x80\x79\xa8\x58\xf4\x9e\x97\x1a\x76\xfe\xfc\x97\x1d\x97\x4e\xa4\x2c\x72\xee\x6b\xe5\xb4\x87\x10\xd3\xa6\x65\xad\x04\x77\x00\xbb\x23\x99\xff\x5b\xdd\x51\xaf\x4d\xd1\x4d\xfc\xff\xf6\x80\x12\x68\xca\x20\xf5\x52\x56\x69\x8c\x11\x79\x4d\x83\xbb\x9d\xd3\xf8\xe0\xba\x9d\x89\x46\xa1\x57\x79\xc9\x3e\xd4\xfd\x43\xca\xb3\xcd\xe2\x03\xb8\xe6\xc0\x52\x2d\x41\x71\xf3\xb4\x2d\x4a\x1a\x74\xa6\x27\x8d\xa2\x68\x6d\xee\xb1\x50\x6a\x44\xe1\x61\xe1\x1c\xab\xeb\x47\xb6\x8f\xfa\x9a\xf3\xd0\xa6\xad\x86\xc3\x76\x0a\x25\xa8\x34\x8c\xe8\xba\x83\xf8\xdf\x6d\x18\xda\xb3\x20\x48\x22\x67\x5a\xd8\xa6\xb5\x95\x97\x6b\x55\xd6\x35\x92\xa6\xde\xc0\x54\xb4\x14\xb7\x16\x05\x9b\x36\xdc\x9f\x57\xab\xd5\xea\x2f\xf0\x67\x77\xf5\x66\xad\x33\xf9\x17\x7c\xfa\xa2\x73\x9a\x72\x8b\xba\x78\xed\x03\x8f\xf6\xec\xb8\x3b\x58\xb8\xf7\x7d\xbf\xe7\x12\x07\xb4\x75\xe3\x8e\x6c\x71\xda\x85\xdd\x22\xb7\x29\x22\xe1\x5e\xad\x41\x4d\x94\xe2\xe0\xc6\x9c\x6e\xef\x54\x63\x1b\xa5\x65\xa5\x3b\xc1\xf9\x91\xf3\x94\xf6\xa0\xec\xb3\xc3\xc1\x01\x8e\x78\x4c\x40\x7d\x5f\x04\x61\x8f\x5a\xb5\x2f\xb1\xc6\x9b\x47\x7c\x40\xe8\x76\xb1\xeb\x9e\x90\xf9\x33\xe3\x65\x17\x2b\x0f\xfa\xbd\x19\xe7\x9d\xe5\x9d\x42\xd7\x57\x9b\x52\x96\x2f\x2a\xb6\xe0\xb0\x90\xb7\x5c\xe5\xeb\xe7\xd8\xa8\x6e\xd1\x84\xd7\x7a\x93\xa3\x7b\x8f\x25\xaf\x1f\xaa\x75\x87\x92\xcf\x92\xd3\x30\x3c\x78\x79\x1c\x9d\x1e\x1e\x1d\x86\x07\x27\x67\xd1\x29\x7f\x99\x1c\xf2\xe3\xa3\xe7\xc9\xd1\xd9\xcb\xe8\xeb\x3b\x94\xbc\x90\xcf\x64\xca\xed\xff\x1e\x79\x24\xb9\x33\xe6\xb1\x37\x49\xfe\xf7\x7f\xe8\x22\xc9\xb3\xc3\x83\xe7\xa7\x70\xc3\x74\xb5\x14\x4a\xc2\x3b\x56\xea\x2a\x93\xa5\xf4\xe0\x77\x19\x2b\xcb\x7c\xf0\xd7\xe2\x0f\x8b\x8c\x89\x14\x57\xfa\xfd\xaf\xe8\x9e\xc9\xcf\x3f\xfd\xdd\x5d\x61\xf8\xf9\xa7\x7f\xfc\x76\xee\x9a\xf4\x7b\xf7\x5c\x47\x7f\xe2\x5d\x93\x7e\xaf\x7b\x1d\xfd\xb7\x7b\xd7\xe4\xe7\x9f\xfe\x4e\xd8\xf6\xf3\x4f\xff\xf8\x57\xdc\x37\xe9\xf7\x1e\x73\x23\xfd\x91\xf7\x4d\xfa\xbd\x8f\xdc\x48\x7f\xe4\x7d\x13\x73\x71\xfd\x81\x1b\xe9\x9f\x76\xdf\xa4\xdf\xbb\xe7\x46\xfa\xd3\xee\x9b\xf4\x7b\x1f\xbd\x70\xd2\xc5\xa6\xfa\xb2\xc9\xd1\xe1\x8b\x38\x3c\x39\x88\x0e\xe3\xd3\xe7\x3c\x3e\x08\x0f\x22\x76\x76\xfc\x3c\x3e\x7e\x19\xbe\xe4\x2f\x8e\xf8\xd7\x87\xeb\xc5\xfb\xc5\x3e\x57\x4a\x2a\xfd\x48\x50\x6f\x06\x6c\x7f\x59\xc8\x89\x07\x23\x76\xcb\xe1\x7c\xc9\x73\xbe\x82\xdf\xc5\xec\x96\xff\x21\xa2\x0f\x83\x9c\x97\xbf\xef\xf7\x7e\xfd\xb7\xe5\xc1\x5d\x96\xff\xdc\x9b\xf2\x9d\x1b\xea\x18\xcd\x7d\xfa\x5d\xf9\x6d\x44\x3c\xf2\xaa\xfc\x26\x1d\x48\xc2\x53\x2e\xcb\xc3\xb6\xbb\xf2\x38\xd9\x63\xaf\xcb\xc3\xda\x6d\xf9\x2f\x72\xe5\xdc\xc5\x9c\xfd\xde\x93\xaf\x9c\xc3\xda\x8d\xf3\x1a\xb6\x3f\xed\xca\xf9\x3d\xb0\x3c\xf5\xfb\xbd\x4f\xb8\x72\x6e\x19\xbd\xff\xca\xb9\xbd\x72\xf8\xb1\x3b\xe7\xf0\x98\x2b\xe7\xfd\xde\x03\x77\xce\xe1\xd3\xae\x9c\xf7\x7b\xdb\xef\x9c\xc3\x13\xae\x9c\xf7\x7b\x1b\x77\xce\xe1\x33\xae\x9c\xf7\x7b\xf6\xce\x39\xfc\xda\xae\x9c\x3f\x08\xc6\xb5\x47\x7a\x91\xf0\x17\x67\x87\xec\xe0\x30\x4e\x4e\x0e\xa3\x17\x2f\x4e\x8e\x92\xa3\xa3\x30\x3e\x48\x4e\x9f\x1f\xbe\x08\x4f\xd8\xd7\xe7\x91\x34\x06\xfc\xd5\xfe\x42\x16\xba\x2a\x45\xfa\x48\xb7\xb4\x36\x8a\x7c\x53\x3d\xc5\x5a\xc2\x6a\xf2\xbf\x57\xb3\x51\xdd\x16\xac\xc1\x33\x36\xf7\x47\x06\x5b\x5d\xdb\xb1\x07\xd7\xc3\x3f\x0e\x6f\x86\xef\x86\x30\xa3\xf5\x7e\x49\x67\xd6\xe9\x79\x7b\xfd\xde\x67\x3b\xb3\xcf\x7e\xef\xcb\x86\x37\x7b\xca\xab\x5f\xbe\xb0\x37\xfb\xb2\xee\xec\xc9\xfe\xec\x9e\x77\xbf\xe0\xdf\xb5\x16\xda\xcb\x27\x5b\x5f\x02\xd3\x3e\x82\xa9\xeb\x53\x16\x4f\x78\x0f\x0c\x6c\x7d\x0d\x8c\xa9\xc8\xfd\xb3\x5e\x05\x03\xad\x37\xc1\xa0\x6f\xfe\x02\x7e\xd9\x0d\xeb\xf7\xfe\x15\x7e\xf9\x11\x6f\x82\x31\x7c\x7e\xbe\x5b\x76\x99\x56\xfb\x15\x29\x4f\x75\xcb\xcd\x9b\x60\xfa\xbd\xcf\x75\xcb\xdd\x37\xc1\x18\x6e\x7f\x3d\x6f\x82\x01\xeb\x95\x9b\xc4\xef\xa9\x6e\xb9\xdf\x7b\x66\xfe\xf5\x7b\x74\x25\x2d\x67\xe9\xbe\xb9\xc2\xbc\x6f\x40\x6a\xb0\x90\x0e\x3e\xda\xfe\x85\x0a\x0e\x31\xa1\x8e\xc1\x6d\x6b\x8e\xe6\x65\x62\xfb\x3c\x8f\x64\x2c\xf2\x45\x33\x89\x5d\xed\xdb\xfb\xc5\xbe\xbd\x5f\xec\xdb\xfb\xc5\xbe\xbd\x5f\xec\xdb\xfb\xc5\xbe\xde\xf7\x8b\x3d\x22\xcb\xa9\xf3\x3d\x1e\x9f\x9e\x1c\x1e\x1e\x9d\x1d\x46\x3c\x3a\x09\x4f\xd9\x8b\xf8\x05\x0f\xc3\x33\x7e\x74\xc0\x79\x7c\x7c\xf0\xd5\x54\x20\x9b\x77\x6c\x7e\xd8\xd7\xab\x87\x8b\x8f\x6b\xcf\x7e\x7b\x2b\xe7\x37\xaf\xf9\xcd\x6b\x7e\xf3\x9a\xdf\xbc\xe6\x57\xed\x35\x1f\xf0\x16\xb5\xb7\x3c\x89\x8f\x5f\x9e\x1c\xf0\x28\x3c\x0d\x0f\x5f\xc4\x87\xd1\x09\x3f\xe6\xa7\xe1\x31\x3f\x88\xe3\xd3\xe3\xd3\x83\xd3\x5f\xda\x5b\xfe\x5f\x00\x00\x00\xff\xff\x3a\x09\xb5\x4e\xc3\x5f\x00\x00")

func creditsBytes() ([]byte, error) {
	return bindataRead(
		_credits,
		"CREDITS",
	)
}

func credits() (*asset, error) {
	bytes, err := creditsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "CREDITS", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"CREDITS": credits,
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
	"CREDITS": &bintree{credits, map[string]*bintree{}},
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
