//go:build constantTime

package compatible

import (
	"io"
	"math/big"

	"go.dedis.ch/kyber/v4/compatible/bigmod"
	"go.dedis.ch/kyber/v4/compatible/compatiblemod"
)

type Int struct {
	Int bigmod.Nat
}

func FromNat(x *bigmod.Nat) *Int { _ = "STUB: not implemented"; return nil }

func FromCompatibleMod(mod *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

func NewInt(x int64) *Int { _ = "STUB: not implemented"; return nil }

func NewUint(x uint64) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) Int64() int64 { _ = "STUB: not implemented"; return 0 }

func (z *Int) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

// SetString sets z to s mod m.
// This is variable time function. Only use if the size of s is public.
// The function requires to pass the modulus, which determines the
// announced length of the Nat
func (z *Int) SetString(s, m string, base int) (*Int, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SetStringM sets z to s mod m.
// This is variable time function. Only use if the size of s is public.
// The function requires to pass the modulus, which determines the
// announced length of the Nat
func (z *Int) SetStringM(s string, m *compatiblemod.Mod, base int) (*Int, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (z *Int) Bytes(m *compatiblemod.Mod) []byte { _ = "STUB: not implemented"; return nil }

// ModInverse sets z to the multiplicative inverse of g in the ring ℤ/nℤ
// Requires n to be prime (uses Fermat: g^(n-2) mod n).
func (z *Int) ModInverse(g *Int, n *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

// p = modulus as *Int
// exp = p - 2

// If g == 0 modulo n, no inverse

// copy of g

// iterate from most-significant bit down to zero

// square: res = res * res mod n

// multiply by base: res = res * base mod n

func (z *Int) ModInverseVartime(g *Int, n *compatiblemod.Mod) *Int {
	_ = "STUB: not implemented"
	return nil
}

// todo, normally fine if it's vartime (from its Kyber's usages (call it BitVartime)
func (z *Int) Bit(i int) uint { _ = "STUB: not implemented"; return 0 }

// copied from saferith.Nat
func (z *Int) FillBytes(buf []byte) []byte { _ = "STUB: not implemented"; return nil }

// LEAK: Number of limbs
// OK: The number of limbs is public
// LEAK: The addresses touched in the out array
// OK: Every member of out is touched

func (z *Int) Text(base int) string { _ = "STUB: not implemented"; return "" }

// one usage in rand.go, maybe can be replaced by big.Int directly
func (z *Int) BitLen() int { _ = "STUB: not implemented"; return 0 }

// vartime wrapper around crypto/rand
func Prime(rand io.Reader, bits int) (*Int, error) { _ = "STUB: not implemented"; return nil, nil }

func (z *Int) String() string { _ = "STUB: not implemented"; return "" }
func (z *Int) Exp(x, y *Int, m *compatiblemod.Mod) *Int {
	_ = "STUB: not implemented"
	// Exp requires y to be reduced modulo m
	return nil
}

func (z *Int) Equal(s2 *Int) bool { _ = "STUB: not implemented"; return false }

func (z *Int) Set(a *Int) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) SetUint(v uint) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) SetUint64(v uint64) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) Add(a, b *Int, mod *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) Sub(a, b *Int, mod *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

// Mul sets the receiver to the result of a * b mod m and returns the receiver
func (z *Int) Mul(a, b *Int, mod *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

// SetBytesMod sets the byte of this Int and then mods the result to the
// given modulus. Ensures that the resulting Int is less than the given
// modulus.
func (z *Int) SetBytesMod(buf []byte, mod *compatiblemod.Mod) *Int {
	_ = "STUB: not implemented"
	// To create the Nat that will be reduced, we need a modulus big enough for it
	// take max between the buffer and the mod byte size to avoid using
	// a mod smaller than the actual modulus
	return nil
}

// +1 to ensure the modulus is larger than the buffer value

// SetBytesWithCheck attempts to set the bytes of this int and returns an error if
// the byte value is larger or equal to the modulus.
// This method simply calls bigmod.Nat.SetBytes() and return an error
// is this method return an error
func (z *Int) SetBytesWithCheck(buf []byte, mod *compatiblemod.Mod) (*Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (z *Int) SetBytesBigBuffer(b []byte, m *compatiblemod.Mod) *Int {
	_ = "STUB: not implemented"
	return nil
}

// Mod computes x mod y, sets the receiver to this result and return
// the receiver
func (z *Int) Mod(x *Int, y *compatiblemod.Mod) *Int {
	_ = "STUB: not implemented"
	// Create a new Int and assign to z since bigmod.Nat.Mod() will overwrite the receiver
	return nil
}

func (z *Int) Sign() int { _ = "STUB: not implemented"; return 0 }

func (z *Int) IsZero() bool { _ = "STUB: not implemented"; return false }

func (z *Int) Cmp(x *Int) int { _ = "STUB: not implemented"; return 0 }

func (z *Int) Abs(x *Int) *Int {
	_ = "STUB: not implemented"

	// FromBigInt creates an Int from the given big.Int
	// this function is var-time
	return nil
}

func FromBigInt(z *big.Int, m *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

// ToBigInt returns this Int as a big.Int.
// this function is var-time
func (z *Int) ToBigInt() *big.Int { _ = "STUB: not implemented"; return nil }

// Create a modulo bigger than this int to be able to call SetBytes with it

// CmpGeqMod returns true if z >= mod otherwise 0
func (z *Int) CmpGeqMod(mod *compatiblemod.Mod) bool { _ = "STUB: not implemented"; return false }
