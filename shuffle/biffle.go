package shuffle

import (
	"crypto/cipher"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/proof"
)

func bifflePred() proof.Predicate {
	_ = "STUB: not implemented"

	// Branch 0 of either/or proof (for bit=0)
	return *new(proof.Predicate)
}

// Branch 1 of either/or proof (for bit=1)

func bifflePoints(suite Suite, G, H kyber.Point,
	X, Y, Xbar, Ybar [2]kyber.Point) map[string]kyber.Point {
	_ = "STUB: not implemented"
	return nil
}

// Biffle is a binary shuffle ("biffle") for 2 ciphertexts based on general ZKPs.
func Biffle(suite Suite, G, H kyber.Point,
	X, Y [2]kyber.Point, rand cipher.Stream) (
	Xbar, Ybar [2]kyber.Point, prover proof.Prover) {
	_ = "STUB: not implemented"

	// Pick the single-bit permutation.
	return nil, nil, *new(proof.Prover)
}

// Pick a fresh ElGamal blinding factor for each pair

// Create the output pair vectors

// BiffleVerifier returns a verifier of the biffle
func BiffleVerifier(suite Suite, G, H kyber.Point,
	X, Y, Xbar, Ybar [2]kyber.Point) (
	verifier proof.Verifier) {
	_ = "STUB: not implemented"
	return *new(proof.Verifier)
}
