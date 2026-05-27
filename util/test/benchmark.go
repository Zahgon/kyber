package test

import (
	"testing"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/sign"
)

func PrepareBLS(numSigs int) (scheme sign.Scheme,
	publics []kyber.Point, privates []kyber.Scalar, msgs [][]byte, sigs [][]byte) {
	_ = "STUB: not implemented"
	return *new(sign.Scheme), nil, nil, nil, nil
}

func BenchCreateKeys(b *testing.B, scheme sign.Scheme, n int) { _ = "STUB: not implemented"; return }

func BenchSign(b *testing.B, scheme sign.Scheme, msg []byte, privates []kyber.Scalar) {
	_ = "STUB: not implemented"
	return
}

func BLSBenchVerify(b *testing.B, sigs [][]byte, scheme sign.Scheme,
	publics []kyber.Point, msgs [][]byte) {
	_ = "STUB: not implemented"
	return
}
