// Code generated by go-bindata. DO NOT EDIT.
// sources:
// internal/catalog/service/memcached/1.6.5.json (195B)
// internal/catalog/service/memcached/1.6.6.json (195B)
// internal/catalog/service/mysql/8.0.17.json (194B)
// internal/catalog/service/postgresql/10.13.json (202B)
// internal/catalog/service/postgresql/11.8.json (199B)
// internal/catalog/service/postgresql/12.3.json (211B)
// internal/catalog/service/postgresql/9.5.22.json (205B)
// internal/catalog/service/postgresql/9.6.18.json (205B)
// internal/catalog/service/redis/6.0.4.json (193B)

package service

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

	info := bindataFileInfo{name: "memcached/1.6.5.json", size: 195, mode: os.FileMode(0644), modTime: time.Unix(1592261973, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x92, 0xab, 0xe1, 0x72, 0xe0, 0x13, 0x8e, 0xa7, 0xce, 0xd8, 0x50, 0xd2, 0x62, 0x8c, 0x3f, 0x59, 0x1b, 0x1d, 0x8, 0x37, 0xf, 0xad, 0x28, 0xef, 0x66, 0xb8, 0xd2, 0xb3, 0xcf, 0xee, 0xf8, 0x3c}}
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

	info := bindataFileInfo{name: "memcached/1.6.6.json", size: 195, mode: os.FileMode(0644), modTime: time.Unix(1592435152, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x80, 0x67, 0x5a, 0xb0, 0x67, 0x88, 0x70, 0x64, 0xb1, 0xec, 0x52, 0xc0, 0x56, 0x45, 0x9e, 0xdf, 0x59, 0xd0, 0x88, 0xe4, 0xed, 0x6f, 0xc2, 0x5, 0x66, 0x8d, 0xe1, 0xc7, 0x50, 0x44, 0x21, 0x20}}
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

	info := bindataFileInfo{name: "mysql/8.0.17.json", size: 194, mode: os.FileMode(0644), modTime: time.Unix(1592435148, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x11, 0x1, 0x48, 0xa5, 0x7d, 0x74, 0x70, 0x22, 0x12, 0x10, 0xb8, 0xb0, 0x7a, 0x51, 0xb, 0x4b, 0x1c, 0xf5, 0xfe, 0xbd, 0x45, 0xca, 0x49, 0xf, 0x46, 0x1c, 0xdd, 0xa0, 0xce, 0x24, 0x8a, 0xd6}}
	return a, nil
}

var _postgresql1013Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\x4d\x0e\x83\x20\x10\x46\xf7\x9c\x62\x32\xeb\x16\xd4\x46\x8d\xde\x06\xff\xc0\x84\x0a\x0e\xa8\xc1\xc6\xbb\x37\xda\x34\x6e\x5f\xde\xfb\xbe\x0f\x03\x40\x67\x7d\x50\xd4\xfb\xd9\x60\x0d\x27\x01\xc0\xb5\x27\x3f\xda\x09\x6b\xc0\x34\xe1\xe9\x0b\x1f\x3f\xbe\xd0\x29\xa1\x0e\xc1\xf9\x5a\x88\x21\x38\x7e\xe7\xdc\x92\x12\x6e\x69\x84\xb7\x0b\xb5\xbd\x58\xaf\x54\xdc\xc2\xf3\x02\x3c\x48\xe2\x6a\xff\x4f\x7a\x2d\xb3\xbc\xb8\x9e\x66\x69\x92\xe8\x4a\x59\x25\x71\x8f\xb2\xd4\x26\x2f\x94\x7f\x6f\xf9\xb0\xca\xd6\x19\xea\xb4\x9b\x4a\x5d\xa9\xa9\x89\x94\x8e\xd9\x18\xb7\xac\x43\x06\x70\xb0\x83\x7d\x03\x00\x00\xff\xff\x5f\x36\x86\x3f\xca\x00\x00\x00")

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

	info := bindataFileInfo{name: "postgresql/10.13.json", size: 202, mode: os.FileMode(0644), modTime: time.Unix(1592433269, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe0, 0xd9, 0x7, 0x35, 0xab, 0x5b, 0xf2, 0x48, 0x2e, 0xee, 0x6e, 0x8e, 0x38, 0x24, 0x2c, 0x80, 0xe3, 0x12, 0x2d, 0x3c, 0x4a, 0xa8, 0x86, 0x8d, 0x81, 0xdb, 0xe2, 0xc5, 0x2b, 0x58, 0x60, 0x57}}
	return a, nil
}

var _postgresql118Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\x5b\x72\x83\x30\x0c\x45\xff\xbd\x0a\x8d\xbe\x5b\xbb\xf4\xc1\x14\x76\x63\x5c\x6a\x03\x01\xdb\x92\x79\x25\xc3\xde\x33\x4e\x26\xc3\xa7\x8e\xee\xb9\xf7\x26\x00\x30\x78\x4e\x96\x5a\x8e\x17\xac\x21\x13\x00\x5c\x5a\xe2\xce\x4f\x58\x03\x16\x85\xfc\xc5\xb7\x27\x9e\x29\x67\xd0\xa5\x14\xb8\x56\xea\x3f\x05\x79\xda\xd2\x93\x55\x61\x6e\x14\xfb\x99\x4c\xab\x96\x6c\xaa\xf3\xff\x9e\x6f\x99\x34\x49\x7b\x7d\x15\xb2\xd3\x9f\x3f\xe5\x63\x26\x0e\x1c\x77\xbd\x6f\xe3\x64\xcc\xd8\xe8\x60\xbf\x74\xcf\x6b\x15\x7a\x1b\x17\xfd\xd1\x4d\x5b\xef\xb6\xf2\x9b\xe2\x5f\x69\x06\x47\xb6\x5a\x83\x46\x01\x70\x88\x43\xdc\x03\x00\x00\xff\xff\xaa\x4b\x9b\x8f\xc7\x00\x00\x00")

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

	info := bindataFileInfo{name: "postgresql/11.8.json", size: 199, mode: os.FileMode(0644), modTime: time.Unix(1592433216, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x85, 0x9a, 0x6a, 0xec, 0xd0, 0xcc, 0x4f, 0x34, 0xd, 0x4f, 0x8, 0x7d, 0x34, 0x47, 0xac, 0xbe, 0x18, 0x44, 0x65, 0xd9, 0x30, 0xb0, 0x87, 0x5, 0xa3, 0x22, 0xec, 0x2e, 0x79, 0x3, 0x16, 0x98}}
	return a, nil
}

var _postgresql123Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8c\x4d\x6e\xc3\x20\x10\x46\xf7\x3e\xc5\x68\xd6\x2d\x18\x26\xc3\x8f\x6f\x03\xc4\x38\x95\xaa\x42\x01\x67\xd1\x2a\x77\xaf\x50\x15\x65\x39\x6f\xbe\xf7\x7e\x17\x00\xac\xa5\x8f\xa3\xed\xfd\xfb\x13\x37\x98\x04\x00\xef\x7b\xeb\x1f\xe5\x0b\x37\x40\xa5\x05\xe1\xdb\x3f\x3e\xdb\xdc\xe0\x6d\x8c\xda\x37\x29\xf3\xa8\xe2\x65\x8b\xd2\x0e\x59\xcf\x28\x7b\x39\x5b\xda\xe5\x7d\x9a\xf2\xf5\x7f\x9f\xb7\x18\xa1\x89\xe3\xe7\x19\xec\xb7\xa0\xd9\xcc\xa6\x5d\x5d\xbe\x72\x24\x1d\xbc\x65\x6b\x8d\xf5\x57\x4a\x8a\x9c\xbe\x24\x43\xe4\xc9\x64\xe7\x4c\xb0\x44\x39\x31\x87\xc8\xc1\xeb\xcb\x1a\x8d\xe2\xa0\x56\xce\xbc\xe2\x02\xf0\x58\x1e\xcb\x5f\x00\x00\x00\xff\xff\xe8\x7d\xff\x9d\xd3\x00\x00\x00")

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

	info := bindataFileInfo{name: "postgresql/12.3.json", size: 211, mode: os.FileMode(0644), modTime: time.Unix(1592435144, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe1, 0xcc, 0x74, 0xa2, 0xfe, 0xc7, 0x83, 0x11, 0xb0, 0xe9, 0x95, 0x13, 0x83, 0x70, 0xa4, 0xbf, 0xd1, 0x2, 0x42, 0x52, 0x77, 0x54, 0x8f, 0xba, 0x39, 0xd2, 0xe, 0x15, 0x4e, 0x60, 0x2e, 0x84}}
	return a, nil
}

var _postgresql9522Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8c\x4d\x6e\x85\x20\x14\x46\xe7\xac\xe2\xe6\x8e\x5b\x78\xfe\x60\xd0\xdd\x50\xab\x28\x51\xc0\x0b\xd2\xd8\xc6\xbd\x37\xc6\xbc\x38\xfd\xbe\x73\xce\x1f\x03\xc0\xe0\x63\x32\x34\xc4\x6d\xc1\x0e\xae\x05\x00\xf3\x40\x71\xf6\x0e\x3b\xc0\x96\x4b\x5e\x96\xf8\x71\x1f\x3b\x5d\x14\x4e\x29\x85\xd8\x09\x31\xa6\xc0\x1f\x9f\x7b\x32\x22\xec\x5f\x22\xfa\x9d\xfa\x41\xe4\xdb\x15\x0f\xf1\x79\x2f\x3c\x69\xe2\xe6\xf7\x1d\x8d\x93\x2e\x65\x73\x75\x5f\x55\xae\xbf\x6b\x47\xed\xa8\x9a\xe3\x35\x17\x56\xda\x35\x8f\xda\xc9\x1f\x75\xd4\x46\x17\xab\x26\xd9\x2e\x53\xef\x6c\x65\x17\xb9\x49\xb5\x6a\x85\x0c\xe0\x64\x27\xfb\x0f\x00\x00\xff\xff\x1a\x8a\x71\xab\xcd\x00\x00\x00")

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

	info := bindataFileInfo{name: "postgresql/9.5.22.json", size: 205, mode: os.FileMode(0644), modTime: time.Unix(1592433536, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8a, 0x55, 0x81, 0x84, 0xd8, 0x62, 0x2b, 0x14, 0xa, 0x1d, 0x6, 0x57, 0x86, 0xb3, 0xd9, 0x90, 0x67, 0x96, 0xc7, 0x81, 0x75, 0xd6, 0x68, 0x4d, 0xc9, 0xf1, 0x73, 0x4e, 0x70, 0x62, 0x78, 0x4a}}
	return a, nil
}

var _postgresql9618Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8c\x5d\x6e\x86\x20\x10\x45\xdf\x59\xc5\x64\x9e\x5b\x50\x13\xb5\xba\x1b\xfc\x03\x13\x2a\x38\x83\x1a\x6c\xdc\x7b\x63\xcc\x17\x5f\xef\x3d\xe7\xfc\x09\x00\x0c\x9e\xa3\xa1\x91\x57\x87\x2d\xdc\x0b\x00\xee\x23\xf1\xec\x17\x6c\x01\x1b\x59\xc9\xfc\x07\xbf\x9e\x63\xa3\x9b\x42\x1b\x63\xe0\x56\xa9\x29\x06\xf9\xfa\xd2\x93\x51\x61\xeb\x14\xfb\x8d\xfa\x51\xed\x8f\xab\x5e\xe2\xfb\x59\x64\xd4\x24\xcd\xf9\x89\xb2\xd5\x45\x59\xdd\xdd\x7c\xd5\x2e\x4b\xa1\xd6\x4d\x96\xce\xa4\x6b\xeb\xca\xca\xf0\xef\x51\x4e\xbb\xee\x83\xa3\xc1\x86\xa5\xb6\x8d\x59\xba\x44\xf9\x5c\xcc\xe9\x28\x06\x14\x00\x97\xb8\xc4\x7f\x00\x00\x00\xff\xff\xb5\x6b\x6a\xe9\xcd\x00\x00\x00")

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

	info := bindataFileInfo{name: "postgresql/9.6.18.json", size: 205, mode: os.FileMode(0644), modTime: time.Unix(1592433346, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa7, 0xc0, 0x7f, 0x90, 0x43, 0x6c, 0x60, 0x9d, 0xf3, 0xbc, 0xfb, 0xa0, 0x86, 0x0, 0x94, 0x8d, 0x63, 0x16, 0x89, 0x28, 0xf2, 0x19, 0x5e, 0x1a, 0x73, 0x7a, 0xbd, 0x3, 0x83, 0xba, 0xfa, 0xd9}}
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

	info := bindataFileInfo{name: "redis/6.0.4.json", size: 193, mode: os.FileMode(0644), modTime: time.Unix(1592435138, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6a, 0xf6, 0xde, 0xf6, 0x98, 0x56, 0x85, 0xbc, 0x5c, 0x81, 0x53, 0xb5, 0x75, 0xb3, 0x4f, 0x9b, 0x49, 0xdc, 0xae, 0x4f, 0x8c, 0x83, 0x6d, 0x2c, 0x49, 0xa2, 0xab, 0xdd, 0x1, 0x88, 0x1c, 0x81}}
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
	"memcached/1.6.5.json":   memcached165Json,
	"memcached/1.6.6.json":   memcached166Json,
	"mysql/8.0.17.json":      mysql8017Json,
	"postgresql/10.13.json":  postgresql1013Json,
	"postgresql/11.8.json":   postgresql118Json,
	"postgresql/12.3.json":   postgresql123Json,
	"postgresql/9.5.22.json": postgresql9522Json,
	"postgresql/9.6.18.json": postgresql9618Json,
	"redis/6.0.4.json":       redis604Json,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"memcached": &bintree{nil, map[string]*bintree{
		"1.6.5.json": &bintree{memcached165Json, map[string]*bintree{}},
		"1.6.6.json": &bintree{memcached166Json, map[string]*bintree{}},
	}},
	"mysql": &bintree{nil, map[string]*bintree{
		"8.0.17.json": &bintree{mysql8017Json, map[string]*bintree{}},
	}},
	"postgresql": &bintree{nil, map[string]*bintree{
		"10.13.json":  &bintree{postgresql1013Json, map[string]*bintree{}},
		"11.8.json":   &bintree{postgresql118Json, map[string]*bintree{}},
		"12.3.json":   &bintree{postgresql123Json, map[string]*bintree{}},
		"9.5.22.json": &bintree{postgresql9522Json, map[string]*bintree{}},
		"9.6.18.json": &bintree{postgresql9618Json, map[string]*bintree{}},
	}},
	"redis": &bintree{nil, map[string]*bintree{
		"6.0.4.json": &bintree{redis604Json, map[string]*bintree{}},
	}},
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