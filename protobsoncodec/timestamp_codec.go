package protobsoncodec

import (
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Timestamp type.
var TypeTimestamp = reflect.TypeOf((*timestamppb.Timestamp)(nil))

// TimestampCodec is the Codec used for *timestamppb.Timestamp values.
type TimestampCodec struct{}

// EncodeValue is the ValueEncoderFunc for *timestamppb.Timestamp.
func (c *TimestampCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeTimestamp {
		return bsoncodec.ValueEncoderError{
			Name:     "TimestampCodec.EncodeValue",
			Types:    []reflect.Type{TypeTimestamp},
			Received: v,
		}
	}
	ts := v.Interface().(*timestamppb.Timestamp)
	if ts == nil {
		return vw.WriteNull()
	}
	return vw.WriteDateTime(ts.AsTime().UnixMilli())
}

// DecodeValue is the ValueDecoderFunc for *timestamppb.Timestamp.
func (c *TimestampCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeTimestamp {
		return bsoncodec.ValueDecoderError{
			Name:     "TimestampCodec.DecodeValue",
			Types:    []reflect.Type{TypeTimestamp},
			Received: v,
		}
	}
	var ts *timestamppb.Timestamp
	switch bsonTyp := vr.Type(); bsonTyp {
	case bsontype.DateTime:
		msec, err := vr.ReadDateTime()
		if err != nil {
			return err
		}
		ts = timestamppb.New(time.UnixMilli(msec))
	case bsontype.Int64:
		msec, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		ts = timestamppb.New(time.UnixMilli(msec))
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		t, err := time.Parse(time.RFC3339Nano, s)
		if err != nil {
			return err
		}
		ts = timestamppb.New(t)
	case bsontype.Null:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		ts = nil
	case bsontype.Undefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		ts = &timestamppb.Timestamp{}
	default:
		return fmt.Errorf("cannot decode %v into a *timestamppb.Timestamp", bsonTyp)
	}
	v.Set(reflect.ValueOf(ts))
	return nil
}

// NewTimestampCodec returns a TimestampCodec.
func NewTimestampCodec() *TimestampCodec {
	return &TimestampCodec{}
}
