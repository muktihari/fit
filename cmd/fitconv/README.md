# Fitconv

Fitconv converts FIT files to CSV format, enabling us to read the FIT data in a human-readable format. Conversely, it also converts CSV files back to FIT format, enabling us to create or edit FIT files in CSV form.

This is designed to work seamlessly with CSVs produced by the Official FIT SDK's `FitCSVTool.jar`.

Note:

- Currently, only conversions between FIT and CSV formats are supported. Other formats may be added in the future as needed or upon request.
- When converting from CSV to FIT, any unknown messages and fields are omitted due to the inability to ascertain their correct value types without additional context.

## Usage Examples

```sh
go run main.go activity.fit activity2.csv

# Output:
# 📄 "activity.fit" -> "activity.csv"
# 🚀 "activity2.csv" -> "activity2.fit". [Info: 2 unknown messages are skipped]

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
| -deg       | FIT to CSV only | Print GPS Positions in degrees instead of semicircles. |
| -trim      | FIT to CSV only | Trim trailing commas in every line (save storage)      |
| -no-expand | FIT to CSV only | [Decode Option] Do not expand components               |

```sh
go run main.go -deg activity.fit activity2.fit
go run main.go -raw activity.fit activity2.fit
go run main.go -deg -no-expand -trim activity.fit activity2.fit
```
