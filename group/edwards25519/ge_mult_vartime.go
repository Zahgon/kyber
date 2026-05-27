package edwards25519

// geScalarMultVartime computes h = a*B, where
//
//	a = a[0]+256*a[1]+...+256^31 a[31]
//	B is the Ed25519 base point (x,4/5) with x positive.
//
// Preconditions:
//
//	a[31] <= 127
func geScalarMultVartime(h *extendedGroupElement, a *[32]byte,
	A *extendedGroupElement) {
	_ = "STUB: not implemented"
	return
}

// A,3A,5A,7A,9A,11A,13A,15A

// Slide through the scalar exponent clumping sequences of bits,
// resulting in only zero or odd multipliers between -15 and 15.

// Form an array of odd multiples of A from 1A through 15A,
// in addition-ready cached group element form.
// We only need odd multiples of A because slide()
// produces only odd-multiple clumps of bits.

// Process the multiplications from most-significant bit downward

// no bits set

// first (most-significant) nonzero clump of bits

// remaining bits
