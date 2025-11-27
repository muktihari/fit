// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"bytes"
)

// lru implements simple lru algorithm. Item search best case: O(1), worst case: O(n), depends how recently used is it.
type lru struct {
	// Holds the actual items representated in bytes.
	items [][]byte

	// Holds items's indexes as its value. The order of value will be shifted based on recent write.
	// Example order: [0 (least recently used), 1, 2, 3, ..., 15 (most recently used)]
	bucket []byte
}

// newLRU creates new lru with fixed size, where size should be > 0.
func newLRU(size byte) *lru {
	l := &lru{}
	l.ResetWithNewSize(size)
	return l
}

// Reset reset variables so lru can be reused again without reallocation.
func (l *lru) Reset() {
	l.bucket = l.bucket[:0]
}

// ResetWithNewSize sets new LRU size and then reset the LRU. If the new size is more than previous size it will re-allocs new storage
// with the new capacity. If the new size is less than previous size it will reslice without re-allocs. Otherwise, only reset.
func (l *lru) ResetWithNewSize(size byte) {
	if size > byte(cap(l.items)) {
		oldItems := l.items[:cap(l.items)]
		l.items = make([][]byte, size)
		copy(l.items, oldItems) // preserve old storage
		l.bucket = make([]byte, 0, size)
		return
	}
	l.bucket = l.bucket[:0]
	l.items = l.items[:size]
}

// Put will compare the equality of item with lru' items and store the item accordingly.
func (l *lru) Put(item []byte) (itemIndex byte, isNewItem bool) {
	if bucketIndex := l.bucketIndex(item); bucketIndex != -1 {
		return l.markAsRecentlyUsed(bucketIndex), false
	}
	if len(l.bucket) == len(l.items) {
		return l.replaceLeastRecentlyUsed(item), true
	}
	itemIndex = byte(len(l.bucket))
	l.bucket = append(l.bucket, itemIndex)
	l.storeAt(item, itemIndex)
	return itemIndex, true
}

func (l *lru) markAsRecentlyUsed(bucketIndex int) (itemIndex byte) {
	itemIndex = l.bucket[bucketIndex]
	l.bucket = append(l.bucket[:bucketIndex], l.bucket[bucketIndex+1:]...) // splice bucketIndex from the bucket
	l.bucket = append(l.bucket, itemIndex)                                 // place at most recent index
	return
}

func (l *lru) replaceLeastRecentlyUsed(item []byte) (itemIndex byte) {
	itemIndex = l.bucket[0]                        // take item's index out of bucket
	copy(l.bucket[:len(l.bucket)-1], l.bucket[1:]) // left shift bucket
	l.bucket[len(l.bucket)-1] = itemIndex          // place at most recent index
	l.storeAt(item, itemIndex)
	return
}

func (l *lru) storeAt(item []byte, itemIndex byte) {
	// PERF: Only alloc when not enough capacity
	if cap(l.items[itemIndex]) < len(item) {
		l.items[itemIndex] = make([]byte, len(item))
	} else {
		l.items[itemIndex] = l.items[itemIndex][:len(item)]
	}
	copy(l.items[itemIndex], item)
}

func (l *lru) bucketIndex(item []byte) int {
	for i := len(l.bucket) - 1; i >= 0; i-- {
		cur := l.bucket[i]
		if bytes.Equal(l.items[cur], item) {
			return i
		}
	}
	return -1
}
