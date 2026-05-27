package dkg

import (
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/share"
)

// Index is an alias to designate the index of a node. The index is used to
// evaluate the share of a node, and is thereafter fixed. A node will use the
// same index for generating a partial signature afterwards for example.
type Index = uint32

// Node represents the public key and its index amongt the list of participants.
// For a fresh DKG, the index can be anything but we usually take the index that
// corresponds to the position in the list of participants. For a resharing, if
// that node is a node that has already ran the DKG, we need to use the same
// index as it was given in the previous DKG in the list of OldNodes, in the DKG
// config.
type Node struct {
	Index  Index
	Public kyber.Point
}

func (n *Node) Equal(n2 *Node) bool { _ = "STUB: not implemented"; return false }

// Result is the struct that is outputted by the DKG protocol after it finishes.
// It contains both the list of nodes that successfully ran the protocol and the
// share of the node.
type Result struct {
	QUAL []Node
	Key  *DistKeyShare
}

func (r *Result) PublicEqual(r2 *Result) bool { _ = "STUB: not implemented"; return false }

// DistKeyShare holds the share of a distributed key for a participant.
type DistKeyShare struct {
	// Coefficients of the public polynomial holding the public key.
	Commits []kyber.Point
	// Share of the distributed secret which is private information.
	Share *share.PriShare
}

// Public returns the public key associated with the distributed private key.
func (d *DistKeyShare) Public() kyber.Point {
	_ = "STUB: not implemented"
	return *

	// PriShare implements the dss.DistKeyShare interface so either pedersen or
	// rabin dkg can be used with dss.
	new(kyber.Point)
}

func (d *DistKeyShare) PriShare() *share.PriShare {
	_ = "STUB: not implemented"

	// Commitments implements the dss.DistKeyShare interface so either pedersen or
	// rabin dkg can be used with dss.
	return nil
}

func (d *DistKeyShare) Commitments() []kyber.Point {
	_ = "STUB: not implemented"

	// Deal holds the Deal for one participant as well as the index of the issuing
	// Dealer.
	return nil
}

type Deal struct {
	// Index of the share holder
	ShareIndex uint32
	// encrypted share issued to the share holder
	EncryptedShare []byte
}

var _ Packet = (*DealBundle)(nil)

// DealBundle is the struct sent out by dealers that contains all the deals and
// the public polynomial.
type DealBundle struct {
	DealerIndex uint32
	Deals       []Deal
	// Public coefficients of the public polynomial used to create the shares
	Public []kyber.Point
	// SessionID of the current run
	SessionID []byte
	// Signature over the hash of the whole bundle
	Signature []byte
}

// Hash hashes the index, public coefficients and deals
func (d *DealBundle) Hash() ([]byte, error) {
	_ = "STUB: not implemented"
	// first order the deals in a  stable order
	return nil, nil
}

func (d *DealBundle) Index() Index { _ = "STUB: not implemented"; return *new(Index) }

func (d *DealBundle) Sig() []byte {
	_ = "STUB: not implemented"

	// Response holds the Response from another participant as well as the index of
	// the target Dealer.
	return nil
}

type Response struct {
	// Index of the Dealer for which this response is for
	DealerIndex uint32
	Status      Status
}

var _ Packet = (*ResponseBundle)(nil)

// ResponseBundle is the struct sent out by share holder containing the status
// for the deals received in the first phase.
type ResponseBundle struct {
	// Index of the share holder for which these reponses are for
	ShareIndex uint32
	Responses  []Response
	// SessionID of the current run
	SessionID []byte
	// Signature over the hash of the whole bundle
	Signature []byte
}

// Hash hashes the share index and responses
func (b *ResponseBundle) Hash() ([]byte, error) {
	_ = "STUB: not implemented"
	// first order the response slice in a canonical order
	return nil, nil
}

func (b *ResponseBundle) Index() Index { _ = "STUB: not implemented"; return *new(Index) }

func (b *ResponseBundle) Sig() []byte { _ = "STUB: not implemented"; return nil }

func (b *ResponseBundle) String() string { _ = "STUB: not implemented"; return "" }

var _ Packet = (*JustificationBundle)(nil)

// JustificationBundle is the struct that contains all justifications for each
// complaint in the precedent phase.
type JustificationBundle struct {
	DealerIndex    uint32
	Justifications []Justification
	// SessionID of the current run
	SessionID []byte
	// Signature over the hash of the whole bundle
	Signature []byte
}

type Justification struct {
	ShareIndex uint32
	Share      kyber.Scalar
}

func (j *JustificationBundle) Hash() ([]byte, error) {
	_ = "STUB: not implemented"
	// sort them in a canonical order
	return nil, nil
}

func (j *JustificationBundle) Index() Index { _ = "STUB: not implemented"; return *new(Index) }

func (j *JustificationBundle) Sig() []byte {
	_ = "STUB: not implemented"

	// Packet is the interface that implements the three messages that this
	// implementation uses during the different phases. This interface allows to
	// verify a DKG packet without knowing its specific type.
	return nil
}

type Packet interface {
	Hash() ([]byte, error)
	Index() Index
	Sig() []byte
}

// VerifyPacketSignature returns an error if the packet has an invalid
// signature. The signature is verified via the information contained in the
// config, namely the old and new nodes public keys.
func VerifyPacketSignature(c *Config, p Packet) error {
	_ = "STUB: not implemented"
	// this method returns the correct dealers wether this config is for a DKG
	// or a resharing. For a DKG, OldNodes is set to nil, so the new nodes are
	// the ones that are going to be dealers as well.
	return nil
}
