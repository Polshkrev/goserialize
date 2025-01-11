package goserialize

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

// A simple serializer.
type Serializer[Type any] struct {
	reader Reader
	writer Writer
}

// Construct a new serializer.
// Returns a pointer to a new serializer of a given type with a given reader and writer.
//
// ? This function could probably be renamed in the future to make use of the domain since this is the main entry point of the library and no other constructors are defined at the moment.
func NewSerializer[Type any](reader Reader, writer Writer) *Serializer[Type] {
	var serializer *Serializer[Type] = new(Serializer[Type])
	serializer.reader = reader
	serializer.writer = writer
	return serializer
}

// Serialize an object into a type pointer.
// Returns a reflected type pointer.
// If the given data can not be reflected, either writing or reading, an appropriate error is returned with a nil data pointer.
func (serializer Serializer[Type]) ReadObject(data Object) (*Type, *gopolutils.Exception) {
	var object *Type = new(Type)
	var rawBytes []byte
	var marshalError error
	rawBytes, marshalError = serializer.writer(object)
	if marshalError != nil {
		return nil, gopolutils.NewNamedException("MarshalError", fmt.Sprintf("Can not marshal data '%+v'.", data))
	}
	var unmarshalError error = serializer.reader(rawBytes, object)
	if unmarshalError != nil {
		return nil, gopolutils.NewNamedException("UnmarshalError", fmt.Sprintf("Can not unmarshal data '%+v'.", data))
	}
	return object, nil
}

// Serialize a slice of objects into a typed view.
// Returns a reflected type view.
// If the given data can not be reflected, either writing or reading, an appropriate error is returned with a nil data pointer.
func (serializer Serializer[Type]) ReadList(data []Object) (collections.View[Type], *gopolutils.Exception) {
	var typeList []Type = make([]Type, 0, len(data))
	var rawObject Object
	for _, rawObject = range data {
		var object *Type
		var readError *gopolutils.Exception
		object, readError = serializer.ReadObject(rawObject)
		if readError != nil {
			return nil, readError
		}
		typeList = append(typeList, *object)
	}
	return SliceToView[Type](typeList), nil
}

// Serialize a given data type as an object.
// Returns an object with the reflected type properties.
// If the given data can not be reflected, either writing or reading, an appropriate error is returned with a nil data pointer.
func (serializer Serializer[Type]) WriteObject(data Type) (Object, *gopolutils.Exception) {
	var object Object = make(Object, 0)
	var rawBytes []byte
	var marshalError error
	rawBytes, marshalError = serializer.writer(data)
	if marshalError != nil {
		return nil, gopolutils.NewNamedException("MarshalError", fmt.Sprintf("Can not write object '%+v': '%s'.", data, marshalError.Error()))
	}
	var unmarshalError error = serializer.reader(rawBytes, &object)
	if unmarshalError != nil {
		return nil, gopolutils.NewNamedException("UnmarshalError", fmt.Sprintf("Can not read object '%+v': %s.", data, unmarshalError.Error()))
	}
	return object, nil
}

// Serialize a slice of objects into a typed view.
// Returns a reflected object slice.
// If the given data can not be reflected, either writing or reading, an appropriate error is returned with a nil data pointer.
func (serializer Serializer[Type]) WriteList(data collections.View[Type]) ([]Object, *gopolutils.Exception) {
	var objectList []Object = make([]Object, 0, data.Size())
	var rawObject Type
	for _, rawObject = range data.Collect() {
		var object Object
		var writeError *gopolutils.Exception
		object, writeError = serializer.WriteObject(rawObject)
		if writeError != nil {
			return nil, writeError
		}
		objectList = append(objectList, object)
	}
	return objectList, nil
}
