# Fitconv

Fitconv converts the FIT file into a CSV file, allowing us to read the FIT file in a human-readable format.

Note: Currently only Convert to CSV format is supported, other format might be added later when necessary.

## Usage Examples

```sh
go run main.go --path activity.fit

# Output:
# Converted! activity.csv

ls
# activity.fit activity.csv
```
