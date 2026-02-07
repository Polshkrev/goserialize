package goserialize

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	csv "github.com/trimmer-io/go-csv"
	"gopkg.in/yaml.v2"
)

// Serializable map object type alias.
type Object = map[string]any

// Slice of serializable objects.
type ObjectList = []Object

// Generic marshal type. The writer takes an object type and returns the raw byte content.
type Writer = func(any) ([]byte, error)

// Generic unmarshal type. The reader type takes in the raw byte content and a pointer to the object.
type Reader = func([]byte, any) error

// Closure around the generic writer function to allow for indenting when serializing json.
// Returns a writer obtained from indent-marshalling of the passed in data.
func jsonIndentWriter(indent string) Writer {
	return func(data any) ([]byte, error) {
		return json.MarshalIndent(data, "", indent)
	}
}

var (
	// Default json reader.
	JSONReader Reader = json.Unmarshal
	// Default yaml reader.
	YAMLReader Reader = yaml.Unmarshal
	// Default toml reader.
	TOMLReader Reader = toml.Unmarshal
	// Default csv reader.
	CSVReader Reader = csv.Unmarshal
	// Default json writer.
	JSONWriter Writer = json.Marshal
	// Indented writer for json types. The indent is based on tabs instead of spaces.
	JSONIndentWriter Writer = jsonIndentWriter("\t")
	// Default yaml writer.
	YAMLWriter Writer = yaml.Marshal
	// Default toml writer.
	TOMLWriter Writer = toml.Marshal
	// Default csv writer.
	CSVWriter Writer = csv.Marshal
)
