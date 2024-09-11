# FIT Activity CLI

A program to manage FIT activity files, including combining multiple FIT activity files into a single continuous activity, concealing start and end GPS positions for privacy, and reducing or removing messages.

Available commands:

1. **combine**: combine multiple activities into one continuous activity.
1. **conceal**: conceal first or last x meters GPS positions for privacy.
1. **reduce**: reduce the size of record messages, available methods:
   - Based on GPS points using RDP [Ramer-Douglas-Peucker]
   - Based on distance interval in meters
   - Based on time interval in seconds
1. **remove**: remove messages based on given message numbers and other parameters

This program is available for download in [Release's Assets](https://github.com/muktihari/fit/releases).

TLDR: [Usage](#Usage)

## Background

As an ultra-cyclist, I used to split the activity recording into smaller chunks of FIT files to prevent unexpected events such as a Cyclocomp hang that required a restart, Cyclocomp memory becoming full, Cyclocomp auto power-off due to low battery, and more. The consequence of this was that when I uploaded it to social media platforms like Strava, it appeared as multiple activities when it should have been one continuous activity. So I need to create a program to combine those FIT files. Actually, there is already an online application available for combining activity files, such as [strava tools by gotoes.org](https://gotoes.org/strava/Combine_GPX_TCX_FIT_Files.php). However, I have some concerns about my privacy, as FIT files contain personal data such as GPS positions and health-related information that are uploaded to someone else's server over the internet.

Furthermore, I need to conceal the start and stop distances for up to X kilometers to prevent revealing my exact location to others. Even though platforms like Strava offer this feature, it is limited to a specific number of kilometers, and the data is still retained by Strava.

## Activity File Specification

FIT Activity File Structure Ref: [https://developer.garmin.com/fit/file-types/activity/](https://developer.garmin.com/fit/file-types/activity/)

Strava Specification: [https://developers.strava.com/docs/uploads](https://developers.strava.com/docs/uploads/)

## How We Combine Multiple Files

First, we will order the files by `FileId.TimeCreated`.
The first file will be the base for the resulting file and we will combine all messages from the next FIT files into the resulting file except: **FileId**, **FileCreator**, **Activity**.

The common messages in an Activity File:

- Activity: we use activity message from first FIT file then update it accordingly.
- Session: fields will be aggregated with the correspoding session of the next FIT file.
- Lap: append as it is.
- Event: append as it is
- Record: field `distance` will be accumulated before append, the rest will be appended as it is
- SplitSummary: fields will be aggregated with the split summary of the next FIT file that has the same `split_type`.

The rest of the messages from the next FIT files will be appended as it is.

### Aggregating Fields:

We will aggregate fields depends on the prefix and suffix of the field name:

- prefix **'total**': sum of the two values. (e.g. **total_elapsed_time**)
- prefix **'num**' and suffix **'s**': sum of the two values. (e.g. **num_splits**)
- prefix **'max**' or **'enhanced_max'**: max value between two values. (e.g. **max_heart_rate**, **enhanced_max_respiration_rate**)
- prefix **'min**' or **'enhanced_min'**: min value between two values. (e.g. **min_cadence**, **enhanced_min_altitude**)
- prefix **'avg**' or **'enhanced_avg'**: average of the two values. (e.g. **avg_temperature**, **enhanced_avg_speed**)

Otherwise, they will be assigned with value from the corresponding field only if they are invalid.

### The process

We will combine last session group (include record, event, and lap) of the first file with the first session group of the next file (and so on).

Example:

- File A: swimming, **cycling**
- File B: **cycling**, _running_
- File C: _running_, walking
- Result: swimming, cycling, running, walking

If the last session sport type is not match with the first session of the next file, error will be returned.

Mismatch example:

- File A: swimming, **cycling**
- File B: **running**, walking
- Result: error sport type mismatch

That is the only guard that we implement, it is user responsibility to select the correct files.

_NOTE: Combining FIT activity files is NOT the same as merging multiple files into a single chained FIT file._

## How We Conceal GPS Positions

1. Conceal Start Position
   We will iterate from the beginning of FIT Messages up to the desired conceal distance and for every record found, we will remove the `position_lat` and `position_long` fields.
1. Conceal End Position
   We will backward-iterate from the end of the FIT messages up to the desired conceal distance and for every record found, we will remove the `position_lat` and `position_long` fields.

For Laps and Sessions, we will update these following fields: `start_position_lat`, `start_position_long`, `end_position_lat`, and `end_position_long` to match the new desired positions.

## How We Reduce Record Messages

We reduce record messages based on provided method, you can only select one of the following methods:

1. Based on GPS points using [Ramer–Douglas–Peucker algorithm](https://w.wiki/B6U3) algorithm

   This method simplifies a curve with fewer points, the precision depends on given epsilon (tolerance). Library that we use [github.com/muktihari/carto](https://github.com/muktihari/carto).

1. Based on distance interval in meters

   This method ensures that the distance gap between two records does not exceed the given distance interval. If the gap is already greater than the interval, no records will be removed.

1. Based on time interval in seconds

   This method ensures that the timestamp gap between two records does not exceed the given time interval. If the gap is already greater than the interval, no records will be removed.

The reduced record messages are simply removed; no aggregation is performed.

## How We Remove Messages

Removing messages in straightforward, we will remove messages based on given message numbers. For instance, giving **160** as input will remove **gps_metadata** message ([List os message numbers ](../../profile/typedef/mesg_num_gen.go)). For unknown messages, we consider message that are not defined by Global Profile (Profile.xlsx) as unknown messages. For developer data, we will remove `developer_data_id` and `field_description` messages, additionally all `DeveloperFields` in all messages will be removed as well.

## Install or Build

_Prerequisite: Install golang: [https://go.dev/doc/install](https://go.dev/doc/install)_

### Install

Install the program directly from the remote repository without manually cloning it. The executable will be saved in $GOPATH/bin, enabling you to call it across directory as long as $GOPATH is exported.

```sh
go install github.com/muktihari/fit/cmd/fitactivity@latest
```

#### Install from local folder:

Clone the repository first, then go to "cmd/fitactivity" folder.

```sh
go install .
```

### Manual Build

Clone the repository first, then go to "cmd/fitactivity" folder.

1. MacOS
   ```sh
   GOOS=darwin GOARCH=amd64 go build -o fitactivity main.go
   ```
2. Windows

   ```sh
   GOOS=windows GOARCH=amd64 go build -o fitactivity.exe main.go
   ```

3. Linux
   ```sh
   GOOS=linux GOARCH=amd64 go build -o fitactivity main.go
   ```

More: [How to build Go Executable for Multiple Platforms](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)

**Alternatively, now you can just download the the binary from [Release's Assets](https://github.com/muktihari/fit/releases)**

## Usage

After the program has been built or installed, you can call the fitactivity executable in terminal (_or command prompt, if you are using windows, use fitactivity.exe instead_).

You can start with run the CLI without any argument, it will print help text to guide you using this CLI.

```sh
fitactivity
```

Output:

```sh
About:
  fitactivity is a program to manage FIT files based on provided command.

Usage:
  fitactivity [command]

Available Commands:
  combine    combine multiple activities into one continuous activity
  conceal    conceal first or last x meters GPS positions for privacy
  reduce     reduce the size of record messages, available methods:
             1. Based on GPS points using RDP [Ramer-Douglas-Peucker]
             2. Based on distance interval in meters
             3. Based on time interval in seconds
  remove     remove messages based on given message numbers and other parameters

Flags:
  -h, --help       Print help
  -v, --version    Print version
```

Every command may have subcommands and its own flags, just run the command to print its help. e.g. `fitactivity combine`.

### Combine Multiple FIT Activity files

```sh
fitactivity combine
```

Output:

```sh
About:
  combine multiple activities into one continuous activity

Usage:
  fitactivity combine [subcommands] [flags] [files]

Available Subcommands (optional):
  conceal    conceal first or last x meters GPS positions for privacy
  reduce     reduce the size of record messages, available methods:
             1. Based on GPS points using RDP [Ramer-Douglas-Peucker]
             2. Based on distance interval in meters
             3. Based on time interval in seconds
  remove     remove messages based on given message numbers and other parameters

Flags:
  (required):
  -o, --out  string    combine output file

  (optional):
  -i, --interleave  uint8    max interleave for message definition [valid: 0-15, default: 15]
  -c, --compress    bool     compress timestamp into message header [default: false; this overrides interleave]

Subcommand Flags (only if subcommand is provided):
  conceal: (select at least one)
   --first  uint32    conceal distance: first x meters
   --last   uint32    conceal distance: last x meters

  reduce: (select only one)
   --rdp       float64    reduce method: RDP [Ramer-Douglas-Peucker] based on GPS points, epsilon > 0
   --distance  float64    reduce method: distance interval in meters
   --time      uint32     reduce method: time interval in seconds

  remove: (select at least one)
   --unknown   bool       remove unknown messages
   --nums      string     remove message numbers (value separated by comma)
   --devdata   bool       remove developer data

Examples:
  fitactivity combine -o result.fit part1.fit part2.fit
  fitactivity combine reduce -o result.fit --rdp 0.0001 part1.fit part2.fit
  fitactivity combine conceal -o result.fit --first 1000 part1.fit part2.fit
  fitactivity combine remove -o result.fit --unknown --nums 160,164 part1.fit part2.fit
  fitactivity combine conceal reduce -o result.fit --last 1000 --time 5 part1.fit part2.fit
```

### Conceal GPS Position (Latitude and Longitude)

Note: conceal value is in meters

```sh
fitactivity conceal
```

Output:

```sh
About:
  conceal first or last x meters GPS positions for privacy

Usage:
  fitactivity conceal [flags] [files]

Flags:
  (select at least one):
    --start         uint32     conceal distance: first x meters
    --end           uint32     conceal distance: last x meters

  (optional):
  -i, --interleave  uint8      max interleave for message definition [valid: 0-15, default: 15]
  -c, --compress    bool       compress timestamp into message header [default: false; this overrides interleave]

Examples:
  fitactivity conceal --first 1000 a.fit b.fit
  fitactivity conceal --first 1000 --last 1000 a.fit b.fit
```

### Reduce Record Messages

```sh
fitactivity reduce
```

Output:

```sh
About:
  reduce the size of record messages, available methods:
  1. Based on GPS points using RDP [Ramer-Douglas-Peucker]
  2. Based on distance interval in meters
  3. Based on time interval in seconds


Usage:
  fitactivity reduce [flags] [files]

Flags:
  (select only one):
    --rdp           float64    reduce method: RDP [Ramer-Douglas-Peucker] based on GPS points, epsilon > 0
    --distance      float64    reduce method: distance interval in meters
    --time          uint32     reduce method: time interval in seconds

  (optional):
  -i, --interleave  uint8      max interleave for message definition [valid: 0-15, default: 15]
  -c, --compress    bool       compress timestamp into message header [default: false; this overrides interleave]


Examples:
  fitactivity reduce --rdp 0.0001 a.fit b.fit
  fitactivity reduce --distance 0.5 a.fit b.fit
  fitactivity reduce --time 5 a.fit b.fit
```

### Remove Messages

```sh
fitactivity remove
```

Output:

```sh
About:
  remove messages based on given message numbers and other parameters

Usage:
  fitactivity remove [flags] [files]

Flags:
  (select at least one):
    --unknown   bool       remove unknown messages
    --nums      string     remove message numbers (value separated by comma)
    --devdata   bool       remove developer data

  (optional):
  -i, --interleave  uint8      max interleave for message definition [valid: 0-15, default: 15]
  -c, --compress    bool       compress timestamp into message header [default: false; this overrides interleave]


Examples:
  fitactivity remove --unknown a.fit b.fit
  fitactivity remove --nums 160,162 a.fit b.fit
  fitactivity remove --devdata a.fit b.fit
  fitactivity remove --unknown --nums 160,162 --devdata a.fit b.fit
```
