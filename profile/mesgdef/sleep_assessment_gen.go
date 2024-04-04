// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// SleepAssessment is a SleepAssessment message.
type SleepAssessment struct {
	AverageStressDuringSleep uint16 // Scale: 100; Excludes stress during awake periods in the sleep window
	CombinedAwakeScore       uint8  // Average of awake_time_score and awakenings_count_score. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AwakeTimeScore           uint8  // Score that evaluates the total time spent awake between sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AwakeningsCountScore     uint8  // Score that evaluates the number of awakenings that interrupt sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	DeepSleepScore           uint8  // Score that evaluates the amount of deep sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepDurationScore       uint8  // Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	LightSleepScore          uint8  // Score that evaluates the amount of light sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	OverallSleepScore        uint8  // Total score that summarizes the overall quality of sleep, combining sleep duration and quality. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepQualityScore        uint8  // Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepRecoveryScore       uint8  // Score that evaluates stress and recovery during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	RemSleepScore            uint8  // Score that evaluates the amount of REM sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepRestlessnessScore   uint8  // Score that evaluates the amount of restlessness during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AwakeningsCount          uint8  // The number of awakenings during sleep.
	InterruptionsScore       uint8  // Score that evaluates the sleep interruptions. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSleepAssessment creates new SleepAssessment struct based on given mesg.
// If mesg is nil, it will return SleepAssessment with all fields being set to its corresponding invalid value.
func NewSleepAssessment(mesg *proto.Message) *SleepAssessment {
	vals := [16]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &SleepAssessment{
		AverageStressDuringSleep: vals[15].Uint16(),
		CombinedAwakeScore:       vals[0].Uint8(),
		AwakeTimeScore:           vals[1].Uint8(),
		AwakeningsCountScore:     vals[2].Uint8(),
		DeepSleepScore:           vals[3].Uint8(),
		SleepDurationScore:       vals[4].Uint8(),
		LightSleepScore:          vals[5].Uint8(),
		OverallSleepScore:        vals[6].Uint8(),
		SleepQualityScore:        vals[7].Uint8(),
		SleepRecoveryScore:       vals[8].Uint8(),
		RemSleepScore:            vals[9].Uint8(),
		SleepRestlessnessScore:   vals[10].Uint8(),
		AwakeningsCount:          vals[11].Uint8(),
		InterruptionsScore:       vals[14].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SleepAssessment into proto.Message. If options is nil, default options will be used.
func (m *SleepAssessment) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumSleepAssessment)

	if m.AverageStressDuringSleep != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Uint16(m.AverageStressDuringSleep)
		fields = append(fields, field)
	}
	if m.CombinedAwakeScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(m.CombinedAwakeScore)
		fields = append(fields, field)
	}
	if m.AwakeTimeScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.AwakeTimeScore)
		fields = append(fields, field)
	}
	if m.AwakeningsCountScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(m.AwakeningsCountScore)
		fields = append(fields, field)
	}
	if m.DeepSleepScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.DeepSleepScore)
		fields = append(fields, field)
	}
	if m.SleepDurationScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(m.SleepDurationScore)
		fields = append(fields, field)
	}
	if m.LightSleepScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(m.LightSleepScore)
		fields = append(fields, field)
	}
	if m.OverallSleepScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.OverallSleepScore)
		fields = append(fields, field)
	}
	if m.SleepQualityScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(m.SleepQualityScore)
		fields = append(fields, field)
	}
	if m.SleepRecoveryScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint8(m.SleepRecoveryScore)
		fields = append(fields, field)
	}
	if m.RemSleepScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint8(m.RemSleepScore)
		fields = append(fields, field)
	}
	if m.SleepRestlessnessScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint8(m.SleepRestlessnessScore)
		fields = append(fields, field)
	}
	if m.AwakeningsCount != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(m.AwakeningsCount)
		fields = append(fields, field)
	}
	if m.InterruptionsScore != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint8(m.InterruptionsScore)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// AverageStressDuringSleepScaled return AverageStressDuringSleep in its scaled value [Scale: 100; Excludes stress during awake periods in the sleep window].
//
// If AverageStressDuringSleep value is invalid, float64 invalid value will be returned.
func (m *SleepAssessment) AverageStressDuringSleepScaled() float64 {
	if m.AverageStressDuringSleep == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.AverageStressDuringSleep, 100, 0)
}

// SetAverageStressDuringSleep sets SleepAssessment value.
//
// Scale: 100; Excludes stress during awake periods in the sleep window
func (m *SleepAssessment) SetAverageStressDuringSleep(v uint16) *SleepAssessment {
	m.AverageStressDuringSleep = v
	return m
}

// SetCombinedAwakeScore sets SleepAssessment value.
//
// Average of awake_time_score and awakenings_count_score. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetCombinedAwakeScore(v uint8) *SleepAssessment {
	m.CombinedAwakeScore = v
	return m
}

// SetAwakeTimeScore sets SleepAssessment value.
//
// Score that evaluates the total time spent awake between sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetAwakeTimeScore(v uint8) *SleepAssessment {
	m.AwakeTimeScore = v
	return m
}

// SetAwakeningsCountScore sets SleepAssessment value.
//
// Score that evaluates the number of awakenings that interrupt sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetAwakeningsCountScore(v uint8) *SleepAssessment {
	m.AwakeningsCountScore = v
	return m
}

// SetDeepSleepScore sets SleepAssessment value.
//
// Score that evaluates the amount of deep sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetDeepSleepScore(v uint8) *SleepAssessment {
	m.DeepSleepScore = v
	return m
}

// SetSleepDurationScore sets SleepAssessment value.
//
// Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetSleepDurationScore(v uint8) *SleepAssessment {
	m.SleepDurationScore = v
	return m
}

// SetLightSleepScore sets SleepAssessment value.
//
// Score that evaluates the amount of light sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetLightSleepScore(v uint8) *SleepAssessment {
	m.LightSleepScore = v
	return m
}

// SetOverallSleepScore sets SleepAssessment value.
//
// Total score that summarizes the overall quality of sleep, combining sleep duration and quality. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetOverallSleepScore(v uint8) *SleepAssessment {
	m.OverallSleepScore = v
	return m
}

// SetSleepQualityScore sets SleepAssessment value.
//
// Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetSleepQualityScore(v uint8) *SleepAssessment {
	m.SleepQualityScore = v
	return m
}

// SetSleepRecoveryScore sets SleepAssessment value.
//
// Score that evaluates stress and recovery during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetSleepRecoveryScore(v uint8) *SleepAssessment {
	m.SleepRecoveryScore = v
	return m
}

// SetRemSleepScore sets SleepAssessment value.
//
// Score that evaluates the amount of REM sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetRemSleepScore(v uint8) *SleepAssessment {
	m.RemSleepScore = v
	return m
}

// SetSleepRestlessnessScore sets SleepAssessment value.
//
// Score that evaluates the amount of restlessness during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetSleepRestlessnessScore(v uint8) *SleepAssessment {
	m.SleepRestlessnessScore = v
	return m
}

// SetAwakeningsCount sets SleepAssessment value.
//
// The number of awakenings during sleep.
func (m *SleepAssessment) SetAwakeningsCount(v uint8) *SleepAssessment {
	m.AwakeningsCount = v
	return m
}

// SetInterruptionsScore sets SleepAssessment value.
//
// Score that evaluates the sleep interruptions. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
func (m *SleepAssessment) SetInterruptionsScore(v uint8) *SleepAssessment {
	m.InterruptionsScore = v
	return m
}

// SetDeveloperFields SleepAssessment's DeveloperFields.
func (m *SleepAssessment) SetDeveloperFields(developerFields ...proto.DeveloperField) *SleepAssessment {
	m.DeveloperFields = developerFields
	return m
}
