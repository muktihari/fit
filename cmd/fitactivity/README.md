# FIT Activity CLI

A program to combine and conceal position (Lat & Long at specified distance) of FIT Activity Files.

## Background

As an ultra-cyclist, I used to split the activity recording into smaller chunks of FIT files to prevent unexpected events such as a Cyclocomp hang that required a restart, Cyclocomp memory becoming full, Cyclocomp auto power-off due to low battery, and more. The consequence of this was that when I uploaded it to social media platforms like Strava, it appeared as multiple activities when it should have been one continuous activity. So I need to create a program to combine those FIT files. Actually, there is already an online application available for combining activity files, such as [strava tools by gotoes.org](https://gotoes.org/strava/Combine_GPX_TCX_FIT_Files.php). However, I have some concerns about my privacy, as FIT files contain personal data such as GPS positions and health-related information that are uploaded to someone else's server over the internet.

Furthermore, I need to conceal the start and stop distances for up to X kilometers to prevent revealing my exact location to others. Even though platforms like Strava offer this feature, it is limited to a specific number of kilometers, and the data is still retained by Strava.

## Activity File Specification

FIT Activity File Structure Ref: [https://developer.garmin.com/fit/file-types/activity/](https://developer.garmin.com/fit/file-types/activity/)

Strava Specification: [https://developers.strava.com/docs/uploads](https://developers.strava.com/docs/uploads/)

## How We Combine

First, we will order the files by `FileId.TimeCreated`.
The first file will be the base for the resulting file and we will combine these following messages from the next FIT files into the resulting file:

- Session: combine session by calculating some fields (list fields will be shown after this)
- Record: field `distance` will be calculated before append, the rest will be appended as it is
- Event: append as it is
- Lap: field `start_position_lat`, `start_position_long`, `end_position_lat`, and `end_position_long` will be removed only if conceal option is specified, the rest will be appended as it is.

  Why lap positions must be removed? GPS Positions saved in lap messages can be vary, user may set new lap every 500m or new lap every 1 hour for example, we don't know the exact distance for each lap. If user want to conceal 1km, we need to find all laps within the conceal distance and decide whether to remove it or change it with new positions, this will add complexity. So, let's just remove it for now, if our upload target is Strava, they don't specify positions in lap message anyway.

Other messages from the next FIT files will not be combined.

### Calculated Session Fields:

Currently we only care these following session fields:

- sport (is used to match two sessions)
- sub_sport (is not used since different devices may have different value)
- start_time (is used to calculate time gap between two sessions, add time gap to total_elapsed_time)
- end_position_lat (will be replaced with next files session's end_position_lat)
- end_position_long (will be replaced with next files session's end_position_long)
- total_elapsed_time
- total_timer_time
- total_distance
- total_ascent
- total_descent
- total_cycles
- total_calories
- avg_speed
- max_speed
- avg_heart_rate
- max_heart_rate
- avg_cadence
- max_cadence
- avg_power
- max_power
- avg_temperature
- max_temperature
- avg_altitude
- max_altitude

### Combine process

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

## Build or Install

_Prerequisite: Install golang: [https://go.dev/doc/install](https://go.dev/doc/install)_

### Build

1. MacOS
   ```sh
   GOOS=darwin GOARCH=amd64 go build -o fitactivity main.go
   ```
2. Windows

   ```sh
   GOOS=windows GOARCH=amd64 go build -o fitactivity main.go
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

## Usage Examples

After the program has been built or installed, you can call the fitactivity executable.

### Combine Multiple FIT files

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
