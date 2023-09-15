// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package byteorder

import "encoding/binary"

func Select(arch byte) binary.ByteOrder {
	if arch == 0 {
		return binary.LittleEndian
	}
	return binary.BigEndian
}
