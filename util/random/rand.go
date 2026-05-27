// Package random provides facilities for generating
// random or pseudorandom cryptographic objects.
package random

import (
	"crypto/cipher"
	"io"

	"go.dedis.ch/kyber/v4/compatible"
	"go.dedis.ch/kyber/v4/compatible/compatiblemod"
)

// Bits chooses a uniform random BigInt with a given maximum BitLen.
// If 'exact' is true, choose a BigInt with _exactly_ that BitLen, not less
func Bits(bitlen uint, exact bool, rand cipher.Stream) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Int chooses a uniform random big.Int less than a given modulus
func Int(mod *compatiblemod.Mod, rand cipher.Stream) *compatible.Int {
	_ = "STUB: not implemented"
	return nil
}

// Bytes fills a slice with random bytes from rand.
func Bytes(b []byte, rand cipher.Stream) { _ = "STUB: not implemented"; return }

type randstream struct {
	Readers []io.Reader
}

func (r *randstream) XORKeyStream(dst, src []byte) { _ = "STUB: not implemented"; return }

// readerBytes is how many bytes we expect from each source

// try to read readerBytes bytes from all readers and write them in a buffer

// we are ok with few sources being insecure (i.e., providing less than
// readerBytes bytes), but not all of them

// create the XOF output, with hash of collected data as seed

// New returns a new cipher.Stream that gets random data from the given
// readers. If no reader was provided, Go's crypto/rand package is used.
// Otherwise, for each source, 32 bytes are read. They are concatenated and
// then hashed, and the resulting hash is used as a seed to a PRNG.
// The resulting cipher.Stream can be used in multiple threads.
func New(readers ...io.Reader) cipher.Stream { _ = "STUB: not implemented"; return *new(cipher.Stream) }
