// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opener

import (
	"context"
	"os"
	"runtime"
	"sync"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

var numCPU = runtime.NumCPU()

// Open opens all paths concurrently using a number of workers equal to the lesser value of len(paths) or runtime.NumCPU().
func Open(paths []string) (fits []*proto.FIT, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	n := len(paths)
	if n > numCPU {
		n = numCPU
	}

	var (
		jobs    = make(chan string, n)
		results = make(chan result, n)
	)

	for i := 0; i < n; i++ {
		go worker(ctx, jobs, results)
	}

	go func() {
		for _, path := range paths {
			jobs <- path
		}
		close(jobs)
	}()

	// Most files has one sequence of FIT activity, grow as needed.
	fits = make([]*proto.FIT, 0, len(paths))
	for i := 0; i < len(paths); i++ {
		res := <-results
		if res.err != nil {
			return nil, res.err
		}
		fits = append(fits, res.fits...)
	}

	return fits, nil
}

type result struct {
	fits []*proto.FIT
	err  error
}

func worker(ctx context.Context, jobs <-chan string, results chan<- result) {
	for path := range jobs {
		fits, err := decode(ctx, path)
		results <- result{fits: fits, err: err}
	}
}

var pool = sync.Pool{New: func() interface{} { return decoder.New(nil) }}

func decode(ctx context.Context, path string) (fits []*proto.FIT, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	dec := pool.Get().(*decoder.Decoder)
	defer pool.Put(dec)

	dec.Reset(f)

	for dec.Next() {
		var fileId *mesgdef.FileId
		fileId, err = dec.PeekFileId()
		if err != nil {
			return
		}

		if fileId.Type != typedef.FileActivity {
			if err = dec.Discard(); err != nil {
				return
			}
			continue
		}

		var fit *proto.FIT
		fit, err = dec.DecodeWithContext(ctx)
		if err != nil {
			return
		}

		fits = append(fits, fit)
	}

	return
}
