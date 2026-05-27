package bigmod

func (x *Nat) CmpGeq(y *Nat) Choice { _ = "STUB: not implemented"; return *new(Choice) }

type Choice uint

func Not(c Choice) Choice { _ = "STUB: not implemented"; return *new(Choice) }

const Yes = Choice(1)
const No = Choice(0)

// CtMask is all 1s if on is yes, and all 0s otherwise.
func CtMask(on Choice) uint { _ = "STUB: not implemented"; return 0 }

// CtEq returns 1 if x == y, and 0 otherwise. The execution time of this
// function does not depend on its inputs.
func CtEq(x, y uint) Choice { _ = "STUB: not implemented"; return *new(Choice) }

func (x *Nat) Assign(on Choice, y *Nat) *Nat { _ = "STUB: not implemented"; return nil }

func (x *Nat) Set(y *Nat) *Nat { _ = "STUB: not implemented"; return nil }

func (x *Nat) SetBytesBigBuffer(b []byte, m *Modulus) (*Nat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// todo check this
func (x *Nat) BitLenAnnounced() int { _ = "STUB: not implemented"; return 0 }

func LimbsSizeInBytes() int { _ = "STUB: not implemented"; return 0 }

func (x *Nat) Bit(i int) uint { _ = "STUB: not implemented"; return 0 }

// 0 <= j < len(x)
