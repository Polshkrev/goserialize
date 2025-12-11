# Types
This module defines some types that standardize the types used in the [serializer](/docs/en-UK/serializer.md) module.
The only serialization types that supported at the moment are `JSON`, `YAML`, `TOML`, and `CSV`.

#### Object
This module defines a type alias that will be familiar to any javascript developers. An object, is essentially a json object. This type alias uses a mapping of strings to an `any` type due to that being the most flexable way of describing what data type can be serialized.

#### Constants
The constants that are defined in this module are the file extensions that can be written and read from.
This module also defines a reader and a writer type to standardize each of the functions that marshal and unmarshal data.
Each of the file extensions that are supported at the moment have a marshal and unmarshal function defined that is type aliased as a reader and writer respectively.