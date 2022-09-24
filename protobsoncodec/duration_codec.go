package protobsoncodec

import (
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Duration type.
var TypeDuration = reflect.TypeOf((*durationpb.Duration)(nil))

// DurationCodec is the Codec used for *durationpb.Duration values.
type DurationCodec struct{}

// EncodeValue is the ValueEncoderFunc for *durationpb.Duration.
func (c *DurationCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeDuration {
		return bsoncodec.ValueEncoderError{
			Name:     "DurationCodec.EncodeValue",
			Types:    []reflect.Type{TypeDuration},
			Received: v,
		}
	}
	dur := v.Interface().(*durationpb.Duration)
	if dur == nil {
		return vw.WriteNull()
	}
	return vw.WriteInt64(int64(dur.AsDuration()))
}

// DecodeValue is the ValueDecoderFunc for *durationpb.Duration.
func (c *DurationCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeDuration {
		return bsoncodec.ValueDecoderError{
			Name:     "DurationCodec.DecodeValue",
			Types:    []reflect.Type{TypeDuration},
			Received: v,
		}
	}
	var dur *durationpb.Duration
	switch bsonTyp := vr.Type(); bsonTyp {
	case bsontype.Int64:
		nsec, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		dur = durationpb.New(time.Duration(nsec))
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		d, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		dur = durationpb.New(d)
	case bsontype.Null:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		dur = nil
	case bsontype.Undefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		dur = &durationpb.Duration{}
	default:
		return fmt.Errorf("cannot decode %v into a *durationpb.Duration", bsonTyp)
	}
	v.Set(reflect.ValueOf(dur))
	return nil
}

// NewDurationCodec returns a DurationCodec.
func NewDurationCodec() *DurationCodec {
	return &DurationCodec{}
}
