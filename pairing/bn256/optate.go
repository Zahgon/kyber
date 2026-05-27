//go:build !constantTime

package bn256

func lineFunctionAdd(r, p *twistPoint, q *curvePoint, r2 *gfP2) (a, b, c *gfP2, rOut *twistPoint) {
	_ = "STUB: not implemented"
	// See the mixed addition algorithm from "Faster Computation of the
	// Tate Pairing", http://arxiv.org/pdf/0904.0854v3.pdf
	return nil, nil, nil, nil
}

func lineFunctionDouble(r *twistPoint, q *curvePoint) (a, b, c *gfP2, rOut *twistPoint) {
	_ = "STUB: not implemented"
	// See the doubling algorithm for a=0 from "Faster Computation of the
	// Tate Pairing", http://arxiv.org/pdf/0904.0854v3.pdf
	return nil, nil, nil, nil
}

func mulLine(ret *gfP12, a, b, c *gfP2) { _ = "STUB: not implemented"; return }

// sixuPlus2NAF is 6u+2 in non-adjacent form.
var sixuPlus2NAF = []int8{
	0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, -1, 0, 1, 0,
	1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, -1, 0, 1, 0, 0, 0, 1, 0, -1,
	0, 0, 0, -1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, -1, 0, -1, 0, 0, 0,
	0, 1, 0, 0, 0, 1,
}

// miller implements the Miller loop for calculating the Optimal Ate pairing.
// See algorithm 1 from http://cryptojedi.org/papers/dclxvi-20100714.pdf
func miller(q *twistPoint, p *curvePoint) *gfP12 { _ = "STUB: not implemented"; return nil }

// In order to calculate Q1 we have to convert q from the sextic twist
// to the full GF(p^12) group, apply the Frobenius there, and convert
// back.
//
// The twist isomorphism is (x', y') -> (xω², yω³). If we consider just
// x for a moment, then after applying the Frobenius, we have x̄ω^(2p)
// where x̄ is the conjugate of x. If we are going to apply the inverse
// isomorphism we need a value with a single coefficient of ω² so we
// rewrite this as x̄ω^(2p-2)ω². ξ⁶ = ω and, due to the construction of
// p, 2p-2 is a multiple of six. Therefore we can rewrite as
// x̄ξ^((p-1)/3)ω² and applying the inverse isomorphism eliminates the
// ω².
//
// A similar argument can be made for the y value.

// For Q2 we are applying the p² Frobenius. The two conjugations cancel
// out and we are left only with the factors from the isomorphism. In
// the case of x, we end up with a pure number which is why
// xiToPSquaredMinus1Over3 is ∈ GF(p). With y we get a factor of -1. We
// ignore this to end up with -Q2.

// finalExponentiation computes the (p¹²-1)/Order-th power of an element of
// GF(p¹²) to obtain an element of GT (steps 13-15 of algorithm 1 from
// http://cryptojedi.org/papers/dclxvi-20100714.pdf)
func finalExponentiation(in *gfP12) *gfP12 {
	_ = "STUB: not implemented"

	// This is the p^6-Frobenius
	return nil
}

func optimalAte(a *twistPoint, b *curvePoint) *gfP12 { _ = "STUB: not implemented"; return nil }
