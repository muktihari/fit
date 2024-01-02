package filedef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Segment files contain data defining a route and timing information to gauge progress against previous performances or other users
type Segment struct {
	FileId mesgdef.FileId

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	SegmentId               *mesgdef.SegmentId
	SegmentLeaderboardEntry *mesgdef.SegmentLeaderboardEntry
	SegmentLap              *mesgdef.SegmentLap
	SegmentPoints           []*mesgdef.SegmentPoint

	UnrelatedMessages []proto.Message
}

var _ File = &Segment{}

func NewSegment(mesgs ...proto.Message) *Segment {
	f := &Segment{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}

	return f
}

func (f *Segment) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.SegmentId:
		f.SegmentId = mesgdef.NewSegmentId(&mesg)
	case mesgnum.SegmentLeaderboardEntry:
		f.SegmentLeaderboardEntry = mesgdef.NewSegmentLeaderboardEntry(&mesg)
	case mesgnum.SegmentLap:
		f.SegmentLap = mesgdef.NewSegmentLap(&mesg)
	case mesgnum.SegmentPoint:
		f.SegmentPoints = append(f.SegmentPoints, mesgdef.NewSegmentPoint(&mesg))
	default:
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

func (f *Segment) ToFit(fac mesgdef.Factory) proto.Fit {
	if fac == nil {
		fac = factory.StandardFactory()
	}

	var size = 4 // non slice fields

	size += len(f.SegmentPoints) + len(f.DeveloperDataIds) +
		len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.Fit{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(fac))

	ToMesgs(&fit.Messages, fac, mesgnum.DeveloperDataId, f.DeveloperDataIds)
	ToMesgs(&fit.Messages, fac, mesgnum.FieldDescription, f.FieldDescriptions)

	if f.SegmentId != nil {
		fit.Messages = append(fit.Messages, f.SegmentId.ToMesg(fac))
	}

	if f.SegmentLeaderboardEntry != nil {
		fit.Messages = append(fit.Messages, f.SegmentLeaderboardEntry.ToMesg(fac))
	}

	if f.SegmentLap != nil {
		fit.Messages = append(fit.Messages, f.SegmentLap.ToMesg(fac))
	}

	ToMesgs(&fit.Messages, fac, mesgnum.SegmentPoint, f.SegmentPoints)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
