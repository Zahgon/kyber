//go:build ((!amd64 && !arm64) || generic) && !constantTime

package bn254

func gfpCarry(a *gfP, head uint64) { _ = "STUB: not implemented"; return }

// If b is negative, then return a.
// Else return b.

func gfpNeg(c, a *gfP) { _ = "STUB: not implemented"; return }

func gfpAdd(c, a, b *gfP) { _ = "STUB: not implemented"; return }

func gfpSub(c, a, b *gfP) { _ = "STUB: not implemented"; return }

func mul(a, b [4]uint64) [8]uint64 { _ = "STUB: not implemented"; return nil }

func halfMul(a, b [4]uint64) [4]uint64 { _ = "STUB: not implemented"; return nil }

func gfpMul(c, a, b *gfP) { _ = "STUB: not implemented"; return }
