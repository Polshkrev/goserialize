# Convert
This module is strictly a utilities module; defining functions to seperate responsibilities of each of the serialization methods.

#### Functions
The utility functions that this module defines are as follows:
- `SliceToView` - this function is mainly self explanatory. This function converts a built-in go slice into a [gopolutils.View](https://github.com/Polshkrev/gopolutils/blob/main/docs/en-UK/collections/view.md).
- `BytesToObject` - this function is used to convert a stream of bytes into an [object](/docs/en-UK/types.md#object). If the stream of bytes can not be unmarshalled, an `UnmarshalError` is returned with a nil data pointer.
- `ObjectToBytes` - this function is the opposite of `BytesToObject`. This function takes in an object and converts the object into a stream of bytes. If the object can not be marshalled, a `MarshalError` is returned with a nil data pointer.
- `BytesToObjectSlice` - this function takes in a stream of bytes to convert to a slice of objects. If the stream of bytes can not be unmarshalled, an `UnmarshalError` is returned with a nil data pointer.
- `ObjectSliceToBytes` - this function takes in a slice of objects to covert to a stream of bytes. If the slice of objects can not be marshalled, a `MarshalError` is returned with a nil data pointer.