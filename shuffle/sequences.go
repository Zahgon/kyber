package shuffle

import (
	"crypto/cipher"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/proof"
)

// SequencesShuffle shuffles a sequence of ElGamal pairs based on Section 5 of
// "Verifiable Mixing (Shuffling) of ElGamal Pairs" by Andrew Neff (April 2004)
//
// The function expects X and Y to be the same dimension, with each row having
// the same length. It also expect X and Y to have at least one element. The
// function will panic if the expectations are not met.
//
// Dim X and Y: [<sequence length, j>, <number of sequences, i>]
//
// The number of rows defines the sequences length. The number of columns
// defines the number of sequences.
//
//	Seq 1  Seq 2  Seq 3
//	(0,0)  (0,1)  (0,2)
//	(1,0)  (1,1)  (1,2)
//	(2,0)  (2,1)  (2,2)
//
// # In the code coordinates are (j,i), where 0 ≤ j ≤ NQ-1, 0 ≤ i ≤ k-1
//
// Last coordinate is (NQ-1, k-1)
//
// Variable names are as representative to the paper as possible.
func SequencesShuffle(
	group kyber.Group,
	G, H kyber.Point,
	X, Y [][]kyber.Point,
	rand cipher.Stream) (xBar, yBar [][]kyber.Point, getProver func(e []kyber.Scalar) (proof.Prover, error)) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Pick a random permutation used in ALL k ElGamal sequences. The permutation
// (π) of an ElGamal pair at index i always outputs to the same index

// Fisher–Yates shuffle

// Pick a fresh ElGamal blinding factor β(j, i) for each ElGamal sequence
// and each ElGamal pair

// Perform the Shuffle

// EGAR 2 (Prover) - Standard ElGamal k-shuffle proof: Knowledge of
// (xUp, yUp), (xDown, yDown) and e[j]

// Need to consolidate beta to a one dimensional array

// assertXY checks that x, y have the same dimensions and at least one element
func assertXY(X, Y [][]kyber.Point) error { _ = "STUB: not implemented"; return nil }

// GetSequenceVerifiable returns the consolidated input and output of sequence
// shuffling elements. Needed by the prover and verifier.
func GetSequenceVerifiable(group kyber.Group, X, Y, Xbar, Ybar [][]kyber.Point, e []kyber.Scalar) (
	xUp, yUp, xDown, yDown []kyber.Point) {
	_ = "STUB: not implemented"

	// EGAR1 (Verifier) - Consolidate input and output
	return nil, nil, nil, nil
}

// No modification could be made for e[0] -> e[0] = 1 if one wanted -
// Remark 7 in the paper
