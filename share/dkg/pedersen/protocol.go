package dkg

import (
	"time"
)

// Board is the interface between the dkg protocol and the external world. It
// consists in pushing packets out to other nodes and receiving in packets from
// the other nodes. A common board would use the network as the underlying
// communication mechanism but one can also use a smart contract based
// approach.
type Board interface {
	PushDeals(*DealBundle)
	IncomingDeal() <-chan DealBundle
	PushResponses(*ResponseBundle)
	IncomingResponse() <-chan ResponseBundle
	PushJustifications(*JustificationBundle)
	IncomingJustification() <-chan JustificationBundle
}

// Phaser must signal on its channel when the protocol should move to a next
// phase. Phase must be sequential: DealPhase (start), ResponsePhase,
// JustifPhase and then FinishPhase.
// Note that if the dkg protocol finishes before the phaser sends the
// FinishPhase, the protocol will not listen on the channel anymore. This can
// happen if there is no complaints, or if using the "FastSync" mode.
// Most of the times, user should use the TimePhaser when using the network, but
// if one wants to use a smart contract as a board, then the phaser can tick at
// certain blocks, or when the smart contract tells it.
type Phaser interface {
	NextPhase() chan Phase
}

// TimePhaser is a phaser that sleeps between the different phases and send the
// signal over its channel.
type TimePhaser struct {
	out   chan Phase
	sleep func(Phase)
}

func NewTimePhaser(p time.Duration) *TimePhaser { _ = "STUB: not implemented"; return nil }

func NewTimePhaserFunc(sleepPeriod func(Phase)) *TimePhaser { _ = "STUB: not implemented"; return nil }

func (t *TimePhaser) Start() { _ = "STUB: not implemented"; return }

func (t *TimePhaser) NextPhase() chan Phase {
	_ = "STUB: not implemented"

	// Protocol contains the logic to run a DKG protocol over a generic broadcast
	// channel, called Board. It handles the receival of packets, ordering of the
	// phases and the termination. A protocol can be ran over a network, a smart
	// contract, or anything else that is implemented via the Board interface.
	return nil
}

type Protocol struct {
	board     Board
	phaser    Phaser
	dkg       *DistKeyGenerator
	canIssue  bool
	res       chan OptionResult
	skipVerif bool
}

func NewProtocol(c *Config, b Board, phaser Phaser, skipVerification bool) (*Protocol, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Protocol) Info(keyvals ...any) { _ = "STUB: not implemented"; return }

func (p *Protocol) Error(keyvals ...any) { _ = "STUB: not implemented"; return }

func (p *Protocol) Start() { _ = "STUB: not implemented"; return }

func (p *Protocol) startFast() { _ = "STUB: not implemented"; return }

// we keep the phase in sync with the dkg phase

// each of the following function returns true or false depending on whether
// the protocol should be aborted or not.

// for all dealers, we should be in the DealPhase

// for all *new* share holders, we should be in the InitPhase

// always return false when we are in the finish phase - we quit the
// protocol.

// whatever happens here, if phaser says it's finished we finish

// we finish only if it's time to do so, maybe we received
// justifications but are not in the right phase yet since it
// may not be the right time or haven't received enough msg from
// previous phase

func (p *Protocol) verify(packet Packet) error { _ = "STUB: not implemented"; return nil }

func (p *Protocol) sendDeals() bool { _ = "STUB: not implemented"; return false }

func (p *Protocol) sendResponses(deals []*DealBundle) bool { _ = "STUB: not implemented"; return false }

// we signal the end since we can't go on

func (p *Protocol) sendJustifications(resps []*ResponseBundle) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Protocol) finish(justifs []*JustificationBundle) { _ = "STUB: not implemented"; return }

func (p *Protocol) WaitEnd() <-chan OptionResult { _ = "STUB: not implemented"; return nil }

type OptionResult struct {
	Result *Result
	Error  error
}

type set struct {
	vals map[Index]Packet
	bad  []Index
}

func newSet() *set { _ = "STUB: not implemented"; return nil }

func (s *set) Push(p Packet) { _ = "STUB: not implemented"; return }

// already misbehaved before

// bad behavior - we evict

// same packet just rebroadcasted - all good

func (s *set) isBad(idx Index) bool { _ = "STUB: not implemented"; return false }

func (s *set) ToDeals() []*DealBundle { _ = "STUB: not implemented"; return nil }

func (s *set) ToResponses() []*ResponseBundle { _ = "STUB: not implemented"; return nil }

func (s *set) ToJustifications() []*JustificationBundle { _ = "STUB: not implemented"; return nil }

func (s *set) Len() int { _ = "STUB: not implemented"; return 0 }
