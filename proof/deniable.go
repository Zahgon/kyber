package proof

import (
	"bytes"

	"go.dedis.ch/kyber/v4"
)

// DeniableProver is a Protocol implementing an interactive Sigma-protocol
// to prove a particular statement to the other participants.
// Optionally the Protocol participant can also verify
// the Sigma-protocol proofs of any or all of the other participants.
// Different participants may produce different proofs of varying sizes,
// and may even consist of different numbers of steps.
func DeniableProver(suite Suite, self int, prover Prover,
	verifiers []Verifier) Protocol {
	_ = "STUB: not implemented"
	return *new(Protocol)
}

type deniableProver struct {
	suite Suite   // Agreed-on ciphersuite for protocol
	self  int     // Our own node number
	sc    Context // Clique protocol context

	// verifiers for other nodes' proofs
	dv []*deniableVerifier

	// per-step state
	key  []byte        // Secret pre-challenge we committed to
	msg  *bytes.Buffer // Buffer in which to build prover msg
	msgs [][]byte      // All messages from last proof step

	pubrand kyber.XOF
	prirand kyber.XOF

	// Error/success indicators for all participants
	err []error
}

func (dp *deniableProver) run(suite Suite, self int, prv Prover,
	vrf []Verifier, sc Context) []error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize error slice entries to a default error indicator,
// so that forgetting to run a verifier won't look like "success"

// Launch goroutines to run whichever verifiers the caller requested

// Run the prover, which will also drive the verifiers.

// Send the last prover message.
// Make sure the verifiers get to run to completion as well

// keySize is arbitrary, make it long enough to seed the XOF
const keySize = 128

// Start the message buffer off in each step with a randomness commitment
func (dp *deniableProver) initStep() error { _ = "STUB: not implemented"; return nil }

// secret random key

// send commitment to it

// The Sigma-Prover will now append its proof content to dp.msg...

func (dp *deniableProver) proofStep() (bool, error) {
	_ = "STUB: not implemented"

	// Send the randomness commit and accumulated message to the leader,
	// and get all participants' commits, via our star-protocol context.
	return false, nil
}

// Distribute this step's prover messages
// to the relevant verifiers as well,
// waking them up in the process so they can proceed.

// send to verifier

// Collect the verifiers' responses,
// collecting error indicators from verifiers that are done.

// collect verifier responses

// get verifier response
// verifier is done

// verifier needs next challenge

func (dp *deniableProver) challengeStep() error {
	_ = "STUB: not implemented"

	// Send our challenge randomness to the leader, and collect all.
	return nil
}

// XOR together all the participants' randomness contributions,
// check them against the respective commits,
// and ensure ours is included to ensure deniability
// (even if all others turn out to be maliciously generated).

// node i's randomness commitment
// node i's committed random key

// ignore participants who dropped out

// mix in this key

// Use the mix to produce the public randomness needed by the prover

// Distribute the master challenge to any verifiers waiting for it

// so send it

// Setup for the next proof step

func (dp *deniableProver) Put(message any) error {
	_ = "STUB: not implemented"
	// Add onto accumulated prover message
	return nil
}

// Prover will call this after Put()ing all commits for a given step,
// to get the master challenge to be used in its challenge/responses.
func (dp *deniableProver) PubRand(data ...any) error { _ = "STUB: not implemented"; return nil }

// finish proof step

// run challenge step

// Get private randomness
func (dp *deniableProver) PriRand(data ...any) error { _ = "STUB: not implemented"; return nil }

// Interactive Sigma-protocol verifier context.
// Acts as a slave to a deniableProver instance.
type deniableVerifier struct {
	suite Suite

	inbox chan []byte   // Channel for receiving proofs and challenges
	prbuf *bytes.Buffer // Buffer with which to read proof messages

	done chan bool // Channel for sending done status indicators
	err  error     // When done indicates verify error if non-nil

	pubrand kyber.XOF
}

func (dv *deniableVerifier) start(suite Suite, vrf Verifier) { _ = "STUB: not implemented"; return }

// Launch a concurrent goroutine to run this verifier

// Await the prover's first message

// Run the verifier, providing dv as its context

// Signal verifier termination

func (dv *deniableVerifier) getProof() {
	_ = "STUB: not implemented"
	// Get the next message from the prover
	return
}

// Read structured data from the proof
func (dv *deniableVerifier) Get(message any) error { _ = "STUB: not implemented"; return nil }

// Get the next public random challenge.
func (dv *deniableVerifier) PubRand(data ...any) error {
	_ = "STUB: not implemented"

	// Signal that we need the next challenge
	return nil
}

// Wait for it

// Produce the appropriate publicly random stream

// Get the next proof message
