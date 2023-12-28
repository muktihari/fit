// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoder

import (
	"testing"

	"github.com/muktihari/fit/proto"
)

func TestLRU(t *testing.T) {
	const size uint8 = proto.LocalMesgNumMask + 1
	l := newLRU(size)

	// place (size * 10) different items, the lru will be shifted in roundroubin order.
	for i := byte(0); i < size*10; i++ {
		localMesgNum, isNew := l.Put([]byte{i})
		if localMesgNum != i%size {
			t.Fatalf("expected: %d, got: %d", i, localMesgNum)
		}
		isNewExpected := true
		if isNew != isNewExpected {
			t.Fatalf("expected: %t, got: %t", isNewExpected, isNew)
		}
	}

	// put same items should shift the lru bucket
	for i := byte(0); i < size; i++ {
		item := l.items[i]
		localMesgNUm, _ := l.Put(item)
		if localMesgNUm != i {
			t.Fatalf("expected: %d, got: %d", i, localMesgNUm)
		}
		if l.bucket[size-1] != i {
			t.Fatalf("expected: %d, got: %d", i, l.bucket[size-1])
		}
	}

	// check index exist
	if lruIndex := l.bucketIndex(l.items[l.bucket[1]]); lruIndex != 1 {
		t.Fatalf("expected lru index: %d, got: %d", 1, lruIndex)
	}

	// check index not exist
	if lruIndex := l.bucketIndex([]byte{255, 255}); lruIndex != -1 {
		t.Fatalf("expected lru index: %d, got: %d", -1, lruIndex)
	}

	l.Reset()
	if len(l.bucket) != 0 {
		t.Fatalf("expected lruBuckt is %d, got: %d", 0, len(l.bucket))
	}
	for i := range l.items {
		if l.items[i] != nil {
			t.Fatalf("[%d] expected nil, got: %v", i, l.items[i])
		}
	}
}

func BenchmarkLRU(b *testing.B) {
	var size byte = proto.LocalMesgNumMask + 1
	b.Run("100k items, zero alloc when item exist", func(b *testing.B) {
		b.StopTimer()
		l := newLRU(size)
		items := make([][]byte, 100_000)
		for i := range items {
			items[i] = []byte{byte(i % int(size))}
		}
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			for i := range items {
				l.Put(items[i])
			}
		}
	})
	b.Run("100k items, alloc since we clone the item", func(b *testing.B) {
		b.StopTimer()
		l := newLRU(size)
		items := make([][]byte, 100_000)
		for i := range items {
			items[i] = []byte{byte(i)}
		}
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			for i := range items {
				l.Put(items[i])
			}
		}
	})
}
