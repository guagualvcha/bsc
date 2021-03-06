// Code generated by go-bindata. DO NOT EDIT.
// sources:
// faucet.html (8.658kB)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x3a\xeb\x93\xdb\xb6\xf1\x9f\xe5\xbf\x62\xc3\x9f\x13\x51\xbf\x3b\x92\x92\x2f\xaf\x91\x48\x75\x6c\x27\xcd\x5c\x67\xea\x64\xea\x64\xda\x4e\x92\x0f\x10\xb1\x12\x71\x07\x02\x34\x00\x4a\xa7\x68\xf4\xbf\x77\x00\x3e\x44\x52\xba\x8b\x5d\xbb\xfe\xa0\xc3\x63\xb1\xbb\xd8\xf7\x82\x8e\x3f\xfb\xee\xc7\xd7\x3f\xff\xfb\xa7\xef\x21\x33\x39\x5f\x3e\x8b\xed\x1f\xe0\x44\x6c\x12\x0f\x85\xb7\x7c\x36\x8a\x33\x24\x74\xf9\x6c\x34\x8a\x73\x34\x04\xd2\x8c\x28\x8d\x26\xf1\x4a\xb3\x0e\xbe\xf5\x4e\x1b\x99\x31\x45\x80\xef\x4a\xb6\x4d\xbc\x7f\x05\xbf\xbc\x0c\x5e\xcb\xbc\x20\x86\xad\x38\x7a\x90\x4a\x61\x50\x98\xc4\xbb\xfd\x3e\x41\xba\xc1\xce\x39\x41\x72\x4c\xbc\x2d\xc3\x5d\x21\x95\xe9\x80\xee\x18\x35\x59\x42\x71\xcb\x52\x0c\xdc\xe4\x1a\x98\x60\x86\x11\x1e\xe8\x94\x70\x4c\x66\xde\xf2\x99\xc5\x63\x98\xe1\xb8\x3c\x1c\xc2\x37\x68\x76\x52\xdd\x1f\x8f\x73\xf8\x2b\x29\x53\x34\x71\x54\xed\x39\x30\xce\xc4\x3d\x64\x0a\xd7\x89\x67\x99\xd5\xf3\x28\x4a\xa9\xb8\xd3\x61\xca\x65\x49\xd7\x9c\x28\x0c\x53\x99\x47\xe4\x8e\x3c\x44\x9c\xad\x74\x64\x76\xcc\x18\x54\xc1\x4a\x4a\xa3\x8d\x22\x45\x74\x13\xde\x84\xdf\x44\xa9\xd6\x51\xbb\x16\xe6\x4c\x84\xa9\xd6\x1e\x28\xe4\x89\xa7\xcd\x9e\xa3\xce\x10\x8d\x07\xd1\xf2\xbf\xa3\xbb\x96\xc2\x04\x64\x87\x5a\xe6\x18\x7d\x19\x7e\x13\x4e\x1d\xc9\xee\xf2\xd3\x54\x2d\x59\x9d\x2a\x56\x18\xd0\x2a\x7d\x6f\xba\x77\xef\x4a\x54\xfb\xe8\x26\x9c\x85\xb3\x7a\xe2\xe8\xdc\x69\x6f\x19\x47\x15\xc2\xe5\x47\xe1\x0e\x84\x34\xfb\xe8\x45\xf8\x65\x38\x8b\x0a\x92\xde\x93\x0d\xd2\x86\x92\xdd\x0a\x9b\xc5\x4f\x46\xf7\x31\x1d\xde\x0d\x55\xf8\x29\x88\xe5\x32\x47\x61\xc2\x3b\x1d\xbd\x08\x67\xdf\x86\xd3\x66\xe1\x1c\xbf\x23\x60\x95\x66\x49\x8d\xc2\x2d\x2a\xc3\x52\xc2\x83\x14\x85\x41\x05\x07\xbb\x3a\xca\x99\x08\x32\x64\x9b\xcc\xcc\x61\x36\x9d\x7e\xbe\xb8\xb4\xba\xcd\xaa\x65\xca\x74\xc1\xc9\x7e\x0e\x6b\x8e\x0f\xd5\x12\xe1\x6c\x23\x02\x66\x30\xd7\x73\xa8\x30\xbb\x8d\xa3\xa3\x59\x28\xb9\x51\xa8\x75\x4d\xac\x90\x9a\x19\x26\xc5\xdc\x5a\x14\x31\x6c\x8b\x97\x60\x75\x41\xc4\xd9\x01\xb2\xd2\x92\x97\x06\x07\x8c\xac\xb8\x4c\xef\xab\x35\xe7\xbf\xdd\x4b\xa4\x92\x4b\x35\x87\x5d\xc6\xea\x63\xe0\x08\x41\xa1\xb0\x46\x0f\x05\xa1\x94\x89\xcd\x1c\xbe\x2e\xea\xfb\x40\x4e\xd4\x86\x89\x39\x4c\x4f\x47\xe2\xa8\x11\x63\x1c\x55\xa1\xea\xd9\x28\x5e\x49\xba\x77\x3a\xa4\x6c\x0b\x29\x27\x5a\x27\xde\x40\xc4\x2e\x04\xf5\x00\x6c\xe4\x21\x4c\x34\x5b\xbd\x3d\x25\x77\x1e\x38\x42\x89\x57\x31\x11\xac\xa4\x31\x32\x9f\xc3\xcc\xb2\x57\x1f\x19\xe0\xe3\x01\xdf\x04\xb3\x17\xcd\xe6\x28\xce\x66\x0d\x12\x83\x0f\x26\x70\xfa\x69\x35\xe3\x2d\x63\xd6\x9c\x5d\x13\x58\x93\x60\x45\x4c\xe6\x01\x51\x8c\x04\x19\xa3\x14\x45\xe2\x19\x55\xa2\xb5\x23\xb6\x84\x6e\xc0\x6b\xe3\x5d\x36\x6b\x38\x89\x28\xdb\xd6\x17\xe9\x0c\x07\x77\x7a\x9c\xed\x6f\xa1\x1e\xc8\xf5\x5a\xa3\x09\x3a\xb7\xe8\x00\x33\x51\x94\x26\xd8\x28\x59\x16\xed\xfe\x28\x76\xab\xc0\x68\xe2\x95\x8a\x7b\x75\x88\x77\x43\xb3\x2f\xea\xcb\x7b\xed\x55\xa5\xca\x03\x2b\x7b\x25\xb9\x07\x05\x27\x29\x66\x92\x53\x54\x89\x77\xeb\xf0\xec\x65\xa9\xe0\x15\x13\x44\xa4\x08\x6f\x73\xa2\x0c\xbc\xce\x08\x13\x40\x28\xb5\x36\x19\x86\x61\x87\xb6\x33\xd0\x73\xee\x82\x95\x11\x27\xa8\x51\xbc\x2a\x8d\x91\x2d\xe0\xca\x08\x58\x19\x11\x50\x5c\x93\x92\x1b\xa0\x4a\x16\x54\xee\x44\x60\xe4\x66\x63\x13\x58\xc5\x77\x75\xc8\x03\x4a\x0c\xa9\xb7\x12\xaf\x81\x6d\x14\x45\x74\x21\x8b\xb2\xa8\x55\x55\x2d\xe2\x43\x41\x04\x45\x6a\x15\xcb\x35\x7a\xcb\x1f\xd8\x16\x21\x47\x78\xf5\xe6\xd5\x68\xa8\xf5\x94\x28\x34\x41\x17\xe5\x99\xee\xe3\xa8\x62\xa5\xba\x10\xd4\xff\xe2\x92\x37\x98\xda\x0b\xe4\x28\x4a\xe8\xcd\x02\x65\x43\x87\xb7\x3c\x1c\x14\x11\x1b\x84\xe7\x8c\x3e\x5c\xc3\x73\x92\xcb\x52\x18\x98\x27\x10\xbe\x74\x43\x7d\x3c\xf6\xb0\x03\xc4\x9c\x2d\x63\xf2\x94\x05\x83\x14\x29\x67\xe9\x7d\xe2\x19\x86\x2a\x39\x1c\x2c\xf2\xe3\x71\x01\x87\x03\x5b\xc3\xf3\xf0\x1f\x98\x92\xc2\xa4\x19\x39\x1e\x37\xaa\x19\x87\xf8\x80\x69\x69\xd0\x9f\x1c\x0e\xc8\x35\x1e\x8f\xba\x5c\xe5\xcc\xf8\xcd\x71\xbb\x2e\xe8\xf1\x68\x79\xae\xf9\x3c\x1e\xe3\x88\x2c\xe3\x88\xb3\x65\xbd\xd9\x97\x44\x54\xf2\x93\x49\x44\xd6\x26\x5a\xeb\x75\xce\xe0\xf8\xe9\xb2\x73\xc1\xb6\x37\x41\xcb\x62\xad\x72\xcd\x0c\xde\xe3\x3e\xf1\x0e\x87\xee\xd9\x7a\x37\x25\x9c\xaf\x88\xbd\x7c\xc5\x7f\x7b\xe8\x0f\xb4\xa6\xb8\x65\xda\x15\x43\xcb\x86\x83\x13\xdb\xef\xe9\xac\x83\x00\x64\x64\x31\x87\x9b\x17\x9d\xe8\x73\xc9\x8f\xbf\x1e\xf8\xf1\xcd\x45\xe0\x82\x08\xe4\xe0\x7e\x03\x9d\x13\xde\x8c\x6b\x87\xe8\xf8\xd7\xf0\x50\x60\x63\x6d\xcb\x5a\x1b\xb3\xa7\x0b\x90\x5b\x54\x6b\x2e\x77\x73\x20\xa5\x91\x0b\xc8\xc9\x43\x9b\xb7\x6e\xa6\xd3\x2e\xdf\xb6\x88\x23\x2b\x8e\x2e\x66\x28\x7c\x57\xa2\x36\xba\x8d\x10\xd5\x96\xfb\xb5\x81\x82\xa2\xd0\x48\x07\xd2\xb0\x14\xad\x68\x1d\x54\x47\xf5\xad\x30\x2f\xf2\xbe\x96\xb2\x4d\x05\x5d\x36\x6a\xd4\x9d\xac\xe5\x2d\x63\xa3\x4e\x70\xa3\xd8\xd0\x0f\x0a\xe5\xca\x96\x6a\x8f\x45\xf2\x2a\x68\xd9\xbb\x17\x88\xaa\xaa\x13\xac\xc9\x82\x9b\xc6\x91\xa1\x1f\x41\xd9\x1a\xe1\x8a\x68\x7c\x1f\xf2\x2e\x63\x9f\xc8\xbb\xe9\xc7\xd2\xcf\x90\x28\xb3\x42\x62\xde\x87\x81\x75\x29\x68\xe7\xfe\xaf\xde\xbc\xfa\x58\xf2\xa5\x60\x5b\x54\x9a\x99\xfd\xfb\xd2\x47\x7a\x62\xa0\x9a\xf7\x59\x88\x23\xa3\x9e\xb6\xb4\xee\xe4\xa2\x6b\xb7\xa3\x7a\x70\x2a\x37\xdd\x76\x14\xc1\x0f\x5c\xae\x08\x87\xad\x65\x79\xc5\x51\x83\x91\x60\xd3\x22\x98\x0c\x21\x2d\x95\x42\x61\x40\x1b\x62\x4a\x0d\x72\xed\x56\xd7\xae\x04\xb0\xe7\xb7\x44\x01\x31\x06\xf3\xc2\x40\x52\x17\x4b\x76\x4d\xa3\xda\xd6\x25\xa0\x9d\xda\x18\xdd\xdb\x6f\x3c\x0f\x12\xf8\xf5\xf7\xc5\xb3\x9a\x95\xef\x70\xcd\x04\x02\xb1\xc2\x48\x6d\xc1\x07\x26\x23\x06\x52\x85\xc4\xa0\x86\x94\x4b\x5d\xaa\x8a\x43\x9b\x68\xc0\x72\xd9\x60\x6a\x30\xdb\x8d\xc2\x51\x6b\x90\xf8\x19\xd1\xd9\xa4\xae\xf5\x14\x9a\x52\x89\xd3\x5e\xb3\x3e\x5a\x4b\x05\xbe\x45\xc0\x92\xe9\x02\x58\xdc\xe0\x0d\x39\x8a\x8d\xc9\x16\xc0\xae\xae\x5a\xe0\x11\x5b\x83\xdf\x40\xfc\xca\x7e\x0f\xcd\x43\x68\xa9\x40\x92\x40\x97\x9a\x23\x58\xe3\xd1\x05\x67\x29\xfa\xec\x1a\x66\x93\x45\xb3\xbb\x52\x48\xee\x9b\x59\x1d\xa2\xab\x3f\xee\xf7\xb8\xe8\x4b\xc6\x09\xbf\x27\x9b\x2a\x01\x68\x20\xb0\x61\xda\x40\xa9\xb8\x95\x8e\x85\xab\x54\xd0\x2a\xc4\xc1\x75\xa5\x72\x96\x98\xea\x41\x9d\x2e\x9a\x2b\x54\x68\x42\x8d\x82\xfa\x7f\x7b\xfb\xe3\x9b\x50\x1b\xc5\xc4\x86\xad\xf7\xfe\xa1\x54\x7c\x0e\xcf\x7d\xef\xff\x6c\xc5\x35\xf9\x75\xfa\x7b\xb8\x25\xbc\xc4\x6b\xa7\xef\xb9\xfb\x3d\xa3\x72\x0d\xf5\x70\x0e\x7d\x82\xc7\xc9\x64\x71\x39\x59\x76\x12\xb8\x42\x8d\xc6\xb7\x80\x6d\x4e\x1b\xca\x88\x40\x8e\x26\x93\xd4\xca\x41\x61\x2a\x85\xc0\xd4\x40\x59\x48\x51\x8b\x04\xb8\xd4\xfa\x64\x88\x0d\x44\x72\x6e\x14\x35\x7c\x02\x02\x77\xf0\x4f\x5c\xbd\x95\xe9\x3d\x1a\xdf\xf7\x77\x4c\x50\xb9\x0b\xb9\x4c\x89\x3d\x60\x9b\x16\x23\x53\xc9\x21\x49\x12\xa8\x5b\x38\x6f\x02\x7f\x01\x6f\xa7\x6d\x33\xe7\xc1\xdc\x0e\xed\x68\x02\x57\x30\x3c\x9e\x49\x6d\xe0\x0a\xbc\xa8\x72\x2c\x9b\x17\x95\x89\x48\xc1\xbc\x49\xe5\x1b\x8d\x16\xa4\xc8\x51\x6b\xb2\xc1\x2e\xb7\xb8\x45\x61\x5a\x8b\xb3\x97\xca\xf5\x06\x12\x70\xda\x2a\x88\xd2\x58\x81\x84\x36\x3e\x37\xa6\x67\x0d\xd8\x81\x25\x09\x88\x92\xf3\x93\xc5\x56\x1e\xb2\x68\x6c\xb1\x07\x1e\xba\xa8\x09\x9f\x25\x09\xd8\x70\x65\xe5\x4d\x4f\x27\xad\x25\x54\x61\x75\x12\xda\x88\x79\x3a\x31\x59\x74\x4d\xbb\x87\x0d\xe9\x9f\xa1\x43\x3a\xc4\x87\xf4\x11\x84\x2e\x8b\x3d\x85\xaf\xca\x7a\x1d\x74\x6e\xe1\x11\x6c\xa2\xcc\x57\xa8\x9e\x42\x57\x65\xb1\x1a\x9d\x13\xf5\xad\x30\x9d\xb3\xd7\x30\xfb\x7a\xf2\x08\x76\x54\x4a\x3e\x8a\x5c\x48\xb3\xf7\x0f\x9c\xec\x65\x69\xe6\x30\x36\xb2\x78\xed\xd2\xce\xf8\x1a\x2c\xad\x39\xb4\x18\xae\x5d\xc7\x30\x87\xb1\x9b\xd9\x7d\x96\xa3\x3b\xf5\xd5\x74\x3a\xbd\x86\xa6\x9f\x7e\x45\xac\x47\xaa\x12\x8f\x8f\xf0\xa3\xcb\x34\xb5\x7d\xf7\xc7\x70\x54\xe3\x68\x79\xaa\xe7\x1f\xc1\x55\x9b\x28\x7a\x6c\xc1\x17\x5f\xc0\xd9\x6e\xdf\x8c\xa3\x08\xfe\x4e\xd4\x3d\xb8\x0a\x53\xe1\x96\xc9\x52\x9f\xd2\x4e\xce\xb4\x66\x62\x03\x44\x03\x95\x02\xeb\x33\x1f\x96\x03\xce\x78\xac\xc1\x60\x09\xd3\x21\x83\x36\x36\x76\x72\xc4\x85\xd4\xd1\xc1\xdb\xcf\x0a\x8d\x44\x2e\x24\x1d\x96\x23\x7c\x96\x80\xe7\x75\x0f\x9f\x41\x58\x80\x16\xd9\x48\xa3\xf9\xb9\xd2\x85\x5f\xa7\xca\x4b\x89\x6c\x72\x6d\x0b\xe7\xe9\xe4\x8c\x89\xe3\x49\xbc\x2f\x8b\x02\x05\x05\x22\xf6\x2e\x3e\xb6\xb2\x65\xc2\x48\xb0\x9d\xb4\x8d\x6f\xdc\x36\x05\x1c\x5d\xa4\xaa\x8f\x5a\x01\xa7\x32\xcf\xa5\x80\x04\x82\xd9\xe2\x42\x4a\xed\x48\xb2\x73\xb5\xa1\x7a\x2e\xc8\x7e\xa8\xa2\xbe\xcc\x06\xc0\xc1\xac\xa7\x94\x9e\xbe\x2e\x2b\x66\xd4\xf2\xcd\x4e\x12\x1d\xa8\xeb\xa4\xaf\xa1\xcc\x3a\xfc\x57\x78\xae\x66\xef\x79\x8d\x76\xbb\x28\x75\xe6\x0f\x18\x9d\x2c\xce\x75\x73\x6b\x50\x11\x83\xae\x33\x72\xba\x40\x61\x98\xc2\x33\x95\x00\x11\xb6\x84\x0a\x14\x0a\x8a\xaa\xa9\x2f\x6c\x63\x55\x75\x41\x3d\x95\xb9\xe7\xf2\x9e\x39\x7d\xa0\xc3\xb8\xfa\x4c\x0a\xb4\x8d\xf3\xc0\x09\x9c\xa1\xf6\x2c\xd5\x02\x23\x27\x85\x46\x0a\x09\x54\xcf\x9b\xfe\x24\x2c\x05\x7b\xf0\x27\x41\x3d\x1f\xe2\x68\xf6\xeb\xb4\xe9\x34\x56\xb1\x7d\x95\x80\x17\x1b\x65\x6b\xef\xb1\x07\x57\x97\x5c\xd0\xa6\xe0\xf1\xf2\xc4\x41\xf7\x28\x40\x6c\xe8\xd2\xf5\x75\x55\x4f\xf0\x9b\x67\x3b\xf0\x8d\x92\xa5\xa0\x73\x5b\x77\xf9\x67\x68\xc9\x96\x18\xa2\x1c\xd6\xc9\x02\x4e\xe0\xae\x51\x9f\x43\x6a\x95\xb3\x80\xaa\xf5\x73\xdd\x35\xb4\x4d\xab\x9b\xad\xa4\xa2\xa8\x02\x45\x28\x2b\xf5\x1c\xbe\x2c\x1e\x16\xbf\x35\x4d\xbd\x6b\x14\x9e\x64\xb5\x50\xb8\x3c\xe3\x28\x4d\xdd\xdb\xcb\x15\x78\x71\x64\x01\xfe\x0c\x4d\x7b\xd9\xee\xb3\x2a\x5c\x68\x87\xa0\x7d\xf4\xac\xd7\x73\x46\x29\x47\xcb\xf0\x09\xbd\x75\x46\xab\xff\xae\x4b\xf5\x49\x42\xdd\x07\x9d\xce\x1c\x01\xb9\xc6\x27\x0e\xb4\x2d\xd5\xd8\x1a\x40\x60\xaf\xcc\x9c\xcc\xeb\xee\xcc\x2d\xab\xb1\x93\x45\xfd\x48\x4e\x4b\xe5\x0a\x2f\x3f\xa8\x0d\xec\x1a\xc6\xda\x16\x82\x54\x8f\x27\x61\x56\xe6\x44\xb0\x3f\xd0\xb7\x79\x69\x52\xc9\xca\xf5\x68\xde\x79\x48\x3e\x63\xe6\xd4\xf8\x8f\x9b\x1c\x37\xae\x85\x38\x6e\xb4\x6b\x15\xd9\x79\x60\x1e\x7f\xa0\x84\x2e\x53\x09\x56\x44\x41\x77\x12\x34\xc9\x17\x94\xb4\xd4\x9b\xbd\x15\x51\xe3\xaa\x3f\x75\xc5\xba\x90\xbb\x64\x7c\x33\x6d\x99\xac\x14\xed\xf4\x3c\xae\x6d\xed\x4c\x19\x96\xcb\xc6\x35\x97\x70\x33\xfd\x14\xdc\x52\x22\x36\x38\xbc\x81\x51\xac\x40\x0a\x24\x35\x6c\x8b\xff\x83\x8b\x7c\x02\x21\x7f\x30\x8b\xd6\x0e\x1b\xe1\x39\x33\xed\xf1\x6b\x77\x5b\xd9\xfe\xbf\xf5\x37\x88\x9c\x84\xaf\xc0\xbb\x78\x91\x47\x2d\x71\x00\x38\x70\xed\xc7\xfd\xde\x3d\x3a\x78\xc3\x9c\x62\xab\xdd\xf6\xb9\x6c\x12\x66\x26\xe7\xbe\x17\x1b\xf7\xf9\xc3\xf2\xdc\x62\x70\x08\xaa\xe5\x7e\x49\x77\xec\x37\x32\xb6\x99\xc7\x41\xd3\x05\x9d\xe2\xa4\x6d\xcc\x9a\x4a\x04\x8e\xa7\xaf\x44\x51\x04\x6f\x0d\x51\x06\x08\xfc\x72\x0b\x65\x41\x89\xb1\xd9\x4b\x82\xcd\x8f\x2e\x8b\xb5\x9f\x91\x56\x44\x69\x58\x4b\xb5\x23\x8a\x42\x29\x0c\xe3\x76\x7f\x0f\x44\x61\x5b\xfa\x69\x34\xb7\x36\x8a\x6d\x09\xf7\xcf\x9a\xc0\xe7\xfe\x38\xec\xaa\x7c\x3c\x09\x91\xa4\xd9\x39\xa0\xcb\x58\x2d\xdd\x04\xde\xb8\x16\xc0\x7f\xee\x9b\x8c\xe9\x49\x48\x8c\x51\xfe\xb8\x67\x0c\xe3\x89\xd5\xeb\xac\xd3\x92\xb5\xc7\xe3\x9e\x5b\x3d\x85\xe3\x54\x4c\xb7\x85\x40\x03\x9e\x6a\xed\x57\x76\x35\xbe\xee\xe0\xee\x9b\xd5\xf8\xf3\x71\xab\xa8\x93\x7b\x9f\xee\x91\x5c\xe4\xa4\x87\x7a\x6c\xbd\x6c\x7c\x46\x9e\x50\xfa\xda\xfa\x8f\xef\x5d\xf0\xf4\xa1\x75\x4c\x5a\x61\x57\xf1\xfa\x49\x29\x33\x41\xf1\xe1\x31\x11\x33\x3a\x9e\x84\xba\x5c\x55\x0f\x15\xfe\x57\x6d\x03\xd6\x80\x39\xe3\x1d\xa6\x82\xb3\x82\xc2\x92\xe8\x17\x15\xc1\xa0\x08\x79\x22\x6b\xd4\x24\xab\x5b\x1d\xaf\xad\xc0\xa7\x93\xf6\x9d\xeb\x7b\x6d\x8b\x2b\xa6\x33\x20\xb0\xc3\x95\x76\xcf\x0a\x50\xdb\xbb\x7b\xda\xa9\x9e\x70\x5e\xfe\x74\xdb\x79\xc6\x69\x3d\xc2\x77\xd8\xdb\x2f\xbc\x97\x1e\x4d\x2e\x7e\x52\xde\xed\x76\xe1\x46\xca\x0d\xaf\x3e\x26\xb7\xaf\x2a\x11\x29\x58\x78\xa7\x3d\x20\x7a\x2f\x52\xa0\xb8\x46\xb5\xec\xa0\xaf\x9f\x5a\xe2\xa8\xfa\xd8\x19\x47\xd5\xff\xe0\xf8\x4f\x00\x00\x00\xff\xff\xed\xb2\xd3\xa7\xd2\x21\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x53, 0xb1, 0xc1, 0x2e, 0x1, 0xc9, 0x33, 0x7a, 0xbf, 0x66, 0xff, 0xed, 0x61, 0x5e, 0xca, 0x9a, 0x73, 0x31, 0xb4, 0xa, 0x95, 0xc3, 0x1b, 0x6, 0xe8, 0xcf, 0x12, 0x4e, 0xb8, 0x2e, 0xcf, 0xdf}}
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

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
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

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"faucet.html": faucetHtml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
