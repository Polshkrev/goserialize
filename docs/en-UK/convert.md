# Convert
This module is strictly a utilities module; defining functions to seperate responsibilities of each of the serialization methods.

#### Functions
The utility functions that this module defines are as follows:
- `SliceToView` - this function is mainly self explanatory. This function converts a built-in go slice into a [gopolutils.Collection](https://github.com/Polshkrev/gopolutils/blob/main/docs/en-UK/collections/collection.md).
- `BytesToObject` - this function is used to convert a slice of bytes representing the raw contents of a file that has been read into memory into an [object](/docs/en-UK/types.md#object) type.