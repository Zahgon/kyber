//go:build !constantTime

package compatiblemod

import (
	"math/big"
)

type Mod struct {
	big.Int
}

func NewInt(x int64) *Mod { _ = "STUB: not implemented"; return nil }

func NewUint(x uint64) *Mod { _ = "STUB: not implemented"; return nil }

func (z *Mod) SetString(s string, base int) (*Mod, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// NewModulusProduct creates a new modulus as the result of
// the multiplication of the two input byte arrays
func NewModulusProduct(a, b []byte) *Mod { _ = "STUB: not implemented"; return nil }

func FromString(s string, base int) (*Mod, error) { _ = "STUB: not implemented"; return nil, nil }

func (z *Mod) SetBytes(buf []byte) *Mod { _ = "STUB: not implemented"; return nil }

func (z *Mod) ToBigInt() *big.Int { _ = "STUB: not implemented"; return nil }

func FromBigInt(x *big.Int) *Mod { _ = "STUB: not implemented"; return nil }
