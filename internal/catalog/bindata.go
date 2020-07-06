// Code generated for package catalog by go-bindata DO NOT EDIT. (@generated)
// sources:
// internal/catalog/data/crystal/0.35.1.json
// internal/catalog/data/crystal/latest.json
// internal/catalog/data/golang/1.13.12.json
// internal/catalog/data/golang/1.14.4.json
// internal/catalog/data/golang/latest.json
// internal/catalog/data/memcached/1.6.5.json
// internal/catalog/data/memcached/1.6.6.json
// internal/catalog/data/memcached/latest.json
// internal/catalog/data/mysql/8.0.17.json
// internal/catalog/data/mysql/latest.json
// internal/catalog/data/node/12.18.1.json
// internal/catalog/data/node/14.4.0.json
// internal/catalog/data/node/latest.json
// internal/catalog/data/postgresql/10.13.json
// internal/catalog/data/postgresql/11.8.json
// internal/catalog/data/postgresql/12.3.json
// internal/catalog/data/postgresql/9.5.22.json
// internal/catalog/data/postgresql/9.6.18.json
// internal/catalog/data/postgresql/latest.json
// internal/catalog/data/redis/6.0.4.json
// internal/catalog/data/redis/latest.json
// internal/catalog/data/ruby/2.6.6.json
// internal/catalog/data/ruby/2.7.1.json
// internal/catalog/data/ruby/latest.json
package catalog

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

var _crystal0351Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\xcd\x6e\x1b\x21\x14\x85\xf7\x7e\x0a\x84\xb2\x2c\x1e\x7e\xc6\x10\xa6\xea\x2a\x0f\x52\x5d\xfe\x6c\x54\x06\x46\x80\x1d\x27\x91\xdf\xbd\xb2\xa7\xf6\x26\x55\xd5\x4d\x96\x47\xe7\xe8\xdc\xef\x48\xf7\x63\x83\x10\xb6\xf5\xad\x75\x48\x78\x42\x57\x89\x10\xee\x6f\x8b\xc7\x13\xc2\xce\xd7\x78\x82\x1e\x4b\xc6\xdf\x56\xe7\xe4\x6b\xbb\xca\x09\x61\xba\x15\xbb\x2d\xbb\x1b\xe6\x18\x93\x7b\x29\xf3\x0c\xd9\x5d\xdd\xf9\x97\x8b\x15\x91\x05\x3d\x95\x63\xff\xde\xa1\x22\x42\x5a\xaf\x71\x21\xb6\xcc\x4b\xc9\x3e\xf7\xf6\x83\x21\xf2\x72\x0b\x20\x72\x0e\xe8\xe9\xa3\x55\x7b\xb9\x17\x1e\x6b\x9a\x60\x76\x72\x9c\x1c\xd4\xd7\x78\x3b\x79\xe8\x7d\x69\xd3\x30\xec\x63\x3f\x1c\xcd\xd6\x96\x79\xf8\xc3\x4e\x12\xe4\xfd\x5d\x0c\xd5\x27\x0f\xcd\xb7\xc1\x95\xd7\x9c\x0a\xb8\x61\x85\x7d\xa4\x57\x49\x18\x59\xbb\xc9\xf9\x59\xfe\x94\xe3\xb6\x43\xdd\xee\xdf\xef\x04\xed\x00\x7c\x27\x3f\x41\x28\xa7\x76\x41\x51\xb9\xa3\x9a\xd2\x00\x3a\x30\x1f\xb4\xe0\x4a\x69\x63\xb9\x08\x4a\x03\xd3\x7a\xe4\xca\x8e\x9c\x69\x4e\xc7\x00\xda\x73\xce\xac\x10\x14\x18\x18\xf9\x79\x60\x8a\xf9\x78\xfe\x9a\x7d\xb7\xea\xff\x98\xf7\x40\x90\x56\x04\x27\x24\x55\xc2\x08\xae\xa9\x12\x94\x19\x0a\x1a\x3c\x04\xa7\xec\xb3\x13\x5e\x1b\x6d\x40\xfa\x91\x8f\xe0\x35\x33\x40\xed\x8e\x51\xe9\x6e\xdb\x83\xc6\x1b\x84\x2e\xd7\x0b\x38\x45\x13\x6d\xc9\xa7\xeb\x5b\x3d\x3e\x2a\xe6\x83\xaf\xb1\xe3\x35\x52\x16\x9f\x5b\x4b\xff\x48\xbc\xa7\x68\xfe\x6e\x6f\x2e\x9b\xdf\x01\x00\x00\xff\xff\xd1\x62\xf6\x8f\xbd\x02\x00\x00")

func crystal0351JsonBytes() ([]byte, error) {
	return bindataRead(
		_crystal0351Json,
		"crystal/0.35.1.json",
	)
}

func crystal0351Json() (*asset, error) {
	bytes, err := crystal0351JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "crystal/0.35.1.json", size: 701, mode: os.FileMode(420), modTime: time.Unix(1593045091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _crystalLatestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\xcd\x6e\x1b\x21\x14\x85\xf7\x7e\x0a\x84\xb2\x2c\x1e\x7e\xc6\x10\xa6\xea\x2a\x0f\x52\x5d\xfe\x6c\x54\x06\x46\x80\x1d\x27\x91\xdf\xbd\xb2\xa7\xf6\x26\x55\xd5\x4d\x96\x47\xe7\xe8\xdc\xef\x48\xf7\x63\x83\x10\xb6\xf5\xad\x75\x48\x78\x42\x57\x89\x10\xee\x6f\x8b\xc7\x13\xc2\xce\xd7\x78\x82\x1e\x4b\xc6\xdf\x56\xe7\xe4\x6b\xbb\xca\x09\x61\xba\x15\xbb\x2d\xbb\x1b\xe6\x18\x93\x7b\x29\xf3\x0c\xd9\x5d\xdd\xf9\x97\x8b\x15\x91\x05\x3d\x95\x63\xff\xde\xa1\x22\x42\x5a\xaf\x71\x21\xb6\xcc\x4b\xc9\x3e\xf7\xf6\x83\x21\xf2\x72\x0b\x20\x72\x0e\xe8\xe9\xa3\x55\x7b\xb9\x17\x1e\x6b\x9a\x60\x76\x72\x9c\x1c\xd4\xd7\x78\x3b\x79\xe8\x7d\x69\xd3\x30\xec\x63\x3f\x1c\xcd\xd6\x96\x79\xf8\xc3\x4e\x12\xe4\xfd\x5d\x0c\xd5\x27\x0f\xcd\xb7\xc1\x95\xd7\x9c\x0a\xb8\x61\x85\x7d\xa4\x57\x49\x18\x59\xbb\xc9\xf9\x59\xfe\x94\xe3\xb6\x43\xdd\xee\xdf\xef\x04\xed\x00\x7c\x27\x3f\x41\x28\xa7\x76\x41\x51\xb9\xa3\x9a\xd2\x00\x3a\x30\x1f\xb4\xe0\x4a\x69\x63\xb9\x08\x4a\x03\xd3\x7a\xe4\xca\x8e\x9c\x69\x4e\xc7\x00\xda\x73\xce\xac\x10\x14\x18\x18\xf9\x79\x60\x8a\xf9\x78\xfe\x9a\x7d\xb7\xea\xff\x98\xf7\x40\x90\x56\x04\x27\x24\x55\xc2\x08\xae\xa9\x12\x94\x19\x0a\x1a\x3c\x04\xa7\xec\xb3\x13\x5e\x1b\x6d\x40\xfa\x91\x8f\xe0\x35\x33\x40\xed\x8e\x51\xe9\x6e\xdb\x83\xc6\x1b\x84\x2e\xd7\x0b\x38\x45\x13\x6d\xc9\xa7\xeb\x5b\x3d\x3e\x2a\xe6\x83\xaf\xb1\xe3\x35\x52\x16\x9f\x5b\x4b\xff\x48\xbc\xa7\x68\xfe\x6e\x6f\x2e\x9b\xdf\x01\x00\x00\xff\xff\xd1\x62\xf6\x8f\xbd\x02\x00\x00")

func crystalLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_crystalLatestJson,
		"crystal/latest.json",
	)
}

func crystalLatestJson() (*asset, error) {
	bytes, err := crystalLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "crystal/latest.json", size: 701, mode: os.FileMode(420), modTime: time.Unix(1593045091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golang11312Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\x4b\x8e\xc3\x20\x10\x04\xd0\x3d\xa7\x68\xb1\x1e\x81\xf9\xd8\x80\x2f\x63\xe1\x6e\x8c\x47\xf2\x0c\x11\x38\x59\x24\xf2\xdd\x23\xe7\xb3\xac\x2a\xbd\x7a\x30\x00\x9e\xcb\xa4\x26\x65\xf8\x08\x67\x04\xe0\xb7\x54\xdb\x6f\xf9\xe7\x23\x70\x25\x94\x11\x4a\xf3\x9f\xf7\x72\xad\xdb\xd9\xae\xfb\x7e\x69\xa3\x94\xb4\x89\x5c\x4a\xde\x92\xc0\xf2\x27\x73\x91\xb9\x7c\x80\x68\x15\xc5\x1e\xab\xc8\xf7\xaf\x6d\x6b\xd4\xfd\xf0\x3a\x75\x73\xd4\x68\x29\x59\x72\xde\x05\x13\xb5\x42\x1c\xfa\x40\x21\x74\x6e\xb1\xa6\x1f\x90\x02\x7a\x4a\x7e\x76\xd4\xf9\x10\x90\x12\x61\x5a\x9c\xd2\x69\x8e\xc6\x72\x06\x70\xb0\x83\x3d\x03\x00\x00\xff\xff\x7d\x6b\xa2\x20\xbe\x00\x00\x00")

func golang11312JsonBytes() ([]byte, error) {
	return bindataRead(
		_golang11312Json,
		"golang/1.13.12.json",
	)
}

func golang11312Json() (*asset, error) {
	bytes, err := golang11312JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang/1.13.12.json", size: 190, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golang1144Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcb\x4b\x8e\xc3\x20\x10\x04\xd0\x3d\xa7\x68\xb1\x1e\xf1\x73\x63\x33\xbe\x0d\x34\x18\x8f\xe4\x49\x47\xe0\x64\x91\xc8\x77\x8f\x2c\x2b\xbb\x52\x55\xbd\xb7\x00\x90\x95\xe5\x0c\x67\x02\x90\xcf\xd2\xfa\x1f\xdf\xe4\x0c\xd2\x2a\x8b\x0a\xe5\xcf\x35\x3c\xda\x76\x96\xeb\xbe\xdf\xfb\xac\x75\xde\x54\x65\xae\x5b\x51\xc4\xff\xba\xb2\xae\x7c\xfd\x55\x6f\xa4\xf6\xd8\x54\x7d\x7d\x69\x5f\xa3\xf3\xe3\xa9\x27\x63\x6d\x5c\x86\x94\xc8\x45\xb2\x26\x64\x9b\x82\x2b\x31\xc4\x94\xc2\x94\x5c\x19\x87\x65\x0a\x01\x71\x31\xce\xff\xa6\xe2\x0c\xe5\x82\x19\x1d\x79\x42\xe3\x03\x4a\x01\x70\x88\x43\x7c\x02\x00\x00\xff\xff\xc0\x84\x2b\xb1\xb7\x00\x00\x00")

func golang1144JsonBytes() ([]byte, error) {
	return bindataRead(
		_golang1144Json,
		"golang/1.14.4.json",
	)
}

func golang1144Json() (*asset, error) {
	bytes, err := golang1144JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang/1.14.4.json", size: 183, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangLatestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcb\x4b\x8e\xc3\x20\x10\x04\xd0\x3d\xa7\x68\xb1\x1e\xf1\x73\x63\x33\xbe\x0d\x34\x18\x8f\xe4\x49\x47\xe0\x64\x91\xc8\x77\x8f\x2c\x2b\xbb\x52\x55\xbd\xb7\x00\x90\x95\xe5\x0c\x67\x02\x90\xcf\xd2\xfa\x1f\xdf\xe4\x0c\xd2\x2a\x8b\x0a\xe5\xcf\x35\x3c\xda\x76\x96\xeb\xbe\xdf\xfb\xac\x75\xde\x54\x65\xae\x5b\x51\xc4\xff\xba\xb2\xae\x7c\xfd\x55\x6f\xa4\xf6\xd8\x54\x7d\x7d\x69\x5f\xa3\xf3\xe3\xa9\x27\x63\x6d\x5c\x86\x94\xc8\x45\xb2\x26\x64\x9b\x82\x2b\x31\xc4\x94\xc2\x94\x5c\x19\x87\x65\x0a\x01\x71\x31\xce\xff\xa6\xe2\x0c\xe5\x82\x19\x1d\x79\x42\xe3\x03\x4a\x01\x70\x88\x43\x7c\x02\x00\x00\xff\xff\xc0\x84\x2b\xb1\xb7\x00\x00\x00")

func golangLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_golangLatestJson,
		"golang/latest.json",
	)
}

func golangLatestJson() (*asset, error) {
	bytes, err := golangLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang/latest.json", size: 183, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _memcached165Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\xd1\x0e\x82\x20\x14\x80\xe1\x7b\x9e\xe2\x8c\xeb\x42\x90\x04\xf4\x6d\xe0\x70\x50\x37\xcd\x06\xd6\x45\xcd\x77\x6f\xda\xaa\xdb\x6f\xfb\xff\x17\x03\xe0\x33\xcd\xe8\x71\xa0\xc8\x3b\xd8\x01\x80\x3f\x28\x97\x71\xb9\xf2\x0e\xb8\x12\x46\x34\xfc\xf4\xf1\x7b\x9e\x76\x1b\xd6\xf5\x56\xba\xaa\xfa\x95\x62\xc9\x7d\x95\xc6\x89\xca\xdf\xce\x47\x29\x56\x9f\x45\xff\xfc\x0e\xca\xe0\xeb\xc6\x1c\xdf\x74\x89\x5e\x5b\x69\x12\x2a\x8d\x5a\x07\x6a\x63\x6a\x6d\xa8\x51\xa1\x8b\x36\x48\xd7\xaa\xd8\x24\x19\xd1\x39\x4f\x01\x8d\xd4\x18\x94\x75\x64\x5c\xa8\x6d\x4c\x9c\x01\x6c\x6c\x63\xef\x00\x00\x00\xff\xff\xa8\xff\x7b\x65\xc3\x00\x00\x00")

func memcached165JsonBytes() ([]byte, error) {
	return bindataRead(
		_memcached165Json,
		"memcached/1.6.5.json",
	)
}

func memcached165Json() (*asset, error) {
	bytes, err := memcached165JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "memcached/1.6.5.json", size: 195, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _memcached166Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\xcd\xae\x82\x30\x10\x40\xe1\x7d\x9f\x62\xd2\xf5\xbd\xa5\xf4\x0f\xca\xdb\xd4\x61\x06\x48\x40\x4c\x8b\x2e\x34\xbc\xbb\x41\xa3\x6e\xbf\xe4\x9c\x87\x00\x90\x0b\x2d\x98\x70\xa4\x5e\x76\x70\x00\x80\xbc\x51\x2e\xd3\x7a\x96\x1d\xc8\x5a\x05\x15\xe4\xdf\xdb\xaf\x79\x3e\x6c\xdc\xb6\x4b\xe9\xaa\xea\x5b\xaa\x35\x0f\x15\x4f\x33\x95\x9f\xfd\xbf\x4a\xb5\xa5\xac\x86\xfb\x67\x50\xc6\x64\x7c\x38\x1e\x51\xb7\xac\x89\x90\x93\xf7\xb1\x36\x11\x23\x39\x47\x3d\xba\xc0\xda\x50\x63\x13\x53\xcb\x09\x93\xf5\xfe\xe4\x88\xd1\x63\x1b\xfa\xa8\x0d\xa3\x25\x6b\xb8\x91\x02\x60\x17\xbb\x78\x06\x00\x00\xff\xff\xe4\x45\x91\x6f\xc3\x00\x00\x00")

func memcached166JsonBytes() ([]byte, error) {
	return bindataRead(
		_memcached166Json,
		"memcached/1.6.6.json",
	)
}

func memcached166Json() (*asset, error) {
	bytes, err := memcached166JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "memcached/1.6.6.json", size: 195, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _memcachedLatestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\xcd\xae\x82\x30\x10\x40\xe1\x7d\x9f\x62\xd2\xf5\xbd\xa5\xf4\x0f\xca\xdb\xd4\x61\x06\x48\x40\x4c\x8b\x2e\x34\xbc\xbb\x41\xa3\x6e\xbf\xe4\x9c\x87\x00\x90\x0b\x2d\x98\x70\xa4\x5e\x76\x70\x00\x80\xbc\x51\x2e\xd3\x7a\x96\x1d\xc8\x5a\x05\x15\xe4\xdf\xdb\xaf\x79\x3e\x6c\xdc\xb6\x4b\xe9\xaa\xea\x5b\xaa\x35\x0f\x15\x4f\x33\x95\x9f\xfd\xbf\x4a\xb5\xa5\xac\x86\xfb\x67\x50\xc6\x64\x7c\x38\x1e\x51\xb7\xac\x89\x90\x93\xf7\xb1\x36\x11\x23\x39\x47\x3d\xba\xc0\xda\x50\x63\x13\x53\xcb\x09\x93\xf5\xfe\xe4\x88\xd1\x63\x1b\xfa\xa8\x0d\xa3\x25\x6b\xb8\x91\x02\x60\x17\xbb\x78\x06\x00\x00\xff\xff\xe4\x45\x91\x6f\xc3\x00\x00\x00")

func memcachedLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_memcachedLatestJson,
		"memcached/latest.json",
	)
}

func memcachedLatestJson() (*asset, error) {
	bytes, err := memcachedLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "memcached/latest.json", size: 195, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql8017Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\x41\x6e\x84\x20\x14\xc6\xf1\x3d\xa7\x78\x61\x5d\x01\x4d\x05\x74\xdd\x65\xbb\x68\x7a\x02\x2a\x8d\xa2\x08\x15\x1c\x47\x9c\x78\xf7\x89\x9a\x59\xbe\xf7\xcf\xf7\x7b\x20\x00\x3c\xa6\x38\x59\xc9\x70\x0d\xc7\x09\x80\x97\xbf\x10\x8d\x77\xb8\x06\x2c\x09\x23\xb9\xc0\x6f\x57\xb8\x05\x7b\x3c\xbb\x79\xfe\x8f\x35\xa5\x8d\x76\xe4\x1c\x93\xc6\x8f\x94\x7e\xf8\xbb\xb3\x5e\xe9\x48\xbf\xd2\xcf\xf7\x67\x26\x09\xa3\x67\xce\x2e\x85\xcc\x2a\x90\x76\x7b\x61\xb1\x53\x45\xc9\x0f\x2f\x1f\xfb\x60\x57\x27\x17\xd3\x1a\x5e\x05\x16\xb8\x78\xef\x0b\x6d\x7e\xb5\x1b\x94\x0a\x2c\xef\x4d\x59\xf1\xb2\x8d\x69\x15\x03\x67\x9b\x98\x12\xc7\x08\x60\x47\x3b\x7a\x06\x00\x00\xff\xff\x12\xbb\x1f\xe4\xc2\x00\x00\x00")

func mysql8017JsonBytes() ([]byte, error) {
	return bindataRead(
		_mysql8017Json,
		"mysql/8.0.17.json",
	)
}

func mysql8017Json() (*asset, error) {
	bytes, err := mysql8017JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/8.0.17.json", size: 194, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysqlLatestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\x41\x6e\x84\x20\x14\xc6\xf1\x3d\xa7\x78\x61\x5d\x01\x4d\x05\x74\xdd\x65\xbb\x68\x7a\x02\x2a\x8d\xa2\x08\x15\x1c\x47\x9c\x78\xf7\x89\x9a\x59\xbe\xf7\xcf\xf7\x7b\x20\x00\x3c\xa6\x38\x59\xc9\x70\x0d\xc7\x09\x80\x97\xbf\x10\x8d\x77\xb8\x06\x2c\x09\x23\xb9\xc0\x6f\x57\xb8\x05\x7b\x3c\xbb\x79\xfe\x8f\x35\xa5\x8d\x76\xe4\x1c\x93\xc6\x8f\x94\x7e\xf8\xbb\xb3\x5e\xe9\x48\xbf\xd2\xcf\xf7\x67\x26\x09\xa3\x67\xce\x2e\x85\xcc\x2a\x90\x76\x7b\x61\xb1\x53\x45\xc9\x0f\x2f\x1f\xfb\x60\x57\x27\x17\xd3\x1a\x5e\x05\x16\xb8\x78\xef\x0b\x6d\x7e\xb5\x1b\x94\x0a\x2c\xef\x4d\x59\xf1\xb2\x8d\x69\x15\x03\x67\x9b\x98\x12\xc7\x08\x60\x47\x3b\x7a\x06\x00\x00\xff\xff\x12\xbb\x1f\xe4\xc2\x00\x00\x00")

func mysqlLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_mysqlLatestJson,
		"mysql/latest.json",
	)
}

func mysqlLatestJson() (*asset, error) {
	bytes, err := mysqlLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/latest.json", size: 194, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _node12181Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8c\x5d\x4e\xc5\x20\x10\x46\xdf\x59\xc5\x84\x67\x2f\x30\x03\x97\x42\x37\x63\xa0\xd0\x1f\x63\x5a\x03\xd8\x18\x4d\xf7\x6e\x6a\xed\xe3\x9c\x33\xdf\xf9\x61\x00\x7c\xdd\x52\x7e\xab\x0f\xa4\xd7\x2f\xde\xc3\x89\x00\xf8\x9e\x4b\x5d\xb6\x95\xf7\xc0\x91\x04\x3a\x81\xfc\xe5\x32\x9f\xe5\xfd\xa4\x73\x6b\x1f\xb5\x97\xf2\x5a\x8b\xad\x4c\x32\x2d\xb5\xc9\xfd\xff\xfd\x4f\x3c\xee\x4b\xb4\x50\xc4\xf4\x7d\x37\xea\x1c\xe8\x69\xcf\x4c\xcc\x4f\x4d\xce\x1b\x44\x4a\xca\x66\xdb\xe1\x10\x75\xf2\xda\xb9\x4c\x41\x8f\x81\x06\xeb\x55\x32\x2a\xa6\x31\xf8\x6c\x34\x79\xc2\xe0\xb3\xea\x5c\xd4\x23\x19\xe4\x0c\xe0\x60\x07\x63\xbf\x01\x00\x00\xff\xff\xd2\x46\x21\x44\xcb\x00\x00\x00")

func node12181JsonBytes() ([]byte, error) {
	return bindataRead(
		_node12181Json,
		"node/12.18.1.json",
	)
}

func node12181Json() (*asset, error) {
	bytes, err := node12181JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node/12.18.1.json", size: 203, mode: os.FileMode(420), modTime: time.Unix(1592875528, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _node1440Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8c\x4d\x0e\x83\x20\x10\x46\xf7\x9c\x62\xc2\xba\x05\x41\x18\xd4\xdb\xf0\xab\x36\x8d\x36\x40\x5d\xb4\xf1\xee\x8d\xb5\x5d\xbe\x37\xf3\xbd\x37\x01\xa0\xcb\x1a\xe2\xad\xd0\x01\x0e\x02\xa0\x5b\xcc\x65\x5e\x17\x3a\x00\x15\x8a\x29\xd6\xd0\xcb\x79\x78\xe6\xfb\x21\xa7\x5a\x1f\x65\xe0\xfc\xdc\xb1\x35\x8f\x3c\xcc\xa5\xf2\xed\xfc\xfe\xfa\xeb\x0f\x58\xb5\x99\x8d\xaf\x7f\xa1\x4c\x56\x6a\x3c\x22\xda\x60\xef\x23\xfa\x84\xa2\x8f\x6d\x6a\xbb\x60\x5d\x1f\x50\x45\xeb\x64\xa3\x8c\x94\x88\xda\x75\xae\x49\xda\x05\xa9\xb1\xb3\x5e\x18\x11\x65\x0c\x1e\x83\x33\xc9\x50\x02\xb0\x93\x9d\x7c\x02\x00\x00\xff\xff\x13\x41\xee\x74\xc2\x00\x00\x00")

func node1440JsonBytes() ([]byte, error) {
	return bindataRead(
		_node1440Json,
		"node/14.4.0.json",
	)
}

func node1440Json() (*asset, error) {
	bytes, err := node1440JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node/14.4.0.json", size: 194, mode: os.FileMode(420), modTime: time.Unix(1592875528, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeLatestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8c\x4d\x0e\x83\x20\x10\x46\xf7\x9c\x62\xc2\xba\x05\x41\x18\xd4\xdb\xf0\xab\x36\x8d\x36\x40\x5d\xb4\xf1\xee\x8d\xb5\x5d\xbe\x37\xf3\xbd\x37\x01\xa0\xcb\x1a\xe2\xad\xd0\x01\x0e\x02\xa0\x5b\xcc\x65\x5e\x17\x3a\x00\x15\x8a\x29\xd6\xd0\xcb\x79\x78\xe6\xfb\x21\xa7\x5a\x1f\x65\xe0\xfc\xdc\xb1\x35\x8f\x3c\xcc\xa5\xf2\xed\xfc\xfe\xfa\xeb\x0f\x58\xb5\x99\x8d\xaf\x7f\xa1\x4c\x56\x6a\x3c\x22\xda\x60\xef\x23\xfa\x84\xa2\x8f\x6d\x6a\xbb\x60\x5d\x1f\x50\x45\xeb\x64\xa3\x8c\x94\x88\xda\x75\xae\x49\xda\x05\xa9\xb1\xb3\x5e\x18\x11\x65\x0c\x1e\x83\x33\xc9\x50\x02\xb0\x93\x9d\x7c\x02\x00\x00\xff\xff\x13\x41\xee\x74\xc2\x00\x00\x00")

func nodeLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_nodeLatestJson,
		"node/latest.json",
	)
}

func nodeLatestJson() (*asset, error) {
	bytes, err := nodeLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node/latest.json", size: 194, mode: os.FileMode(420), modTime: time.Unix(1592875528, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgresql1013Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcd\x5b\xae\x83\x20\x10\xc6\xf1\x77\x56\x31\xe1\xf9\x1c\x2e\x36\x6a\x74\x33\x0d\xde\xc0\x84\x0a\x0e\xa8\xc1\xc6\xbd\x37\xda\x34\xbe\xfe\x32\xf3\xff\xde\x04\x80\x7a\x17\xa2\xc6\x3e\xcc\xf6\x29\x05\xad\xe1\x44\x00\xba\xf6\x18\x46\x37\xd1\x1a\xa8\x14\x4c\x3e\xe8\xdf\xd7\x17\xb4\xa7\x99\x18\x7d\xa8\x39\x1f\xa2\x67\x77\x81\x39\xd4\xdc\x2f\x0d\x0f\x6e\xc1\xb6\xe7\xeb\xf5\xca\xef\x83\xff\x0b\x58\x54\xc8\xf4\xfe\x4b\x06\xa3\xb2\xbc\xb8\x96\x66\x65\x45\xf2\xa5\xaa\x44\xda\x93\x2a\x8d\xcd\x0b\x1d\x5e\x5b\x3e\xac\xaa\xf5\x16\x3b\xe3\xa7\xd2\x54\x7a\x6a\x12\xca\x31\x1b\xd3\x96\x75\x94\x00\x1c\xe4\x20\x9f\x00\x00\x00\xff\xff\x09\x01\xe8\x75\xcd\x00\x00\x00")

func postgresql1013JsonBytes() ([]byte, error) {
	return bindataRead(
		_postgresql1013Json,
		"postgresql/10.13.json",
	)
}

func postgresql1013Json() (*asset, error) {
	bytes, err := postgresql1013JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgresql/10.13.json", size: 205, mode: os.FileMode(420), modTime: time.Unix(1594056234, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgresql118Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\x5b\x6e\xc3\x20\x14\x44\xff\x59\xc5\x15\xdf\x2d\x94\x3e\xac\xda\x9b\xa9\x30\x75\xc0\x4f\xe0\x5e\xfc\x4a\xe4\xbd\x47\x24\x8a\xfc\x39\x67\x34\x67\x6e\x0c\x80\x07\x4f\xc9\x62\x43\x71\xf8\x53\x8a\x57\x90\x21\x00\x5f\x1a\xa4\xd6\x4f\xbc\x02\xae\x94\xf8\xe5\x6f\x4f\x3c\xe3\x90\x91\x4b\x29\x50\x25\xe5\x25\x05\x71\x0a\x84\x47\x2b\xc3\x5c\x4b\xf2\x33\x9a\x46\x2e\x79\x29\xcf\xfe\x3d\x67\x91\x34\x0a\x7b\x7d\x09\xc9\xe9\xcf\x9f\xe2\x71\x13\x7b\x8a\xbb\xde\xb7\x71\x32\x66\xac\x75\xb0\x5f\xba\xa3\xb5\x0c\x9d\x8d\x8b\xfe\x68\xa7\xad\x73\x5b\xf1\x8d\xf1\xbf\x30\xbd\x43\x5b\xae\x41\x73\x06\x70\xb0\x83\xdd\x03\x00\x00\xff\xff\x8a\x59\xad\xe5\xca\x00\x00\x00")

func postgresql118JsonBytes() ([]byte, error) {
	return bindataRead(
		_postgresql118Json,
		"postgresql/11.8.json",
	)
}

func postgresql118Json() (*asset, error) {
	bytes, err := postgresql118JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgresql/11.8.json", size: 202, mode: os.FileMode(420), modTime: time.Unix(1594056241, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgresql123Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\x4b\x6e\xc3\x20\x10\xc6\xf1\xbd\x4f\x31\x62\xdd\xf2\x9a\x0c\x0f\x5f\xa6\x02\x62\x9c\x4a\x51\xa1\x80\xb3\x68\x95\xbb\x57\xa8\x8a\xbc\x9c\xff\xe8\xfb\xfd\x2e\x00\xac\x96\x3e\xf6\xb6\xf5\xef\xfb\x87\xd2\x6c\x85\x19\x01\xd8\x63\x6b\xfd\xb3\x7c\xb1\x15\x98\xd2\x1c\xd9\xdb\x7f\x3e\xda\x7d\xa6\xdb\x18\xb5\xaf\x42\xe4\x51\xf9\x09\xf0\xd2\x76\x51\x8f\x28\x7a\x39\x5a\xda\xc4\x63\x2e\xc5\xf9\x7f\x9f\x37\x1f\xa1\xf1\xfd\xe7\x05\xf6\x5b\xd0\x64\xa6\x69\xa5\xcb\x57\x8a\xa8\x83\xb7\x64\xad\xb1\xfe\x8a\x49\xa1\xd3\x97\x64\x10\x3d\x9a\xec\x9c\x09\x16\x31\x27\xa2\x10\x29\x78\x7d\x91\xd1\x28\x0a\x4a\x52\x26\xc9\x16\x80\xe7\xf2\x5c\xfe\x02\x00\x00\xff\xff\xb9\x40\x86\xcf\xd6\x00\x00\x00")

func postgresql123JsonBytes() ([]byte, error) {
	return bindataRead(
		_postgresql123Json,
		"postgresql/12.3.json",
	)
}

func postgresql123Json() (*asset, error) {
	bytes, err := postgresql123JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgresql/12.3.json", size: 214, mode: os.FileMode(420), modTime: time.Unix(1594056473, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgresql9522Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\x51\x6e\x84\x20\x10\xc6\xf1\x77\x4e\x31\xe1\xb9\x05\xa5\x81\x14\x2f\xb3\xc1\x56\x91\x85\x08\x3b\x83\x64\xb3\x8d\x77\x6f\x8c\xd9\xf8\xfa\x9f\xf9\x7e\x7f\x0c\x80\x97\x4c\xd5\xe3\x44\x8f\x74\xb3\x37\xcd\x07\x38\x2a\x00\x6f\x13\x52\xc8\x2b\x1f\x80\x5b\xa1\x85\x52\xfc\xe3\x3c\x6c\x98\x8e\xb8\xd4\x5a\x68\x90\x72\xae\x45\x5c\x86\xc8\xe8\x65\xd9\x46\x49\x79\xc3\x9f\x49\xb6\x73\x2b\xaf\x8f\xcf\xb3\x88\xea\x50\xf8\xd7\x1b\xa5\xc5\x29\x6d\x0e\xb7\x6b\xae\x04\xdd\x79\x13\xba\xc7\xa2\xc7\xdf\x6f\x6a\xb3\xb2\xee\x6e\xee\x7d\xd2\xcf\xaf\x11\x63\x1f\xd7\x1e\xd7\x48\x31\x98\xd4\x66\x7a\xbe\x2c\x67\x00\x3b\xdb\xd9\x7f\x00\x00\x00\xff\xff\x0b\xd6\x15\x11\xd1\x00\x00\x00")

func postgresql9522JsonBytes() ([]byte, error) {
	return bindataRead(
		_postgresql9522Json,
		"postgresql/9.5.22.json",
	)
}

func postgresql9522Json() (*asset, error) {
	bytes, err := postgresql9522JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgresql/9.5.22.json", size: 209, mode: os.FileMode(420), modTime: time.Unix(1594056224, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgresql9618Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\x4b\x6e\xc3\x30\x0c\x44\xf7\x3e\x05\xc1\x75\xab\x8f\x23\xd1\xa2\x2f\x13\xf8\x43\x39\x05\x82\x4a\x95\xe4\x2c\x5a\xe4\xee\x85\x61\x04\xd9\xbe\x99\x79\xf3\xd7\x01\x60\x4e\xb5\x6d\x45\xea\xcf\xfd\xca\x57\xc2\x11\x0e\x0a\x80\x0f\x29\xf5\x2b\x7d\xe3\x08\xc8\x8a\x94\x0d\xf8\x71\x06\x7b\xb9\x1f\xf0\xd6\x5a\xae\xa3\xd6\xb1\x65\xf5\x76\xa8\x54\x36\x9d\xf7\x59\xd7\xb4\x97\x45\xf4\xe3\xdc\xea\x77\xe3\xf3\x24\xaa\x4d\x45\x6d\xbf\x2f\x69\xbd\x4d\xbd\x3f\xde\xd1\x8b\x31\xf3\xea\x44\x58\xb8\x77\xc3\x65\x15\x36\xd1\x4f\x8b\xf1\xeb\x10\x02\x13\xb3\x84\x78\xb1\x83\x9f\x28\xfa\x99\x8d\xb7\x8e\x2c\x0d\xce\x52\xe8\xa3\x5b\x08\x3b\x80\x67\xf7\xec\xfe\x03\x00\x00\xff\xff\x5e\xbe\x5a\x0c\xdd\x00\x00\x00")

func postgresql9618JsonBytes() ([]byte, error) {
	return bindataRead(
		_postgresql9618Json,
		"postgresql/9.6.18.json",
	)
}

func postgresql9618Json() (*asset, error) {
	bytes, err := postgresql9618JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgresql/9.6.18.json", size: 221, mode: os.FileMode(420), modTime: time.Unix(1594056225, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgresqlLatestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\x4b\x6e\xc3\x20\x10\xc6\xf1\xbd\x4f\x31\x62\xdd\xf2\x9a\x0c\x0f\x5f\xa6\x02\x62\x9c\x4a\x51\xa1\x80\xb3\x68\x95\xbb\x57\xa8\x8a\xbc\x9c\xff\xe8\xfb\xfd\x2e\x00\xac\x96\x3e\xf6\xb6\xf5\xef\xfb\x87\xd2\x6c\x85\x19\x01\xd8\x63\x6b\xfd\xb3\x7c\xb1\x15\x98\xd2\x1c\xd9\xdb\x7f\x3e\xda\x7d\xa6\xdb\x18\xb5\xaf\x42\xe4\x51\xf9\x09\xf0\xd2\x76\x51\x8f\x28\x7a\x39\x5a\xda\xc4\x63\x2e\xc5\xf9\x7f\x9f\x37\x1f\xa1\xf1\xfd\xe7\x05\xf6\x5b\xd0\x64\xa6\x69\xa5\xcb\x57\x8a\xa8\x83\xb7\x64\xad\xb1\xfe\x8a\x49\xa1\xd3\x97\x64\x10\x3d\x9a\xec\x9c\x09\x16\x31\x27\xa2\x10\x29\x78\x7d\x91\xd1\x28\x0a\x4a\x52\x26\xc9\x16\x80\xe7\xf2\x5c\xfe\x02\x00\x00\xff\xff\xb9\x40\x86\xcf\xd6\x00\x00\x00")

func postgresqlLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_postgresqlLatestJson,
		"postgresql/latest.json",
	)
}

func postgresqlLatestJson() (*asset, error) {
	bytes, err := postgresqlLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgresql/latest.json", size: 214, mode: os.FileMode(420), modTime: time.Unix(1594056473, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _redis604Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\xdb\x6a\x85\x30\x10\x85\xe1\xfb\x3c\xc5\x90\xeb\x36\x07\x27\xe6\xe0\xdb\x8c\xc9\xb4\x0a\x62\x4a\x62\x5b\x68\xf1\xdd\x37\xba\xd9\x97\xeb\x83\xf5\xff\x0b\x00\xd9\xb8\xac\x5d\x4e\x70\x0d\x00\xf9\xc3\xad\xaf\x75\x97\x13\x48\xaf\x8c\x72\xf2\xed\xe9\xdf\x6d\xbb\x6c\x39\x8e\xaf\x49\xeb\x52\x7f\xf7\xad\x52\x51\xf7\x5b\xad\x55\x37\xde\x98\x3a\x77\x7d\xcb\xfb\xfd\x55\x07\x35\xf5\xf9\xf7\x4a\xf4\x85\x86\xd1\x5f\x15\x44\x0c\xc6\x8c\x64\xd9\x64\x24\x1a\x12\xe6\x18\x32\x5a\x74\x3e\x30\x45\xca\xd6\xa6\xe8\xd2\x60\x3f\x68\x36\x31\x9a\x90\x52\x9c\x29\xf8\x31\xa7\xe4\xb0\xb0\x14\x00\xa7\x38\xc5\x23\x00\x00\xff\xff\xc4\x02\x07\xeb\xc1\x00\x00\x00")

func redis604JsonBytes() ([]byte, error) {
	return bindataRead(
		_redis604Json,
		"redis/6.0.4.json",
	)
}

func redis604Json() (*asset, error) {
	bytes, err := redis604JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "redis/6.0.4.json", size: 193, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _redisLatestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\xdb\x6a\x85\x30\x10\x85\xe1\xfb\x3c\xc5\x90\xeb\x36\x07\x27\xe6\xe0\xdb\x8c\xc9\xb4\x0a\x62\x4a\x62\x5b\x68\xf1\xdd\x37\xba\xd9\x97\xeb\x83\xf5\xff\x0b\x00\xd9\xb8\xac\x5d\x4e\x70\x0d\x00\xf9\xc3\xad\xaf\x75\x97\x13\x48\xaf\x8c\x72\xf2\xed\xe9\xdf\x6d\xbb\x6c\x39\x8e\xaf\x49\xeb\x52\x7f\xf7\xad\x52\x51\xf7\x5b\xad\x55\x37\xde\x98\x3a\x77\x7d\xcb\xfb\xfd\x55\x07\x35\xf5\xf9\xf7\x4a\xf4\x85\x86\xd1\x5f\x15\x44\x0c\xc6\x8c\x64\xd9\x64\x24\x1a\x12\xe6\x18\x32\x5a\x74\x3e\x30\x45\xca\xd6\xa6\xe8\xd2\x60\x3f\x68\x36\x31\x9a\x90\x52\x9c\x29\xf8\x31\xa7\xe4\xb0\xb0\x14\x00\xa7\x38\xc5\x23\x00\x00\xff\xff\xc4\x02\x07\xeb\xc1\x00\x00\x00")

func redisLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_redisLatestJson,
		"redis/latest.json",
	)
}

func redisLatestJson() (*asset, error) {
	bytes, err := redisLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "redis/latest.json", size: 193, mode: os.FileMode(420), modTime: time.Unix(1592873462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruby266Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xdd\x8a\x83\x30\x10\x85\xef\x7d\x8a\x30\xd7\xdb\xfc\xa9\x11\x7d\x9b\x4c\x32\x6a\x20\x44\x49\x62\xa1\x5d\xfa\xee\x4b\x5a\xf6\x6e\xd9\xbb\xc3\x37\x67\x3e\x38\xdf\x1d\x63\x90\x2f\x7c\xc0\xc2\x5a\x66\x0c\xee\x94\x4b\x38\x12\x2c\x0c\x34\x37\xdc\xc0\xd7\x87\x5f\x39\x36\xb6\xd7\x7a\x96\x45\x08\x67\xdd\x4e\xbc\xbd\xde\xa2\x4d\x1b\x3f\xf2\x26\xce\x0b\x45\x23\x42\x73\xf3\x0e\xb7\xb7\x81\x57\x9b\xf9\xf6\xfc\x15\x95\xdd\xea\xd1\x34\x57\x6f\x06\x54\x43\xef\x69\xed\x8d\x44\xeb\x14\x4e\x03\xe1\x68\xc8\x1b\x89\xca\x4a\xef\xac\x19\xfa\x19\xa5\x54\xe3\x64\x49\xa9\x75\x9d\x26\x3f\x3a\xaf\x69\xd6\x7a\x56\xd0\x31\xf6\x6a\x56\x88\x01\x83\x3b\xd2\xbd\xcd\x80\xfa\x38\xa9\xe9\x43\xda\x29\x87\x0a\x9f\xca\x71\x52\x2a\x25\xfe\xd3\x78\xc6\x80\x7f\x9f\xbb\x57\xf7\x13\x00\x00\xff\xff\x2c\x71\x74\x04\x2a\x01\x00\x00")

func ruby266JsonBytes() ([]byte, error) {
	return bindataRead(
		_ruby266Json,
		"ruby/2.6.6.json",
	)
}

func ruby266Json() (*asset, error) {
	bytes, err := ruby266JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ruby/2.6.6.json", size: 298, mode: os.FileMode(420), modTime: time.Unix(1593974633, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ruby271Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xc1\x8e\xeb\x20\x0c\x45\xf7\xf9\x0a\xe4\xf5\x2b\x18\x1a\x42\x5e\x7e\xa6\xc2\x89\x95\x20\x21\x12\x01\xa9\xd4\x8e\xfa\xef\x23\x5a\xcd\x6e\x34\x3b\xfb\xf8\xfa\x48\xf7\xab\x13\x02\xf2\x49\x8f\x9b\xb9\x39\x98\x44\xdb\x85\x80\x3b\xe7\x12\xf6\x04\x93\x00\x23\x9d\xd4\xf0\xef\xc3\xcf\x1c\x1b\xdb\x6a\x3d\xca\xa4\xd4\xec\xe7\x8d\x65\x7b\xbf\x44\x9f\x56\xb9\xe7\x55\x1d\x27\xa9\x46\x94\x91\xee\x3d\x5c\xde\x06\x59\x7d\x96\xeb\xf3\x47\x54\x36\x6f\xec\xd0\x5c\x4b\xaf\xc7\x7e\xbc\xd2\xb2\x20\x22\x5a\x37\xcc\xfa\xea\xd0\x3a\xad\x8d\xf6\x03\x93\xe9\xed\x68\xb4\x1e\x16\x42\x72\x44\x06\xd1\xf2\x7f\x64\x63\x91\x3d\xf7\x7a\x84\x4e\x88\x57\xb3\x42\x0c\x14\xe6\x3d\xdd\x5b\x0d\xa8\x8f\x83\x9b\x3e\xa4\x8d\x73\xa8\xf0\x89\xec\x07\xa7\x52\xe2\x1f\x89\x67\x0c\xf4\xfb\xb9\x7b\x75\xdf\x01\x00\x00\xff\xff\x59\xd0\x5e\x39\x2e\x01\x00\x00")

func ruby271JsonBytes() ([]byte, error) {
	return bindataRead(
		_ruby271Json,
		"ruby/2.7.1.json",
	)
}

func ruby271Json() (*asset, error) {
	bytes, err := ruby271JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ruby/2.7.1.json", size: 302, mode: os.FileMode(420), modTime: time.Unix(1593045091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rubyLatestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xc1\x8e\xeb\x20\x0c\x45\xf7\xf9\x0a\xe4\xf5\x2b\x18\x1a\x42\x5e\x7e\xa6\xc2\x89\x95\x20\x21\x12\x01\xa9\xd4\x8e\xfa\xef\x23\x5a\xcd\x6e\x34\x3b\xfb\xf8\xfa\x48\xf7\xab\x13\x02\xf2\x49\x8f\x9b\xb9\x39\x98\x44\xdb\x85\x80\x3b\xe7\x12\xf6\x04\x93\x00\x23\x9d\xd4\xf0\xef\xc3\xcf\x1c\x1b\xdb\x6a\x3d\xca\xa4\xd4\xec\xe7\x8d\x65\x7b\xbf\x44\x9f\x56\xb9\xe7\x55\x1d\x27\xa9\x46\x94\x91\xee\x3d\x5c\xde\x06\x59\x7d\x96\xeb\xf3\x47\x54\x36\x6f\xec\xd0\x5c\x4b\xaf\xc7\x7e\xbc\xd2\xb2\x20\x22\x5a\x37\xcc\xfa\xea\xd0\x3a\xad\x8d\xf6\x03\x93\xe9\xed\x68\xb4\x1e\x16\x42\x72\x44\x06\xd1\xf2\x7f\x64\x63\x91\x3d\xf7\x7a\x84\x4e\x88\x57\xb3\x42\x0c\x14\xe6\x3d\xdd\x5b\x0d\xa8\x8f\x83\x9b\x3e\xa4\x8d\x73\xa8\xf0\x89\xec\x07\xa7\x52\xe2\x1f\x89\x67\x0c\xf4\xfb\xb9\x7b\x75\xdf\x01\x00\x00\xff\xff\x59\xd0\x5e\x39\x2e\x01\x00\x00")

func rubyLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_rubyLatestJson,
		"ruby/latest.json",
	)
}

func rubyLatestJson() (*asset, error) {
	bytes, err := rubyLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ruby/latest.json", size: 302, mode: os.FileMode(420), modTime: time.Unix(1593045091, 0)}
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
	"crystal/0.35.1.json":    crystal0351Json,
	"crystal/latest.json":    crystalLatestJson,
	"golang/1.13.12.json":    golang11312Json,
	"golang/1.14.4.json":     golang1144Json,
	"golang/latest.json":     golangLatestJson,
	"memcached/1.6.5.json":   memcached165Json,
	"memcached/1.6.6.json":   memcached166Json,
	"memcached/latest.json":  memcachedLatestJson,
	"mysql/8.0.17.json":      mysql8017Json,
	"mysql/latest.json":      mysqlLatestJson,
	"node/12.18.1.json":      node12181Json,
	"node/14.4.0.json":       node1440Json,
	"node/latest.json":       nodeLatestJson,
	"postgresql/10.13.json":  postgresql1013Json,
	"postgresql/11.8.json":   postgresql118Json,
	"postgresql/12.3.json":   postgresql123Json,
	"postgresql/9.5.22.json": postgresql9522Json,
	"postgresql/9.6.18.json": postgresql9618Json,
	"postgresql/latest.json": postgresqlLatestJson,
	"redis/6.0.4.json":       redis604Json,
	"redis/latest.json":      redisLatestJson,
	"ruby/2.6.6.json":        ruby266Json,
	"ruby/2.7.1.json":        ruby271Json,
	"ruby/latest.json":       rubyLatestJson,
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
	"crystal": &bintree{nil, map[string]*bintree{
		"0.35.1.json": &bintree{crystal0351Json, map[string]*bintree{}},
		"latest.json": &bintree{crystalLatestJson, map[string]*bintree{}},
	}},
	"golang": &bintree{nil, map[string]*bintree{
		"1.13.12.json": &bintree{golang11312Json, map[string]*bintree{}},
		"1.14.4.json":  &bintree{golang1144Json, map[string]*bintree{}},
		"latest.json":  &bintree{golangLatestJson, map[string]*bintree{}},
	}},
	"memcached": &bintree{nil, map[string]*bintree{
		"1.6.5.json":  &bintree{memcached165Json, map[string]*bintree{}},
		"1.6.6.json":  &bintree{memcached166Json, map[string]*bintree{}},
		"latest.json": &bintree{memcachedLatestJson, map[string]*bintree{}},
	}},
	"mysql": &bintree{nil, map[string]*bintree{
		"8.0.17.json": &bintree{mysql8017Json, map[string]*bintree{}},
		"latest.json": &bintree{mysqlLatestJson, map[string]*bintree{}},
	}},
	"node": &bintree{nil, map[string]*bintree{
		"12.18.1.json": &bintree{node12181Json, map[string]*bintree{}},
		"14.4.0.json":  &bintree{node1440Json, map[string]*bintree{}},
		"latest.json":  &bintree{nodeLatestJson, map[string]*bintree{}},
	}},
	"postgresql": &bintree{nil, map[string]*bintree{
		"10.13.json":  &bintree{postgresql1013Json, map[string]*bintree{}},
		"11.8.json":   &bintree{postgresql118Json, map[string]*bintree{}},
		"12.3.json":   &bintree{postgresql123Json, map[string]*bintree{}},
		"9.5.22.json": &bintree{postgresql9522Json, map[string]*bintree{}},
		"9.6.18.json": &bintree{postgresql9618Json, map[string]*bintree{}},
		"latest.json": &bintree{postgresqlLatestJson, map[string]*bintree{}},
	}},
	"redis": &bintree{nil, map[string]*bintree{
		"6.0.4.json":  &bintree{redis604Json, map[string]*bintree{}},
		"latest.json": &bintree{redisLatestJson, map[string]*bintree{}},
	}},
	"ruby": &bintree{nil, map[string]*bintree{
		"2.6.6.json":  &bintree{ruby266Json, map[string]*bintree{}},
		"2.7.1.json":  &bintree{ruby271Json, map[string]*bintree{}},
		"latest.json": &bintree{rubyLatestJson, map[string]*bintree{}},
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
