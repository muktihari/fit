# Fitprint

Fitconv prints the FIT file into stdout in human-readable format.

## Usage Examples

```sh
go run main.go ../../testdata/from_official_sdk/Settings.fit

# Output:
# file_id:
#     manufacturer: 1
#     product: 988
#     serial_number: 123456
#     type: 2
# user_profile:
#     weight: 90 kg
#     gender: 1
#     age: 28 years
#     height: 1.9 m
#     language: 0
# hrm_profile:
#     hrm_ant_id: 100
#
# took: 426.537µs
```
