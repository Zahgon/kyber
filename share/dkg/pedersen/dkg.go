package dkg

import (
	"errors"
	"io"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/share"
	"go.dedis.ch/kyber/v4/sign"
)

type Suite interface {
	kyber.Group
	kyber.HashFactory
	kyber.XOFFactory
	kyber.Random
}

// Config holds all required information to run a fresh DKG protocol or a
// resharing protocol. In the case of a new fresh DKG protocol, one must fill
// the following fields: Suite, Longterm, NewNodes, Threshold (opt). In the case
// of a resharing protocol, one must fill the following: Suite, Longterm,
// OldNodes, NewNodes. If the node using this config is creating new shares
// (i.e. it belongs to the current group), the Share field must be filled in
// with the current share of the node. If the node using this config is a new
// addition and thus has no current share, the PublicCoeffs field be must be
// filled in.
type Config struct {
	Suite Suite

	// Longterm is the longterm secret key.
	Longterm kyber.Scalar

	// Current group of share holders. It will be nil for new DKG. These nodes
	// will have invalid shares after the protocol has been run. To be able to issue
	// new shares to a new group, the group member's public key must be inside this
	// list and in the Share field. Keys can be disjoint or not with respect to the
	// NewNodes list.
	OldNodes []Node

	// PublicCoeffs are the coefficients of the distributed polynomial needed
	// during the resharing protocol. The first coefficient is the key. It is
	// required for new share holders.  It should be nil for a new DKG.
	PublicCoeffs []kyber.Point

	// Expected new group of share holders. These public-key designated nodes
	// will be in possession of new shares after the protocol has been run. To be a
	// receiver of a new share, one's public key must be inside this list. Keys
	// can be disjoint or not with respect to the OldNodes list.
	NewNodes []Node

	// Share to refresh. It must be nil for a new node wishing to
	// join or create a group. To be able to issue new fresh shares to a new group,
	// one's share must be specified here, along with the public key inside the
	// OldNodes field.
	Share *DistKeyShare

	// The threshold to use in order to reconstruct the secret with the produced
	// shares. This threshold is with respect to the number of nodes in the
	// NewNodes list. If unspecified, default is set to
	// `vss.MinimumT(len(NewNodes))`. This threshold indicates the degree of the
	// polynomials used to create the shares, and the minimum number of
	// verification required for each deal.
	Threshold uint32

	// OldThreshold holds the threshold value that was used in the previous
	// configuration. This field MUST be specified when doing resharing, but is
	// not needed when doing a fresh DKG. This value is required to gather a
	// correct number of valid deals before creating the distributed key share.
	// NOTE: this field is always required (instead of taking the default when
	// absent) when doing a resharing to avoid a downgrade attack, where a resharing
	// the number of deals required is less than what it is supposed to be.
	OldThreshold uint32

	// Reader is an optional field that can hold a user-specified entropy
	// source.  If it is set, Reader's data will be combined with random data
	// from crypto/rand to create a random stream which will pick the dkg's
	// secret coefficient. Otherwise, the random stream will only use
	// crypto/rand's entropy.
	Reader io.Reader

	// When UserReaderOnly is set to true, only the user-specified entropy
	// source Reader will be used. This should only be used in tests, allowing
	// reproducibility.
	UserReaderOnly bool

	// FastSync is a mode where nodes sends pre-emptively responses indicating
	// that the shares they received are good. If a share is invalid, a
	// complaint is still sent as usual. This has two consequences:
	//  - In the event all shares are good, nodes don't need to wait for the
	//  timeout; they can already finish the protocol at the point.
	//  - However, it requires nodes to send more messages on the network. We
	//  pass from a O(f) where f is the number of faults to a O(n^2). Note that
	//  the responses messages are small.
	FastSync bool

	// Nonce is required to avoid replay attacks from previous runs of a DKG /
	// resharing. The required property of the Nonce is that it must be unique
	// across runs. A Nonce must be of length 32 bytes. User can get a secure
	// nonce by calling `GetNonce()`.
	Nonce []byte

	// Auth is the scheme to use to authentify the packets sent and received
	// during the protocol.
	Auth sign.Scheme

	// Log enables the DKG logic and protocol to log important events (mostly
	// errors).  from participants. Errors don't mean the protocol should be
	// stopped, so logging is the best way to communicate information to the
	// application layer. It can be nil.
	Log Logger
}

// Phase is a type that represents the different stages of the DKG protocol.
type Phase int

const (
	InitPhase Phase = iota
	DealPhase
	ResponsePhase
	JustifPhase
	FinishPhase
)

func (p Phase) String() string { _ = "STUB: not implemented"; return "" }

// PhaseError is an error recorded at a given phase
type PhaseError struct {
	DkgPhase    Phase // "open", "unlink", etc.
	ActualPhase Phase
}

func (e *PhaseError) Error() string { _ = "STUB: not implemented"; return "" }

// DistKeyGenerator is the struct that runs the DKG protocol.
type DistKeyGenerator struct {
	// config driving the behavior of DistKeyGenerator
	c     *Config
	suite Suite

	long     kyber.Scalar
	pub      kyber.Point
	dpriv    *share.PriPoly
	dpub     *share.PubPoly
	statuses *StatusMatrix
	// the valid shares we received
	validShares map[uint32]kyber.Scalar
	// all public polynomials we have seen
	allPublics map[uint32]*share.PubPoly
	// list of dealers that clearly gave invalid deals / responses / justifs
	evicted []uint32
	// list of share holders that misbehaved during the response phase
	evictedHolders []Index
	state          Phase
	// index in the old list of nodes
	oidx Index
	// index in the new list of nodes
	nidx Index
	// old threshold used in the previous DKG
	oldT uint32
	// new threshold to use in this round
	newT uint32
	// indicates whether we are in the re-sharing protocol or basic DKG
	isResharing bool
	// indicates whether we are able to issue shares or not
	canIssue bool
	// Indicates whether we are able to receive a new share or not
	canReceive bool
	// indicates whether the node holding the pub key is present in the new list
	newPresent bool
	// indicates whether the node is present in the old list
	oldPresent bool
	// public polynomial of the old group
	olddpub *share.PubPoly
}

// NewDistKeyHandler takes a Config and returns a DistKeyGenerator that is able
// to drive the DKG or resharing protocol.
func NewDistKeyHandler(c *Config) (*DistKeyGenerator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// canReceive is true by default since in the default DKG mode everyone
// participates

// if we are not in the new list of nodes, then we definitely can't
// receive anything

// fresk DKG present

// if the user provided a reader, use it alone or combined with crypto/rand

// in fresh dkg case, we consider the old nodes same a new nodes

// resharing case

// resharing case and we are included in the new list of nodes

// take the commits of the share, no need to duplicate information

// oldThreshold is only useful in the context of a new share holder, to
// make sure there are enough correct deals from the old nodes.

// in fast sync mode, we set every shares to complaint by default and
// expect everyone to send success for correct shares

// in normal mode, every shares of other nodes is expected to be
// correct, unless honest nodes send a complaint

// we set the statuses of the shares we expect to receive as complaint
// by default, so if we miss one share or there's an invalid share,
// it'll generate a complaint

func (d *DistKeyGenerator) Deals() (*DealBundle, error) { _ = "STUB: not implemented"; return nil, nil }

// compute share

// we set our own share as true, because we are not malicious!

// we don't send our own share - useless

// ProcessDeals process the deals from all the nodes. Each deal for this node is
// decrypted and stored. It returns a response bundle if there is any invalid or
// missing deals. It returns an error if the node is not in the right state, or
// if there is not enough valid shares, i.e. the dkg is failing already.
func (d *DistKeyGenerator) ProcessDeals(bundles []*DealBundle) (*ResponseBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// oldnode member is not in the right state

// newnode member which is not in the old group is not in the riht state

// a node that is only in the old group should not process deals
// he moves on to the next phase silently

//nolint:nilnil // protocol defined this way

// dont look at our own deal
// Note that's why we are not checking if we are evicted at the end of this function and return an error
// because we're supposing we are honest and we don't look at our own deal

// invalid public polynomial is clearly cheating
// so we evict him from the list
// since we assume broadcast channel, every honest player will evict
// this party as well

// already saw a bundle from the same dealer - clear sign of
// cheating so we evict him from the list

// invalid index for share holder is a clear sign of cheating
// so we evict him from the list
// and we don't even need to look at the rest

// we dont look at other's shares

// check if share is valid w.r.t. public commitment

// invalid share - will issue complaint

// check that the evaluation this public polynomial at 0,
// corresponds to the commitment of the previous the dealer's index

// inconsistent share from old member

// share is valid -> store it

// we set to true the status of each node that are present in both list
// for their respective index -> we assume the share a honest node creates is
// correct for himself - that he won't create an invalid share for himself

// producing response part

// if the node is evicted, we don't even need to send a complaint or a
// response since every honest node evicts him as well.
// XXX Is that always true ? Should we send a complaint still ?

// we send success responses only in fast sync

// dealer i did not give a successful share (or absent etc)

func (d *DistKeyGenerator) ExpectedResponsesFastSync() int { _ = "STUB: not implemented"; return 0 }

// ProcessResponses takes the response from all nodes if any and returns a
// triplet:
// - the result if there is no complaint. If not nil, the DKG is finished.
// - the justification bundle if this node must produce at least one. If nil,
// this node must still wait on the justification phase.
// - error if the dkg must stop now, an unrecoverable failure.
func (d *DistKeyGenerator) ProcessResponses(bundles []*ResponseBundle) (
	res *Result,
	jb *JustificationBundle,
	err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// if we are a old node that will leave

// if we are not in fastsync, we expect only complaints
// if there is no complaints all is good

// just in case we don't treat our own response

// the index of the dealer doesn't exist - clear violation
// so we evict

// we should only receive complaint if we are not in fast sync
// mode - clear violation
// so we evict

// In case of fast sync, we want to make sure all share holders have sent a
// valid response (success or complaint). All share holders that did not
// will be evicted from the final group. Since we are using a broadcast
// channel, if a node is honest, its response will be received by all honest
// nodes.

// we only need to look at the nodes that did not sent any response,
// since the invalid one are already marks as evicted

// we dont evict ourself

// there is no complaint in the responses received and the status matrix
// is all filled with success that means we can finish the protocol -
// regardless of the mode chosen (fast sync or not).

// old nodes that are not present in the new group

// check if there are some node who received at least t complaints.
// In that case, they must be evicted already since their polynomial can
// now be reconstructed so any observer can sign in its place.

// new node that is expecting some justifications

// check if there are justifications this node needs to produce

// create justifications for the requested share

// mark those shares as resolved in the statuses

// no justifications required from us !

// ProcessJustifications takes the justifications of the nodes and returns the
// results if there is enough QUALified nodes, or an error otherwise. Note that
// this method returns "nil,nil" if this node is a node only present in the old
// group of the dkg: indeed a node leaving the group don't need to process
// justifications, and can simply leave the protocol.
func (d *DistKeyGenerator) ProcessJustifications(bundles []*JustificationBundle) (*Result, error) {
	_ = "STUB: not implemented"

	// an old node leaving the group do not need to process justifications.
	// Here we simply return nil to avoid requiring higher level library to
	// think about which node should receive which packet
	//
	//nolint:nilnil // protocol defined this way
	return nil, nil
}

// bundle contains duplicate - clear violation
// so we evict

// we dont treat our own justifications

// index is invalid

// already evicted node

// invalid index - clear violation
// so we evict

// dealer hasn't given any public polynomial at the first phase
// so we evict directly - no need to look at its justifications

// compare commit and public poly

// invalid justification - evict

// check that the evaluation this public polynomial at 0,
// corresponds to the commitment of the previous the dealer's index

// inconsistent share from old member

// valid share -> mark OK

// store the share if it's for us

// check if we are evicted or not

// check if there is enough dealer entries marked as all success

// this dealer has some unjustified shares

// we need enough old QUAL dealers, more than the threshold the old
// group uses

// that should not happen in the threat model but we still returns the
// fatal error here so DKG do not finish

// otherwise it's all good - let's compute the result

func (d *DistKeyGenerator) computeResult() (*Result, error) {
	_ = "STUB: not implemented"
	return nil,

		// add a full complaint row on the nodes that are evicted
		nil
}

// add all the shares and public polynomials together for the deals that are
// valid ( equivalently or all justified)

// instead of adding, in this case, we interpolate all shares

func (d *DistKeyGenerator) computeResharingResult() (*Result, error) {
	_ = "STUB: not implemented"
	// only old nodes sends shares
	return nil, nil
}

// this dealer has some unjustified shares
// no need to check for th e evicted list since the status matrix
// has been set previously to complaint for those

// share of dist. secret. Invertion of rows/column

// the private polynomial is generated from the old nodes, thus inheriting
// the old threshold condition

// recover public polynomial by interpolating coefficient-wise all
// polynomials
// the new public polynomial must however have "newT" coefficients since it
// will be held by the new nodes.

// take all i-th coefficients

// using the old threshold / length because there are at most
// len(d.c.OldNodes) i-th coefficients since they are the one generating one
// each, thus using the old threshold.

// Reconstruct the final public polynomial

// To compute the QUAL in the resharing case, we take each new nodes whose
// column in the status matrix contains true for all valid dealers.
// That means:
// 1. we only look for valid deals
// 2. we only take new nodes, i.e. new participants, that correctly ran the
// protocol (i.e. absent nodes will not be counted)

// look if this node is also a dealer which have been misbehaving

// it's a valid dealer as well

// it's an invalid dealer, so we evict him

// we also check if he has been misbehaving during the response phase
// only

func (d *DistKeyGenerator) computeDKGResult() (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this dealer has some unjustified shares
// no need to check the evicted list since the status matrix
// has been set previously to complaint for those

// however we do need to check for evicted share holders since in this
// case (DKG) both are the same.

var ErrEvicted = errors.New("our node is evicted from list of qualified participants")

// checkIfEvicted returns an error if this node is in one of the two eviction list. This is useful to detect
// our own misbehaviour or lack of connectivity: for example if this node can receive messages from others but is
// not able to send, everyone will send a complaint about this node, and thus it is going to be evicted.
// This method checks if you are and returns an error from the DKG to stop it. Once evicted a node's messages are
// not processed anymore and it is left out of the protocol.
func (d *DistKeyGenerator) checkIfEvicted(phase Phase) error { _ = "STUB: not implemented"; return nil }

// For DKG -> for all phases look at evicted dealers since both lists are the same anyway
// For resharing ->  only at response phase we evict some new share holders
// 			otherwise, it's only dealers we evict (since deal and justif are made by dealers)

// we can't be evicted as an old node leaving the group here

// we can't be evicted as a new node in this setting

func findPub(list []Node, toFind kyber.Point) (Index, bool) {
	_ = "STUB: not implemented"
	return *new(Index), false
}

func findIndex(list []Node, index Index) (kyber.Point, bool) {
	_ = "STUB: not implemented"
	return *new(kyber.Point), false
}

func MinimumT(n uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func isIndexIncluded(list []Node, index uint32) bool { _ = "STUB: not implemented"; return false }

// NonceLength is the length of the nonce
const NonceLength = 32

// GetNonce returns a suitable nonce to feed in the DKG config.
func GetNonce() []byte { _ = "STUB: not implemented"; return nil }

func (d *DistKeyGenerator) sign(p Packet) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DistKeyGenerator) Info(keyvals ...any) { _ = "STUB: not implemented"; return }

func (d *DistKeyGenerator) Error(keyvals ...any) { _ = "STUB: not implemented"; return }

func (c *Config) Info(keyvals ...any) { _ = "STUB: not implemented"; return }

func (c *Config) Error(keyvals ...any) { _ = "STUB: not implemented"; return }

// CheckForDuplicates looks at the lits of node indices in the OldNodes and
// NewNodes list. It returns an error if there is a duplicate in either list.
// NOTE: It only looks at indices because it is plausible that one party may
// have multiple indices for the protocol, i.e. a higher "weight".
func (c *Config) CheckForDuplicates() error { _ = "STUB: not implemented"; return nil }
