// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

//go:generate go run . -out ../nat_amd64.s -pkg bigmod

func main() {
	Package("go.dedis.ch/kyber/v4/compatible/bigmod")
	ConstraintExpr("!purego")

	addMulVVW(1024)
	addMulVVW(1536)
	addMulVVW(2048)

	Generate()
}

func addMulVVW(bits int) { _ = "STUB: not implemented"; return }

// zero out carry

// implicit MULQ inputs and outputs

// The ADX strategy implements the following function, where c1 and c2 are
// the overflow and the carry flag respectively.
//
//    func addMulVVW(z, x []uint, y uint) (carry uint) {
//        var c1, c2 uint
//        for i := range z {
//            hi, lo := bits.Mul(x[i], y)
//            lo, c1 = bits.Add(lo, z[i], c1)
//            z[i], c2 = bits.Add(lo, carry, c2)
//            carry = hi
//        }
//        return carry + c1 + c2
//    }
//
// The loop is fully unrolled and the hi / carry registers are alternated
// instead of introducing a MOV.

// implicit source of MULXQ

// zero out carry

// unset flags and zero out z0
