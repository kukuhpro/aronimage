package storage

import "encoding/base64"

func Base64ToByte(base64EncodeString string) []byte {
	e64 := base64.StdEncoding
	enc := []byte(base64EncodeString)
	maxDecLen := e64.DecodedLen(len(enc))
	decBuf := make([]byte, maxDecLen)

	n, err := e64.Decode(decBuf, enc)
	_ = err

	return decBuf[0:n]
}
