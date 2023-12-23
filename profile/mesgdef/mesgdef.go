package mesgdef

import (
	"sync"

	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

type Factory interface {
	CreateMesgOnly(num typedef.MesgNum) proto.Message
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}

var fieldsPool = sync.Pool{
	New: func() any {
		fields := [256]proto.Field{}
		return &fields
	},
}
