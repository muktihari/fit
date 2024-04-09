# Fitconv

Fitconv converts the FIT file into a CSV file, allowing us to read the FIT file in a human-readable format. And vice-versa, it converts CSV file back to FIT file allowing us editing FIT file in its CSV form using only code editor.

This is designed to work seamlessly with CSVs produced by the Official FIT SDK's `FitCSVTool.jar`.

Note:

- Currently only `FIT to CSV` conversion and vice-versa (`CSV to FIT`) are supported, other format might be added later when necessary (or requested).
- When converting `CSV to FIT`, unknown message and unknown fields are both skipped since we can't get the correct value's type since we don't have any context about it.

## Usage Examples

```sh
go run main.go activity.fit activity2.csv

# Output:
# 📄 "activity.fit" -> "activity.csv"
# 🚀 "activity2.csv" -> "activity2.fit". [In total, 2 unknown messages are skipped]

ls
# activity.fit activity.csv activity2.fit activity2.csv
```

### Options

| Options    | Valid For       | Description                                            |
| ---------- | --------------- | ------------------------------------------------------ |
| -v         | All             | Show version                                           |
| -disk      | FIT to CSV only | Use disk instead of load everything in memory          |
| -unknown   | FIT to CSV only | Print 'unknown(68)' instead of 'unknown'               |
| -valid     | FIT to CSV only | Print only valid value and omit invalid value          |
| -raw       | FIT to CSV only | Use raw value instead of scaled value                  |
| -deg       | FIT to CSV only | Print GPS Positions in degress instead of semicircles. |
| -trim      | FIT to CSV only | Trim trailing commas in every line (save storage)      |
| -no-expand | FIT to CSV only | [Decode Option] Do not expand components               |

```sh
go run main.go -deg activity.fit activity2.fit
go run main.go -raw activity.fit activity2.fit
go run main.go -deg -no-expand -trim activity.fit activity2.fit
```
