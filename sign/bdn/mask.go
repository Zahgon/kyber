package bdn

import (
	"go.dedis.ch/kyber/v4"
)

// Mask is a bitmask of the participation to a collective signature.
type Mask struct {
	// The bitmask indicating which public keys are enabled/disabled for aggregation. This is
	// the only mutable field.
	mask []byte

	// The following fields are immutable and should not be changed after the mask is created.
	// They may be shared between multiple masks.

	// Public keys for aggregation & signature verification.
	publics []kyber.Point
	// Coefficients used when aggregating signatures.
	publicCoefs []kyber.Scalar
	// Terms used to aggregate public keys
	publicTerms []kyber.Point
}

// NewMask creates a new mask from a list of public keys. If a key is provided, it
// will set the bit of the key to 1 or return an error if it is not found.
//
// The returned Mask will contain pre-computed terms and coefficients for all provided public
// keys, so it should be re-used for optimal performance (e.g., by creating a "base" mask and
// cloning it whenever aggregating signatures and/or public keys).
func NewMask(group kyber.Group, publics []kyber.Point, myKey kyber.Point) (*Mask, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mask returns the bitmask as a byte array.
func (m *Mask) Mask() []byte { _ = "STUB: not implemented"; return nil }

// Len returns the length of the byte array necessary to store the bitmask.
func (m *Mask) Len() int { _ = "STUB: not implemented"; return 0 }

// SetMask replaces the current mask by the new one if the length matches.
func (m *Mask) SetMask(mask []byte) error { _ = "STUB: not implemented"; return nil }

// GetBit returns true if the given bit is set.
func (m *Mask) GetBit(i int) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// SetBit turns on or off the bit at the given index.
func (m *Mask) SetBit(i int, enable bool) error { _ = "STUB: not implemented"; return nil }

// forEachBitEnabled is a helper to iterate over the bits set to 1 in the mask
// and to return the result of the callback only if it is positive.
func (m *Mask) forEachBitEnabled(f func(i, j, n int) int) int { _ = "STUB: not implemented"; return 0 }

// IndexOfNthEnabled returns the index of the nth enabled bit or -1 if out of bounds.
func (m *Mask) IndexOfNthEnabled(nth int) int { _ = "STUB: not implemented"; return 0 }

// NthEnabledAtIndex returns the sum of bits set to 1 until the given index. In other
// words, it returns how many bits are enabled before the given index.
func (m *Mask) NthEnabledAtIndex(idx int) int { _ = "STUB: not implemented"; return 0 }

// Publics returns a copy of the list of public keys.
func (m *Mask) Publics() []kyber.Point { _ = "STUB: not implemented"; return nil }

// Participants returns the list of public keys participating.
func (m *Mask) Participants() []kyber.Point { _ = "STUB: not implemented"; return nil }

// CountEnabled returns the number of bit set to 1
func (m *Mask) CountEnabled() int { _ = "STUB: not implemented"; return 0 }

// CountTotal returns the number of potential participants
func (m *Mask) CountTotal() int { _ = "STUB: not implemented"; return 0 }

// Merge merges the given mask to the current one only if
// the length matches
func (m *Mask) Merge(mask []byte) error { _ = "STUB: not implemented"; return nil }

// Clone copies the mask while keeping the precomputed coefficients, etc. This method is thread safe
// and does not modify the original mask. Modifications to the new Mask will not affect the original.
func (m *Mask) Clone() *Mask { _ = "STUB: not implemented"; return nil }
