package static


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct{
  path string
  root string
  data []byte
}

var (
  assets = map[string][]string{
    
      ".tml": []string{  // all .tml assets.
        
          "mongo-api-json.tml",
        
          "mongo-api-readme.tml",
        
          "mongo-api-test.tml",
        
          "mongo-api.tml",
        
          "mongo-solo-readme.tml",
        
          "mongo-solo.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        "mongo-api-json.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\xc1\x6b\xc2\x30\x14\x87\xcf\x0d\xe4\x7f\xf8\xcd\xc3\xb0\x20\xf5\xbe\xe1\x61\x0c\x76\xd8\x41\x07\xe2\x69\x0c\x8c\xf1\x55\xeb\xda\x44\x5e\xa3\x32\x42\xfe\xf7\x91\x54\x99\x87\x6d\xac\xec\xd2\xd2\xbe\xc7\xf7\x7d\xbc\xa3\x62\x78\x5f\xdb\x13\x31\x8a\xb9\xe3\x83\x76\xc5\x6c\xb5\x23\xed\x8a\xa9\x6a\x28\x3d\x42\x78\x9e\xcf\xa6\x98\x60\xe9\x3d\x1a\xb5\x4f\x5f\xe7\x65\x0c\x56\xad\x35\x03\x0c\x76\xe9\x15\xc2\x52\x0a\x29\xfe\x46\x7d\x64\x52\x8e\xbe\x63\x77\x93\x07\xed\x2a\x6b\xfe\x63\x58\xec\xd7\x3f\x18\xba\xc9\x6f\x86\xf1\x18\xb5\x55\xeb\xb8\xfe\x64\x19\x4c\xee\xc0\xa6\x85\x82\xa1\x13\x2a\xd3\x3a\x65\x34\xc1\x96\x50\xf0\xfe\x92\xf0\xa2\xf4\xbb\xda\x50\x08\xc5\xd7\xbf\xab\xac\x10\x50\xb2\x6d\xe0\xb6\x84\x3d\xdb\x63\xb5\x26\x44\x29\xb4\x35\x8e\x8c\x2b\xa4\x28\x0f\x46\x5f\x8b\x87\xe7\x11\x5a\xc7\x95\xd9\xe4\x18\xf6\xb0\x8d\x40\xcc\x96\x73\x78\x29\xb2\x78\x32\xaa\xa9\xe9\x93\x1b\x0f\x91\x55\x65\xc4\xe0\x6e\x92\x5a\x8b\x85\x69\x14\xb7\x5b\x55\x0f\x5f\xdf\x56\x1f\x8e\x2e\x85\xf9\x08\xb7\x91\x9f\xdf\xa7\xf5\x9b\x09\x4c\x55\x27\x73\xd6\x1d\xaf\x8f\xd8\x77\xed\x52\x64\x5d\xc3\x99\x10\xf9\xa3\xc8\x95\x22\x48\xf1\x19\x00\x00\xff\xff\xfe\xdd\xd6\xbe\xbd\x02\x00\x00"),
          path: "mongo-api-json.tml",
          root: "mongo-api-json.tml",
        },
      
        "mongo-api-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xdf\x8a\x9c\x30\x14\xc6\xef\x05\xdf\xe1\x80\x37\x4e\xd9\xe6\x01\x0a\xbd\x98\x6a\x59\xf6\xa2\xad\xb4\xdd\xab\x52\x30\x9b\x39\xc6\x74\x63\x8e\x24\x47\x2a\x2c\xbe\x7b\x89\xce\x30\x7f\x70\xe8\x94\xae\x37\xd1\x83\xf9\xbe\x5f\x7e\x79\x79\x11\xdf\xd8\x0f\x8a\xc5\x97\xa7\x5f\xa8\x58\x7c\x96\x1d\x4e\x13\x7c\x22\xa7\xa9\xfc\x00\xdb\xea\x21\x4d\xde\xff\xfd\x49\x93\x34\xb9\x21\x0b\x4c\x00\x09\x72\x60\x7a\xab\xd1\xa1\x97\x8c\x3b\x28\xbe\x3e\x96\x60\xba\xde\x62\x87\x8e\x25\x1b\x72\xd0\x90\x07\x6e\x11\xea\xd5\xd0\x7d\x72\x0d\xc6\x41\x2f\xd5\xb3\xd4\xcb\x9f\xd5\xb3\x16\x95\xe4\x76\x9a\x6a\x11\x89\xbe\xb7\x08\x0d\x59\x4b\xbf\x8d\xd3\xd0\x21\xb7\xb4\x03\x1c\x4d\xe0\x30\x37\xa8\x21\x30\x75\x40\x7d\x24\x31\xe4\xc2\xbb\xb8\x2b\xcb\xe0\xe3\x88\x2a\xbe\xd6\x75\xad\x29\x4d\xe2\x67\xae\x78\x04\x45\x8e\x71\x64\x51\x2c\xeb\x1d\x34\x23\x34\x83\x53\xb9\x22\x0b\x6f\x3a\x4d\xa2\x20\x6b\x51\xc5\xb0\x0d\xa0\xf7\xe4\xf7\xcb\x9c\x75\x8d\x29\x1c\xa0\x8c\x9b\x4f\x7d\x74\x13\x9d\xc9\x00\x3d\x7a\x96\xc6\xc5\x1d\x4c\xb3\xb0\x03\x69\xe1\x51\x32\x9e\xb0\x2e\x83\x75\x5a\xb4\xd8\xc1\xd1\x68\xb5\xa8\x9b\x26\x71\xe5\xea\x2e\xd9\xb3\x0c\xee\x91\x4f\xca\xee\x91\xd7\x9b\xfa\xe1\xc9\x1a\xf5\x50\x42\x60\x6f\x9c\xde\x40\xfe\x0f\xb5\x77\x7b\x73\x17\xbd\xb0\xb5\xf6\xbc\x7b\x6b\xed\x5a\xfd\x06\xf2\x1f\x3f\xff\xb3\xef\xb1\xdf\x9d\x7b\x5d\x06\x37\x9d\xf6\x55\x44\x97\x68\xf1\x0c\x60\x19\xdc\xa8\xfb\x18\xf7\x27\x00\x00\xff\xff\xfb\x03\x21\x07\xe2\x03\x00\x00"),
          path: "mongo-api-readme.tml",
          root: "mongo-api-readme.tml",
        },
      
        "mongo-api-test.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x5b\x6f\xe2\x46\x14\xc7\xdf\x23\xe5\x3b\x9c\x5a\xaa\x64\xb6\xee\x64\xf7\x95\x8a\x07\x08\xdd\xa4\x95\x82\xd1\x9a\x74\x1f\xd1\x30\x73\x0c\xd3\x1d\x7b\xd2\x99\x63\x48\x84\xf8\xee\x95\x2f\x5c\xb2\x09\x59\x43\x36\xad\x4a\x1d\x29\x91\x62\xe6\xdc\xfe\xf3\x3f\x3f\xc9\xcc\xb9\x05\xff\xfc\x0c\x00\x00\xe7\x98\x92\x83\x0e\x24\x48\x56\x09\xc7\x06\xb8\xf0\x45\xe6\xc8\x24\x2c\x22\x2e\xbe\xf4\x95\xbb\xd3\xfc\xc1\x37\x8e\x45\x24\x4d\x46\xad\xd6\xf9\x59\x19\x2b\x4c\x1a\xab\x69\x1e\x6b\xd2\xa9\x61\x97\xc5\xbf\xcb\xf2\xb3\xfc\xe7\xc6\x48\x6c\x43\x32\x35\xec\xc6\xa4\x86\x4c\xaa\x44\xb0\xfd\xb4\xdf\x6b\x83\x71\xec\x0a\x09\xd3\xb9\xef\x2d\x97\x2c\x22\x9b\x09\x62\x43\x2e\xbe\xf0\x29\xae\x56\xe3\x9b\x70\x70\x15\x8e\xfb\x3d\xaf\xb5\x13\x77\x6d\x1c\xd5\x8c\xbc\x0e\xa3\xd1\xa3\xd8\x5b\x87\xb6\x66\xec\x6d\xf4\xeb\xa7\x47\xb1\xdd\x8c\x66\xb5\x7b\xee\xde\x8e\xae\xbf\xea\x7b\xc8\x9d\x5b\x18\x2b\x6b\x66\x18\x76\xa3\xe8\x73\xf8\xa9\xbf\xc9\xb1\x5a\xcb\x4e\xe8\xe8\xd2\x68\xe8\x80\xb7\x5c\x6a\xb3\x40\x0b\xeb\x24\xe1\xe4\x4f\x14\xc4\x06\x3c\xc1\xe2\xcf\x6a\x35\xce\x4f\x8f\x85\xd1\x1a\x05\x29\x93\x7a\xe7\x67\xc5\xfd\x5d\x5c\xc0\x08\x1d\x5d\x21\x6d\x5b\xd8\x89\x5e\xad\x60\xce\xb5\x92\x9c\xd0\x01\xcd\x10\x6c\x6e\x0f\x9c\x73\x0d\x26\x06\x0e\x7b\x82\x8a\xbc\x16\x85\xb1\x12\x62\x6b\x12\xe0\xa5\x37\xe4\x84\x9d\x9f\xc5\x59\x2a\xbe\x51\xd4\x27\x78\x97\x37\xac\xd2\x29\x1b\xb5\x2a\x2b\xf1\x3b\x05\xed\x0e\x24\x72\x52\x98\xb3\x1a\x3f\xa8\xac\x1b\x54\xee\x2b\x7c\x5b\x38\x70\xc7\xa0\x74\x9f\x47\x0a\x93\x12\xde\x13\xfb\xac\x68\x36\x52\x09\x9a\x8c\xfc\xf5\xb3\x01\x2e\xfe\xe0\x3a\xc3\x1e\x9f\xfa\xad\x00\x3e\xbc\x87\x77\x40\x2a\x41\x16\xa1\x30\xa9\xdc\xa4\x42\x8d\x49\x00\x68\x6d\x9e\x50\x1b\x2e\x7f\x8f\xc2\xc1\x47\x63\xfd\x6f\x5f\xc1\xa5\x45\x4e\x98\x9f\x6f\x95\xb9\x54\x5c\x24\xfa\xa1\x03\xa9\xd2\xb0\xb3\x30\xf9\x68\x8e\x7d\xe4\x4a\xa3\xf4\xbd\x28\x13\x02\x9d\x8b\x33\xad\x1f\x8a\x92\x28\x21\xcf\x02\xb1\xb1\xfb\x2e\xa0\x12\xbf\x0d\x3f\xfe\xf4\x17\xf3\x8a\x86\x5b\x1b\xfb\x6c\x4b\xe4\x5e\x7c\x65\x09\x6f\x23\x8d\xc4\x18\x6d\x7e\x4b\xac\x8f\x1a\x09\x7d\x41\xf7\x41\x21\x18\x1b\x66\x13\xad\xc4\x6f\xfd\xcd\xd9\x6a\xf4\x76\xa7\x38\x5f\x2a\xb3\x3d\xdf\xfa\xe5\x70\x61\xb8\xcc\x9b\x5e\x7b\xee\x85\xb6\x55\x4a\x06\xe4\xe4\x28\x69\x0e\x2d\xc2\xb6\xea\x8c\x4b\xd7\x94\x03\x5f\x21\x3d\xaf\xce\xb1\xb6\xa8\xd6\x12\x25\x38\x32\xb6\x5e\x8f\xc5\x62\x1e\x29\xc4\x2b\xea\x15\x9a\xac\x1e\xb3\xa7\xab\xf5\x31\xf8\xd1\xfa\xd5\x00\xda\x5f\xb9\x61\x50\xc3\xa0\x13\x63\x50\x19\xe4\x82\x35\x8c\xda\x1b\x1a\x75\xb5\x1e\xa2\x1d\xf2\x69\xa5\x80\xc7\x9d\xf0\x02\xf0\xee\x0a\xd1\xc6\x4a\x7a\x01\xfc\xfc\x21\xff\xfd\x2e\x90\xca\x37\xb7\x6a\xe6\x1f\x40\xd4\x81\xd5\x76\x04\x53\x31\x68\x4c\xfd\x2a\xb8\x05\x9d\x0e\xbc\x3f\x7c\x58\xd2\xc8\x1d\xc1\x87\x43\x21\x79\xf0\x9c\xc7\x16\xaa\x4d\xe3\xd0\x4a\xb4\xbd\x87\x7f\x0f\xca\xbd\x87\xa2\x85\x86\xcd\x0d\x9b\x4f\x95\xcd\x4f\xc0\xbc\xf6\xfc\x1e\x30\x37\x40\x3e\x6d\x20\xef\x39\x5f\xee\xc4\x57\x24\x16\xf9\x43\x65\xd2\xda\x2f\xe7\x0b\x45\xb3\xe7\x31\xfc\x62\xd9\x86\xbf\x0d\x7f\x4f\x86\xbf\x35\x96\xed\xf6\x4e\x3e\x5d\xb6\xac\x7c\xf8\x66\xab\x56\x16\x6d\x56\xad\x59\xb5\x93\x59\xb5\x32\xeb\xc5\xc5\x28\xec\x87\x6d\xa8\xb6\xca\x99\x04\x69\x96\xfb\xfb\x79\x49\xaa\x3d\x78\x2a\xe1\xf1\x0a\x95\xab\xfb\xe6\x1a\x1d\x5e\xa6\x2e\x90\x4a\x5b\x3d\x79\x0f\x4b\xcc\xf7\xf9\x66\xfe\xc5\xaa\x0d\x91\xfe\x83\x44\xfa\xff\x51\xe6\xf1\xc4\xfb\x39\x7c\xc4\xe8\xc5\xa2\xbd\xfd\xf0\x87\x97\x59\xe3\xe3\xef\x00\x00\x00\xff\xff\x54\xd7\x6c\x96\xd9\x1d\x00\x00"),
          path: "mongo-api-test.tml",
          root: "mongo-api-test.tml",
        },
      
        "mongo-api.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x5d\x6f\xdb\xc6\xd2\xbe\xb6\x01\xff\x87\x29\xf1\x22\x20\x5b\x95\x6e\xdf\x4b\xb5\x2e\x50\xdb\xc9\x41\x80\x36\xc9\x49\xd2\xd3\x8b\xa2\x08\x56\xdc\x91\xbc\x35\xb9\xd4\xe1\xae\x22\x0b\x82\xfe\xfb\xc1\x7e\xf1\x4b\x22\x45\xc9\xb4\x23\x3b\xea\x45\x1d\x52\xfb\x31\xf3\xcc\xec\xcc\x33\x4b\x2e\xcf\xcf\x61\xb9\x0c\x3f\xc8\x6c\x16\xc9\xf0\xed\xe8\x1f\x8c\x64\xf8\x86\x24\xb8\x5a\xbd\x62\x18\x53\x01\x14\xc7\x8c\xa3\x00\xc2\x81\x71\x89\xd9\x98\x44\x08\xf3\x1b\x16\xdd\x00\xde\x4d\x53\x81\x02\x12\x94\x37\x29\x05\x99\x42\x86\x72\x96\x71\x20\x90\x90\x29\xa4\x63\x20\x71\x7c\x76\x7a\x7e\x0e\x44\xca\x8c\x8d\x66\x52\x8d\x23\x44\x1a\x31\x22\x91\xc2\x9c\xc9\x1b\x90\x37\x68\xe7\xa0\x20\xb4\x18\xb3\x0c\x81\xa8\x89\x23\x46\x91\xc2\x68\xa1\xdb\xe4\xbf\x85\x67\xa7\x72\x31\xc5\x2d\x62\x97\x84\x5d\x9e\x9d\x9e\x98\xbb\x7e\x00\x7e\x42\xa6\x7f\x09\x99\x31\x3e\xf9\x3b\x6f\xb2\x5c\x0d\x00\xb3\x2c\xcd\x82\xb3\xd3\xd5\xd9\xa9\x16\xb9\x61\xf8\xab\x94\x8b\x59\x82\x59\x1b\x2e\x24\x8a\x70\x2a\x45\x01\x03\x25\x92\xd8\xdf\xe6\x2c\x8e\x61\x84\x10\x99\x71\xa8\x9e\x8b\x71\x99\x6a\x25\x27\xec\x33\xe3\x13\x60\xc9\x34\xc6\x04\xb9\x54\x17\x7d\x80\x92\x4b\x5d\x45\xc5\xde\x6e\xc0\x24\x30\x98\x14\x90\xfc\x9e\xf2\x49\x4a\x0b\xcd\x1b\x1d\x82\x38\x97\x18\xa7\x99\xf2\x89\x8c\xa1\xd6\x8b\xe8\x61\x12\x35\x4c\x78\x4d\x24\x19\x11\x81\x40\x38\xb5\xb7\x3e\xa0\x10\x2c\xe5\x4e\x17\x3b\x5d\x55\xe4\x37\x38\x57\x56\xfc\x36\x29\x0d\x31\x00\x7d\x69\xbb\x77\x36\xe5\xf5\x65\x49\x95\x02\x65\xa3\xca\x34\x4b\x3f\x33\x8a\x70\x7d\x09\x57\xef\xff\xb8\x86\x74\x8a\x19\x91\x2c\xe5\x42\x0f\x39\x13\x4a\x1f\x2d\xb6\x32\x8b\xb2\xc5\x8c\x53\xcc\x62\xc6\x11\xe8\x68\x8b\x39\xae\x2f\xed\x7c\xcb\xb3\x53\x80\x28\x8d\xc1\xa0\xaf\xae\xe8\xc8\x2a\xae\x2e\x12\x85\x5d\x24\xdc\xdf\xf0\x77\xf3\x57\xfd\x84\x5c\xcc\x32\xa4\xaf\x39\xc5\x3b\x18\xa5\x69\xac\x6e\x32\x1e\xa5\xca\x77\x24\xd6\xef\x53\xbc\x43\x01\x7f\xfd\xad\x80\xd2\xbf\x15\xf0\xbc\xc1\xb9\x5d\xb8\x0a\x08\x8e\x73\x60\x5c\x48\xc2\x23\x54\xbe\xdb\xa8\x42\x78\x76\x3a\x9e\xf1\x48\x75\xf7\x0b\x15\x06\x90\xd4\xa5\x1d\x00\x24\xa9\x55\x6a\x90\x8b\x12\x86\x61\x2e\x4b\x00\xdf\x36\x4e\xa3\x21\x02\x17\x59\x5e\x6c\x6b\xa7\xfe\xa3\xa3\x21\x24\xe9\xa0\xb8\x11\xa5\xf1\x50\xfd\xaf\x74\xcb\xca\x38\x84\xa4\x74\xd3\xca\x36\x74\xff\xb0\x3f\xad\x0a\xac\x0c\xea\x06\x5c\x22\x25\x26\x6a\x9d\xcb\xd4\xde\x57\xf1\xce\x79\x0e\xcd\x35\xcd\x57\xb7\x98\x62\xc4\xc6\x2c\x52\xa2\xc4\x18\x49\xe3\xeb\x1a\x44\x3f\xa1\xa3\x16\x10\x82\xf2\xc4\xbe\x5d\x98\x60\x55\x66\x63\x48\xe8\x28\xac\x38\x44\x09\x0d\x8b\x1c\x67\x71\xae\x8d\xfa\x63\xfe\x9f\xc0\xf0\x22\xb7\xd7\x1b\x9c\x7f\xcc\x48\x84\xbe\xd7\x6c\xf4\x92\x1c\x5e\x60\xc6\xa0\x38\xc6\x4c\x8b\xe0\x06\x7a\x99\x30\xe9\xbb\x8b\xd7\x7c\x9c\x76\x1e\x71\xe0\x7a\xfd\xc9\xe4\x8d\x11\x26\x09\x5f\x72\xea\x07\x41\xa0\x4c\x60\xf5\x8d\x91\x2b\xc4\x42\x0b\x71\x00\x17\x17\xf0\xc3\x16\xa5\xcd\x3f\xce\xcf\xe1\xf5\x18\xe6\x08\x37\x84\xaa\xd8\x6d\x90\x1c\xe1\x38\xcd\xd0\x58\x0c\xe6\x44\x80\x5b\x46\x03\x65\x38\x0e\xe2\x96\x4d\x07\xaa\x57\x44\xb8\x04\x9e\xca\x7c\x30\x21\xd3\xa9\x36\x7b\x3a\x15\x30\xc2\x88\xcc\x84\x5e\x36\x63\xc2\x62\xe7\x03\x61\x2e\xf7\x37\x6b\x86\x7a\xf1\x02\x8c\x22\xd5\x85\xdb\x45\x15\x9a\x87\x3f\x51\x8a\x7c\xda\xa0\x74\x14\xd2\x51\xa8\x63\x65\x3e\xb7\xfa\xed\x9b\x0b\x35\x4e\x79\xf4\x46\xb3\xbd\x54\xc0\x8c\x7d\xef\x95\x51\x44\xa6\x10\x65\x48\x24\xba\xc9\x74\x6c\x67\x1b\xac\xe6\x7b\x85\x7b\x7b\x03\x3d\x41\x94\xc6\xf5\x36\x1a\x77\x4f\x4b\x6c\xa6\x32\x16\xae\x29\x8d\x59\x56\x57\x5a\x3b\x9b\x95\x21\xbc\x8a\x53\x81\x7e\xee\x19\xc5\xc4\x0a\x05\x87\x4f\x78\xe5\x3b\x21\x5c\x43\x25\xfb\x27\x1b\x8b\x54\xd3\x8c\xf0\x09\x42\xc9\xa3\xca\x10\x59\xec\x86\x17\xe5\x75\xfb\xb2\xb4\x1e\x75\x9f\xe0\xa7\x06\x84\x77\x44\xd9\x46\x12\x87\xf2\xfe\x08\x9b\x9e\x56\xc9\x8e\xf0\xaf\x4b\x5d\x77\xcc\x0b\x90\xd9\x0c\xab\xed\xea\xc6\xaa\x18\xac\x55\x7d\x13\x1b\x3e\xcc\xa2\x08\xd1\x84\x4c\xa3\xbf\x4a\xae\x25\x63\xf6\x05\x42\x50\x73\xa6\xb5\xd5\xe8\xb4\x2b\x7e\x6e\x11\xfb\x15\xe3\x4c\xdc\x20\x05\x42\xa9\x66\x6d\xdd\xa5\x0c\x2a\x49\x4d\x2f\x6d\x97\x61\xae\xd2\x19\x97\x95\xdc\x62\x5b\xa9\x0c\x22\x53\x49\x62\xe0\xb3\x64\x84\x99\x8a\x32\x19\x46\x69\x46\x61\x9c\xa5\x89\x21\xd2\xa3\xce\x09\x45\xcf\xe3\x47\xf2\x4e\x51\x51\x89\x77\x32\xbc\x32\x7f\x03\xf0\x19\x97\x8e\x45\x39\x4f\xde\x39\x51\xe8\xf1\x7b\x4a\x11\x76\xac\x8e\xc9\xc1\xe9\xf3\x5a\xbc\xbc\x9b\xb2\x0c\xa9\xd2\x32\x28\x2f\x49\xbb\x9c\xc7\x89\xcc\x17\xa0\x55\x1e\x6e\x88\x50\x34\x56\x75\xf3\x82\x4e\x3e\xbc\xbe\x84\x27\x28\x9d\x65\xa2\x0d\x92\xf7\x19\x20\xbf\xff\x71\xb0\x21\x48\x16\x11\xab\xf0\x70\x4b\x1d\x9a\xa2\xd4\x0e\xea\x91\xe9\x34\x5e\xf4\x1f\xfa\x3b\xea\xf6\xd8\x59\xef\x21\x8d\xd9\x55\xe5\xb6\x9c\xf7\xdf\x19\x66\x0b\xa5\xfe\x48\xa4\x3c\xfc\x7d\x99\xf7\xd2\xb1\x22\x07\x67\x43\x32\x0c\x5f\x31\x4e\x7d\xdd\x3f\x30\x4b\xcc\x0f\x7e\x3a\x30\xe0\xb4\x74\xde\xc0\x68\xd9\x2f\xaa\x5b\xa2\xd1\x35\xaa\xa4\x47\xad\x0a\x3d\x08\x9f\x4b\xe6\xe2\x79\x6e\x9f\x22\xf8\x9b\x49\x6b\xd1\x3f\x49\x3f\xa3\x8e\xee\xeb\xd1\xde\xd6\xa1\xea\x22\x2f\x3b\xa6\xb3\x51\xcc\xa2\xd7\xd7\xa1\x1e\xf1\xbd\xee\x23\xf2\x86\x4c\xa8\x92\x36\x99\x09\x15\xea\x3e\x23\x10\xdb\x1e\x18\x85\xcf\x24\x9e\xe1\x40\x85\xbf\x0c\x85\x40\x0a\xc8\xe4\x0d\x66\x30\x5a\x00\xd1\xee\x05\x69\x06\xff\xa8\xbf\x92\x4c\xf4\xe8\x29\x77\xdb\x16\xc8\x4b\x95\xe2\x3b\x12\xdd\x92\x09\xae\x56\x61\x43\x4c\xb7\xe5\x6f\xe7\x64\x65\x70\xd9\x94\xad\x06\xb9\xbe\xb6\x00\xad\x15\x47\x3b\xe7\x2d\x33\x55\x4f\x89\xcb\x0d\x56\x73\x0c\x27\xb2\x57\x48\xff\x25\x92\xdb\x1e\x8b\x9a\x1a\x07\x6d\x58\x14\x5b\xf5\x7a\xd8\xda\xe0\x49\xa6\xbd\xad\x5a\x3d\x76\xc2\x3b\x68\x13\xef\x94\x0a\x8b\xf1\x0a\xb1\x87\xb9\xd8\x83\x46\xef\xd9\x94\x2b\xdf\xeb\x30\x6c\xb3\x65\x0f\xde\xd4\x8e\xf2\xbd\x53\x63\x07\x33\xed\x6b\x82\xc7\xce\x9b\x1d\x34\xab\xa7\xd6\x6a\x41\x65\xb6\x2d\xca\x39\x95\x50\x5a\x4e\xa8\xf9\x06\xdd\xe6\x84\x5a\xde\x0d\x95\x37\x58\xdb\x52\xde\x9a\xeb\xbe\x60\x1e\xbe\x57\xce\x35\xb8\x6d\xce\xb9\x18\x63\xb2\x0b\x06\xf7\x4d\xca\x46\x96\xbe\xaa\x49\x3b\x58\xb3\x63\x29\xf5\xc2\x77\x4f\x2b\x33\xdb\xdd\xb9\xed\x61\xbb\x45\xb9\x63\x7a\x3e\xfc\xf4\x5c\xdd\x85\x3d\x48\x43\xb7\x26\xe9\xe5\x52\xc1\xe0\x2b\x8f\x7f\xa5\xe2\x90\x5d\xa5\xe0\x99\x47\xb4\x1e\x40\x00\xab\x55\x31\xcb\x58\xdf\xce\x41\xd5\x97\x98\x85\xee\x81\xee\xda\x36\x6d\xc7\x0d\xd8\xea\xcf\x50\x3c\x11\x6a\xa8\x64\xf3\x07\xc8\x63\x15\xc6\x1a\x22\x6c\xbe\xf8\x9a\x47\x6f\x43\x7f\x5b\x2f\x65\x50\x6b\xd7\x0e\x8d\x37\x18\xaf\xd6\x29\xd8\x71\x1b\xb7\x9d\x21\xbd\xe6\x02\x33\xe9\x1b\xee\xe5\x1b\xb3\x05\x7d\xed\x8a\x5b\xaf\xdf\x0a\xfc\x1e\x4e\x5e\x06\x75\x27\xff\xef\x84\xd9\x96\x2c\x75\xd5\x1a\xb5\x77\x95\x7f\xd3\xe6\x47\xfe\xdc\x68\xb9\xc4\x58\x60\x79\x69\xd5\xf8\xb2\xbf\x5c\xea\x17\x16\x6c\x82\xfc\xd5\xec\xbd\x9b\x09\xc0\x53\x6d\x3c\xf0\xfe\xd1\x7f\x56\xab\x60\x67\xc7\x68\xa7\xce\x07\xe3\x0f\xfb\xec\x38\x1d\x96\x47\xac\x6d\x3b\x59\xeb\x73\xba\x5a\xb5\x11\xe5\x7f\xa1\xfc\x35\x8e\xdf\x61\xf6\x8e\x4c\xd0\xbd\xb6\x81\x42\x3f\xdd\xcc\x2c\x7d\x2d\xef\x3d\x11\x4e\x4b\xef\x0d\x88\x98\xd5\x5f\x18\xd8\xbe\x0d\x24\x17\x53\x7c\xaa\xfc\xb8\x02\xd7\x66\x9a\x9c\x66\x54\x25\x42\xfb\x62\x84\xbe\xba\x5c\xe4\xd7\x53\x05\xb3\x7e\xd4\x92\xa1\x98\xa6\x5c\xa0\xc3\x9e\x71\x19\x00\xf8\x7f\xfd\xbd\x03\x96\x03\x80\x3e\x1e\xdb\x18\xad\x7a\x62\xda\x6e\xb0\x6d\x0c\xfa\x44\xcc\x99\x8c\x6e\x2c\x32\x22\xfc\x98\xfe\x96\xce\x31\xf3\x35\x62\x5a\x97\x93\x88\x08\x04\x8f\x8a\xc8\x1b\x80\x47\x51\x44\xde\xb0\x58\x52\x0e\xd9\x0b\xf0\xbe\xf7\xe0\x3b\x77\x7d\x76\x7a\xb2\x3a\x1c\x82\xee\x16\xd4\xbd\x82\x7d\x27\x46\xc6\x59\x3c\x78\x4e\x4f\x86\xba\xa9\x77\xc2\xc6\x66\x49\xfd\x7c\x01\x3f\xc0\x8b\x17\x6b\xab\xea\x67\xfb\x8e\xc8\xc9\x89\x8d\x66\x15\xa6\x6e\x7c\xf5\x72\xf1\x56\xf9\x8e\xf2\x0c\xbb\x60\xf3\x75\x1b\x98\x9e\x5a\x84\x7c\x80\x18\xb9\x6f\x2f\x02\x2b\x90\xf1\xb9\x13\x13\x50\x1b\x9e\xd8\x8a\xf0\xec\xf4\x44\xff\xf4\x7e\x83\x28\xf9\xa3\xd9\xc0\x68\x55\xb5\x4b\x2e\x44\x15\x07\x3b\xed\x67\x92\x99\x39\xff\x24\x5c\xa2\x7b\xdf\xea\x63\xfa\x41\x92\x4c\xaa\x08\x51\x87\xea\xc7\x4d\x50\xfd\xe2\x90\x2a\x0d\x05\x17\xf5\x66\xaa\x41\x65\xf8\x0b\xf8\x41\x09\x02\x8a\x68\x74\xe8\x0f\xdf\x6a\x29\x36\x0c\x53\xee\x76\x0e\xff\xaf\x65\xce\x85\xfe\x05\x7e\x34\x83\x57\x7a\x7d\xf7\x9d\xba\xb5\xca\x81\x68\x61\xfc\xb5\x0d\xab\xcb\xe1\xbf\x55\xa6\x1c\x56\x02\xba\x17\x0c\x60\xbd\x87\x72\x53\x5b\x09\xb8\x5b\xfa\xb2\x46\x64\x3c\xa1\x24\x62\x7c\xf2\xc9\xac\x85\xa1\xe3\x49\x25\x79\x6b\x5c\xdc\xd3\x2a\x7f\xb2\xee\xf1\x69\xae\x75\xf7\x86\x15\x5b\xd6\x7a\x68\xbf\xcc\xc7\x86\x4a\x30\xdc\xd8\xf6\x72\x51\x6f\x6d\x6f\xd7\x5b\x2b\x98\xd7\x07\xd6\xe8\xd7\x9b\xd6\x4c\xea\x7a\xd5\x6e\x97\x7a\xad\x5c\x15\x12\x1c\x68\x05\xfd\x28\x71\x78\xbf\x47\xba\x8d\xb5\xb3\x7d\x91\x77\xad\x78\x56\xd1\x80\x49\x4c\xf4\x9b\x9f\x1b\x5f\xf2\xed\x4e\xe3\xcb\x4f\x8b\x3f\xdc\xb2\xa9\x5f\xf6\xe6\x20\xfc\x8d\x29\xd4\x4b\xee\x1a\x84\x1f\xd2\x4c\xfa\x2e\x7a\x86\xbf\xc6\xb1\xff\x42\x0b\xd3\x57\x11\x90\x67\xd4\x32\x43\x6d\x7e\x65\x55\xb3\x4d\xc3\x60\xe9\xe8\x0b\x3c\x8d\x6e\xf5\x0c\xa8\x97\x09\xca\x74\x99\xb3\xdd\x0e\x3c\xb0\x3c\x86\x7b\xcd\x4d\x62\x52\xbc\xe5\x66\xc6\xac\xa1\xae\x66\xdb\x75\x53\xb7\xae\x5b\xe1\x3e\x7a\xe7\xc9\xbd\x5c\xae\xe6\x6b\xb5\x78\x1d\x97\x0a\x26\x6b\xb8\xe8\xd6\x46\x87\x0b\x45\x38\x90\x53\xdf\x5c\xdb\x22\xb8\x01\x4f\x97\xb9\x6d\xd3\x52\x59\x6c\xd2\x55\xd3\xb2\xd9\x13\x90\xe3\x5a\x3a\xa0\xb5\x64\x5b\x5a\xd3\x57\x79\x57\xc5\x11\x38\xd5\x7e\x50\xad\x88\x2d\x27\x3c\x56\xc4\xdd\x2a\xe2\x12\x85\x6e\xa8\x88\xeb\xa5\xf0\x3e\xb5\x6e\x2f\x65\xae\x15\xb5\xd7\x6a\x37\x1f\xf3\xf9\x15\xbd\x4f\xa1\xec\x7d\xc6\x25\xef\xc1\x3f\x7b\x7a\x78\x7d\x9f\x38\x6b\x3e\xe6\xf0\xb6\x1c\x7e\xe4\xc2\x5b\x30\x59\xc3\x05\x9e\x1d\x17\x3e\xae\x90\xfd\x56\x48\x95\xe1\x6e\xe1\xb4\x97\x0b\xbd\x69\x55\x26\xb4\x1d\xdf\x2d\xd6\x4f\x75\xe1\x16\x17\x9a\xea\x6a\xd6\xa9\x07\x75\xa4\x57\x35\xfe\x8a\xd8\xae\x05\x72\x33\xd5\x55\x28\xb9\x47\x3d\x5a\x62\xa8\x9c\xad\x06\xff\xb1\x09\xaf\x95\xb6\x3f\xb6\x9b\x0f\x58\xf5\xeb\x5b\x5c\x58\x8d\x9f\xc8\xab\x53\x5b\x28\x6a\x93\x3e\xfd\x1e\xfd\xd8\xc1\x1b\xec\xe7\x0a\x9e\x03\xcf\xed\x1f\x86\xc3\xe2\xc4\x4f\xc4\x77\x76\x22\xd6\xb7\xb8\x18\x1a\x9d\xee\x4f\xb1\xa1\x4f\x7e\xfd\x96\xa3\xe1\x0b\x47\xba\xb0\xaf\x47\xac\x11\x8b\xfb\x92\xe1\x7d\x89\x70\xdf\x7a\xb8\xd7\x72\x62\x4c\x1a\x08\xef\x06\xe7\x7c\x70\xbe\x7b\xd8\x1e\xdb\xea\x8d\xbd\xc4\xaf\x07\xb4\xb4\x42\x75\x2b\x13\xde\x81\x02\xbb\x53\x66\xe5\x3d\xde\xaf\x8c\xee\x76\x3c\x7f\xf7\xf8\xcc\xb6\x3f\x4a\xfb\x8c\x8e\xea\x6d\xdb\x78\x7d\xd4\x93\x5c\x5f\x27\xbb\xed\x1b\x84\xc3\xe2\xb6\x4f\xcc\x83\x76\xe2\xb8\x56\xb7\x4f\x8c\x96\x4e\x0c\x1e\xe9\xee\x01\xd0\xdd\x87\xa4\x0e\x47\xb2\x7b\x24\xbb\xcf\x94\xec\xfe\x31\xa5\x2a\x94\xcf\xc4\x91\xea\x6e\xa5\xba\x06\xab\x4e\x6c\xf7\xf1\x8f\xc2\x1a\xe1\x7a\x62\xbc\x6e\xb0\x67\x43\x7a\xc7\xfa\x9b\x5c\x03\x27\x46\x75\xc4\x3d\x16\x75\x41\x03\xfa\x3f\x48\x7f\xd8\x94\xf6\xa1\x34\x7f\xce\x4c\xf6\x21\xbd\x65\xdb\xa1\xda\xff\x9b\x9a\x87\x73\xee\x8c\xec\xe5\x42\xad\xf2\x82\x9a\xba\x83\xc0\x5e\x4e\x4c\x75\x27\x49\x26\xaa\xcb\x04\xe5\x47\x32\xc9\x07\xa9\x9c\xf6\x73\xed\xd7\x38\xf2\x72\xa9\xfb\x87\xff\xd1\x7b\xc1\xab\x4d\x4c\x99\x8d\xdd\x89\xdd\x01\xa4\xb7\xaa\x73\x89\xf6\xfa\xfa\x19\x75\xe8\xb7\x7e\xc9\x39\xf8\x49\x75\x2c\x59\xfd\x78\x20\xf8\x89\x1c\x08\xb6\x69\x54\xbb\xcd\xc0\x9a\xad\x2f\xea\x37\x33\x74\x66\xc7\xf3\x9f\x1d\x3e\x2c\xd2\xbe\x78\x7b\x28\x67\x3b\x01\xbb\xc5\x3b\xab\xe9\xbd\x7e\x62\x74\x73\xdb\xee\x7e\xd6\x8a\x59\x5b\x53\xe5\x04\x5e\x6e\xe9\xd6\xa6\x9b\xc1\x2e\x7a\xb4\x9e\x64\x5e\x55\xaa\xf6\x6b\x22\xc9\x7d\x0f\x31\xef\xe2\xc8\xf9\xa4\x7d\x64\xe9\xfd\xfc\xf8\xde\xb5\xb5\xb5\x54\xa1\xca\x23\x26\xb2\x2d\x64\xd5\x16\x2a\xf7\xd0\xbc\x8b\xf0\xed\x9f\xe6\xab\x1c\x8b\x7e\x79\x87\x91\x7b\x03\x46\xd5\x4e\xaa\x86\xd0\x4e\x65\x3f\xf6\x1f\xc7\xe9\xdc\xd4\x47\x78\x87\xd1\x4c\xff\x94\x8e\x81\x40\x34\x13\x32\x4d\x8a\xf6\x64\x42\x18\x17\x52\x37\xdd\xe3\x0b\xe0\x4a\x8e\xcd\x35\xc9\xf8\x4e\x4f\xa2\xbf\xc2\xae\xbf\x84\x7f\x95\x8f\x1e\xb8\xdd\xf5\xfb\x15\x1d\x6a\xee\x9e\x4a\x0e\x33\xd4\xd3\xa8\x26\x8c\x41\xb1\xf8\xf8\xff\xa3\x7e\x59\xe6\xb0\xeb\x83\x67\xc3\xfd\xbf\xd8\x97\x70\x0a\xfb\x8e\xef\xfc\x0d\x79\xa7\x0f\x2b\x7f\x59\x0f\xde\x12\x24\xde\x3a\xa1\x9c\x98\x2a\xde\x77\x21\x2c\x4d\xf1\xfa\x7f\x01\x00\x00\xff\xff\x23\x3d\x14\x51\xcd\x66\x00\x00"),
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
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error){
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
func MustFindFileReader(path string) (*bytes.Reader, int64){
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error){
  item, ok := assetFiles[path]
  if !ok {
    return nil,0, fmt.Errorf("File %q not found in file system", path)
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
func ReadFile(path string, doGzip bool) (string, error){
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
func ReadFileByte(path string, doGzip bool) ([]byte, error){
  reader, _, err := FindFile(path, doGzip)
  if err != nil {
    return nil, err
  }

  if closer, ok := reader.(io.Closer); ok {
    defer closer.Close()
  }

  var bu bytes.Buffer

  _, err = io.Copy(&bu, reader);
  if err != nil && err != io.EOF {
   return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
  }

  return bu.Bytes(), nil
}
