// Package share implements Shamir secret sharing and polynomial commitments.
// Shamir's scheme allows you to split a secret value into multiple parts, so called
// shares, by evaluating a secret sharing polynomial at certain indices. The
// shared secret can only be reconstructed (via Lagrange interpolation) if a
// threshold of the participants provide their shares. A polynomial commitment
// scheme allows a committer to commit to a secret sharing polynomial so that
// a verifier can check the claimed evaluations of the committed polynomial.
// Both schemes of this package are core building blocks for more advanced
// secret sharing techniques.
package share

import (
	"crypto/cipher"
	"errors"

	"go.dedis.ch/kyber/v4"
)

// Some error definitions
var errGroups = errors.New("non-matching groups")
var errCoeffs = errors.New("different number of coefficients")

// PriShare represents a private share.
type PriShare struct {
	I uint32       // Index of the private share
	V kyber.Scalar // Value of the private share
}

// Hash returns the hash representation of this share
func (p *PriShare) Hash(s kyber.HashFactory) []byte { _ = "STUB: not implemented"; return nil }

func (p *PriShare) String() string { _ = "STUB: not implemented"; return "" }

// PriPoly represents a secret sharing polynomial.
type PriPoly struct {
	g      kyber.Group    // Cryptographic group
	coeffs []kyber.Scalar // Coefficients of the polynomial
}

// NewPriPoly creates a new secret sharing polynomial using the provided
// cryptographic group, the secret sharing threshold t, and the secret to be
// shared s. If s is nil, a new s is chosen using the provided randomness
// stream rand.
func NewPriPoly(group kyber.Group, t uint32, s kyber.Scalar, rand cipher.Stream) *PriPoly {
	_ = "STUB: not implemented"
	return nil
}

// CoefficientsToPriPoly returns a PriPoly based on the given coefficients
func CoefficientsToPriPoly(g kyber.Group, coeffs []kyber.Scalar) *PriPoly {
	_ = "STUB: not implemented"
	return nil
}

// Threshold returns the secret sharing threshold.
func (p *PriPoly) Threshold() uint32 { _ = "STUB: not implemented"; return 0 }

// Secret returns the shared secret p(0), i.e., the constant term of the polynomial.
func (p *PriPoly) Secret() kyber.Scalar {
	_ = "STUB: not implemented"

	// Eval computes the private share v = p(i).
	return *new(kyber.Scalar)
}

func (p *PriPoly) Eval(i uint32) *PriShare { _ = "STUB: not implemented"; return nil }

// Shares creates a list of n private shares p(1),...,p(n).
func (p *PriPoly) Shares(n uint32) []*PriShare { _ = "STUB: not implemented"; return nil }

// Add computes the component-wise sum of the polynomials p and q and returns it
// as a new polynomial.
func (p *PriPoly) Add(q *PriPoly) (*PriPoly, error) { _ = "STUB: not implemented"; return nil, nil }

// Equal checks equality of two secret sharing polynomials p and q. If p and q are trivially
// unequal (e.g., due to mismatching cryptographic groups or polynomial size), this routine
// returns in variable time. Otherwise it runs in constant time regardless of whether it
// eventually returns true or false.
func (p *PriPoly) Equal(q *PriPoly) bool { _ = "STUB: not implemented"; return false }

// Commit creates a public commitment polynomial for the given base point b or
// the standard base if b == nil.
func (p *PriPoly) Commit(b kyber.Point) *PubPoly { _ = "STUB: not implemented"; return nil }

// Mul multiples p and q together. The result is a polynomial of the sum of
// the two degrees of p and q. NOTE: it does not check for null coefficients
// after the multiplication, so the degree of the polynomial is "always" as
// described above. This is only for use in secret sharing schemes. It is not
// a general polynomial multiplication routine.
func (p *PriPoly) Mul(q *PriPoly) *PriPoly { _ = "STUB: not implemented"; return nil }

// Coefficients return the list of coefficients representing p. This
// information is generally PRIVATE and should not be revealed to a third party
// lightly.
func (p *PriPoly) Coefficients() []kyber.Scalar {
	_ = "STUB: not implemented"

	// RecoverSecret reconstructs the shared secret p(0) from a list of private
	// shares using Lagrange interpolation.
	return nil
}

func RecoverSecret(g kyber.Group, shares []*PriShare, t, n uint32) (kyber.Scalar, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), nil
}

type byIndexScalar []*PriShare

func (s byIndexScalar) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s byIndexScalar) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s byIndexScalar) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// xyScalar returns the list of (x_i, y_i) pairs indexed. The first map returned
// is the list of x_i and the second map is the list of y_i, both indexed in
// their respective map at index i.
func xyScalar(g kyber.Group, shares []*PriShare, t, n uint32) (map[uint32]kyber.Scalar, map[uint32]kyber.Scalar) {
	_ = "STUB: not implemented"
	// we are sorting first the shares since the shares may be unrelated for
	// some applications. In this case, all participants needs to interpolate on
	// the exact same order shares.
	return nil, nil
}

func minusConst(g kyber.Group, c kyber.Scalar) *PriPoly { _ = "STUB: not implemented"; return nil }

// RecoverPriPoly takes a list of shares and the parameters t and n to
// reconstruct the secret polynomial completely, i.e., all private
// coefficients.  It is up to the caller to make sure that there are enough
// shares to correctly re-construct the polynomial. There must be at least t
// shares.
func RecoverPriPoly(g kyber.Group, shares []*PriShare, t, n uint32) (*PriPoly, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Notations follow the Wikipedia article on Lagrange interpolation
// https://en.wikipedia.org/wiki/Lagrange_polynomial

// add all L_j * y_j together

func (p *PriPoly) String() string { _ = "STUB: not implemented"; return "" }

// PubShare represents a public share.
type PubShare struct {
	I uint32      // Index of the public share
	V kyber.Point // Value of the public share
}

// Hash returns the hash representation of this share.
func (p *PubShare) Hash(s kyber.HashFactory) []byte { _ = "STUB: not implemented"; return nil }

// PubPoly represents a public commitment polynomial to a secret sharing polynomial.
type PubPoly struct {
	g       kyber.Group   // Cryptographic group
	b       kyber.Point   // Base point, nil for standard base
	commits []kyber.Point // Commitments to coefficients of the secret sharing polynomial
}

// NewPubPoly creates a new public commitment polynomial.
func NewPubPoly(g kyber.Group, b kyber.Point, commits []kyber.Point) *PubPoly {
	_ = "STUB: not implemented"
	return nil
}

// Info returns the base point and the commitments to the polynomial coefficients.
func (p *PubPoly) Info() (base kyber.Point, commits []kyber.Point) {
	_ = "STUB: not implemented"
	return *

	// Threshold returns the secret sharing threshold.
	new(kyber.Point), nil
}

func (p *PubPoly) Threshold() int64 { _ = "STUB: not implemented"; return 0 }

// Commit returns the secret commitment p(0), i.e., the constant term of the polynomial.
func (p *PubPoly) Commit() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// Eval computes the public share v = p(i).
func (p *PubPoly) Eval(i uint32) *PubShare { _ = "STUB: not implemented"; return nil }

// x-coordinate of this share

// Shares creates a list of n public commitment shares p(1),...,p(n).
func (p *PubPoly) Shares(n uint32) []*PubShare { _ = "STUB: not implemented"; return nil }

// Add computes the component-wise sum of the polynomials p and q and returns it
// as a new polynomial. NOTE: If the base points p.b and q.b are different then the
// base point of the resulting PubPoly cannot be computed without knowing the
// discrete logarithm between p.b and q.b. In this particular case, we are using
// p.b as a default value which of course does not correspond to the correct
// base point and thus should not be used in further computations.
func (p *PubPoly) Add(q *PubPoly) (*PubPoly, error) { _ = "STUB: not implemented"; return nil, nil }

// Equal checks equality of two public commitment polynomials p and q. If p and
// q are trivially unequal (e.g., due to mismatching cryptographic groups, or threshold issues),
// this routine returns in variable time. Otherwise it runs in constant time
// regardless of whether it eventually returns true or false.
func (p *PubPoly) Equal(q *PubPoly) bool { _ = "STUB: not implemented"; return false }

// Check a private share against a public commitment polynomial.
func (p *PubPoly) Check(s *PriShare) bool { _ = "STUB: not implemented"; return false }

type byIndexPub []*PubShare

func (s byIndexPub) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s byIndexPub) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s byIndexPub) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// xyCommits is the public version of xScalars.
func xyCommit(g kyber.Group, shares []*PubShare, t, n uint32) (map[uint32]kyber.Scalar, map[uint32]kyber.Point) {
	_ = "STUB: not implemented"
	// we are sorting first the shares since the shares may be unrelated for
	// some applications. In this case, all participants needs to interpolate on
	// the exact same order shares.
	return nil, nil
}

// RecoverCommit reconstructs the secret commitment p(0) from a list of public
// shares using Lagrange interpolation.
func RecoverCommit(g kyber.Group, shares []*PubShare, t, n uint32) (kyber.Point, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Point), nil
}

// RecoverPubPoly reconstructs the full public polynomial from a set of public
// shares using Lagrange interpolation.
func RecoverPubPoly(g kyber.Group, shares []*PubShare, t, n uint32) (*PubPoly, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// compute the L_j * y_j polynomial in point space

// add all L_j * y_j together

// lagrangeBasis returns a PriPoly containing the Lagrange coefficients for the
// i-th position. xs is a mapping between the indices and the values that the
// interpolation is using, computed with xyScalar().
func lagrangeBasis(g kyber.Group, i uint32, xs map[uint32]kyber.Scalar) *PriPoly {
	_ = "STUB: not implemented"
	return nil
}

// compute lagrange basis l_j

// Compute den = xi - xm
// Compute den = 1 / den
// Compute acc = acc * den

// multiply all coefficients by the denominator
