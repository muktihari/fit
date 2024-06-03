package mesgdef

import (
	"sync"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// Factory defines a contract that any Factory containing these method can be used by mesgdef's structs.
type Factory interface {
	// CreateField create new field based on defined messages in the factory. If not found, it returns new field with "unknown" name.
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}

var pool = sync.Pool{New: func() any { return new([255]proto.Field) }}

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
