# Serializer
This module is the actual workhorse of this library. This module defines a `Serializer` type that can use its reader and writer properties to convert custom **reflectable** types into a built-in go maps.

## Table of Contents
1. [Usage](#usage)
    1. [Construction](#construction)
    2. [Read](#read)
    3. [Write](#write)

### Usage
#### Construction
To constuct a new serializer, there is only one method defined: `NewSerializer`. This function could simply be renamed to `New` in future. This function uses injects both the reader and writer dependency into the serializer.

#### Read
There are two methods defined for reading. To unmarshal an object into a custom type, both methods `ReadList` and `ReadObject` are defined. As the names of each would suggest, each respective method either marshals a slice or a single object as a slice of custom types or a single custom type.
`ReadObject` returns a pointer to the serialized type. If the object passed into the method can not be serialized, either a `MarshalError` or an `UnmarshalError` is returned with a `nil` data pointer.
`ReadList` returns a [collections.View](https://github.com/Polshkrev/gopolutils/blob/main/docs/en-UK/collections/view.md) of a custom type. As this method uses `ReadObject` to serialize each object, all error handling that is used in that method applies to this method.

#### Write
There are two method defined for writing. To marshal an custom type into an object, both methods `WriteList` and `WriteObject` are defined. As the names of each would suggest, each respective method either marshals a `collections.View` of a custom type or a single custom type as a slice of objects or a single object.
`WriteObject` takes in a single custom type and returns an object. If the custom type passed into the method can not be serialized, either a `MarshalError` or an `UnmarshalError` is returned with a `nil` data pointer.
`WriteList` takes in a `collections.View` of a custom type and returns a slice of a objects. As this method uses `WriteObject` to serialize each custom type, all error handling that is used in that method applies to this method.