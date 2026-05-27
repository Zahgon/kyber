package protobuf

import (
	"reflect"
	"sync"
)

type TagPrefix int

// Possible tag options.
const (
	TagNone TagPrefix = iota
	TagOptional
	TagRequired
)

func ParseTag(field reflect.StructField) (id int, opt TagPrefix, name string) {
	_ = "STUB: not implemented"
	return 0, *new(TagPrefix), ""
}

// ProtoField contains cached reflected metadata for struct fields.
type ProtoField struct {
	ID     int64
	Prefix TagPrefix
	Name   string // If non-empty, tag-defined field name.
	Index  []int
	Field  reflect.StructField
}

func (p *ProtoField) Required() bool { _ = "STUB: not implemented"; return false }

var cache = map[reflect.Type][]*ProtoField{}
var cacheLock sync.Mutex

func ProtoFields(t reflect.Type) []*ProtoField { _ = "STUB: not implemented"; return nil }

func innerFieldIndexes(id *int, v reflect.Type) []*ProtoField {
	_ = "STUB: not implemented"
	return nil
}
