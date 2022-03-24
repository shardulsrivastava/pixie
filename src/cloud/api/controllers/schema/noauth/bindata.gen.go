// Code generated for package noauth by go-bindata DO NOT EDIT. (@generated)
// sources:
// 01_base_schema.graphql
// 02_unauth_schema.graphql
// 03_auth_schema.graphql
package noauth

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

var __01_base_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\x14\x54\x5c\x2a\x10\x4a\x49\x4f\x81\xe0\x07\x1c\x7b\x38\x47\x72\xbc\x3e\xef\x46\x47\x84\xf8\x77\x94\xcb\x5d\x77\xd5\x6c\x31\xf3\xb4\x4f\x43\xe2\xec\xf1\xeb\x80\xd3\xc2\xb6\x0e\xf8\xd8\xc2\x01\xf3\x62\xde\x26\x29\x03\xde\xaf\x97\xfb\x73\xae\xc3\x57\x22\xb4\x32\x20\x0a\xb5\x3c\x1a\x7c\xce\x72\x06\xe7\x6a\x2b\x6c\xad\xd4\x83\xeb\xf0\x29\x38\x13\xa1\xd1\x1b\x51\x7d\x0e\x4c\x92\x23\x9b\x22\xb1\x11\xbe\xc4\xeb\xce\x12\x95\xfb\x0e\x26\x18\xe9\x3a\xf0\xc7\x58\x22\x23\xc6\x15\x62\x89\x0d\xdf\x53\xde\xb9\xc9\xac\xea\xd0\xf7\xc7\xc9\xd2\x32\x1e\x82\xcc\xfd\xb1\xf9\x9a\x4e\xf9\x96\x4f\xdb\x73\xfd\xa4\xba\x50\xfb\xe7\x97\x57\xe7\x36\xf8\xae\x75\xf1\x2c\x22\x75\xc0\x9b\x48\xa6\x2f\x0f\x9b\xd4\xa5\x70\xb3\xbc\xdf\xf9\x0f\x00\x00\xff\xff\x6f\xc4\xb8\xef\x28\x01\x00\x00")

func _01_base_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__01_base_schemaGraphql,
		"01_base_schema.graphql",
	)
}

func _01_base_schemaGraphql() (*asset, error) {
	bytes, err := _01_base_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_base_schema.graphql", size: 296, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_unauth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\xb1\x0a\xc2\x40\x0c\x86\xf7\x7b\x8a\xdc\x56\x5f\xe1\x36\x1d\x84\x0e\x0a\xa2\x9b\x38\x84\x9a\xd6\x60\x2f\x29\x77\xb1\x58\xc4\x77\x17\x0b\xa5\x15\xb7\x9f\xe4\xfb\xbf\x9f\x9e\x46\x72\x05\x1b\x3a\x82\xc3\x83\xd2\x00\x2f\x07\x80\xc9\xb8\xc6\xca\x72\x31\xa5\x3d\x46\x0a\x70\xb4\xc4\xd2\xf8\x55\x80\xf5\x44\x94\x52\xab\x77\x00\x3d\x25\xae\x87\x52\x7a\x36\x3a\xe9\x9d\xa4\xe0\x39\x2f\x9b\x1b\xd5\x96\x50\xbc\x7b\x3b\x37\xce\xfe\xa8\xc6\x79\x36\x8a\x39\xc0\x79\xfa\xf8\xcb\x3f\x3d\x82\x3d\xa5\xcc\x3a\xeb\x1d\x40\x75\x43\x69\xa8\xd5\x66\x79\x34\x8e\x94\x0d\x63\xb7\xcb\x01\xb6\xad\xa2\x7d\x85\x9f\x00\x00\x00\xff\xff\xef\x77\x02\x34\xfc\x00\x00\x00")

func _02_unauth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__02_unauth_schemaGraphql,
		"02_unauth_schema.graphql",
	)
}

func _02_unauth_schemaGraphql() (*asset, error) {
	bytes, err := _02_unauth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_unauth_schema.graphql", size: 252, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_auth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x5a\xdd\x8f\xe3\xb6\x11\x7f\xf7\x5f\x31\x97\x7b\xc8\x2e\xb0\x3d\x04\x45\x13\x14\x7e\xaa\x62\xeb\x72\xea\xee\x7a\x5d\xdb\x7b\x69\x10\x1c\x0e\xb4\x34\xb6\x88\x95\x48\x85\xa4\x76\xd7\x2d\xf2\xbf\x17\x43\x52\x12\x29\x6b\x73\x1f\x2d\xfa\x66\xf1\x63\xe6\x37\x1f\x9c\x19\x0e\x8d\xcf\x06\x45\x01\xe6\xd4\x20\xfc\xa3\x45\x75\x82\x7f\xcf\x00\x5a\x8d\x6a\x0e\xf7\x1a\x55\x26\x0e\xf2\xd5\x0c\x40\xaa\xe3\x1c\xee\xd4\xb1\xfb\xa6\x15\x5b\x34\x86\x8b\xa3\x76\x2b\xbb\xaf\x6e\x36\x31\x46\xf1\x7d\x6b\xd0\xcf\x0f\xdf\x9e\x1e\x0d\xea\x39\xfc\xda\xb3\xf9\x40\x13\x79\xd5\x6a\x83\xea\x82\x17\x73\xc8\x96\xaf\x2e\xe7\xb0\x70\x23\x1d\x67\xbf\xe0\xc7\xd3\x8a\xd5\x78\x21\x58\x8d\x73\xd8\x1a\xc5\xc5\xf1\xe5\xc5\xc4\x26\x9c\x09\x39\x2d\xa4\x10\x98\x1b\x2e\xc5\x39\xcf\x61\x6e\x20\xc8\x13\x65\xf8\x81\xe5\xe6\x82\xf9\x1f\xbb\x53\x83\x73\x48\x82\x2f\x4b\xe2\x26\xeb\x86\x68\x23\x6b\x8d\xcc\x65\xdd\x54\x68\xf0\x82\x8b\xa6\x35\x1d\xec\x2b\xc8\x5b\xa5\xa5\x5a\x4b\x3d\x87\x4c\x98\x2b\x60\x96\xe5\x1c\x92\x60\x4f\x62\xc7\x88\xf8\x55\x87\xfc\x3e\x5b\x76\x34\x2e\xe3\xc5\x1b\xd4\x6d\x75\xc6\xf6\x2d\xc7\xaa\x18\xf3\x3e\xd0\xa0\x97\x20\x58\x9b\x0a\xc3\xcd\xe9\x9a\x8b\xe2\x6a\x06\x00\xa0\xf0\xb7\x96\x2b\x2c\x12\x75\xa4\xc5\xa4\xd0\xe9\xe5\x1f\x3e\x03\x9e\x05\xd2\x61\x9c\x01\xbc\x86\x6d\xae\x78\x63\xea\xa3\x02\x14\x45\x23\xb9\x30\xfa\x0a\x14\x1e\x50\x81\x91\x50\xc8\x5c\x03\x17\x90\x57\xb2\x2d\x58\xc3\xdf\x34\x4a\x1a\x39\x03\xa8\xf8\x23\xbe\xe7\xf8\x44\x70\x6e\xfc\xef\x5b\x34\xac\x60\x86\x39\x23\x77\x2b\x16\x52\x18\x14\x46\x07\x36\xbe\x19\x4d\xd1\x72\x6d\x71\x10\x39\x87\x28\x26\xe6\x66\x27\x48\x6d\xa3\x09\x2f\xd3\x12\x9b\x4a\x9e\xe0\x01\x4f\x7a\x06\x50\xd8\xaf\x1a\x85\xb9\xc6\x13\x31\x58\x86\x03\x31\x9f\x68\x6d\xc0\x26\xda\xe2\xb9\x24\xeb\xac\x63\xc1\x1a\xee\x69\x27\xeb\xec\x8c\xa8\x9b\x0d\xa8\xb9\x45\x9e\xcc\xba\x6a\x8f\x5c\xcc\x00\x1a\xfb\x43\x5f\x3c\x70\x51\xcc\xfd\x30\xd9\xf5\x72\x0e\xbf\xba\x2f\x47\x4e\x21\xc9\xca\xa5\x70\x83\x74\x3e\x2c\x6d\x7f\x14\xaf\x3c\xa1\xf7\xa8\xb4\xf5\xe5\xe1\x88\x0e\x1b\x7c\x20\xd8\xc4\xa4\x16\x52\x1c\xf8\x31\x24\x36\xf0\x76\x73\x23\x04\xdb\xde\x68\x9b\x78\x68\x72\xdd\x88\xf0\x12\x0d\xe3\x15\x16\xe3\xad\xb3\xdf\x67\xb3\x30\x38\xde\xb6\x86\xd1\xb4\x8d\x8f\x0b\x85\xcc\xa0\x0f\x12\x51\xd0\x81\xbf\x15\xd8\x28\xcc\x99\xc1\xe2\x42\x21\xd3\x24\xfb\x37\x7e\x81\x06\xa6\x10\x84\x7c\x82\xdc\x12\x28\xe0\x91\x33\x68\x9e\xbd\xc1\xbf\xb9\x9c\x01\xdc\x37\x05\x33\xf8\x9e\xff\x8b\xdb\xf0\x43\x9a\xf0\xe7\x89\x8e\x53\xb6\x7c\x75\x05\x8f\xc1\xe4\x1c\xd2\x82\x1b\xb6\xaf\xa2\x2d\x13\x91\xd0\x41\x8e\x3c\xe8\xcc\xa1\x00\x96\x48\xe7\x73\xf9\x82\xff\xfd\x28\x65\x85\x4c\x0c\xe4\x9c\x0b\x0d\xae\xd4\x11\x70\xdf\xd3\x3b\x9d\x80\x61\xc6\xb8\xd0\x7d\x22\xe9\x84\x89\x12\xca\xe5\x79\x82\xd9\xa2\x89\x73\xca\x05\x0b\xd2\x4d\x48\x25\x48\x3b\x97\x53\x89\x28\x13\x8f\xdc\xc1\xb9\xc0\x9a\xf1\x2a\xf0\xe0\x03\x57\xda\xac\xc2\x04\x73\x05\x15\x1b\x0d\x5d\x76\x79\x92\xc8\xc4\xf2\xad\x51\xd5\x5c\x93\xfb\xeb\x0b\xca\x88\xbd\x01\xdb\x78\x32\x06\x1c\x4c\x0c\xc4\x43\x1b\xde\xa9\xe3\x85\x54\xc7\x31\x8a\x6c\x39\x70\xbf\x53\xc7\x5e\xb9\x52\x1d\x7b\xc6\x72\x18\x1f\x98\x06\x8b\x89\x4e\x90\xe4\x1d\x3f\x27\xda\x4e\x3e\xa0\x08\x88\x5d\xf6\xbc\x67\x00\x1b\x7c\x94\x0f\x98\x54\x55\xb0\x56\xc7\x8b\x03\x0f\xd8\x60\x2d\x1f\xad\xac\x6f\x95\xac\x49\x9c\x40\x3b\xe1\xd2\x38\x32\x39\xd1\x3e\x19\x2c\xae\x00\x05\x89\x55\xf4\x84\xfa\x91\x51\x30\xba\x82\xdc\xee\x0e\x74\x11\x12\xd5\x93\x8e\xbb\x99\x88\x26\x56\xb7\x2e\x39\x0c\xa4\x46\x0b\xa7\x8e\xcf\x98\xd6\xa7\x49\x90\x91\x7f\x9f\xcd\x6c\x44\xea\x3c\xc3\x46\x24\x0f\x63\x06\x10\x55\x44\x33\x80\xd8\xab\x29\xc2\xf3\xdc\xb4\x2a\x5a\x33\x76\x27\x37\x34\xe4\x6f\x1a\xe0\x3a\x69\x1a\x25\x1f\x03\xc5\x0e\x58\xb2\x65\xba\x66\xa6\xb4\x50\xb2\x65\x3a\x26\xd6\x30\x53\x0e\xdf\xdd\x26\xef\x69\x9f\xc0\x5f\xc8\x9a\x71\x31\xa6\xe8\x2c\xea\x10\xb1\x4a\x47\xca\xe5\x05\x12\x18\x4a\x08\x1e\x17\x25\x82\x50\x6d\x9d\xbf\x5b\xd6\x4c\xb0\xea\x64\x78\xae\xef\x1a\x23\xa9\x34\x8a\x48\x9d\xeb\x7c\x88\x1f\x76\xbb\x91\xad\xda\x22\x8a\x97\xf6\xd9\x7a\xeb\x85\x90\x34\x4d\x60\x7a\xd7\x67\x61\xee\x81\xc6\x15\xc0\x48\xc5\x3e\xfb\x24\xe6\x56\xcf\xe1\x6d\x25\x99\x71\x55\x87\xce\xcf\x8d\xe4\x08\x8d\x08\x3c\x50\xd0\x1f\x8c\xf1\x25\xf4\x26\xcb\x9e\xff\x02\x5f\x44\xef\x7f\x02\x13\x45\x5b\x4f\xd4\xc2\x5b\xc3\x0c\x5a\x06\x49\xba\xfd\x78\xbf\xba\x5e\xdd\xfd\xbc\xf2\x5f\xeb\x74\xb5\xcc\x56\x3f\xf9\xaf\xcd\xfd\x6a\x35\x7c\xbd\x4d\xb2\x9b\x74\xe9\x3f\x76\xe9\xe6\x36\x5b\x25\xbb\x74\x39\xc9\x69\x28\xf2\x1d\xa3\x64\x17\x30\x7a\x0d\x89\x00\x2c\xb8\xf1\xf7\x03\x90\x39\x5d\x1c\x80\x1f\x80\xd9\x94\x02\x25\xd3\x50\xcb\x82\x1f\x38\x16\x60\x4a\x04\xe7\x45\x06\x9f\x0d\xec\x4f\xc0\x85\x46\x45\x3e\x04\x52\x41\x41\x89\x9a\x7e\xe7\x25\x53\x2c\xa7\xea\xe4\x8d\x65\xb2\x2b\x39\x15\xdb\x79\xd5\x16\xa8\xa9\xf6\xb1\x1b\x84\xa5\xf7\x80\xa7\xbd\x64\xaa\x00\x26\x0a\x68\x98\x76\x04\x64\x5d\x33\x51\xd8\xed\x84\x38\x5d\x66\x3b\x07\x17\x34\x56\x98\x0f\x78\x45\x75\x9a\x06\x9d\x97\x52\xa3\x00\x26\xa2\xfb\x0a\xe8\xf6\x78\x44\x4d\x7b\xdf\x74\xb0\x0a\x4e\xa5\x95\x06\x5b\xfe\xbf\xb6\xa0\xa2\x2d\xa6\x64\x06\xb8\x01\x5d\xca\xb6\x2a\x80\x12\x8d\x5d\x44\xac\xbe\xd5\xfe\xa6\x45\x77\x0a\x1a\x14\xa4\x18\x46\x31\xa4\x51\x9c\xac\x6b\xd8\xbe\x93\x62\x9b\xde\xa4\x8b\xdd\x1f\xf8\x03\x15\xc5\xde\x1d\xae\x23\x77\xb8\xfe\xb8\xbe\x5b\xfa\x5f\xdb\xf7\x8b\xee\xd7\x62\x93\xad\x77\xfe\x63\x95\xdc\xa6\xdb\x75\xb2\x48\xbb\xef\xbb\x65\x3a\x9c\xb8\x80\xd5\xb6\xd7\x80\x65\xe5\x8a\xf2\x69\x2c\xa3\xd0\xe9\x3d\x9b\xf2\x46\x90\xf2\x66\x00\x35\x33\x79\x89\x45\x26\x0a\x7c\xb6\xf7\xb8\x4c\x98\x0f\x74\xb9\x21\xff\x9e\x22\x6e\x1d\xbf\x47\xb7\x63\xfb\x11\x28\x72\x19\x72\xb5\x02\x9f\x41\x1e\xac\x62\x0d\xdb\x3b\x4b\x98\x12\x75\x68\x47\x57\x01\x1f\xa4\x22\x35\x1b\xb6\xb7\x28\xec\xad\xd7\x12\xfa\xb9\x44\x53\xa2\xf2\x7e\x43\xce\xc5\x82\xcd\xb4\x0f\x0c\xf9\x01\xd1\x77\x0c\x9f\x78\x55\x41\xcd\x1e\x9c\x95\xbd\x2b\x02\x3e\x63\xde\xda\xc8\x49\x7c\x86\xaf\xe4\x60\x28\x90\x12\xf1\x21\x64\x42\x88\x6f\x74\xaf\x1d\x44\xfd\x30\x69\x1f\x77\x89\x0d\xd4\x70\x90\xaa\x66\x86\x4a\x7b\x77\xf6\x08\x6c\x7f\x10\xb5\x2f\x3b\x9e\x4a\x9e\x97\xd6\xf1\xf7\x88\x02\x1a\xa6\x34\x16\x74\x42\xcf\xdd\x59\xf6\x3e\xef\xfc\x9d\xed\xb7\x46\x36\xd0\x48\xcd\x2d\x5e\x92\xaf\xe7\x99\x85\x57\xfb\x48\xa1\x63\x0c\x84\x8b\xc1\x23\xab\x78\x71\x15\xe8\xa7\x53\xe0\x1b\x9b\xef\xd3\x7e\x3c\x54\xd6\x6b\x48\xaa\x2a\x32\x29\x99\x05\x59\x5e\x06\xd6\x27\x90\xda\xdb\x78\x1b\x69\x37\xf2\x9f\x69\xa5\x06\xed\x81\x40\xb3\x2f\x44\x06\xed\xbd\xa2\x93\x8f\x0a\x02\x5e\x60\xf1\xb9\x66\x7d\x15\xe9\x49\x2a\x10\xd2\xba\x2d\x5d\x1b\x5b\x25\xb0\x00\x65\x91\x38\xcf\x6d\x98\x32\x9c\x55\x70\x61\x54\x8b\x97\xb4\xbc\x87\x74\x71\x60\x95\x46\xba\xc2\x95\x4c\x27\x45\x61\xed\xc3\xaa\x5b\x7b\xdc\xf4\x44\xcd\xb4\x90\xc2\x30\x2e\x50\xd1\x01\x6b\x5d\x5e\x1f\x17\x3f\xd3\x29\xcb\x1f\xd5\x61\x59\x8d\x5a\xb3\x63\x34\xd4\xdd\x3d\xc3\x11\x6d\x98\x32\x0b\xd9\x0a\x63\x8f\xdc\x00\xe5\xfa\xaf\x3a\x7d\x44\xe1\xd4\x3d\x41\xcc\xde\x84\x76\xbc\xc6\x08\x06\xdd\x85\x46\x83\x1d\xc1\xb5\x2c\xbe\x4a\xaa\x56\x7f\xb1\x58\x79\xa7\x46\xdb\xe4\x8b\x75\xea\x1a\x00\x48\xa2\xd1\x6c\x27\x66\xd7\x17\x98\xd2\x87\x8d\xf6\xfe\xee\x1c\x88\xe0\x7c\xb0\xc0\x03\x23\xaf\xb4\x06\xa0\x1c\x26\xa4\x29\xfd\x71\x7a\x10\xf2\x49\x90\xcb\x2f\xb6\x51\xd2\xa6\x7d\x7e\xbd\x86\x12\x59\x65\xca\x13\x6d\x2d\x91\x29\xb3\x47\xe6\x3d\x4b\x61\x8e\xfc\x11\x0b\x4a\xb5\x0a\x8f\x6d\xc5\x14\x70\x61\x50\x51\x79\x6b\xf3\xad\x29\x5d\x0c\xf0\xed\x00\x22\xa7\x50\x37\x52\x14\x84\xc0\x48\xdb\xa3\x43\x6d\xb4\x07\xf1\x2e\x4d\x6e\x76\xef\x7e\x39\x07\xd1\x8a\x00\x86\x0d\x9b\x03\xc5\xdc\x75\x3c\xa9\x7e\x90\xb0\xe6\xcf\x1c\x61\x51\xc9\xd6\x65\x7c\xae\xfd\xf1\xea\xc2\xcb\x20\xc3\x15\xec\x6d\xb4\x13\xdf\x1a\xf8\xad\x45\x75\xb2\xe1\x84\x8e\xa6\x96\x35\x7a\xb3\xf9\x2c\xae\x50\x63\xbd\xaf\x50\xc3\xbb\xdd\x6e\xfd\xad\x86\xef\xbf\xfb\xce\x5b\xbf\xd7\xdf\x34\x78\x1b\xed\x8f\xd2\xf6\x04\xb9\x1e\xb0\x7a\x39\x7e\xda\xac\x17\x9d\x04\x94\x2f\xf6\x0a\xd9\x83\x7e\x63\x09\x94\xb2\x41\x17\x8d\x99\xe9\x4b\x87\x4e\x70\x4b\x37\x27\xa0\x7b\x96\x3f\x50\xa1\xc2\x05\x5a\x91\xe9\xf0\xd7\x14\x5b\xc0\x23\x72\x48\x3c\xce\x65\xb6\x5d\xdc\xad\x56\xe9\x62\x67\x2b\xbc\xb1\x9e\xe9\xc2\x48\xb6\x79\x2a\x51\x8c\x15\xcd\xdd\x48\xa3\x64\x8e\x5a\x53\xe8\xec\x96\x77\x3a\x58\x2f\x93\x9d\x2b\x23\x1d\x5d\xd7\xf6\x71\xf5\x52\x27\xb9\x53\x3b\x0d\x51\xd8\xd2\x74\x84\x99\x38\x81\xb4\xc1\xec\xd0\x2a\x97\x4d\x9d\x1b\x5b\xfa\xa8\x81\xed\x65\xeb\x54\xf0\xe4\xa3\x1e\x37\xa1\x6f\x4a\x35\x86\x72\x2e\xa3\xc7\xf2\xc4\x34\x18\x75\xf2\xfe\xe7\x18\x38\x48\x07\xdb\x54\xeb\xbc\x46\xc8\x27\x12\x98\xc1\x9e\x15\x91\x02\xad\x90\xe9\x50\x23\x8f\x34\x58\xe0\x51\xb1\x62\x30\x70\xa0\xbf\x8a\x3f\x60\x75\x22\xb6\x7b\x0c\x3c\x8e\x78\xd7\xfc\x58\x1a\x1a\xb6\x7d\x14\xef\xaa\x74\xcd\xe8\xac\x96\xfe\xb4\x49\x96\xae\x04\xb7\xd1\x2a\x6c\x9d\xd9\xd3\xde\x30\xad\x4d\xa9\x64\x7b\x2c\xd3\x51\x33\x21\x88\xdf\x41\xd7\x2f\xbe\x77\x74\x91\x2c\x0a\x23\x5d\xc4\x7c\xd7\x9d\x99\x28\xf8\xc5\x3d\xbd\xa8\x97\xd7\xcf\x8e\x7b\xa9\xc3\x1b\xc6\xcb\x33\x67\x37\x71\x85\xc6\x9c\x16\xd3\x93\xe7\x8d\xfb\x2e\xc0\x2a\x59\xad\x2b\x26\xb0\x8f\xeb\xb6\x72\xec\xbf\x5c\x40\xed\xe3\xca\x92\x19\xf6\xe9\xe5\xa2\xad\x57\xb2\x40\xed\x63\xaf\x1d\xc8\x84\x36\xaa\xa5\xdb\x1c\x16\xf1\xa4\xd3\xe9\xed\x79\x46\x68\x14\x3e\x72\xd9\xea\xed\x94\xd2\xcf\xe6\xa3\x7c\x35\x36\x65\xfc\xdc\xe3\x8c\xda\x24\x45\xa1\x50\x47\x79\xc9\xc8\x07\x14\xe7\x57\xd1\xa1\x0f\x68\xb7\x9e\x35\x5e\xb8\x9d\xbb\xe1\xe2\x21\xda\xfb\x1a\x36\x9f\x78\xe8\xb0\xd4\xc7\xef\x1b\x9f\x6a\x9b\x8c\xaf\xb5\x5f\xc8\xa6\x7b\xcc\xf0\x25\x81\xe3\x39\x3f\x43\x61\x2d\xf0\x5c\x75\xab\x43\x04\x8f\x5c\xff\x7d\x7b\xb7\xfa\x1a\x10\xf1\xe3\xcb\x17\x49\x6a\xcb\xaf\x0e\x65\x7c\x6a\xbf\x88\xf9\x0b\xf2\x8f\x9e\x85\xfc\xf1\x88\x45\xef\xef\x8c\xc1\x8b\xa0\x25\x03\x10\x5d\xe8\xed\xe7\x4d\xb6\xba\xff\xe7\xc7\xe4\x76\xf9\xc3\x5f\xba\xa1\x65\xb2\xf9\x39\x5b\xc5\x63\x8b\xbb\xd5\x2e\xc9\x56\xe9\xe6\xe3\x36\xdd\x7d\xfc\x25\xb9\xbd\xd9\x4e\x4f\x4d\xd0\x8b\x17\xec\xd2\xdb\xf5\x0d\x05\x5d\x47\xa4\x3f\x02\xc3\x73\xa5\x7b\x02\x56\x91\xef\xea\x92\xfd\xf9\xfb\x1f\x22\x19\xe3\x16\xd5\x97\xc4\xd0\xe9\x06\x57\xd0\xf8\x76\x16\x3f\x6f\x38\x9e\x6f\x0c\x9a\xd7\xee\xd0\xbd\xd0\x17\x74\xf6\x77\xed\xdd\x3f\x29\xac\xec\xeb\x0b\x09\xae\xdf\x74\x25\xab\x9d\x9b\xac\x57\x83\xee\xf2\xf4\xb5\xda\xc6\x76\x79\x94\xc1\xdd\x8b\x38\x68\x33\x11\x98\x75\xdb\x34\x52\x19\xdd\x37\x77\xa3\xc6\x61\xff\x58\x75\x9e\x77\xe0\x85\x36\x76\xef\x6f\xc3\x63\x9d\x95\x62\x1d\xb6\x25\xd6\xd7\x1f\x37\xe9\x2e\x5d\xed\xb2\xbb\xd5\x50\xa5\x07\xdd\xee\x49\xc1\x1f\x59\xd5\xe2\x79\xa0\x1b\x5e\xf0\xec\xae\xbe\x8b\x1e\x3d\xd2\x6d\xf3\x12\x6b\x16\x76\x5e\xcf\x67\x27\x99\x4e\xea\xf7\xcc\xf0\x5f\x81\xfd\x65\x02\x7a\x24\xc7\xd4\x1a\x2f\xc9\x6b\xa0\xec\x36\xd8\x29\x76\xa2\x51\xc7\xfe\x33\x22\xd7\x84\x27\x1d\x6c\x21\x2f\xf2\xd3\xb6\xcf\x7e\xe3\x27\x8d\xd1\x7f\x1b\xb2\xa5\x4b\xa9\xee\xcd\x35\x4c\xdf\x43\x9b\x74\xf2\x99\xf3\xff\x0d\xf0\x3c\x58\x4e\x81\x06\xc8\x5b\x6d\x64\x9d\x3e\xd3\x49\xb9\xdf\xdc\x84\xae\x1e\x5b\x71\x4a\x9e\xcf\x6a\x83\x8d\x64\x98\x10\xe1\x5c\x82\x09\x01\x26\xf0\xff\x21\xfc\xff\x04\x00\x00\xff\xff\x17\xc3\x7b\x23\x70\x23\x00\x00")

func _03_auth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__03_auth_schemaGraphql,
		"03_auth_schema.graphql",
	)
}

func _03_auth_schemaGraphql() (*asset, error) {
	bytes, err := _03_auth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_auth_schema.graphql", size: 9072, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"01_base_schema.graphql":   _01_base_schemaGraphql,
	"02_unauth_schema.graphql": _02_unauth_schemaGraphql,
	"03_auth_schema.graphql":   _03_auth_schemaGraphql,
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
	"01_base_schema.graphql":   &bintree{_01_base_schemaGraphql, map[string]*bintree{}},
	"02_unauth_schema.graphql": &bintree{_02_unauth_schemaGraphql, map[string]*bintree{}},
	"03_auth_schema.graphql":   &bintree{_03_auth_schemaGraphql, map[string]*bintree{}},
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
