// Package pvss implements public verifiable secret sharing as introduced in
// "A Simple Publicly Verifiable Secret Sharing Scheme and its Application to
// Electronic Voting" by Berry Schoenmakers. In comparison to regular verifiable
// secret sharing schemes, PVSS enables any third party to verify shares
// distributed by a dealer using zero-knowledge proofs. PVSS runs in three steps:
//  1. The dealer creates a list of encrypted public verifiable shares using
//     EncShares() and distributes them to the trustees.
//  2. Upon the announcement that the secret should be released, each trustee
//     uses DecShare() to first verify and, if valid, decrypt his share.
//  3. Once a threshold of decrypted shares has been released, anyone can
//     verify them and, if enough shares are valid, recover the shared secret
//     using RecoverSecret().
package pvss

import (
	"errors"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/proof/dleq"
	"go.dedis.ch/kyber/v4/share"
)

// Suite describes the functionalities needed by this package in order to
// function correctly.
type Suite interface {
	kyber.Group
	kyber.HashFactory
	kyber.Encoding
	kyber.XOFFactory
	kyber.Random
}

var ErrTooFewShares = errors.New("not enough shares to recover secret")
var ErrDifferentLengths = errors.New("inputs of different lengths")
var ErrEncVerification = errors.New("verification of encrypted share failed")
var ErrDecVerification = errors.New("verification of decrypted share failed")
var ErrGlobalChallengeVerification = errors.New("failed to verify global challenge")
var ErrDecShareChallengeVerification = errors.New("failed to verify the share decryption challenge")

// PubVerShare is a public verifiable share.
type PubVerShare struct {
	S share.PubShare // Share
	P dleq.Proof     // Proof
}

// EncShares creates a list of encrypted publicly verifiable PVSS shares for
// the given secret and the list of public keys X using the sharing threshold
// t and the base point H. The function returns the list of shares and the
// public commitment polynomial.
func EncShares(
	suite Suite,
	H kyber.Point,
	X []kyber.Point,
	secret kyber.Scalar,
	t uint32,
) (shares []*PubVerShare, commit *share.PubPoly, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create secret sharing polynomial

// Create secret set of shares

// Create public polynomial commitments with respect to basis H

// Prepare data for encryption consistency proofs ...

// Create NIZK discrete-logarithm equality proofs

func computeCommitments(suite Suite, n uint32, polyComs []kyber.Point) []kyber.Point {
	_ = "STUB: not implemented"
	return nil
}

// Compute Xi = C0 + iC1 + (i^2)C2 + ... + (i^(t-1))C_(t-1) for i in [1, ..., n]
// Using Horner's method: Xi = C0 + i(C1 + i(C2 + i(....)))

// From j=t-1 to j = 1 since last C0 is not multiplied by ith

func computeGlobalChallenge(suite Suite, n uint32, commit *share.PubPoly,
	encShares []*PubVerShare) (kyber.Scalar, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), nil
}

// VerifyEncShare checks that the encrypted share sX satisfies
// log_{H}(sH) == log_{X}(sX) where sH is the public commitment computed by
// evaluating the public commitment polynomial at the encrypted share's index i.
func VerifyEncShare(suite Suite, H, X, sH kyber.Point, expGlobalChallenge kyber.Scalar, encShare *PubVerShare) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyEncShareBatch provides the same functionality as VerifyEncShare but for
// slices of encrypted shares. The function returns the valid encrypted shares
// together with the corresponding public keys.
func VerifyEncShareBatch(
	suite Suite,
	H kyber.Point,
	X, sH []kyber.Point,
	commit *share.PubPoly,
	encShares []*PubVerShare,
) ([]kyber.Point, []*PubVerShare, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// good public keys
// good encrypted shares

// Need to compute the global challenge and verify the encrypted shares

// DecShare first verifies the encrypted share against the encryption
// consistency proof and, if valid, decrypts it and creates a decryption
// consistency proof.
func DecShare(
	suite Suite,
	H, X, sH kyber.Point,
	x, expGlobalChallenge kyber.Scalar,
	encShare *PubVerShare,
) (*PubVerShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decryption: x^{-1} * (xS)

// DecShareBatch provides the same functionality as DecShare but for slices of
// encrypted shares. The function returns the valid encrypted and decrypted
// shares as well as the corresponding public keys.
func DecShareBatch(
	suite Suite,
	H kyber.Point,
	X, sH []kyber.Point,
	x kyber.Scalar,
	expGlobalChallenges []kyber.Scalar,
	encShares []*PubVerShare,
) ([]kyber.Point, []*PubVerShare, []*PubVerShare, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// good public keys
// good encrypted shares
// good decrypted shares

// VerifyDecShare checks that the decrypted share sG satisfies
// log_{G}(X) == log_{sG}(sX). Note that X = xG and sX = s(xG) = x(sG).
func VerifyDecShare(suite Suite, G, X kyber.Point, encShare *PubVerShare, decShare *PubVerShare) error {
	_ = "STUB: not implemented"
	// Compute challenge for the decShare
	return nil
}

// VerifyDecShareBatch provides the same functionality as VerifyDecShare but for
// slices of decrypted shares. The function returns the valid decrypted shares.
func VerifyDecShareBatch(
	suite Suite,
	G kyber.Point,
	X []kyber.Point,
	encShares []*PubVerShare,
	decShares []*PubVerShare,
) ([]*PubVerShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// good decrypted shares

// RecoverSecret first verifies the given decrypted shares against their
// decryption consistency proofs and then tries to recover the shared secret.
func RecoverSecret(
	suite Suite,
	G kyber.Point,
	X []kyber.Point,
	encShares []*PubVerShare,
	decShares []*PubVerShare,
	t, n uint32,
) (kyber.Point, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Point), nil
}
