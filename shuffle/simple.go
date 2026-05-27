package shuffle

import (
	"crypto/cipher"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/proof"
)

// XX the Zs in front of some field names are a kludge to make them
// accessible via the reflection API,
// which refuses to touch unexported fields in a struct.

// P (Prover) step 0: public inputs to the simple k-shuffle.
type ssa0 struct {
	X []kyber.Point
	Y []kyber.Point
}

// V (Verifier) step 1: random challenge t
type ssa1 struct {
	Zt kyber.Scalar
}

// P step 2: Theta vectors
type ssa2 struct {
	Theta []kyber.Point
}

// V step 3: random challenge c
type ssa3 struct {
	Zc kyber.Scalar
}

// P step 4: alpha vector
type ssa4 struct {
	Zalpha []kyber.Scalar
}

// SimpleShuffle is the "Simple k-shuffle" defined in section 3 of
// Neff, "Verifiable Mixing (Shuffling) of ElGamal Pairs", 2004.
type SimpleShuffle struct {
	grp kyber.Group
	p0  ssa0
	v1  ssa1
	p2  ssa2
	v3  ssa3
	p4  ssa4
}

// Simple helper to compute G^{ab-cd} for Theta vector computation.
func thenc(grp kyber.Group, G kyber.Point,
	a, b, c, d kyber.Scalar) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Init initializes the simple shuffle with the given group and the k parameter
// from the paper.
func (ss *SimpleShuffle) Init(grp kyber.Group, k int) *SimpleShuffle {
	_ = "STUB: not implemented"
	return nil
}

// Prove the  "Simple k-shuffle" defined in section 3 of
// Neff, "Verifiable Mixing (Shuffling) of ElGamal Pairs", 2004.
// The Scalar vector y must be a permutation of Scalar vector x
// but with all elements multiplied by common Scalar gamma.
//
//nolint:funlen // 51 statement instead of authorized 50
func (ss *SimpleShuffle) Prove(g kyber.Point, gamma kyber.Scalar,
	x, y []kyber.Scalar, _ cipher.Stream,
	ctx proof.ProverContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Step 0: inputs
// (4)

// V step 1

// P step 2

// (5) and (6) xhat,yhat vectors

// (7) theta and Theta vectors

// V step 3

// P step 4

// (8)

// Simple helper to verify Theta elements,
// by checking whether A^a*B^-b = T.
// P,Q,s are simply "scratch" kyber.Point/Scalars reused for efficiency.
func thver(A, B, T, P, Q kyber.Point, aS, bS, s kyber.Scalar) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify for Neff simple k-shuffle proofs.
func (ss *SimpleShuffle) Verify(G, Gamma kyber.Point,
	ctx proof.VerifierContext) error {
	_ = "STUB: not implemented"

	// extract proof transcript
	return nil
}

// Validate all vector lengths

// check verifiable challenges (usually by reproducing a hash)

// fills in v1

// fills in v3

// Verifier step 5

// scratch variables
