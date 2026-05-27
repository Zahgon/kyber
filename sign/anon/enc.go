package anon

import (
	"go.dedis.ch/kyber/v4"
)

func header(suite Suite, _ kyber.Point, x kyber.Scalar,
	xb1, xb2 []byte, anonymitySet Set) []byte {
	_ = "STUB: not implemented"

	// Encrypt the master scalar key with each public key in the set
	return nil
}

// compute DH shared secret

// Create and encrypt a fresh key decryptable only by the given receivers.
// Returns the secret key and the ciphertext.
func encryptKey(suite Suite, anonymitySet Set) (k, c []byte) {
	_ = "STUB: not implemented"
	// Choose a keypair and encode its representation
	return nil, nil
}

// Generate the ciphertext header

// Decrypt and verify a key encrypted via encryptKey.
// On success, returns the key and the length of the decrypted header.
func decryptKey(
	suite Suite,
	ciphertext []byte,
	anonymitySet Set,
	mine int,
	privateKey kyber.Scalar,
) ([]byte, int, error) {
	_ = "STUB: not implemented"
	// Decode the (supposed) ephemeral public key from the front
	return nil, 0, nil
}

// Decode the (supposed) master secret with our private key

// Make sure it reproduces the correct ephemeral public key

// Regenerate and check the rest of the header,
// to ensure that any of the anonymitySet members could decrypt it

// constantTimeAllEq returns 1 iff all bytes in slice x have the value y.
// The time taken is a function of the length of the slices
// and is independent of the contents.
func constantTimeAllEq(x []byte, y byte) int { _ = "STUB: not implemented"; return 0 }

// macSize is how long the hashes are that we extract from the XOF.
// This constant of 16 is taken from the previous implementation's behavior.
const macSize = 16

// Encrypt a message for reading by any member of an explit anonymity set.
// The caller supplies one or more keys representing the anonymity set.
// If the provided set contains only one public key,
// this reduces to conventional single-receiver public-key encryption.
func Encrypt(suite Suite, message []byte,
	anonymitySet Set) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We now know the ciphertext layout

// Now encrypt and MAC the message based on the master secret

// Decrypt a message encrypted for a particular anonymity set.
// Returns the cleartext message on success, or an error on failure.
//
// The caller provides the anonymity set for which the message is intended,
// and the private key corresponding to one of the public keys in the set.
// Decrypt verifies that the message is encrypted correctly for this set -
// in particular, that it could be decrypted by ALL of the listed members -
// before returning successfully with the decrypted message.
//
// This verification ensures that a malicious sender
// cannot de-anonymize a receiver by constructing a ciphertext incorrectly
// so as to be decryptable by only some members of the set.
// As a side-effect, this verification also ensures plaintext-awareness:
// that is, it is infeasible for a sender to construct any ciphertext
// that will be accepted by the receiver without knowing the plaintext.
func Decrypt(suite Suite, ciphertext []byte, anonymitySet Set, mine int, privateKey kyber.Scalar) ([]byte, error) {
	_ = "STUB: not implemented"
	// Decrypt and check the encrypted key-header.
	return nil, nil
}

// Determine the message layout

// Decrypt the message and check the MAC
