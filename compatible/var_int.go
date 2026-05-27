//go:build !constantTime

package compatible

import (
	"io"
	"math/big"

	"go.dedis.ch/kyber/v4/compatible/compatiblemod"
)

type Int struct {
	big.Int
}

func NewInt(x int64) *Int { _ = "STUB: not implemented"; return nil }

func NewUint(x uint64) *Int { _ = "STUB: not implemented"; return nil }

func Jacobi(x, y *Int) int { _ = "STUB: not implemented"; return 0 }

func Prime(randR io.Reader, bits int) (*Int, error) { _ = "STUB: not implemented"; return nil, nil }

func (z *Int) ToCompatibleMod() *compatiblemod.Mod { _ = "STUB: not implemented"; return nil }

func FromCompatibleMod(mod *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

func FromBigInt(z *big.Int, _ *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) ToBigInt() *big.Int {
	_ = "STUB: not implemented"

	// SetString sets this receiver to s. Return this receiver
	// and a boolean indicating success
	return nil
}

func (z *Int) SetString(s, _ string, base int) (*Int, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SetStringM sets this receiver to s mod m. Returns this receiver
// and a boolean indicating success.
func (z *Int) SetStringM(s string, m *compatiblemod.Mod, base int) (*Int, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (z *Int) Mul(a, b *Int, mod *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) Sub(a, b *Int, mod *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) Add(a, b *Int, mod *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) SetUint64(x uint64) *Int { _ = "STUB: not implemented"; return nil }

// Mod computes x mod y, sets the receiver to this result and return
// the receiver
func (z *Int) Mod(x *Int, y *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

// SetBytesMod sets the byte of this Int and then mods the result to the
// given modulus. Ensures that the resulting Int is less than the given
// modulus.
func (z *Int) SetBytesMod(buf []byte, mod *compatiblemod.Mod) *Int {
	_ = "STUB: not implemented"
	return nil
}

func (z *Int) SetBytesWithCheck(buf []byte, mod *compatiblemod.Mod) (*Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (z *Int) Cmp(y *Int) (r int) { _ = "STUB: not implemented"; return 0 }

func (z *Int) Exp(x, y *Int, m *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) ModInverse(g *Int, n *compatiblemod.Mod) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) SetInt64(x int64) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) Set(x *Int) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) SetBit(x *Int, i int, b uint) *Int { _ = "STUB: not implemented"; return nil }

func (z *Int) Bytes(_ *compatiblemod.Mod) []byte { _ = "STUB: not implemented"; return nil }

// CmpGeqMod returns true if z >= mod otherwise 0
func (z *Int) CmpGeqMod(mod *compatiblemod.Mod) bool { _ = "STUB: not implemented"; return false }
