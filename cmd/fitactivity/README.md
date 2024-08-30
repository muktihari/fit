# FIT Activity CLI

A program to combine multiple FIT (\*.fit) activity files into one continuous activity and conceal its position (Lat & Long at specified distance) for privacy. Available for download in [Release's Assets](https://github.com/muktihari/fit/releases).

TLDR: [Usage](#Usage)

## Background

As an ultra-cyclist, I used to split the activity recording into smaller chunks of FIT files to prevent unexpected events such as a Cyclocomp hang that required a restart, Cyclocomp memory becoming full, Cyclocomp auto power-off due to low battery, and more. The consequence of this was that when I uploaded it to social media platforms like Strava, it appeared as multiple activities when it should have been one continuous activity. So I need to create a program to combine those FIT files. Actually, there is already an online application available for combining activity files, such as [strava tools by gotoes.org](https://gotoes.org/strava/Combine_GPX_TCX_FIT_Files.php). However, I have some concerns about my privacy, as FIT files contain personal data such as GPS positions and health-related information that are uploaded to someone else's server over the internet.

Furthermore, I need to conceal the start and stop distances for up to X kilometers to prevent revealing my exact location to others. Even though platforms like Strava offer this feature, it is limited to a specific number of kilometers, and the data is still retained by Strava.

## Activity File Specification

FIT Activity File Structure Ref: [https://developer.garmin.com/fit/file-types/activity/](https://developer.garmin.com/fit/file-types/activity/)

Strava Specification: [https://developers.strava.com/docs/uploads](https://developers.strava.com/docs/uploads/)

## How We Combine

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
- prefix **'max**': max value between two values. (e.g. **max_heart_rate**)
- prefix **'min**': min value between two values. (e.g. **min_cadence**)
- prefix **'avg**': average of the two values. (e.g. **avg_temperature**)

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
   We will iterate from the beginning of FIT Messages up to the desired conceal distance and for every record found, we will remove the `position_lat` and `position_long` fields. And also, we will update the corresponding session fields: `start_position_lat` and `start_position_long`.
1. Conceal End Position
   We will backward-iterate from the end of the FIT messages up to the desired conceal distance and for every record found, we will remove the `position_lat` and `position_long` fields. And also, we will update the corresponding session fields: `end_position_lat` and `end_position_long`.

We will remove `start_position_lat`, `start_position_long`, `end_position_lat`, and `end_position_long` fields from Laps. But why? GPS Positions saved in lap messages can be vary, user may set new lap every 500m or new lap every 1 hour for example, we don't know the exact distance for each lap. If user want to conceal 1km, we need to find all laps within the conceal distance and decide whether to remove it or change it with new positions, this will add complexity. So, let's just remove it for now, if our upload target is Strava, they don't specify positions in lap message anyway.

## Build or Install

_Prerequisite: Install golang: [https://go.dev/doc/install](https://go.dev/doc/install)_

### Build

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

### Install

Or you can install the program instead of manual build, this will build base on your current OS and CPU architecture and save the executable in $GOPATH/bin, enabling you to call it across directory as long as $GOPATH is exported.

```sh
go install .
```

**Alternatively, now you can just download the the binary from [Release's Assets](https://github.com/muktihari/fit/releases)**

## Usage

After the program has been built or installed, you can call the fitactivity executable in terminal (_or command prompt, if you are using windows, use fitactivity.exe instead_)

### Combine Multiple FIT Activity files

1. Using -o argument
   ```sh
   fitactivity --combine -o result.fit part1.fit part2.fit
   ```
2. Using redirection pipe '>'
   ```sh
   fitactivity --combine part1.fit part2.fit > result.fit
   ```

### Conceal GPS Position (Latitude and Longitude)

Note: conceal value is in meters

1. Conceal GPS position for 1km away from the actual start position for each files.

   ```sh
   fitactivity --conceal-start 1000 file1.fit file2.fit

   # ls output
   # file1.fit file1_concealed_1000_0.fit file2.fit file2_concealed_1000_0.fit
   ```

2. Conceal GPS position for 1km away from the actual end position for each files.

   ```sh
   fitactivity --conceal-end 1000 file1.fit file2.fit

   # ls output
   # file1.fit file1_concealed_0_1000.fit file2.fit file2_concealed_0_1000.fit
   ```

3. Conceal GPS position for 1km away from both start and end position for each files

   ```sh
   fitactivity --conceal-start 1000 --conceal-end 1000 file1.fit file2.fit

   # ls output
   # file1.fit file1_concealed_1000_1000.fit file2.fit file2_concealed_1000_1000.fit
   ```

### Combine Multiple Files and Conceal GPS Position of the Resulting File.

This will combine the FIT activity files, `part1.fit` and `part2.fit`, into one continues FIT activity file `result.fit` and then concealing the start and end GPS position of `result.fit` for 1km away from the actual start and end position.

```sh
fitactivity --combine --conceal-start 1000 --conceal-end 1000 part1.fit part2.fit > result.fit

# ls output
# part1.fit part2.fit result.fit
```

### Available Options

NOTE: We can only use either `interleave` or `compress` at a time.

<table class="table table-bordered table-hover table-condensed">
<thead>
<tr>
    <th style="width: 80px">Option</th>
    <th>Type</th>
    <th>Default</th>
    <th>Description</th>
</tr>
</thead>
<tbody>
<tr>
    <td>-interleave</td>
    <td>uint</td>
    <th>15</th>
    <td>Max interleave allowed to reduce writing the same message definition on encoding process. Valid value: 0-15. <strong>NOTE: If your target is an embedded device, consider using smaller value or just use 0, since its RAM is relatively small.</strong></td>
</tr>
<tr>
    <td>-compress</td>
    <td>boolean</td>
    <th>false</th>
    <td>Compress timestamp in message header. Save 7 bytes per message for every message written in 31s interval.</td>
</tr>
</tbody>
</table>

Example using options:

```sh
fitactivity --combine -o result.fit --interleave 7 part1.fit part2.fit
```

```sh
fitactivity --combine -o result.fit --compress part1.fit part2.fit
```
