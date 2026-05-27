//go:build !constantTime

package gnark

import (
	"crypto/cipher"
	"io"

	"go.dedis.ch/kyber/v4/compatible/compatiblemod"

	fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"go.dedis.ch/kyber/v4"
)

var _ kyber.Scalar = &Scalar{}

type Scalar struct{ inner fr.Element }

func (s *Scalar) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scalar) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s *Scalar) String() string { _ = "STUB: not implemented"; return "" }

func (s *Scalar) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (s *Scalar) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Scalar) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Scalar) Equal(s2 kyber.Scalar) bool { _ = "STUB: not implemented"; return false }

func (s *Scalar) Set(a kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) Clone() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

func (s *Scalar) SetInt64(v int64) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) Zero() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

func (s *Scalar) Add(a, b kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) Sub(a, b kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) Neg(a kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) One() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

func (s *Scalar) Mul(a, b kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) Div(a, b kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) Inv(a kyber.Scalar) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) Pick(stream cipher.Stream) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) SetBytes(data []byte) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}

func (s *Scalar) ByteOrder() kyber.ByteOrder {
	_ = "STUB: not implemented"
	return *new(kyber.ByteOrder)
}

func (s *Scalar) GroupOrder() *compatiblemod.Mod { _ = "STUB: not implemented"; return nil }
