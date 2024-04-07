# Fitconv

Fitconv converts the FIT file into a CSV file, allowing us to read the FIT file in a human-readable format.

Note:

- Currently only FIT to CSV conversion is supported, other format might be added later when necessary (or requested).

## Usage Examples

```sh
go run main.go activity.fit activity2.fit

# Output:
# Converted! activity.csv
# Converted! activity2.csv

ls
# activity.fit activity.csv activity2.fit activity2.csv
```

### Options

| Options    | Description                                           |
| ---------- | ----------------------------------------------------- |
| -v         | Show version                                          |
| -csv       | Convert FIT to CSV (default if not specified)         |
| -disk      | Use disk instead of load everything in memory         |
| -unknown   | Print 'unknown(68)' instead of 'unknown'              |
| -valid     | Print only valid value and omit invalid value         |
| -raw       | Use raw value instead of scaled value                 |
| -deg       | Print GPS Positions in degress instead of semicircles |
| -trim      | Trim trailing commas in every line (save storage)     |
| -no-expand | [Decode Option] Do not expand components              |

```sh
go run main.go -deg activity.fit activity2.fit
go run main.go -raw activity.fit activity2.fit
go run main.go -deg -no-expand -trim activity.fit activity2.fit
```
