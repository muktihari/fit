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
- Record: field distance will be calculated before append, the rest will be appended as it is
- Event: append as it is
- Lap: append as it is

Other messages from the next FIT files will not be combined.

### Calculated Session Fields:

Currently we only care these following session fields:

- sport
- sub_sport
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

- File A: swimming, cycling
- File B: cycling, running
- File C: running, walking
- Result: swimming, cycling, running, walking

If the last session sport type is not match with the first session of the next file, error will be returned.

Mismatch example:

- File A: cycling
- File B: running
- Result: error sport type mismatch

That is the only guard that we implement, it is user responsibility to select the correct files.

## Build or Install

> Prerequisite: Install golang: [https://go.dev/doc/install](https://go.dev/doc/install)

### Build

1. MacOS
   ```sh
   GOOS=darwin GOARCH=amd64 go build -out fitactivity main.go
   ```
2. Windows

   ```sh
   GOOS=windows GOARCH=amd64 go build -out fitactivity main.go
   ```

3. Linux
   ```sh
   GOOS=linux GOARCH=amd64 go build -out fitactivity main.go
   ```

More: [How to build Go Executable for Multiple Platforms](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)

### Install

Or you can install the program instead of manual build, this will build base on your current OS and CPU architecture and save the executable in $GOPATH/bin, enabling you to call it across directory as long as $GOPATH is exported.

```sh
go install .
```

## Usage Examples

After the program has been built or installed, you can call the fitactivity executable.

### Combine FIT files

1. Using -o argument
   ```sh
   fitactivity --combine -o result.fit part1.fit, part2.fit
   ```
2. Using redirection pipe '>'
   ```sh
   fitactivity --combine part1.fit part2.fit > result.fit
   ```

### Conceal Position (Latitude and Longitude)

1. Conceal position for 1km from the start position for each files.

   Results: `part1_concealed_1000_0.fit`, `part2_concealed_1000_0.fit`

   ```sh
   fitactivity --conceal-start 1000 part1.fit part2.fit
   ```

2. Conceal position for 1km from the end position for each files.

   Results: `part1_concealed_0_1000.fit`, `part2_concealed_0_1000.fit`

   ```sh
   fitactivity --conceal-end 1000 part1.fit part2.fit
   ```

3. Conceal position for 1km from both start and end position for each files

   Results: part1_concealed_1000_1000.fit part2_concealed_1000_1000.fit

   ```sh
   fitactivity --conceal-start 1000 --conceal-end 1000 part1.fit part2.fit
   ```

### Combine and Conceal Position

This will combine the activity and then conceal position of the resulting `result.fit`

```sh
fitactivity --combine --conceal-start 1000 --conceal-end 1000 part1.fit part2.fit > result.fit
```
