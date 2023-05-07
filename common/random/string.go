package random

import (
	"crypto/rand"
	"encoding/hex"
)

func String() string {
	var chunk [64]byte
	_, err := rand.Read(chunk[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(chunk[:])
}
