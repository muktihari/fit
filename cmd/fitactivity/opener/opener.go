// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opener

import (
	"bufio"
	"context"
	"os"
	"sync"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// Open opens all given paths concurrently.
func Open(paths ...string) ([]proto.Fit, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultc := make(chan result, len(paths))
	var wg sync.WaitGroup

	wg.Add(len(paths))
	for i := range paths {
		path := paths[i]
		go worker(ctx, path, resultc, &wg)
	}

	go func() {
		wg.Wait()
		close(resultc)
	}()

	fits := make([]proto.Fit, 0, len(paths))
	for res := range resultc {
		if res.err != nil {
			return nil, res.err
		}
		fits = append(fits, *res.fit)
	}

	return fits, nil
}

type result struct {
	fit *proto.Fit
	err error
}

func worker(ctx context.Context, path string, resultc chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(path)
	if err != nil {
		resultc <- result{err: err}
		return
	}
	defer f.Close()

	dec := decoder.New(bufio.NewReader(f),
		decoder.WithNoComponentExpansion(),
	)

	for {
		fileId, err := dec.PeekFileId()
		if err != nil {
			resultc <- result{err: err}
			return
		}

		if fileId.Type != typedef.FileActivity {
			// Chained file may be composed of activity files and other file types, skip the non-activity file.
			continue
		}

		fit, err := dec.DecodeWithContext(ctx)
		if err != nil {
			resultc <- result{err: err}
			return
		}

		resultc <- result{fit: fit}

		if !dec.Next() {
			break
		}
	}
}
