package circl

import (
	"crypto/cipher"
	"hash"
	"io"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
)

var _ pairing.Suite = Suite{}

type Suite struct{}

func NewSuite() (s Suite) { _ = "STUB: not implemented"; return *new(Suite) }

func (s Suite) String() string  { _ = "STUB: not implemented"; return "" }
func (s Suite) G1() kyber.Group { _ = "STUB: not implemented"; return *new(kyber.Group) }
func (s Suite) G2() kyber.Group { _ = "STUB: not implemented"; return *new(kyber.Group) }
func (s Suite) GT() kyber.Group { _ = "STUB: not implemented"; return *new(kyber.Group) }

func (s Suite) Pair(p1, p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (s Suite) ValidatePairing(p1, p2, p3, p4 kyber.Point) bool {
	_ = "STUB: not implemented"
	return false
}

func (s Suite) Read(_ io.Reader, _ ...any) error { _ = "STUB: not implemented"; return nil }

func (s Suite) Write(_ io.Writer, _ ...any) error { _ = "STUB: not implemented"; return nil }

func (s Suite) Hash() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func (s Suite) XOF(seed []byte) kyber.XOF { _ = "STUB: not implemented"; return *new(kyber.XOF) }

func (s Suite) RandomStream() cipher.Stream { _ = "STUB: not implemented"; return *new(cipher.Stream) }
