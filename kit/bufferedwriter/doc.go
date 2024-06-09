// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bufferedwriter provides functionality to wrap an io.Writer while keep maintaining
// the underlying capability to write at specific bytes such as when it's implementing io.WriterAt or io.WriteSeeker.
// This package differs from bufio.Writer which encapsulates the writer's implementation details as io.Writer.
//
// Deprecated: Encoder now implements built-in write buffer so this package is no longer needed.
// This package will no longer be maintained and might be deleted in the next major/minor release.
package bufferedwriter
