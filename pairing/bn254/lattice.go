//go:build !constantTime

package bn254

import (
	"math/big"
)

var half = new(big.Int).Rsh(Order, 1)

var curveLattice = &lattice{
	vectors: [][]*big.Int{
		{bigFromBase10("147946756881789319000765030803803410728"), bigFromBase10("147946756881789319010696353538189108491")},
		{bigFromBase10("147946756881789319020627676272574806254"), bigFromBase10("-147946756881789318990833708069417712965")},
	},
	inverse: []*big.Int{
		bigFromBase10("147946756881789318990833708069417712965"),
		bigFromBase10("147946756881789319010696353538189108491"),
	},
	det: bigFromBase10("43776485743678550444492811490514550177096728800832068687396408373151616991234"),
}

//nolint:lll,unused // maybe useful
var targetLattice = &lattice{
	vectors: [][]*big.Int{
		{bigFromBase10("9931322734385697761"), bigFromBase10("9931322734385697761"), bigFromBase10("9931322734385697763"), bigFromBase10("9931322734385697764")},
		{bigFromBase10("4965661367192848881"), bigFromBase10("4965661367192848881"), bigFromBase10("4965661367192848882"), bigFromBase10("-9931322734385697762")},
		{bigFromBase10("-9931322734385697762"), bigFromBase10("-4965661367192848881"), bigFromBase10("4965661367192848881"), bigFromBase10("-4965661367192848882")},
		{bigFromBase10("9931322734385697763"), bigFromBase10("-4965661367192848881"), bigFromBase10("-4965661367192848881"), bigFromBase10("-4965661367192848881")},
	},
	inverse: []*big.Int{
		bigFromBase10("734653495049373973658254490726798021314063399421879442165"),
		bigFromBase10("147946756881789319000765030803803410728"),
		bigFromBase10("-147946756881789319005730692170996259609"),
		bigFromBase10("1469306990098747947464455738335385361643788813749140841702"),
	},
	det: new(big.Int).Set(Order),
}

type lattice struct {
	vectors [][]*big.Int
	inverse []*big.Int
	det     *big.Int
}

// decompose takes a scalar mod Order as input and finds a short, positive decomposition of it wrt to the lattice basis.
func (l *lattice) decompose(k *big.Int) []*big.Int {
	_ = "STUB: not implemented"

	// Calculate closest vector in lattice to <k,0,0,...> with Babai's rounding.
	return nil
}

// Transform vectors according to c and subtract <k,0,0,...>.

func (l *lattice) Precompute(add func(i, j uint)) { _ = "STUB: not implemented"; return }

func (l *lattice) Multi(scalar *big.Int) []uint8 { _ = "STUB: not implemented"; return nil }

// round sets num to num/denom rounded to the nearest integer.
func round(num, denom *big.Int) { _ = "STUB: not implemented"; return }

// todo CondAssignment
