package dkg

type Status int32

const (
	Success   Status = 0
	Complaint Status = 1
)

type BitSet map[uint32]Status
type StatusMatrix map[uint32]BitSet

func NewStatusMatrix(dealers []Node, shareHolders []Node, status Status) *StatusMatrix {
	_ = "STUB: not implemented"
	return nil
}

func (s *StatusMatrix) StatusesForShare(shareIndex uint32) BitSet {
	_ = "STUB: not implemented"
	return *new(BitSet)
}

func (s *StatusMatrix) StatusesOfDealer(dealerIndex uint32) BitSet {
	_ = "STUB: not implemented"
	return *

	// can panic if indexes are not from the original list of nodes
	new(BitSet)
}

func (s *StatusMatrix) Set(dealer, share uint32, status Status) { _ = "STUB: not implemented"; return }

func (s *StatusMatrix) SetAll(dealer uint32, status Status) { _ = "STUB: not implemented"; return }

func (s *StatusMatrix) AllTrue(dealer uint32) bool { _ = "STUB: not implemented"; return false }

func (s *StatusMatrix) CompleteSuccess() bool { _ = "STUB: not implemented"; return false }

// can panic if indexes are not from the original list of nodes
func (s *StatusMatrix) Get(dealer, share uint32) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}

func (s *StatusMatrix) String() string {
	_ = "STUB: not implemented"
	// get dealer indexes
	return ""
}

// get shareholder indexes

func (b BitSet) LengthComplaints() uint32 { _ = "STUB: not implemented"; return 0 }
