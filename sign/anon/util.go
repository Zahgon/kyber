package anon

import (
	"go.dedis.ch/kyber/v4"
)

func BenchGenSig(suite Suite, nkeys int, benchMessage []byte, benchPub []kyber.Point, benchPri kyber.Scalar) []byte {
	_ = "STUB: not implemented"
	return nil
}

func BenchGenKeys(g kyber.Group,
	nkeys int) ([]kyber.Point, kyber.Scalar) {
	_ = "STUB: not implemented"
	return nil,

		// Create an anonymity set of random "public keys"
		*new(kyber.Scalar)
}

// pick random points

// Make just one of them an actual public/private keypair (X[mine],x)

func BenchSign(suite Suite, pub []kyber.Point, pri kyber.Scalar,
	niter int, benchMessage []byte) {
	_ = "STUB: not implemented"
	return
}

func BenchVerify(suite Suite, pub []kyber.Point,
	sig []byte, niter int, benchMessage []byte) {
	_ = "STUB: not implemented"
	return
}
