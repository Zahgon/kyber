package protobuf

import (
	"io"
	"reflect"
	"regexp"
)

const protoTemplate = `[[range $name, $values := .Enums]]
enum [[$name|$.Renamer.TypeName]] {[[range $values]]
  [[.Name|$.Renamer.ConstName]] = [[.Value]];[[end]]
}

[[end]][[range .Types]]
message [[.Name|$.Renamer.TypeName]] {[[range .|Fields]]
  [[.|TypeName]] [[.|$.Renamer.FieldName]] = [[.ID]][[.|Options]];[[end]]
}
[[end]]
`

var splitName = regexp.MustCompile(`((?:ID)|(?:[A-Z][a-z_0-9]+)|([\w\d]+))`)

func typeIndirect(t reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func typeName(f ProtoField, enums enumTypeMap, renamer GeneratorNamer) (s string) {
	_ = "STUB: not implemented"
	return ""
}

func fieldPrefix(f ProtoField, def TagPrefix) string { _ = "STUB: not implemented"; return "" }

func innerTypeName(t reflect.Type, enums enumTypeMap, renamer GeneratorNamer) string {
	_ = "STUB: not implemented"
	return ""
}

// we have to do this again (otherwise we'll end up with an empty name for the value):

// here we can just use the value's type:

func options(f ProtoField) string { _ = "STUB: not implemented"; return "" }

type GeneratorNamer interface {
	FieldName(ProtoField) string
	TypeName(name string) string
	ConstName(name string) string
}

// DefaultGeneratorNamer renames symbols when mapping from Go to .proto files.
//
// The rules are:
// - Field names are mapped from SomeFieldName to some_field_name.
// - Type names are not modified.
// - Constants are mapped form SomeConstantName to SOME_CONSTANT_NAME.
type DefaultGeneratorNamer struct{}

func (d *DefaultGeneratorNamer) FieldName(f ProtoField) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DefaultGeneratorNamer) TypeName(name string) string { _ = "STUB: not implemented"; return "" }

func (d *DefaultGeneratorNamer) ConstName(name string) string { _ = "STUB: not implemented"; return "" }

type reflectedTypes []reflect.Type

func (r reflectedTypes) Len() int           { _ = "STUB: not implemented"; return 0 }
func (r reflectedTypes) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (r reflectedTypes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type EnumMap map[string]interface{}

type enumValue struct {
	Name  string
	Value Enum
}

type enumValues []enumValue

func (e enumValues) Len() int           { _ = "STUB: not implemented"; return 0 }
func (e enumValues) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (e enumValues) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type enumTypeMap map[string]enumValues

// GenerateProtobufDefinition generates a .proto file from a list of structs via reflection.
// fieldNamer is a function that maps ProtoField types to generated protobuf field names.
func GenerateProtobufDefinition(w io.Writer, types []interface{}, enumMap EnumMap, renamer GeneratorNamer) (err error) {
	_ = "STUB: not implemented"
	return nil
}
