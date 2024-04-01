# FIT SDK for Go

![GitHub Workflow Status](https://github.com/muktihari/fit/workflows/CI/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/muktihari/fit.svg)](https://pkg.go.dev/github.com/muktihari/fit)
[![CodeCov](https://codecov.io/gh/muktihari/fit/branch/master/graph/badge.svg)](https://codecov.io/gh/muktihari/fit)
[![Go Report Card](https://goreportcard.com/badge/github.com/muktihari/fit)](https://goreportcard.com/report/github.com/muktihari/fit)
[![Profile Version](https://img.shields.io/badge/profile-v21.132-lightblue.svg?style=flat)](https://developer.garmin.com/fit/download)

The Flexible and Interoperable Data Transfer (FIT) protocol is a protocol developed by Garmin for storing and sharing data originating from sports, fitness, and health devices.
When recording an activity using devices such as cycling computer, smartwatch, and similar devices, chances are the resulting file is often in FIT file format (\*.fit).
The FIT file is a binary file format known for its compact size, making it the preferred choice for manufacturers to use in their embedded devices.
However, despite having gained widespread adoption, Garmin has not yet released an official SDK for Go, this is where this SDK comes in to bridge the gap, enabling Go developers to be able to interact with FIT file format.

More about FIT: [developer.garmin.com/fit](https://developer.garmin.com/fit)

This SDK is designed with efficiency in mind, but it places a higher priority on clarity, simplicity and extensibility. While other Go implementations for decoding or encoding FIT files may exist, we offer greater correctness, completeness, and similar semantics to the Official SDK.

## Usage

Please see [Usage](/docs/usage.md).

## Protocol Version 2.0 is supported

Version 2.0 introduced [**Developer Fields**](https://developer.garmin.com/fit/cookbook/developer-data) as a way to add custom data fields to existing messages. We strives to support **Developer Fields** and carefully thought about how to implement it since the inception of the SDK. While this may still need to be battle-tested to ensure correctness, this is generally work and usable.

Here is the sample of what **Developer Fields** would look like in a **.fit** that have been converted to **.csv** by `fitconv`. The **device_info** message has some **Developer Fields** defined in **field_description** (the ones that are being bold):

<table class="table table-bordered table-hover table-condensed">
<thead>
<tr>
    <th>Type</th>
    <th>Local Number</th>
    <th>Message</th>
    <th>Field 1</th>
    <th>Value 1</th>
    <th>Units 1</th>
    <th>Field 2</th>
    <th>Value 2</th>
    <th>Units 2</th>
    <th>Field 3</th>
    <th>Value 3</th>
    <th>Units 3</th>
    <th>Field 4</th>
    <th>Value 4</th>
    <th>Units 4</th>
    <th>Field 5</th>
    <th>Value 5</th>
    <th>Units 5</th>
    <th>Field 6</th>
    <th>Value 6</th>
    <th>Units 6</th>
</tr>
</thead>
<tbody>
<tr>
    <td>Definition</td>
    <td>0</td>
    <td>developer_data_id</td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>application_id</td>
    <td>16</td>
    <td> </td>
    <td>application_version</td>
    <td>1</td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>developer_data_id</td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>application_id</td>
    <td>&lt;omitted&gt;</td>
    <td> </td>
    <td>application_version</td>
    <td>40113</td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Definition</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>1</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>1</td>
    <td></td>
    <td>field_name</td>
    <td>13</td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>7</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>5</td>
    <td></td>
    <td>field_name</td>
    <td><strong>device_model</strong></td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Definition</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>1</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>1</td>
    <td></td>
    <td>field_name</td>
    <td>20</td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>7</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>4</td>
    <td></td>
    <td>field_name</td>
    <td><strong>device_manufacturer</strong></td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>7</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>6</td>
    <td></td>
    <td>field_name</td>
    <td><strong>device_os_version</strong></td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>7</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>7</td>
    <td></td>
    <td>field_name</td>
    <td><strong>mobile_app_version</strong></td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Definition</td>
    <td>0</td>
    <td>device_info</td>
    <td>manufacturer</td>
    <td>1</td>
    <td> </td>
    <td>product</td>
    <td>1</td>
    <td> </td>
    <td>device_model</td>
    <td>11</td>
    <td></td>
    <td>device_manufacturer</td>
    <td>6</td>
    <td> </td>
    <td>device_os_version</td>
    <td>5</td>
    <td></td>
    <td>mobile_app_version</td>
    <td>8</td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>device_info</td>
    <td>manufacturer</td>
    <td>265</td>
    <td> </td>
    <td>product</td>
    <td>101</td>
    <td> </td>
    <td><strong>device_model</strong></td>
    <td><strong>iPhone14,4</strong></td>
    <td> </td>
    <td><strong>device_manufacturer</strong></td>
    <td><strong>apple</strong></td>
    <td> </td>
    <td><strong>device_os_version</strong></td>
    <td><strong>16</strong>.6</td>
    <td> </td>
    <td><strong>mobile_app_version</strong></td>
    <td><strong>332.0.0</strong></td>
    <td> </td>
</tr>
</tbody>
</table>

## CLIs

We provide some CLI programs to interact with FIT files that can be found in `cmd` folder.

1. `fitactivity`: Combines multiple FIT activity files into one continuous FIT activity (and conceal the start and end GPS positions for privacy). [README.md](/cmd/fitactivity/README.md)
2. `fitconv`: Converts FIT file to a CSV file, allowing us to read the FIT file in a human-readable format. [README.md](/cmd/fitconv/README.md)

The programs are automatically built during release; for Linux, Windows, and macOS platforms. They are available for download in [Release's Assets](https://github.com/muktihari/fit/releases).

## Custom FIT SDK

A FIT file may contain manufacturer specific messages (`Product Profile`) that are not defined in `Global Profile (Profile.xlsx)` since it's specific to a manufacturer (other manufacturers may have different meaning for that messages)

To be able to decode or create the manufacturer specific messages, we provide options to pick based on your need:

1. Register manufacturer specific messages at runtime

   For those who prefer using this SDK as it is without need to generate their own custom SDK, we provide `factory` package as an abstraction to hold the profile messages. For example, please see [Register Manufacturer Specific Types & Messages at Runtime](/docs/runtime_registration.md).

2. Generate custom FIT SDK

   Please see [Generate Custom FIT SDK](/docs/generating_code.md#Generate-Custom-FIT-SDK)

## Contributing

Please see [CONTRIBUTING.md](/CONTRIBUTING.md).
Thank you, [contributors](https://github.com/muktihari/fit/graphs/contributors)!
