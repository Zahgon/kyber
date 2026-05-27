// Package encoding package provides helper functions to encode/decode a Point/Scalar in
// hexadecimal.
package encoding

import (
	"io"

	"go.dedis.ch/kyber/v4"
)

// ReadHexPoint reads a point from r in hex representation.
func ReadHexPoint(group kyber.Group, r io.Reader) (kyber.Point, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Point), nil
}

// WriteHexPoint writes a point in hex representation to w.
func WriteHexPoint(w io.Writer, point kyber.Point) error { _ = "STUB: not implemented"; return nil }

// ReadHexScalar takes a hex-encoded scalar and returns that scalar,
// optionally an error
func ReadHexScalar(group kyber.Group, r io.Reader) (kyber.Scalar, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), nil
}

// WriteHexScalar converts a scalar key to a hex-string
func WriteHexScalar(_ kyber.Group, w io.Writer, scalar kyber.Scalar) error {
	_ = "STUB: not implemented"
	return nil
}

// PointToStringHex converts a point to a hexadecimal representation
func PointToStringHex(_ kyber.Group, point kyber.Point) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// StringHexToPoint reads a hexadecimal representation of a point from a string.
func StringHexToPoint(group kyber.Group, s string) (kyber.Point, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Point), nil
}

// ScalarToStringHex encodes a scalar to hexadecimal.
func ScalarToStringHex(_ kyber.Group, scalar kyber.Scalar) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// StringHexToScalar reads a scalar in hexadecimal from string
func StringHexToScalar(group kyber.Group, str string) (kyber.Scalar, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), nil
}

func getHex(r io.Reader, l int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
