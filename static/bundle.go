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
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\xdd\x6f\xdb\xb6\x16\x7f\x4e\x80\xfc\x0f\x67\xc2\x50\x48\x9b\xa7\x6c\xf7\xd1\x5b\x06\x2c\x49\x7b\x51\x60\x6b\x7b\xdb\xee\xee\x61\x18\x0a\x5a\x3c\x76\xd8\x48\x94\xaf\x48\xc7\x31\x02\xff\xef\x17\xfc\xd2\x97\x65\x4b\xb6\x95\xd4\x49\xd3\x87\x3a\x92\xc8\xc3\xf3\xc5\x73\x7e\x87\xa4\x74\x7a\x0a\x77\x77\xe1\x07\x99\xcd\x22\x19\xbe\x1d\x7d\xc6\x48\x86\x6f\x48\x82\xcb\xe5\x2b\x86\x31\x15\x40\x71\xcc\x38\x0a\x20\x1c\x18\x97\x98\x8d\x49\x84\x30\xbf\x62\xd1\x15\xe0\xed\x34\x15\x28\x20\x41\x79\x95\x52\x90\x29\x64\x28\x67\x19\x07\x02\x09\x99\x42\x3a\x06\x12\xc7\x27\xc7\xa7\xa7\x40\xa4\xcc\xd8\x68\x26\x15\x1d\x21\xd2\x88\x11\x89\x14\xe6\x4c\x5e\x81\xbc\x42\x3b\x06\x05\xa1\xd9\x98\x65\x08\x44\x0d\x1c\x31\x8a\x14\x46\x0b\xdd\x26\x7f\x16\x9e\x1c\xcb\xc5\x14\x5b\xd8\x2e\x31\x7b\x77\x72\x7c\x64\xee\xfa\x01\xf8\x09\x99\xfe\x2d\x64\xc6\xf8\xe4\x9f\xbc\xc9\xdd\x72\x00\x98\x65\x69\x16\x9c\x1c\x2f\x4f\x8e\x35\xcb\x6b\xc8\x5f\xa4\x5c\xcc\x12\xcc\x36\xe9\x85\x44\x11\x4e\xa5\x28\xd4\x40\x89\x24\xf6\xd9\x9c\xc5\x31\x8c\x10\x22\x43\x87\xea\xb1\x18\x97\xa9\x16\x72\xc2\x6e\x18\x9f\x00\x4b\xa6\x31\x26\xc8\xa5\xba\xe8\x43\x29\x39\xd7\x55\xad\xd8\xdb\x6b\x74\x12\x18\x9d\x14\x2a\xf9\x23\xe5\x93\x94\x16\x92\xaf\x75\x08\xe2\x5c\x62\x9c\x66\xca\x27\x32\x86\x5a\x2e\xa2\xc9\x24\x8a\x4c\x78\x49\x24\x19\x11\x81\x40\x38\xb5\xb7\x3e\xa0\x10\x2c\xe5\x4e\x16\x3b\x5c\x95\xe5\x37\x38\x57\x56\xfc\x2e\x29\x91\x18\x80\xbe\xb4\xdd\x3b\x9b\xf2\xf2\xbc\x24\x4a\xa1\x65\x23\xca\x34\x4b\x6f\x18\x45\xb8\x3c\x87\x8b\xf7\x7f\x5e\x42\x3a\xc5\x8c\x48\x96\x72\xa1\x49\xce\x84\x92\x47\xb3\xad\xcc\xa2\x6c\x31\xe3\x14\xb3\x98\x71\x04\x3a\x6a\x31\xc7\xe5\xb9\x1d\xef\xee\xe4\x18\x20\x4a\x63\x30\xda\x57\x57\x74\x64\x05\x57\x17\x89\xd2\x5d\x24\xdc\x6f\xf8\x87\xf9\x55\x8f\x90\x8b\x59\x86\xf4\x35\xa7\x78\x0b\xa3\x34\x8d\xd5\x4d\xc6\xa3\x54\xf9\x8e\xc4\xfa\x7d\x8a\xb7\x28\xe0\xef\x7f\x94\xa2\xf4\xb3\x42\x3d\x6f\x70\x6e\x27\xae\x52\x04\xc7\x39\x30\x2e\x24\xe1\x11\x2a\xdf\x5d\x2b\x42\x78\x72\x3c\x9e\xf1\x48\x75\xf7\x0b\x11\x06\x90\xd4\xb9\x1d\x00\x24\xa9\x15\x6a\x90\xb3\x12\x86\x61\xce\x4b\x00\xdf\xad\x1d\x46\xab\x08\x5c\x64\x79\xd1\xd6\x4e\xfd\xa3\xa3\x21\x24\xe9\xa0\xb8\x11\xa5\xf1\x50\xfd\x57\xba\x65\x79\x1c\x42\x52\xba\x69\x79\x1b\xba\x3f\xec\xa3\x65\xa1\x2b\xa3\x75\xa3\x5c\x22\x25\x26\x6a\x9e\xcb\xd4\xde\x57\xf1\xce\x79\x0e\xcd\x25\xcd\x67\xb7\x98\x62\xc4\xc6\x2c\x52\xac\xc4\x18\x49\xe3\xeb\x5a\x89\x7e\x42\x47\x1b\x94\x10\x94\x07\xf6\xed\xc4\x04\x2b\x32\x1b\x43\x42\x47\x61\xc5\x21\x4a\xda\xb0\x9a\xe3\x2c\xce\xa5\x51\x3f\xe6\xff\x04\x86\x67\xb9\xbd\xde\xe0\xfc\x63\x46\x22\xf4\xbd\xf5\x46\x2f\xf1\xe1\x05\x86\x06\xc5\x31\x66\x9a\x05\x47\xe8\x65\xc2\xa4\xef\x2e\x5e\xf3\x71\xda\x99\xe2\xc0\xf5\xfa\x8b\xc9\x2b\xc3\x4c\x12\xbe\xe4\xd4\x0f\x82\x40\x99\xc0\xca\x1b\x23\x57\x1a\x0b\xad\x8a\x03\x38\x3b\x83\x1f\x5b\x84\x36\x7f\x9c\x9e\xc2\xeb\x31\xcc\x11\xae\x08\x55\xb1\xdb\x68\x72\x84\xe3\x34\x43\x63\x31\x98\x13\x01\x6e\x1a\x0d\x94\xe1\x38\x88\x6b\x36\x1d\xa8\x5e\x11\xe1\x12\x78\x2a\x73\x62\x42\xa6\x53\x6d\xf6\x74\x2a\x60\x84\x11\x99\x09\x3d\x6d\xc6\x84\xc5\xce\x07\xc2\x9c\xef\x6f\x56\x0c\xf5\xe2\x05\x18\x41\xaa\x13\xb7\x8b\x28\x34\x0f\x7f\xa2\x14\xf9\xb4\x41\xe9\x28\xa4\xa3\x50\xc7\xca\x7c\x6c\xf5\xec\x9b\x33\x45\xa7\x4c\x7d\xad\xd9\x5e\x2a\xc5\x8c\x7d\xef\x95\x11\x44\xa6\x10\x65\x48\x24\xba\xc1\x74\x6c\x67\x0d\x56\xf3\xbd\xc2\xbd\xbd\x81\x1e\x20\x4a\xe3\x7a\x1b\xad\x77\x4f\x73\x6c\x86\x32\x16\xae\x09\x8d\x59\x56\x17\x5a\x3b\x9b\xe5\x21\xbc\x88\x53\x81\x7e\xee\x19\xc5\xc0\x4a\x0b\x4e\x3f\xe1\x85\xef\x98\x70\x0d\x15\xef\x9f\x6c\x2c\x52\x4d\x33\xc2\x27\x08\x25\x8f\x2a\xab\xc8\xea\x6e\x78\x56\x9e\xb7\x2f\x4b\xf3\x51\xf7\x09\x7e\x5e\xa3\xe1\x2d\xb5\x6c\x23\x89\xd3\xf2\xee\x1a\x36\x3d\xad\x90\x1d\xd5\xbf\xca\x75\xdd\x31\xcf\x40\x66\x33\xac\xb6\xab\x1b\xab\x62\xb0\x8d\xe2\x9b\xd8\xf0\x61\x16\x45\x88\x26\x64\x1a\xf9\x55\x72\x2d\x19\xb3\x2f\x25\x04\x35\x67\x5a\x99\x8d\x4e\xba\xe2\xf1\x06\xb6\x5f\x31\xce\xc4\x15\x52\x20\x94\x6a\xd4\xd6\x9d\xcb\xa0\x92\xd4\xf4\xd4\x76\x19\xe6\x22\x9d\x71\x59\xc9\x2d\xb6\x95\xca\x20\x32\x95\x24\x06\x3e\x4b\x46\x98\xa9\x28\x93\x61\x94\x66\x14\xc6\x59\x9a\x18\x20\x3d\xea\x9c\x50\xf4\x38\x7e\x24\x6f\x15\x14\x95\x78\x2b\xc3\x0b\xf3\x1b\x80\xcf\xb8\x74\x28\xca\x79\xf2\xd6\x89\x42\xd3\xef\x29\x45\x58\x5a\x1d\x93\x83\x93\xe7\xb5\x78\x79\x3b\x65\x19\x52\x25\x65\x50\x9e\x92\x76\x3a\x8f\x13\x99\x4f\x40\x2b\x3c\x5c\x11\xa1\x60\xac\xea\xe6\x05\x9d\x7c\x78\x75\x0a\x4f\x50\x3a\xcb\x44\x0d\x9c\xf7\x19\x20\x7f\xf8\x69\xd0\x10\x24\x8b\x88\x55\x78\xb8\x85\x0e\xeb\xa2\xd4\x16\xe2\x91\xe9\x34\x5e\xf4\x1f\xfa\x3b\xca\xf6\xd0\x59\xef\x3e\x8d\xd9\x55\xe4\x4d\x39\xef\x7f\x33\xcc\x16\x4a\xfc\x91\x48\x79\xf8\xc7\x5d\xde\x4b\xc7\x8a\x5c\x39\x0d\xc9\x30\x7c\xc5\x38\xf5\x75\xff\xc0\x4c\x31\x3f\xf8\xf9\xc0\x14\xa7\xb9\xf3\x06\x46\xca\x7e\xb5\xda\x12\x8d\x2e\x51\x25\x3d\x6a\x45\xe8\x81\xf9\x9c\x33\x17\xcf\x73\xfb\x14\xc1\xdf\x0c\x5a\x8b\xfe\x49\x7a\x83\x3a\xba\xaf\x46\x7b\x5b\x87\xaa\x8b\xbc\xec\x98\xce\x46\x31\x8b\x5e\x5f\x86\x9a\xe2\x7b\xdd\x47\xe4\x0d\x99\x50\x25\x6d\x32\x13\x2a\xd4\xdd\x20\x10\xdb\x1e\x18\x85\x1b\x12\xcf\x70\xa0\xc2\x5f\x86\x42\x20\x05\x64\xf2\x0a\x33\x18\x2d\x80\x68\xf7\x82\x34\x83\xcf\xea\x57\x92\x89\xa6\x9e\x72\xb7\x6c\x81\xbc\x54\x29\xbe\x23\xd1\x35\x99\xe0\x72\x19\xae\x89\xe9\xb6\xfc\xed\x9c\xac\x8c\x5e\x9a\xb2\xd5\x20\x97\xd7\x16\xa0\xb5\xe2\x68\xeb\xbc\x65\x86\xea\x29\x71\x39\x62\x35\xc7\x70\x2c\x7b\x05\xf7\x5f\x22\xb9\xed\x30\xa9\xa9\x71\xd0\x35\x93\xa2\x55\xae\xfb\xad\x0d\x1e\x65\xda\x6b\x95\xea\xa1\x13\xde\x41\x9b\x78\xab\x54\x58\xd0\x2b\xd8\x1e\xe6\x6c\x0f\xd6\x7a\x4f\x53\xae\x7c\xaf\xc3\xb0\xcd\x96\x3d\x78\xd3\x66\x2d\xef\x9d\x1a\x3b\x98\x69\x57\x13\x3c\x74\xde\xec\x20\x59\x3d\xb5\x56\x0b\x2a\xb3\x6c\x51\xce\xa9\x84\xd2\x72\x42\xcd\x17\xe8\x9a\x13\x6a\x79\x35\x54\x5e\x61\x6d\x49\xb9\x35\xd7\x7d\xc1\x3c\xbc\x57\xce\x35\x7a\x6b\xce\xb9\x18\x63\xb2\x8d\x0e\xf6\x4d\xca\x86\x97\xbe\xaa\x49\x4b\x6c\xbd\x63\x29\xf1\xc2\x77\x8f\x2b\x33\xdb\xd5\xb9\xf6\xb0\xbd\x41\xb8\xe7\xf4\x7c\xf8\xe9\xb9\xba\x0a\x7b\x90\x86\xde\x98\xa4\xef\xee\x94\x1a\x7c\xe5\xf1\xaf\x54\x1c\xb2\xb3\x14\x3c\xb3\x45\xeb\x01\x04\xb0\x5c\x16\xa3\x8c\xf5\xed\x5c\xa9\x5a\x28\xb7\x9b\xbb\xb2\x46\xdb\x71\xf5\xb5\xfa\x18\x8a\xed\xa0\x35\x65\x6c\xbe\x7b\x3c\x56\x31\x6c\x4d\x78\xcd\x67\xde\x7a\xea\x9b\x54\xdf\xd6\x4b\x09\x6e\x8d\xda\xa1\x71\x83\xe5\x6a\x9d\x82\x2d\xd7\x70\x37\xc3\xa3\xd7\x5c\x60\x26\x7d\x03\xbc\x7c\x63\xb3\xa0\xaf\x25\x71\xeb\xf2\xad\x8a\xdf\xc1\xc3\xcb\x4a\xdd\xca\xf9\x3b\xe9\xac\x25\x45\x5d\x6c\x0c\xd9\xdb\xf2\x1f\xb8\xf9\x85\xb1\xc0\xf2\x0c\xaa\xc1\x62\xff\xee\x4e\x9f\x4b\xb0\x79\xf0\x37\xb3\xc4\x6e\x48\x81\xa7\xda\x78\xe0\x7d\xd6\x3f\xcb\x65\xb0\xb5\x0b\x6c\x46\xc8\x07\x63\xf9\x5d\x16\x96\x0e\xcb\xf6\x2b\xab\x4b\xd6\xfa\x9c\x2e\x97\x9b\xf0\xf0\xbf\x51\xfe\x16\xc7\xef\x30\x7b\x47\x26\xe8\x4e\x67\xa0\xd0\x9b\x98\x99\x45\xa9\xe5\x25\x26\xc2\x69\xe9\x78\x80\x88\x59\xfd\x5c\x40\xfb\x6a\x8f\x5c\x4c\xf1\xb1\xc2\xe0\x8a\xba\x9a\xd1\x70\x9a\x51\x95\xef\xec\xf9\x07\x7d\x75\xbe\xc8\xaf\xa7\x4a\xcd\x7a\x47\x25\x43\x31\x4d\xb9\x40\xa7\x7b\xc6\x65\x00\xe0\xff\xfd\xcf\x16\xba\x1c\x00\xf4\xb1\x3b\x63\xa4\xea\x09\x50\x3b\x62\x6d\x40\xf9\x48\xcc\x99\x8c\xae\xac\x66\x44\xf8\x31\xfd\x3d\x9d\x63\xe6\x6b\x8d\x69\x59\x8e\x22\x22\x10\x3c\x2a\x22\x6f\x00\x1e\x45\x11\x79\xc3\x62\x4a\x39\xcd\x9e\x81\xf7\x83\x07\xdf\xbb\xeb\x93\xe3\xa3\xe5\xe1\xe0\x70\x37\xa1\xf6\x0a\xeb\x9d\x80\x17\x67\xf1\xe0\x29\x6d\x00\x75\x13\xef\x88\x8d\xcd\x94\xfa\xe5\x0c\x7e\x84\x17\x2f\x56\x66\xd5\x2f\xf6\x28\xc8\xd1\x91\x8d\x66\x15\x40\x6e\x7c\xf5\x7c\xf1\x56\xf9\x8e\xf2\x0c\x3b\x61\xf3\x79\x1b\x98\x9e\x9a\x85\x9c\x40\x8c\xdc\xb7\x17\x81\x65\xc8\xf8\xdc\x91\x09\xa8\x6b\x36\x66\x45\x78\x72\x7c\xa4\x1f\xbd\x6f\x60\x25\xdf\x81\x0d\x8c\x54\x55\xbb\xe4\x4c\x54\xf5\x60\x87\xbd\x21\x99\x19\xf3\x2f\xc2\x25\xba\x63\x55\x1f\xd3\x0f\x92\x64\x52\x45\x88\xba\xaa\x7e\x6a\x52\xd5\xaf\x4e\x53\x25\x52\x70\x56\x6f\xa6\x1a\x54\xc8\x9f\xc1\x8f\x8a\x11\x50\x40\xa3\x43\x7f\xf8\x4e\x73\xd1\x40\xa6\xdc\xed\x14\xfe\xa5\x79\xce\x99\xfe\x15\x7e\x32\xc4\x2b\xbd\xbe\xff\x5e\xdd\x5a\xe6\x8a\xd8\x80\xed\x6b\xeb\x52\xe7\xc3\xff\xa8\x4c\x39\xac\x04\x74\x2f\x18\xc0\x6a\x0f\xe5\xa6\x16\xf3\xbb\x5b\xfa\xb2\x06\x64\x3c\xa1\x38\x62\x7c\xf2\xc9\xcc\x85\xa1\xc3\x49\x25\x7e\x6b\xa8\xdb\xd3\x22\x7f\xb2\xee\xf1\x69\xae\x65\xf7\x86\x15\x5b\xd6\x7a\x68\xbf\xcc\x69\x43\x25\x18\x36\xb6\x3d\x5f\xd4\x5b\xdb\xdb\xf5\xd6\x4a\xcd\xab\x84\xb5\xf6\xeb\x4d\x6b\x26\x75\xbd\x6a\xb7\x4b\xbd\x96\xae\xde\x08\x0e\xb4\x50\x7e\x90\x38\xbc\xdb\xce\xad\x9a\xdc\x19\x93\x98\x08\xd8\x0a\x1b\xb4\x14\xd8\xf6\xb4\xef\x4a\x85\xad\x86\xa3\x6e\xb8\xe6\xa3\xc0\x5d\x6b\x80\xf2\x8e\xf2\x87\x6b\x36\xf5\xcb\x53\x21\x08\x7f\x67\xca\x64\x25\x5f\x0f\xc2\x0f\x69\x26\x7d\x17\x7a\xc3\xdf\xe2\xd8\x7f\x61\x78\xe9\xab\x84\xc8\xf3\x71\x19\xdf\xae\x3f\xd7\xaa\xb1\xaa\xc1\xbf\x74\xf4\x25\x2b\x8b\x55\xaf\x82\x7a\x89\xe1\x4e\xb1\x49\x4c\x8a\x43\x6c\xd6\x92\x35\x85\x29\x1b\x6f\xbb\x68\x5b\xa5\x50\x18\x5e\x2f\xc1\xb8\xa3\xe3\x6a\xb4\x8d\xa6\x6a\x15\x09\x8a\x33\xa9\x79\x7b\x23\xc3\x99\x42\x1a\xc8\xa9\x6f\xae\x6d\x9d\x5b\x51\x86\x75\x78\x9d\x89\x96\x0f\xeb\xa6\xd9\xb3\x9b\x6e\x72\x53\x67\x19\x4e\x61\xa5\x1e\x75\xf6\xac\x42\xa3\x86\x2a\xd5\xe2\xb4\xe7\x2a\xb5\x5b\x95\x5a\x82\xb5\x6b\xaa\xd4\x7a\x79\xba\x4b\xfd\xd9\x4b\xe9\x69\x59\xed\xb5\x02\xcd\x69\x3e\xbd\x42\xf4\x31\x94\xa2\x4f\xb8\x0c\x3d\xf8\x6d\x9f\xfb\x97\x77\x27\x24\x7b\x30\x58\xb4\x1f\x94\xb9\xcd\xbe\xd1\x9e\x09\xfd\x7e\x36\x91\x6a\x49\xff\xde\xb6\x91\xaa\x4e\x04\x75\xe0\xba\x5f\x81\x03\x8f\x15\xfa\xf6\x07\x7b\x57\x66\xad\x6b\x9a\xbf\x25\xd4\x04\x8a\x95\x26\xf6\xd2\xfa\x5e\x13\xee\x19\x2e\xb7\x4f\x0c\xdb\xaa\xc9\x98\x16\x47\x97\xf0\xf1\xf9\x42\x2f\x4a\x95\xc1\x71\xc7\x23\xc2\x7a\x7f\x16\xae\x71\xa1\x61\xb3\x46\xb0\x9a\xa8\x03\xd0\xaa\xf1\x57\x84\x9c\xad\x22\x9b\x61\xb3\xd2\x92\xdb\xca\xd1\x1c\x43\xe5\x15\x69\xf0\x1f\x1a\x3c\x5b\x6e\xfb\x43\xce\x39\xc1\xaa\x57\x5f\xe3\xc2\x4a\xfc\x48\x4e\x40\xb5\xc0\xdd\x75\xf2\xf4\xfb\x06\xc7\x16\xde\x60\xbf\x3a\xf0\x14\x30\x73\xff\x6a\x38\x2c\x7c\xfd\x48\x7c\x67\x2b\x90\x7e\x8d\x8b\xa1\x91\x69\x3f\xb8\xae\x31\xd8\x3a\xa8\xbe\x0b\x76\x78\xcb\xd1\xa0\x85\xaf\x0a\x2c\x34\xc2\x85\x5d\x3d\x62\x05\x58\xec\x06\x82\x9b\xac\xb7\x1d\x04\xee\x5b\x0e\x77\xec\x26\xc6\x64\x0d\xd8\x6d\x70\xce\xbd\x65\x7e\xdc\x1e\xbb\xd1\x1b\x7b\x89\x5f\xf7\x68\x69\xa5\xd5\x56\x24\xbc\x05\x04\x76\x2f\x8b\x95\xd7\x8b\xbf\x32\xb8\xdb\xf1\x35\xba\x87\x47\xb6\xfd\x41\xda\x27\xf4\xc6\x5d\xdb\x22\xee\x83\xbe\x90\xf5\x75\xa2\xdb\xbe\x95\x70\x58\xd8\xf6\x91\x79\xd0\x56\x18\xd7\xca\xf6\x89\xd1\xd2\x8b\x7f\xcf\x70\xf7\x00\xe0\xee\x7d\x42\x87\x67\xb0\xfb\x0c\x76\x9f\x28\xd8\xfd\x73\x4a\x55\x28\x9f\x89\x67\xa8\xdb\x0a\x75\x8d\xae\x3a\xa1\xdd\x87\x7f\xa3\xd5\x30\xd7\x13\xe2\x75\xc4\x9e\x0c\xe8\x1d\xeb\x4f\x6b\x0d\x1c\x1b\x55\x8a\x3b\x4c\xea\x02\x06\xf4\xff\x3e\xfc\x61\x43\xda\xfb\x92\xfc\x29\x23\xd9\xfb\xf4\x96\xb6\x77\x63\xbf\x9d\x9a\xcd\x39\x35\xb5\xd4\x1f\xe7\x0b\x35\xcb\x0b\x68\xea\xde\xe7\xf5\xa0\x74\xde\xf7\xdb\xa9\x24\x13\xd5\x65\x82\xf2\x23\x99\xe4\x44\x2a\x6f\xf3\xb9\xf6\x2b\x18\xf9\xee\x4e\xf7\x0f\xff\xab\xd7\x82\x97\x4d\x48\x99\x8d\x0d\x37\xfa\x00\xd9\xb5\xea\x5c\x82\xbd\xbe\xde\x9f\x0e\xfd\x8d\x1f\x64\x0e\x7e\x56\x1d\x4b\x56\xaf\xbd\xd7\x6b\xc9\x3f\xbf\xda\x7b\x6f\x67\x32\x76\x7d\xb5\xd7\xa6\x51\xed\x36\x03\x6b\xb6\xbe\xa0\xdf\xcc\xc0\x99\x2d\xdf\xef\xec\xf0\x7d\x90\xcd\x93\xb7\x87\x72\xb6\x93\x62\x5b\xbc\xb3\x9a\xde\xeb\x6f\x84\x36\xb7\xed\xee\x67\x5b\x9c\xfc\xa9\x36\x55\x4e\xe0\xe5\x96\xde\xd8\xb4\x59\xd9\x45\x8f\xa6\x5d\xae\xd5\x0f\xd9\x6a\x9e\x2e\x89\x24\xfb\xbe\xa4\xbc\x8d\x23\xe7\x83\xf6\x91\xa5\x77\xf3\xe3\xbd\x6b\x6b\x6b\xa9\x42\x94\x07\x4c\x64\x2d\x60\xd5\x16\x2a\x7b\x48\xde\x85\xf9\xcd\x5f\xd8\xab\x1c\x28\x7f\x79\x8b\x91\x3b\x01\xa3\x6a\x27\x55\x43\x68\xa7\xb2\xdf\xec\x8f\xe3\x74\x6e\xea\x23\xbc\xc5\x68\xa6\x1f\xa5\x63\x20\x10\xcd\x84\x4c\x93\xa2\x3d\x99\x10\xc6\x85\xd4\x4d\x77\xf8\x90\xb7\xe2\xa3\xb9\x26\x19\xdf\xea\x41\xf4\xc7\xd4\xf5\x07\xed\x2f\x72\xea\x81\x5b\x5d\xdf\xaf\xe8\x50\x63\xf7\x54\x72\x18\x52\x8f\xa3\x9a\x30\x06\xc5\xe2\x1b\xfe\x0f\xfa\x81\x98\xc3\xae\x0f\x9e\x0c\xf6\xff\x62\x1f\xb4\x29\xec\x3b\xbe\xf5\x1b\xf2\x4e\x1f\x56\xfe\xb2\x1e\xdc\x12\x24\xde\x3a\xa6\x1c\x9b\x2a\xde\x77\x01\x2c\xeb\xe2\xf5\xff\x03\x00\x00\xff\xff\x66\xd6\xed\xe3\x94\x66\x00\x00"),
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
