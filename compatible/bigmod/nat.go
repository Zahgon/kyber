// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bigmod implements constant-time big integer arithmetic modulo large
// moduli. Unlike math/big, this package is suitable for implementing
// security-sensitive cryptographic operations. It is a re-exported version the
// standard library package crypto/internal/fips140/bigmod used to implement
// crypto/rsa amongst others.
//
// The API is NOT stable. The caller is responsible for ensuring that Nats are
// reduced modulo the Modulus they are used with.
package bigmod

import (
	"math/bits"
)

const (
	// _W is the size in bits of our limbs.
	_W = bits.UintSize
	// _S is the size in bytes of our limbs.
	_S = _W / 8
)

// Note: These functions make many loops over all the words in a Nat.
// These loops used to be in assembly, invisible to -race, -asan, and -msan,
// but now they are in Go and incur significant overhead in those modes.
// To bring the old performance back, we mark all functions that loop
// over Nat words with //go:norace. Because //go:norace does not
// propagate across inlining, we must also mark functions that inline
// //go:norace functions - specifically, those that inline add, addMulVVW,
// assign, cmpGeq, rshift1, and sub.

// choice represents a constant-time boolean. The value of choice is always
// either 1 or 0. We use an int instead of bool in order to make decisions in
// constant time by turning it into a mask.
type choice uint

func not(c choice) choice { _ = "STUB: not implemented"; return *new(choice) }

const yes = choice(1)
const no = choice(0)

// ctMask is all 1s if on is yes, and all 0s otherwise.
func ctMask(on choice) uint {
	_ = "STUB: not implemented"

	// ctEq returns 1 if x == y, and 0 otherwise. The execution time of this
	// function does not depend on its inputs.
	return 0
}

func ctEq(x, y uint) choice {
	_ = "STUB: not implemented"
	// If x != y, then either x - y or y - x will generate a carry.
	return *new(choice)
}

// Nat represents an arbitrary natural number
//
// Each Nat has an announced length, which is the number of limbs it has stored.
// Operations on this number are allowed to leak this length, but will not leak
// any information about the values contained in those limbs.
type Nat struct {
	// limbs is little-endian in base 2^W with W = bits.UintSize.
	limbs []uint
}

// preallocTarget is the size in bits of the numbers used to implement the most
// common and most performant RSA key size. It's also enough to cover some of
// the operations of key sizes up to 4096.
const preallocTarget = 2048
const preallocLimbs = (preallocTarget + _W - 1) / _W

// NewNat returns a new nat with a size of zero, just like new(Nat), but with
// the preallocated capacity to hold a number of up to 2048 bits.
// NewNat inlines, so the allocation can live on the stack.
func NewNat() *Nat { _ = "STUB: not implemented"; return nil }

// expand expands x to n limbs, leaving its value unchanged.
func (x *Nat) expand(n int) *Nat { _ = "STUB: not implemented"; return nil }

// reset returns a zero nat of n limbs, reusing x's storage if n <= cap(x.limbs).
func (x *Nat) reset(n int) *Nat { _ = "STUB: not implemented"; return nil }

// Clear both the returned limbs and the previously used ones.

// resetToBytes assigns x = b, where b is a slice of big-endian bytes, resizing
// n to the appropriate size.
//
// The announced length of x is set based on the actual bit size of the input,
// ignoring leading zeroes.
func (x *Nat) resetToBytes(b []byte) *Nat { _ = "STUB: not implemented"; return nil }

// trim reduces the size of x to match its value.
func (x *Nat) trim() *Nat {
	_ = "STUB: not implemented"
	// Trim most significant (trailing in little-endian) zero limbs.
	// We assume comparison with zero (but not the branch) is constant time.
	return nil
}

// set assigns x = y, optionally resizing x to the appropriate size.
func (x *Nat) set(y *Nat) *Nat { _ = "STUB: not implemented"; return nil }

// Bits returns x as a little-endian slice of uint. The length of the slice
// matches the announced length of x. The result and x share the same underlying
// array.
func (x *Nat) Bits() []uint {
	_ = "STUB: not implemented"

	// Bytes returns x as a zero-extended big-endian byte slice. The size of the
	// slice will match the size of m.
	//
	// x must have the same size as m and it must be less than or equal to m.
	return nil
}

func (x *Nat) Bytes(m *Modulus) []byte { _ = "STUB: not implemented"; return nil }

// SetBytes assigns x = b, where b is a slice of big-endian bytes.
// SetBytes returns an error if b >= m.
//
// The output will be resized to the size of m and overwritten.
//
//go:norace
func (x *Nat) SetBytes(b []byte, m *Modulus) (*Nat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetOverflowingBytes assigns x = b, where b is a slice of big-endian bytes.
// SetOverflowingBytes returns an error if b has a longer bit length than m, but
// reduces overflowing values up to 2^⌈log2(m)⌉ - 1.
//
// The output will be resized to the size of m and overwritten.
func (x *Nat) SetOverflowingBytes(b []byte, m *Modulus) (*Nat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setBytes would have returned an error if the input overflowed the limb
// size of the modulus, so now we only need to check if the most significant
// limb of x has more bits than the most significant limb of the modulus.

// bigEndianUint returns the contents of buf interpreted as a
// big-endian encoded uint value.
func bigEndianUint(buf []byte) uint { _ = "STUB: not implemented"; return 0 }

func (x *Nat) setBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

// SetUint assigns x = y.
//
// The output will be resized to a single limb and overwritten.
func (x *Nat) SetUint(y uint) *Nat { _ = "STUB: not implemented"; return nil }

// Equal returns 1 if x == y, and 0 otherwise.
//
// Both operands must have the same announced length.
//
//go:norace
func (x *Nat) Equal(y *Nat) uint {
	_ = "STUB: not implemented"
	// Eliminate bounds checks in the loop.
	return 0
}

// IsZero returns 1 if x == 0, and 0 otherwise.
//
//go:norace
func (x *Nat) IsZero() uint {
	_ = "STUB: not implemented"
	// Eliminate bounds checks in the loop.
	return 0
}

// IsOne returns 1 if x == 1, and 0 otherwise.
//
//go:norace
func (x *Nat) IsOne() uint {
	_ = "STUB: not implemented"
	// Eliminate bounds checks in the loop.
	return 0
}

// IsMinusOne returns 1 if x == -1 mod m, and 0 otherwise.
//
// The length of x must be the same as the modulus. x must already be reduced
// modulo m.
//
//go:norace
func (x *Nat) IsMinusOne(m *Modulus) uint { _ = "STUB: not implemented"; return 0 }

// IsOdd returns 1 if x is odd, and 0 otherwise.
func (x *Nat) IsOdd() uint { _ = "STUB: not implemented"; return 0 }

// TrailingZeroBitsVarTime returns the number of trailing zero bits in x.
func (x *Nat) TrailingZeroBitsVarTime() uint { _ = "STUB: not implemented"; return 0 }

// cmpGeq returns 1 if x >= y, and 0 otherwise.
//
// Both operands must have the same announced length.
//
//go:norace
func (x *Nat) cmpGeq(y *Nat) choice {
	_ = "STUB: not implemented"
	// Eliminate bounds checks in the loop.
	return *new(choice)
}

// If there was a carry, then subtracting y underflowed, so
// x is not greater than or equal to y.

// assign sets x <- y if on == 1, and does nothing otherwise.
//
// Both operands must have the same announced length.
//
//go:norace
func (x *Nat) assign(on choice, y *Nat) *Nat {
	_ = "STUB: not implemented"
	// Eliminate bounds checks in the loop.
	return nil
}

// add computes x += y and returns the carry.
//
// Both operands must have the same announced length.
//
//go:norace
func (x *Nat) add(y *Nat) (c uint) {
	_ = "STUB: not implemented"
	// Eliminate bounds checks in the loop.
	return 0
}

// sub computes x -= y. It returns the borrow of the subtraction.
//
// Both operands must have the same announced length.
//
//go:norace
func (x *Nat) sub(y *Nat) (c uint) {
	_ = "STUB: not implemented"
	// Eliminate bounds checks in the loop.
	return 0
}

// ShiftRightVarTime sets x = x >> n.
//
// The announced length of x is unchanged.
//
//go:norace
func (x *Nat) ShiftRightVarTime(n uint) *Nat {
	_ = "STUB: not implemented"
	// Eliminate bounds checks in the loop.
	return nil
}

// BitLenVarTime returns the actual size of x in bits.
//
// The actual size of x (but nothing more) leaks through timing side-channels.
// Note that this is ordinarily secret, as opposed to the announced size of x.
func (x *Nat) BitLenVarTime() int {
	_ = "STUB: not implemented"
	// Eliminate bounds checks in the loop.
	return 0
}

// bitLen is a version of bits.Len that only leaks the bit length of n, but not
// its value. bits.Len and bits.LeadingZeros use a lookup table for the
// low-order bits on some architectures.
func bitLen(n uint) int {
	_ = "STUB: not implemented"

	// We assume, here and elsewhere, that comparison to zero is constant time
	// with respect to different non-zero values.
	return 0
}

// Modulus is used for modular arithmetic, precomputing relevant constants.
//
// A Modulus can leak the exact number of bits needed to store its value
// and is stored without padding. Its actual value is still kept secret.
type Modulus struct {
	// The underlying natural number for this modulus.
	//
	// This will be stored without any padding, and shouldn't alias with any
	// other natural number being used.
	nat *Nat

	// If m is even, the following fields are not set.
	odd   bool
	m0inv uint // -nat.limbs[0]⁻¹ mod _W
	rr    *Nat // R*R for montgomeryRepresentation
}

// rr returns R*R with R = 2^(_W * n) and n = len(m.nat.limbs).
func rr(m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// We start by computing R = 2^(_W * n) mod m. We can get pretty close, to
// 2^⌊log₂m⌋, by setting the highest bit we can without having to reduce.

// Then we double until we reach 2^(_W * n).

// Next we need to get from R to 2^(_W * n) R mod m (aka from one to R in
// the Montgomery domain, meaning we can use Montgomery multiplication now).
// We could do that by doubling _W * n times, or with a square-and-double
// chain log2(_W * n) long. Turns out the fastest thing is to start out with
// doublings, and switch to square-and-double once the exponent is large
// enough to justify the cost of the multiplications.

// The threshold is selected experimentally as a linear function of n.

// We calculate how many of the most-significant bits of the exponent we can
// compute before crossing the threshold, and we do it with doublings.

// Then we process the remaining bits of the exponent with a
// square-and-double chain.

// minusInverseModW computes -x⁻¹ mod _W with x odd.
//
// This operation is used to precompute a constant involved in Montgomery
// multiplication.
func minusInverseModW(x uint) uint {
	_ = "STUB: not implemented"
	// Every iteration of this loop doubles the least-significant bits of
	// correct inverse in y. The first three bits are already correct (1⁻¹ = 1,
	// 3⁻¹ = 3, 5⁻¹ = 5, and 7⁻¹ = 7 mod 8), so doubling five times is enough
	// for 64 bits (and wastes only one iteration for 32 bits).
	//
	// See https://crypto.stackexchange.com/a/47496.
	return 0
}

// NewModulus creates a new Modulus from a slice of big-endian bytes. The
// modulus must be greater than one.
//
// The number of significant bits and whether the modulus is even is leaked
// through timing side-channels.
func NewModulus(b []byte) (*Modulus, error) { _ = "STUB: not implemented"; return nil, nil }

// NewModulusProduct creates a new Modulus from the product of two numbers
// represented as big-endian byte slices. The result must be greater than one.
//
//go:norace
func NewModulusProduct(a, b []byte) (*Modulus, error) { _ = "STUB: not implemented"; return nil, nil }

func newModulus(n *Nat) (*Modulus, error) { _ = "STUB: not implemented"; return nil, nil }

// Size returns the size of m in bytes.
func (m *Modulus) Size() int { _ = "STUB: not implemented"; return 0 }

// BitLen returns the size of m in bits.
func (m *Modulus) BitLen() int { _ = "STUB: not implemented"; return 0 }

// Nat returns m as a Nat.
func (m *Modulus) Nat() *Nat {
	_ = "STUB: not implemented"
	// Make a copy so that the caller can't modify m.nat or alias it with
	// another Nat in a modulus operation.
	return nil
}

// shiftIn calculates x = x << _W + y mod m.
//
// This assumes that x is already reduced mod m.
//
//go:norace
func (x *Nat) shiftIn(y uint, m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// Eliminate bounds checks in the loop.

// Each iteration of this loop computes x = 2x + b mod m, where b is a bit
// from y. Effectively, it left-shifts x and adds y one bit at a time,
// reducing it every time.
//
// To do the reduction, each iteration computes both 2x + b and 2x + b - m.
// The next iteration (and finally the return line) will use either result
// based on whether 2x + b overflows m.

// Like in maybeSubtractModulus, we need the subtraction if either it
// didn't underflow (meaning 2x + b > m) or if computing 2x + b
// overflowed (meaning 2x + b > 2^_W*n > m).

// Mod calculates out = y mod m.
//
// This works regardless how large the value of y is.
//
// The output will be resized to the size of m and overwritten.
//
//go:norace
func (x *Nat) Mod(y *Nat, m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// Working our way from the most significant to the least significant limb,
// we can insert each limb at the least significant position, shifting all
// previous limbs left by _W. This way each limb will get shifted by the
// correct number of bits. We can insert at least N - 1 limbs without
// overflowing m. After that, we need to reduce every time we shift.

// For the first N - 1 limbs we can skip the actual shifting and position
// them at the shifted position, which starts at min(N - 2, i).

// We shift in the remaining limbs, reducing modulo m each time.

// ExpandFor ensures x has the right size to work with operations modulo m.
//
// The announced size of x must be smaller than or equal to that of m.
func (x *Nat) ExpandFor(m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// resetFor ensures x has the right size to work with operations modulo m.
//
// x is zeroed and may start at any size.
func (x *Nat) resetFor(m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// maybeSubtractModulus computes x -= m if and only if x >= m or if "always" is yes.
//
// It can be used to reduce modulo m a value up to 2m - 1, which is a common
// range for results computed by higher level operations.
//
// always is usually a carry that indicates that the operation that produced x
// overflowed its size, meaning abstractly x > 2^_W*n > m even if x < m.
//
// x and m operands must have the same announced length.
//
//go:norace
func (x *Nat) maybeSubtractModulus(always choice, m *Modulus) { _ = "STUB: not implemented"; return }

// We keep the result if x - m didn't underflow (meaning x >= m)
// or if always was set.

// Sub computes x = x - y mod m.
//
// The length of both operands must be the same as the modulus. Both operands
// must already be reduced modulo m.
//
//go:norace
func (x *Nat) Sub(y *Nat, m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// If the subtraction underflowed, add m.

// SubOne computes x = x - 1 mod m.
//
// The length of x must be the same as the modulus.
func (x *Nat) SubOne(m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// Sub asks for x to be reduced modulo m, while SubOne doesn't, but when
// y = 1, it works, and this is an internal use.

// Add computes x = x + y mod m.
//
// The length of both operands must be the same as the modulus. Both operands
// must already be reduced modulo m.
//
//go:norace
func (x *Nat) Add(y *Nat, m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// montgomeryRepresentation calculates x = x * R mod m, with R = 2^(_W * n) and
// n = len(m.nat.limbs).
//
// Faster Montgomery multiplication replaces standard modular multiplication for
// numbers in this representation.
//
// This assumes that x is already reduced mod m.
func (x *Nat) montgomeryRepresentation(m *Modulus) *Nat {
	_ = "STUB: not implemented"
	// A Montgomery multiplication (which computes a * b / R) by R * R works out
	// to a multiplication by R, which takes the value out of the Montgomery domain.
	return nil
}

// montgomeryReduction calculates x = x / R mod m, with R = 2^(_W * n) and
// n = len(m.nat.limbs).
//
// This assumes that x is already reduced mod m.
func (x *Nat) montgomeryReduction(m *Modulus) *Nat {
	_ = "STUB: not implemented"
	// By Montgomery multiplying with 1 not in Montgomery representation, we
	// convert out back from Montgomery representation, because it works out to
	// dividing by R.
	return nil
}

// montgomeryMul calculates x = a * b / R mod m, with R = 2^(_W * n) and
// n = len(m.nat.limbs), also known as a Montgomery multiplication.
//
// All inputs should be the same length and already reduced modulo m.
// x will be resized to the size of m and overwritten.
//
//go:norace
func (x *Nat) montgomeryMul(a *Nat, b *Nat, m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// Attempt to use a stack-allocated backing array.

// This loop implements Word-by-Word Montgomery Multiplication, as
// described in Algorithm 4 (Fig. 3) of "Efficient Software
// Implementations of Modular Exponentiation" by Shay Gueron
// [https://eprint.iacr.org/2011/239.pdf].

// bounds check elimination hint

// Step 1 (T = a × b) is computed as a large pen-and-paper column
// multiplication of two numbers with n base-2^_W digits. If we just
// wanted to produce 2n-wide T, we would do
//
//   for i := 0; i < n; i++ {
//       d := bLimbs[i]
//       T[n+i] = addMulVVW(T[i:n+i], aLimbs, d)
//   }
//
// where d is a digit of the multiplier, T[i:n+i] is the shifted
// position of the product of that digit, and T[n+i] is the final carry.
// Note that T[i] isn't modified after processing the i-th digit.
//
// Instead of running two loops, one for Step 1 and one for Steps 2–6,
// the result of Step 1 is computed during the next loop. This is
// possible because each iteration only uses T[i] in Step 2 and then
// discards it in Step 6.

// Step 6 is replaced by shifting the virtual window we operate
// over: T of the algorithm is T[i:] for us. That means that T1 in
// Step 2 (T mod 2^_W) is simply T[i]. k0 in Step 3 is our m0inv.

// Step 4 and 5 add Y × m to T, which as mentioned above is stored
// at T[i:]. The two carries (from a × d and Y × m) are added up in
// the next word T[n+i], and the carry bit from that addition is
// brought forward to the next iteration.

// Finally for Step 7 we copy the final T window into x, and subtract m
// if necessary (which as explained in maybeSubtractModulus can be the
// case both if x >= m, or if x overflowed).
//
// The paper suggests in Section 4 that we can do an "Almost Montgomery
// Multiplication" by subtracting only in the overflow case, but the
// cost is very similar since the constant time subtraction tells us if
// x >= m as a side effect, and taking care of the broken invariant is
// highly undesirable (see https://go.dev/issue/13907).

// The following specialized cases follow the exact same algorithm, but
// optimized for the sizes most used in RSA. addMulVVW is implemented in
// assembly with loop unrolling depending on the architecture and bounds
// checks are removed by the compiler thanks to the constant size.

// compiler hint

// compiler hint

// compiler hint

// addMulVVW multiplies the multi-word value x by the single-word value y,
// adding the result to the multi-word value z and returning the final carry.
// It can be thought of as one row of a pen-and-paper column multiplication.
//
//go:norace
func addMulVVW(z, x []uint, y uint) (carry uint) {
	_ = "STUB: not implemented"
	// bounds check elimination hint
	return 0
}

// We use bits.Add with zero to get an add-with-carry instruction that
// absorbs the carry from the previous bits.Add.

// Mul calculates x = x * y mod m.
//
// The length of both operands must be the same as the modulus. Both operands
// must already be reduced modulo m.
//
//go:norace
func (x *Nat) Mul(y *Nat, m *Modulus) *Nat {
	_ = "STUB: not implemented"

	// A Montgomery multiplication by a value out of the Montgomery domain
	// takes the result out of Montgomery representation.
	return nil
}

// xR = x * R mod m
// x = xR * y / R mod m

// Attempt to use a stack-allocated backing array.

// T = x * y

// x = T mod m

// The following specialized cases follow the exact same algorithm, but
// optimized for the sizes most used in RSA. See montgomeryMul for details.

// compiler hint

// compiler hint

// compiler hint

// Exp calculates x = y^e mod m.
//
// The exponent e is represented in big-endian order. The output will be resized
// to the size of m and overwritten. y must already be reduced modulo m.
//
// m must be odd, or Exp will panic.
//
//go:norace
func (x *Nat) Exp(y *Nat, e []byte, m *Modulus) *Nat { _ = "STUB: not implemented"; return nil }

// We use a 4 bit window. For our RSA workload, 4 bit windows are faster
// than 2 bit windows, but use an extra 12 nats worth of scratch space.
// Using bit sizes that don't divide 8 are more complex to implement, but
// are likely to be more efficient if necessary.

// table[i] = x ^ (i+1)
// newNat calls are unrolled so they are allocated on the stack.

// Square four times. Optimization note: this can be implemented
// more efficiently than with generic Montgomery multiplication.

// Select x^k in constant time from the table.

// Multiply by x^k, discarding the result if k = 0.

// ExpShortVarTime calculates out = x^e mod m.
//
// The output will be resized to the size of m and overwritten. x must already
// be reduced modulo m. This leaks the exponent through timing side-channels.
//
// m must be odd, or ExpShortVarTime will panic.
func (x *Nat) ExpShortVarTime(y *Nat, e uint, m *Modulus) *Nat {
	_ = "STUB: not implemented"
	return nil
}

// For short exponents, precomputing a table and using a window like in Exp
// doesn't pay off. Instead, we do a simple conditional square-and-multiply
// chain, skipping the initial run of zeroes.

// InverseVarTime calculates x = a⁻¹ mod m and returns (x, true) if a is
// invertible. Otherwise, InverseVarTime returns (x, false) and x is not
// modified.
//
// a must be reduced modulo m, but doesn't need to have the same size. The
// output will be resized to the size of m and overwritten.
//
//go:norace
func (x *Nat) InverseVarTime(a *Nat, m *Modulus) (*Nat, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GCDVarTime calculates x = GCD(a, b) where at least one of a or b is odd, and
// both are non-zero. If GCDVarTime returns an error, x is not modified.
//
// The output will be resized to the size of the larger of a and b.
func (x *Nat) GCDVarTime(a, b *Nat) (*Nat, error) { _ = "STUB: not implemented"; return nil, nil }

// extendedGCD computes u and A such that a = GCD(a, m) and u = A*a - B*m.
//
// u will have the size of the larger of a and m, and A will have the size of m.
//
// It is an error if either a or m is zero, or if they are both even.
func extendedGCD(a, m *Nat) (u, A *Nat, err error) {
	_ = "STUB: not implemented"
	// This is the extended binary GCD algorithm described in the Handbook of
	// Applied Cryptography, Algorithm 14.61, adapted by BoringSSL to bound
	// coefficients and avoid negative numbers. For more details and proof of
	// correctness, see https://github.com/mit-plv/fiat-crypto/pull/333/files.
	//
	// Following the proof linked in the PR above, the changes are:
	//
	//  1. Negate [B] and [C] so they are positive. The invariant now involves a
	//     subtraction.
	//  2. If step 2 (both [x] and [y] are even) runs, abort immediately. This
	//     case needs to be handled by the caller.
	//  3. Subtract copies of [x] and [y] as needed in step 6 (both [u] and [v]
	//     are odd) so coefficients stay in bounds.
	//  4. Replace the [u >= v] check with [u > v]. This changes the end
	//     condition to [v = 0] rather than [u = 0]. This saves an extra
	//     subtraction due to which coefficients were negated.
	//  5. Rename x and y to a and n, to capture that one is a modulus.
	//  6. Rearrange steps 4 through 6 slightly. Merge the loops in steps 4 and
	//     5 into the main loop (step 7's goto), and move step 6 to the start of
	//     the loop iteration, ensuring each loop iteration halves at least one
	//     value.
	//
	// Note this algorithm does not handle either input being zero.
	return nil, nil, nil
}

// Before and after each loop iteration, the following hold:
//
//   u = A*a - B*m
//   v = D*m - C*a
//   0 < u <= a
//   0 <= v <= m
//   0 <= A < m
//   0 <= B <= a
//   0 <= C < m
//   0 <= D <= a
//
// After each loop iteration, u and v only get smaller, and at least one of
// them shrinks by at least a factor of two.

// If both u and v are odd, subtract the smaller from the larger.
// If u = v, we need to subtract from v to hit the modified exit condition.

// Exactly one of u and v is now even.

// Halve the even one and adjust the corresponding coefficient.

// v.IsOdd() == 0

//go:norace
func rshift1(a *Nat, carry uint) { _ = "STUB: not implemented"; return }

// DivShortVarTime calculates x = x / y and returns the remainder.
//
// It panics if y is zero.
//
//go:norace
func (x *Nat) DivShortVarTime(y uint) uint { _ = "STUB: not implemented"; return 0 }
