//go:build !constantTime

package bn256

import (
	"math/big"
)

type gfP [4]uint64

func newGFp(x int64) (out *gfP) { _ = "STUB: not implemented"; return nil }

func newGFpFromBigInt(bigInt *big.Int) *gfP { _ = "STUB: not implemented"; return nil }

func hashToBase(msg, dst []byte) *gfP { _ = "STUB: not implemented"; return nil }

func (e *gfP) String() string { _ = "STUB: not implemented"; return "" }

func (e *gfP) Set(f *gfP) { _ = "STUB: not implemented"; return }

func (e *gfP) exp(f *gfP, bits [4]uint64) { _ = "STUB: not implemented"; return }

func (e *gfP) Invert(f *gfP) { _ = "STUB: not implemented"; return }

func (e *gfP) Sqrt(f *gfP) {
	_ = "STUB: not implemented"
	// Since p = 4k+3, then e = f^(k+1) is a root of f.
	return
}

func (e *gfP) Marshal(out []byte) { _ = "STUB: not implemented"; return }

func (e *gfP) Unmarshal(in []byte) { _ = "STUB: not implemented"; return }

func (e *gfP) BigInt() *big.Int { _ = "STUB: not implemented"; return nil }

func montEncode(c, a *gfP) { _ = "STUB: not implemented"; return }
func montDecode(c, a *gfP) { _ = "STUB: not implemented"; return }

func sign0(e *gfP) int { _ = "STUB: not implemented"; return 0 }

func legendre(e *gfP) int {
	_ = "STUB: not implemented"

	// Since p = 4k+3, then e^(2k+1) is the Legendre symbol of e.
	return 0
}
