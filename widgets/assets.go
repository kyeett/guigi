// Code generated by go-bindata.
// sources:
// arrow.png
// checkbox.png
// DO NOT EDIT!

package widgets

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

var _arrowPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9b\x09\x50\x53\xc9\xba\xc7\x9b\xd1\x61\xdf\x54\x5c\x11\x6f\x08\x2a\x88\x64\x39\x21\x01\x12\x43\x58\xc2\x2a\x44\x04\x82\x04\x57\x4e\x92\x93\x10\x21\x8b\x49\x94\x80\x23\x03\x22\x8a\xcb\x15\x11\x15\x04\x41\xd4\x51\x10\x04\x17\x14\x45\x50\x11\x15\xdc\x00\xd1\x71\x05\x45\x05\x44\x74\xc0\x37\xa2\x83\x1b\xf0\x2a\xb8\x21\xe0\xd4\xdc\x37\xf7\xdd\x57\xaf\xea\x9c\x2a\xa8\xca\xd7\xfd\xff\xff\x4e\x6f\x5f\x77\x43\x65\xed\xec\x59\x9e\x06\xba\x13\x74\x01\x00\x06\xde\x5e\x6e\x01\x00\x00\x6d\x00\x80\x8e\xb6\x26\x00\xc0\x28\x50\x8c\x57\x7f\x90\x79\x85\x28\x00\xd0\x1b\xa5\xfe\xd1\x00\xe9\x3b\xc6\x01\xa0\xeb\x27\x62\x32\x67\xcf\x0e\x93\x2a\xa5\x8a\x30\xa9\x0c\xe3\xcd\x64\x62\x64\x72\xa9\x40\x14\x81\x00\xa0\xba\x97\x19\x28\x64\x07\xb6\x8d\x71\xec\x7a\xf0\xca\xd5\x27\x31\x76\x8b\x8f\x34\x60\xb4\x36\x26\xc0\x75\x7d\xac\x49\xf2\x54\x6b\x73\x1d\xe3\x99\x89\xe6\xbb\xee\x8e\xf6\x3f\x3b\xc2\xdd\x7d\xb8\xe9\xc5\xec\xc4\x61\x1b\x37\xc6\x6e\x18\xed\x6f\xa3\xbf\x4e\xf7\x8e\x56\xb3\xf9\xee\x8d\xf1\x7b\x12\xd7\x5d\x7e\xff\x68\xf9\x9e\xf0\x5b\x65\x5d\x4f\xaa\x7b\x8e\xdd\x66\xb4\x1e\xe8\xcc\x2c\xaa\xd5\x2a\xd7\xd6\xdf\xe6\x65\xeb\x4f\x89\xd5\xdf\xef\x3a\x69\xc4\xc3\xf8\xaa\xf3\x95\x8f\x9a\xf1\xce\x71\xba\x93\x65\xc0\x58\xbb\x88\x6f\xae\xe8\xb6\xd0\x00\x6f\x62\xe8\xf4\xe9\xd8\x72\xad\xd3\x40\x43\x75\x43\x4f\x1b\x94\xb3\x32\xca\x89\x13\xd7\xf4\x1a\x3d\x73\xdd\x3e\x5f\x23\x76\x95\x46\xb9\x32\xd9\xc1\x47\x3b\xd6\x04\x38\x47\x6f\x74\xdd\x07\x9c\x3d\x34\x62\x33\xd3\xa6\x06\x82\xec\xe1\x20\xf4\x0a\x8f\xdb\x04\x66\xe3\x40\xa8\xe5\xcf\x4f\x6e\x80\xd8\xcc\x37\xdb\x97\x6b\x80\x79\x5b\xcc\x47\x6b\x64\x27\x03\xcc\x28\x24\xd1\x0d\x84\xcd\x00\x07\x6a\xe7\xac\x73\x07\x1c\x22\x30\x11\xf8\x9c\xa7\x82\x3b\x38\x40\xf4\x0e\x16\x78\x81\x43\xc7\x41\x79\x8d\xb1\xde\x11\xa0\xad\x0f\x88\xfe\x89\x09\xd3\xc0\xf0\xe5\x20\xf4\xa2\x85\x85\x0a\xac\xda\x0e\x4c\x3c\x5e\x06\xd3\x7f\x9f\x5e\x68\xdc\x6a\x0d\x80\x6a\xfb\x11\x7a\xc9\x64\xaa\x7b\xdc\x3c\x6d\xdc\x12\xad\xe0\x60\xfc\x24\xab\x1a\xef\xb1\xd3\x4d\xec\xe0\x1f\xe0\x74\x67\x7c\x45\xd4\xf6\xf1\xa4\x38\xb2\xd1\xc3\xce\xab\x00\x64\x27\x8f\x1e\xf1\x30\xbe\xb3\x5b\x55\x63\x58\x58\x53\x63\xb7\xb1\xd8\x70\x11\xee\xd9\x19\xcd\x5e\x2c\x97\xdb\xd4\xd3\x52\x9b\x2f\x73\x06\xe0\x91\x32\xa6\xae\xc7\x86\x50\x34\x21\x76\x86\x66\xac\xb8\x67\xf7\x94\x8e\xe1\x61\x0b\xf5\xb2\x97\x76\xad\x17\x1c\x31\x72\x3e\x06\xb2\xdb\xee\x05\x77\xca\xd4\x7d\xe3\xbe\x6b\x53\xc9\xdd\xbb\x2d\xcd\xcd\x77\x66\x9e\x77\x9d\x07\x5f\x0e\x5e\xd1\x23\xac\x58\x54\x1e\xdc\x1d\xf1\x36\x86\xfe\xa1\xeb\xf5\xc3\xd3\x8f\x27\xc7\x91\x16\xc4\x79\x0e\x7f\xf3\xb0\xe2\xea\x2b\xdf\xbd\x13\x6b\x37\xeb\xac\xe2\x36\x3c\x49\xf2\xfc\x70\x5c\xaf\xf7\x83\xd9\x0d\xec\x79\x6f\x6e\x09\xc7\xda\xa4\xc9\x7f\xe4\x92\x8d\x6e\xe9\x1b\x3a\x49\x7b\xb1\x6b\x5d\x37\x4f\xbb\x30\xe5\x7a\x2f\xff\x71\xa3\xcd\xfb\x61\x8c\x46\x0a\x38\xb4\x98\xa1\xb1\x3c\x19\x97\xbb\x80\x37\xec\x0f\x1f\x4b\x63\x8e\xb2\xfc\x71\x17\x00\x2f\x4b\xa5\x67\x7f\xb5\xd2\x1e\x16\x1b\x16\xff\xa8\x36\xb2\xf7\x85\x53\xb7\x67\xf6\x14\x10\x2b\xf0\xde\xbc\x04\x80\x85\x6e\x93\xf1\xec\xcb\x45\x4e\x55\xda\x00\xb8\x65\xc7\x59\x1f\x72\x31\x7d\x75\x76\x9c\x4d\xb9\xe6\xc4\xb3\x3f\x36\x9c\xd5\xfb\x10\x6a\xb7\xca\x75\xf2\xf9\x0a\x57\x63\x57\x5d\x7e\xec\x44\x79\xe8\xb4\x03\x2b\x5d\xad\x56\x57\x1f\x32\xb3\x0e\x25\xfc\xea\x6c\x46\x2e\xf7\x0f\x5b\x9d\x22\x33\xc4\x57\x04\x9d\x30\x74\x94\x19\xbd\x83\xb5\xb3\xce\x5a\x95\xc5\xe9\x0c\x77\xb9\xa8\x3d\x0a\x36\xbf\x97\xa8\xc1\x5d\xc7\xb1\xd8\xa8\x3d\x76\x7b\xc2\x1b\xf3\xb3\x5e\xc3\x47\xae\x9f\x8c\xb9\x94\x68\xe2\x8f\xf5\x49\xa6\x78\x41\x01\xeb\xd8\x98\x20\x2f\xb3\x97\x8e\xf2\x1f\xa8\x6b\xe6\x7b\x9c\x4c\x9e\x71\x11\x90\x56\x7e\x58\x5d\xe6\x09\x6d\x89\x3b\xe2\xef\xf8\xab\xfe\x31\x66\x49\xae\xfd\x84\x55\x1e\xbb\xf5\xae\xf1\x95\x50\x7a\xd2\x98\xec\x8c\x6b\xf2\xa3\x13\x1c\x37\xd8\xed\x6a\xbd\x96\x1f\x66\x74\x2f\x89\xc1\xac\x22\x96\x3f\x1d\xa7\x65\x7b\x3e\x82\x88\xaf\x3c\xe7\x7a\xd8\xc1\x46\x27\x85\x77\xe5\x30\xf6\x80\x7e\x7e\xe5\x84\x43\x73\x0e\x0c\x2b\x5d\xfb\x9a\x57\x7b\x38\x72\xf8\xc8\x38\xf7\x72\x1d\xbb\xe1\x91\xae\x18\x78\x92\x8b\x17\xc7\x97\x33\xeb\x8e\x57\x2e\x66\xe2\x64\xf2\x88\x7b\x3a\xf6\x3a\xc3\x56\xf9\x57\x58\x62\x53\x72\x7d\x8e\x4e\x3d\x39\xe2\xed\xf9\x71\x5c\x12\xb4\x86\xe3\x65\x63\x71\x6b\x64\xd2\x28\x82\xf1\xf2\x44\xe6\x79\xfb\xc9\x23\xd7\x32\x49\x96\x25\x9b\x7e\xf0\x3e\xfc\x0b\xfb\xfa\xc8\xeb\x1e\xd7\x25\x01\x16\x6d\x36\xfb\x33\xc8\x5e\x13\x2c\x52\x2e\xdd\xe3\x87\x1f\xd1\xa2\x6e\xb1\xc1\x4c\xbf\x98\xd0\x9c\xd9\x7c\xae\xd9\xb6\xd9\xb4\xd9\xfc\xe5\x3c\xdd\x4c\xae\xf3\x92\x5f\xde\x04\x3e\xb6\x0a\xb8\xbc\x78\xe6\xf8\x16\xe3\x16\x72\x8b\x26\x92\x46\x72\x09\xda\x93\x1e\x70\x3f\x77\x17\x7b\x04\x25\x69\xae\x6b\xc1\xae\xe2\x80\x83\xfe\x6b\x72\x47\x91\x97\xa6\x53\xab\xe3\xc7\x78\xfb\xec\x2d\xd8\xbd\xf5\xc6\x84\xb0\xb1\x61\x5d\xa2\x83\x4f\x95\x27\x47\xdc\xdf\x3a\x69\x56\xea\xb5\xe0\xeb\x57\xda\x3c\x9f\xa6\x3e\x1d\xf6\x2a\x52\xdf\x60\xf5\xe8\xf8\xaa\x44\xa1\xd5\xf8\x8c\x71\x5a\xe3\x84\xe3\x23\xc7\x1d\x4c\xf3\xdd\x37\xfe\xc6\xfa\xf1\x55\x6c\xa2\x3d\xf4\x3c\xad\x75\xbb\xd1\x76\x69\x10\xc7\xba\xd8\xfa\xa7\xd1\xc9\x97\x75\xb7\xce\xdc\x3a\x75\xab\xa9\x35\x81\xbd\x3f\x6f\x5f\xde\x83\x3c\xbd\xa0\xce\xa0\x2b\xec\xa4\xfd\x0b\x03\x57\x14\x7a\x04\xd9\x06\x36\xed\xd7\xf9\x75\x5b\x7e\x50\x5e\xe6\xec\xe7\x81\xa6\x81\xc2\xfd\x33\x72\xd7\xe4\x09\x73\xa5\x6c\xa7\x5f\x62\x42\xba\xd7\x6a\x86\x9c\xf3\x3d\xe7\xc7\xfc\xc5\x3f\xfb\xf0\xa2\xa6\xb2\x4a\x8c\x1e\x6f\x7c\x91\xca\xc6\x3c\xdc\x20\x7f\xf5\xd2\x38\xca\xf9\xa9\x87\xab\xef\xfe\xb4\xcc\xba\x34\xad\xdb\xac\x74\xbd\x63\xde\xfe\x5d\x54\xa6\xed\x4c\xdb\xe0\xa2\x80\xce\x2d\x02\x27\xad\xa6\xb4\x25\x6f\x76\x28\x67\xe9\xb4\x5a\xe5\xa4\x76\xc6\xac\x7b\x60\xf2\x62\xda\xe3\x69\xa6\xfb\xc8\xbe\x25\x10\x32\x3f\x3c\x6f\x5b\xea\x36\x4e\xbe\x57\xbe\x7f\xbe\xe7\xf3\x53\x76\x85\xed\x39\x99\xfb\x4e\xb9\x97\x86\xbc\x8f\x33\x74\xa9\x9a\x69\x39\xdf\xd2\x53\x98\x54\xf3\x63\x66\x6b\xe0\xdd\xc0\xfb\x39\xba\x39\x63\x58\x8e\xb4\xd7\x85\xf8\xbc\xf0\xbd\xa6\xbb\x4a\xdd\x22\xcf\x48\x72\x3a\xd2\xb2\x4e\x2f\x8c\xcd\xf2\x95\xe6\x3c\x88\x3a\xd5\xfa\x6e\xe2\x0a\x42\x37\xbb\x5b\xfc\x2e\xe7\xb5\xa7\x3e\x47\xb3\x48\x7f\xb2\x66\xbd\x7e\xcd\x38\xfa\x05\xca\x12\x29\xdd\x8c\x50\x55\xff\xbb\x3f\x73\xc1\x39\x67\xe6\x5e\x78\x2b\x9c\xfa\x68\x55\x72\x49\x1d\x7d\xcf\x4d\x8f\x7f\x7a\x18\x6d\x10\x55\xce\x6f\x9a\xdc\x24\xae\x14\x57\xee\x99\xaa\x39\xd5\x7c\xaa\x0f\xeb\x19\xab\x23\x3d\x80\x55\xb2\x6f\xda\xbe\x19\xbe\x33\x7c\xab\xaa\xaf\x56\xef\xae\xbe\xb5\x9d\x92\xe1\x08\xdd\x26\xb5\x66\xb4\x66\xdc\xce\x68\x3c\x1e\x3a\x57\x35\xd7\xa6\x38\xb7\xf8\x04\xe2\x7d\xa4\x79\xee\xbe\x90\x8e\xe2\x32\x49\xc2\x5c\x6a\xc8\x5e\x4e\xd0\xdc\x05\x45\x1e\x05\x99\x05\x53\x6e\x95\x16\x94\xe6\x8c\xc8\x29\x99\x95\x32\x2b\x3f\xeb\x96\xf0\xdc\xa1\xdf\x8b\x2f\x15\xef\x3f\x3e\xfe\x50\xcb\x1d\xcd\xfa\xc2\x62\x71\xf1\x0a\x58\xc8\x1d\x2b\x3c\x7e\xa1\xf0\x5c\x47\xca\xa9\x94\x15\x27\x57\xfc\xb3\xdb\x51\xcb\x60\xcd\x5d\xc3\x7c\xb3\xb9\x66\x31\x8a\xa3\x72\x13\xda\x1c\xda\x83\xf4\x9a\x1d\x1d\x4e\x2b\x2e\xbc\xda\xd6\x46\x10\x72\x6a\x8d\x6b\xef\xd8\x36\x15\x8f\xba\x32\xed\xbc\x28\x33\x30\xe3\x86\xa7\x70\xac\xb0\x44\x11\xdf\x99\xb8\x7e\xad\xcf\xca\x2b\xf8\x54\x06\x3e\xc3\xb4\xf2\xfe\xc5\x53\xcd\x33\x95\xaf\x97\x2c\x78\x14\x7d\xfd\x75\x66\x46\x51\x46\x45\x7b\x6e\xfd\xc1\x0e\xbb\x0e\x51\xc7\x9a\x86\x9d\x57\xc9\x3b\xa7\x65\x85\xd8\x2a\xae\x9d\xbb\xc8\x7c\xe8\xdf\x52\xe7\xa4\xb7\x73\x2b\x31\xbc\xf1\x40\x8c\x1f\xd3\x2f\x4e\x78\xac\xd5\x82\xc0\x23\x24\x07\x17\xdc\x3e\x74\x7c\x6f\x64\x5e\x17\xfc\x62\x2c\xe3\x0f\x86\x7c\x67\x13\xa3\x89\x6e\xff\x80\xf4\x20\x64\xd1\x55\xfb\xf2\xe3\x17\x8f\x5f\xf1\x3b\x27\xcd\x5a\x94\xd0\xc8\xfc\xd9\xa2\x97\xdf\x1b\xdd\x5b\x07\x72\x62\xa7\x68\x08\x87\xad\x8d\x73\x89\x9b\xf7\xc3\xfb\x37\xd1\x3f\x99\x9f\xd9\xe8\x44\xab\x47\xfe\x38\x72\xb7\xdd\x65\x43\x76\x89\xcb\x68\x97\xa4\x95\x05\x2b\x5b\xce\x16\x98\x87\x5d\xfc\x6d\x53\x81\x79\x57\x6b\x5a\x4d\x72\x8d\xc8\x30\x61\x76\xcc\xd3\x5b\xd7\xb4\xe0\x27\xf0\xfa\x73\x31\xfa\x96\x09\x6b\x12\x82\x57\x3b\x6d\x94\x25\x72\x26\xd4\x41\xa9\x94\xc8\xba\x06\xa6\xa2\xd1\xed\x9d\xdb\x36\xb7\x07\x19\x42\xc8\x8f\x7a\x10\x37\x93\x62\xcd\xc8\xc3\x2f\xa6\xdf\xb3\x3f\x49\x3e\xe9\x54\x89\x17\x84\x5c\xbd\xc9\xc9\xe4\x78\x47\xf8\x4c\xa4\xe2\xad\x0b\xed\x24\x11\xa7\x1b\x3f\xb4\x77\xae\x79\x9c\xf9\x58\xaf\x69\xf9\xd8\x34\xc2\xc5\x05\xb2\xa8\x5d\x35\x15\xaf\x12\x72\xc6\xec\xc7\x8f\xb4\xa1\x6d\x73\xf6\x4d\x9c\xe5\xb0\xe5\xbf\x5c\x8c\xad\x73\xc7\x6c\x98\xf2\x72\x54\xe4\x94\x8d\xba\x51\x15\x66\x15\x42\x2c\xa1\xce\xae\x2e\xf1\x0e\x73\xbe\xf5\xa4\x59\x2f\xbd\xac\x59\xd1\x9b\x7f\x4b\x2e\xa3\x44\x4d\xbe\x60\xbb\xd7\xa6\xa5\x64\x95\xe5\xfe\x4d\x65\xbe\x06\xbe\xf7\x33\x68\x59\x63\xd2\x97\x61\xc4\x24\x47\xda\x29\x28\x7e\x73\xf2\xa6\x25\xae\x79\xe6\xac\x94\x62\xe1\x33\xe1\x13\xc1\xb2\x2b\x07\x90\x84\x23\x3d\x17\x24\xeb\x8f\x13\x26\x6d\xd5\xba\x14\x5e\xb4\xe9\xc0\x8f\x21\x38\xdd\x62\xd7\x7d\x4f\xd9\x73\xad\x6e\x5b\x6b\x1d\x11\x22\xc1\x82\x33\xd5\x6b\x2f\x53\xb7\xe5\x67\x4d\xaf\x6e\xbf\x5a\x74\xb9\x7e\x93\xfd\x8e\x17\x3b\x7b\xdb\x47\x56\x8c\x7c\xb4\xb9\x78\xef\x6a\x9f\x60\x9c\x19\xe7\xc9\xa9\x03\x8b\xa9\x47\xe8\x4b\xa2\xea\x75\x9b\x46\xff\x38\x7d\xf5\x1c\xcd\x03\xf4\x13\xeb\x8a\xad\xc5\x37\xfd\xf2\xdb\x83\x53\x0b\x63\x4e\x4d\x8d\x66\x85\xc7\x2d\x3c\x7b\x79\xe5\xd9\x63\x1a\xe1\xcf\xf5\x82\x74\x5b\x12\x56\x3c\x6f\xb8\xdd\xbe\xc0\x8f\x43\x08\x69\x3c\x06\xbf\x5b\x9d\x69\xca\x18\xbd\x76\x5d\xa1\x49\xfd\xd8\xd4\xd1\x3f\xdf\xb1\x6a\x50\xb5\x98\xae\x91\x5e\x3e\xb3\xbd\x61\xcb\xd1\xfd\xf9\xe7\x4b\xf8\xb5\x82\xeb\x82\xb6\x27\x0f\x70\x59\xf3\x0c\xb7\xe5\xa5\x3e\x49\x35\x94\x18\xdc\xbd\x77\xd4\xa0\x71\xce\xd1\xf6\xd7\x7f\xb8\xd7\x13\x52\xab\xad\xea\x4b\xe3\x8b\x99\x25\x7b\xce\xd0\x4f\x94\x21\x39\x97\x6a\xaa\x67\x38\xb1\x7f\x61\xbf\x62\xbf\x63\xdb\x75\x5c\xbf\x7f\x0c\xbe\xfd\x46\x52\xd7\x61\xba\xbc\xec\x0f\xf2\x6f\x37\xeb\xe7\x95\x75\x75\x33\x2e\x9d\xae\x0f\xd2\x08\x62\xdd\x14\xdc\x44\xde\x7b\xbf\x2f\xea\x72\x2a\x5c\x9b\xd7\xf6\x66\xdb\x7b\xf1\x82\x82\x59\xe2\x25\xed\x55\xc3\xae\x0f\xeb\xd6\xb2\x30\x2c\xbe\x5d\xf2\xeb\x4d\xb3\xeb\x8e\xec\x80\xa4\x5b\xd3\x9f\x7b\x1a\x56\xfd\xe3\x74\x4f\xe6\x8b\x48\x9c\x81\x9d\x41\xd0\x3f\x77\x54\x3e\xb6\x90\xb9\x45\x35\xb7\x75\xe9\x34\x76\x4d\x32\xa0\xdd\x5b\xb6\x61\xf9\x81\x47\x5a\x37\xfe\xb1\xcc\x4c\x68\x46\x58\xec\x9a\xfe\x32\xdd\x23\x23\x30\xfd\xfd\x3c\xee\xbc\x93\x7e\x57\x19\x75\xb7\x9e\x3d\x88\xfe\xe0\x59\xff\x0f\xa3\x1d\x64\x7a\x48\xd4\xcc\x4e\xf6\x25\xd1\xbd\xb6\x95\x37\x52\x84\x2f\x6b\x37\x64\x6e\xcd\x5c\x11\x83\x7f\xdb\x10\x5a\xbf\xbb\x78\x79\x2b\xaf\xa6\x71\xb1\x91\xf4\xc5\x49\x63\xe5\xa3\xeb\xee\xbd\xeb\x5b\x67\x94\xcd\xa8\x99\xff\xf2\x84\xf2\xe5\x91\xfa\x96\x51\xc1\x27\x82\xd3\x0f\xd2\x16\x89\x9f\x29\x9f\xd1\x7b\x52\x6e\xf8\x55\x95\xee\x28\x4d\x2f\x2c\x99\x1f\x5e\xd2\x11\x7d\xea\xd8\x72\x41\x57\x9b\x69\x96\x7d\xcd\xfd\x57\xa7\xa2\x7e\x1f\xd3\x7b\xeb\x66\x16\xd5\xf1\xfe\xa9\xce\x68\x69\xd7\x85\xae\xd7\x8d\x63\xee\xc7\xc8\x2d\xbb\x6b\xae\x94\xde\x8d\xb4\xec\x28\x28\xf3\x2b\x5b\xf4\xca\xb9\xdd\xe5\x2e\xf3\xfe\x9a\x9b\x4e\xf3\x7a\x1a\xee\xbc\x68\xf8\x79\x71\x79\x64\x76\x67\x5d\x83\x67\x97\x66\x42\x7c\x6d\x4f\xf2\xeb\x24\x23\xe6\x8f\x49\xdb\x93\xae\xac\x36\x58\xfd\x3c\xcb\x81\x41\xa5\x3e\x76\x92\x77\xb7\x7d\xc8\x61\xd6\xf2\x6b\xa7\x66\x8d\xed\xd9\xdd\x73\xa5\x70\xec\x18\x42\xcc\x9b\x9b\xcf\x6f\xd7\xd4\xa5\xd5\xed\x4a\x7d\x99\x9a\x72\x46\xfc\xd3\xde\xee\x67\xe5\xcf\xc7\x1f\xac\xaf\xac\xa8\x9c\x74\xfa\x92\xca\x16\x5a\xf8\x6a\xce\xfb\xf6\xb3\x91\x0d\x91\x66\x2f\xf6\x34\x1a\xf9\x35\xff\x84\xf9\xd9\xaa\x37\xec\xed\xe1\xd7\x81\xd7\x72\xde\xa7\x8c\x4f\x49\xe9\x1d\x3e\x2c\xad\xf3\x03\xcf\x16\x57\x0f\x80\x53\xaa\x88\xcd\x51\x72\x58\xbe\x34\x9e\x54\x8c\x87\xf9\x52\x2e\x82\x57\x89\x65\x40\xfd\xd0\x9d\x54\x32\x98\x17\x8e\x28\x31\x5c\x44\x28\x92\x38\x62\x5f\x94\x9e\xc1\x62\x44\x7c\x47\x6c\x30\x85\x45\x64\xc9\x98\x48\x98\xc8\x2b\x5a\x8e\x04\x46\xcf\x62\xf3\xa2\xc3\x79\x54\x3e\xd6\x89\xa1\x4b\x57\xd1\x54\x62\x99\x18\x51\xc2\x18\x95\x38\x42\xa2\xa0\xa9\x1c\xb1\x7d\xbe\x34\x89\x82\xa6\x0e\x13\xb0\x98\xbe\x2a\xca\x70\x47\xac\x8b\xba\x00\xc3\x61\xcd\xc6\x30\xa5\x72\x04\x43\xc1\xdb\xe1\x78\x10\x04\x61\xec\xa9\x78\x88\xe2\x60\x4b\xa2\xd8\x60\x48\x44\x88\x42\x20\x52\x09\x10\x11\x47\x84\x68\x10\x91\x46\x22\x62\x3e\x3d\x58\x86\x2e\x06\x83\xa1\xcb\xf9\x02\x5a\x80\x9b\xc7\x27\x9c\x9c\x2f\x70\xc4\x86\x29\x95\x32\x1a\x81\x10\x19\x19\x89\x8f\xb4\xc5\x4b\xe5\x42\x02\x44\xa5\x52\x09\x44\x12\x81\x44\xc2\xc9\xf9\x02\x9c\x22\x4a\xa2\x84\x55\x38\x89\xc2\xe2\xa3\xc9\x67\x1f\x37\x44\xc1\x93\x8b\x64\x4a\x91\x54\x82\x51\x7f\x86\xb9\xd2\xa5\x4a\x47\x2c\x56\x17\xd3\xef\xf9\xd4\x2e\xb1\xec\x0b\x48\xa2\xf8\xd4\x77\x3c\xa9\x98\xa0\x82\x65\x04\x08\x4f\x24\x0c\x25\xe2\xf3\xbe\x68\x64\x4b\xe5\x11\x7d\xaf\xc6\xe7\x11\x90\x08\x44\x8c\x48\x94\x0a\x02\x84\x87\x86\xd4\xa9\xc4\x32\x16\xeb\xcf\x71\x62\xf1\x90\x4a\x85\xd2\x7d\x99\xf2\xcf\x95\x0a\x76\x94\x0c\x21\x04\x20\x0a\xe9\x52\x39\x0f\x71\x5f\x86\x48\x94\x16\x43\x5b\x05\x20\x82\x7f\xc5\x2a\x00\x11\x0c\x69\x24\xfb\x7c\x95\x19\xda\xec\x4b\xf1\x77\xbb\x51\x29\x12\x7c\xe7\x45\xd4\x25\xdf\x95\x21\x2a\xd1\x77\x64\xea\x92\x8f\x32\xc6\x57\x1d\x5d\x25\x96\xd1\x98\x72\x04\x56\x4a\xe5\x6c\xa9\x34\x82\xf1\x71\xba\x7e\xbd\x88\x31\x99\x7d\xd3\x13\x63\xc5\x82\x79\x22\x89\x3a\x38\x8d\x4e\x18\x28\x1a\xca\x0f\x71\x83\x95\x08\x83\x44\x24\x11\x71\x44\x07\x1c\xd1\x81\x0d\x51\x68\x64\x32\x8d\x0c\x4d\x27\x92\x68\x44\x62\x3f\x93\x8f\x35\x07\x78\xb0\x10\x25\xcc\x87\x95\xf0\x10\x2e\x14\x7b\x1a\x44\xe9\xef\xf2\x4d\xdd\x81\x3e\x52\xbe\x48\x10\xf5\x97\x5c\xbe\xd6\xec\xe7\xc1\xe7\xd1\x04\x52\xb9\x18\x56\x32\x44\x62\x58\x88\x10\x64\x12\x21\x9d\xf0\x35\xf8\x2d\x8d\xc5\xa2\x79\x4b\x14\x4a\x58\xc2\x43\xbc\xdd\x18\x2a\xb1\x0c\x2f\x12\xf1\x69\x10\xd7\x81\x47\xa4\x22\x64\x1c\xcc\xb5\x75\xc0\x91\x79\x64\x07\x9c\x03\xcc\x85\x70\xf6\x30\xc2\x27\x51\x89\x02\x04\x12\x50\xfb\x5e\xe1\x5b\xf9\x20\x6b\x37\x29\x6f\xa9\x7a\x09\x79\xbb\x31\x3e\x66\x1b\xbe\x94\x27\xe2\x7f\x9d\x69\x34\x7b\x7b\x98\x44\xb2\xb5\x83\x70\x10\x4c\xe4\xe3\x20\xc8\x1e\xc1\x39\x38\xd8\xc2\x38\xae\x3d\x11\x11\x90\x20\x08\x26\x21\xe4\xcf\xa0\x7e\x66\x83\x40\x7e\x72\x91\x50\x24\x81\x23\xfa\xd5\x51\xb7\x85\x2f\xe2\xd3\xf8\x24\x7b\x07\x01\xc5\x1e\xc2\x21\x3c\x1e\x8c\x23\xf3\x1d\xec\x71\x5c\x1e\x09\xc1\x71\x6d\x79\x76\x14\x9e\x9d\x80\x48\xe1\x42\x9f\x11\x43\xd8\x0c\x42\x79\x89\x14\x4a\xa9\x3c\x8a\xf1\xcd\x4c\xee\x4b\x53\x81\xc8\x92\x6f\xa3\x9f\x0b\x22\x44\x7d\x69\x4b\x06\xcb\x15\x88\x7a\x29\x3a\x62\x3f\xaf\x45\xec\x20\x81\x5a\xd3\x97\x1d\x68\x30\x4f\x9d\xf0\x18\xbc\xbe\x09\xc7\xa7\x13\xbe\x89\x7e\x5f\x26\x1a\x3c\x9c\x7f\xad\x0b\x06\xc9\xbf\xcf\x88\x0c\x43\x24\x7f\xb6\x50\xfa\xd5\xfa\xbe\x89\x42\x2a\x50\x46\xc2\x72\xc4\x45\x88\x48\x94\x7f\x6d\x15\x0f\x25\x1c\xd4\xe3\x84\x8f\x5d\xfe\xbf\x30\x14\x0a\x78\xd9\xdf\x1c\x08\x7b\x2a\x0f\xe6\x0a\x1c\x70\x10\x17\x16\xe0\xc8\x7c\xb2\x1d\x0e\xe6\xf2\x48\x38\xb2\xc0\x01\xa2\xc0\x44\x92\x03\x59\x40\xfa\xbb\x03\xf1\x4d\x96\xf8\xbf\x1c\x88\xaf\x00\x5e\x18\x2c\x11\x22\x7c\x06\xe1\xb3\xf0\x73\xe0\x3f\x38\x76\x3c\xa9\x64\x19\x22\xff\x57\x16\x92\x0c\x96\xc3\x62\x44\x89\xc8\x15\x0c\x81\x5c\x2a\xc6\xc0\x32\x59\x84\x88\x07\xab\x55\x84\x65\x12\xfe\xa7\x7d\xea\x4b\x2e\xc3\x28\xa5\x98\x7e\x09\x77\x90\xc7\x7f\xb0\xad\x7c\x44\x2e\xfa\x57\x66\x6a\xbf\xb7\xfc\xd2\x4d\x98\xff\x67\x6d\xfe\xdb\x6b\xf3\xaf\xed\x79\xe8\xda\x1c\xb0\xf1\x11\x06\xef\x7c\x9f\x77\xd3\xc1\x3b\xe5\xe7\x63\xc1\xc7\xf9\xe9\xa1\x9e\x61\x7f\x6d\xc4\xe9\x7d\xa7\xdc\xbf\x93\x52\x07\xc8\x87\x72\xe7\xff\x4f\x0f\x0f\x83\xe4\x43\xb9\x4b\xff\xee\x11\xe5\xbb\x36\x83\x7b\xbe\x5f\x07\xf7\x2f\xfd\x7a\xee\x62\x4a\x23\xa4\x72\x96\x94\x8f\x30\x6c\xe9\x84\xa1\xc2\x43\xaa\xbc\x99\xcc\xd9\x1f\xff\xc3\xc1\x50\x04\x78\xba\x62\xbc\xdd\x99\x76\x10\xd5\xce\x0e\x47\xc2\x43\xfd\x6d\xfa\xd5\xeb\xe7\xa3\xbe\x04\xa8\x4f\x58\x88\x44\xd9\x97\x50\x18\x10\x9d\x30\x28\x36\xb0\x3e\x47\x3d\x21\x22\x96\xf6\x95\xd9\x93\x88\x44\x22\x91\x00\xa9\x7f\x7f\x92\xf6\x2f\x1e\x28\x0d\xf9\x73\x69\xc8\x9f\x48\xbf\x16\x05\x49\x44\x4a\x06\xe9\x93\x64\x40\xb8\x9f\x4a\x7d\x53\xf9\xd8\x7b\x81\x32\x98\x87\xa8\x9b\x36\x30\x34\xb0\xf6\x6c\x91\x0a\x89\xe0\xb8\x89\xc4\x88\x44\xa1\x7e\x07\x87\x4f\x92\x81\xf1\x21\x75\x21\x5f\xcb\xa9\xfd\x75\x21\x83\x74\x1f\x17\x68\xbf\x1b\xf4\xc7\xeb\x39\xe1\xd3\xfd\x9c\xa1\x4b\x27\x7c\xf9\xdb\xc0\x50\x69\xe2\xdf\xff\xa0\x10\x14\x82\x42\x50\x08\x0a\x41\x21\x28\x04\x85\xa0\x10\x14\x82\x42\x50\x08\x0a\x41\x21\x28\x04\x85\xa0\x10\x14\x82\x42\x50\x08\x0a\x41\x21\x28\x04\x85\xa0\x10\x14\x82\x42\x50\x08\x0a\x41\x21\x28\x04\x85\xa0\x10\x14\x82\x42\x50\x08\x0a\x41\x21\x28\x04\x85\xa0\x10\x14\x82\x42\x50\x08\x0a\x41\x21\xff\x66\x88\xee\xd7\x6f\xdb\x23\x12\xbe\x23\x36\x12\xeb\xc4\xe8\x7c\x8b\x85\x00\x00\x18\x9e\x57\x00\x0b\x80\xe8\x29\x00\xc4\xc6\x03\xf0\xb6\x17\x80\xd8\x36\x00\x96\x12\x01\x78\x16\x0a\x00\x2d\x0d\x80\xb1\xd2\xcd\x8b\x2e\x78\x00\x00\xd8\xde\x6e\x2e\x6c\xd5\xbd\x55\xd7\x76\x19\xc4\x3a\x1b\xbb\xd5\x39\xae\xb5\x34\x71\xf1\x3c\xfa\x03\x2b\x34\x37\x21\x3a\x17\xc4\xd1\x7f\x37\x6e\xd6\x2f\x89\xdf\xb9\xe1\xfd\xdd\xdf\x7c\x43\xab\x12\x73\xa7\xd0\xf9\x11\xe3\xe6\xba\x3e\xaa\x58\x59\x38\xc5\xc4\xd0\xd7\x4b\x67\x77\xf5\xbb\xe5\x8f\x70\x3b\xd2\xc2\x9d\xb6\x8c\x59\x51\x70\xa1\xd5\xb0\x0b\x58\xef\x7e\x68\x39\xa1\x66\xcb\x48\x00\x00\xf0\x76\x9f\xe5\x56\xe8\x1a\xba\xf2\xbf\x03\x00\x00\xff\xff\x5c\x4d\xb5\xf4\xcc\x4a\x00\x00")

func arrowPngBytes() ([]byte, error) {
	return bindataRead(
		_arrowPng,
		"arrow.png",
	)
}

func arrowPng() (*asset, error) {
	bytes, err := arrowPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "arrow.png", size: 19148, mode: os.FileMode(420), modTime: time.Unix(1596977096, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _checkboxPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9b\x0b\x6c\x14\xc7\xfd\xc7\xc7\xf9\x03\x7f\x73\xc4\xad\xdb\xb8\x3c\xd2\x44\xd9\x2e\x0d\xe1\xb5\xb7\xaf\x5b\xdf\xed\xea\xbc\xd8\xdc\xe1\xfa\x28\x47\xfc\x02\x6c\x04\x0a\x7b\xbb\x73\xe7\x2d\xb7\x8f\xec\xae\xf1\x99\x96\xca\xc1\x02\x11\xd2\x26\x25\xa2\x11\xe0\x24\x48\x7d\x22\x45\x89\x4a\x93\x80\x1b\xd2\x90\x12\x27\x54\x4d\x05\x75\x0a\x6d\x45\x03\x24\x01\x95\xa8\x4d\x03\xa5\x51\x43\x02\x4d\xaa\xbd\xf3\xd9\xeb\xbb\x33\x32\x75\x9a\xaa\xd2\xac\x64\x5b\x9e\xdf\x7c\xbf\x9f\xdd\xdf\xcc\xfc\x66\xf6\xa4\xbb\xbf\x71\xf9\x57\x2a\x7c\xb7\xfa\x00\x00\x15\xb1\x86\x68\xb3\xfb\x17\x00\xe0\x2b\x9f\x02\x00\x30\xde\x3d\x31\x00\x00\x98\x6a\x36\xb4\xdb\x00\x4c\xbb\xc5\xfd\x29\x03\x7b\x1e\x9d\x09\x40\xd9\x56\x35\x12\x69\x6c\xec\x30\x1c\xc3\xee\x30\x4c\x2c\x16\x89\x60\xa6\x65\x24\xd5\x34\x04\x20\x73\xea\xc9\x1d\xfb\x97\xce\x6e\xaa\xda\x7a\xb8\x6f\x60\x05\xbd\xae\x6d\xf3\xea\x6d\x6b\x1a\x07\x12\xf8\x20\x11\x18\x9c\x53\x4f\x6d\xdf\x12\xdd\xf2\xe5\xd3\x0a\xf3\xca\x96\x98\xb5\x7b\x4d\xd3\x1d\x2f\x1e\xab\x9f\x36\x58\xdf\x74\x4f\xdd\x33\xfb\x7f\x52\x35\x8d\x98\x54\x57\x2e\xf6\x62\xe5\x38\x33\xbd\x65\xd9\xd3\xef\x54\x5f\x5d\x71\xf6\x83\x07\x5e\x7a\xf3\x5a\xe5\x6b\x3f\x6a\x5e\x3e\x69\xdf\xc9\x67\xfb\xf6\xf6\xed\xd8\x92\x88\x1d\x38\xf8\x36\xfe\xab\xbb\x5f\x7b\x6f\xcf\xca\x16\xf8\xc8\x23\xd3\x3f\x60\x6f\x3a\xb3\xb7\x12\x3c\x34\xb5\x65\x5e\x4a\xfb\xe8\x5c\xf7\xc6\xd4\xdf\xcf\xbf\x75\xf4\xc8\x60\xed\xea\x2f\x5e\x7b\xec\xa9\xed\xa7\x9f\xa9\x75\xc0\xca\x1f\xaa\x27\x6b\x43\x3d\xd2\xbc\x55\xbd\xa1\x49\x9f\xef\x5b\xad\x56\x1e\xac\x7d\x9d\xa8\xfa\x2a\xb6\xba\xf6\xf9\x3f\x5f\x5d\x50\x49\xce\xd8\x7e\x67\xf9\x92\x9e\xfe\x43\x0f\x1f\x5f\xbc\xef\xb3\xdb\xbe\xbb\xe9\x8e\xb2\xc4\x8a\xb3\x77\xdd\xb7\xeb\x59\x7d\x25\xb6\xf3\xff\x6e\xfe\x78\x0a\x10\xec\xa7\x2d\x7c\xf7\x52\xe5\xca\xf1\x13\x8f\x56\x2c\xb8\x32\xab\xe9\xd0\x0b\xcf\xdd\x3c\xd0\xf3\x4d\xeb\xd7\xf8\x6e\x2c\xb6\x7d\xd2\x2c\xcd\x50\x7b\xd2\x11\xed\xd0\xe1\x27\xfe\xb2\x3a\xb9\xd3\x4a\xef\xd6\x4e\x9c\xd8\x7a\x60\x6d\xe4\xd8\xef\x9f\x6c\x8a\x3c\x78\xd7\xd1\x87\xc1\x96\x44\xec\xc3\x27\xce\x5e\x2e\xaf\x7d\x2f\x3c\xa9\xfd\x5c\x55\x73\x6f\xe5\x99\x8a\x9f\x5e\x2e\x5b\xb3\x36\x72\x76\xed\x77\x2a\x6e\x8f\xbc\x5f\xb5\xe6\xcc\x82\xdf\x72\xbd\x5b\xcf\xce\xfe\x78\x4f\xdf\xfa\xc3\xaf\x6c\x7a\xe8\x58\xed\x1f\x1f\xb8\xf0\xf3\x8a\xc8\x55\xf1\x1b\x7f\xed\x3e\x73\x61\xf7\x85\x4d\x8f\x45\xfa\x27\x07\x8f\x05\xe7\xec\x53\xdf\xfc\xf6\xf3\x2b\x7f\x09\xc0\xa2\x9d\x6a\x6b\x9b\xd3\x16\x5f\x26\xc8\x86\xe6\x97\x14\x23\x01\xfd\x19\xcd\x04\xee\x15\x5e\x94\x31\x25\x79\x3d\x74\xb0\x04\x4c\xa9\x7a\x0d\x7e\xf1\xd0\x61\x1c\x53\x95\x1a\x7c\x15\x17\xa7\xe2\x66\x04\x76\xa8\x0d\x1b\x2d\xd8\xb2\x71\x79\xab\xbc\x71\xbd\xcc\x2b\xf8\x22\xd1\x17\xce\x08\x19\xcd\xd4\xa0\x23\x61\x19\x2d\xad\xdb\x42\xa6\x06\xcf\xfa\x0a\xba\x2d\xb8\xcd\x24\x8e\x65\xbb\x38\xeb\x6b\xf0\x3a\x37\x80\xb5\xc5\x1b\xb1\x88\x61\x41\x8c\xf3\x57\x13\x32\x4d\xd3\x58\x90\xf7\xd3\x5c\x88\x65\xb8\x85\x18\x43\xd1\x1c\x49\xf1\x24\x4d\x11\x14\x2d\xd0\x94\xc0\x50\xd8\xd0\x85\x8b\x3e\x0c\xc3\xc2\x96\x92\x14\x9a\xa3\xf5\x43\x38\x4b\x49\xd6\xe0\x1d\x8e\x63\x0a\x24\xd9\xd5\xd5\xe5\xef\x62\xfd\x86\x95\x22\x69\x9e\xe7\x49\x8a\x21\x19\x86\xb0\x94\x24\x61\x77\xeb\x8e\x94\x21\x74\x7b\x76\xce\x24\xef\x13\x85\xb6\x6c\xa9\xa6\xa3\x1a\x3a\xe6\xfe\x2f\x25\x8c\x4e\xa7\x06\xc7\x7d\x98\xe7\x1a\x7a\x2e\xcd\x1c\x06\xe9\xf6\x50\xee\x64\x43\x23\x33\x92\x49\xd2\x7e\x8a\x2c\x25\x52\xe4\x61\x8d\xd9\x69\xa5\xb3\xb7\xa6\xc8\x24\x4c\x43\x0d\xea\x8e\x4d\xd2\x7e\xba\xa4\x2e\xa3\x99\xf1\xf8\xf5\x71\x9a\x56\x52\x69\x3b\x4b\x36\x38\xd7\x57\xda\xad\xdd\x26\x24\x9b\xa1\x6d\x74\x5a\x32\x5c\xb2\x01\xea\xce\xec\xd2\x56\xcd\x30\x79\x23\x56\xcd\x30\x59\xd2\xc8\xcc\x57\x80\xd2\x66\xc3\xe1\x31\xd3\xe8\xa8\xc9\x31\x6e\xc4\x8d\x8c\x29\x83\x19\x75\x0c\x99\x1b\xc9\xc9\xc4\x11\x5d\x38\xa3\x99\x42\xc4\x82\x92\x63\x58\xad\x86\x91\x16\x73\xd3\x75\xa4\x7e\x45\x22\xd9\xe9\x89\xcd\x8d\x4b\xb2\xaa\xbb\x8d\xf3\xc2\x64\xa1\xa8\x94\x1f\x8c\x4a\x0e\x14\x19\x8a\xa1\x08\x2a\x44\x50\x7c\x2b\xcd\x0b\x2c\x27\xb0\xd5\x0b\x28\x46\xa0\x28\x8f\x49\xae\x67\x81\x47\x1c\x3a\x92\x22\x39\x52\x29\x97\x6a\x81\xa1\xbd\x2e\xa3\xfa\x16\xfa\x18\x8a\x9a\xec\x1e\x97\xcb\x48\x4f\x8f\x87\x22\x0b\x49\xc3\xd2\x24\x47\x54\x35\x29\x05\x49\x53\x4f\x85\xc9\x91\xc6\xd1\xb4\x78\x5c\x88\xe9\xb6\x23\xe9\x32\x8c\x45\xc5\x8c\x66\xfa\x55\x55\x11\x14\x16\x72\x0a\x45\x73\x04\xe4\x42\x41\x22\xc0\x27\x59\x22\x94\x90\x25\x22\xa4\x28\x90\x81\xb4\x22\xd1\x9c\x92\xbd\x85\xd1\xf2\x22\xeb\xa8\x21\x77\xba\x4b\x28\x16\x15\x73\xd5\x46\x31\x64\x55\x19\x99\x69\x02\x0b\xab\x99\x10\x4b\xb3\x04\x2d\x25\x39\x82\xa6\x83\x90\x48\x04\x82\x3c\x01\x21\x4b\x85\xb8\x20\x43\xd3\x5c\x28\x0f\xf2\x98\x15\x81\xee\xb6\xd4\x94\xaa\x4b\x69\x4f\x1f\xf7\x59\x14\x55\x11\x42\xbc\xcc\xf0\xc1\x84\x42\x28\x81\x40\x35\x11\x90\x25\x96\x90\x78\x96\x23\xaa\x61\x88\x97\xd8\x10\x95\x94\x68\x36\x8f\x28\x61\x53\x84\x6a\x50\x6d\xc7\xb0\xba\xc5\x51\x33\x39\x5b\xa6\x5a\xe0\xbd\xa3\x5b\xf3\x81\xb4\x9a\x2d\x5b\xa6\x64\xd9\xd0\x5d\x8a\x35\x78\x7e\x2d\xe2\x45\x02\x57\x93\xad\x0e\x82\x24\xbb\x05\x4f\x94\xb3\x13\x4e\x09\x93\xa3\x5a\xc7\x96\xa9\xc5\xc3\x39\xbe\x14\x14\xc9\xc7\x66\x74\x75\x40\xfd\x7a\x0b\xc5\xd3\x6b\x6c\x13\xdb\x48\x3a\x5d\x92\x05\xeb\x52\x50\x77\xc6\xb7\x8a\x4b\x09\x8b\x32\x4e\xe6\x52\xfe\x1f\x18\x0a\x5b\xda\x30\xb1\x81\x90\x64\x86\x4b\x24\x14\x86\x48\x04\xa8\x24\x11\x50\xb8\x24\x11\x62\x94\x20\xc1\x48\x12\x1b\x60\xf9\x40\x90\xe6\xf9\x09\x0f\x84\xb7\x4a\xfc\x37\x07\x62\x04\x20\x77\x48\x7a\x0a\x2a\x22\x99\x17\xe6\x1b\x3e\xc5\xb1\x93\x0d\x7d\x03\xb4\x6e\x64\x21\x99\x92\x25\x69\xd0\x81\x96\x2d\x26\x2d\x43\xc3\x24\xd3\x4c\xab\xb2\xe4\xaa\xc8\x0d\xba\x32\xb4\x4f\x0d\xd7\x32\xcc\x31\x30\x4f\xc1\x2d\xf2\xf8\x14\x9f\x55\x81\x96\x7a\x23\x33\xd5\x73\x97\xc3\x69\xc2\xfe\xc7\x9e\x79\xc2\x6b\x73\x7c\x7b\x1e\x5a\x9b\x05\x1b\x1f\x59\xbc\xf3\xe5\x77\xd3\xe2\x9d\x32\x7f\x2c\xc8\xcd\xcf\x7a\x77\x86\x8d\x6f\xc4\xc3\xd9\x53\xee\x44\x4a\x6a\x81\xbc\x94\xbb\xf2\xef\x1e\x1e\x8a\xe4\xa5\xdc\x8d\x89\x1e\x51\xc6\xb4\x29\xce\xbc\x27\xc1\xde\xe8\xc8\xb9\x2b\x62\xa4\x0d\x2b\x6e\x28\x50\x64\xc3\x64\xa9\xe6\x92\xaa\x58\x24\xd2\x98\xfb\x60\x40\x6c\x88\x62\x41\x8a\x27\xea\xbc\x6a\x4f\xd8\x23\x77\xcf\xfe\xee\xc1\x0a\xea\x4e\xb6\x8e\x88\x74\x98\x2c\x6a\x2b\xec\xdf\xe6\xce\x83\x74\x67\x36\x16\x64\x28\x8a\xa2\x48\xda\xfd\x3d\x24\xf5\x86\x0b\xa5\xed\xd7\x97\xb6\x5f\x47\x3a\x12\x5a\xa1\xab\x8e\xc8\x0c\x49\x0a\x9a\x3d\x2a\xf7\x05\x25\x97\xb4\x16\x53\x92\xa1\x58\xcd\x71\x2c\x17\x26\x0b\x9b\x0b\x15\x8d\x6a\x06\xa6\xdb\xa2\xaa\x06\x75\x3b\x9b\x11\x76\x48\x53\x18\x28\x29\x6c\xf7\x08\x29\xaf\xb0\xbd\x48\x98\x5b\x9d\x9e\xd7\xe7\xdc\xbb\x39\x39\xf4\x72\x2e\xfa\xc2\xe4\xf0\x07\x03\xa5\x6a\xc4\x27\x7f\x21\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x7c\xc2\x10\xdf\xc8\x57\xed\xa1\xae\xd4\xe0\x5d\xf8\x22\xf1\xf2\x5e\x87\x01\x00\x60\x72\x43\x73\x1c\x00\x6d\x17\x00\xf6\x83\x00\xfc\x83\x05\xe0\xbe\x5a\x00\x4c\x1d\x80\x73\x09\x00\xe8\x0c\x00\xb7\x58\x3b\x7f\x71\x7e\x19\x00\xe0\xb9\x58\xb4\xae\x35\x73\xea\x5b\x83\xf1\xa5\x8b\x9b\x66\x4e\x79\xe3\xe8\xab\x2f\x4f\x56\xe6\x04\xd6\x7f\x0f\x5b\x7c\xa0\xb7\x5e\x8a\xdf\xf6\x85\x6d\xdb\x5e\x35\x67\xbc\xbb\x78\x1f\x37\xd7\xf7\xcf\x9e\xfe\xcd\xbf\xfb\xcd\xb2\x23\x3b\x4e\x3b\x97\xd7\x56\x36\xf5\xa7\xac\x2b\x6f\x84\x2f\x3d\xfe\xfe\x82\x59\x17\xab\xc2\x33\x32\x2f\xcf\x25\xca\xe7\xef\x7a\xeb\x33\x87\x7e\x00\xfa\x07\x16\xbe\xe8\x23\xf9\xf9\x7b\xcb\x4e\x02\x6d\xe1\xdf\x7a\x3a\x37\xfb\x5f\x68\x5a\x74\xd3\xaa\x5d\xe7\xa6\xdf\xf3\xf6\xba\x8a\xa9\x8f\xdf\xdb\x3b\xf9\xd6\xb2\xaf\xd1\x07\x2f\xdd\x7f\xf4\xb6\xf3\x1d\xdb\x82\x6b\x8e\xbe\x74\xb5\xbf\xef\x4a\xf9\xa5\xaf\xcf\xfc\xd2\x92\xad\xe2\x47\x83\xd3\xc0\xc5\xcc\x91\x3f\x5d\xf9\x7e\xf7\xe7\x66\x4d\xd9\x35\x70\x6d\xc6\xfe\xe0\x8f\x4f\xf1\xaf\xb7\x4c\x6f\x3a\x38\xf8\x87\xe3\x3f\xbb\xfd\x9d\xff\x07\xeb\x3e\xbc\x53\x3a\x78\x71\xbe\x0e\x00\x00\xb1\x25\xcb\xa3\x4f\x2d\x5e\xb7\xf9\x5f\x01\x00\x00\xff\xff\x3d\x80\xda\x3c\x66\x42\x00\x00")

func checkboxPngBytes() ([]byte, error) {
	return bindataRead(
		_checkboxPng,
		"checkbox.png",
	)
}

func checkboxPng() (*asset, error) {
	bytes, err := checkboxPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "checkbox.png", size: 16998, mode: os.FileMode(420), modTime: time.Unix(1596994582, 0)}
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
	"arrow.png": arrowPng,
	"checkbox.png": checkboxPng,
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
	"arrow.png": &bintree{arrowPng, map[string]*bintree{}},
	"checkbox.png": &bintree{checkboxPng, map[string]*bintree{}},
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

