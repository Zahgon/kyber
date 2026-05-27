package protobuf

import (
	"bytes"
	"reflect"
	"time"
)

// Message fields declared to have exactly this type
// will be transmitted as fixed-size 32-bit unsigned integers.
type Ufixed32 uint32

// Message fields declared to have exactly this type
// will be transmitted as fixed-size 64-bit unsigned integers.
type Ufixed64 uint64

// Message fields declared to have exactly this type
// will be transmitted as fixed-size 32-bit signed integers.
type Sfixed32 int32

// Message fields declared to have exactly this type
// will be transmitted as fixed-size 64-bit signed integers.
type Sfixed64 int64

// Protobufs enums are transmitted as unsigned varints;
// using this type alias is optional but recommended
// to ensure they get the correct type.
type Enum uint32

type encoder struct {
	bytes.Buffer
}

// Encode a Go struct into protocol buffer format.
// The caller must pass a pointer to the struct to encode.
func Encode(structPtr interface{}) (bytes []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (en *encoder) message(sval reflect.Value) { _ = "STUB: not implemented"; return }

// Encode all fields in-order

// Skip blank/padding fields

var timeType = reflect.TypeOf(time.Time{})
var durationType = reflect.TypeOf(time.Duration(0))

func (en *encoder) value(key uint64, val reflect.Value, prefix TagPrefix) {
	_ = "STUB: not implemented"

	// Non-reflectively handle some of the fixed types
	return
}

// Encode time.Time as sfixed64

// Handle pointer or interface values (possibly within slices).
// Note that this switch has to handle all the cases,
// because custom type aliases will fail the above typeswitch.

// Varint-encoded 32-bit and 64-bit signed integers.
// Note that protobufs don't support 8- or 16-bit ints.

// Varint-encoded 32-bit and 64-bit unsigned integers.

// Fixed-length 32-bit floats.

// Fixed-length 64-bit floats.

// Length-delimited string.

// Embedded messages.

// Length-delimited slices or byte-vectors.

// Optional field: encode only if pointer is non-nil.

// Abstract interface field.

// If the object support self-encoding, use that.

// add the length of the type tag

// Only write the tag if a generator exists

// Encode from the object the interface points to.

func (en *encoder) slice(key uint64, slval reflect.Value) {
	_ = "STUB: not implemented"

	// First handle common cases with a direct typeswitch
	return
}

// Write the whole byte-slice as one key,value pair

// We'll need to use the reflective path

// Encode packed representation key/value pair

// Handle the encoding of an arbritary map[K]V
func (en *encoder) handleMap(key uint64, mpval reflect.Value, prefix TagPrefix) {
	_ = "STUB: not implemented"
	/*
		A map defined as
			map<key_type, value_type> map_field = N;
		is encoded in the same way as
			message MapFieldEntry {
				key_type key = 1;
				value_type value = 2;
			}
			repeated MapFieldEntry map_field = N;
	*/return
}

// illegal map entry values
// - nil message pointers.

var bytesType = reflect.TypeOf([]byte{})

func (en *encoder) sliceReflect(key uint64, slval reflect.Value) { _ = "STUB: not implemented"; return }

// Write the byte-slice as one key,value pair

// Write each element as a separate key,value pair

// Encode packed representation key/value pair

func (en *encoder) uvarint(v uint64) { _ = "STUB: not implemented"; return }

func (en *encoder) svarint(v int64) { _ = "STUB: not implemented"; return }

func (en *encoder) u32(v uint32) { _ = "STUB: not implemented"; return }

func (en *encoder) u64(v uint64) { _ = "STUB: not implemented"; return }
