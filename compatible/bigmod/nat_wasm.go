// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

package bigmod

// The generic implementation relies on 64x64->128 bit multiplication and
// 64-bit add-with-carry, which are compiler intrinsics on many architectures.
// Wasm doesn't support those. Here we implement it with 32x32->64 bit
// operations, which is more efficient on Wasm.

func idx(x *uint, i uintptr) *uint { _ = "STUB: not implemented"; return nil }

func addMulVVWWasm(z, x *uint, y uint, n uintptr) (carry uint) { _ = "STUB: not implemented"; return 0 }

func addMulVVW1024(z, x *uint, y uint) (c uint) { _ = "STUB: not implemented"; return 0 }

func addMulVVW1536(z, x *uint, y uint) (c uint) { _ = "STUB: not implemented"; return 0 }

func addMulVVW2048(z, x *uint, y uint) (c uint) { _ = "STUB: not implemented"; return 0 }
