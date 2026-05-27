//go:build !constantTime

package bn254

import (
	"crypto/cipher"
	"io"

	"go.dedis.ch/kyber/v4"
)

var marshalPointID1 = [8]byte{'b', 'n', '2', '5', '4', '.', 'g', '1'}
var marshalPointID2 = [8]byte{'b', 'n', '2', '5', '4', '.', 'g', '2'}
var marshalPointIDT = [8]byte{'b', 'n', '2', '5', '4', '.', 'g', 't'}

type pointG1 struct {
	g   *curvePoint
	dst []byte
}

func newPointG1(dst []byte) *pointG1 { _ = "STUB: not implemented"; return nil }

func (p *pointG1) Equal(q kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (p *pointG1) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG1) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG1) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG1) Set(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Clone makes a deep copy of the point
func (p *pointG1) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG1) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG1) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	// XXX: An approach to implement this is:
	// - Encode data as the x-coordinate of a point on y²=x³+3 where len(data)
	//   is stored in the least significant byte of x and the rest is being
	//   filled with random values, i.e., x = rand || data || len(data).
	// - Use the Tonelli-Shanks algorithm to compute the y-coordinate.
	// - Convert the new point to Jacobian coordinates and set it as p.
	return *new(kyber.Point)
}

func (p *pointG1) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pointG1) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// p = a + b

func (p *pointG1) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG1) Neg(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG1) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG1) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	// Clone is required as we change the point
	return nil, nil
}

// Take a copy so that p is not written to, so calls to MarshalBinary
// are threadsafe.

func (p *pointG1) MarshalID() [8]byte { _ = "STUB: not implemented"; return nil }

func (p *pointG1) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointG1) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

// This is the point at infinity

func (p *pointG1) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointG1) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG1) ElementSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG1) String() string { _ = "STUB: not implemented"; return "" }

func (p *pointG1) Hash(m []byte) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func hashToPoint(domain, m []byte) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func hashToField(domain, m []byte) (*gfP, *gfP) { _ = "STUB: not implemented"; return nil, nil }

// `mapToPoint` implements the general Shallue-van de Woestijne mapping to BN254 G1
// RFC9380, 6.6.1. https://datatracker.ietf.org/doc/html/rfc9380#name-shallue-van-de-woestijne-me
func mapToPoint(domain []byte, u *gfP) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Compute tv1 = u^2 * g(Z)

// Compute tv2 = 1 + tv1

// Compute tv1 = 1 - tv1

// Compute tv3 = inv0(tv1 * tv2)

// Compute tv5 = u * tv1 * tv3 * tv4

// Compute x1 = -Z / 2 - tv5

// Compute x2 = -Z / 2 + tv5

// Compute tv7 = tv2^2

// Compute x3 = tv8^2
// Compute x3 = c4 * x3
// Compute x3 = 1 + x3

// `expandMsgXmdKeccak256` implements expand_message_xmd from IETF RFC9380 Sec 5.3.1
// Borrowed from: https://github.com/kilic/bls12-381/blob/master/hash_to_field.go
func expandMsgXmdKeccak256(domain, msg []byte, outLen int) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Compute DST_prime = DST || I2OSP(len(DST), 1)
// Compute b_0 = H(Z_pad || msg || l_i_b_str || I2OSP(0, 1) || DST_prime)

// Compute b_1 = H(b_0 || I2OSP(1, 1) || DST_prime)

// Compute b_i = H(strxor(b_0, b_(i - 1)) || I2OSP(i, 1) || DST_prime)

// Create b_i = H(strxor(b_0, b_(i - 1)) || I2OSP(i, 1) || DST_prime)

// b_1 || ... || b_(ell - 1)

// b_ell

type pointG2 struct {
	g   *twistPoint
	dst []byte
}

func newPointG2(dst []byte) *pointG2 { _ = "STUB: not implemented"; return nil }

func (p *pointG2) Equal(q kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (p *pointG2) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG2) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG2) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) Set(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Clone makes a deep copy of the point
func (p *pointG2) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG2) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG2) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pointG2) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// p = a + b

func (p *pointG2) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) Neg(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	// Clone is required as we change the point during the operation
	return nil, nil
}

func (p *pointG2) MarshalID() [8]byte { _ = "STUB: not implemented"; return nil }

func (p *pointG2) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointG2) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

// This is the point at infinity.

func (p *pointG2) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointG2) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG2) ElementSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG2) String() string { _ = "STUB: not implemented"; return "" }

type pointGT struct {
	g *gfP12
}

func newPointGT() *pointGT { _ = "STUB: not implemented"; return nil }

func (p *pointGT) Equal(q kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (p *pointGT) Null() kyber.Point {
	_ = "STUB: not implemented"
	// TODO: This can be a precomputed constant
	return *new(kyber.Point)
}

func (p *pointGT) Base() kyber.Point {
	_ = "STUB: not implemented"
	// TODO: This can be a precomputed constant
	return *new(kyber.Point)
}

func (p *pointGT) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Set(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Clone makes a deep copy of the point
func (p *pointGT) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointGT) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (p *pointGT) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pointGT) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Neg(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pointGT) MarshalID() [8]byte { _ = "STUB: not implemented"; return nil }

func (p *pointGT) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:funlen
func (p *pointGT) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

// TODO: check if point is on curve

func (p *pointGT) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointGT) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointGT) ElementSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointGT) String() string { _ = "STUB: not implemented"; return "" }

func (p *pointGT) Finalize() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointGT) Miller(p1, p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Pair(p1, p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}
