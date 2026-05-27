// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build purego || !(386 || amd64 || arm || arm64 || loong64 || ppc64 || ppc64le || riscv64 || s390x || wasm)

package bigmod

func addMulVVW1024(z, x *uint, y uint) (c uint) { _ = "STUB: not implemented"; return 0 }

func addMulVVW1536(z, x *uint, y uint) (c uint) { _ = "STUB: not implemented"; return 0 }

func addMulVVW2048(z, x *uint, y uint) (c uint) { _ = "STUB: not implemented"; return 0 }
