package goserialize

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

// Convert a slice to a view.
// Returns a new view containing the elements of the given slice.
func SliceToView[Type any](items []Type) collections.View[Type] {
	var item Type
	var result collections.Collection[Type] = collections.NewArray[Type]()
	for _, item = range items {
		result.Append(item)
	}
	return result
}

// Convert a stream of bytes to an object type.
// Returns an object type representing the stream of bytes.
func BytesToObject(content []byte, reader Reader) (Object, *gopolutils.Exception) {
	var object Object = make(Object, 0)
	var unmarshalError error = reader(content, &object)
	if unmarshalError != nil {
		return nil, gopolutils.NewNamedException("UnmarshalError", fmt.Sprintf("Can not read data '%+v': '%s'.", content, unmarshalError.Error()))
	}
	return object, nil
}
