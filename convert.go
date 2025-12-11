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
// If the stream of bytes can not be unmarshalled, an UnmarshalError is returned with a nil data pointer.
func BytesToObject(content []byte, reader Reader) (Object, *gopolutils.Exception) {
	var object Object = make(Object, 0)
	var unmarshalError error = reader(content, &object)
	if unmarshalError != nil {
		return nil, gopolutils.NewNamedException("UnmarshalError", fmt.Sprintf("Can not read data '%+v': '%s'.", content, unmarshalError.Error()))
	}
	return object, nil
}

// Convert an object type to a stream of bytes.
// Returns a byte slice representing the object.
// If the object can not be marshalled, a MarshallError is returned with a nil data pointer.
func ObjectToBytes(object Object, writer Writer) ([]byte, *gopolutils.Exception) {
	var result []byte
	var marshalError error
	result, marshalError = writer(object)
	if marshalError != nil {
		return nil, gopolutils.NewNamedException("MarhsalError", fmt.Sprintf("Can not write data '%+v': %s", object, marshalError.Error()))
	}
	return result, nil
}

// Convert a stream of bytes to a slice of objects.
// Returns a slice of objects representing the stream of bytes.
// If the stream of bytes can not be marshalled, a MarshalError is returned with a nil data pointer.
func BytesToObjectSlice(content []byte, reader Reader) ([]Object, *gopolutils.Exception) {
	var result []Object = make([]Object, 0)
	var unmarshalError error = reader(content, &result)
	if unmarshalError != nil {
		return nil, gopolutils.NewNamedException("MarshalError", fmt.Sprintf("Can not read data '%+v': %s", content, unmarshalError.Error()))
	}
	return result, nil
}

// Convert a slice of objects to a stream of bytes.
// Returns a stream of bytes representing the slice of objects.
// If the slice of objects can not be marshalled, a MarshallError is returned with a nil data pointer.
func ObjectSliceToBytes(objects []Object, writer Writer) ([]byte, *gopolutils.Exception) {
	var result []byte
	var marshalError error
	result, marshalError = writer(objects)
	if marshalError != nil {
		return nil, gopolutils.NewNamedException("MarshalError", fmt.Sprintf("Can not write data '%+v': %s", objects, marshalError.Error()))
	}
	return result, nil
}
