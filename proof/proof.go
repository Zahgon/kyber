// Package proof implements generic support for Sigma-protocols
// and discrete logarithm proofs in the Camenisch/Stadler framework.
// For the cryptographic foundations of this framework see
// "Proof Systems for General Statements about Discrete Logarithms" at
// ftp://ftp.inf.ethz.ch/pub/crypto/publications/CamSta97b.pdf.
package proof

import (
	"go.dedis.ch/kyber/v4"
)

// Suite defines the functionalities needed for this package to operate
// correctly. It provides a general abstraction to easily change the underlying
// implementations.
type Suite interface {
	kyber.Group
	kyber.HashFactory
	kyber.Encoding
	kyber.XOFFactory
	kyber.Random
}

/*
A Predicate is a composable logic expression in a knowledge proof system,
representing a "knowledge specification set" in Camenisch/Stadler terminology.
Atomic predicates in this system are statements of the form P=x1*B1+...+xn*Bn,
indicating the prover knows secrets x1,...,xn that make the statement true,
where P and B1,...,Bn are public points known to the verifier.
These atomic Rep (representation) predicates may be combined
with logical And and Or combinators to form composite statements.
Predicate objects, once created, are immutable and safe to share
or reuse for any number of proofs and verifications.

After constructing a Predicate using the Rep, And, and Or functions below,
the caller invokes Prover() to create a Sigma-protocol prover.
Prover() requires maps defining the values of both the Scalar variables
and the public Point variables that the Predicate refers to.
If the statement contains logical Or operators, the caller must also pass
a map containing branch choices for each Or predicate
in the "proof-obligated path" down through the Or predicates.
See the examples provided for the Or function for more details.

Similarly, the caller may invoke Verifier() to create
a Sigma-protocol verifier for the predicate.
The caller must pass a map defining the values
of the public Point variables that the proof refers to.
The verifier need not be provided any secrets or branch choices, of course.
(If the verifier needed those then they wouldn't be secret, would they?)

Currently we require that all Or operators be above all And operators
in the expression - i.e., Or-of-And combinations are allowed,
but no And-of-Or predicates.
We could rewrite expressions into this form as Camenisch/Stadler suggest,
but that could run a risk of unexpected exponential blowup in the worst case.
We could avoid this risk by not rewriting the expression tree,
but instead generating Pedersen commits for variables that need to "cross"
from one OR-domain to another non-mutually-exclusive one.
For now we simply require expressions to be in the appropriate form.
*/
type Predicate interface {

	// Create a Prover proving the statement this Predicate represents.
	Prover(suite Suite, secrets map[string]kyber.Scalar,
		points map[string]kyber.Point, choice map[Predicate]int) Prover

	// Create a Verifier for the statement this Predicate represents.
	Verifier(suite Suite, points map[string]kyber.Point) Verifier

	// Produce a human-readable string representation of the predicate.
	String() string

	// precedence-sensitive helper stringifier.
	precString(prec int) string

	// prover/verifier: enumerate the variables named in a predicate
	enumVars(prf *proof)

	// prover: recursively produce all commitments
	commit(prf *proof, w kyber.Scalar, v []kyber.Scalar) error

	// prover: given challenge, recursively produce all responses
	respond(prf *proof, c kyber.Scalar, r []kyber.Scalar) error

	// verifier: get all the commitments required in this predicate,
	// and fill the r slice with empty secrets for responses needed.
	getCommits(prf *proof, r []kyber.Scalar) error

	// verifier: check all commitments against challenges and responses
	verify(prf *proof, c kyber.Scalar, r []kyber.Scalar) error
}

// stringification precedence levels
const (
	precNone = iota
	precOr
	precAnd
	precAtom
)

// Internal prover/verifier state
type proof struct {
	s Suite

	nsvars     int            // number of Scalar variables
	npvars     int            // number of Point variables
	svar, pvar []string       // Scalar and Point variable names
	sidx, pidx map[string]int // Maps from strings to variable indexes

	pval map[string]kyber.Point // values of public Point variables

	// prover-specific state
	pc     ProverContext
	sval   map[string]kyber.Scalar   // values of private Scalar variables
	choice map[Predicate]int         // OR branch choices set by caller
	pp     map[Predicate]*proverPred // per-predicate prover state

	// verifier-specific state
	vc VerifierContext
	vp map[Predicate]*verifierPred // per-predicate verifier state
}
type proverPred struct {
	w  kyber.Scalar   // secret pre-challenge
	v  []kyber.Scalar // secret blinding factor for each variable
	wi []kyber.Scalar // OR predicates: individual sub-challenges
}
type verifierPred struct {
	V kyber.Point    // public commitment produced by verifier
	r []kyber.Scalar // per-variable responses produced by verifier
}

////////// Rep predicate //////////

// A term describes a point-multiplication term in a representation expression.
type term struct {
	S string // Scalar multiplier for this term
	B string // Generator for this term
}

type repPred struct {
	P string // Public point of which a representation is known
	T []term // Terms comprising the known representation
}

// Rep creates a predicate stating that the prover knows
// a representation of a point P with respect to
// one or more secrets and base point pairs.
//
// In its simplest usage, Rep indicates that the prover knows a secret x
// that is the (elliptic curve) discrete logarithm of a public point P
// with respect to a well-known base point B:
//
//	Rep(P,x,B)
//
// Rep can take any number of (Scalar,Base) variable name pairs, however.
// A Rep statement of the form Rep(P,x1,B1,...,xn,Bn)
// indicates that the prover knows secrets x1,...,xn
// such that point P is the sum x1*B1+...+xn*Bn.
func Rep(P string, SB ...string) Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// Return a string representation of this proof-of-representation predicate,
// mainly for debugging.
func (rp *repPred) String() string { _ = "STUB: not implemented"; return "" }

func (rp *repPred) precString(_ int) string { _ = "STUB: not implemented"; return "" }

func (rp *repPred) enumVars(prf *proof) { _ = "STUB: not implemented"; return }

func (rp *repPred) commit(prf *proof, w kyber.Scalar, pv []kyber.Scalar) error {
	_ = "STUB: not implemented"

	// Create per-predicate prover state
	return nil
}

// Compute commit V=wY+v1G1+...+vkGk

// We're on a non-obligated branch

// We're on a proof-obligated branch, so w=0

// current term

// Choose a blinding secret the first time
// we encounter each variable

// Encode and send the commitment to the verifier

func (rp *repPred) respond(prf *proof, c kyber.Scalar,
	pr []kyber.Scalar) error {
	_ = "STUB: not implemented"

	// Create a response array for this OR-domain if not done already
	return nil
}

// current term

// Produce a correct response for each variable
// the first time we encounter that variable.

// We're on a non-proof-obligated branch:
// w was our challenge, v[s] is our response.

// We're on a proof-obligated branch,
// so we need to calculate the correct response
// as r = v-cx where x is the secret variable

// Send our responses if we created the array (i.e., if pr == nil)

func (rp *repPred) getCommits(prf *proof, pr []kyber.Scalar) error {
	_ = "STUB: not implemented"

	// Create per-predicate verifier state
	return nil
}

// Get the commitment for this representation

// Fill in the r vector with the responses we'll need.

// current term

func (rp *repPred) verify(prf *proof, c kyber.Scalar, pr []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil

	// Get the needed responses if a parent didn't already
}

// Recompute commit V=cY+r1G1+...+rkGk

// current term

func (rp *repPred) Prover(suite Suite, secrets map[string]kyber.Scalar,
	points map[string]kyber.Point,
	choice map[Predicate]int) Prover {
	_ = "STUB: not implemented"
	return *new(Prover)
}

func (rp *repPred) Verifier(suite Suite,
	points map[string]kyber.Point) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}

////////// And predicate //////////

type andPred []Predicate

// And predicate states that all of the constituent sub-predicates are true.
// And predicates may contain Rep predicates and/or other And predicates.
func And(sub ...Predicate) Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// Return a string representation of this AND predicate, mainly for debugging.
func (ap *andPred) String() string { _ = "STUB: not implemented"; return "" }

func (ap *andPred) precString(prec int) string { _ = "STUB: not implemented"; return "" }

func (ap *andPred) enumVars(prf *proof) { _ = "STUB: not implemented"; return }

func (ap *andPred) commit(prf *proof, w kyber.Scalar, pv []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil

	// Create per-predicate prover state
}

// Recursively generate commitments

func (ap *andPred) respond(prf *proof, c kyber.Scalar, pr []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil

	// Recursively compute responses in all sub-predicates
}

func (ap *andPred) getCommits(prf *proof, pr []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil

	// Create per-predicate verifier state
}

func (ap *andPred) verify(prf *proof, c kyber.Scalar, pr []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil
}

func (ap *andPred) Prover(suite Suite, secrets map[string]kyber.Scalar,
	points map[string]kyber.Point,
	choice map[Predicate]int) Prover {
	_ = "STUB: not implemented"
	return *new(Prover)
}

func (ap *andPred) Verifier(suite Suite,
	points map[string]kyber.Point) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}

////////// Or predicate //////////

type orPred []Predicate

// Or predicate states that the prover knows
// at least one of the sub-predicates to be true,
// but the proof does not reveal any information about which.
func Or(sub ...Predicate) Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// Return a string representation of this OR predicate, mainly for debugging.
func (op *orPred) String() string { _ = "STUB: not implemented"; return "" }

func (op *orPred) precString(prec int) string { _ = "STUB: not implemented"; return "" }

func (op *orPred) enumVars(prf *proof) { _ = "STUB: not implemented"; return }

func (op *orPred) commit(prf *proof, w kyber.Scalar, pv []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil
}

// only happens within an AND expression

// Create per-predicate prover state

// Choose pre-challenges for our subs.

// We're on a proof-obligated branch;
// choose random pre-challenges for only non-obligated subs.

// else wi[i] == nil for proof-obligated sub

// Since w != nil, we're in a non-obligated branch,
// so choose random pre-challenges for all subs
// such that they add up to the master pre-challenge w.
// index of last sub

// choose all but last

func commitmentProducer(prf *proof, wi []kyber.Scalar, sub []Predicate) error {
	_ = "STUB: not implemented"
	// Now recursively choose commitments within each sub
	return nil
}

// Fresh variable-blinding secrets for each pre-commitment

func (op *orPred) respond(prf *proof, c kyber.Scalar, pr []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil
}

// Calculate the challenge for the proof-obligated subtree

// If there's more than one choice, send all our sub-challenges.

// Recursively compute responses in all subtrees

// Get from the verifier all the commitments needed for this predicate
func (op *orPred) getCommits(prf *proof, _ []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil
}

func (op *orPred) verify(prf *proof, c kyber.Scalar, pr []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the prover's sub-challenges

// Make sure they add up to the parent's composite challenge

// trivial single-sub OR

// Recursively verify all subs

func (op *orPred) Prover(suite Suite, secrets map[string]kyber.Scalar,
	points map[string]kyber.Point,
	choice map[Predicate]int) Prover {
	_ = "STUB: not implemented"
	return *new(Prover)
}

func (op *orPred) Verifier(suite Suite,
	points map[string]kyber.Point) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}

/*
type lin struct {
	a1,a2,b kyber.Scalar
	x1,x2 PriVar
}
*/

// Construct a predicate asserting a linear relationship a1x1+a2x2=b,
// where a1,a2,b are public values and x1,x2 are secrets.
/*
func (p *Prover) Linear(a1,a2,b kyber.Scalar, x1,x2 PriVar) {
	return &lin{a1,a2,b,x1,x2}
}
*/

func (prf proof) init(suite Suite, pred Predicate) *proof {
	prf.s = suite

	// Enumerate all the variables in a consistent order.
	// Reserve variable index 0 for convenience.
	prf.svar = []string{""}
	prf.pvar = []string{""}
	prf.sidx = make(map[string]int)
	prf.pidx = make(map[string]int)
	pred.enumVars(&prf)
	prf.nsvars = len(prf.svar)
	prf.npvars = len(prf.pvar)

	return &prf
}

func (prf *proof) enumScalarVar(name string) { _ = "STUB: not implemented"; return }

func (prf *proof) enumPointVar(name string) { _ = "STUB: not implemented"; return }

// Make a response-array if that wasn't already done in a parent predicate.
func (prf *proof) makeScalars(pr []kyber.Scalar) []kyber.Scalar {
	_ = "STUB: not implemented"
	return nil
}

// Transmit our response-array if a corresponding makeScalars() created it.
func (prf *proof) sendResponses(pr []kyber.Scalar, r []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil
}

// Send responses only for variables
// that were used in this OR-domain.

// In the verifier, get the responses at the top of an OR-domain,
// if a corresponding makeScalars() call created it.
func (prf *proof) getResponses(pr []kyber.Scalar, r []kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil
}

func (prf *proof) prove(p Predicate, sval map[string]kyber.Scalar,
	pval map[string]kyber.Point,
	choice map[Predicate]int, pc ProverContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate all commitments

// Generate top-level challenge from public randomness

// Generate all responses based on master challenge

func (prf *proof) verify(p Predicate, pval map[string]kyber.Point,
	vc VerifierContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the commitments from the verifier,
// and calculate the sets of responses we'll need for each OR-domain.

// Produce the top-level challenge

// Check all the responses and sub-challenges against the commitments.

// Produce a higher-order Prover embodying a given proof predicate.
func (prf *proof) prover(p Predicate, sval map[string]kyber.Scalar,
	pval map[string]kyber.Point,
	choice map[Predicate]int) Prover {
	_ = "STUB: not implemented"
	return *new(Prover)
}

// Produce a higher-order Verifier embodying a given proof predicate.
func (prf *proof) verifier(p Predicate, pval map[string]kyber.Point) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}
