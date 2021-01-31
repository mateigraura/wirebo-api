package crypto

import (
	"bytes"
	"crypto/elliptic"
	"testing"
)

func TestPubExchange_SecretsShouldMatch(t *testing.T) {
	t.Parallel()

	alice, _ := NewKeyPair(elliptic.P256())
	bob, _ := NewKeyPair(elliptic.P256())

	aliceSecret := alice.ComputeSecret(bob.PublicKey)
	bobSecret := bob.ComputeSecret(alice.PublicKey)

	if !bytes.Equal(aliceSecret, bobSecret) {
		t.Error("Secrets do not match")
	}

	t.Logf("\nAlice secret %x\n", aliceSecret)
	t.Logf("\nBob secret %x\n", bobSecret)
}

func TestPubExchange_IncorrectCurve(t *testing.T) {
	t.Parallel()

	alice, _ := NewKeyPair(elliptic.P256())
	// different curve, secrets should not match
	bob, _ := NewKeyPair(elliptic.P224())

	aliceSecret := alice.ComputeSecret(bob.PublicKey)
	bobSecret := bob.ComputeSecret(alice.PublicKey)

	if bytes.Equal(aliceSecret, bobSecret) {
		t.Error("Secrets should not match")
	}

	t.Logf("\nAlice secret %x\n", aliceSecret)
	t.Logf("\nBob secret %x\n", bobSecret)
}
