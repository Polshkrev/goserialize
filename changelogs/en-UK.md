# Changelog
## v0.9.0 - 2026-02-11
`Chnaged`
- All serialization functions and methods are now concurrent-safe.

## v0.8.0 - 2026-02-07
`Added`
- `types`
    - Added `ObjectList` type.

`Removed`
- `types`
    - Removed all deprecated types.
## v0.7.0 - 2026-02-06
`Added`
- A new changelog has appeared.
- `types`
    - Added `JSONIndentWriter`.

`Deprecated`
- `types`
    - Added deprecation notices for `JSONType`, `YAMLType`, `TOMLType`, `CSVType`, due to the introduction of the [Suffix enum in the github.com/Polshkrev/fayl](https://github.com/Polshkrev/gopolutils/blob/main/fayl/sufixes.go#L11) dependency.
## v0.6.0 - 2026-01-16
`Added`
- `types`
    - Added deprecation comments for each of the `types`.
## v0.5.0 - 2025-12-11: The CSV Update
`Added`
- Added `github.com/trimmer-io/go-csv` dependency.
- `types`
    - Added `CSVType`.
    - Added `CSVReader`.
    - Added `CSVWriter`.
## v0.4.0 - 2025-01-15: The Conversion Update
`Added`
- `convert`
    - Added `ObjectToBytes` function.
    - Added `BytesToObjectSlice` function.
    - Added `ObjectSliceToBytes` function.
## v0.3.3 - 2025-01-11
`Fixed`
- `types`
    - Fixed some minor formatting changes.
## v0.3.2 - 2025-01-11
`Added`
- `Docs`
    - Added main documentation in `README.md`.
    - Added the `convert` module documentation.
    - Added documentation for the `Serializer` type.
    - Added documentation for the `types` module.

`Changed`
- `Docs`
    - All english documentation is now in a seperate localized folder.

`Fixed`
- `serializer`
    - In the `WriteObject` method, instead of returning a `MarshalError` upon an error state after reading, an `UnmarshalError` is returned.
## v0.3.1 - 2025-01-11
`Added`
- Added `meta.toml`.

`Fixed`
- `serializer`
    - In the `WriteList` method, the size of the passed in view is used to initialize the initial size of the result object slice.
## v0.3.0 - 2024-12-26
`Added`
- `serializer`
    - Added `ReadObject` method.
    - Added `ReadList` method.
    - Added `WriteObject` method.
    - Added `WriteList` method.
## v0.2.0 - 2024-12-26
`Added`
- `convert`
    - Added `SliceToView` function.
    - Added `BytesToObject` function.
- `serializer`
    - Added `github.com/Polshkrev/gopolutils` dependency.
    - Added `Serializer` type.
    - Added `Serializer` constructor.
## v0.1.0 - 2024-12-26
`Added`
- `types`
    - Added `github.com/BurntSushi/toml` dependency.
    - Added `gopkg.in/yaml.v2` dependency.
    - Added `Object` type.
    - Added `JSONType`.
    - Added `YAMLType`.
    - Added `TOMLType`.
    - Added `Writer` type alias.
    - Added `Reader` type alias.
    - Added `JSONReader`.
    - Added `YAMLReader`.
    - Added `TOMLReader`.
    - Added `JSONWriter`.
    - Added `YAMLWriter`.
    - Added `TOMLWriter`.