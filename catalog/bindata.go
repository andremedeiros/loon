// Code generated by go-bindata. DO NOT EDIT.
// sources:
// catalog/nix/golang/1.13.12.json (189B)
// catalog/nix/golang/1.14.4.json (207B)
// catalog/nix/memcached/1.6.6.json (215B)
// catalog/nix/mysql/8.0.20.json (224B)
// catalog/nix/postgresql/12.3.json (231B)
// catalog/nix/redis/6.0.4.json (213B)
// catalog/nix/ruby/2.6.6.json (198B)
// catalog/nix/ruby/2.7.1.json (218B)

package catalog

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

var _nixGolang11312Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\x4b\x6e\xc4\x20\x10\x04\xd0\x3d\xa7\x68\xb1\x8e\xc0\x7c\x6c\xc0\xb7\xc1\xdd\x18\x47\x22\x26\x02\x27\x8b\x19\xf9\xee\x23\xcf\x67\x59\x55\x7a\x75\x67\x00\x3c\xd7\x12\xf7\xcc\x67\xb8\x12\x00\xff\x4f\xad\x7f\xd7\x9d\xcf\xc0\x95\x50\x46\x28\xcd\xbf\x5e\xcb\x5f\x2b\x57\xbb\x1d\xc7\x6f\x9f\xa5\xa4\x22\x72\xad\xb9\x24\x81\xf5\x47\xe6\x2a\x73\x7d\x03\xd1\x1b\x8a\x23\x36\x91\x6f\x1f\xdb\xb7\xa8\xc7\xe9\x79\xea\x96\xa8\xd1\x52\xb2\xe4\xbc\x0b\x26\x6a\x85\x38\x8d\x81\x42\x18\xdc\x6a\xcd\x38\x21\x05\xf4\x94\xfc\xe2\x68\xf0\x21\x20\x25\xc2\xb4\x3a\xa5\xd3\x12\x8d\xe5\x0c\xe0\x64\x27\x7b\x04\x00\x00\xff\xff\x01\xa6\xcf\x0d\xbd\x00\x00\x00")

func nixGolang11312JsonBytes() ([]byte, error) {
	return bindataRead(
		_nixGolang11312Json,
		"nix/golang/1.13.12.json",
	)
}

func nixGolang11312Json() (*asset, error) {
	bytes, err := nixGolang11312JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nix/golang/1.13.12.json", size: 189, mode: os.FileMode(0644), modTime: time.Unix(1591831486, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xef, 0xc6, 0x11, 0xc1, 0x5a, 0xda, 0x8c, 0xdc, 0x20, 0x5c, 0x0, 0xba, 0xa7, 0x6e, 0x3c, 0x84, 0xf4, 0xe7, 0xd7, 0x2f, 0x69, 0xb1, 0x37, 0xb1, 0x98, 0x6b, 0x9, 0xac, 0x32, 0xe5, 0x17, 0x7b}}
	return a, nil
}

var _nixGolang1144Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8e\x4b\x6a\xc5\x20\x14\x86\xe7\xae\xe2\xe0\xb8\xf8\xba\x7a\x63\xdd\xcd\x51\x4f\x4c\xc1\xc6\xa2\xa6\x83\x96\xec\xbd\x84\xd0\x3b\xfc\x1f\x1f\x7c\xbf\x0c\x80\x97\x56\x71\x2f\x3c\xc0\x95\x00\xf8\x37\xf5\xf1\xd1\x76\x1e\x80\x6b\xa1\xad\xb0\xfc\xed\x1e\x8e\x5e\xaf\x72\x9b\xf3\x6b\x04\x29\x73\x15\xa5\xb5\x52\x49\xa4\xf6\x29\x4b\x93\xa5\xdd\x7f\x31\x7a\x12\x13\xbb\x28\x3f\xff\xe8\xd8\xd0\xb8\xe7\x45\x2f\x4a\x6b\x5c\x1f\x31\x26\x83\x49\x2b\x9f\x75\xf4\x86\xd0\x63\x8c\x7e\x89\x86\x9e\x8f\x75\xf1\xde\xda\x55\x19\xf7\x1e\xc9\xa8\x94\xc9\x66\x6b\x92\x4b\x56\x39\xff\x92\xa9\x38\x69\x4c\x1e\x60\xf6\x83\x18\xc0\xc9\x4e\xf6\x17\x00\x00\xff\xff\x6d\x22\xd1\x66\xcf\x00\x00\x00")

func nixGolang1144JsonBytes() ([]byte, error) {
	return bindataRead(
		_nixGolang1144Json,
		"nix/golang/1.14.4.json",
	)
}

func nixGolang1144Json() (*asset, error) {
	bytes, err := nixGolang1144JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nix/golang/1.14.4.json", size: 207, mode: os.FileMode(0644), modTime: time.Unix(1591831602, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x55, 0xeb, 0x9e, 0x4d, 0x1f, 0xb, 0x42, 0x65, 0x96, 0xeb, 0xe0, 0xfd, 0xb2, 0xf8, 0xcc, 0xef, 0x4b, 0x5f, 0x6f, 0x8d, 0x97, 0xb4, 0x55, 0xd, 0x6, 0x24, 0xce, 0xc6, 0x5e, 0x19, 0xdb, 0x6e}}
	return a, nil
}

var _nixMemcached166Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\x4d\x4e\xc3\x30\x10\x40\xe1\xbd\x4f\x31\xf2\x1a\x1c\xd7\x7f\x8d\x7d\x9b\x61\x32\xd3\x54\x4a\x09\xb2\x1d\x16\xa0\xdc\x1d\x05\x14\xba\xfd\xa4\xf7\xbe\x15\x80\x7e\xf0\x83\x90\x66\x9e\x74\x81\x03\x00\xf4\x27\xd7\x76\x5f\xdf\x75\x01\x7d\x31\xc9\x24\xfd\xf2\xe7\x5b\x5d\x0e\x9b\x7b\xff\x68\x65\x18\xfe\x4b\xb3\xd6\xdb\x20\xf7\x85\xdb\xd3\x5e\x7f\x4b\xd3\xb1\x9a\xdb\xd7\x39\x68\x33\xba\x98\x8e\x47\xb6\xa3\x58\x66\x12\x8c\x31\x5f\x5c\xa6\xcc\x21\xf0\x44\x21\x89\x75\x7c\xf5\x28\x3c\x0a\x12\xfa\x18\xdf\x02\x0b\x45\x1a\xd3\x94\xad\x13\xf2\xec\x9d\x5c\xcf\xe3\x82\x9d\x5b\xd7\x05\x7a\xdd\x58\x01\xec\x6a\x57\x3f\x01\x00\x00\xff\xff\x0b\x1f\x09\xa8\xd7\x00\x00\x00")

func nixMemcached166JsonBytes() ([]byte, error) {
	return bindataRead(
		_nixMemcached166Json,
		"nix/memcached/1.6.6.json",
	)
}

func nixMemcached166Json() (*asset, error) {
	bytes, err := nixMemcached166JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nix/memcached/1.6.6.json", size: 215, mode: os.FileMode(0644), modTime: time.Unix(1591831372, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x90, 0xb2, 0x24, 0xb3, 0x1b, 0x29, 0x8d, 0xad, 0x15, 0x9, 0xe9, 0xde, 0x62, 0x68, 0x26, 0xbc, 0xc1, 0xee, 0xfb, 0x71, 0x4f, 0xdf, 0x15, 0xfc, 0xfd, 0xa7, 0xdf, 0x98, 0xa5, 0x3d, 0x5b, 0xa5}}
	return a, nil
}

var _nixMysql8020Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\xc1\x4e\xc3\x30\x0c\xc6\xf1\x7b\x9f\xc2\xf2\x99\x25\x6e\xb6\x34\x5e\xcf\x1c\xe1\x80\x78\x02\x2f\x4e\x18\x52\xd7\x42\x93\x81\x06\xda\xbb\x4f\xdd\xd4\xa3\xf5\xf7\xf7\xfb\x6f\x00\xf0\x74\x29\xdf\x03\xf6\xb0\x1c\x00\xf8\x93\xe6\xf2\x39\x8d\xd8\x03\xb2\x21\xe3\x08\x9f\x1e\xe1\x3c\x2f\x5f\x78\xac\xf5\xab\xf4\xd6\x46\x1d\xcd\x7d\x6a\xe2\x74\xb2\xf6\x79\xfa\x1d\x87\x49\xb4\xd8\xd7\xcb\xfb\xdb\xcb\x86\x0d\xd9\x7b\xde\x3c\x14\x53\x65\x36\x1f\x7f\x2b\x56\x8e\xe2\x7c\xb7\x78\x3b\xed\x5c\x92\x28\xbb\x6d\xd4\x7d\x64\xcf\x07\x89\xd2\x92\x92\x63\x1f\xb6\xec\x98\x83\xa7\x20\x6d\xc8\x59\x0f\x59\xf6\x41\x25\xb7\x14\x94\x62\xe7\x52\xf2\x79\x15\x07\xa9\xa9\x54\xec\xa1\xce\xe7\xd4\x00\x5c\x9b\x6b\x73\x0b\x00\x00\xff\xff\xc7\xd9\x31\xea\xe0\x00\x00\x00")

func nixMysql8020JsonBytes() ([]byte, error) {
	return bindataRead(
		_nixMysql8020Json,
		"nix/mysql/8.0.20.json",
	)
}

func nixMysql8020Json() (*asset, error) {
	bytes, err := nixMysql8020JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nix/mysql/8.0.20.json", size: 224, mode: os.FileMode(0644), modTime: time.Unix(1591831377, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x50, 0xad, 0x58, 0xf2, 0xe4, 0x4e, 0x74, 0x77, 0x21, 0x6b, 0xa7, 0x5c, 0xd, 0x10, 0x76, 0x4d, 0x39, 0x18, 0x90, 0xc7, 0xf8, 0xc5, 0x4e, 0xc7, 0x80, 0xb2, 0x6a, 0x86, 0xac, 0x8c, 0xaf, 0x6f}}
	return a, nil
}

var _nixPostgresql123Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8c\xc1\x6e\x83\x30\x0c\x86\xef\x3c\x85\xe5\xf3\x96\x40\x5c\x27\xc0\xdb\x18\x9a\xd0\x49\xd5\x60\x89\xd3\xc3\xa6\xbe\xfb\x14\x55\x15\x47\x7f\xfe\xbf\xef\xaf\x03\xc0\x63\x2f\xba\xe5\x58\x7e\xee\x38\x43\x23\x00\xf8\x88\xb9\x7c\xed\xdf\x38\x03\x0e\xce\x10\x7e\xbc\x70\xcd\x6d\x83\x37\xd5\xa3\xcc\xd6\x26\x3d\xcc\x69\x9b\x3d\x6f\xf6\xa8\x8b\x2d\x7b\xcd\x6b\xb4\x8f\x66\xda\xf3\xff\xd9\x6e\xa3\x92\xcd\xf6\xfb\x0e\x96\x9b\x38\xf6\xad\x19\xfa\x31\x5d\x79\x21\x27\x53\xe0\x10\x7c\x98\xae\xb4\x0e\x34\xba\xcb\xea\x89\x26\xf2\x69\x1c\xbd\x04\xa2\xb4\x32\xcb\xc2\x32\xb9\x4b\xbf\xf8\x81\x65\xe8\x39\x71\xff\x2e\xde\x45\x63\x51\x9c\x41\x73\x8d\x1d\xc0\xb3\x7b\x76\xff\x01\x00\x00\xff\xff\xfe\x3d\xba\xb3\xe7\x00\x00\x00")

func nixPostgresql123JsonBytes() ([]byte, error) {
	return bindataRead(
		_nixPostgresql123Json,
		"nix/postgresql/12.3.json",
	)
}

func nixPostgresql123Json() (*asset, error) {
	bytes, err := nixPostgresql123JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nix/postgresql/12.3.json", size: 231, mode: os.FileMode(0644), modTime: time.Unix(1591837324, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7f, 0x8f, 0xad, 0x21, 0xab, 0x74, 0xca, 0xbe, 0xa5, 0x23, 0x11, 0xfc, 0xca, 0x7b, 0xf6, 0x18, 0x6c, 0xc4, 0x5c, 0xbf, 0x70, 0xbb, 0xed, 0x41, 0xb, 0x5, 0x70, 0xc2, 0x22, 0x77, 0x23, 0x5c}}
	return a, nil
}

var _nixRedis604Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\x5b\x6a\xc3\x30\x10\x85\xe1\x77\xaf\x62\xd0\x73\xab\x8b\x47\xd6\x6d\x37\x13\x69\xda\x18\x44\x54\x24\xa5\x85\x96\xec\xbd\xd8\x21\x8f\xe7\x87\xf3\xfd\x2d\x00\xa2\x73\xd9\x87\x48\x70\x0c\x00\xf1\xcd\x7d\xec\xed\x26\x12\x08\x27\xb5\xb4\xe2\xed\xd9\xef\xbd\x1e\xed\x3a\xe7\x57\x52\xaa\xb4\x9f\x5b\x6d\x54\xe4\xf9\x96\x7b\x53\x9d\x2b\xd3\xe0\xa1\xce\xf2\x7e\x7e\xe5\xa4\x2e\x3f\x7f\x5f\xc4\xb8\xd2\xba\xb9\x43\x41\x44\xaf\xf5\x46\x86\x75\x46\xa2\x35\x62\x0e\x3e\xa3\x41\xeb\x3c\x53\xa0\x6c\x4c\x0c\x36\xae\xe6\x83\x2e\x3a\x04\xed\x63\x0c\x17\xf2\x6e\xcb\x31\x5a\x2c\xfc\x12\x2b\x4d\x1e\x53\x24\x98\xfd\xce\x0b\xc0\x63\x79\x2c\xff\x01\x00\x00\xff\xff\x58\xc2\xf3\xf2\xd5\x00\x00\x00")

func nixRedis604JsonBytes() ([]byte, error) {
	return bindataRead(
		_nixRedis604Json,
		"nix/redis/6.0.4.json",
	)
}

func nixRedis604Json() (*asset, error) {
	bytes, err := nixRedis604JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nix/redis/6.0.4.json", size: 213, mode: os.FileMode(0644), modTime: time.Unix(1591829878, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x73, 0xfa, 0x2e, 0x87, 0x2b, 0xd2, 0xf9, 0x8b, 0x99, 0xed, 0xce, 0x9a, 0x3c, 0xbf, 0xb5, 0x3e, 0xb4, 0xc6, 0x22, 0xc3, 0xb5, 0x89, 0x79, 0xb4, 0xf4, 0xd1, 0xb3, 0xdc, 0xc1, 0xf4, 0x79, 0xc6}}
	return a, nil
}

var _nixRuby266Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\xdd\x6a\x84\x30\x10\xc5\xf1\xfb\x3c\xc5\x90\xeb\x6e\xbe\x33\x41\xdf\x66\x26\x89\x5a\x58\x56\x89\x5a\x68\x8b\xef\x5e\x62\xd9\xbb\xc3\x1f\xce\xef\x57\x00\xc8\x76\xf2\xb7\x1c\xa1\x6f\x00\xf9\x55\xdb\xfe\xb9\xbe\xe4\x08\xd2\x29\x54\x28\x3f\xfe\xfb\xd9\x9e\xbd\x2d\xc7\xb1\xed\xa3\xd6\x99\xf2\x52\x55\xbf\x3e\x9e\xf4\x9a\xd5\xda\x66\xbd\x9d\xac\x7b\xd1\x4e\xe1\x3d\x1e\xb7\xa0\x0e\x6a\x6a\xfe\x79\x43\xfb\x42\x2e\x62\xb7\x3c\x06\xb6\xc1\x97\x3a\x79\x34\x4c\xd9\x72\x0a\x95\x23\xd6\x82\x86\x2d\x99\x92\x09\x83\x1f\xd8\x18\x1b\x13\x55\x6b\xa7\x29\xa5\x12\x73\x71\x75\x70\x6e\xb0\x52\x00\x5c\xe2\x12\x7f\x01\x00\x00\xff\xff\xc8\x18\x25\x82\xc6\x00\x00\x00")

func nixRuby266JsonBytes() ([]byte, error) {
	return bindataRead(
		_nixRuby266Json,
		"nix/ruby/2.6.6.json",
	)
}

func nixRuby266Json() (*asset, error) {
	bytes, err := nixRuby266JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nix/ruby/2.6.6.json", size: 198, mode: os.FileMode(0644), modTime: time.Unix(1591831101, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3, 0x18, 0x6e, 0xdd, 0xfd, 0x3e, 0x54, 0xa2, 0x91, 0x31, 0x6f, 0x43, 0x93, 0x37, 0x81, 0x92, 0x71, 0x6d, 0x44, 0x57, 0xf0, 0x37, 0xa, 0x1c, 0x46, 0xb4, 0xaf, 0xb, 0x15, 0x75, 0x33, 0x14}}
	return a, nil
}

var _nixRuby271Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\xdd\x6a\xc3\x30\x0c\xc5\xf1\xfb\x3c\x85\xf0\xf5\x6a\x4b\x6a\x1c\x67\x79\x1b\x2b\x16\xc9\x20\xb4\xc5\x1f\x83\x6d\xf4\xdd\x87\x5b\x72\x77\xf8\xc3\xf9\xfd\x0d\x00\x26\x37\xf9\x31\x0b\xf4\x0d\x60\xbe\x35\x97\xaf\xfb\xcd\x2c\x60\xd8\x06\x4b\xe6\xe3\xdd\x5b\x3e\x7a\xdb\x6b\x7d\x94\xc5\xb9\x35\xae\xbb\xda\x7e\xbd\x1c\xf1\xb6\xd9\x7b\xde\xdc\xa3\x89\xeb\xc5\xb1\x0d\xaf\x71\x79\x09\xb6\xc6\x6c\xb7\xdf\x13\x2a\x7b\x64\x3f\x75\x2b\x8d\x34\x8f\xf3\x55\x52\x42\x44\xf4\x61\x5a\xe9\x1a\xd0\x07\x22\xa6\x38\xa9\xf0\xe8\x67\x26\x9a\x92\xa0\x04\x11\x46\xf4\xfa\x89\xca\x1e\x35\xea\x48\xf3\x29\x1e\xb1\x6a\xa9\x66\x81\x9a\x9b\x0e\x00\xcf\xe1\x39\xfc\x07\x00\x00\xff\xff\x7f\x81\xcf\x8f\xda\x00\x00\x00")

func nixRuby271JsonBytes() ([]byte, error) {
	return bindataRead(
		_nixRuby271Json,
		"nix/ruby/2.7.1.json",
	)
}

func nixRuby271Json() (*asset, error) {
	bytes, err := nixRuby271JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nix/ruby/2.7.1.json", size: 218, mode: os.FileMode(0644), modTime: time.Unix(1591831228, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd5, 0x30, 0x1e, 0x5d, 0x4e, 0x0, 0x72, 0x5b, 0x70, 0x1b, 0xe9, 0xbd, 0x61, 0x65, 0x58, 0xfb, 0x80, 0x19, 0x45, 0x47, 0x56, 0x88, 0x50, 0xe5, 0x62, 0x8a, 0x3b, 0x91, 0xcd, 0x1, 0x81, 0xae}}
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
	"nix/golang/1.13.12.json":  nixGolang11312Json,
	"nix/golang/1.14.4.json":   nixGolang1144Json,
	"nix/memcached/1.6.6.json": nixMemcached166Json,
	"nix/mysql/8.0.20.json":    nixMysql8020Json,
	"nix/postgresql/12.3.json": nixPostgresql123Json,
	"nix/redis/6.0.4.json":     nixRedis604Json,
	"nix/ruby/2.6.6.json":      nixRuby266Json,
	"nix/ruby/2.7.1.json":      nixRuby271Json,
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
	"nix": &bintree{nil, map[string]*bintree{
		"golang": &bintree{nil, map[string]*bintree{
			"1.13.12.json": &bintree{nixGolang11312Json, map[string]*bintree{}},
			"1.14.4.json":  &bintree{nixGolang1144Json, map[string]*bintree{}},
		}},
		"memcached": &bintree{nil, map[string]*bintree{
			"1.6.6.json": &bintree{nixMemcached166Json, map[string]*bintree{}},
		}},
		"mysql": &bintree{nil, map[string]*bintree{
			"8.0.20.json": &bintree{nixMysql8020Json, map[string]*bintree{}},
		}},
		"postgresql": &bintree{nil, map[string]*bintree{
			"12.3.json": &bintree{nixPostgresql123Json, map[string]*bintree{}},
		}},
		"redis": &bintree{nil, map[string]*bintree{
			"6.0.4.json": &bintree{nixRedis604Json, map[string]*bintree{}},
		}},
		"ruby": &bintree{nil, map[string]*bintree{
			"2.6.6.json": &bintree{nixRuby266Json, map[string]*bintree{}},
			"2.7.1.json": &bintree{nixRuby271Json, map[string]*bintree{}},
		}},
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
