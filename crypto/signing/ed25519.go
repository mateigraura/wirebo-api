package signing

import (
	libed25519 "crypto/ed25519"

	"github.com/mateigraura/wirebo-api/crypto"
)

func VerifySignature(publicKey, message, signature []byte) error {
	if len(publicKey) != libed25519.PublicKeySize {
		return crypto.ErrInvalidPublicKey
	}

	isValidSig := libed25519.Verify(publicKey, message, signature)
	if !isValidSig {
		return crypto.ErrInvalidSignature
	}

	return nil
}
