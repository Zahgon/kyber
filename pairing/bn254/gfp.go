//go:build !constantTime

package bn254

type gfP [4]uint64

func newGFp(x int64) (out *gfP) { _ = "STUB: not implemented"; return nil }

func newGFpFromBase10(x string) *gfP { _ = "STUB: not implemented"; return nil }

func (e *gfP) String() string { _ = "STUB: not implemented"; return "" }

func (e *gfP) Set(f *gfP) { _ = "STUB: not implemented"; return }

func (e *gfP) Invert(f *gfP) { _ = "STUB: not implemented"; return }

// Borrowed from: https://github.com/cloudflare/bn256/blob/master/gfp.go#L63
func (e *gfP) Exp(f *gfP, bits [4]uint64) { _ = "STUB: not implemented"; return }

// Borrowed from: https://github.com/cloudflare/bn256/blob/master/gfp.go#L85
func (e *gfP) Sqrt(f *gfP) {
	_ = "STUB: not implemented"
	// Since p = 4k+3, then e = f^(k+1) is a root of f.
	return
}

func (e *gfP) Marshal(out []byte) { _ = "STUB: not implemented"; return }

func (e *gfP) Unmarshal(in []byte) error {
	_ = "STUB: not implemented"
	// Unmarshal the bytes into little endian form
	return nil
}

// Ensure the point respects the curve modulus

func montEncode(c, a *gfP) { _ = "STUB: not implemented"; return }
func montDecode(c, a *gfP) { _ = "STUB: not implemented"; return }

// https://datatracker.ietf.org/doc/html/rfc9380/#name-the-sgn0-function
func sgn0(e *gfP) int { _ = "STUB: not implemented"; return 0 }

// Borrowed from: https://github.com/cloudflare/bn256/blob/master/gfp.go#L123
func legendre(e *gfP) int {
	_ = "STUB: not implemented"

	// Since p = 4k+3, then e^(2k+1) is the Legendre symbol of e.
	return 0
}
