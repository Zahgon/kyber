//go:build !constantTime

package kilic

import (
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/compatible/compatiblemod"
)

var curveOrder, _ = new(compatiblemod.Mod).SetString(
	"73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001", 16)

func NewScalar() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }
