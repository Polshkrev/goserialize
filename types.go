package goserialize

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

// Serializable map object type alias.
type Object = map[string]any

const (
	// Default json file extension.
	JSONType string = "json"
	// Default yaml file extenstion.
	YAMLType string = "yaml"
	// Default toml file extension.
	TOMLType string = "toml"
)

// Generic marshal type. The writer takes an object type and returns the raw byte content.
type Writer = func(any) ([]byte, error)

// Generic unmarshal type. The reader type takes in the raw byte content and a pointer to the object.
type Reader = func([]byte, any) error

var (
	// Default json reader.
	JSONReader Reader = json.Unmarshal
	// Default yaml reader.
	YAMLReader Reader = yaml.Unmarshal
	// Default toml reader.
	TOMLReader Reader = toml.Unmarshal
	// Default json writer.
	JSONWriter Writer = json.Marshal
	// Default yaml writer.
	YAMLWriter Writer = yaml.Marshal
	// Default toml writer.
	TOMLWriter Writer = toml.Marshal
)
