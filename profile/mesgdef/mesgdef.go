package mesgdef

import (
	"sync"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

type Factory interface {
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}

var pool = sync.Pool{New: func() any { return new([256]proto.Field) }}

type Options struct {
	Factory               Factory // If not specified, factory.StandardFactory() will be used.
	IncludeExpandedFields bool
}

var defaultOptions = DefaultOptions()

func DefaultOptions() *Options {
	return &Options{
		Factory:               factory.StandardFactory(),
		IncludeExpandedFields: false,
	}
}
