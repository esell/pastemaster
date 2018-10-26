// Code generated by go-bindata.
// sources:
// data/index.html
// data/show.html
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _dataIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xcd\x72\xdb\x20\x10\xbe\xe7\x29\x76\x36\x39\xc8\xd3\x8e\x48\xd3\x4b\x47\x41\x3a\xf7\xd6\xcc\x24\x2f\x80\x31\xb1\x71\x10\xa8\xb0\x38\xf1\x78\xf4\xee\x1d\x81\xec\x91\x22\xf7\x90\x8b\x2d\x96\x4f\xfb\xfd\x2c\x88\xef\xa8\x35\xcd\x0d\x00\x00\xdf\x29\xb1\xc9\x8f\x69\x49\x9a\x8c\x6a\x9e\xe3\xba\xd5\x04\x4f\x22\x90\xe2\x2c\xd7\x2e\x18\xe0\x41\x7a\xdd\x11\x04\x2f\x6b\xdc\x11\x75\xa1\x62\x4c\xec\xc5\x47\xb9\x75\x6e\x6b\x94\xe8\x74\x28\xa5\x6b\x53\x8d\x19\xbd\x0e\x6c\xff\x37\x2a\x7f\x64\x3f\xcb\x87\xf2\xc7\xb8\x28\x5b\x6d\xcb\x7d\xc0\x86\xb3\xdc\x6f\x14\xc4\xb2\xa2\xbc\x58\xbb\xcd\xb1\xe1\x52\x59\x52\x7e\xa2\x00\x00\x7e\xbb\x40\x15\x9c\x4e\x50\x42\xdf\xcf\x76\xf8\xab\xf3\xed\x1c\x9c\x1b\xaf\x7d\x93\x7e\x96\x5b\xa4\x3e\x48\x78\x25\x40\x6f\x6a\x6c\x55\x08\x62\xab\x10\xbc\x7b\x0f\x35\x3e\xdc\x23\x48\x67\x42\x8d\xbf\xee\x07\xb1\x67\x6c\x03\x5f\xa2\xd0\xb6\x8b\x04\x74\xec\x54\x8d\xeb\x48\xe4\x2c\x82\xb3\xd2\x68\xf9\x56\x63\x48\x79\xbf\x78\x61\x83\x90\xa4\x9d\x2d\x56\x08\x07\x61\xa2\xaa\x31\xcf\x02\xe7\x2d\x39\x5b\x9a\xe4\x03\xf9\x67\x6e\xde\x25\x4f\x5e\x85\x68\x68\x90\xdf\x4d\x86\xcd\xa6\xc1\x72\x96\xc2\x1e\x83\x1f\x47\x9c\xf5\x0e\x96\xd9\x5e\x1c\x44\xae\x4e\xb4\xbc\x46\x9b\xf4\xc2\x15\x07\x70\x9a\x29\x39\x08\x0f\x63\xb4\xf5\x5d\x81\xb7\xe7\x98\x57\xe5\x41\x98\x62\xf5\x38\x03\xdf\x95\xc3\xd9\x29\x4e\x8b\x1c\x07\x41\x15\xe0\xd3\x9f\xe7\x17\xfc\xbe\xd8\x8d\xde\x54\x80\xcc\xaa\xf7\x74\x74\xaf\x20\x36\x82\x44\x05\xa7\x91\xbc\x1a\xff\xfb\x25\x30\x44\x29\x55\x08\xd5\xc5\x62\x31\xa4\xf0\xd9\xd4\x45\x6f\x81\xb7\x63\xc6\xab\x72\xb8\x5d\x05\x1e\x5d\xf4\x60\xb4\x7d\x03\x1d\x2a\x40\xf8\x06\xa9\xc1\xe3\xe2\xfd\x2b\xe4\xca\x7b\xe7\x27\xd4\xca\xfb\xff\x31\x0b\xa3\x3c\x25\xc0\x95\xce\xb3\x4a\x3f\x41\xf4\xe3\xc8\xcf\x37\x8f\xb3\xfc\x49\xb8\xf9\x17\x00\x00\xff\xff\xf2\xeb\x14\xb5\x1b\x04\x00\x00")

func dataIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataIndexHtml,
		"data/index.html",
	)
}

func dataIndexHtml() (*asset, error) {
	bytes, err := dataIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/index.html", size: 1051, mode: os.FileMode(420), modTime: time.Unix(1540584027, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataShowHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xc1\x0a\x83\x30\x0c\x86\xef\x3e\x45\xe8\x03\x2c\xb2\xd3\x18\xb1\x97\x5d\x76\xdc\x2b\x54\x0d\x53\xd0\x76\x34\x81\x6d\x48\xdf\x7d\x68\x19\xda\x1e\x4a\xfb\xf3\xff\xc9\x97\xd0\xa0\xf3\x64\x2b\x00\x00\x1a\xd8\xf5\xf9\xb9\x7d\x75\xd4\x89\xed\xc3\x89\x32\xdc\x82\x57\xf6\x4a\x98\xc5\xec\xc7\x3d\x40\x6d\xe8\xbf\x87\x6c\xc7\x5e\x39\xee\xc2\x7a\xee\x41\xf4\x0a\xcb\x02\xa7\x21\x88\x7a\x37\x33\xa4\x54\x38\x08\xdb\x68\xb7\x0b\x4a\x5d\xf9\xa3\x2e\xb2\x83\xb1\x6f\xcc\xcc\x22\xee\xc9\x06\x62\x78\x4b\x63\xce\xb5\x81\x2e\x4c\xd2\x98\x4b\x6d\xec\x5a\xfe\xb5\x21\xa7\x44\xf8\xcf\x95\x20\xb4\x76\x69\x0f\x74\x84\x47\x5e\xc2\x3c\x0c\x61\xde\x4d\xf5\x0b\x00\x00\xff\xff\xc4\x65\x96\x9a\x24\x01\x00\x00")

func dataShowHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataShowHtml,
		"data/show.html",
	)
}

func dataShowHtml() (*asset, error) {
	bytes, err := dataShowHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/show.html", size: 292, mode: os.FileMode(420), modTime: time.Unix(1540583902, 0)}
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
	"data/index.html": dataIndexHtml,
	"data/show.html": dataShowHtml,
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
		"index.html": &bintree{dataIndexHtml, map[string]*bintree{}},
		"show.html": &bintree{dataShowHtml, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
