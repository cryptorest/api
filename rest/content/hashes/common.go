package hashes

import (
	"hash"
)

func hashSum(data []byte, h hash.Hash) []byte {
	h.Write(data)

	return h.Sum(nil)
}
