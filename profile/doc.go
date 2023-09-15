// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package profile defines an abstraction layer of types used in the Fit file above the base types (primitive-types) such as sint, uint, etc.
// Here is an example to help understanding it better:
//   - Type DateTime is a time representation decoded in uint32 format in the Fit binary proto. The value of uint32 is a number counted
//     since Fit Epoch (time since 31 Dec 1989 00:00:000 UTC).
//
// Using an abstraction layer like the profile type allows time to be stored in binary files as compact as uint32 values.
// This means that when we encounter a field with the DateTime type, we can decode it into time.Time{}.
package profile
