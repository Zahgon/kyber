package test

import (
	"go.dedis.ch/kyber/v4"
)

// GroupBench is a generic benchmark suite for kyber.groups.
type GroupBench struct {
	g kyber.Group

	// Random secrets and points for testing
	x, y kyber.Scalar
	X, Y kyber.Point
	xe   []byte // encoded Scalar
	Xe   []byte // encoded Point
}

// NewGroupBench returns a new GroupBench.
func NewGroupBench(g kyber.Group) *GroupBench { _ = "STUB: not implemented"; return nil }

// ScalarAdd benchmarks the addition operation for scalars
func (gb GroupBench) ScalarAdd(iters int) { _ = "STUB: not implemented"; return }

// ScalarSub benchmarks the subtraction operation for scalars
func (gb GroupBench) ScalarSub(iters int) { _ = "STUB: not implemented"; return }

// ScalarNeg benchmarks the negation operation for scalars
func (gb GroupBench) ScalarNeg(iters int) { _ = "STUB: not implemented"; return }

// ScalarMul benchmarks the multiplication operation for scalars
func (gb GroupBench) ScalarMul(iters int) { _ = "STUB: not implemented"; return }

// ScalarDiv benchmarks the division operation for scalars
func (gb GroupBench) ScalarDiv(iters int) { _ = "STUB: not implemented"; return }

// ScalarInv benchmarks the inverse operation for scalars
func (gb GroupBench) ScalarInv(iters int) { _ = "STUB: not implemented"; return }

// ScalarPick benchmarks the Pick operation for scalars
func (gb GroupBench) ScalarPick(iters int) { _ = "STUB: not implemented"; return }

// ScalarEncode benchmarks the marshalling operation for scalars
func (gb GroupBench) ScalarEncode(iters int) { _ = "STUB: not implemented"; return }

// ScalarDecode benchmarks the unmarshalling operation for scalars
func (gb GroupBench) ScalarDecode(iters int) { _ = "STUB: not implemented"; return }

// PointAdd benchmarks the addition operation for points
func (gb GroupBench) PointAdd(iters int) { _ = "STUB: not implemented"; return }

// PointSub benchmarks the subtraction operation for points
func (gb GroupBench) PointSub(iters int) { _ = "STUB: not implemented"; return }

// PointNeg benchmarks the negation operation for points
func (gb GroupBench) PointNeg(iters int) { _ = "STUB: not implemented"; return }

// PointMul benchmarks the multiplication operation for points
func (gb GroupBench) PointMul(iters int) { _ = "STUB: not implemented"; return }

// PointBaseMul benchmarks the base multiplication operation for points
func (gb GroupBench) PointBaseMul(iters int) { _ = "STUB: not implemented"; return }

// PointPick benchmarks the pick-ing operation for points
func (gb GroupBench) PointPick(iters int) { _ = "STUB: not implemented"; return }

// PointEncode benchmarks the encoding operation for points
func (gb GroupBench) PointEncode(iters int) { _ = "STUB: not implemented"; return }

// PointDecode benchmarks the decoding operation for points
func (gb GroupBench) PointDecode(iters int) { _ = "STUB: not implemented"; return }
