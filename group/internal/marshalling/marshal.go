// Package marshalling provides a common implementation of (un)marshalling method using Writer and Reader.
package marshalling

import (
	"io"
	"reflect"

	"go.dedis.ch/kyber/v4"
)

// PointMarshalTo provides a generic implementation of Point.EncodeTo
// based on Point.Encode.
func PointMarshalTo(p kyber.Point, w io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// PointUnmarshalFrom provides a generic implementation of Point.DecodeFrom,
// based on Point.Decode, or Point.Pick if r is a Cipher or cipher.Stream.
// The returned byte-count is valid only when decoding from a normal Reader,
// not when picking from a pseudorandom source.
func PointUnmarshalFrom(p kyber.Point, r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// no byte-count when picking randomly

// ScalarMarshalTo provides a generic implementation of Scalar.EncodeTo
// based on Scalar.Encode.
func ScalarMarshalTo(s kyber.Scalar, w io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ScalarUnmarshalFrom provides a generic implementation of Scalar.DecodeFrom,
// based on Scalar.Decode, or Scalar.Pick if r is a Cipher or cipher.Stream.
// The returned byte-count is valid only when decoding from a normal Reader,
// not when picking from a pseudorandom source.
func ScalarUnmarshalFrom(s kyber.Scalar, r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// no byte-count when picking randomly

var tScalar = reflect.TypeFor[kyber.Scalar]()
var tPoint = reflect.TypeFor[kyber.Point]()

// GroupNew is the Default implementation of reflective constructor for Group
func GroupNew(g kyber.Group, t reflect.Type) any { _ = "STUB: not implemented"; return *new(any) }
