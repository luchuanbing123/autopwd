package asset

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _key_autopwd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x61\x02\x9e\xfd\x30\x82\x02\x5d\x02\x01\x00\x02\x81\x81\x00\xc5\x7a\x3d\x95\x13\xf6\x24\x49\x81\xe7\xba\xe6\xa9\x70\xcb\x5d\xc3\xe1\xde\xc0\xcc\xc3\xe0\x5a\x94\xf2\x24\x7d\xa5\xbf\x64\x7a\xf3\x6d\x5a\xbc\xad\x99\x07\x14\x75\x37\xd5\x48\xe4\xc6\x90\xb6\x5c\x11\x1f\xfb\x83\x17\x77\x61\x32\xab\xf7\x34\x85\xb3\xc3\x8f\x8b\x17\x9e\xe6\x3c\xd3\x04\x82\x00\xdf\xd4\x0a\x5a\x64\x09\x88\xd7\x99\x37\xfe\xbd\x24\xef\x26\x8e\x26\x80\x5f\x99\x58\x2a\xb2\x3a\x1a\xc9\x58\x23\xb1\xbe\xeb\x83\x94\xa7\x4b\xa8\x7c\x7c\xf9\x24\x1f\x1c\x4b\xa6\xc6\x5d\xcc\x02\xc4\x80\x39\x38\x30\xea\x67\x02\x03\x01\x00\x01\x02\x81\x80\x07\xd8\xaf\x91\xb8\xd5\x52\xe3\xc8\xd1\x98\x4f\x89\xf3\xfd\x60\xa0\x63\xdd\x21\xf4\x00\xe4\x6d\x34\xf5\xda\x15\x53\xfc\xa2\xd5\x6f\xce\xac\x88\x4a\xad\x85\xe6\x7e\x31\xd5\xbb\xab\xbb\x68\x0c\x96\xe3\xe7\x4d\x6b\x0c\x07\xa6\x49\x09\x24\x2d\x66\x12\x9a\x76\x8b\xc0\xa0\x82\xef\xa1\x8d\xa4\xa1\x4d\xb7\x2d\x73\x43\x9d\xac\x3b\x5a\x9d\xff\x54\x30\x63\x07\xf8\x82\x39\xe7\xfc\x59\x3b\xeb\x82\x74\x9d\x2a\xa1\xe3\x5e\xc4\xd0\xce\xe7\xc1\xbe\x1c\x89\xde\x30\x62\xaf\x71\x78\xfa\x45\xce\x81\xf7\xd3\x9a\xf7\x03\x1b\xc1\x02\x41\x00\xd6\x12\x5c\x46\x44\x95\xfc\xe7\xa1\x10\x41\x43\x35\x22\xba\x2b\xa4\x4b\xee\xfd\xaf\x4c\xc7\x76\xb4\x7c\x48\xa9\x54\x60\x8b\x85\x5d\x87\xa1\x33\x06\x50\x26\xf3\x55\xd6\xe8\xce\x93\xb0\x42\xa3\xa0\xc7\x41\x90\x79\xa9\xa6\xc1\xfe\xe9\x16\x1c\x10\x73\x4a\x99\x02\x41\x00\xec\x27\xd6\xeb\x75\x3d\xb0\xe0\x14\xdd\x4a\x1c\x3d\x23\x54\xe2\x68\x80\xaa\xba\xd7\x9a\xe2\xd5\xa5\x8a\x85\x5b\x35\x82\xcd\x81\x2c\xc8\xac\xbd\xc9\x07\xde\xbc\x0d\xea\x84\x9e\x52\xb4\xa2\xab\x04\xa2\x31\x9f\x36\xf7\xde\xae\x04\xce\xad\x15\x6d\x23\xfc\xff\x02\x41\x00\xcb\x9c\xd5\x15\xa4\xe0\xc0\x97\x94\x33\xc0\xcf\x6a\xf7\xe7\xf9\xe0\xa0\xb4\xe5\xa8\x5a\x30\x12\x71\x23\x11\x7d\xb5\x9a\xca\xba\x04\x37\x20\x15\x47\x38\xef\x83\x5d\xaf\x36\xea\x5d\xc1\x35\xf1\x2a\xd8\xbb\x3c\xee\xeb\x75\x88\x69\xac\x4b\xf2\x4f\x97\x49\x69\x02\x41\x00\xce\x37\x0b\xc9\x98\x1b\x42\xb4\xaa\x03\x82\x8f\x9e\x23\x98\x20\x0c\xb9\x2e\x5c\xf1\xa0\xe2\xab\x45\x5e\xec\x8c\x74\xf1\xbf\x9a\x29\x0d\x23\x2a\xa6\xbd\x24\x96\x7a\x60\x6b\x66\x05\x3a\x8f\x41\x5f\x4e\x72\x9e\x61\x55\x99\x31\xc4\x28\xe5\xa3\x6c\x26\x9c\xe9\x02\x40\x12\x09\x6d\x49\xa1\x7e\xb3\xe9\x11\x76\xee\x1b\x94\x16\x58\xcc\x25\xc5\x54\x9c\x2f\x15\x8b\xc7\x30\x45\x87\xc7\xf7\x4c\x32\xdc\xce\x03\x85\xfe\x68\x5b\x97\x71\x69\x18\x83\x38\x61\xb1\x5b\xe1\x6a\x3c\xe0\x05\xb0\xa0\xb9\xbd\xfb\x7f\x49\x39\x56\x2d\x43\xd2\x01\x00\x00\xff\xff\x67\x09\xef\x20\x61\x02\x00\x00")

func key_autopwd() ([]byte, error) {
	return bindata_read(
		_key_autopwd,
		"key.autopwd",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"key.autopwd": key_autopwd,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"key.autopwd": &_bintree_t{key_autopwd, map[string]*_bintree_t{
	}},
}}
