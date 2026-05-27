package protobuf

import (
	"reflect"
)

// Constructors represents a map defining how to instantiate any interface
// types that Decode() might encounter while reading and decoding structured
// data. The keys are reflect.Type values denoting interface types. The
// corresponding values are functions expected to instantiate, and initialize
// as necessary, an appropriate concrete object type supporting that
// interface. A caller could use this capability to support
// dynamic instantiation of objects of the concrete type
// appropriate for a given abstract type.
type Constructors map[reflect.Type]func() interface{}

// String returns an easy way to visualize what you have in your constructors.
func (c *Constructors) String() string { _ = "STUB: not implemented"; return "" }

type DecodingFieldError struct {
	Field string
	Err   error
}

// Implement error interface
func (e *DecodingFieldError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *DecodingFieldError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newDecodingFieldError(field string, err error) *DecodingFieldError {
	_ = "STUB: not implemented"
	return nil
}

// Decoder is the main struct used to decode a protobuf blob.
type decoder struct {
	nm Constructors
}

// Decode a protocol buffer into a Go struct.
// The caller must pass a pointer to the struct to decode into.
//
// Decode() currently does not explicitly check that all 'required' fields
// are actually present in the input buffer being decoded.
// If required fields are missing, then the corresponding fields
// will be left unmodified, meaning they will take on
// their default Go zero values if Decode() is passed a fresh struct.
func Decode(buf []byte, structPtr interface{}) error { _ = "STUB: not implemented"; return nil }

// DecodeWithConstructors is like Decode, but you can pass a map of
// constructors with which to instantiate interface types.
func DecodeWithConstructors(buf []byte, structPtr interface{}, cons Constructors) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// if its NOT a pointer, it is bad return an error

// Decode a Protocol Buffers message into a Go struct.
// The Kind of the passed value v must be Struct.
func (de *decoder) message(buf []byte, sval reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// Interface are not reset because the decoder won't
// be able to instantiate it again in some scenarios.

// Decode all the fields

// Parse the key

// Lookup the corresponding struct field.
// Leave field with a zero Value if fieldnum is out-of-range.
// In this case, as well as for blank fields,
// value() will just skip over and discard the field content.

// For fields within embedded structs, ensure the embedded values aren't nil.

// For more debugging output, uncomment the following three lines.
// if fieldi < len(fields){
//   fmt.Printf("Decoding FieldName %+v\n", fields[fieldi].Field)
// }
// Decode the field's value

// Pull a value from the buffer and put it into a reflective Value.
func (de *decoder) value(wiretype int, buf []byte,
	val reflect.Value) ([]byte, error) {
	_ = "STUB: not implemented"

	// Break out the value from the buffer based on the wire type
	return nil, nil
}

// varint

// 32-bit

// 64-bit

// length-delimited

// We've gotten the value out of the buffer,
// now put it into the appropriate reflective Value.

func (de *decoder) decodeSignedInt(wiretype int, v uint64) (int64, error) {
	_ = "STUB: not implemented"
	// encoded as varint
	return 0, nil
}

// sfixed32

// sfixed64

func (de *decoder) putvalue(wiretype int, val reflect.Value,
	v uint64, vb []byte) error {
	_ = "STUB: not implemented"
	// If val is not settable, it either represents an out-of-range field
	// or an in-range but blank (padding) field in the struct.
	// In this case, simply ignore and discard the field's content.
	return nil
}

// Signed integers may be encoded either zigzag-varint or fixed
// Note that protobufs don't support 8- or 16-bit ints.

// Varint-encoded 32-bit and 64-bit unsigned integers.

// ufixed32

// ufixed64

// Fixed-length 32-bit floats.

// Fixed-length 64-bit floats.

// Length-delimited string.

// Embedded message

// Optional field
// Instantiate pointer's element type.

// Repeated field or byte-slice

// make(map[k]v):

// Abstract field: instantiate via dynamic constructor.

// Backwards compatible usage of the default constructors

// As pointers to interface are discouraged in Go, we use
// the generator only for interface types

// If the object support self-decoding, use that.

// Decode into the object the interface points to.
// XXX perhaps better ONLY to support self-decoding
// for interface fields?

// Instantiate an arbitrary type, handling dynamic interface types.
// Returns a Ptr value.
func (de *decoder) instantiate(t reflect.Type) reflect.Value {
	_ = "STUB: not implemented"

	// If it's an interface type, lookup a dynamic constructor for it.
	return *new(reflect.Value)
}

// Otherwise, for all concrete types, just instantiate directly.

var sfixed32type = reflect.TypeOf(Sfixed32(0))
var sfixed64type = reflect.TypeOf(Sfixed64(0))
var ufixed32type = reflect.TypeOf(Ufixed32(0))
var ufixed64type = reflect.TypeOf(Ufixed64(0))

// Handle decoding of slices
func (de *decoder) slice(slval reflect.Value, vb []byte) error {
	_ = "STUB: not implemented"
	// Find the element type, and create a temporary instance of it.
	return nil
}

// Decide on the wiretype to use for decoding.

// Packed 32-bit representation

// Packed 64-bit representation

// Packed 32-bit representation

// Packed 64-bit representation

// Packed varint representation

// Packed 32-bit representation

// Packed 64-bit representation

// Unpacked byte-slice

// no SetByte method in reflect so has to pass down by uint64

// Other unpacked repeated types
// Just unpack and append one value from vb.

// Decode packed values from the buffer and append them to the slice.

// Handles the entry k,v of a map[K]V
func (de *decoder) mapEntry(slval reflect.Value, vb []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// for repeated values (slices etc)

// We did not decode the key or the value in the map entry.
// Either way, it's an invalid map entry.
