package internal

import (
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/share"
)

// Suite defines the capabilities required by the v3marshalling package.
type Suite interface {
	// Group is needed for Group.Scalar
	kyber.Group
}

// compatiblePriShare is a struct for PriShare used when marshaling to
// ensure compatibility with V3
type compatiblePriShare struct {
	I int64
	V kyber.Scalar
}

// MarshalPriShare marshals a share.PriShare into bytes or returns an error
// if the encoding did not work. Encoding is compatible with Kyber V3
func MarshalPriShare(priShare *share.PriShare) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalPriShare unmarshals a share.PriShare from bytes or returns an error
// if the decoding did not work. Decoding is compatible with Kyber V3
func UnmarshalPriShare(data []byte, suite Suite) (*share.PriShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check for overflow on I
