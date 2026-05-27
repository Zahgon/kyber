// Package suites allows callers to look up Kyber suites by name.
//
// Currently, only the "ed25519" suite is available with a constant
// time implementation and the other ones use variable time algorithms.
package suites

import (
	"errors"

	"go.dedis.ch/kyber/v4"
)

// Suite is the sum of all suites mix-ins in Kyber.
type Suite interface {
	kyber.Encoding
	kyber.Group
	kyber.HashFactory
	kyber.XOFFactory
	kyber.Random
}

var suites = map[string]Suite{}

var requireConstTime = false

// register is called by suites to make themselves known to Kyber.
func register(s Suite) { _ = "STUB: not implemented"; return }

// ErrUnknownSuite indicates that the suite was not one of the
// registered suites.
var ErrUnknownSuite = errors.New("unknown suite")

// Find looks up a suite by name.
func Find(name string) (Suite, error) { _ = "STUB: not implemented"; return *new(Suite), nil }

// MustFind looks up a suite by name and panics if it is not found.
func MustFind(name string) Suite { _ = "STUB: not implemented"; return *new(Suite) }

// RequireConstantTime causes all future calls to Find and MustFind to only
// search for suites where the implementation is constant time.
// It should be called in an init() function for the main package
// of users of Kyber who need to be sure to avoid variable time implementations.
// Once constant time implementations are required, there is no way to
// turn it back off (by design).
//
// At this time, the only constant time crypto suite is "Ed25519".
func RequireConstantTime() { _ = "STUB: not implemented"; return }
