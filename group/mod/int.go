//go:build !constantTime

// Package mod contains a generic implementation of finite field arithmetic
// on integer fields with a constant modulus.
package mod

import (
	"crypto/cipher"
	"io"

	"go.dedis.ch/kyber/v4/compatible"
	"go.dedis.ch/kyber/v4/compatible/compatiblemod"

	"go.dedis.ch/kyber/v4"
)

var marshalScalarID = [8]byte{'m', 'o', 'd', '.', 'i', 'n', 't', ' '}

// Int is a generic implementation of finite field arithmetic
// on integer finite fields with a given constant modulus,
// built using Go's built-in big.Int or with the saferith package,
// depending on whether the constantTime build tag is chosen.
// Int satisfies the kyber.Scalar interface,
// and hence serves as a basic implementation of kyber.Scalar,
// e.g., representing discrete-log exponents of Schnorr groups
// or scalar multipliers for elliptic curves.
//
// Int offers an API similar to and compatible with big.Int,
// but "carries around" a pointer to the relevant modulus
// and automatically normalizes the value to that modulus
// after all arithmetic operations, simplifying modular arithmetic.
// Binary operations assume that the source(s)
// have the same modulus, but do not check this assumption.
// Unary and binary arithmetic operations may be performed on uninitialized
// target objects, and receive the modulus of the first operand.
// For efficiency the modulus field M is a pointer,
// whose target is assumed never to change.
type Int struct {
	V  compatible.Int     // Integer value from 0 through M-1
	M  *compatiblemod.Mod // Modulus for finite field arithmetic
	BO kyber.ByteOrder    // Endianness which will be used on input and output
}

// NewInt creaters a new Int with a given compatible.Int and a compatible.Mod modulus.
func NewInt(v *compatible.Int, m *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

// NewInt64 creates a new Int with a given int64 value and compatible.Mod modulus.
func NewInt64(v int64, m *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

// NewIntBytes creates a new Int with a given slice of bytes and a compatible.Mod
// modulus.
func NewIntBytes(a []byte, m *compatiblemod.Mod, byteOrder kyber.ByteOrder) *Int {
	_ = "STUB: not implemented"
	return nil
}

// NewIntString creates a new Int with a given string and a compatible.Mod modulus.
// The value is set to a rational fraction n/d in a given base.
func NewIntString(n, d string, base int, m *compatiblemod.Mod) *Int {
	_ = "STUB: not implemented"
	return nil
}

// Init a Int with a given compatible.Int value and modulus pointer.
// Note that the value is copied; the modulus is not.
func (i *Int) Init(v *compatible.Int, m *compatiblemod.Mod) *Int {
	_ = "STUB: not implemented"
	return nil
}

// Init64 creates an Int with an int64 value and compatible.Int modulus.
func (i *Int) Init64(v int64, m *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

// InitBytes init the Int to a number represented in a big-endian byte string.
func (i *Int) InitBytes(a []byte, m *compatiblemod.Mod, byteOrder kyber.ByteOrder) *Int {
	_ = "STUB: not implemented"
	return nil
}

// InitString inits the Int to a rational fraction n/d
// specified with a pair of strings in a given base.
func (i *Int) InitString(n, d string, base int, m *compatiblemod.Mod) *Int {
	_ = "STUB: not implemented"
	return nil
}

// Return the Int's integer value in hexadecimal string representation.
func (i *Int) String() string { _ = "STUB: not implemented"; return "" }

// SetString sets the Int to a rational fraction n/d represented by a pair of strings.
// If d == "", then the denominator is taken to be 1.
// Returns (i,true) on success, or
// (nil,false) if either string fails to parse.
func (i *Int) SetString(n, d string, base int) (*Int, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Cmp compares two Ints for equality or inequality
func (i *Int) Cmp(s2 kyber.Scalar) int { _ = "STUB: not implemented"; return 0 }

// s2 is not an Int

// Equal returns true if the two Ints are equal
func (i *Int) Equal(s2 kyber.Scalar) bool { _ = "STUB: not implemented"; return false }

// s2 is not an Int

// Nonzero returns true if the integer value is nonzero.
func (i *Int) Nonzero() bool { _ = "STUB: not implemented"; return false }

// Set both value and modulus to be equal to another Int.
// Since this method copies the modulus as well,
// it may be used as an alternative to Init().
func (i *Int) Set(a kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Scalar)
}

// Clone returns a separate duplicate of this Int.
func (i *Int) Clone() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

// Zero set the Int to the value 0.  The modulus must already be initialized.
func (i *Int) Zero() kyber.Scalar {
	_ = "STUB: not implemented"
	return *

	// One sets the Int to the value 1.  The modulus must already be initialized.
	new(kyber.Scalar)
}

func (i *Int) One() kyber.Scalar {
	_ = "STUB: not implemented"
	return *

	// SetInt64 sets the Int to an arbitrary 64-bit "small integer" value.
	// The modulus must already be initialized.
	new(kyber.Scalar)
}

func (i *Int) SetInt64(v int64) kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

// Int64 returns the int64 representation of the value.
// If the value is not representable in an int64 the result is undefined.
func (i *Int) Int64() int64 {
	_ = "STUB: not implemented"

	// SetUint64 sets the Int to an arbitrary uint64 value.
	// The modulus must already be initialized.
	return 0
}

func (i *Int) SetUint64(v uint64) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

// Uint64 returns the uint64 representation of the value.
// If the value is not representable in an uint64 the result is undefined.
func (i *Int) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

// Add sets the target to a + b mod M, where M is a's modulus..
func (i *Int) Add(a, b kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Scalar)
}

//nolint:errcheck // Design pattern to emulate generics

// Sub sets the target to a - b mod M.
// Target receives a's modulus.
func (i *Int) Sub(a, b kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Scalar)
}

//nolint:errcheck // Design pattern to emulate generics

// Neg sets the target to -a mod M.
func (i *Int) Neg(a kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Scalar)
}

// Mul sets the target to a * b mod M.
// Target receives a's modulus.
func (i *Int) Mul(a, b kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Scalar)
}

//nolint:errcheck // Design pattern to emulate generics

// Div sets the target to a * b^-1 mod M, where b^-1 is the modular inverse of b.
func (i *Int) Div(a, b kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Scalar)
}

//nolint:errcheck // Design pattern to emulate generics

// Inv sets the target to the modular inverse of a with respect to modulus M.
func (i *Int) Inv(a kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

// Exp sets the target to a^e mod M,
// where e is an arbitrary compatible.Int exponent (not necessarily 0 <= e < M).
func (i *Int) Exp(a kyber.Scalar, e *compatible.Int) kyber.Scalar {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Scalar)
}

// to protect against golang/go#22830

// Jacobi computes the Jacobi symbol of (a/M), which indicates whether a is
// zero (0), a positive square in M (1), or a non-square in M (-1).
func (i *Int) Jacobi(as kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Scalar)
}

// Sqrt computes some square root of a mod M of one exists.
// Assumes the modulus M is an odd prime.
// Returns true on success, false if input a is not a square.
func (i *Int) Sqrt(as kyber.Scalar) bool {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return false
}

// Pick a [pseudo-]random integer modulo M
// using bits from the given stream cipher.
func (i *Int) Pick(rand cipher.Stream) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

// ByteOrder return the byte representation type (big or little endian)
func (i *Int) ByteOrder() kyber.ByteOrder {
	_ = "STUB: not implemented"

	// GroupOrder returns the order of the underlying group
	return *new(kyber.ByteOrder)
}

func (i *Int) GroupOrder() *compatiblemod.Mod {
	_ = "STUB: not implemented"

	// MarshalSize returns the length in bytes of encoded integers with modulus M.
	// The length of encoded Ints depends only on the size of the modulus,
	// and not on the value of the encoded integer,
	// making the encoding is fixed-length for simplicity and security.
	return nil
}

func (i *Int) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

// MarshalBinary encodes the value of this Int into a byte-slice exactly Len() bytes long.
// It uses i's ByteOrder to determine which byte order to output.
func (i *Int) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// may be shorter than l

// MarshalID returns a unique identifier for this type
func (i *Int) MarshalID() [8]byte { _ = "STUB: not implemented"; return nil }

// UnmarshalBinary tries to decode a Int from a byte-slice buffer.
// Returns an error if the buffer is not exactly Len() bytes long
// or if the contents of the buffer represents an out-of-range integer.
func (i *Int) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

// Still needed here because of the comparison with the modulo

// MarshalTo encodes this Int to the given Writer.
func (i *Int) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// UnmarshalFrom tries to decode an Int from the given Reader.
func (i *Int) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// BigEndian encodes the value of this Int into a big-endian byte-slice
// at least min bytes but no more than max bytes long.
// Panics if max != 0 and the Int cannot be represented in max bytes.
func (i *Int) BigEndian(minBytes, maxBytes int) []byte { _ = "STUB: not implemented"; return nil }

// SetBytes set the value to a number represented
// by a byte string.
// Endianness depends on the endianess set in i.
func (i *Int) SetBytes(a []byte) kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

// LittleEndian encodes the value of this Int into a little-endian byte-slice
// at least min bytes but no more than max bytes long.
// Panics if max != 0 and the Int cannot be represented in max bytes.
func (i *Int) LittleEndian(minByte, maxBytes int) []byte { _ = "STUB: not implemented"; return nil }

// reverse copies src into dst in byte-reversed order and returns dst,
// such that src[0] goes into dst[len-1] and vice versa.
// dst and src may be the same slice but otherwise must not overlap.
func reverse(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }
