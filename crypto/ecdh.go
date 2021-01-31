package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
)

type KeyPair struct {
	PublicKey  ecdsa.PublicKey
	privateKey *ecdsa.PrivateKey
}

func NewKeyPair(curve elliptic.Curve) (KeyPair, error) {
	pair, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return KeyPair{}, err
	}

	return KeyPair{
		privateKey: pair,
		PublicKey:  pair.PublicKey,
	}, nil
}

func (kp KeyPair) ComputeSecret(pubkey ecdsa.PublicKey) []byte {
	derived, _ := kp.PublicKey.Curve.ScalarMult(pubkey.X, pubkey.Y, kp.privateKey.D.Bytes())
	secret := sha256.Sum256(derived.Bytes())
	return secret[:]
}
