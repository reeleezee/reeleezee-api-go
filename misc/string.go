/*
Licensed under MIT license
(c) 2017 Reeleezee BV
*/
package misc

import (
	"crypto/rand"
	"fmt"
)

func MaxString(data interface{}, length int) string {
	if data != nil {
		data := data.(string)
		if len(data) > length {
			return data[:length]
		}
		return data
	}
	return ""
}

// NO REAL UUID!
// Consider a better library:
//	https://github.com/satori/go.uuid
//	https://github.com/pborman/uuid
func PseudoUuidV4() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40 // Version 4
	b[8] = (b[8] & 0xbf) | 0x80 // Variant is 10
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}
