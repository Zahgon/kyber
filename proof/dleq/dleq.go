// Package dleq provides functionality to create and verify non-interactive
// zero-knowledge (NIZK) proofs for the equality (EQ) of discrete logarithms (DL).
// This means, for two values xG and xH one can check that
//
//	log_{G}(xG) == log_{H}(xH)
//
// without revealing the secret value x.
package dleq

import (
	"errors"

	"go.dedis.ch/kyber/v4"
)

// Suite wraps the functionalities needed by the dleq package.
type Suite interface {
	kyber.Group
	kyber.HashFactory
	kyber.XOFFactory
	kyber.Random
}

var ErrDifferentLengths = errors.New("inputs of different lengths")
var ErrInvalidProof = errors.New("invalid proof")

// Proof represents a NIZK dlog-equality proof.
type Proof struct {
	C  kyber.Scalar // challenge
	R  kyber.Scalar // response
	VG kyber.Point  // public commitment with respect to base point G
	VH kyber.Point  // public commitment with respect to base point H
}

// NewDLEQProof computes a new NIZK dlog-equality proof for the scalar x with
// respect to base points G and H. It therefore randomly selects a commitment v
// and then computes the challenge c = H(xG,xH,vG,vH) and response r = v - cx.
// Besides the proof, this function also returns the encrypted base points xG
// and xH.
func NewDLEQProof(
	suite Suite,
	G kyber.Point,
	H kyber.Point,
	x kyber.Scalar,
) (proof *Proof, xG kyber.Point, xH kyber.Point, err error) {
	_ = "STUB: not implemented"
	// Encrypt base points with secret
	return nil, *new(kyber.Point), *new(kyber.Point), nil
}

// Commitment

// Challenge

// Response

// NewDLEQProofBatch computes lists of NIZK dlog-equality proofs and of
// encrypted base points xG and xH. Note that the challenge is computed over all
// input values.
func NewDLEQProofBatch(
	suite Suite,
	G []kyber.Point,
	H []kyber.Point,
	secrets []kyber.Scalar,
) (proof []*Proof, xG []kyber.Point, xH []kyber.Point, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// Encrypt base points with secrets

// Commitments

// Collective challenge

// Responses

// Verify examines the validity of the NIZK dlog-equality proof.
// The proof is valid if the following two conditions hold:
//
//	vG == rG + c(xG)
//	vH == rH + c(xH)
func (p *Proof) Verify(suite Suite, G kyber.Point, H kyber.Point, xG kyber.Point, xH kyber.Point) error {
	_ = "STUB: not implemented"
	return nil
}
