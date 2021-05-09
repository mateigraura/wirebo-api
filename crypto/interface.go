package crypto

type Signer interface {
	Verify(publicKey, message, signature []byte) error
}

type Hasher interface {
	HashObj(obj interface{}) ([]byte, error)
	Hash(content []byte) ([]byte, error)
}
