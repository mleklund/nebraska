// Code generated by go-bindata. DO NOT EDIT.
// sources:
// db/drop_all_tables.sql (674B)
// db/sample_data.sql (15.065kB)
// db/migrations/0001_initial.sql (7.178kB)

package api

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

var _dbDrop_all_tablesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x41\x8e\x85\x30\x0c\x43\xf7\x9c\xa2\xf7\xe0\x30\x91\x09\x19\x88\x28\x6d\xd5\x04\x34\xdc\x7e\x24\xc4\x6c\xbe\xbe\x14\xf6\x2f\xb6\x63\xcf\xbd\xb6\xe4\x98\xb2\x24\xfd\x49\xf2\xab\xe6\x96\x5c\xb0\x27\x86\x31\x66\x19\x87\xaf\xc8\x61\xd2\x2d\x60\xd0\x5a\x56\x86\x6b\x2d\x01\xd9\xc0\x1b\x16\x09\x28\xae\x5d\xaa\x11\xf8\x85\x22\xaf\x28\x45\x72\x40\x2d\xbd\x1e\x2d\x7a\x43\x8b\x39\x0a\x47\xe9\xfe\x31\x32\x87\x1f\x6f\x45\xe9\x7d\x49\x1f\x06\xb4\xaa\x79\xed\x57\x70\x25\xa7\x14\x27\xbf\x5a\x94\xff\x06\xa3\x45\xd9\xf5\x54\x8f\x3c\x9f\x39\xe9\x19\x81\xa6\x0c\xde\xb2\x5a\x24\x3f\xc3\x31\xc1\x84\x76\x5d\xfa\x5d\x89\x8d\xc3\x5f\x00\x00\x00\xff\xff\xda\x81\x6e\xb2\xa2\x02\x00\x00")

func dbDrop_all_tablesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbDrop_all_tablesSql,
		"db/drop_all_tables.sql",
	)
}

func dbDrop_all_tablesSql() (*asset, error) {
	bytes, err := dbDrop_all_tablesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/drop_all_tables.sql", size: 674, mode: os.FileMode(420), modTime: time.Unix(1537251402, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2c, 0xfc, 0xeb, 0x41, 0x72, 0xb1, 0x9f, 0x5b, 0xf7, 0xb3, 0x6a, 0xf1, 0xba, 0xb6, 0x59, 0x6d, 0xe3, 0xa0, 0x60, 0xc2, 0x4a, 0x87, 0xe8, 0x26, 0xfc, 0x10, 0x23, 0x2a, 0xf1, 0xbb, 0xd4, 0x16}}
	return a, nil
}

var _dbSample_dataSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5b\x59\x73\xd3\xc8\x16\x7e\xe7\x57\x74\xd5\x7d\x30\x14\x6e\xa7\xf7\x25\x53\xf7\x01\xc2\x3a\x13\x18\x42\xc8\x50\xcc\x0b\xd5\xcb\x69\x47\x44\x96\x3c\x92\x9c\x90\xfc\xfa\x5b\x92\x63\xe3\x24\xe4\x46\xce\x02\x19\x8a\xb2\x91\x25\x4e\xf7\xf9\xbe\xb3\xf4\x39\xad\x06\x63\xf4\xba\xc8\x9a\xcc\xe5\x28\xba\xc6\x3d\x78\x80\x31\x7a\x06\xc9\xcd\xf2\x06\x35\xe0\x26\xc8\x15\x11\xcd\x6a\xa8\xd0\x43\x17\x27\x59\xb1\xd1\x7d\x3f\x7a\x90\x15\x35\x54\x0d\xca\x8a\xa6\x9c\xff\xbb\x87\x59\x1c\xa2\xc2\x4d\xe0\x11\x3a\x74\xf9\x0c\x6a\xf4\x70\x10\x8d\xe5\x82\xc5\x80\x2d\xa3\x02\x0b\x41\x23\x76\xc2\x25\xec\x63\x34\x5c\x3b\xee\x19\xb7\x83\x21\x1a\xc4\xf9\x84\x83\x47\xbf\x9d\x19\xb7\x9d\xb6\x46\x0f\xdb\xbf\xda\x81\x87\xa8\x86\x50\x41\x33\xec\x26\xfc\x9c\xc5\x95\x99\x3a\xad\xda\xa1\x8c\xe7\x94\x59\x16\x85\xd6\x46\x1a\x16\x08\xc8\xe4\xac\x72\x00\x52\x52\x9e\x68\x37\x5d\x1f\xb5\x1e\xfd\xd6\x71\xf1\xfc\x10\x8a\x06\x35\xc7\x53\xa8\xcf\xe8\x06\xed\xfd\xcf\xed\x7d\xf4\xb0\xfd\x1e\xa2\x0a\xea\x59\xde\x0c\x51\x84\x3a\x54\xd9\xb4\xc9\xca\xe2\x9b\x82\x7c\x88\xc8\x10\x0d\x5e\x17\x75\xe3\x8a\x00\xa8\x82\x69\x59\x35\x10\x91\x2b\x10\x54\x55\x59\xa1\x38\xab\xb2\x62\xdc\xfe\x9e\x4d\xa3\x6b\x00\xd5\x0d\x4c\x47\xe7\x29\x59\x7f\x5a\x3a\x44\x83\xbd\x6e\xc4\x0a\xed\xbb\x1a\x4d\xab\x32\x40\x5d\x77\x53\x47\xe4\xa6\xd3\x3c\x83\x88\xa6\x2e\x1c\xb8\x31\xdc\xc2\x7c\x6c\x05\x66\x8d\x66\xd3\x71\xe5\x22\x44\xd4\x94\x28\xcc\xaa\xaa\x25\x33\xec\xbb\xa2\x80\x1c\x1d\x42\x55\x67\x65\x71\xd3\x39\xe9\x29\xc8\x67\xe5\x51\x91\x97\x2e\xb6\x2c\xe6\xae\x81\xba\xb9\xb5\x19\xc4\x2a\x8d\x0b\xae\x90\xab\xaa\xec\x10\x22\xaa\x67\xa1\x65\x34\xcd\xf2\xfc\xf8\xa6\x53\x19\x42\xe6\x73\x75\x14\xe6\xf9\x62\xf0\x11\x3a\x9d\x3c\x94\x93\x69\x0e\xad\x10\x9a\x56\xdd\xe0\x10\x91\x3f\x46\xd9\x29\xe5\xa3\x85\xe3\x6e\x95\x15\xfc\xb9\x3b\x37\x70\x70\xad\xc0\x19\xbd\x56\xee\x7f\x0b\xdd\x33\x4a\x7d\x2f\xcc\xc0\x2a\x66\xa8\x53\x38\x52\x97\xb0\xf0\x11\xb0\x75\xc4\x61\xab\xbd\x56\x20\x55\x0c\x52\xb7\x11\x36\x9f\xbc\xbd\xda\xce\x8a\xd9\x57\x94\xca\x0a\x4d\x5c\x5d\x67\x87\x80\x6a\xa8\x0e\xa1\x42\x11\xa6\x79\x79\x3c\x81\xa2\xa9\xd7\x89\xc9\x55\x0c\x0b\x43\x2c\xd5\x63\xde\x89\x60\x8d\xc0\x12\xac\xc7\x82\x52\xc0\x5e\x07\x8e\x3d\x07\x4f\x45\xd2\x8e\xa9\x36\x01\xb4\xfc\x6a\xa5\x46\x7c\x44\xda\xa9\xf7\x9b\x66\x5a\x6f\x6e\x6c\x84\x72\x32\x29\x8b\x36\x0b\xd6\x4d\x59\xb5\xd1\x30\x2e\xcb\x71\x0e\x6e\x9a\xd5\xa3\x50\x4e\x36\xe6\x81\x89\x17\x4f\x43\x59\x01\x2e\xeb\x51\x01\xcd\x86\x9b\x44\x25\xf0\xac\xae\x36\x4e\x07\xde\x68\x47\x9e\x0b\x8c\xc6\x27\x83\x21\x7a\xbb\xb7\xbd\x3d\x44\x03\x2a\x85\x55\x5a\x48\xd3\x3e\xcf\xc5\x1f\x47\xfa\x09\x3c\xdd\xae\xfe\x7a\xfd\xcc\xfe\xee\xd3\xf1\x9b\xf2\x77\x78\x2a\x8f\xff\x18\xff\xb7\x7d\xce\x08\x95\x98\x58\xcc\x08\x22\x64\x93\xb2\x4d\xae\x47\x92\x71\xcb\x3b\xe9\x5e\xc6\xb8\x8a\x32\xce\xb5\xe7\x49\x03\x4e\x89\x59\x2c\x34\x18\xec\x88\x64\x38\x11\xc3\x45\x64\xd2\x47\x2f\x57\x28\x13\x77\x45\x99\xf8\x7f\x94\x49\x42\x8d\xa5\xac\x7d\x9e\xaa\x83\x03\xf7\xf8\xe9\xc6\xc9\x87\x43\xfd\xe7\xbb\x8f\xe3\x2c\x7e\x7a\x7c\x20\x76\x8b\x67\xbb\x17\x29\x53\x9b\x54\x6e\x32\x3b\xa2\xc4\x30\xa5\x6e\x8d\xb2\xc0\x1c\x57\x9c\x32\xec\xad\xb1\x58\x10\x0e\xd8\x79\xa9\x31\x51\x81\x48\xe9\x34\xb8\xc0\x4e\x29\x33\xc4\x8c\xc8\x5d\x50\x76\x3a\xf0\xa5\x94\x69\xad\xa9\x16\x54\xb4\xcf\xfd\x3f\x3c\xed\xbc\x7f\xf5\xce\x7c\x7d\xca\xdf\xbf\xd8\xfb\xb2\x15\x9f\x24\x7e\xb4\xf3\x69\x8b\x3d\xff\x8e\x97\x11\xbb\x49\xd4\xc8\x70\x6b\x8d\xbd\x35\xca\x04\x97\x86\x18\xcb\x70\x70\xd1\x60\xa1\x8c\xc3\x8e\x78\x8f\xc1\xdb\x48\x80\x58\x08\x4e\x2c\x28\xa3\xf2\x8e\x28\x9b\x0f\x7c\x39\x65\x46\x09\x2e\x75\x07\xfa\xe0\xad\x70\x93\xf2\x8f\x4f\x7f\xfd\xbd\xf7\x92\x7d\x2c\x77\xe3\xce\x2b\xfa\xee\xd5\xbb\x93\x4a\x3e\x39\x4b\x99\x44\x94\x6f\x4a\xb9\xc9\xc8\xa8\xa5\x9b\xde\x1e\x65\xcc\x88\xc8\xac\xf4\x58\x52\x93\xb0\x88\x4a\x63\x6b\x2d\x60\x2b\xac\x32\x91\x00\x44\x4b\x16\x94\x31\x7b\x47\x94\xcd\x07\xbe\x94\x32\xa3\x98\x90\x72\xee\x65\x2c\xdf\x2f\xf7\x0e\x0f\x8b\xf2\x13\x97\x76\x9a\xb1\x17\x85\xdb\xdd\xf8\x5a\x8f\x9b\x6c\x25\x30\x29\xc1\x94\x20\xc6\x37\x29\xdd\xa4\x64\x64\x98\xb4\x46\x5e\x97\xb2\x65\x1d\xb1\x5c\x9d\x88\x22\x4a\xb8\xd8\x2e\x1f\x02\x0b\x4b\x04\xb6\x4e\x01\x4e\x51\x28\x29\x2d\x8f\xd4\x77\x39\xa4\x6e\x9c\xcf\xa1\xbd\xfa\x0f\x15\xde\x46\xb5\x6a\x52\x6a\x11\x91\x6d\x14\x70\x31\x62\x8a\x32\x41\x7b\xeb\x37\x44\xfd\xd2\xe9\x55\x38\x28\x33\xde\x04\x66\xb1\x24\xb2\x8d\x16\xc1\xb1\x01\xa5\xb1\xa3\x8e\x00\x0f\x4a\xd0\xd0\xf9\x99\x87\xc6\x75\x28\x52\xd0\x89\xf3\xcb\x51\x08\xce\xc5\x0f\x47\xe1\x8c\x76\x84\xb7\xd6\x68\x57\x64\xa1\x1d\xc5\x26\x06\x81\xb9\x24\xda\x3b\xb0\x14\xa0\xa3\xd6\xe5\xd3\xfd\x39\x0c\x9a\xbc\x37\xff\xc7\x18\x52\xcb\xb5\x60\xf4\x0a\xa1\x73\x30\xc6\x55\x39\x9b\xd6\xdf\x50\x58\xc7\x22\x78\x4d\x30\xd7\x1e\xb0\x20\x4c\x61\x23\x79\xc2\x3e\xc5\xa8\x3c\x17\xda\xfb\xce\x93\x76\x97\x3e\xf5\xa2\xac\xda\x22\x3b\xce\x42\x57\x5c\x85\x7c\x56\x37\x50\xb5\x45\x4e\x72\x79\x0d\x43\xd4\x54\xb3\xe5\xf7\xe9\xad\xc1\x93\x59\xdd\x54\x2e\xcf\xdc\xc6\xee\x71\x2c\xe0\x78\xd0\xad\x7b\x68\x92\x15\xb3\x06\x5a\xd9\xb6\xac\x56\x64\xe5\xc6\x25\x1c\x59\xa2\xd8\x3a\x1c\xf5\x8a\x99\x2b\x38\xe2\x09\x28\x11\x96\xe0\x18\x35\xc7\xc2\x0b\x8b\xbd\x66\x0e\x33\xe3\x02\xb5\x2e\xa4\x10\x43\x3b\xd7\xd3\x53\x7f\x7d\x57\x95\x93\xb2\x6b\x7e\x5a\xd3\xa3\x0a\x72\x70\x35\xd4\xc3\xae\x49\x70\x4d\xd8\x47\x7e\x36\xae\x51\x3d\x85\x90\xa5\x2c\xb4\xb7\x8f\xcb\x59\x85\x42\x59\xa4\x6c\x3c\xab\xba\xaa\x75\x70\x96\xc8\xbb\xa1\x53\x73\x26\xd6\x72\xb9\x5e\xa1\x7b\x05\x9d\xd2\x1b\x4a\x94\x21\x18\xb8\x72\x58\x18\x6d\xdb\x4a\xc3\x61\x91\xac\xb1\x60\x88\x37\xb6\x0b\xff\x27\x8b\xc0\xf9\x50\xb9\x70\x50\x2f\xfb\xab\x08\x87\x90\x97\xd3\xb6\xb4\x46\x47\x65\x75\xd0\x35\x7a\x59\xbd\xe0\x39\xa2\x54\xc1\x3f\x33\x28\x9a\xfc\xf8\x86\x4e\xd9\x2e\x3e\xbc\x07\x8b\xc2\xd2\xb5\xb2\x68\xaf\xd4\x71\x3e\xff\x94\x15\x94\xf5\x67\x37\x8f\xba\x25\x99\x9e\x79\xaa\x80\x01\x96\x3a\x19\x2c\xb4\x96\xd8\x30\x9d\xb0\x49\xc4\x53\xea\xc0\xfa\xd8\xc5\xcb\xb4\xac\x9b\x6c\xde\x73\xb5\x3f\xbb\x22\xc0\x6c\x3d\x35\xcd\x1b\x20\x6f\xcc\xb3\xe3\xbf\xe5\xfb\xbf\x8f\x4e\x9e\x6d\x1f\x7f\x88\x07\xaf\xbe\xfc\xb9\xf1\x69\x9c\xfe\xf8\xab\x60\xef\xc7\x7b\x6f\xca\x83\xf0\xdf\x6f\x44\x9e\xe1\x73\x30\x58\xf9\x7c\xb7\x9a\xe7\x8c\x99\x8e\x9b\x5e\xbd\x4b\x3f\xcc\x51\x3a\x16\x7c\xe2\xb8\x75\x24\x2c\xc0\x04\x6c\x0c\x18\xac\x62\xb2\x34\x05\xa6\x44\x50\x97\x60\xde\xd9\x7b\x59\x4c\xde\x49\xba\x3f\xd5\x27\xc7\x8f\x1f\x3f\x2e\x65\x7a\xfa\xfa\xe8\x79\xfe\xba\xf8\xf0\x64\x52\xeb\x8d\xe2\x4b\x71\xf0\x75\xd6\x14\x1b\x3b\xaf\xd7\xc6\xbc\x2c\xc7\xa9\x9a\x2f\xfa\xd7\x5a\x66\xbe\x0f\x99\x59\x1b\xa4\x88\x14\x73\x2e\x04\x16\x0e\x2c\x36\x2e\x32\x2c\x83\xe2\x51\xaa\xa8\xc2\xbc\x4a\xf9\x0e\xe4\xdd\xad\x43\x63\x5f\x7e\x3a\xf9\xaa\x5f\x7f\x7d\xfc\x21\xff\xf2\x8f\x7f\x5b\x47\x5d\x16\x4a\x96\xe5\xc7\x7f\x9e\x9e\x84\xad\x6a\xfb\xc5\xb6\x38\xda\xda\xdf\xb9\x86\x99\xe7\xe5\xb4\x65\x5a\x74\xad\x67\xbf\xe6\xa1\x1f\x66\x2d\x4c\x4c\x32\x05\x4c\x99\x93\x58\x44\x17\xb1\xd3\x14\xb0\x50\x32\x04\xaa\x94\x21\xc2\x5c\x82\xd9\xbe\xda\xab\xc5\xd1\x7e\x76\x92\x8e\x0f\xbd\x98\x8c\xf3\xc7\x1f\xdd\x5b\xf7\x91\xff\xb5\xbd\xf3\xe9\xa8\xfe\x48\x5f\xbe\x7a\xfb\xea\xf7\xb7\xd3\xf0\x62\x2c\xd6\xc3\xbc\x52\x0f\x1b\x26\x99\xe8\x62\xab\x57\xf5\xdf\x0f\xb3\x0d\x51\x68\x11\x24\x86\xe4\x78\x9b\x1e\x6c\xbb\xa4\x33\x2c\x4d\x8c\xde\x48\x06\xd1\x5c\x66\x67\xba\x6b\x4f\x76\xb6\xb6\x5f\x7e\x99\x40\xfa\x54\x3c\xdf\x70\x2f\xc2\x74\xeb\xcb\x36\x7d\x5b\x8f\x67\xaf\xf6\x77\x5e\x7e\x21\x7b\x5b\x13\x99\xc8\x1b\x6b\xfa\x63\x3e\x57\xd0\x5a\xca\xb5\x36\xeb\xd4\x1e\x0f\x30\x46\xbb\x6e\x32\xcd\xe1\xcc\x96\x0c\xbd\x85\xcd\x1a\xaf\x84\x34\x84\x48\x6c\x92\x20\x58\x28\xa6\xb1\xe7\xdc\x63\x0f\x9a\x38\x4d\x4d\x10\x66\x5e\xba\x5c\x98\xbe\xbd\xfb\xfb\xac\x6e\x90\x2b\xce\xcc\xdd\x94\xa8\xde\x2f\x8f\x50\xfb\x09\x65\x99\x77\x9b\x4c\xef\xcb\x3c\x87\xaa\x5d\x60\x36\x1f\xdd\x74\x27\xa7\x03\x37\xdf\x26\x9b\x55\xf9\x10\xa5\x2c\x87\x39\xd6\xd3\xad\xbc\xe1\xaa\x3e\x67\xe1\x4a\x6a\x65\x9b\xf9\xb0\x4c\x86\x61\x4a\x41\x62\x1b\x35\xc1\x09\x52\x32\xd4\x86\x18\x6c\x1a\x0c\x91\x38\xd3\x25\x55\x50\x75\xea\x8f\xca\x6a\xdc\xb5\x39\x0d\xd4\xcd\x67\x3a\x22\xa3\xce\x73\x97\x17\xbd\xb8\xbc\x03\x50\x94\x29\xab\x93\x13\x2d\x28\x7e\x73\x50\x7c\x01\x8a\xff\x4c\x50\x86\x10\xe1\xdb\x54\x20\x93\xd5\x37\x07\x25\x16\xa0\xc4\x75\x41\x2d\x7a\x97\x95\xd8\x0a\x65\x5e\x56\xe7\x31\x0c\x17\xf0\xcf\x05\x5a\x02\xce\xbc\x70\xad\x91\xc2\xe5\x78\x06\x6f\x5c\xdb\x09\x74\xad\x0e\x21\x5b\x5b\x84\xf4\xd6\xb7\xdb\x4a\xea\x41\xda\xed\xe2\x0a\x9e\x45\x00\x67\xae\x70\xbe\x95\xde\xe7\x3f\x84\x58\xfb\xe2\xc5\x3a\xb8\x7a\x79\xf8\x15\x15\xb3\x0f\xce\x29\xe3\xc3\x15\xa1\xdf\x36\x1e\x11\x3d\xdf\x62\x68\x56\xe3\x23\xa8\x1b\xcc\x16\x77\x4f\x9b\xb5\xf9\xde\x74\x3d\x44\xed\x53\x14\x4a\x57\x37\x3f\xb8\x6f\xeb\x4b\x5b\x2f\xdb\x5c\x41\x9b\x26\x5a\x30\x25\x1c\x66\x44\x13\x2c\xbc\x11\xd8\x2a\x88\xd8\x38\xa6\x6c\xf4\x4e\x12\x46\xcf\xd3\x06\xae\x6e\x30\xbd\x8c\xb6\xf6\xe9\x2f\x4f\x9b\xa7\x94\x18\xca\xdd\x55\xde\xb6\xe3\xf0\x33\x38\xec\xae\x9e\x74\x3d\xd8\x6a\x63\xb6\xa4\x6c\x49\xc0\x7d\xa4\xab\x57\x66\x3b\x47\xd7\xe2\x4d\xd4\x3c\xeb\x64\xd3\x95\x84\xb2\x78\xd4\xf9\x0f\x25\xa3\xf6\xcf\x85\xa6\xa6\x87\x3c\x5b\x91\xbf\x50\x2d\xf7\x90\xe7\x2b\xf2\xfc\x1a\xf2\x62\x45\xfe\x42\xe5\xda\x43\x5e\xae\xc8\x5f\xe8\x70\x7a\xc8\xab\x15\x79\x75\x0d\x79\xbd\x22\x7f\x61\x5b\xb5\x87\xbc\x59\x91\x37\xd7\x90\xb7\x2b\xf2\x17\x6a\xc1\x3e\xfe\x43\x56\x1d\xe8\xc2\x1e\x5e\x9f\x11\xce\xb8\xe0\xa5\x3e\xf8\x79\xb5\xfa\xfd\x56\x94\x9d\x16\x50\xa7\xc7\x16\xba\x4a\x65\x79\xbd\xbc\x98\x6f\x92\x2f\x7f\x9e\x75\xfe\xde\x01\xd8\x67\x65\xfb\x81\xca\xaf\x97\x3d\xee\x44\x79\x76\x6d\xe5\xfb\x17\xbd\xf7\x90\xf9\xfe\xc5\xed\x10\xf5\x5b\xd9\x7f\xa0\xf2\xf2\xe7\x2b\x7f\x7d\xb7\x51\x3f\x5f\x79\xba\xaa\x3c\x5f\x47\x79\xfd\xf3\x95\x17\xd7\x66\xde\xac\x15\xb0\x7d\xca\xb2\x9b\xf9\xbc\x5e\x47\x79\xfb\xf3\x95\x3f\xe3\xf3\x6c\xad\x45\x6a\xad\xde\xf4\x6e\xb4\xbf\xbe\xd3\xd3\xf5\xd6\xd8\x6b\x68\xef\x42\x93\x1d\x66\xcd\xf1\x52\xe3\xa5\x12\x45\x79\xf4\xf0\x11\x72\x0d\x6a\xb2\x09\xa0\x93\xb2\x00\x34\x98\x35\x61\x80\x70\x2b\x09\xd5\xa1\xcb\xd1\x80\xa3\xfd\x72\x56\x9d\xbe\x42\x11\x6b\xef\xc9\xf4\x0d\xd5\xbe\x0d\xcf\x99\xea\xe4\x76\x91\xaa\x25\x52\xd9\x19\xf1\xd7\x45\x4a\xd9\x12\xea\xe9\x89\xc4\x5f\x17\xaa\x59\x42\x15\xbf\xb8\xff\x32\xb1\x84\xca\xe6\x7d\xf6\xfd\x85\xba\x4c\xa9\x75\xe3\x9a\x59\xfd\x79\x3f\xab\x9b\xb2\xfa\x0e\xf2\x55\x8b\xf5\x62\xe1\x9b\xbd\x7f\x42\x49\x77\x15\x1a\x79\x4d\x34\x68\x75\x07\xe5\x1e\xe2\x52\xeb\xe2\xb2\xf7\xd9\x4a\xfa\x9a\x68\x90\xb8\xe7\x66\x62\x37\x07\x86\x28\x41\x35\x84\xb2\x88\xf7\x14\xa3\x58\x79\x17\xd6\xaf\xc0\x51\xf7\xd9\x17\xe5\x75\xe1\xfc\x5b\x52\x46\x7f\x60\xfa\x3e\xdb\x49\x5f\x17\xce\xbf\x26\x69\xdc\x04\xd9\x8f\xcc\x1a\x97\x1c\x53\x60\xb7\x70\x4c\x41\x1b\x12\x95\x15\x04\x5b\x27\xcc\xfc\x18\xa2\x31\x31\x61\x2b\xbd\x53\xdc\x7b\xb0\xc1\x7f\xff\x98\x02\xea\x76\x23\x9f\x14\x65\xb3\x0f\x15\xaa\x2f\x3c\x1f\xa2\x04\x90\xa3\x54\x01\xa0\xa6\x44\x15\x4c\xca\x43\x40\x13\xf8\x99\xa7\x14\x20\x79\x6a\x54\xb0\x38\xca\xe0\xb1\x88\x89\x61\xcb\x0d\xc3\x41\x18\x46\x41\xc4\x10\x84\x5f\x79\xf7\xbd\xb9\xb1\x91\x97\xc1\xe5\xfb\x65\xdd\x6c\x1a\x42\xe6\x67\xb1\x23\x4c\xca\xcf\x87\x74\x71\xe8\x7b\x79\xd1\x8b\xca\x3b\x00\xe5\x1d\x33\x2e\x09\x83\xbd\xf4\x16\x0b\x45\x00\x5b\xa1\x1c\x06\xef\x2c\x51\x01\x74\x74\x69\x0d\x50\x74\x01\x8a\x5e\x17\xd4\x0d\x5f\x7c\x3b\x1d\x4c\xb0\x4e\xe0\xc8\x1c\xc7\x42\xcb\x88\x3d\x28\x81\x2d\xa5\x29\x99\xa8\xc0\x5a\x7d\xee\x85\xfe\xb7\x83\xe4\x7d\xfd\xb9\x97\x27\x5c\x75\x54\x94\x3a\xce\x9c\xb3\x98\x4b\xc9\xb0\xb0\x29\x60\xc7\x4c\xc0\x4a\x0a\xee\x43\x48\x84\x28\xfb\x4d\x51\x84\x51\x9c\xbf\x90\xfc\xb0\x0f\x8b\xff\x23\x57\x37\xb3\x94\xd0\x51\x96\xe7\xc8\x03\x72\xf9\x91\x3b\xae\xd1\x3e\x54\xf0\x83\x5f\x49\xf6\xa5\xad\x97\x6d\x1e\xfd\xf6\xe0\x7f\x01\x00\x00\xff\xff\x5f\xc2\xa1\xad\xd9\x3a\x00\x00")

func dbSample_dataSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbSample_dataSql,
		"db/sample_data.sql",
	)
}

func dbSample_dataSql() (*asset, error) {
	bytes, err := dbSample_dataSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/sample_data.sql", size: 15065, mode: os.FileMode(420), modTime: time.Unix(1537795722, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x26, 0x13, 0x18, 0xfa, 0x7d, 0x45, 0xd6, 0xfb, 0xc0, 0xaa, 0x93, 0x65, 0x4f, 0x78, 0xb5, 0xe9, 0x97, 0x20, 0xe2, 0x17, 0x6e, 0xf4, 0xda, 0xf7, 0xbb, 0xb6, 0x18, 0xd0, 0xa4, 0x76, 0xf2, 0x24}}
	return a, nil
}

var _dbMigrations0001_initialSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\xc1\x6e\xe3\x36\x10\x3d\x5b\x5f\x41\xec\xc5\x36\xea\x00\xe9\x62\x93\x4b\xda\xbd\xb4\x5f\x50\xa0\x67\x62\x42\x8e\x6c\x22\x14\xa9\x92\x94\x1b\xe7\xeb\x0b\xca\x94\x44\x49\x94\xac\xb5\x5d\x34\xe8\xd5\x7a\x43\x0e\xdf\xbc\x37\x1a\xca\x0f\x0f\xe4\xa7\x42\xec\x0d\x38\x24\x7f\x96\x59\xf6\xf0\x40\x7e\xd3\x06\xff\xd0\x52\xa2\x21\x96\x1d\xb0\x80\x2c\x03\xe9\xd0\x10\x0e\x0e\x5e\xc1\x22\x61\xda\xa0\x09\x08\x74\xc4\x89\x02\x3f\xb4\x42\xf2\x2b\x59\x57\x8e\xad\x5f\xb2\x8c\x19\xf4\x2b\xe2\xbb\x43\x65\x85\x56\x44\xe4\x44\x69\x47\xf0\x5d\x58\x67\xc9\x97\xaa\x12\xfc\x41\x5b\x5b\x7e\xe9\xc0\x0e\x5e\x25\x12\x87\x50\x90\x4d\xb6\x12\x9c\x78\x10\x29\x8d\x28\xc0\x9c\xc8\x1b\x9e\x08\xc7\x1c\x2a\xe9\xea\x07\x74\x8f\x0a\x7d\xda\xf4\xf8\x6d\xb3\xdd\x65\x2b\x05\x05\x92\x23\x18\x76\x00\xb3\xf9\xfa\xb4\xad\xf7\x53\x95\x94\x84\x1d\x90\xbd\x91\x4d\x0d\xf8\xe5\x3b\x59\xaf\xb7\xa4\x52\xe2\xaf\x0a\x77\xd9\xea\xbc\x37\xa7\xce\xd6\xc7\xb0\x0e\x8a\xd2\x7d\xb4\x3b\xb1\xca\x18\x54\x8e\xb6\xcf\xda\x55\xb3\xed\x30\xf3\xca\xa2\xb1\x57\xa5\xee\x23\x2f\xa6\xdf\x82\x86\x47\xb0\xc8\x0c\xba\x36\xf6\xe9\x71\x1c\x1b\x20\xe7\xc8\x5b\x4e\xbd\xcb\x56\xbe\x40\xb4\x39\x62\xbb\x91\xc1\x1c\x0d\x2a\x86\x36\x54\x50\xf0\x2d\xd1\x8a\x70\x94\xe8\x90\x30\xb0\x0c\x38\xc6\xa4\x09\xc5\xf1\xdd\x43\x02\x6f\x61\xdd\x11\xad\x50\x96\x52\x30\x70\x5e\x45\x37\xeb\x22\x45\x4e\x44\xea\x2e\x5b\x71\xb4\xcc\x88\xb2\xde\xce\xe1\xbb\xfb\x0f\xd9\xf2\xca\xa8\x8b\xdc\x92\xb3\x23\x3e\xd9\xed\x58\x7b\x25\xb0\x37\xd8\xe3\x55\x04\xb9\x53\xe9\xab\xe1\x46\xc4\xd4\x0f\xbe\x93\x47\x0f\x3a\xa2\xa9\x7d\xdc\x29\x34\x21\xd1\x06\xd4\x92\x59\x19\x19\x45\x3c\x27\x44\x6d\x64\x87\xce\x85\xc4\x5e\xb5\x7e\x7e\x7c\x9c\x28\x89\x15\x1f\x91\x5b\x6a\xd4\x01\xec\xa1\xfd\xe9\xf9\xdb\xad\x32\x8f\x64\x37\x5b\xbf\x9e\x3c\xe7\xcb\xb8\xe9\xaf\xb9\x23\x81\xaf\x44\x3d\x7d\x87\xd5\x96\x02\xbb\x5a\xf6\x78\x44\xe5\x62\x8a\x5a\xf8\xba\xd4\xd6\x09\x65\x1d\x48\xb9\xde\x91\x6c\xc5\x0e\x46\x17\x7e\xbb\x64\x91\xdb\xb0\x1a\x6b\x0f\xf0\xf5\xe9\x79\xc0\xb3\x42\xe4\x96\x02\x2f\x84\x22\xaf\x5a\x4b\x04\xd5\x86\xe5\x20\xad\xa7\x40\x58\xca\x51\x3a\x98\x04\x70\x61\xfd\xd1\x69\x09\x27\xa9\x81\xd3\x57\x60\x6f\x3a\xcf\x47\x78\x67\xea\xae\x57\xa0\x03\xff\x42\xa2\x56\xec\x15\xb8\xca\x20\x35\x16\xfa\x6a\x8b\x52\xef\x05\x7c\xf4\x25\xd6\xc7\x71\x04\x2e\x85\x9a\x83\xdc\xa2\xab\xe0\xd4\x59\x4d\xb5\x6e\x5e\xde\x44\x07\x82\xe9\x76\x19\x4b\xeb\x00\x4a\xa1\xfc\xd7\xdf\xb1\x9e\x27\x2d\xb5\x49\xe2\x3f\x93\x39\x87\x15\xb9\x5c\x08\x3f\xf4\x84\x4c\x9a\x06\x3d\xb4\xf6\xa8\x4f\x77\x95\x6a\xf8\x9f\xa9\xd1\xde\xe8\xaa\xbc\x6e\x96\xb8\xf1\x75\x17\xb3\xec\x47\x3c\x5d\x39\x2a\x14\x2d\x8d\xde\x1b\xb4\x36\xed\xde\x9e\xbe\xb5\x14\xec\x44\xab\x92\x83\x43\x4b\x51\xf9\x03\xf1\xa4\x8b\x13\x61\x16\x72\xa4\x85\xe6\xb8\x34\x40\xe7\xb9\x60\x48\x0f\xba\x32\xcb\x93\x6b\x07\xd6\x86\xa7\x6f\xf5\x1b\x24\x3c\x2d\xd1\x08\xcd\xa9\x50\x0e\xcd\x11\x64\xaf\x87\x0e\xc9\x9c\x08\x69\xe9\x0d\xcf\x0b\x78\x6f\x19\x29\xd1\x04\xbc\x7f\xe9\xe2\x1e\xcd\xd4\xaa\x13\x51\xe1\x8d\xdc\x23\xba\x3e\x92\xae\xdc\x92\x64\x07\x11\x77\x19\x0a\xef\x6a\xc8\xe0\x90\x94\x21\x5b\xf3\xdc\xd1\x90\x8d\xd9\xba\x6d\x47\x7e\xac\xdf\x96\x8a\x35\xf3\x55\x6c\xaf\xd8\x98\x81\x67\xc1\x3b\x52\x45\x49\x84\x42\x77\x97\xc6\x37\x1e\x13\x9a\xbc\xa8\x75\xe0\xaa\xa6\x61\x58\x34\x02\x64\x9c\xd9\xa8\x75\x3f\xf6\x5b\xf1\x5c\x9f\x16\x2c\x9e\x08\xa2\xc0\x99\x74\x06\x33\xfb\x75\xc3\xe3\x2d\x7a\x0c\x7c\x04\x83\xed\xb2\x95\x04\xeb\x68\xbd\x13\xcd\xb5\x69\x6c\xf5\xe3\x0b\xd7\xeb\x04\x07\xed\x0d\xa8\x71\x82\x03\x50\xf2\xec\xe1\xc8\x29\x5c\x37\x3b\x9f\x7f\xff\xa1\xe6\xdb\x16\x60\x20\xd2\x94\x15\x3b\x4d\x4f\xf9\xf0\xae\xa6\xae\x5d\x96\xb2\x74\x63\xbf\x49\x47\xc7\x0e\xdb\x44\x27\xdc\x91\x7e\x82\x69\x77\xa7\x35\x19\x2d\xb3\x7d\x59\x18\x33\xd8\x6d\x69\x58\x73\xf0\x4b\xe6\xa5\x07\x61\x9d\x36\xa7\x39\x13\x8f\x64\x3d\xa7\xae\xbb\x1a\xea\xff\xa3\xac\xb9\x29\x7a\xb2\x24\x0b\xf5\x32\x0c\x5b\x2e\x99\x61\xe4\xa4\x6a\xea\x7b\x1d\xad\xaf\xe6\x33\x42\x69\xee\xf4\xbd\xf1\xc2\x8f\x74\x68\x7d\xa9\x13\x4f\xe2\x41\xb0\x77\xf3\x99\x6e\xf7\xe7\x2b\xe6\x4c\x16\x37\xdd\x93\x0c\x1e\x85\xae\xd2\x97\x52\x7f\xbf\x35\x46\x1b\xca\xfc\xa8\x38\xf8\x5c\xf0\x89\xa5\xda\x15\xcf\xaf\x35\x1a\xff\xa2\xe5\xe2\x32\xfb\xd5\x92\x72\x0d\x05\x98\x17\x67\x00\x5d\x94\x62\xc0\xf5\x52\x1c\x7f\x8c\x63\x4e\x1c\x85\x9b\x6d\x52\xb7\x54\x9d\x49\xb0\x36\x25\x4f\x8b\x47\x34\x7e\xe3\xc4\xb3\xeb\xc6\x8b\x4f\xd0\x83\xae\x9a\x75\xbb\xd8\x29\xa1\x2f\xd7\x77\x52\x55\x5d\x8d\x2f\x6a\xa6\x83\x46\xfd\x6a\x1a\xd4\x9b\xb0\xa7\x61\x7d\x3d\x27\x3f\x74\xd2\x66\xa9\x57\x09\xec\x4d\x0a\x5b\xf7\xa1\xdb\xbf\xad\x24\x2a\x92\x5a\xe4\x62\x69\x7a\x33\x4b\x97\xd6\x8e\x44\x14\xd4\xe4\xc7\xff\xbf\xfc\xae\xff\x56\x59\xc6\x8d\x2e\x9b\x09\x21\x6f\xfe\x30\xa9\xbf\x12\x87\xe5\x5f\xd2\x90\xf3\x67\xf4\x79\x4c\xac\xe5\x79\x64\x43\xd3\x3c\xaa\xff\xd5\xe9\x02\x36\x70\x36\x8f\x0a\x7e\x99\x07\xb5\xaa\x5e\x06\x6b\x6e\x47\x0b\xd1\xcb\x49\x9a\x7a\x7b\xcf\x47\x45\xad\x7d\x01\xf0\x52\x45\x1b\xd7\x2c\x2a\x67\xc2\x36\x6d\xdc\x3f\x01\x00\x00\xff\xff\x82\x87\x12\x03\x0a\x1c\x00\x00")

func dbMigrations0001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0001_initialSql,
		"db/migrations/0001_initial.sql",
	)
}

func dbMigrations0001_initialSql() (*asset, error) {
	bytes, err := dbMigrations0001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0001_initial.sql", size: 7178, mode: os.FileMode(420), modTime: time.Unix(1537795722, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf2, 0x23, 0x0, 0x88, 0xc5, 0x3b, 0x98, 0x60, 0xb7, 0xd, 0x75, 0xbc, 0x68, 0x62, 0x3a, 0x5, 0x31, 0xe0, 0xa7, 0x8, 0x87, 0x4f, 0xb2, 0x73, 0x96, 0x31, 0xb, 0xc5, 0x5f, 0x80, 0xdf, 0x70}}
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
	"db/drop_all_tables.sql": dbDrop_all_tablesSql,

	"db/sample_data.sql": dbSample_dataSql,

	"db/migrations/0001_initial.sql": dbMigrations0001_initialSql,
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
	"db": {nil, map[string]*bintree{
		"drop_all_tables.sql": {dbDrop_all_tablesSql, map[string]*bintree{}},
		"migrations": {nil, map[string]*bintree{
			"0001_initial.sql": {dbMigrations0001_initialSql, map[string]*bintree{}},
		}},
		"sample_data.sql": {dbSample_dataSql, map[string]*bintree{}},
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
