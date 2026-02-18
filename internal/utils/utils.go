package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func FindNonce(input []byte) uint64 {
	var nonce uint64
	buf := make([]byte, len(input)+8)

	copy(buf, input)

	for nonce = 0; nonce <= 10000000; nonce++ {
		for i := 0; i < 8; i++ {
			x := byte(nonce >> (8 * i))
			buf[len(input)+i] = x
		}

		hash := sha256.Sum256(buf)

		hexStr := hex.EncodeToString(hash[:])

		if hexStr[0] == '0' && hexStr[1] == '0' {
			return nonce
		}
	}

	return 0
}
