// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package bindata generated by go-bindata.// sources:
// assets/nrpc.tmpl
package bindata

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _nrpcTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x6d\x8f\xdc\xc6\x79\x9f\xc9\x5f\xf1\x78\x21\x3b\xe4\x69\x8f\xab\x00\x75\x81\x6c\x73\x05\x8c\xb3\x6a\x0b\x88\x4e\xea\xdd\xd5\xf9\x60\x18\x07\x1e\x77\x76\x97\x15\x97\x5c\x71\x86\xf7\x82\x05\x01\xcb\x88\x6d\x35\xb1\x22\x05\x4e\x54\xa1\x96\x95\xc8\x70\x9c\x7c\xa9\x92\xb6\x80\x6b\x4b\x4e\xfc\x67\xbc\xa7\xd3\xbf\x28\xe6\x8d\x9c\x21\x87\xdc\xbd\x3b\xcb\x46\x8b\xea\x83\xee\x6e\x39\xf3\xcc\x33\xcf\xfb\x1b\x77\x36\x5b\x85\x73\xd3\x6b\xa3\x0d\x7f\x82\xa0\xbf\x06\xaf\x25\x57\xfd\xe0\x9a\x3f\x42\xec\x03\x2f\xcf\x6d\xb9\x62\x2b\xdb\xfd\x67\x14\x10\xb6\x08\x91\xab\xe5\x07\x86\x45\x57\x53\x34\x0c\x0f\x6a\x4b\xc5\xc7\xa6\x0d\x7e\xea\x4f\x70\x7d\x03\xff\xd8\x83\xd5\x3c\xb7\x7b\x3d\x58\x4f\x06\x08\x46\x28\x46\xa9\x4f\xd0\x00\x76\x0f\x61\x9a\x26\x24\x09\x56\x47\x28\x5e\x8d\xd3\x69\xe0\xc1\xab\x57\x60\xe3\xca\x36\x5c\x7c\xf5\xd2\xb6\x07\x36\xdd\x24\x2e\x04\xb3\x99\xbc\x68\x9e\x2f\x00\x82\x93\x2c\x0d\x10\xee\xc3\x6c\xe6\xbd\x86\x08\xdf\x63\x4f\x0d\x80\x6c\x3b\x9c\x4c\x93\x94\x80\x63\x5b\x9d\x20\x89\x09\x3a\x20\x1d\xdb\xea\x44\xc9\x88\xfe\x20\xe1\x04\x75\x6c\xdb\xea\x8c\x92\x64\x14\x21\x6f\x94\x44\x7e\x3c\xf2\x92\x74\xd4\x63\x87\xee\x66\x43\xfe\x0b\x5d\x3c\x0a\xc9\x38\xdb\xf5\x82\x64\xd2\x8b\x7d\x82\x57\xc3\x84\xfd\xf4\x46\xf4\x29\x25\x58\xea\xc7\x23\x04\x94\x44\x17\x0f\x48\xea\x5f\x62\x27\x63\x46\x4f\x6b\x36\x13\x3f\x56\x01\xc5\x03\xf9\x6b\x38\x84\xab\x69\x32\x41\x64\x8c\x32\x4c\x3f\x54\x0f\x99\x16\x4f\x7a\x41\x14\xa2\x98\xec\x70\xf4\x94\x07\x1d\x0d\xa2\xba\x99\x20\x7f\x12\x85\x31\xe9\x51\x8a\x75\x6c\xd7\x66\x2c\xed\xad\x3c\xbb\x71\x6b\xfe\xcb\xf7\x8e\xee\xdf\x9a\xff\xfc\xe1\x4a\x4f\x70\x9a\x23\xee\x6d\xa1\x74\x2f\x0c\x90\x64\x3f\xde\x0b\xa4\xd8\x29\x54\x96\x8f\x74\x79\x13\x5b\x6b\x32\x57\x2e\xd4\x44\x48\x5f\x2e\xc5\x48\xd9\x74\x19\x91\x71\x32\x60\xcb\x3d\xfe\x3b\xe5\x65\xaf\x47\x99\x2b\xf0\xca\x73\x01\x05\x42\x0c\x64\x8c\x20\x8c\x09\x4a\x87\x7e\x80\x80\x8c\x7d\x42\xa5\x66\x2f\x1c\xa0\x14\x43\x32\x64\xcf\x31\x5f\x5e\x85\x02\x78\x9c\x64\xd1\x00\xc2\xc9\x34\x42\x13\x14\x13\xcf\x26\x87\x53\x64\x3e\xa9\x38\x62\xa6\x10\x4e\x41\x58\x5c\xa1\xb7\x32\xff\xec\x9d\xa3\x07\xf7\xbf\xf9\xe2\xd6\xd1\xbd\x3f\x6d\x24\x9b\xe8\x7a\x86\x30\x39\xfa\xe3\xc3\xf9\x57\xb7\x0b\xb2\x87\x43\x88\x11\xa3\xed\xa5\x78\x9a\x91\x6d\x7a\x6a\xc7\x63\x22\x5e\x6c\xe9\x80\x94\x95\x73\x29\xc2\x59\xc4\x57\x71\x2a\x6e\x96\x1f\x14\xc2\x75\x6e\xc2\x10\xa9\x33\x0e\x38\x5e\x47\x77\xbf\x3c\xfa\xaf\xdf\xcc\xef\xdc\x7a\xfa\x87\x3f\x33\x44\xac\xd9\x4c\xd9\x93\xe7\x4e\x40\x0e\x40\x68\x8a\xb7\xce\x7f\xda\x96\x22\xdf\xaf\x21\xc2\x2f\x6b\x30\x02\x9a\x8c\xb1\x93\xbe\xf9\xe2\xc9\xb3\x4f\xee\xcd\x6f\xbf\x73\xf4\x9b\x3f\xaf\xf4\x98\xa1\xb0\xac\x2e\xcc\x66\xe0\x01\xa5\x3d\x49\xc3\x78\x24\xe0\xa3\x78\x20\x16\xd4\x69\xc8\xa8\xf2\x46\x12\x0e\xe6\xef\xfe\x7e\x7e\xfb\x9d\x02\xd4\x02\x3a\xd2\x1d\x9d\xe2\xd4\x14\x5d\x87\x95\xd9\xec\xb5\x84\xd3\x4c\xdd\x50\x00\x33\x23\x21\x31\xc0\x29\xf2\x27\x68\xb0\x89\xa6\xd1\x61\x15\x87\xd7\x7d\xbc\x45\x94\xe7\xc2\x2e\xd2\x83\xa7\x19\x1e\xc3\x30\x8b\x03\x47\x3f\xfe\x4a\x46\x8a\xf3\xdd\x2a\x02\xae\x2d\x50\x38\xfe\xeb\x87\xf3\xf7\x1f\x4b\x1a\x2a\xb6\x23\x46\x9a\x54\x28\xb2\x33\x8d\x0e\xa9\xe4\x50\xb3\x57\xbd\xc7\xd1\xef\xde\x9f\xbf\xff\xde\x26\xba\xde\x63\xcb\xa8\x4c\x3e\xf8\x45\xf5\x2e\x7e\x3c\x00\xc7\x0c\x9e\x91\xd4\x05\x27\x4e\x08\x38\xf5\x3b\xbb\xae\x80\x54\xde\x54\x01\x92\xe7\x5d\xa8\x11\x1a\xa5\x69\x92\xba\xb6\xf6\x69\x61\xd6\x84\xdc\x3e\xfb\xf8\x41\x45\x8d\x9e\x3e\xf9\xf0\xe8\xe3\xdf\x16\xca\xa4\x2d\x56\x04\x50\x59\x96\xdb\xb6\xc9\xe8\xda\x7b\x7e\x4a\x09\xd5\xeb\xc1\xf6\x18\x51\x21\xa1\x87\x40\x90\x50\x83\x40\xc2\x24\x06\xea\x26\xba\x30\x41\x3e\xce\x52\x34\x00\x9f\x00\x37\xc8\xab\x38\x1c\x20\xcf\xb6\xf8\x5f\x9b\xeb\xdb\xff\x90\xa4\x9a\x69\x59\x83\xd2\x56\x7b\x1b\x68\x7f\x2b\x9b\x4c\xfc\xf4\xf0\x0d\x14\x50\xc6\x28\xcf\xc4\x83\x2b\x53\x82\x67\xb6\x65\x59\x74\x7f\x1f\xf8\xbf\x0e\x25\xfb\x8e\x70\x01\x02\xbb\x9d\x12\xbb\x1d\x8a\xdd\x0e\x46\x41\x12\x0f\x70\xa7\x4b\x77\xbf\x8e\xa2\x69\xb1\xbb\xe5\x4e\x30\x4c\x52\x08\xfc\x28\xc2\xca\xed\xd4\xab\x71\x70\x57\x98\x9e\x87\x7b\xd4\xe3\x4e\xfc\xe9\x9b\xc3\x28\xf1\xc9\xdf\xfe\xcd\x5b\xe2\xe7\xec\x82\xf7\xa3\x3e\x5c\xf0\x2e\xfc\xb0\x0b\x17\xbc\x1f\xbd\xac\xfc\xce\x3f\xbf\xf0\xc3\x9c\xc1\x59\x4f\x62\x4c\x7e\xe2\xef\xa2\x48\x00\xe2\xda\xff\x16\xff\xc1\xee\x6d\x75\x84\x95\xee\xf4\xa1\xa3\x92\x92\x63\xc2\xe0\xb0\xff\xde\x94\xbb\x3a\xdc\x78\x75\x72\xd7\x2e\x58\x38\xf6\xe3\x41\x84\x52\x40\x07\x28\xc8\x1a\x38\x48\xcf\x41\xa9\xe4\xa0\x38\xf5\xf5\x8b\xcf\x93\x85\xe2\x90\x1d\x81\xde\x4e\x81\xde\x32\x2c\x34\xdf\xc9\xc8\x41\xf5\x6a\xff\x4b\x39\x18\x24\x59\x4c\x98\xe7\x66\x97\x83\x89\x3f\x40\x34\x14\xa4\x7e\x9c\x0b\x68\x17\x82\xc8\xc7\x38\x1c\x86\x3c\x48\xe4\x36\x06\xa8\xeb\x2e\x34\x72\x9d\xee\x5d\xc4\xd0\x75\x7a\x14\x4a\xeb\x0c\x15\x0f\x2a\x0c\xd5\xb4\x91\x21\xb7\xc3\x90\x55\xd9\xd6\x29\xae\xd0\x76\x03\xc1\x9b\xe7\x4a\xd3\x2e\x74\x50\x1c\x24\x83\x30\x1e\xd1\xdf\x39\x91\x76\x28\x91\x1a\xc8\x2d\x2c\x05\x16\xf2\x36\x90\x28\x0b\x04\xda\xa9\x2e\x16\x09\x33\xcd\x29\xaf\x44\x20\xdf\x0a\xe5\xa5\x12\x49\x44\x17\x51\x7f\xc1\x85\xbe\x7f\x26\xb8\xa5\xfb\x12\x11\x6e\x49\xb2\xd7\x85\xd2\x8b\x58\x16\x83\x0f\x1b\xaf\x6c\x6f\x01\xce\x76\x71\x90\x86\x53\x66\x07\xa4\x65\x60\x51\x6f\xe0\xc7\x5c\xff\xc1\xa7\xc0\xb4\x85\x19\x0e\xe3\x11\xf8\x30\x0a\xf7\x50\xac\x9d\x53\xc4\xb7\x32\x06\xf6\xe9\x06\x4f\xc4\xb2\x1b\x9b\x57\xd7\xe1\xf8\xd1\x27\xcf\xee\xbd\x2b\x10\x62\xee\xb4\x16\x25\x4b\x6c\x31\x49\xb3\x80\xc0\xcc\xb6\x68\x2c\xc9\xff\xd5\x22\xca\xfd\x24\xbd\x46\x83\x73\x80\x15\x16\x56\xfc\x94\xfd\x7d\x35\x49\x22\xdb\x8a\x03\xb1\x0b\x78\x40\xe3\x13\xbc\x9e\xc4\x71\x21\x60\x00\xa6\xe8\xdc\xb6\x24\x95\x31\x48\x1e\x48\x9f\xdf\x5b\x79\xf6\xf0\xf1\xfc\xf1\x6d\x15\x7f\x1a\x8f\xc1\x06\xda\x37\x5c\xc1\x14\x04\x77\x21\x0e\x74\x7c\xba\x80\x4d\x78\xb8\x34\xc6\xac\x53\x65\x66\x5b\x29\x22\x59\x1a\xc3\x4b\x86\xc7\x54\xc4\x02\x72\x20\x4d\x7e\x40\x0e\xa8\x28\xc5\x41\xbf\xa0\x44\x40\x3f\x10\x04\xa0\x9f\x62\xfa\x77\x71\xe3\x3e\x94\x62\x27\x53\xd7\x0e\x15\xc7\xbc\x24\xc1\xfc\xcb\xcf\xe7\xb7\x7f\xb5\x88\x04\xeb\x49\x1c\x64\x69\x8a\x62\x22\x89\x21\x59\x55\x65\xd4\xf3\xa7\x88\x38\xb9\x0f\xe2\x17\x8d\x26\x55\x8a\x60\x79\xdb\x5e\x0f\xb6\x10\xb9\x58\x08\x03\x46\x84\x67\x88\x09\x0b\xb7\xa1\x14\x93\xfd\x31\x2a\xd5\xe2\x07\x2b\x57\xb3\xdd\x28\xc4\xe3\x1f\xb0\x50\x9d\x6a\x80\x20\xdc\xf1\xa3\xbf\x3e\xfd\xcb\xa3\xa3\xcf\x6f\x1e\xdd\xf8\xd3\xd3\xaf\xee\x3e\xfd\xdd\x8d\xa3\xbb\x5f\xca\x4c\x8e\xd1\xd0\x19\x1b\xaf\xe8\x6a\x88\x38\x75\xf9\x74\x29\x11\xc6\x5e\xf9\x60\xad\xc4\xae\x64\xdc\xf1\x2f\xff\x7b\x7e\xfb\x2e\x3f\x9f\xeb\x21\x4f\xab\x96\x39\x9f\x27\x69\x8e\x2b\x52\x2d\x85\xe8\x1d\x5e\x28\xd1\x8a\x3f\x32\xb9\x10\x89\x6d\xad\x02\x24\x42\x7b\x4f\x0f\xd6\x2d\x7e\xb4\x58\x5a\x81\x51\x2b\x01\x70\x18\xde\x4a\x05\x86\xf7\xf7\x1d\x3b\x37\x54\x29\x78\x18\x5f\xa9\x55\xd4\x53\xee\xc6\xc4\xd7\xae\xe4\x3f\x2d\xd9\x38\xba\xbe\x28\x1b\xa7\x36\xba\x9d\xe0\x95\x74\x5a\x88\x94\x03\xad\x74\xcd\xf3\xe9\xb5\x11\x2b\x11\x09\x36\xd1\x0c\x59\x50\x46\xe4\x34\xf3\x0f\xde\x35\xe5\xd2\x2d\x84\xce\x73\xbc\x17\xb4\x03\xe5\x14\x6e\x87\xdb\x94\xef\xe7\xf9\x84\x2c\x80\xde\x58\x01\x98\xe0\x11\xb4\x64\xc2\xc0\x32\x42\x2a\xaa\x34\xbc\xdd\xe9\x16\x3a\x41\x39\xcb\xb1\x52\x75\x86\x5a\x8a\xd4\xdf\xbf\x8c\x47\x5d\xba\x93\x2e\x62\x6c\xbb\xec\xa7\x78\xec\x47\x85\xda\x75\x61\x82\x47\x34\xcd\xa6\xac\x4e\x53\x78\x61\x0d\xe2\x30\x62\xdb\xad\x28\x19\x79\x57\xd3\x30\x26\x43\xa7\x63\xe0\xab\x67\x66\x6b\x5f\x60\x3a\xe1\x27\x51\x0c\xa9\x9d\x99\x20\x8c\xfd\x11\xea\xc3\x8b\xb8\xc3\x50\xa2\x87\x4a\xb5\x43\x69\x4a\x23\x05\x6a\xba\xca\x12\x9a\xae\x8b\x79\xee\x75\xb4\x9a\x8b\x41\x60\xe0\x3c\x14\x32\x73\x1e\x3a\x5e\xa7\xa0\xbe\x6d\x59\xe7\x65\x84\xd2\x08\xb0\x2e\x2c\x70\x1e\x0a\x79\x31\x03\xac\x4a\x02\x95\x81\xce\x52\xb5\x21\x09\x12\xce\x83\x10\x1a\x06\x9c\x81\xa6\xcc\x90\xec\x7d\x61\x0d\x4a\xef\xc5\x19\x23\x89\x74\x7e\x4d\x00\x90\x8b\x05\x11\x05\x2f\xfb\x6b\x30\xf6\xe2\xc0\x93\xfa\x26\xb6\x75\x81\x0b\x86\xfb\x77\x35\x8e\x57\xd9\x91\x17\x86\x31\x0e\x23\xbb\x5e\x54\xe0\x46\x00\xa3\xa5\xcb\x0f\x8a\xf5\x52\x16\x4b\x5b\x23\x5c\xe3\x06\x42\x03\x2c\x1d\xa2\xd7\x52\xee\xe8\x3e\xbb\xff\xf6\xf1\x67\x37\xb8\x2b\xa2\x41\x20\x77\x07\x97\xf1\x48\x75\xe8\x0b\x8c\x53\xb9\xda\x61\x2a\xc8\xea\xd7\x94\x3c\x94\x24\x7b\x7e\x0a\xc6\xe2\x5f\x38\x84\xb1\x27\x83\x00\x85\x84\x74\xf1\x5a\xf9\xa8\xdc\x90\x03\x8a\x30\xd2\xd6\x04\xe4\x40\x90\x98\x57\x22\xa4\x8e\x6e\xa0\x7d\x71\x4d\x1a\x73\x75\x19\x13\x99\x9e\x7a\x5b\x92\x83\xf4\x0f\x56\x5c\x72\x59\xbe\x82\x0e\x48\xea\x07\x04\xb8\x3a\x42\x4c\x0d\xfe\x4b\xa5\x08\x0d\xd3\x64\x02\x82\xfb\xd4\x8c\x89\x3a\xd9\x05\x70\x22\x14\xd7\x15\xc9\x65\xa6\x97\xff\x3e\x9b\x51\xb4\xf3\x7c\xa7\x90\xfc\xae\x5a\x6a\x93\x20\xaa\xaa\xe3\x52\x13\xb8\x17\xe8\x30\x60\x47\x51\x9f\x2e\xc3\xb2\x0b\xc4\x0f\x23\xdd\x44\x5d\xf5\x53\x2c\xcb\xe0\x34\x09\xaa\xda\x81\x0e\xb5\xaa\x46\xc4\xf3\xbc\x2b\x74\xb2\x5e\x7b\x2f\x77\xd5\xd5\x5c\x23\xae\x6b\x1b\x6c\x61\x8b\x29\x8c\x06\x28\xed\x57\xaa\xe7\xe2\xd8\xa9\x9f\xb2\x38\x6a\xe8\x87\x11\x1a\xf4\xe1\xc5\xbd\xd2\xf4\x71\xbd\xa2\xfc\x2f\x04\x40\x54\xf5\x99\xbb\x5e\x63\xe4\x29\x1f\x09\xe4\xb6\xfd\x30\x82\x35\x46\x34\x5b\x73\x74\x61\x17\xce\xc5\xc2\xcf\xd7\x43\x94\x3c\x57\x20\x21\x22\xda\x4b\xec\x21\xbb\x4f\x2c\x92\x36\x28\xf8\xfe\xe6\x6c\x76\x2e\xcc\xf3\xb7\x94\x12\x64\xc5\xb5\x6a\x27\xd6\x02\x9a\xca\x89\x82\x1f\x86\x13\x0b\x29\x31\x9f\xc8\xc4\x3b\xf0\xa3\xa8\x48\xe9\xfc\x78\x00\xc3\x24\x9d\xd0\x3c\x7b\x9a\xc4\x18\x71\x2d\x0d\x27\x13\x34\x08\x7d\x82\x2e\x32\xef\xc3\xe3\x72\xf6\xbb\x6d\xe1\xfd\x90\x04\x63\xae\x16\x33\xcd\x38\xeb\x41\x53\xa3\xad\xe2\x15\x66\xaa\x39\x0d\x21\x15\xbb\x6f\xe0\x63\xd4\xe8\x13\xfa\x65\x21\x79\xa9\x90\xca\xb2\x78\x4f\x46\xf5\xb0\x10\xd2\x6c\x37\x4e\x56\xa5\xc5\xe0\x0f\x3d\xb8\x34\x8a\x93\x14\x41\x48\x3c\x45\xb2\x38\x19\x99\xe6\xcd\x66\xbd\x15\x78\xa1\x80\x0f\xc5\x7d\x2a\x7a\xec\x34\x39\x2b\xd7\x65\x1b\x28\xa1\x27\xf2\xc3\x22\x9f\xb4\xd4\x86\x5b\xe5\x92\x65\x0c\x53\x2d\xc5\xb3\xc5\x52\x44\xc4\x87\x54\xb4\xd3\x0c\xd5\x20\x72\x78\xcb\x20\xb9\x23\x8d\x8d\x44\x53\x38\xd5\xae\x2c\x7a\x78\x17\x8b\xc8\x87\xea\xb8\xc1\xe6\x50\x1d\x73\xb8\xad\x68\x3e\x4a\x85\xa8\xec\x5b\x2e\x92\x52\x79\xaa\x5a\x10\x8d\xd7\xcb\x18\x11\x6b\x37\x45\xfe\x35\xe1\xf5\x25\x99\x9a\x9b\x3e\x92\xe4\xcc\xd4\xa2\x7d\xa7\x78\xc2\xfb\x2c\x9c\x70\xfa\x92\xa6\xce\x90\xd8\xa1\xc4\x2b\x8a\x05\xff\xa7\x58\x04\x7f\x4e\x9d\xe8\xd4\xd2\xbe\xea\x13\x9f\x91\xcf\x10\x82\xb4\x93\x8a\x5a\x80\x3a\xa9\xa4\x3a\x64\xf2\x5c\x33\xb1\x2a\x26\x62\x0d\x5e\x2a\x8d\x04\x2f\x67\xd1\xbb\xf5\xa1\xfc\x74\x67\xfd\x27\x97\x2e\x6e\x6c\xb3\x62\x96\x75\x59\x46\xb0\x9d\x5d\x7f\x50\x9c\x99\xa2\x00\x85\x7b\xf4\x2c\x16\x88\xa5\x29\xdf\xe9\xb8\xbc\x02\x46\xff\x33\x76\xb5\x2d\x63\x65\xb0\x74\x22\xde\x4f\x43\x32\x66\xd5\xb7\x37\xfc\x28\x43\xd8\xe1\xf5\xb6\xca\xd5\x3b\x26\xb9\xee\x14\x74\xd8\xa1\x74\xe8\xb8\xde\xa5\x38\x70\x5c\x89\x8a\x64\x9a\x8c\x48\x64\x32\x75\xfb\xd6\xfc\xf1\xed\xf9\xcd\x7f\x9d\x7f\x70\xf7\xe8\xe6\x9d\xf9\xcf\x7f\x2b\xac\x44\x73\x2f\xaf\x4c\xc4\x64\x22\x89\x89\xd6\x0c\xe4\xfb\x4b\x0c\xfd\xdd\x08\x69\x30\x1c\x57\x5d\x20\x63\xbe\x35\xde\x17\x34\x04\x5e\xae\xc3\x42\x61\x4f\xf0\xa2\xcb\x73\x0d\x97\x73\xaf\x88\x79\x65\x1d\xd3\xd0\xc6\x65\x0b\x2b\x5e\x6c\x2a\x7a\xc7\x8d\x2d\x5c\xb6\xa9\x5b\x58\xbe\x37\xa9\x1f\x08\x21\xcf\xdf\x2a\xa0\x95\x59\xc1\x89\xda\xaf\xa2\x01\x6b\x04\xd3\xe5\x44\x48\xd1\xb4\x2d\x2d\xe4\x37\x57\x5c\x6d\x3c\xe0\x04\xe6\xe4\x4d\xd1\x94\x51\xd8\xca\xf9\x8f\x32\x9a\xef\x8a\x28\xbf\x14\x51\x26\x0d\x3a\xc7\xce\xc8\x10\xd5\x22\x35\x3b\x82\xe2\xc6\x61\x1c\xa3\x74\x13\xe1\xa9\x66\x9f\xc4\x32\x8e\xff\x72\xbc\x65\x37\x69\x3b\x5a\xe7\x81\xf9\x5c\x69\x17\x4f\x28\x57\x82\x8a\x42\x31\xa0\xe8\x29\x83\x2c\xad\xd4\x8f\xed\xc2\xc9\x4e\xa8\xc8\xda\x49\x24\x79\x61\x15\xe2\x7b\x10\x74\x4e\x64\x83\xd7\x34\x8a\x2b\x97\x57\xf9\x44\xa1\x21\xcd\x50\x35\x71\x16\x21\xab\x95\xab\xe1\x04\x8b\x84\xe2\x84\xd4\xad\x19\xb5\x56\x82\x3c\x2c\x90\x9b\x7f\xfa\xb3\xa7\x77\xde\x53\xb2\x54\xab\x9a\xd3\xaa\x11\x22\x8f\x19\xd5\xd5\x03\x34\xf4\xb3\x88\xf4\x65\x50\xf9\xe4\xde\xf1\xa3\x4f\x79\x82\xca\x41\x0b\xf3\xb8\xa0\xca\xd2\x87\x2c\xbe\x16\x27\xfb\x31\x0f\x5f\x5f\xbc\xde\xe1\x99\x13\x8b\x34\x16\x39\xb4\x36\x7f\x56\xba\x33\xf5\x00\xee\xc6\x58\x6a\xa6\x90\xae\xee\xbd\x4e\xe1\xbc\x2a\x0d\x22\x93\xe3\xa2\xe7\x56\x7d\x96\xe2\xb2\x72\xbb\x61\x44\x4c\x82\x7a\x65\x48\xa8\x38\xf0\x48\x52\x98\x4f\x11\xf1\x32\x22\x08\x74\xbb\x80\xb3\x20\x40\x18\x53\x24\xa6\xd1\xe1\x16\xff\x0b\x76\x93\x24\x62\x69\x3f\x95\xc6\x17\xb4\x47\xb3\xd3\x7a\xec\x7a\x6e\x67\xbc\x39\x46\xf1\x80\x1d\x58\xbd\xbe\x08\xaf\xf0\x77\x85\x07\x3f\x47\x45\xa0\x2c\x5c\x3c\xcf\x83\xe5\x00\x80\xe1\xfa\xbd\x1e\x65\x53\x92\xb2\x7c\x27\x0d\x03\x20\x89\xc2\xff\x52\x18\xeb\x93\x0a\x35\x8c\xea\xc8\xb8\xde\x95\x5d\xd6\x1c\x74\xb4\x50\x25\xf2\xa7\x18\x0d\x1c\xd7\xdb\xe2\x13\x08\x8e\xeb\xb2\x5a\x8d\xd2\x97\xa4\x6c\xa9\x6a\x60\x69\xbd\x9a\x0a\x44\xac\xbd\x9c\x1e\xc2\xf5\x0c\x65\xb2\x36\x2a\x4e\xb5\x2d\xbd\x74\x27\x6b\x48\xff\x98\xa1\x4c\xd2\x5c\xde\xc0\x10\x39\x6b\x96\x84\x8a\x7b\x1f\x38\x56\x86\xa3\x2a\x65\xd8\xbc\xc2\xe7\x5e\x0f\x36\xb3\x98\x6d\x90\x69\x37\x3e\x8c\x83\x71\x9a\xc4\x49\x86\xa3\x43\x95\x54\x9b\x59\xfc\x4a\xac\x44\x73\x42\x4f\xeb\xc4\x79\x41\x23\x8e\xb8\xa5\x1a\xb8\x70\x18\xcc\xd8\xeb\x5b\x17\xa7\x09\x75\xab\xa9\x15\x64\xe4\x25\x78\x5a\x40\xe5\x67\xca\x6b\xa2\x82\x26\xbc\x90\x50\x21\xca\xb7\x18\xb6\x9f\x4d\x07\xf4\xa8\x7d\x91\x49\x7e\xee\x5a\xa0\x5b\xe4\x5c\x6d\xd5\xd7\xda\xdf\xeb\x6c\xaa\x43\xe9\x7e\x97\x6d\xec\x86\x76\x76\xbd\xb6\x58\x56\xb3\x58\x11\x42\x19\xe3\xe6\xff\x64\x15\xa2\x32\x47\xdc\xd8\x25\x60\x20\xd8\x1f\xd5\x89\x4b\x05\x82\x7e\x44\xe5\x94\xa2\xb0\xdf\x38\xb3\x6b\x6d\x89\xb2\x56\xcb\x11\x92\xf1\xf5\x33\xb6\xc3\x09\x4a\x32\xe5\x70\x36\xd7\xe4\xbd\x9a\xa5\x6c\x06\xc1\x96\x35\xed\x6a\x97\x9a\x11\xdb\xa9\x36\xa0\x97\xe8\x5e\x96\x85\x3f\xb5\x7b\x55\x6b\x68\x36\xf7\x2f\x25\x18\xdc\x7e\x6d\x31\xc4\xa9\x57\xe2\x85\x8c\x34\x74\xbe\xf9\xd3\x59\xbd\xc9\xbd\x50\x50\x14\x49\xe9\xd7\x9a\x49\x9d\xae\xa1\x6c\xd5\x22\x32\x8a\xcc\x78\x79\xde\xd7\xe9\x55\x85\xa5\x9d\xaa\x76\x9c\x8a\x53\x1b\x3b\x4e\x74\x77\x41\x43\x76\x12\x56\xff\xac\x9e\x24\x65\xa8\xaf\xf4\x88\xe8\x1a\x21\x41\x7d\x78\x19\x56\xb8\xf0\x70\x0d\x96\xe3\x00\xbc\x25\x12\x98\x18\xa1\xf5\xe7\x8b\x3e\x21\x28\xbd\xf9\xa0\xb0\x59\x4a\x6b\x7e\x19\xa8\x02\x2b\x87\xe8\xf2\x2c\x80\x4a\xa9\x5f\x03\x62\x37\xf6\xb5\xdb\x4b\xb4\xa7\x9a\xf6\x5e\x66\x90\x5c\xec\x3d\xfe\xfa\xd7\xf3\x8f\x1e\xd0\x28\xff\xc1\xfd\xa7\xff\xf1\x84\x0d\x09\x9f\xac\x55\x2e\x91\xe0\xfd\x32\x06\xe2\xe8\xe3\x87\xdf\x3c\xf9\x7c\xfe\xd5\xdb\xf3\x3f\xfc\xa2\xdb\xda\x92\x6f\x27\x70\x25\x67\x2c\x3b\x28\xb3\xd9\x09\xbb\xd6\xb2\xb0\xa7\x0e\x48\x28\x8d\xd9\x96\x06\x92\x1c\x74\x0e\x3c\xc5\x4c\xcb\x36\x67\x59\x2f\x5c\x3c\x4e\xc1\x00\x28\x16\x49\x83\x51\xac\xa9\x9f\xd0\x3e\x60\x11\x78\x9a\x72\x19\xa1\x3e\xaf\xa6\x6e\x38\x04\x45\x71\xea\x5d\xdd\x5a\x53\xb7\x5c\xad\x36\x62\x65\x1b\xcf\xe4\x6d\xeb\x12\x50\x0e\xba\x95\x2e\x98\x77\x38\xd5\x87\xe5\x94\x18\x14\x23\x62\x34\x51\x9e\xdf\xf9\xe0\xe8\xdf\x7f\xcf\xa7\x6b\x6a\x99\x2b\x97\x45\xac\xcb\x62\x0b\x06\x2e\x6c\xa0\x03\xe2\x10\xa1\xe5\x15\x0b\xd0\xf2\x12\x80\xac\x29\x71\x0f\x41\x73\x3b\x53\x6d\xba\xf2\xda\xc0\x44\x19\x7c\xd0\x2f\xeb\x51\x34\x2e\xe3\x91\xc4\xc4\xd8\xef\xab\x15\x1e\x72\xdb\x52\xda\x05\x65\x81\x1b\x7b\xc8\x58\xda\xe6\x15\xab\xa5\x00\x17\xac\x65\xbb\xba\x45\xab\xbd\xb7\x32\xbf\xf9\xd1\xfc\xc9\x63\x95\x0b\xc7\x6f\xbf\x3b\xff\xf2\xf3\xe3\xaf\x3f\x3a\x7e\xf8\x01\xff\x44\xe1\xc5\xd2\x76\x81\xd2\x62\x17\x6d\x1d\xc6\xc1\xd9\xac\x83\xb3\x2c\xef\x55\x16\x2a\x86\x24\xa8\xd6\xb9\x94\xae\xef\xd2\x68\x15\xf8\x58\xae\x6d\x51\xd1\xde\xca\x76\x0b\xd6\x07\x5e\x1c\x78\xfa\x8d\x71\x4b\x9b\xd7\xc4\x76\x53\x50\xd2\x72\xd5\x59\x81\x42\xa9\xbf\xb9\xc6\xd4\xaf\xde\x29\xd8\x79\x7a\xe6\x39\xa7\x1a\x47\x52\xec\x5c\x91\x12\xb1\xd3\x5b\x2a\xcc\x94\xcd\x35\x8b\xf1\x9d\x31\x14\x37\x32\xb3\x1c\x67\x61\x85\x1e\x7d\x76\x83\x33\x73\x49\x5b\x61\x99\x7b\x57\x41\x43\xd7\x4a\x14\xa3\x4f\x30\x30\xc5\x99\x69\xa0\x0d\xbf\x89\xcc\xd3\x07\x48\x1e\x67\x9a\x91\x12\xb9\x9f\xe0\x9b\x23\xf1\xa0\x17\x28\x5d\x03\x17\xdb\xba\xa0\x89\xc0\xe5\xd1\xc3\x67\x37\x3e\x3c\xbd\xd0\xad\x8f\xfd\xf8\x8c\x16\xe3\xc7\xab\xc1\xd8\x8f\x5b\x3a\x1a\x5d\x68\x97\xb6\x60\x4c\x59\x35\xf1\xaf\x21\x67\x01\xa4\x9a\xf4\xb4\xe8\xd2\x49\x94\xa9\xab\x85\x0b\x9c\x90\x5c\xd6\xda\xc6\xf7\xd8\xe0\xcf\x18\x7e\xbc\xca\x25\x48\x63\x5c\x30\xee\x56\x99\xb7\xe0\xc5\x2e\xd1\xe0\x38\x5d\x4b\xee\x44\xcc\x77\xf8\xc4\x7a\x75\xf0\x7b\x19\xf3\x23\xd3\x46\x35\x15\xec\x9a\xde\xf7\x6d\xed\x20\x88\xfa\x6e\xf3\x4b\x8a\x3a\xc8\x60\x57\x98\xb4\xda\xa8\x7a\x0b\x77\xba\xb6\x32\x5f\xc9\xc3\x5b\xbd\xb8\x82\x89\x9f\x32\x0b\xc7\xc2\x96\x8d\x64\xdf\x71\xb5\x53\xff\x3f\x3a\x5e\xea\x75\x58\x0d\x28\xfb\x7b\x36\x2b\x27\x7d\x54\x6d\x65\xfc\xe7\xb2\xbd\xee\x47\x11\x1f\x84\x0b\xd8\x20\x5c\x31\xc8\x76\xe2\x17\x5d\x95\x99\x98\xe2\xe3\x97\x8a\xa5\xb3\xfa\xeb\xae\xaa\x13\xa7\xbf\x6f\xb7\x84\x8c\xe6\xa2\x5c\xf3\x4b\x57\x8d\x3d\x92\x4a\x7f\x5f\xc5\xa0\xc3\x84\xd4\xdc\x1c\xd1\x86\x37\x79\x54\xb9\xa4\xff\x1b\x72\xb9\x17\x01\x2e\xce\x76\x59\x88\xec\xb4\x79\x3a\x75\xee\x24\xd8\xe5\xdc\x29\x3c\x52\x41\x9d\xb5\xb5\xa2\xf7\x74\xf1\xca\x96\x72\x06\xeb\xd5\x35\xbd\xef\x5f\x96\xfb\x49\x42\xfc\x88\xbf\xd3\x47\xfc\x6b\x28\xae\x16\xfe\x11\xaf\x4f\x16\x7a\xb9\x15\xc6\x01\x72\x98\xae\x2a\x25\xcb\xe6\x77\x51\x6b\x2c\xa8\x91\xbf\xac\x8b\x8a\xb3\xdc\xb6\xf7\xe8\x4c\x2c\x5d\xc4\xd1\x5a\xc3\x45\x61\xa9\xc2\x51\xee\x15\x34\x5b\x6e\xf0\x09\x65\x7d\xc1\xfc\x42\xf3\x29\x8c\xff\x29\xcd\x3c\x08\xc5\x5e\x5d\xd6\xcc\xf3\x95\x0b\xde\x46\x57\xb4\xd3\x05\xa7\x04\x9c\x10\x70\xd0\xf5\xc6\x37\xb3\xab\x13\x07\xc2\x0e\x37\xbe\x8d\xad\x18\xbc\x32\xf0\x30\xca\xea\xb7\xe8\x17\xe0\x6c\x5e\x01\x9e\x83\x4f\x80\xef\xd8\x23\x28\xe5\x12\x7b\xd1\xb0\xa4\x16\x1b\x18\xe6\x36\xea\x51\x46\x8b\x80\x70\x38\xe6\x01\x90\xca\x40\x49\xcb\x97\x0b\x18\xa1\x94\xe3\x2b\xa5\x92\xea\xcb\xcc\x52\xa8\x5f\x41\xf5\x8a\xcc\x1f\xa6\xe8\x7a\x97\x35\xae\x2a\x5e\xf1\x3b\xf5\x58\x8b\x6c\x5b\xe0\x47\xb5\x11\x34\xb5\x81\x2a\x31\x38\x91\xee\xd6\x7c\x9d\x46\x58\xc3\xc8\xc8\x32\x03\x04\xff\x87\x7c\xce\x19\x1c\xce\x32\xaa\x62\x16\xfa\xe2\x1d\x91\x8a\x94\x8b\xba\x56\x31\x9f\x53\x9e\x55\x7f\x45\xc4\xe8\xb3\x5a\x5e\x29\xa9\xd7\xd0\xdb\x16\xb7\xcd\xea\xf0\x4a\xea\xe2\x5e\x65\xa5\xad\x26\x26\x7c\xca\x3e\x9e\xec\x76\x89\x07\x65\x1b\x4f\x6f\xdf\x2d\x6e\x5d\x4d\xcf\xde\xe3\x6c\x6d\xe8\x15\xbc\xae\x7e\xcb\x91\xa1\xbd\x52\xce\x3e\x35\x7c\xfb\x91\xa5\xb5\xda\x0d\x61\x45\x7d\x72\x4a\x85\xaf\x7f\x19\x89\xec\x6b\x9e\xb6\x97\x09\xe6\x56\x26\x68\x91\xc3\x4a\xd9\x72\x0c\xe8\x65\x0c\x3d\x46\xd0\xdf\x31\xae\xb0\xba\x0f\x95\x7e\x9b\xce\x71\x0a\xe3\xe5\x15\xbd\xed\xb6\x44\xbf\x72\xaa\xf4\x2b\xc5\xb7\x5c\x9c\xb5\x6b\x39\x5d\xbe\x6b\x99\x2f\x2d\x11\xbc\x55\xb5\x8c\x5c\x9c\xe2\xdb\xae\x14\x13\xd4\xf6\x72\x91\x56\xe4\xf8\xcf\x87\x47\xf7\xff\xc5\xf4\xde\x26\x2b\x08\x78\x95\xaf\xba\x68\xea\x9a\x2f\x95\x79\x53\x63\xaa\x51\x51\x25\x62\x3d\xec\x58\x42\xe8\x8b\x02\x73\xa0\xb5\x50\x97\x6f\xc6\x56\xcd\x90\xd2\x93\x35\x72\x54\xb4\xa2\xb4\xaf\xa1\x50\xdf\x60\xd3\x9e\x18\x5b\xbd\xfa\x20\x54\x03\xd6\xed\xcd\xde\x8a\x89\x5c\x03\xf2\x6d\x20\xab\x81\xab\x22\x79\x12\x7b\xd7\x22\xd6\xa7\x92\xea\xba\xf2\x2f\x14\x6a\x93\x44\x77\xe1\x9b\xbf\x7c\xfd\xf4\xd7\x7f\x7c\xf6\xc9\xc7\xf3\x4f\xef\x3e\xfd\xb7\x9f\x71\x87\x26\xde\xc5\xbf\x79\x67\xfe\xab\x7b\x95\x24\x4f\x61\x88\x2a\xf4\x8e\x3c\xeb\x8b\xcf\x38\x64\x0a\xab\x62\xee\x5b\x66\x14\xb4\xea\xae\x4a\x66\xc1\xdd\xa5\x34\x8e\xb1\xaf\x9c\xe5\x68\x33\x63\x15\x9d\x6b\xb1\x82\xc6\x89\x0a\x91\x9b\x56\xb6\xb9\x55\x4c\x55\x59\xaf\xa9\x55\x6d\x71\x29\x6b\x55\x69\x36\x6a\x44\x85\x06\xce\x59\x68\xdc\x30\x3e\xc3\x83\x44\x2a\x8a\xdf\x3f\xc1\xeb\x9b\x8a\x28\x76\x11\x99\xf9\xa2\x16\xf2\x16\x65\x73\x1e\x5d\xd4\x83\xbd\xf9\x17\x9f\xd5\x75\xa7\x36\xcf\xbd\xa4\x7d\x36\x7f\x3f\x1a\xe3\x6f\x18\x87\xc4\x71\xb5\xef\x3d\x54\x2d\x17\x4b\x2a\x46\x21\x26\x28\x15\x93\xab\x98\x7d\x37\x95\xfc\xc6\x16\xed\x9b\x3a\xd5\xef\xfa\xb9\x9c\x61\xb2\x29\x76\x3a\x7a\x22\x51\xee\x70\x9b\xb7\x54\x26\x01\x97\xda\x53\xcd\x30\x4e\x72\x50\xc3\x97\x1b\xa9\xdf\xe2\xa3\x8e\x09\xfe\x4f\x00\x00\x00\xff\xff\x96\xc4\xbf\xe8\xc1\x55\x00\x00")

func nrpcTmplBytes() ([]byte, error) {
	return bindataRead(
		_nrpcTmpl,
		"nrpc.tmpl",
	)
}

func nrpcTmpl() (*asset, error) {
	bytes, err := nrpcTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nrpc.tmpl", size: 21953, mode: os.FileMode(420), modTime: time.Unix(1600425086, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"nrpc.tmpl": nrpcTmpl,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"nrpc.tmpl": &bintree{nrpcTmpl, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
