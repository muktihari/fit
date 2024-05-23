# FIT Dump CLI

Fitdump dumps the FIT file(s) into segmented bytes.

```txt
SEGMENT      |LOCAL NUM|  HEADER    BYTES
file_header                         [14 32 133 82 230 44 4 0 46 70 73 84 12 58]
message_definition  | 0|  01000000  [64 0 1 0 0 7 3 4 140 4 4 134 7 4 134 1 2 132 2 2 132 5 2 132 0 1 0]
message_data        | 0|  00000000  [0 203 230 191 170 63 85 233 108 255 255 255 255 0 1 15 152 255 255 4]
message_definition  | 1|  01000001  [65 0 1 0 49 5 2 20 7 0 2 132 1 1 2 3 1 0 4 1 0]
message_data        | 1|  00000001  [1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 6 166 255 255 255]
...
```

## Usage

### Run

```sh
go run main.go triathlon_summary_last.fit

# FIT dumped: "triathlon_summary_last.fit" -> "triathlon_summary_last-fitdump.txt"
```

```sh
go run main.go -hex triathlon_summary_last.fit
```

### Build or Install

#### Build

```sh
go build -o fitdump main.go
```

#### Install

```sh
go install .
```

#### Run the resulting Binary

```sh
fitdump triathlon_summary_last.fit

# Output:
# FIT dumped: "triathlon_summary_last.fit" -> "triathlon_summary_last-fitdump.txt"
```

### Options

| Options | Description                |
| ------- | -------------------------- |
| -hex    | Print bytes in hexadecimal |

```sh
go run main.go -hex triathlon_summary_last.fit
fitdump -hex triathlon_summary_last.fit
```
