# FIT SDK for Go Generator

The FIT SDK for Go Generator, also known as "fitgen", is a program designed to create several \*.go files using `Profile.xlsx`, file retrieved from the Official FIT SDK release. The generated files enable this FIT SDK for Go to carry out the decoding and encoding process of FIT files.

The files are organized into distinct packages:

- profile: mesgdef, typedef, untyped
- factory

To define your manufacturer specifications, duplicate the `Profile.xlsx` file and incorporate your specifications within it. Afterward, utilize the provided command-line interface (CLI) to generate a customized SDK. When executing the CLI command, specify the path to the edited-file such as `Profile-copy.xlsx` using the "--path" option.

Example:

- "./fitgen --profile-file Profile-copy.xlsx --path ../../../../ --builders all --profile-version 21.115 -v -y"
- "./fitgen -f Profile-copy.xlsx -p ../../ -b all --profile-version 21.115 -v -y"

Note: The existing Garmin SDK specifications must not be altered, as such modifications could lead to data that does not align with the terms and conditions of the FIT Protocol.
