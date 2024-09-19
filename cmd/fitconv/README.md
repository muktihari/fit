# Fitconv

Fitconv converts FIT files to CSV format, enabling us to read the FIT data in a human-readable format. Conversely, it also converts CSV files back to FIT format, enabling us to create or edit FIT files in CSV form.

This is designed to work seamlessly with CSVs produced by the Official FIT SDK's `FitCSVTool.jar`. However, interoperability is not guaranteed to be 100%.

Note:

- Currently, only conversions between FIT and CSV formats are supported. Other formats may be added in the future as needed or upon request.
- When converting from CSV to FIT, any unknown messages and fields are omitted by default, unless the CSV file is created by fitconv using `--verbose` flags (see Usage for details).

This program is available for download in [Release's Assets](https://github.com/muktihari/fit/releases).

## Usage

If you run it from source code, replace "fitconv" with "go run main.go"

```sh
fitconv activity.fit activity2.csv

# Output:
# 📄 "activity.fit" -> "activity.csv"
# 🚀 "activity2.csv" -> "activity2.fit". [Info: 2 unknown messages are skipped]

ls
# activity.fit activity.csv activity2.fit activity2.csv
```

To enable editing unknown data in CSV then parse it back to FIT file format, use `--verbose`:

```sh
fitconv --verbose activity.fit

# Output:
# 📄 "activity.fit" -> "activity.csv"
```

Edit activity.csv, then run:

```sh
fitconv activity.csv

# Output:
# 🚀 "activity.csv" -> "activity.fit".
```

All unknown data successfully converted.

Be careful: the previous command replace existing activity.fit. Ensure you create a backup if needed.

**NOTE: When --verbose is specified, it becomes non-interoperable with `FitCSVTool.jar`, but the unknown data can be converted back to FIT as it was.**

### Build or Install

#### Build

```sh
go build -o fitconv main.go
```

#### Install

```sh
go install .
```

#### Run the resulting Binary

```sh
fitconv activity.fit activity2.csv

# Output:
# 📄 "activity.fit" -> "activity.csv"
# 🚀 "activity2.csv" -> "activity2.fit". [Info: 2 unknown messages are skipped]
```

### Options

| Options       | Valid For       | Description                                              |
| ------------- | --------------- | -------------------------------------------------------- |
| -v, --version | All             | Show version                                             |
| -about        | All             | Show about this program                                  |
| -disk         | FIT to CSV only | Use disk instead of load everything in memory            |
| -verbose      | FIT to CSV only | Print 'unknown(68)' instead of 'unknown'                 |
| -valid        | FIT to CSV only | Print only valid value and omit invalid value            |
| -raw          | FIT to CSV only | Use raw value instead of scaled value                    |
| -deg          | FIT to CSV only | Print GPS Positions in degrees instead of semicircles.   |
| -trim         | FIT to CSV only | Trim trailing commas in every line (save storage)        |
| -no-expand    | FIT to CSV only | [Decode Option] Do not expand components                 |
| -no-checksum  | FIT to CSV only | [Decode Option] Do not check integrity (no CRC checksum) |

```sh
fitconv -deg activity.fit activity2.fit
fitconv -raw activity.fit activity2.fit
fitconv -deg -no-expand -trim activity.fit activity2.fit
```
