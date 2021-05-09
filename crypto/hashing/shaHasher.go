package hashing

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type ShaHasher struct {
}

func (sh *ShaHasher) HashObj(obj interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(obj)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(buf.Bytes())
	return hash[:], nil
}

func (sh *ShaHasher) Hash(content []byte) ([]byte, error) {
	return nil, nil
}
