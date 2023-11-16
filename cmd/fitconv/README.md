# Fitconv

Fitconv converts the FIT file into a CSV file, allowing us to read the FIT file in a human-readable format.

Note:

- Currently only Convert FIT to CSV format is supported, other format might be added later when necessary.

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

| Options             | Description                                   |
| ------------------- | --------------------------------------------- |
| -f2c                | Convert FIT to CSV (default if not specified) |
| -f2c-raw-value      | Use raw value instead of scaled value         |
| -f2c-unknown-number | Print 'unknown(68)' instead of 'unknown'      |
| -f2c-use-disk       | Use disk instead of load everything in memory |
| -v                  | Show version                                  |
