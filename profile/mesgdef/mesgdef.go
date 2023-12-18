package mesgdef

import (
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

type Factory interface {
	CreateMesgOnly(num typedef.MesgNum) proto.Message
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}
