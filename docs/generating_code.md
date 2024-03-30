# Generating Code

This SDK relies heavily on code generation, as the specification is defined in the `Profile.xlsx` file provided by Garmin in their Official SDK release. This document serves as guidance for generating the files, outlines considerations before generation, and provides recovery steps in case issues arise during code generation.

We only generate `Global Profile` specified in `Profile.xlsx` retrieved from the Official SDK. For generating your own custom SDK containing `Product Profile`, please see [How to customize the Fit SDK](#How-to-customize-the-Fit-SDK)

## Conventions to follow

1. [Go cmd][Go cmd] says: To convey to humans and machine tools that code is generated, generated source should have a line that matches the following regular expression (in Go syntax):
   ```regex
   ^// Code generated .* DO NOT EDIT\.$
   ```
1. While there is no naming convention for generated go files, we use the suffix `_gen` for all files created by `fitgen` to easily distinguish code-generated files from others. For example: `activity_gen.go`
1. Items not adequately defined in `Profile.xlsx` or lacking sufficient details are encouraged to code manually when possible to enhance robustness. For example: `Fit Base Type` segregated into `profile/basetype/basetype.go`

## How we generate

When Garmin releases a new SDK, we review the changes they've made. If they are trivial, we update the `Profile.xlsx` from the latest release to trigger CI to generate the files. For major changes, we open an [issue][issues] to have discussions before proceeding further to ensure the stability of our SDK.

After generating code we make sure all existing tests are passes locally and on CI. In case the code breaks, we should revert first and fix the issue locally. The fix can be submitted later using [pull requests][prs].

[Go cmd]: https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source
[issues]: https://github.com/muktihari/fit/issues
[prs]: https://github.com/muktihari/fit/pulls

## Generate Custom FIT SDK

FIT is designed to have the ability to include a `Product Profile`, which contains manufacturer-specific specifications not defined in the `Global Profile`. The `Product Profile` declares manufacturer-specific types and messages used in their FIT-generated files.

To work with these files perfectly, we must include the specification in `Profile.xlsx` and then generate the code. Otherwise, if we encounter a manufactuer specific message, it will be an unknown message.

Here are the steps to follow:

1. Clone this repository to your local machine.
1. Make a copy of `Profile.xlsx` (the file retrieved from official Fit SDK), let's say `Profile-copy.xlsx`.
1. Add manufacturer specific types and messages to `Profile-copy.xlsx` file.
1. Remember, modifying official specifications is prohibited, you can only add/edit your own specifications.
1. Go to fit/internal/cmd/fitgen, run the CLI and point it to the updated `Profile-copy.xlsx`. Example:

   ```sh
   go run main.go --profile-file Profile-copy.xlsx --path ../../../../ --builders all --profile-version 21.115 --verbose -y
   ```

   The path `../../../../` is targeting the root of fit directory (github.com/muktihari/fit).

   For details of how to use the CLI, please do:

   ```sh
   go run main.go --help
   ```
