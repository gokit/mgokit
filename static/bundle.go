package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct {
	path string
	root string
	data []byte
}

var (
	assets = map[string][]string{

		".tml": []string{ // all .tml assets.

			"mongo-api-backend.tml",

			"mongo-api-json.tml",

			"mongo-api-readme.tml",

			"mongo-api-test.tml",

			"mongo-api.tml",

			"mongo-solo-readme.tml",

			"mongo-solo.tml",
		},
	}

	assetFiles = map[string]fileData{

		"mongo-api-backend.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x6e\xa3\x30\x10\x86\xef\x3c\xc5\x1c\x13\x09\xc1\x2b\xec\x12\xb4\xd1\x5e\x9a\x48\x55\x4f\x55\x0f\xc6\xfc\x01\x37\xc6\xb6\xec\x49\x1a\x14\xf1\xee\x55\x02\x4a\x69\x15\xa2\x36\x3d\x59\x1a\x66\xe6\xff\xbe\x11\x69\x4a\xc7\x63\xf2\xc8\x7e\x27\x39\x59\x15\xaf\x90\x9c\x3c\x88\x06\x5d\x97\x67\x99\x90\x5b\x98\x92\x4a\x6c\x94\x41\x20\x41\xc5\x50\x79\xab\x95\xac\xc9\xc3\x79\x04\x18\x0e\xc4\x35\xa8\x52\x7b\x65\xaa\x28\x4d\xa9\x01\xd7\xb6\x0c\x84\x83\xb3\x01\x25\x15\xed\xb9\x21\xcf\x48\x35\x4e\xa3\x81\x61\xc1\xca\x1a\xda\x58\x3f\x1a\x25\x6e\x1d\xa6\x70\x92\xd3\xe2\x3f\x97\xf9\xe8\x56\xef\x07\xba\x32\x0c\xbf\x11\x12\xc7\x88\x28\x87\x06\x63\x26\xf9\x40\xd2\x1a\xc6\x81\x93\x45\xff\xc6\xe4\x76\x85\x56\xf2\x7f\x4e\x81\xbd\x32\xd5\x9c\xe0\xbd\xf5\x11\xd1\xc2\x43\x4c\x0d\x41\xa3\x19\x31\xac\x85\xdc\x8a\xea\xc4\x3a\xc1\x35\x6c\xa5\x88\x68\x09\xfe\x26\x08\xcd\x7e\x90\x10\x53\x1f\x31\x8f\x88\x9e\x5c\x39\x49\xfe\x25\xe5\x6e\x95\xde\xe4\xaf\xd6\x59\xbb\xf2\x25\xfc\xf5\x34\x7b\xfa\x34\x3c\x59\x3b\x32\x7b\x7e\xb9\xd3\x6d\x09\xce\xda\x7f\x0a\xba\xbc\x9e\xb8\x45\x7b\x51\xdb\x0b\xbd\xc3\xe8\x4f\xe8\x7e\x71\xd3\x5e\xf6\x86\xe5\x25\xf5\xb3\x6c\x4c\x4e\x54\x67\x88\x98\x3c\x82\xb3\x26\x60\x0d\xbf\x1e\x8a\xf7\xdc\xe2\xbc\x6b\x00\xeb\xa2\xf7\x00\x00\x00\xff\xff\x54\x7a\x5b\xb6\xc8\x03\x00\x00"),
			path: "mongo-api-backend.tml",
			root: "mongo-api-backend.tml",
		},

		"mongo-api-json.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\xcf\x4a\xc4\x30\x10\xc6\xcf\x0d\xe4\x1d\xc6\x3d\x48\x0b\x4b\xf6\xae\xf4\x05\x44\x76\xc5\xc5\x93\x08\x3b\xcd\x4e\xd7\xd6\x76\x22\x49\xea\x1f\x42\xde\x5d\x92\xae\xe0\x41\xd0\x5e\x12\x18\x92\xdf\xf7\xe3\x1b\x29\x36\x1b\xe8\x9d\x61\x68\xbb\x0f\x3f\x59\x72\xa0\x94\x92\xe2\x0d\x2d\x94\x52\x40\x08\x6a\xef\xed\xa4\xbd\xda\x35\x3d\x69\xaf\xb6\x38\x52\x3e\x62\xbc\xd9\xef\xb6\x50\xc3\x21\x04\x18\xf1\xf5\x1e\xf9\x68\xc6\x3c\x3b\x7f\x81\x55\xe3\x0c\xaf\x60\xd5\xe7\x2b\xc6\x83\x14\x95\x14\x39\xf2\xd6\xe0\xf1\x4f\xb6\x25\x3f\x59\x76\x80\xc0\xf4\x0e\x1d\x3b\x8f\xac\x09\x4c\x0b\xf8\x43\xec\x0e\xf5\x0b\x9e\x28\x46\xf5\x2b\x30\x46\x25\x45\x3b\xb1\xfe\x57\x66\xa9\x0d\x7b\x62\x0f\xce\xdb\x8e\x4f\x15\x94\x0b\x82\xd6\x40\xd6\x1a\x5b\x41\x90\xa2\x48\x0d\xd2\x40\xe3\x12\xd3\x54\x4e\xd1\xb5\x09\x03\x57\x75\xde\x8b\x7a\xe0\x11\xad\x7b\xc6\xa1\x7c\x7c\x6a\x3e\x3d\x7d\x1b\x56\x6b\xb8\x4c\xfc\xea\x3a\x3f\xbf\xa8\x81\xbb\x21\x27\x17\x73\x6f\x4b\x82\xc3\xec\x2e\x45\x31\x3b\x9c\x09\x89\xbf\x4e\x5c\x29\xf2\xfc\x2b\x00\x00\xff\xff\x39\x66\x83\xa1\x2e\x02\x00\x00"),
			path: "mongo-api-json.tml",
			root: "mongo-api-json.tml",
		},

		"mongo-api-readme.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x41\x8b\xdb\x30\x10\x85\xef\x06\xff\x87\x29\xb9\x38\x65\xab\xdc\x17\x7a\x48\x9d\x12\xf6\xd0\x36\x6c\xbb\xa7\xb0\x60\x45\x9e\xc8\xea\xca\x1a\x23\x8d\xa9\x21\xf8\xbf\x17\xd9\x59\xb2\xd9\x3a\x6d\x4a\xeb\x8b\xac\x41\xf3\xde\x37\x6f\x0e\x07\xf1\x95\x7d\xab\x58\x7c\xd9\x7d\x47\xc5\xe2\xb3\xac\xb1\xef\xe1\x13\x39\x4d\xab\x0f\xb0\xdc\xdc\xa5\xc9\xfb\x3f\x7f\x69\xb2\x7d\xb3\x5d\x13\xdc\x63\x43\x9e\x21\x97\xbe\x7c\xcc\x2a\xe6\x26\xdc\x2e\x16\x9a\xfc\x50\x56\xd2\x97\x42\x51\xbd\xd8\xc9\x52\xe3\xe2\x70\x10\x1b\xa9\x9e\xa4\xc6\x8d\xe4\xaa\xef\xe7\xbf\xe9\x18\xaf\xbf\xb6\xa4\x49\x9a\x5c\x31\x03\x98\x00\x12\x64\xcb\xf4\x4e\xa3\x43\x2f\x19\x4b\xc8\xef\x1f\x56\x60\xea\xc6\x62\x8d\x8e\x25\x1b\x72\xb0\x27\x0f\x5c\x21\x14\x93\xa2\x47\xe5\x02\x8c\x83\x66\xe4\x18\x5e\x6e\x9e\xb4\x18\x81\x0a\x11\x89\xbe\x55\x08\x7b\xb2\x96\x7e\x18\xa7\xa1\x46\xae\xa8\x04\xec\x4c\xe0\x30\x38\xa8\x36\x30\xd5\x40\x4d\x24\x31\xe4\xc2\x6d\xec\x9a\xcd\xe0\x63\x87\x2a\xfe\x16\x45\xa1\x29\x4d\xe2\x35\x53\xdc\x81\x22\xc7\xd8\xb1\xc8\xc7\xf3\x06\xf6\x1d\xec\x5b\xa7\x32\x45\x16\xde\xd6\x9a\x44\x4e\xd6\xa2\x8a\x62\x73\x40\xef\xc9\x1f\x8f\x41\xeb\x12\x53\x78\x86\x32\x6e\x98\xfa\x94\x4d\xcc\x4c\x06\x68\xd0\xb3\x34\x2e\x76\x30\x0d\x81\x3d\x93\xe6\x1e\x25\xe3\x0b\xd6\xb1\x30\x4d\x8b\x16\x6b\x38\x25\x7a\x5c\x61\xdf\x8b\x0b\xab\x7b\xcd\x3e\x9b\xc1\x1a\xf9\x85\xd9\x1a\x79\xda\xa9\x69\x77\xd6\xa8\xbb\x15\x04\xf6\xc6\xe9\x39\x64\x7f\x61\x7b\x73\x4c\xee\x95\x2f\x2c\xad\x3d\xf7\x5e\x5a\x3b\x65\x3f\x87\x6c\xfb\xf8\x8f\x7e\x0f\x4d\x79\x9e\xeb\x58\xb8\x6a\xda\xff\x12\xf4\x0a\x2d\x9e\x01\x8c\x85\x2b\xe3\x3e\xc9\xfd\x0c\x00\x00\xff\xff\x8d\x90\x2c\x35\x5a\x04\x00\x00"),
			path: "mongo-api-readme.tml",
			root: "mongo-api-readme.tml",
		},

		"mongo-api-test.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xef\x6f\xe2\x46\x10\xfd\x8e\xc4\xff\x30\xb5\x54\xc9\x5c\xa9\xef\xf2\x95\x8a\x0f\x10\xee\x92\x56\x0d\x44\x67\xd2\xfb\x88\x96\xdd\x31\xd9\xbb\xb5\x37\xdd\x1d\x43\x22\xe4\xff\xbd\xf2\xda\xfc\xba\x84\x1c\x26\x49\x4f\x4a\x1c\x29\x91\x62\x76\xe6\xcd\x7b\xfb\xe6\x09\xc4\x9c\x19\xf0\x9b\x0d\x00\x00\x9c\x63\x42\x16\xba\x10\x23\x19\xc9\x6d\x30\xc4\x85\xcf\x53\x4b\x3a\x0e\x42\x62\xfc\xdb\x40\xda\x1b\xc5\xee\x7c\x6d\x83\x90\x84\x4e\xa9\xd5\x6a\x36\x8a\x5a\xae\x93\x48\xce\xf2\x5a\x9d\xcc\x74\x70\xea\xfe\x5d\x16\xaf\xe5\x3f\x17\x5a\x60\x07\xe2\x99\x0e\x2e\x74\xa2\x49\x27\x92\xb7\x37\xaf\x0e\xfa\x1d\xd0\x36\x38\x43\xc2\x64\xee\x7b\xcb\x65\x10\x92\x49\x39\x05\x97\x8c\x7f\x63\x33\xcc\xb2\xc9\xc5\x68\x78\x36\x9a\x8c\x3f\x86\xe3\xc9\xa0\xef\xb5\xb6\x8a\xcf\xb5\xa5\x2a\xe5\xe7\xa3\x70\xbc\xd3\xe0\xca\xa2\xa9\xd2\xe0\x2a\xfc\xf8\x79\xa7\x41\x2f\xa5\xeb\x6a\x14\x7a\x57\xe3\xf3\xef\x68\x5c\x32\x6b\x17\xda\x88\x2a\x6d\x2e\x7b\x61\xf8\x65\xf4\x79\xb0\x6e\x94\xad\xee\x83\xd0\xd2\xa9\x56\xd0\x05\x6f\xb9\x54\x7a\x81\x06\x56\x9d\x46\xd3\xaf\xc8\x29\x18\xb2\x18\xdd\x9f\x2c\x9b\xe4\xa7\x27\x5c\x2b\x85\x9c\xa4\x4e\xbc\x66\xc3\x5d\xec\xfb\xf7\x30\x46\x4b\x67\x48\x9b\x39\xb6\xaa\xb3\x0c\xe6\x4c\x49\xc1\x08\x2d\xd0\x35\x82\xc9\x7d\x83\x73\xa6\x40\x47\xc0\x60\x4f\x91\xeb\x6b\x90\x6b\x23\x20\x32\x3a\x06\x56\x98\x46\x4c\x83\x66\x23\x4a\x13\xfe\x03\x50\x9f\xe0\x5d\x3e\xb0\x4c\x66\xc1\xb8\x55\x7a\x8c\xdd\x48\xe8\x74\x21\x16\x53\xe7\xda\x92\x7e\xbb\xf4\x74\xbb\xb4\xa5\x33\xb4\xb3\xe6\x96\x73\xe9\x36\xaf\xe4\x3a\x21\xbc\xa5\xe0\x8b\xa4\xeb\xb1\x8c\x51\xa7\xe4\xaf\x9e\x0d\x71\xf1\x0f\x53\x29\xf6\xd9\xcc\x6f\xb5\xe1\xe4\x03\xbc\x03\x92\x31\x06\x21\x72\x9d\x88\x75\x2b\x54\x18\xb7\x01\x8d\xc9\x1b\xfe\xad\x99\x78\x90\x41\x49\xe3\xaf\x70\x34\xf4\x7f\x7c\x37\xa7\x06\x19\x61\x7e\xb8\x55\x80\xc8\xc8\x21\xfc\xd2\x85\x44\x2a\xd8\x5a\xb1\x9c\xb3\x0d\x3e\x31\xa9\x50\xf8\x5e\x98\x72\x8e\xd6\x46\xa9\x52\x77\xa0\x34\x13\x28\x20\xef\x02\x91\x36\xfb\x6e\xa6\xbc\x95\x0e\xfc\xfa\xdb\xbf\x81\xe7\x98\xb4\xd6\xbe\xda\x40\xe4\x4e\x7d\x22\x84\xb7\xd6\x4c\x60\x84\x26\xbf\xbe\x60\x80\x0a\x09\x7d\x4e\xb7\x6d\xa7\x64\x70\x99\x4e\x95\xe4\x7f\x0e\xd6\x67\x4b\xea\x9d\xae\x3b\x5f\x28\xb3\x39\xdf\xfa\xa3\xba\x30\x4c\xe4\x43\xaf\xcc\xf8\xc8\xd8\x32\x21\x0d\x62\x7a\x94\x34\x55\x41\x82\x8d\x3a\x93\xc2\x4e\x05\xe1\x33\xa4\x87\xd5\x39\xd6\x16\xe5\xbe\xa2\x00\x4b\xda\x1c\x36\xa3\xdb\xd8\x23\x85\x78\x02\x9e\xd3\x24\xdb\x0d\xa5\x9e\x52\xc7\xe4\x92\x52\x4f\x4e\xa6\xfd\xc8\x75\x38\xd5\xe1\xf4\x56\xc2\xa9\x28\xb2\xed\x55\x4a\x75\xd6\x31\xd5\x53\xaa\xa0\xee\x31\xcb\xbd\x36\x78\x37\x4e\xad\x89\x14\x5e\x1b\x7e\x3f\xc9\x7f\x9f\x25\xb6\xf2\x5d\x2e\xa7\xf8\x1f\x42\xab\x22\xda\x96\x52\x32\x02\x85\x89\x5f\x16\xb7\xa0\xdb\x85\x0f\xd5\xc9\x92\x42\x66\x09\x4e\xaa\xc6\x66\x65\x9e\xc7\x02\x1d\x9c\xcf\x23\x23\xd0\xf4\xef\x7e\x5e\x4c\xf7\xef\xdc\x08\x75\x5a\xd7\x69\xfd\xe6\xd2\xfa\x5e\x54\xaf\x96\x61\x4f\x62\xd7\x49\xfd\xba\x93\x7a\xcf\xf9\x62\x27\xbe\x8b\x68\x9e\x3f\x94\x3a\x39\xf8\x03\xfe\x42\xd2\xf5\xc3\xf9\xfc\x28\x6c\x1d\xcc\x75\x30\xbf\xfe\x60\x3e\x60\x0b\xaf\x6e\xc4\xfd\x2d\x4c\x8b\x87\x2f\xb6\x83\x05\x68\xbd\x83\xf5\x0e\xbe\xfe\x1d\x2c\xba\x2e\x97\x60\x58\x22\x3e\x49\x54\x62\x65\x12\xf0\x72\xb6\x1e\x78\x5f\xad\x4e\xbc\xed\x77\x44\x90\x65\x0f\x4b\x55\x2e\xce\x7d\x69\x8f\x57\xae\xd8\xf5\x17\xd7\xae\x3a\xcc\xa1\x09\x56\xd8\xed\xde\x47\xbd\x58\x3f\xcf\xf7\x04\x8f\xa2\xd6\x11\xf6\x9a\x22\xec\xed\xc5\xd2\x2e\xe3\xfd\xc1\x7d\x04\x75\xb7\x81\x2f\x4f\xbe\x3a\xcc\x2e\xfd\x03\xbe\x00\x29\xc8\x77\xab\x90\x8f\xdc\x43\x20\x0d\x33\x24\x10\x4e\xd7\xca\x33\x1e\x26\xc0\xb3\x40\x65\xcd\xc6\x7f\x01\x00\x00\xff\xff\x15\x68\x91\xc8\xa8\x1f\x00\x00"),
			path: "mongo-api-test.tml",
			root: "mongo-api-test.tml",
		},

		"mongo-api.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x59\x6f\xdb\xc6\xf6\x7f\xb6\x01\x7f\x87\x53\xa2\x08\xc8\x56\xa5\xdb\xff\xa3\x5b\x15\xf8\x7b\xc9\x45\x80\x36\xc9\x4d\xd2\xdb\x87\xa2\x08\x46\x9c\x23\x99\x0d\x39\xa3\xcb\x19\xd9\x16\x04\x7d\xf7\x8b\xd9\xb8\x89\x92\x48\x89\xb6\x65\xc7\x7d\xa8\x42\x72\x96\xb3\x9f\xdf\x99\xe1\xd0\xa7\xa7\x80\x59\xc6\x33\x01\x61\x18\x9e\x1c\xdf\x90\x0c\xfc\x93\x63\x00\x80\xab\x2c\x7b\xcb\xe5\x6b\x3e\x63\x14\x86\xb6\x51\xf8\x16\x6f\x7d\x2f\xc3\x88\x67\x14\x18\x97\x30\x56\x8f\xbd\xe0\xe4\x38\x38\x39\x3e\x39\x3e\x3d\x85\xc5\x22\xfc\x28\xb3\x59\x24\xc3\x77\xa3\x7f\x30\x92\xe1\x5b\x92\xe2\x72\xf9\x3a\xc6\x84\x0a\xa0\x38\x8e\x19\x0a\x20\x0c\x62\x26\x31\x1b\x93\x08\xe1\xf6\x3a\x8e\xae\x01\xef\xa6\x5c\xa0\x80\x14\xe5\x35\xa7\x20\x39\x64\x28\x67\x19\x03\x02\x29\x99\x02\x1f\x03\x49\x12\x3d\x05\x91\x32\x8b\x47\x33\xa9\xc6\x11\x82\x47\x31\x91\x48\xe1\x36\x96\xd7\x20\xaf\xd1\xce\x41\x41\x68\x32\x66\x19\x02\x51\x13\x47\x31\x45\x0a\xa3\xb9\x6e\x93\x3f\x0b\x4f\x8e\xe5\x7c\x8a\x5b\xc8\x2e\x11\xbb\x38\x39\x3e\x32\x77\xfd\x00\xfc\x94\x4c\xff\x12\x32\x8b\xd9\xe4\xef\xbc\xc9\x62\x39\x30\xd2\x0a\x4e\x8e\x97\x9b\xa5\x72\xc1\x99\x98\xa5\x98\x6d\x92\x0b\x89\x22\x9c\x4a\x51\x88\x81\x12\x49\xec\xb3\xdb\x38\x49\x60\x84\x10\x99\x71\xa8\x9e\x2b\x66\x92\x6b\x26\x27\xf1\x4d\xcc\x26\x10\xa7\xd3\x04\x53\x64\x52\x5d\xf4\x21\x94\x9c\xea\xaa\x54\xec\xed\x35\x32\x09\x8c\x4c\x0a\x91\xfc\xce\xd9\x84\xd3\x82\xf3\xb5\x06\x41\x9c\x49\x8c\x79\xa6\x6c\x22\x8b\x51\xf3\x45\xf4\x30\xa9\x1a\x26\xbc\x24\x92\x8c\x88\x40\x20\x8c\xda\x5b\x1f\x51\x88\x98\x33\xc7\x8b\x9d\xae\x4a\xb2\xb2\xe6\x00\xfc\xef\xd2\xd2\x10\x03\xd0\x97\xb6\x7b\x6b\x55\x5e\x9e\x97\x58\x29\xa4\x6c\x58\x99\x66\xfc\x26\xa6\x08\x97\xe7\x70\xf1\xe1\x8f\x4b\xe0\x53\xcc\x88\x8c\x39\x13\x7a\xc8\x99\x50\xfc\x68\xb2\x95\x5a\x94\x2e\x66\x8c\x62\x96\xc4\x0c\x81\x8e\xb6\xa8\xe3\xf2\xdc\xce\xb7\x50\x7e\x1b\xf1\x04\x8c\xf4\xd5\x15\x1d\x59\xc6\xd5\x45\xaa\x64\x17\x09\xf7\x1b\xfe\x6e\x7e\xd5\x23\x64\x62\x96\x21\x7d\xc3\x28\xde\xc1\x88\xf3\x44\xdd\x8c\x59\xc4\x95\xed\x48\xac\xdf\xa7\x78\x87\x02\xfe\xfa\x5b\x09\x4a\x3f\x2b\xc4\xf3\x16\x6f\xad\xe3\x2a\x41\x30\xbc\x85\x98\x09\x49\x58\x84\xca\x76\xd7\xb2\x10\x9e\x1c\x8f\x67\x2c\x52\xdd\xfd\x82\x85\x01\xa4\x75\x6a\x07\x00\x29\xb7\x4c\x0d\x72\x52\xc2\x30\xcc\x69\x09\xe0\xbb\xb5\xd3\x2c\x4c\x68\xb3\x91\xe5\xd5\xb6\x76\xea\x3f\x3a\x3a\x83\x94\x0f\x8a\x1b\x11\x4f\xce\xd4\xff\x4a\xb7\x2c\x8d\x67\x90\x96\x6e\x5a\xda\xce\xdc\x3f\xec\xa3\x65\x21\x2b\x23\x75\x23\x5c\x22\x25\xa6\xca\xcf\x25\xb7\xf7\x55\xbc\x73\x96\x43\x73\x4e\x73\xef\x16\x53\x8c\xe2\x71\x1c\x29\x52\x12\x8c\xa4\xb1\x75\x2d\x44\x3f\xa5\xa3\x0d\x42\x08\xca\x13\xfb\xd6\x31\xc1\xb2\x1c\x8f\x21\xa5\xa3\xb0\x62\x10\x25\x69\x58\xc9\xb1\x38\xc9\xb9\x51\x3f\xe6\xff\x29\x9c\x0d\x73\x7d\xbd\xc5\xdb\x4f\x19\x89\xd0\xf7\xd6\x2b\xbd\x44\x87\xca\x21\x5a\xdc\x38\xc6\x4c\x93\xe0\x06\xba\x4a\x63\xe9\xbb\x8b\x37\x6c\xcc\x5b\x8f\x38\x70\xbd\xfe\x8c\xe5\xb5\x21\x26\x0d\xaf\x18\xf5\x83\x40\xa7\x2b\xcb\x6f\x82\x4c\x49\x2c\xb4\x22\x0e\x60\x38\x84\x1f\xb7\x30\x6d\xfe\x71\x7a\x0a\x6f\xc6\x70\x8b\x70\x4d\xa8\x8a\xdd\x46\x92\x23\x1c\xf3\x0c\x8d\xc6\xe0\x96\x08\x70\x6e\x34\x50\x8a\x63\x20\xbe\xc4\xd3\x81\xea\x15\x11\x26\x55\x16\xcd\x07\x13\x92\x4f\xb5\xda\xf9\x54\xc0\x08\x23\x32\x13\xda\x6d\xc6\x24\x4e\x9c\x0d\x84\x39\xdd\xdf\xac\x28\xea\xd5\x2b\x30\x8c\x54\x1d\xb7\x0d\x2b\x34\x0f\x7f\xa2\x14\xf9\xb4\x42\xe9\x28\xa4\x23\x9d\xf9\x83\x7c\x6e\xf5\xec\x9b\xa1\x1a\xa7\x3c\xfa\x5a\xb5\x5d\x29\xc1\x8c\x7d\xef\xb5\x61\x44\x72\x88\x32\x24\x12\xdd\x64\x3a\xb6\xc7\x0d\x5a\xf3\xbd\xc2\xbc\xbd\x81\x9e\x20\xe2\x49\xbd\x8d\x96\xbb\xa7\x29\x36\x53\x19\x0d\xd7\x98\xc6\x2c\xab\x33\xad\x8d\xcd\xd2\x10\x5e\x24\x5c\xa0\x9f\x5b\x46\x31\xb1\x92\x82\x93\x4f\x78\xe1\x3b\x22\x5c\x43\x45\xfb\x67\x1b\x8b\x54\xd3\x8c\xb0\x09\x42\xc9\xa2\xca\x22\xb2\xb2\x3b\x1b\x96\xfd\xf6\xaa\xe4\x8f\xba\x4f\xf0\xf3\x1a\x09\x77\x94\xb2\x8d\x24\x4e\xca\xbb\x4b\xd8\xf4\xb4\x4c\xb6\x14\xff\x2a\xd5\x75\xc3\x1c\x82\xcc\x66\x58\x6d\x57\x57\x56\x45\x61\x1b\xd9\x37\xb1\xe1\xe3\x2c\x8a\x10\x4d\xc8\x34\xfc\xab\xe4\x5a\x52\x66\x5f\x42\x08\x6a\xc6\xb4\xe2\x8d\x8e\xbb\xe2\xf1\x06\xb2\x5f\xc7\x2c\x16\xd7\x48\x81\x50\xaa\x51\x5b\x7b\x2a\x83\x4a\x52\xd3\xae\xed\x32\xcc\x05\x9f\x31\x59\xc9\x2d\xb6\x95\xca\x20\x92\x4b\x92\x00\x9b\xa5\x23\xcc\x54\x94\xb1\xb0\x7e\x9c\xf1\xd4\x00\xe9\x51\xeb\x84\xa2\xe7\xf1\x23\x79\xa7\xa0\xa8\xc4\x3b\x19\x5e\x98\xdf\x00\xfc\x98\x49\x87\xa2\x9c\x25\x77\x4e\x14\x7a\xfc\x9e\x52\x84\x1d\xab\x65\x72\x70\xfc\xbc\x11\x57\x77\xd3\x38\x43\xaa\xb8\x0c\xca\x2e\x69\xdd\x79\x9c\xca\xdc\x01\x2d\xf3\x70\x4d\x84\x82\xb1\xaa\x9b\x17\xb4\xb2\xe1\x55\x17\x9e\xa0\x74\x9a\x89\x1a\x28\xef\x33\x40\xfe\xf0\xd3\xa0\x21\x48\x16\x11\xab\xb0\x70\x0b\x1d\xd6\x45\xa9\x0e\xec\x91\xe9\x34\x99\xf7\x1f\xfa\x5b\xf2\xf6\xd0\x59\xef\x3e\x95\xd9\x96\xe5\x4d\x39\xef\xbf\x33\xcc\xe6\x8a\xfd\x91\xe0\x2c\xfc\x7d\x61\xe1\x9d\x8e\x14\xb9\x68\x1a\x52\x61\xf8\x3a\x66\xd4\xd7\xbd\x03\xe3\x60\x7e\xf0\xf3\x81\x89\x4d\x53\xe7\x0d\x0c\x8f\xfd\xca\x74\x4b\x2c\xba\x44\x95\xf2\xa8\x65\xa1\x07\xe2\x73\xca\x5c\x34\xcf\xf5\x53\x84\x7e\x33\x69\x2d\xf6\xa7\xfc\x06\x75\x6c\x5f\x8d\xf5\xb6\x0a\x55\x17\x79\xd1\x31\x9d\x8d\x92\x38\x7a\x73\x19\xea\x11\x3f\xe8\x3e\x22\x6f\x18\x0b\x55\xd0\xa6\x33\xa1\x02\xdd\x0d\x02\xb1\xed\x21\xa6\x70\x43\x92\x19\x0e\x54\xf0\xcb\x50\x08\xa4\x80\xb1\xbc\xc6\x0c\x46\x73\x20\xda\xb8\x80\x67\xf0\x8f\xfa\x95\x64\xa2\x47\xe7\xcc\x2d\x5a\x20\x2b\xd5\x89\xef\x49\xf4\x85\x4c\x70\xb9\x0c\xd7\x44\x74\x5b\xfc\xb6\x4e\x55\x46\x2e\x4d\xb9\x6a\x90\xf3\x6b\xcb\xcf\x5a\x69\xd4\x39\x6b\x99\xa9\x7a\x4a\x5b\x6e\xb0\x9a\x61\x38\x92\xbd\x82\xfa\xc7\x48\x6d\x3b\x38\x35\x35\x06\xba\xc6\x29\xb6\xf2\x75\xbf\x95\xc1\x93\x4c\x7a\x5b\xb9\x7a\xe8\x74\x77\xd0\x2a\xee\x94\x08\x8b\xf1\x0a\xb2\xcf\x72\xb2\x07\x6b\xad\xa7\x29\x57\x7e\xd0\x61\xd8\x66\xcb\x1e\xac\x69\xb3\x94\xf7\x4e\x8d\x2d\xd4\xb4\x45\x05\x56\x1c\xc3\x21\xa4\x13\x1e\x96\xf7\x15\x16\x8d\xb5\x5f\xa9\x45\xa5\x06\xdc\xaa\xd2\x87\xce\xc3\x2d\x24\x55\x4f\xd5\xd5\xf2\xcc\x2c\x82\x94\x73\x34\xa1\xb4\x9c\xa0\xf3\xe5\xbe\xe6\x04\x5d\x5e\x5b\x95\xd7\x58\x5b\xa0\xde\x9a\x3b\x1f\x31\xaf\xef\x95\xc3\x8d\xdc\x9a\x73\x38\x26\x98\x76\x91\xc1\xbe\x49\xde\xd0\xd2\x57\x6d\x6a\x07\x5b\x6f\x58\x8a\xbd\xf0\xfd\xd3\xca\xf4\x76\xad\x6f\x7b\x1a\xd8\xc0\xdc\x4b\xba\x3f\xfc\x74\x5f\x5d\xd3\x3d\x48\x45\x6f\x4c\xfa\x8b\x85\x12\x83\xaf\x2c\xfe\xb5\x8a\x43\xd6\x4b\xc1\x33\x1b\xbe\x1e\x40\x00\xcb\x52\x1e\x1a\xeb\xdb\xb9\x50\x35\x53\x6e\x6f\x78\x25\xfb\xb5\x5c\xcb\xad\x3e\x86\x62\x73\x69\x4d\x59\x9c\xef\x45\x8f\x55\x0c\x5b\x13\x5e\x73\xcf\x5b\x3f\xfa\x26\xd1\x6f\xeb\xa5\x18\xb7\x4a\x6d\xd1\xb8\x41\x73\xb5\x4e\x41\xc7\x15\xe1\xcd\x70\xeb\x0d\x13\x98\x49\xdf\x00\x39\xdf\xe8\x2c\xe8\x6b\x81\xdd\x9a\xfc\x56\xc1\xef\x60\xe1\x65\xa1\x76\x32\xfe\x56\x32\xdb\x92\xa2\x2e\x36\x86\xec\xae\xf4\x07\xce\xbf\x30\x11\x58\xf6\xa0\x1a\xcc\xf6\x17\x0b\xfd\x96\x43\xee\x78\x7a\x10\xf0\xd4\x53\x0f\xbc\x7f\xf4\xcf\x72\x19\x74\x56\xfe\x66\xac\x7d\x30\x3a\xdf\x65\x89\xea\xb0\xb4\xbe\xb2\x4e\x65\xf5\xce\xe8\x72\xb9\x09\x09\xff\x0b\xe5\xff\x27\x89\x7b\xbd\x03\x85\xde\x05\xcd\x2c\x30\x2d\xaf\x52\x11\x46\x4b\xef\x17\x88\x24\xae\xbf\x58\xb0\x7d\xc1\x48\xce\xa7\xf8\x54\x91\xaf\x91\x53\x33\xf2\xe5\x19\x55\xb9\xcd\xbe\x39\xa1\xaf\xce\xe7\xf9\xf5\x94\x4c\x10\xf4\x5e\x4c\x86\x62\xca\x99\xc0\xf7\x98\xbd\xb7\x37\x03\x00\xff\xaf\xbf\x3b\x08\x71\x00\xd0\xc7\xbe\x8e\x61\xa7\x27\xf0\xec\x06\xdb\x06\x8a\x8f\xc4\x6d\x2c\xa3\x6b\x2b\x19\x11\x7e\xe2\xbf\xf1\x5b\xcc\x7c\x2d\x31\xcd\xcb\x51\x44\x04\x82\x47\x45\xe4\x0d\xc0\xa3\x28\x22\xef\xac\x70\x22\x27\xd9\x21\x78\x3f\x78\xf0\xbd\xbb\x3e\x39\x3e\x5a\x1e\x0e\xe6\x76\x9e\xb4\x57\x08\x6f\x05\xb2\x58\x9c\x0c\x9e\xd3\xd6\x51\x3b\xf6\x8e\xe2\xb1\x71\xa9\x5f\x86\xf0\x23\xbc\x7a\xb5\xe2\x55\xbf\xd8\x97\x48\x8e\x8e\x6c\x18\xab\x80\x6f\x63\xab\xe7\xf3\x77\xca\x76\x94\x65\x58\x87\xcd\xfd\x36\x30\x3d\x35\x09\xf9\x00\x09\x32\xdf\x5e\x04\x96\x20\x63\x73\x47\x26\x84\xae\xd9\xd2\x15\xe1\xc9\xf1\x91\x7e\xf4\xa1\x81\x94\x7c\xef\x36\x30\x5c\x55\xf5\x92\x13\x51\x95\x83\x9d\xf6\x86\x64\x66\xce\x3f\x09\x93\xe8\x5e\xc8\xfa\xc4\x3f\x4a\x92\x49\x15\x21\xea\xa2\xfa\xa9\x49\x54\xbf\x3a\x49\x95\x86\x82\x61\xbd\x99\x6a\x50\x19\x7e\x08\x3f\x2a\x42\x40\x81\x8a\x16\xfd\xe1\x3b\x4d\x45\xc3\x30\xe5\x6e\xa7\xf0\x7f\x9a\xe6\x9c\xe8\x5f\xe1\x27\x33\x78\xa5\xd7\xf7\xdf\xab\x5b\xcb\x5c\x10\x1b\x70\x7c\x6d\x0d\xea\xfc\xec\xdf\x2a\x37\x9e\x19\x0b\xb0\xb4\x79\xc1\x00\x56\x7b\x28\x33\xb5\xf8\xde\xdd\xd2\x97\x35\xe8\xe2\x09\x45\x51\xcc\x26\x9f\x8d\x2f\x9c\x39\x64\x54\xa2\xb7\x86\xb0\x3d\xcd\xf2\x67\x6b\x1e\x9f\x6f\x35\xef\xde\x59\x45\x97\xb5\x1e\xda\x2e\xf3\xb1\xa1\x12\x0c\x1b\xdb\x9e\xcf\xeb\xad\xed\xed\x7a\x6b\x25\xe6\xd5\x81\xb5\xf4\xeb\x4d\x6b\x2a\x75\xbd\x6a\xb7\x4b\xbd\x96\xae\xb6\x08\x0e\xb4\x28\x7e\x90\x38\xdc\x75\xcf\xd7\x3c\x50\xce\x9d\xc5\x12\x53\x01\x9d\xb0\xc1\x96\x62\xda\xbe\x27\xbc\x52\x4d\xab\xe9\xa8\x9b\xae\xf9\x25\xe2\xb6\xa8\xbf\xbc\x1b\xfd\xf1\x4b\x3c\xf5\xcb\xae\x10\x84\xbf\xc5\x4a\x65\x25\x5b\x0f\xc2\x8f\x3c\x93\xbe\x0b\xbd\xa1\x82\x58\xaf\x0c\x2d\x7d\x15\x0d\x79\x3e\x2e\x03\xdb\xf5\x6f\xc4\x6a\x90\x6a\x80\x2f\x1d\x3d\x4a\x2d\xd1\x7a\xd1\xbe\xc9\x04\x1b\x57\xf0\xa1\xba\x8a\xbf\xd1\x74\xa1\x5e\xb9\xb8\x97\xec\x24\xa6\xc5\x3b\x76\xd6\x5c\x6a\x04\x29\x43\xea\xba\x0a\xdc\xc8\xbb\x5b\xd3\x71\x6f\xb6\xab\xd9\x36\xda\xc3\x56\x96\x9a\x44\x60\x78\x18\x2a\x38\x83\x8c\xfa\xe6\xda\x16\xce\x2b\xdb\x1f\x8b\x85\x49\x77\xcb\x87\xf5\x85\xec\xc5\x17\x1e\xdd\x17\x9c\xfa\x19\x85\x95\x5a\xda\x19\x4d\x15\xe4\x35\x54\xd8\x16\x71\xbe\x14\xda\xed\x0a\xed\x12\x40\x5f\x53\x6f\xd7\x0b\xed\x5d\x2a\xe9\x5e\x8a\x68\x4b\x6a\xaf\xb5\x74\x3e\xe6\xf3\x2b\xa9\x9f\x42\x51\xfd\x8c\x0b\xea\x83\xdf\xac\xba\x7f\x7e\x77\xc2\xe4\x07\x83\xaa\xfb\xc1\xcb\x5d\x76\xbb\xf6\x44\x0d\xf7\xb3\xf5\x55\x43\x16\xfd\x6f\x7e\xed\x0c\x3d\x76\x81\x1d\xeb\xe1\xf7\x7e\xb5\x20\x3c\x55\x00\xdf\x1f\x78\x5f\x09\x0b\xae\x69\x7e\x14\xab\x09\xda\x2b\x49\xec\x25\xf5\xbd\x3c\xfa\x05\xf4\x1f\x80\xe7\xd9\x56\x4d\xd6\x62\x2b\x81\x12\xc2\x3f\x9f\xeb\x05\xc2\x32\xbc\x6f\xf9\xaa\xb7\xde\x17\x87\x2f\x38\xd7\xc0\x5f\x63\x70\x3d\xa8\x2b\x01\x54\xe3\xaf\x08\xfb\x5b\x41\x36\x03\x7f\x25\x25\xb7\xad\xa6\x29\x86\xca\x41\x77\xf0\x1f\x1a\xfe\x5b\x6a\xfb\xc3\xfe\xf9\x80\x55\xb7\xf9\x82\x73\xcb\xf1\x13\x79\xf3\x6c\x0b\x60\x5f\xc7\x4f\xbf\xe7\x70\x3a\x58\x83\xfd\x76\xc4\x73\x40\xfd\xfd\x8b\xe1\xb0\x2a\x84\x27\x62\x3b\x9d\xca\x8c\x2f\x38\x3f\x33\x3c\xed\x57\x70\x68\x90\xb7\xae\xd8\xd8\x05\x9c\xbc\x63\x68\xe0\xc8\x0b\x1a\xe9\x80\x46\x3a\x1a\x4f\x47\xdc\xb2\xab\x69\xae\x20\x9c\xdd\xe0\x7e\x93\x19\x75\x03\xfb\x7d\xf3\xe1\xde\xbe\x4a\x30\x5d\x03\xeb\x1b\xbc\x64\x6f\x9e\x9f\xb6\xeb\x6c\x74\x8b\x1e\xd6\x60\x9e\xb9\xeb\x94\x8a\x83\xad\xb5\x41\x87\xa2\xc0\x1d\x83\x2c\xef\x01\x7c\x65\x05\x40\xcb\x03\xa2\x0f\x8f\xf5\xfb\x03\xf9\xcf\xe8\x2c\xe9\xb6\x85\xf9\x07\x3d\x6a\xf8\x75\xe2\xfd\xbe\x85\x70\x58\x68\xff\x89\x59\x50\x27\xd4\x6f\x79\xfb\x1c\xd3\xd2\x91\xd6\x97\x02\xe0\xa5\x00\x78\x28\x14\xf3\x52\x00\xbc\x14\x00\x2f\x05\xc0\xfd\x16\x00\x7f\x4c\xa9\x4a\x6f\x33\xf1\x02\xff\xb7\xc2\x7f\x23\xab\x56\x15\xc0\xc3\x9f\x37\x37\xc4\xf5\x54\x05\xb8\xc1\x9e\x4d\x21\x30\xd6\x9f\xd1\x1b\x38\x32\xaa\x23\xee\x10\x5d\x0a\x68\xb4\xf3\xd7\x2f\x9e\xe8\xa1\xf3\xfb\xe2\xfc\x39\xa3\xfb\xfb\xb4\x96\x6d\x27\xd7\xbf\x9d\x9a\x2d\x5c\xe5\x5a\xea\x1f\xe7\x73\xe5\xe5\x05\x5c\x77\xa7\xed\x3d\x28\xbd\xa1\xff\xed\x54\x92\x89\xea\x32\x41\xf9\x89\x4c\xf2\x41\x2a\x27\x6e\x5d\xfb\x95\xba\x61\xb1\xd0\xfd\xc3\xff\xe8\x1d\x83\x65\xa7\xea\xe1\xe5\x48\xfd\xc6\x5e\x87\x7d\xa4\xde\x26\x48\x6d\x10\x03\xab\xb3\xbe\xd0\xe5\xcc\x00\x95\x8e\xa7\xab\x5b\x7c\x97\x67\xb3\x5b\x3e\x22\xd4\xec\x88\x1d\x3b\x9c\xee\xae\x29\xa1\x02\x0a\xea\x67\xbd\x9b\xdb\xb6\xb7\xe1\x0e\xef\xc5\x55\x9b\x2a\x03\xf3\x72\x2b\xda\xd8\xb4\x59\x91\x45\x8f\xcd\xdf\x19\xb8\x24\x92\xec\xf2\xad\x81\x5d\x7c\x22\x9f\xf0\x71\xdd\x62\xef\x85\x09\xab\x9c\x82\x9d\xfe\x33\xde\x23\xb9\x0d\xac\x7e\x98\x60\x0b\x8c\xb6\x25\xd4\x1e\xa2\x6e\x23\xad\xcd\x5f\xfa\xac\x1c\xe9\xb8\xba\xc3\xc8\xbd\xc1\xa5\xaa\x3a\x55\xdd\xe8\x4f\x5c\xdb\xbf\x1c\x92\x24\xfc\xd6\x54\x6e\x78\x87\xd1\x4c\x3f\xe2\x63\x20\x10\xcd\x84\xe4\x69\xd1\x9e\x4c\x48\xcc\x84\xd4\x4d\x77\xf8\x73\x02\x8a\x8e\xe6\x6a\x69\x7c\xa7\x27\xd1\x7f\xd2\x41\xff\x59\x8d\x8b\x7c\xf4\xc0\xed\x85\xec\x57\x0e\xa9\xb9\x7b\x2a\x86\xcc\x50\x4f\xa3\xce\x31\x0a\xc5\xe2\x2f\x89\x3c\xe8\x87\xa5\x0e\xbb\x72\x79\x36\x55\xc9\xa3\x7d\x08\xab\xd0\xef\xf8\xce\x6f\x48\x76\x7d\x68\xf9\x41\x2c\xf8\x50\xbe\x3c\xf9\xce\x31\xe9\xd8\x56\xf9\xa3\x0d\xc8\x5a\x17\xff\xff\x17\x00\x00\xff\xff\xde\x3d\xe6\xb1\xb7\x6b\x00\x00"),
			path: "mongo-api.tml",
			root: "mongo-api.tml",
		},

		"mongo-solo-readme.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xcd\x4a\x43\x31\x10\x85\xf7\x79\x8a\x03\xdd\x54\x29\xf7\x01\x04\x37\x56\x85\x2e\x5c\xbb\xed\x34\x99\xe6\x0e\xe6\xce\x94\x64\x6a\xef\xe3\x4b\x5a\x7f\x10\x8a\x66\x33\xe1\x70\xe6\xfb\x60\x5e\x4c\xb3\x3d\x3e\x60\x33\x1d\x0a\x4f\xac\x4e\x2e\xa6\xe1\xfe\xef\x17\xc2\xf5\x3d\x48\x03\x21\x5a\x62\x64\x56\xae\x97\xf0\x40\xf1\x8d\x32\xc3\x47\x72\x1c\xaa\xbd\x4b\xe2\xde\xdb\x51\x93\x08\xf9\x6d\xde\x5b\x85\xa8\x73\xa5\xe8\xa2\x19\x27\xf1\x11\x84\xa3\x26\xae\xa2\x8c\xa9\x8b\xd3\x0e\x89\x9c\x76\xd4\x78\x08\xe1\xd9\x4a\xb1\x53\x2f\x4f\xec\xa3\xa5\x06\xaa\xfc\xc3\xe5\x74\x17\xc2\x62\x81\x57\xf1\x71\xa3\x89\xe7\x10\xb6\xdb\x6d\xb6\xf0\x1d\x2c\xa3\xcf\x88\xa6\xce\xb3\x0f\xeb\xcb\x5c\x21\x5a\x41\xf3\x2a\x9a\x57\x90\x5e\xe3\x86\x61\x18\xa6\x6c\xc3\x79\xeb\x06\x5c\xab\xd5\xcf\xd1\x99\x67\xcd\xd3\xcc\xf1\xcb\xd0\xff\xff\xc3\xf7\x33\xf6\x47\x8d\xcb\x9e\xdd\x76\xfc\xda\x4a\xe1\xd8\xcf\x71\xcd\x11\x3e\x02\x00\x00\xff\xff\x3a\x8b\x2f\x3b\xb4\x01\x00\x00"),
			path: "mongo-solo-readme.tml",
			root: "mongo-solo-readme.tml",
		},

		"mongo-solo.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x4d\x8f\xdb\x36\x10\x3d\x57\xbf\x62\xea\x43\x21\x05\x06\xdd\xb3\x0b\x5f\xfc\x11\x60\x0f\xd9\x02\x4d\x8b\x1e\x0b\x8a\x1c\xd9\x04\x24\xd2\x20\xa9\x58\xc5\x22\xff\xbd\x18\x7e\xc8\xb2\xb0\xdb\x75\x8a\x62\x37\xbe\x18\x14\x39\x6f\xde\xbc\x19\x3e\xae\x56\xf0\xc9\xe8\xa3\xd9\x6f\x41\x62\xa3\x34\x3a\xe0\xa0\xb4\x47\xdb\x70\x81\x70\x39\x29\x71\x02\x1c\xce\xc6\x85\x9d\x0e\xfd\xc9\x48\x68\x8c\x05\x8b\xde\x2a\xfc\xa2\xf4\x11\x78\xb1\x5a\x41\x47\x30\x6c\xcf\x3d\xaf\xb9\x43\xe0\x5a\xa6\x4f\x9f\xd1\x39\x65\x34\x2b\xfc\xdf\x67\x1c\xb3\x5d\x73\x3c\x15\x3f\x3c\xe2\xa5\xac\xa0\xfc\xd0\x4d\x10\x96\x10\x96\x29\x7a\x09\x68\xad\xb1\x55\xf1\xb5\xa0\x64\x37\x74\x9d\xb7\xbd\xf0\xbd\xcd\x74\xcf\xd6\x7c\x51\x12\xe9\xd0\xee\xb7\x3f\xf6\x60\xce\x68\xb9\x57\x46\x3b\x0a\xed\x1d\x51\x0e\xd4\x80\x3b\xf0\x27\x84\x5e\x4b\xb4\xad\xd2\x08\xb2\x4e\x34\xf7\xdb\x04\xfb\x54\x00\xc8\x3a\xd3\x2e\x80\x24\xb0\x4a\xb8\xfc\xcf\x3e\xc5\xff\x44\xec\x11\x2f\xa4\x4c\x6f\x35\x31\xd3\x78\x01\xa5\x9d\xe7\x5a\x20\x98\x06\xf6\x5b\x56\x34\xbd\x16\x74\xac\xec\xe6\x10\x4b\x80\xce\xe4\x4c\x15\x7c\xd8\x6f\x29\x39\x24\x3c\xf8\x29\xaf\xe9\x27\xeb\x35\x74\x66\x39\xae\x13\xd4\x1a\xba\xf8\xed\x6b\xe2\xf3\xa7\xf2\xa7\x07\x2d\x71\x00\x7e\x3e\xb7\x0a\x63\xc1\x49\x21\x09\x2a\x6c\xb9\x56\x09\x04\x6f\x6e\xf7\x84\x69\x5b\x14\xa4\x1b\x08\xa3\x1b\x75\xec\xa3\x8a\xa9\x84\xb2\x93\x35\x71\xac\xae\x39\x4a\xe1\x07\x3a\xeb\x71\xf0\x6c\x17\xff\x97\x84\x43\x5a\x2a\x7d\x5c\xc6\x7c\xe8\x80\x31\x46\xcd\x0d\x51\x55\x6c\x2d\xc4\xe2\x3a\x58\x6f\x46\x61\x1e\xf1\xf2\xbb\xe5\x02\xcb\xc5\x7e\xcb\xc6\x34\x8b\x2a\x9c\x94\xd8\xa0\x85\x4e\xd6\x2c\x1f\x3f\x74\xca\x97\x79\xf1\xa0\x1b\x33\x8f\x5b\xe6\x4d\xfa\x16\x91\x3b\x76\xd0\xb2\xac\xaa\xaa\x08\xa0\xaa\x81\x16\x75\x99\x78\x56\xb0\xd9\xc0\xcf\x70\x95\x3d\xb5\x42\xab\x36\xa9\x1c\x99\x8c\x23\xeb\x26\xd3\x1a\x0a\x91\x35\x93\x35\x0b\xf3\x9d\xf1\x69\xeb\xc7\x0d\x61\x4c\x80\x5f\x2c\xe3\x40\xda\x34\xe5\xe2\x23\x57\x2d\x4a\x6a\x92\xb0\xc8\x3d\xe6\x54\xe1\x2e\x26\xba\xb3\x02\xcb\xc5\xb5\x85\x8b\xd0\x87\xf9\x7e\x10\x7e\x11\xd8\xc6\x44\x41\x88\x59\xb1\x68\xed\x4d\xb1\x41\xf6\x94\x9d\xed\x5a\xe3\xb0\x4c\xe2\x4d\x26\x66\xbd\x19\x55\x61\xbb\x92\x52\xc7\x23\xc4\xf6\xaf\x34\x07\x74\xc8\x72\x7d\xc4\x71\x2c\x9e\x8a\x31\x79\x12\x6a\xbd\x99\xa0\xb2\x83\x76\xbd\xc5\x38\x6b\x21\xa6\xfa\xe5\x79\x39\xbf\x51\x52\x0c\xb8\xa3\xa4\xea\x99\x69\x79\x5d\xcc\x18\x95\x6a\xbb\x4f\xe9\x1b\xbe\x33\xb9\x27\x92\xff\x6b\x35\x71\xce\x3f\xf7\x42\x20\xc6\x2b\x1d\xcb\x21\x97\x9b\x34\xe4\xff\xa8\xa9\xca\x73\x50\xdc\xc1\xe8\xa3\xd2\xca\x9d\x50\x02\x97\x92\xb8\x24\x02\x2c\xc4\xbe\x90\x3d\x29\x32\xb9\x65\xd1\xc5\x0e\x03\x8a\x6c\x4c\xe4\xab\x64\x40\xa1\xaa\x68\xf8\xbc\x6d\xcd\x25\x3a\x1b\x0e\x28\xfa\xb0\x65\x1a\xe0\x20\x7a\xe7\x4d\x77\x3d\xcf\x8f\x9c\xfc\x38\x1c\x9d\xcc\xd5\xdc\xd1\x28\xdf\xeb\x66\xd6\x0c\x01\x98\xa6\x3b\xbe\x55\xbb\x11\x31\x79\xda\xbd\xd6\x46\xf9\xbe\xd1\xd5\x62\xc8\x5d\x86\x96\x8b\x78\x70\x87\xe1\xac\x2c\x4a\x2a\xad\x9a\xdc\x96\x74\xd1\x9a\xce\x8f\x77\x23\x15\x0c\x27\xee\xe8\xf9\xa7\xa8\x45\xf5\x5f\xdc\x2a\x36\x04\xaf\x8f\xf0\xdb\x98\xd4\x3b\x3a\xf2\x7b\xbb\xf0\x7b\x77\xbc\x09\x17\x7f\x99\x49\xdc\x02\xbe\x81\x34\xd7\x97\xa3\x19\xca\xd9\x23\xf4\xd2\x7b\xf1\xdd\x0f\xf4\x2b\x86\xf0\x6b\x26\x93\xe9\xdd\x21\xf5\x73\x66\xfb\x4f\x00\x00\x00\xff\xff\x5b\x42\x5e\xea\x0f\x0c\x00\x00"),
			path: "mongo-solo.tml",
			root: "mongo-solo.tml",
		},
	}
)

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
	return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) (io.Reader, int64) {
	reader, size, err := FindFile(path, doGzip)
	if err != nil {
		panic(err)
	}

	return reader, size
}

// FindDecompressedGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error) {
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64) {
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error) {
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64) {
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error) {
	reader, size, err := FindFileReader(path)
	if err != nil {
		return nil, size, err
	}

	if !doGzip {
		return reader, size, nil
	}

	gzr, err := gzip.NewReader(reader)
	return gzr, size, err
}

// MustFindFileReader returns bytes.Reader for path else panics.
func MustFindFileReader(path string) (*bytes.Reader, int64) {
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error) {
	item, ok := assetFiles[path]
	if !ok {
		return nil, 0, fmt.Errorf("File %q not found in file system", path)
	}

	return bytes.NewReader(item.data), int64(len(item.data)), nil
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
	body, err := ReadFile(path, doGzip)
	if err != nil {
		panic(err)
	}

	return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error) {
	body, err := ReadFileByte(path, doGzip)
	return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
	body, err := ReadFileByte(path, doGzip)
	if err != nil {
		panic(err)
	}

	return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error) {
	reader, _, err := FindFile(path, doGzip)
	if err != nil {
		return nil, err
	}

	if closer, ok := reader.(io.Closer); ok {
		defer closer.Close()
	}

	var bu bytes.Buffer

	_, err = io.Copy(&bu, reader)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
	}

	return bu.Bytes(), nil
}
