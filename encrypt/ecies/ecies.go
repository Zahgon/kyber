// Package ecies implements the Elliptic Curve Integrated Encryption Scheme (ECIES).
package ecies

import (
	"hash"

	"go.dedis.ch/kyber/v4"
)

// Encrypt first computes a shared DH key using the given public key, then
// HKDF-derives a symmetric key (and nonce) from that, and finally uses these
// values to encrypt the given message via AES-GCM. If the hash input parameter
// is nil then SHA256 is used as a default. Encrypt returns a byte slice
// containing the ephemeral elliptic curve point of the DH key exchange and the
// ciphertext or an error.
func Encrypt(group kyber.Group, public kyber.Point, message []byte, hash func() hash.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generate an ephemeral elliptic curve scalar and point

// Compute shared DH key

// Derive symmetric key and nonce via HKDF (NOTE: Since we use a new
// ephemeral key for every ECIES encryption and thus have a fresh
// HKDF-derived key for AES-GCM, the nonce for AES-GCM can be an arbitrary
// (even static) value. We derive it here simply via HKDF as well.)

// Encrypt message using AES-GCM

// Serialize ephemeral elliptic curve point and ciphertext

// Decrypt first computes a shared DH key using the received ephemeral elliptic
// curve point (stored in the first part of ctx), then HKDF-derives a symmetric
// key (and nonce) from that, and finally uses these values to decrypt the
// given ciphertext (stored in the second part of ctx) via AES-GCM. If the hash
// input parameter is nil then SHA256 is used as a default. Decrypt returns the
// plaintext message or an error.
func Decrypt(group kyber.Group, private kyber.Scalar, ctx []byte, hash func() hash.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reconstruct the ephemeral elliptic curve point

// Compute shared DH key and derive the symmetric key and nonce via HKDF

// Decrypt message using AES-GCM

func deriveKey(hash func() hash.Hash, dh kyber.Point, l int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
