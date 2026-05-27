package ibe

import (
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
)

type Ciphertext struct {
	// Random point rP
	U kyber.Point
	// Sigma attached to ID: sigma XOR H(rG_id)
	V []byte
	// ciphertext of the message M XOR H(sigma)
	W []byte
}

// H2Tag is the domain separation tag for the H2 hash function
func H2Tag() []byte { _ = "STUB: not implemented"; return nil }

// H3Tag is the domain separation tag for the H3 hash function
func H3Tag() []byte { _ = "STUB: not implemented"; return nil }

// H4Tag is the domain separation tag for the H4 hash function
func H4Tag() []byte { _ = "STUB: not implemented"; return nil }

// EncryptCCAonG1 implements the CCA identity-based encryption scheme from
// https://crypto.stanford.edu/~dabo/pubs/papers/bfibe.pdf for more information
// about the scheme.
// - master is the master key on G1
// - "identities" (rounds) are on G2
// - the Ciphertext.U point will be on G1
// - ID is the ID towards which we encrypt the message
// - msg is the actual message
// - seed is the random seed to generate the random element (sigma) of the encryption
// The suite must produce points which implements the `HashablePoint` interface.
//
//nolint:dupl // unavoidable
func EncryptCCAonG1(s pairing.Suite, master kyber.Point, ID, msg []byte) (*Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Compute Gid = e(master,Q_id)

// 2. Derive random sigma

// 3. Derive r from sigma and msg

// 4. Compute U = rP

// 5. Compute V = sigma XOR H2(rGid)
// even in Gt, it's additive notation

// 6. Compute M XOR H(sigma)

// DecryptCCAonG1 decrypts ciphertexts encrypted using EncryptCCAonG1 given a G2 "private" point
func DecryptCCAonG1(s pairing.Suite, private kyber.Point, c *Ciphertext) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Compute sigma = V XOR H2(e(rP,private))

// 2. Compute M = W XOR H4(sigma)

// 3. Check U = rP

// EncryptCCAonG2 implements the CCA identity-based encryption scheme from
// https://crypto.stanford.edu/~dabo/pubs/papers/bfibe.pdf for more information
// about the scheme.
// - master is the master key on G2
// - identities ("round") are on G1
// - the Ciphertext.U point will be on G2
// - ID is the ID towards which we encrypt the message
// - msg is the actual message
// - seed is the random seed to generate the random element (sigma) of the encryption
// The suite must produce points which implements the `HashablePoint` interface.
//
//nolint:dupl // unavoidable
func EncryptCCAonG2(s pairing.Suite, master kyber.Point, ID, msg []byte) (*Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Compute Gid = e(Q_id, master)

// 2. Derive random sigma

// 3. Derive r from sigma and msg

// 4. Compute U = rP

// 5. Compute V = sigma XOR H2(rGid)
// even in Gt, it's additive notation

// 6. Compute M XOR H(sigma)

// DecryptCCAonG2 decrypts ciphertexts encrypted using EncryptCCAonG2 given a G1 "private" point
func DecryptCCAonG2(s pairing.Suite, private kyber.Point, c *Ciphertext) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Compute sigma = V XOR H2(e(rP,private))

// 2. Compute M = W XOR H4(sigma)

// 3. Check U = rP

// hash sigma and msg to get r
func h3(s pairing.Suite, sigma, msg []byte) (kyber.Scalar, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), nil
}

// we hash it a first time: buffer = hash("IBE-H3" || sigma || msg)

// We will hash iteratively: H(i || H("IBE-H3" || sigma || msg)) until we get a
// value that is suitable as a scalar.

// We then apply masking to our resulting bytes at the bit level
// but we assume that toMask is a few bits, at most 8.
// For instance when using BLS12-381 toMask == 1.

// NOTE: Here we unmarshal as a test if the buffer is within the modulo
// because we know unmarshal does this test. This implementation
// is almost generic if not for this line. TO make it truly generic
// we would need to add methods to create a scalar from bytes without
// reduction and a method to check if it is within the modulo on the
// Scalar interface.

// if we didn't return in the for loop then something is wrong

func h4(s pairing.Suite, sigma []byte, length int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gtToHash(s pairing.Suite, gt kyber.Point, length int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func xor(a, b []byte) []byte { _ = "STUB: not implemented"; return nil }

type CiphertextCPA struct {
	// commitment
	RP kyber.Point
	// ciphertext
	C []byte
}

// EncryptCPAonG1 implements the CPA identity-based encryption scheme from
// https://crypto.stanford.edu/~dabo/pubs/papers/bfibe.pdf for more information
// about the scheme.
// SigGroup = G2 (large secret identities)
// KeyGroup = G1 (short master public keys)
// P random generator of G1
// dist master key: s, Ppub = s*P \in G1
// H1: {0,1}^n -> G1
// H2: GT -> {0,1}^n
// ID: Qid = H1(ID) = xP \in G2
//
//	secret did = s*Qid \in G2
//
// Encrypt:
//   - random r scalar
//   - Gid = e(Ppub, r*Qid) == e(P, P)^(x*s*r) \in GT
//     = GidT
//   - U = rP \in G1,
//   - V = M XOR H2(Gid)) = M XOR H2(GidT)  \in {0,1}^n
func EncryptCPAonG1(s pairing.Suite, basePoint, public kyber.Point, ID, msg []byte) (*CiphertextCPA, error) {
	_ = "STUB: not implemented"
	return nil,

		// we're using blake2 as XOF which only outputs 2^16-1 length
		nil
}

// e(Qid, Ppub) = e( H(round), s*P) where s is dist secret key

// H(gid)

// DecryptCPAonG1 implements the CPA identity-based encryption scheme from
// https://crypto.stanford.edu/~dabo/pubs/papers/bfibe.pdf for more information
// about the scheme.
// SigGroup = G2 (large secret identities)
// KeyGroup = G1 (short master public keys)
// Decrypt:
//   - V XOR H2(e(U, did)) = V XOR H2(e(rP, s*Qid))
//     = V XOR H2(e(P, P)^(r*s*x))
//     = V XOR H2(GidT) = M
func DecryptCPAonG1(s pairing.Suite, private kyber.Point, c *CiphertextCPA) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
