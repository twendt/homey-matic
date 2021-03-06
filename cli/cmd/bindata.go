// Code generated by go-bindata.
// sources:
// data/driver/assets/icon.svg
// data/driver/device.js
// data/driver/driver.compose.json
// data/driver/driver.flow.compose.json
// data/driver/driver.js
// DO NOT EDIT!

package cmd

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

var _dataDriverAssetsIconSvg = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4f\x73\xeb\xc6\x0d\xbf\xfb\x53\x6c\x99\x4b\x32\x25\xa1\x05\xf6\x1f\xe0\x5a\xce\x4c\xfb\x9a\x49\x0e\xb9\xb4\x69\x3a\x93\x1b\x1f\x49\x4b\x9c\x27\x89\x1a\x92\x96\xec\xf7\xe9\x3b\xbb\x24\x25\xd9\x7e\xed\x94\x17\x62\x7f\xf8\xb3\xc0\x6f\x81\x25\x1f\x7e\x7c\xd9\xef\xd4\xa9\xe9\x87\xb6\x3b\xac\x33\x04\x9d\xa9\xe6\x50\x75\x75\x7b\xd8\xac\xb3\x7f\xfd\xf6\x53\xc1\x99\x1a\xc6\xf2\x50\x97\xbb\xee\xd0\xac\xb3\x43\x97\xfd\xf8\x78\xf7\xf0\xa7\xa2\x50\x7f\xeb\x9b\x72\x6c\x6a\x75\x6e\xc7\xad\xfa\xe5\xf0\x65\xa8\xca\x63\xa3\xbe\xdf\x8e\xe3\xf1\x7e\xb5\x3a\x9f\xcf\xd0\xce\x20\x74\xfd\x66\xf5\x83\x2a\x8a\xc7\xbb\xbb\x87\xe1\xb4\xb9\x53\x4a\xbd\xec\x77\x87\xe1\xbe\xae\xd6\xd9\xec\x70\x7c\xee\x77\xc9\xb0\xae\x56\xcd\xae\xd9\x37\x87\x71\x58\x21\xe0\x2a\xbb\x9a\x57\x57\xf3\x2a\xee\xde\x9e\x9a\xaa\xdb\xef\xbb\xc3\x90\x3c\x0f\xc3\x77\x37\xc6\x7d\xfd\x74\xb1\x8e\xd9\x9c\x4d\x32\x42\x11\x59\x69\x5a\x11\x15\x7d\xfd\x54\x0c\xaf\x87\xb1\x7c\x29\xde\xba\x0e\xa7\xcd\xb7\x5c\x49\x6b\xbd\x1a\x4e\x9b\xab\xe5\xff\x67\x75\x3f\x74\x75\x7b\xec\xea\xf6\x62\xbe\x00\x30\x74\xcf\x7d\xd5\x3c\x75\xfd\xa6\x81\x43\x33\xae\x3e\xfd\xf6\xe9\xa2\x2c\x34\xd4\x63\x7d\x13\x66\xe1\xf3\xcd\xae\x6f\x48\x3e\x94\xfb\x66\x38\x96\x55\x33\xac\x16\x3c\xf9\x9f\xdb\x7a\xdc\xae\x33\xa7\x75\x5a\x6e\x9b\x76\xb3\x1d\xaf\xeb\x53\xdb\x9c\xff\xda\xbd\xac\x33\xad\xb4\x42\x43\x40\x82\xde\x5f\xa4\x30\x19\x5d\xbb\x04\x13\xd0\xd6\xeb\x6c\x38\x6d\x78\x5a\xcc\xdb\xdd\x5f\xcc\x34\x08\x81\x55\xdf\xbb\xba\xf4\x2c\x95\x41\x93\x2b\xd2\x28\x85\xc6\x02\xed\x0f\xc9\x6b\x29\xf5\xbe\xee\xaa\x98\xfb\x3a\x6b\xab\xee\x00\x91\xbd\xc7\x3b\xa5\x1e\xea\xe6\x69\x88\x76\xd3\x66\x71\x45\x99\x5a\x25\xd5\xc5\x35\xfa\xd5\xb1\x82\xab\xe1\xe7\x72\x98\xea\x56\xea\x58\x6e\x9a\xaa\xdb\x75\xfd\x3a\xfb\xee\x29\x3d\xb3\xe2\x73\xd7\xd7\x4d\xbf\xa8\x7c\x7a\xde\xa8\xba\x63\x59\xb5\xe3\xeb\x34\x15\x73\xec\xa5\xc8\x18\xf5\xa2\xd7\xdf\xd6\x0f\xdb\xb2\xee\xce\xeb\x8c\xde\x2b\xbf\x76\xdd\x3e\x46\x35\x2c\xec\x49\xde\xab\xab\x97\x75\x16\x2c\x88\x21\xf2\xf6\x83\xf2\x75\x9d\x11\x23\xb0\xb5\x1c\xde\x2b\xeb\xae\x7a\x8e\x73\x53\x3c\x1f\xda\x71\x58\x67\xfb\xfd\x07\xf7\xe7\xbe\x8f\x06\xbb\xf2\xb5\xe9\xd7\x59\x7a\xe1\x6c\x34\x6c\xbb\xf3\xa6\x8f\xf4\x3d\x95\xbb\x0b\x7f\x73\xa8\xe3\xcb\xfb\x50\xe7\xf6\x50\x77\xe7\x62\xee\x2c\x14\xfa\x40\xc2\x6c\xb1\x34\x1b\x6a\xfc\x90\xf1\x6c\xf2\xb2\xce\x0a\xfe\x2f\xba\xd7\xff\xa1\xdb\x97\x2f\xed\xbe\xfd\xda\xd4\xeb\x0c\x97\xbe\xd8\x37\x63\x59\x97\x63\x79\xed\x86\x05\x71\xa9\xa7\x94\x7a\xe8\xeb\xa7\xfb\x7f\x7c\xfa\x69\x5a\x29\xf5\x50\x55\xf7\xff\xee\xfa\x2f\xf3\x52\x29\x15\x0d\xca\xcf\xdd\xf3\xb8\xce\xb2\xc7\x0b\xfc\x50\x57\xf7\x4f\x5d\xbf\x2f\xc7\xc7\x76\x5f\x6e\x9a\x38\xe4\x7f\x7e\xd9\xef\x1e\x56\x57\xc5\x1b\xe3\xf1\xf5\xd8\x5c\x83\x4e\x61\xfb\x66\x1a\xf9\x6f\xde\x7b\x75\xb5\x6f\xa3\xd3\xea\x9f\x63\xbb\xdb\xfd\x12\x37\x99\xcb\xba\x09\xda\x8e\xbb\xe6\x31\xed\x39\x89\x4b\x15\xab\xb9\x8c\xb9\xc8\xd5\x4d\x95\x0f\xab\x85\x83\xb4\xda\xbc\x63\x73\x57\x7e\x6e\x76\xeb\xec\xef\x9f\x9b\x43\xa3\xf0\x3d\xd7\x9b\xbe\x7b\x3e\xee\xbb\xba\x99\xfb\x25\xbb\x32\xfb\xa6\x7f\xc6\xbe\x3c\x0c\x91\x86\x75\x96\xc4\x5d\x39\x36\xdf\xeb\xbc\x40\x6f\x21\x68\x36\xf4\xc3\xc2\xff\xb1\x1c\xb7\x4b\x4d\xc3\xf8\xba\x6b\xd6\xd9\x53\xbb\xdb\xdd\x7f\xa7\xd3\xf3\x97\x61\xec\xbb\x2f\xcd\xd4\x5a\xf7\x1a\x5c\x30\x3a\x38\x5c\x9a\x40\xa9\x78\xa6\x0a\x2d\xa0\x38\x6f\x42\x4e\xc6\x80\x71\xce\x92\x3a\xa9\xc2\x69\x70\x12\xd8\xa8\x9d\x22\xb0\x5e\x5b\xa2\xbc\x20\xb0\x68\x03\x2d\x08\x5e\x11\x87\x40\xde\x19\xce\x35\xa0\xf5\x26\x7c\x04\xd2\xdb\x52\xee\x34\x58\xed\x3e\xac\x51\xcd\xb1\x38\x9f\xa2\x87\x05\x08\x33\xc0\xea\x67\xe5\x1d\x78\x67\x34\xdb\x6b\xda\xea\x0f\xf5\xab\x42\x74\x20\xe8\x2d\xe7\xc4\x04\xac\x03\xaa\x4a\x15\x1a\xc8\x60\xa0\xbc\xd0\x60\x82\x15\xab\x0a\x0c\xe0\x58\x93\x33\x11\x8b\x7c\x70\xaa\x14\x43\x60\xe1\x2b\xf6\xb3\x4a\xf1\x2c\x7b\x13\xb9\xb0\x0c\xe2\x6d\x0c\xa9\xf3\xc2\x10\x68\xb2\x2e\xc4\xf0\x28\x4c\x3e\x2f\xac\x80\x66\xef\x8d\x9a\x03\x84\x09\x33\xc8\xc4\x11\x13\xd2\xe2\x52\x74\x2f\x98\x90\x10\xd8\x04\x9f\x4b\x00\x47\xd6\x25\x2e\x0c\x79\xc9\x85\xc1\x05\x34\xa8\x34\x78\xcd\x28\x98\x6b\xf0\xa2\x03\x2b\x06\x0a\xda\xb9\x5c\x03\xa3\x17\x52\x4e\x43\x70\x5a\x42\x24\x98\x4d\x3c\x34\x63\x80\xd0\xe7\x1a\x34\x5a\xa7\xac\x80\xa0\x73\x31\x95\x58\x9f\x66\x13\x21\xcf\xda\xda\x04\x39\x87\x14\xd4\x57\xb5\x57\xf1\xa0\x5d\x70\x79\x21\x0c\x22\x1c\xcc\x4c\x1d\xba\x89\x3a\xe7\x9d\xf3\xaa\x40\x0d\xec\xbc\x49\x75\x04\xf4\xc6\x4e\x3d\x62\x45\xe8\x06\x32\x06\x98\x8d\x71\x94\xeb\xa4\xd7\x1c\x82\x93\xd8\x04\xcc\x82\x2a\xd2\xe2\x25\x10\x9b\x3c\x52\x25\x46\x94\x06\xc3\x86\x7d\x4c\x3c\x0a\x41\x61\x00\x12\x6f\x6c\xac\xdd\x51\xf4\x8a\x27\x84\xce\x52\xc8\x35\x58\x12\x46\xb5\x8b\xc5\x38\xe3\x35\x4d\xf5\x59\x71\xa2\x26\x21\xd8\xdc\x32\x58\xb2\x96\xa7\x42\x34\x93\xcd\xc9\x83\x37\xc6\x7a\xa5\x41\x5b\xef\xa3\x09\x23\xda\xd8\x86\xc4\xde\xd9\xdc\x0a\x10\xdb\x64\x60\xbd\xb5\x36\xd7\x91\x3f\x1b\xf3\x13\x42\x91\xbc\x90\x18\x82\x29\x9e\x96\x75\x1c\x0f\x54\x02\xb0\x23\xc2\x44\xa3\x01\xb2\x88\x2e\x17\x0f\xcc\x2e\x38\x85\xe0\x75\x30\x94\x17\x08\xce\x88\x60\xea\x24\x01\xf6\x14\xec\x45\x8a\xf3\x55\x10\x68\xed\x42\xc8\x11\xc4\x5a\xed\x17\xc0\x5f\x80\x53\x2c\xd7\x3a\x76\x9c\x7a\x90\x02\x90\x26\xa2\xd8\x37\x41\x0c\xe6\x17\xad\x06\x23\x6c\xdd\x2d\x40\x28\xe8\x72\xad\x10\x90\x30\xd8\x48\x97\x17\x12\xaf\x96\x4d\xe6\xfc\x52\x11\x85\x05\x2f\xde\x61\x5e\xd8\x00\xc1\x5a\xe7\xd5\xef\x0a\xd9\x46\xb2\xc8\xc5\x01\xf4\xa0\x1d\x8b\xa5\x78\x4c\xda\xa3\xc6\x38\x21\x96\x01\x0d\x1b\xbc\x0a\xdb\x28\x8a\x04\x21\xb9\x48\x3c\x6d\xc1\x0e\xc4\x23\xa5\xce\x4d\xdc\x47\x62\x8c\x07\x23\x9a\x82\xda\x2a\xe3\x41\xbc\x31\x6e\x11\x2c\xab\x93\xba\xe8\x2f\xc2\x6d\x2e\x24\xa0\xc5\x39\x21\xf5\x87\xda\xab\x10\xaf\x31\xe3\x89\x73\x1d\xb3\x97\x00\x16\x59\xde\x3b\x78\xcf\xde\x87\x18\xda\x01\x63\xa0\xf9\x8d\x32\xa5\x90\x36\x71\x17\x69\xce\xdd\x73\xbc\x75\xc4\x61\xae\x53\xd6\x08\x64\xc8\x47\x0f\x04\xd6\xce\x07\x73\x23\x9d\xd4\xa2\x9f\xde\xee\x6d\x0e\xc6\x02\xb9\x38\xf5\x29\x69\x1b\x00\x8d\xb3\xc6\xe6\x26\x5e\x5b\x3a\x5d\x60\x0c\x8e\x59\x7c\x3a\x34\x9d\xe6\x87\x08\xbc\x35\xde\x98\x2b\xb4\x04\x8f\xe5\x5e\x3d\x16\x2d\x82\xf3\xe8\x29\x8d\xa0\x27\xcd\xf1\x66\x4b\xc0\x0d\x12\xe2\x05\x28\xd1\x46\x91\x05\xd2\x8e\xe6\x2d\x03\xc9\x8d\x7d\xb1\x38\x7c\x55\xbf\x2a\xf1\x10\x90\xad\x77\xe9\xb3\x41\xc6\xa3\xfa\x5d\x91\x36\x10\x28\x76\xf5\x9b\x42\x1d\x18\x11\x8b\xa2\x4e\x91\x78\x87\xc6\x84\xab\xb0\x55\x46\x83\x77\xe2\xc8\xdc\x48\x91\x6d\x22\xd0\xde\x98\x38\x69\x8e\x20\x56\xcd\x71\xa8\x6c\xc0\x34\x54\x01\x85\x28\x6e\x24\x10\x10\x99\xbc\xc2\xd8\x68\xd6\x4f\x43\x35\x5b\x52\xbe\x58\x2e\x88\x99\x11\xa3\xb6\xe9\x16\xc5\x80\xce\xdf\x48\x5f\x2f\xdf\xc7\xf8\x69\x8e\xdf\x58\x23\xc6\x5e\xc1\xcb\x3f\x60\x77\x38\x34\xd5\xd8\xf5\x45\xf5\xdc\x9f\xca\xf1\xb9\x6f\xd6\x99\x5e\x7e\x9f\x56\x9b\xc7\xbb\x87\xf8\x67\xf3\x78\xf7\x9f\x00\x00\x00\xff\xff\x35\x32\xd3\xa6\x1d\x0e\x00\x00")

func dataDriverAssetsIconSvgBytes() ([]byte, error) {
	return bindataRead(
		_dataDriverAssetsIconSvg,
		"data/driver/assets/icon.svg",
	)
}

func dataDriverAssetsIconSvg() (*asset, error) {
	bytes, err := dataDriverAssetsIconSvgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/driver/assets/icon.svg", size: 3613, mode: os.FileMode(420), modTime: time.Unix(1604862634, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDriverDeviceJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x4d\x6f\xdb\x30\x0c\xbd\xfb\x57\x10\xbe\x38\x46\x52\x75\xbd\xc6\xf3\x80\x6e\xcd\xd6\x02\xfd\x42\x93\xfb\x20\xdb\x4c\xa2\x4d\x91\x32\x89\xce\x92\x19\xfa\xef\x83\x1c\x3b\x75\x3e\xd0\xf6\xb2\x0f\x9f\x04\x3e\x52\x7c\xe2\x7b\x66\x54\x5a\x04\x4b\x46\xe4\x14\x25\x41\x90\x6b\x65\x09\xae\xf5\x02\x37\x90\x82\xc1\x1f\xa5\x30\xd8\x8b\xe6\x3e\x10\xc5\x49\x83\x5f\xe1\x4a\xe4\xd8\x4d\x60\xec\x9c\xb1\x73\x29\xb2\xf3\xa2\xc6\xd8\x37\xeb\xd3\x83\xaa\x02\x31\x05\xa5\x09\xd8\x5d\x29\x49\x34\x95\x67\xce\x05\x55\x75\xe6\x31\xf6\x89\x2f\x79\x26\xa4\x20\x81\xb6\x06\xb6\x3d\xf2\x36\xbc\xb9\xe3\x4b\x48\xa1\x0a\x00\x00\x7c\x91\xe1\x6a\x86\x20\x94\x20\xc1\xe5\x41\xbd\x73\x75\x5a\x58\x55\xcc\xb9\x70\xd8\x54\xd5\xa1\x7c\xce\x95\x42\x19\x0e\xe1\x62\xf0\x1c\xfd\x8e\x9b\x70\x08\xe1\x78\x72\x39\x19\x85\x9d\xf8\x8a\xcb\x12\x27\x9b\x25\x7a\x34\xd3\x5a\x22\x57\x61\x0d\xbb\xc1\x8e\x09\xaa\xa2\xd3\x51\x72\x4b\xfb\x74\xfe\x18\x85\xc0\xd5\xa3\x45\x69\xf1\xa5\x91\x35\x59\xaa\xd8\x0d\xbc\x3d\xe7\x92\x5b\x5b\xeb\xbc\xe0\x24\xf2\x46\x16\x5c\x13\xaa\xc2\xb6\xfa\x56\x41\xdd\x4d\xab\x1b\x25\xa8\x17\x77\x5e\xb2\xe2\x06\x44\xb1\x86\x14\x68\x2e\x2c\x9b\x21\x5d\x71\xe2\xbd\x98\x71\x22\x23\xb2\x92\xd0\xb2\x1b\x55\xe0\x3a\xd9\x95\xb4\x72\x77\x6d\xd0\xcc\x6e\x0b\x9f\x36\x43\xb7\xe3\x69\x4b\x74\x1b\xbc\xc1\x1a\xbb\xf1\x1e\x5b\xe4\x65\x9d\x5e\xd3\xeb\x0d\xba\xb5\x9f\x1b\x1c\x31\xef\x58\xa9\xc3\xf0\x55\x4b\xfd\x3d\xca\xc1\xf1\xa9\xe6\xed\x2d\xf8\x9a\x4e\x07\x25\xfb\x4f\x3d\x11\xb2\xe5\x12\x0d\x6b\x7c\xb7\x77\x5b\x9c\x34\x3f\x40\xbb\x3e\x14\x02\xbb\x2f\x17\x19\x9a\x8f\x25\x91\x56\x16\xde\xf9\x9b\xea\xac\xc6\x07\xe2\x17\x8e\xd6\x64\xf8\x68\x85\x8a\x6e\x85\x25\x54\x68\xec\x91\x9f\x2d\xca\x69\x63\xe8\x67\xd7\x4e\xb5\x81\x9e\x44\x82\xac\xbe\x1d\x52\xb8\x48\xda\xf3\xfb\xd4\xbb\xf6\xa0\xbb\x73\x2d\xde\xef\xc7\x07\x42\xf9\x0e\x2c\x33\xa2\x98\x21\xd3\xaa\x17\xa1\x27\x74\x16\x41\x7f\x8b\xd4\x5b\xf7\x46\x11\x9a\x29\xcf\xf1\x9e\x2f\x10\xfa\x10\x3d\xe3\xdb\xcd\x7a\x59\x14\x06\xad\xf5\xd0\xd0\x43\x0d\x19\x9f\xf9\xf8\x34\x1a\x8f\xbf\x8e\xaf\x1f\x9e\x26\xd1\x00\x7a\xb5\xa8\x31\xa4\x1f\x4e\xf8\x65\x7b\xa1\x11\x2b\x34\x8c\x8c\x98\xcd\xda\x17\x3c\xfa\xcb\xb1\xf8\x2c\xf5\xcf\x9e\x4f\x1a\x40\x05\xe1\xb6\x47\x38\x6c\x9b\xb9\x53\xd1\x01\x84\x4b\x5f\xdc\xda\xc8\xce\xb5\xa1\x10\x5c\xbc\xef\xa3\x38\xf9\x07\x33\xb9\x7d\xb8\xff\xf2\x1f\x8c\x44\x6a\x35\x7b\x71\x22\xae\x71\x77\xe7\x9f\xf0\xab\x7e\xa1\x8b\x52\x22\xc3\xf5\x52\x1b\xb2\x90\x1e\x2e\xee\x24\xf8\x1d\x00\x00\xff\xff\x5c\xf8\xbf\xae\xc7\x07\x00\x00")

func dataDriverDeviceJsBytes() ([]byte, error) {
	return bindataRead(
		_dataDriverDeviceJs,
		"data/driver/device.js",
	)
}

func dataDriverDeviceJs() (*asset, error) {
	bytes, err := dataDriverDeviceJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/driver/device.js", size: 1991, mode: os.FileMode(420), modTime: time.Unix(1604862634, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDriverDriverComposeJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x95\x4f\x6b\xe3\x30\x10\xc5\xef\xf9\x14\x83\xcf\x21\xb9\xf7\x56\xd2\xc3\xe6\xb2\x14\x96\x3d\x95\xb2\x4c\x6c\x61\x0f\x48\xb2\xd0\x8c\xc3\x96\xe0\xef\xbe\x48\x72\xfe\xf8\x5f\xc8\x36\xa1\xbd\x05\xe5\xcd\x7b\xef\x27\xd9\xf2\x61\x01\x00\x90\x51\x91\x3d\x41\x76\x38\xc0\xea\xc5\xd3\x5e\xf9\x9f\x68\x14\xb4\x6d\xb6\x4c\x7f\x5b\x34\x2a\x7b\x82\x24\x8e\x2b\xca\x4e\x0f\x44\x45\xdb\x8d\xe5\x1a\x99\xfb\xba\x4d\x58\xba\x70\xce\xd1\xe1\x8e\x34\x09\xa9\xa0\x7c\x7b\xef\xd6\xc9\x60\x19\x57\x2e\x32\x35\xfa\x32\xd4\xc8\xd6\x45\xf4\xe2\xf5\x28\x7f\x8d\xcc\x4a\x78\x9d\xc6\xd7\x71\x62\xe5\x6c\xd9\xc5\x45\x1b\x36\xa8\xf5\x7f\xd9\xc4\x89\x68\xd3\xc3\x73\x48\x3e\x74\x3e\x59\x9f\xbb\x5e\x6c\xaa\x26\x96\x3f\x3b\x4f\x45\xe0\x59\xf6\x15\xa2\x8c\xd3\x28\xea\xa4\x2b\xd4\x9e\xf2\xb1\xae\x76\x42\xb5\xed\x6f\xc7\x99\x87\x6c\xd9\x68\x0c\x55\xc4\x37\xaa\x27\x68\x07\x46\x16\xf7\x54\x62\x30\x9b\xf6\xb2\xea\xaf\x8c\xca\xf4\x1d\x17\x13\xde\xf3\xe0\x33\x40\xb7\x82\xdf\xda\x17\x8b\xe2\xde\xba\x97\x16\x57\xda\x4e\x26\xa5\x94\xe3\xb3\xcb\x4a\x84\x6c\xc9\xd7\x9e\x0d\xf9\x70\xd1\xae\xf4\x75\xe3\x86\x79\x1a\x77\x4a\x4f\x03\xa7\xf7\xee\x25\x16\x80\x02\x05\xb3\xab\xe7\x9d\x57\xa4\x0b\x1f\x87\xde\x46\x66\x63\xfb\xcb\xed\x70\xc3\x5a\xa3\xee\xa9\xe6\x8c\x68\x9e\x61\xc0\xf2\xec\x5c\x36\xa9\x69\x67\x9c\xf7\xa8\x9b\x98\xff\xa3\x36\xca\xa0\x50\x3e\x9e\x9f\x98\xbd\x4a\x9b\x6e\x82\xaf\x01\x4e\x37\xcd\x27\x99\xe7\xd2\x2b\xb2\x72\x4b\xf8\xd6\x82\xa9\x59\x20\x47\x56\x0c\x52\x29\x48\xe8\x10\xae\x78\xa0\xb4\xc4\xe1\x37\x76\x7f\xa7\x47\x2d\x6c\xc1\x12\x76\x8d\x00\xd7\xe6\x38\xc4\xc0\x8d\x73\xb5\x17\x30\x8d\x16\x72\xba\x27\xe7\x15\x6c\x2d\x48\x55\xb3\x3a\xc7\x7d\x80\xa1\xb2\x12\xb0\xb5\x80\x41\xc9\xab\xd5\xcc\x46\xdc\x7d\xa4\x58\x14\x5e\xf1\xf0\x5d\x3e\x89\x1e\x7c\xa6\x09\xfb\xb9\xcb\xfc\xe4\xd9\xde\xcd\x9c\xe7\xcd\xf6\xf5\x6b\x88\x37\x9b\xdf\xb0\x7d\xfd\x4e\xd2\x5f\xca\x13\xce\x82\x1c\x69\x25\x7c\x21\x1e\x01\xdb\xc5\x3d\x0c\xb8\xb7\xf2\x3e\xfc\x94\x2c\xda\xc5\xbf\x00\x00\x00\xff\xff\x19\xa1\x48\x70\x9b\x09\x00\x00")

func dataDriverDriverComposeJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataDriverDriverComposeJson,
		"data/driver/driver.compose.json",
	)
}

func dataDriverDriverComposeJson() (*asset, error) {
	bytes, err := dataDriverDriverComposeJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/driver/driver.compose.json", size: 2459, mode: os.FileMode(420), modTime: time.Unix(1604867793, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDriverDriverFlowComposeJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\xcd\x6a\xf3\x30\x10\xbc\xfb\x29\x16\x91\x63\x92\x07\x30\x7c\x97\xf0\xf5\x50\x28\xa1\xd0\xde\x4a\x0e\x0a\xde\xba\xa6\xb2\x64\xd6\xb2\x4b\x31\x7a\xf7\x22\xd9\x75\xfc\x27\xdb\x04\x37\xa7\xb0\x9a\x1d\xcd\x8c\xc6\x55\x00\x00\xc0\x34\x25\x71\x8c\x94\xb3\x10\xde\xdc\xc4\xfe\xaa\xf6\x9f\xc3\x24\x11\x0b\x81\x55\x15\x1c\xff\x53\x52\x22\x9d\x79\x8a\x60\xcc\x21\x23\xcc\x73\xb6\xef\x83\x75\xa2\x05\xb2\x70\xc0\xe1\x8e\x50\x5a\x9e\x53\xa1\xb5\x92\xf0\x50\xa2\xd4\xac\x07\x32\x43\x2a\xf5\x89\xb2\xaf\x6c\x5a\x61\xbb\x21\x79\x6a\xef\x66\x57\x77\xc7\x40\xda\x8d\xf7\x3b\x73\x28\x59\xa4\x57\x24\x2f\xca\x6b\x64\xd2\x10\x9b\x84\x99\xd1\xb4\x3f\xb9\x0c\x2c\x73\x8a\xff\xc2\x70\x63\xe5\x57\xea\x42\x2e\x11\xa9\x2c\x52\x5f\x5e\x5c\xc9\x45\x81\xd3\x32\x5b\xb9\xd5\x01\x76\xb5\xa8\x1c\xc2\x7f\x70\x3c\xbb\xac\x4f\xcd\xc4\x8c\x73\xe9\x6e\x12\x97\x31\xc2\xae\xe4\xc2\xee\x3e\x6a\x24\xae\xf1\xc6\x37\xb7\xed\x3d\x81\x5e\x93\x1d\xb7\x31\x1e\x87\x2d\x5e\xf0\x2b\x8a\xd9\x0e\xb4\xd0\xba\x0b\x1d\xea\xd9\x15\xbf\x03\x63\x03\x48\xde\x41\x36\x01\x74\x5c\xef\xed\x11\xca\x68\x29\xbe\x19\xc8\x65\xdc\xc7\x71\x04\x0b\x5d\x73\x9f\xfd\xab\xed\xca\x52\xdd\x9e\x2d\x12\x66\xa1\x1b\x56\x6e\xcd\xd3\xe7\x1f\x8a\xf4\xe6\xaf\xfe\xe2\x58\xef\x7d\x70\xbf\x9a\x55\x96\x84\x92\xf1\xe6\x8e\x9e\x2c\xe9\xbd\x86\xd6\x36\x2f\x98\x3e\xaf\xe7\x97\xc0\x04\x3f\x01\x00\x00\xff\xff\x08\xab\x54\xd2\xa4\x06\x00\x00")

func dataDriverDriverFlowComposeJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataDriverDriverFlowComposeJson,
		"data/driver/driver.flow.compose.json",
	)
}

func dataDriverDriverFlowComposeJson() (*asset, error) {
	bytes, err := dataDriverDriverFlowComposeJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/driver/driver.flow.compose.json", size: 1700, mode: os.FileMode(420), modTime: time.Unix(1604862634, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDriverDriverJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x41\x6f\xe2\x3c\x10\xbd\xe7\x57\xcc\xa9\x4e\x24\x30\xdf\x1d\xe5\x3b\x6c\xab\xd5\xae\xb4\xaa\xaa\x55\x6f\xab\x6a\x65\x92\x21\x98\x75\x6c\x76\x3c\x81\x56\x51\xfe\xfb\xca\x0e\x81\x50\xa0\xd4\x27\x3c\xf3\xe6\x79\x66\xde\x23\xa2\xf1\x08\x9e\x49\x17\x2c\xe6\x49\x52\x38\xeb\x19\xbe\xb9\x1a\xdf\x20\x07\xc2\xbf\x8d\x26\x4c\xc5\x2a\x04\x44\x36\xdf\xe7\x1f\x48\x6f\x91\xc6\x00\x29\x67\x52\xce\x8c\x5e\xcc\xca\x98\x93\x6b\x1f\xe0\x49\x61\x94\xf7\x91\xaf\x56\xac\x8b\x7d\x21\xbe\x32\xda\xd2\x0f\x3c\x6d\x92\x00\x00\x38\xfb\xdd\x6a\x4e\x33\x68\xe3\x35\x1c\xdf\x6c\x90\xe4\x90\x98\x1f\xe2\x6d\x3b\x05\xbd\x04\x79\xaf\x36\x6a\xa1\x8d\x66\x8d\x1e\xba\xee\x90\xe7\x95\xf6\xb2\x18\x27\x73\xf8\x75\xc8\x0e\x0c\xa4\x6c\x85\xa0\xad\x66\xad\xcc\x75\xb2\x70\x44\xdb\x82\x84\xae\x13\x93\x33\x16\xb4\xe5\x18\xbd\x6f\xcd\x28\xcf\xb7\x29\x2f\xa2\x44\xf2\x01\xfd\xcb\xe9\x12\xd0\x78\xbc\x39\xf8\xcb\x47\x84\xb1\x62\x35\x08\xf4\xfc\xb6\xe9\x6b\xe2\xc0\xbd\x3e\x8f\xaa\x0e\x6f\x88\xd1\xd3\xb1\xc8\xb8\x2a\xbd\x50\x2d\xd7\x4e\xdb\x54\x4c\x44\x36\x01\xb1\x52\x1e\x16\x88\x36\x6e\x19\xcb\xe8\x89\x63\x33\x61\x51\x16\x41\x3e\x36\xf5\x02\xe9\x4b\xc3\xec\xac\x87\xff\x60\xfa\xbe\xbf\xdf\x4b\xe3\x76\xcf\xa4\xab\x6a\x80\x3d\x11\x7a\x8f\x25\xe4\x60\x71\xd7\x1b\x56\x7e\x35\x6e\x77\xaf\xa8\xdc\x03\x1f\x70\xab\x0b\x4c\xcf\x27\x99\x6e\x42\xb1\xc8\x4e\xe4\x90\x84\x95\xf6\x8c\x94\x5e\x89\xff\x6c\xec\x8f\xf0\xc3\x22\xa5\xa9\xa2\xca\x4f\xc0\xb3\x62\xcc\x20\xff\x7f\x64\xd9\xe1\xe8\x25\x44\x94\x5c\xc4\x7e\x21\xcf\x7b\xf8\x70\xbf\xbb\x83\x98\x8e\xbd\x84\xc5\x1d\x11\x87\x50\x76\x81\x37\x1c\x42\x6e\xc8\xc2\x13\xb9\x5a\x7b\x94\x84\xde\x99\x2d\xa6\x4c\x0d\x66\x67\x05\x5d\x6f\x92\x4f\x32\xad\xb1\xe0\x74\xa9\x8c\xbf\xc4\x74\x12\xe9\xb2\x6b\xb6\xea\x7a\x8d\x6f\xeb\xcb\x17\x14\x0d\x22\xa6\x65\x94\x6e\x02\xec\xfe\xa0\x3d\x2e\xba\xfd\xa4\x2b\x4e\x05\xdc\xbf\x72\x85\xf4\x14\x5b\x28\x2e\x56\xbd\xab\x91\xc8\x51\x96\x1c\x07\x1f\x4d\xd9\x25\x49\xed\xca\xc6\xa0\xc4\xd7\x8d\x23\x0e\x7f\x99\x77\x1f\xb9\x79\xf2\x2f\x00\x00\xff\xff\x65\xe6\x70\xfd\x5b\x05\x00\x00")

func dataDriverDriverJsBytes() ([]byte, error) {
	return bindataRead(
		_dataDriverDriverJs,
		"data/driver/driver.js",
	)
}

func dataDriverDriverJs() (*asset, error) {
	bytes, err := dataDriverDriverJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/driver/driver.js", size: 1371, mode: os.FileMode(420), modTime: time.Unix(1604862634, 0)}
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
	"data/driver/assets/icon.svg": dataDriverAssetsIconSvg,
	"data/driver/device.js": dataDriverDeviceJs,
	"data/driver/driver.compose.json": dataDriverDriverComposeJson,
	"data/driver/driver.flow.compose.json": dataDriverDriverFlowComposeJson,
	"data/driver/driver.js": dataDriverDriverJs,
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
		"driver": &bintree{nil, map[string]*bintree{
			"assets": &bintree{nil, map[string]*bintree{
				"icon.svg": &bintree{dataDriverAssetsIconSvg, map[string]*bintree{}},
			}},
			"device.js": &bintree{dataDriverDeviceJs, map[string]*bintree{}},
			"driver.compose.json": &bintree{dataDriverDriverComposeJson, map[string]*bintree{}},
			"driver.flow.compose.json": &bintree{dataDriverDriverFlowComposeJson, map[string]*bintree{}},
			"driver.js": &bintree{dataDriverDriverJs, map[string]*bintree{}},
		}},
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

