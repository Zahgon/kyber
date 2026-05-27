//go:build constantTime

package compatiblemod

import (
	"math/big"

	"go.dedis.ch/kyber/v4/compatible/bigmod"
)

type Mod struct {
	bigmod.Modulus
}

func (m *Mod) Nat() *bigmod.Nat { _ = "STUB: not implemented"; return nil }

// SetString sets this receiver to the value of s in the given
// base. This methods leaks the size of s.
func (m *Mod) SetString(s string, base int) (*Mod, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func FromString(s string, base int) (*Mod, bool) { _ = "STUB: not implemented"; return nil, false }

func (m *Mod) SetBytes(b []byte) *Mod { _ = "STUB: not implemented"; return nil }

// one usage in group/edwards22519/point_test.go @ TestPointIsCanonical
func (m *Mod) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func NewInt(x int64) *Mod { _ = "STUB: not implemented"; return nil }

func NewUint(x uint64) *Mod { _ = "STUB: not implemented"; return nil }

// NewModulusProduct creates a new modulus as the result of
// the multiplication of the two input byte arrays
func NewModulusProduct(a, b []byte) *Mod { _ = "STUB: not implemented"; return nil }

func FromBigInt(x *big.Int) *Mod { _ = "STUB: not implemented"; return nil }

func (m *Mod) ToBigInt() *big.Int { _ = "STUB: not implemented"; return nil }

func (m *Mod) SetBigInt(big *big.Int) *Mod { _ = "STUB: not implemented"; return nil }

func (m *Mod) Bit(i int) uint { _ = "STUB: not implemented"; return 0 }

func (m *Mod) String() string { _ = "STUB: not implemented"; return "" }
