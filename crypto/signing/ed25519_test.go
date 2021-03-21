package signing

import (
	"encoding/base64"
	"testing"

	"github.com/mateigraura/wirebo-api/crypto"
	"github.com/stretchr/testify/assert"
)

func TestVerifyForCorrectSignature_ShouldPass(t *testing.T) {
	pubKey := "XXnByecIEjQ8Ir/10T/YnCGWX6W48BW+fgmF+PP7iWQ="
	sig := "TJwsQOAbAw9n0twlVXJ2P7FmthrVWaIX5N7j5j6ebxPY0FgpTbRWm7TkbN1jepvQvAXQpsAp8ZLR5OseZnVjBQ=="

	pubKeyBytes, _ := base64.StdEncoding.DecodeString(pubKey)
	sigBytes, _ := base64.StdEncoding.DecodeString(sig)
	msg := []byte("msg")

	err := VerifySignature(pubKeyBytes, msg, sigBytes)
	assert.Nil(t, err)
}

func TestVerifyForIncorrectPubKeyLength_ShouldErr(t *testing.T) {
	pubKey := "XXnByecIEjQ8Ir/10T/YnCGWX6W48BW+fgmF+PP7iWQQ=="
	sig := "TJwsQOAbAw9n0twlVXJ2P7FmthrVWaIX5N7j5j6ebxPY0FgpTbRWm7TkbN1jepvQvAXQpsAp8ZLR5OseZnVjBQ=="

	pubKeyBytes, _ := base64.StdEncoding.DecodeString(pubKey)
	sigBytes, _ := base64.StdEncoding.DecodeString(sig)
	msg := []byte("msg")

	err := VerifySignature(pubKeyBytes, msg, sigBytes)
	assert.Equal(t, crypto.ErrInvalidPublicKey, err)
}

func TestVerifyForIncorrectSig_ShouldErr(t *testing.T) {
	pubKey := "XXnByecIEjQ8Ir/10T/YnCGWX6W48BW+fgmF+PP7iWQ="
	sig := "TJwsQOAbAw9n0twlVXJ2P7FmthrVWaIX5n7j5j6ebxPY0FgpTbRWm7TkbN1jepvQvAXQpsAp8ZLR5OseZnVjBQ=="

	pubKeyBytes, _ := base64.StdEncoding.DecodeString(pubKey)
	sigBytes, _ := base64.StdEncoding.DecodeString(sig)
	msg := []byte("msg")

	err := VerifySignature(pubKeyBytes, msg, sigBytes)
	assert.Equal(t, crypto.ErrInvalidSignature, err)
}
