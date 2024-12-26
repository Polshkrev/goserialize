package goserialize

// A simple serializer.
type Serializer[Type any] struct {
	reader Reader
	writer Writer
}

// Construct a new serializer.
// Returns a pointer to a new serializer of a given type with a given reader and writer.
func NewSerializer[Type any](reader Reader, writer Writer) *Serializer[Type] {
	var serializer *Serializer[Type] = new(Serializer[Type])
	serializer.reader = reader
	serializer.writer = writer
	return serializer
}
