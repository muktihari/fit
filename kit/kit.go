// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kit

// Ptr returns new pointer of v
func Ptr[T any](v T) *T { return &v }
