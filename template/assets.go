// Code generated by go-bindata.
// sources:
// property/default/config.hcl
// property/default/field.tmpl
// property/default/method.tmpl
// type/default/config.hcl
// type/default/type.tmpl
// DO NOT EDIT!

package template

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

var _propertyDefaultConfigHcl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xcb\x4c\xcd\x49\x51\xb0\x55\x50\x02\x33\xf4\x4a\x72\x0b\x72\x94\xb8\x72\x53\x4b\x32\xf2\xc1\xa2\x10\x16\x54\x18\x10\x00\x00\xff\xff\xda\x37\xa8\x72\x2c\x00\x00\x00")

func propertyDefaultConfigHclBytes() ([]byte, error) {
	return bindataRead(
		_propertyDefaultConfigHcl,
		"property/default/config.hcl",
	)
}

func propertyDefaultConfigHcl() (*asset, error) {
	bytes, err := propertyDefaultConfigHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "property/default/config.hcl", size: 44, mode: os.FileMode(420), modTime: time.Unix(1488374780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _propertyDefaultFieldTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x57\xa8\xae\x4e\x4e\x2c\xc8\x2c\x49\xcc\xc9\xac\x4a\x55\xd0\x73\xcb\x4c\xcd\x49\xf1\x4b\xcc\x4d\xad\xad\x05\xca\x40\xb8\x2e\xa9\xc5\xc9\x45\x99\x05\x25\x99\xf9\x79\xb5\xb5\x5c\x84\xd5\x87\x54\x16\x80\xb8\x09\x59\xc5\xf9\x79\x56\x4a\x30\x51\x88\x22\xa5\x04\x2e\x40\x00\x00\x00\xff\xff\x1c\x2b\xdd\xa2\x74\x00\x00\x00")

func propertyDefaultFieldTmplBytes() ([]byte, error) {
	return bindataRead(
		_propertyDefaultFieldTmpl,
		"property/default/field.tmpl",
	)
}

func propertyDefaultFieldTmpl() (*asset, error) {
	bytes, err := propertyDefaultFieldTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "property/default/field.tmpl", size: 116, mode: os.FileMode(420), modTime: time.Unix(1488479462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _propertyDefaultMethodTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x51\xcb\x4e\xeb\x30\x10\xdd\xf7\x2b\x46\xd1\x5d\xb4\x5d\xb8\x77\x8d\xc4\xa2\x94\x22\x01\xa2\x48\x55\xf7\xc8\x4a\x27\xa9\x25\xc7\x09\xb6\x83\x54\x5c\xff\x3b\x13\x3b\xa9\x1b\x04\x2c\x61\xe7\xc7\x99\x33\xe7\xe1\x1c\xec\xb1\x10\x0a\x21\xe3\xba\x6c\x2b\x54\xd6\x64\xe0\x3d\x5d\x0c\xcc\x8d\xd5\x6d\x6e\xdd\x04\xc0\x39\xcd\x55\x89\xc0\xbc\x77\x8e\x6d\x78\x85\x70\x82\x9c\x37\xc2\x72\x29\xde\xd1\x7b\x42\xb0\xdd\xb1\xa1\x53\x40\xa3\xda\xd3\x89\xb0\x40\x27\xe2\x9b\x38\x37\xec\xd1\x98\xa3\x78\x43\x9d\x75\x54\xa2\x00\x61\x5e\x68\xab\x3e\x02\x23\xdc\x16\x4d\x2d\xe9\x93\x18\xa4\xc1\xb0\x6c\xf4\xd8\xd1\x0e\xec\x61\x1a\x5f\x21\xec\x7d\x14\xb4\x27\x7b\xbe\x79\x58\xaf\x76\x59\xf8\xfc\x77\xe0\x66\x39\x78\x82\xab\x6b\x28\x2d\x4c\xd9\x13\xda\x43\xbd\x4f\xef\x27\x90\xa8\x66\xf0\x9f\x46\x16\x0b\x12\x9e\x3c\x41\x8f\xed\xcc\x46\x7f\xf1\x7e\x8b\x26\xd7\xa2\xb1\xa2\x56\x34\x54\xb4\x2a\x87\xa9\x86\xb9\x73\x16\xab\x46\x72\x7b\x69\x31\x6a\x8b\x0c\xb3\x1f\xd8\xa7\xc1\xcb\x48\x71\xe7\x33\x31\x5e\x94\xf3\xd9\xc2\x39\x90\x59\xd2\xb8\x45\xdb\x6a\x15\xfb\x80\xd8\xdf\x28\xe9\xa4\x4a\x07\x24\x28\x21\x87\xc8\xfb\x17\xcd\xfa\x46\x23\x8e\x7d\x25\x3e\xae\x49\x8d\x74\xb1\x7f\xdb\xcd\xfd\x66\xb7\xde\xde\x2d\x57\xeb\xdf\xac\xe7\x6f\x22\x3f\xc7\xf0\x11\x00\x00\xff\xff\xa5\x2d\xa9\x13\x5f\x03\x00\x00")

func propertyDefaultMethodTmplBytes() ([]byte, error) {
	return bindataRead(
		_propertyDefaultMethodTmpl,
		"property/default/method.tmpl",
	)
}

func propertyDefaultMethodTmpl() (*asset, error) {
	bytes, err := propertyDefaultMethodTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "property/default/method.tmpl", size: 863, mode: os.FileMode(420), modTime: time.Unix(1488479494, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeDefaultConfigHcl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\xa9\x2c\x48\x55\xb0\x55\x50\x02\xd1\x7a\x25\xb9\x05\x39\x4a\x5c\x80\x00\x00\x00\xff\xff\x3a\x12\xfd\xa1\x13\x00\x00\x00")

func typeDefaultConfigHclBytes() ([]byte, error) {
	return bindataRead(
		_typeDefaultConfigHcl,
		"type/default/config.hcl",
	)
}

func typeDefaultConfigHcl() (*asset, error) {
	bytes, err := typeDefaultConfigHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/default/config.hcl", size: 19, mode: os.FileMode(420), modTime: time.Unix(1488404377, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeDefaultTypeTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x55\xdf\x6f\x9b\x30\x10\x7e\x9e\xff\x8a\x1b\xaa\x26\xa8\x32\xfa\xbe\xa9\x0f\x5b\x4b\xa5\x74\x2b\x44\x29\xe9\xcb\x34\x55\x0e\x38\x89\x57\x62\x53\xdb\x54\xca\x18\xff\xfb\x6c\x08\xc4\xe4\x47\xa3\x69\xad\xa6\x3d\x25\xbe\xfb\x7c\xdf\xc7\x7d\x07\x77\x76\x06\xf1\x82\x4a\x48\x78\x4a\x40\xff\xce\x09\x23\x82\x60\x45\x52\x98\xae\x60\x2e\x70\xbe\x78\xcc\xde\x9b\xac\xce\x20\x8d\xbe\x8c\x20\x8c\x62\x08\x2e\x87\xf1\x5b\x84\x72\x9c\x3c\xe0\x39\x81\xb2\xf4\x2f\x38\x9b\xd1\xb9\x3f\x6a\x22\x55\xf5\x11\x21\x44\x97\x39\x17\x0a\x5c\x54\x96\x74\x06\xe4\x11\xfc\x2f\x94\xa5\xe0\x44\x9f\xaf\x83\x8b\xd8\xa9\x2a\x04\x50\xa7\x18\xd7\x28\x2a\xef\x09\x53\x62\x05\x7e\xbc\xca\x49\x88\x97\xc4\x83\x1a\xe2\x10\xa6\x15\x50\x36\x3f\xfb\x21\x39\x73\xea\x4b\x84\xa5\x3a\xd7\xfe\x9a\x88\xc0\x4c\x2b\xf1\x87\x35\xa7\xac\x83\x26\xec\xaf\xd3\x0d\xd0\x3b\x2c\xe5\xa8\x10\xfd\xf4\xba\x5c\x1b\xaa\xaa\xf6\x74\x49\x64\x22\x68\xae\x28\x67\x1a\xa5\x74\x64\x0b\x27\x95\x28\x12\x05\xa5\x2d\xf3\x8a\x92\x2c\xd5\x2a\x6b\x81\xad\xba\x0a\xed\x90\x8c\x89\xe4\xd9\x13\x11\x20\xda\x3f\x33\x2e\xfa\x90\x3d\x94\xdd\xad\x1e\xb5\x7d\x67\xd3\xbb\x4e\xd2\x0d\x51\x0b\xde\x69\xb2\xf2\x47\xfa\x32\x2b\x58\x02\xae\x80\xd3\xbd\x12\x3c\xb8\xc1\x42\x2e\x70\x76\x7d\x1b\x85\xae\x07\xee\xb7\xef\xd3\x95\x22\x03\x20\x42\x70\x9d\x35\xd2\x04\x51\x85\x60\x60\xdc\xf5\xd7\x68\xf7\x9d\xf0\x7b\xf5\x3c\xd3\x9d\x63\x54\x13\xb6\xb4\xc8\x52\xac\x30\x34\x74\x5e\x43\xb7\xc3\xd6\x5d\xa8\xc1\x03\xd8\xc7\xba\x3d\x6c\xdb\x23\x34\x0c\xe3\x60\x7c\xf5\xe9\x22\x70\xfe\x66\x48\x28\x53\x44\xcc\x70\x42\xfa\x73\xd2\x37\xe5\x1f\x0d\x0a\x9c\xa8\x75\x00\x3e\x9c\x6f\xdc\x07\x6b\x7a\x4e\x72\x2e\x25\x9d\x66\xc4\x24\x6b\xd4\xc8\x0a\x34\xaf\xa3\xe5\x5e\x57\xd0\x76\x2f\xe6\x3a\x61\xd7\xa9\x2a\x33\x30\xa7\x3b\xd1\xf6\xca\x00\xa6\x9c\x67\xcd\x0c\x01\x24\x03\xe0\x0f\x86\xda\x78\x68\x11\xf8\xcf\x54\xf0\xd0\x1b\xe8\x26\xa2\x2e\xa0\x4b\xf5\xac\xde\xef\xf9\x24\x1c\x46\xe1\x3e\xbf\x5f\xc5\x06\xf8\x05\xba\x75\x38\xa7\x0a\x67\xf4\x67\x6f\x5a\xca\xff\xdf\xa1\x9d\xa7\x7b\x0d\xc3\x82\x70\x72\xd3\x7c\xe5\x9f\x69\xd5\x1f\xbd\xbc\x56\x83\x8c\x6d\x7a\x3d\xa1\x84\x33\xd9\xac\xbb\x75\xcb\x9f\x70\x56\x34\x34\x01\x2b\x96\x77\xe6\x24\xd7\xfb\xc8\xba\xae\x0f\x35\x50\x17\x3a\xef\x67\x5c\x67\x93\x73\x3c\x64\x6f\xb1\x43\xdf\xa2\xd1\x24\xbe\xdf\x2c\xb5\x17\xdd\x59\x43\x96\x17\xea\xc0\xe2\x3a\x24\x68\x1c\xdc\x46\x5f\xef\x82\xf1\xcb\x88\xd9\xf0\xfc\x0e\x00\x00\xff\xff\x1d\xf6\x64\xc8\xb9\x08\x00\x00")

func typeDefaultTypeTmplBytes() ([]byte, error) {
	return bindataRead(
		_typeDefaultTypeTmpl,
		"type/default/type.tmpl",
	)
}

func typeDefaultTypeTmpl() (*asset, error) {
	bytes, err := typeDefaultTypeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/default/type.tmpl", size: 2233, mode: os.FileMode(420), modTime: time.Unix(1488478841, 0)}
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
	"property/default/config.hcl": propertyDefaultConfigHcl,
	"property/default/field.tmpl": propertyDefaultFieldTmpl,
	"property/default/method.tmpl": propertyDefaultMethodTmpl,
	"type/default/config.hcl": typeDefaultConfigHcl,
	"type/default/type.tmpl": typeDefaultTypeTmpl,
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
	"property": &bintree{nil, map[string]*bintree{
		"default": &bintree{nil, map[string]*bintree{
			"config.hcl": &bintree{propertyDefaultConfigHcl, map[string]*bintree{}},
			"field.tmpl": &bintree{propertyDefaultFieldTmpl, map[string]*bintree{}},
			"method.tmpl": &bintree{propertyDefaultMethodTmpl, map[string]*bintree{}},
		}},
	}},
	"type": &bintree{nil, map[string]*bintree{
		"default": &bintree{nil, map[string]*bintree{
			"config.hcl": &bintree{typeDefaultConfigHcl, map[string]*bintree{}},
			"type.tmpl": &bintree{typeDefaultTypeTmpl, map[string]*bintree{}},
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

