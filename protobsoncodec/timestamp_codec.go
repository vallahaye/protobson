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
func (tsc *TimestampCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeTimestamp {
		return bsoncodec.ValueEncoderError{
			Name:     "TimestampCodec.EncodeValue",
			Types:    []reflect.Type{TypeTimestamp},
			Received: v,
		}
	}
	ts := v.Interface().(*timestamppb.Timestamp)
	return vw.WriteTimestamp(uint32(ts.GetSeconds()), uint32(ts.GetNanos()))
}

// DecodeValue is the ValueDecoderFunc for *timestamppb.Timestamp.
func (tsc *TimestampCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeTimestamp {
		return bsoncodec.ValueDecoderError{
			Name:     "TimestampCodec.DecodeValue",
			Types:    []reflect.Type{TypeTimestamp},
			Received: v,
		}
	}
	ts := &timestamppb.Timestamp{}
	switch typ := vr.Type(); typ {
	case bsontype.Timestamp:
		t, i, err := vr.ReadTimestamp()
		if err != nil {
			return err
		}
		ts.Seconds = int64(t)
		ts.Nanos = int32(i)
	case bsontype.DateTime:
		dt, err := vr.ReadDateTime()
		if err != nil {
			return err
		}
		ts.Seconds = dt / 1000
		ts.Nanos = int32(dt % 1000 * 1000000)
	case bsontype.Int64:
		sec, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		ts.Seconds = sec
	case bsontype.Int32:
		sec, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		ts.Seconds = int64(sec)
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
	case bsontype.Undefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot decode %v into a *timestamppb.Timestamp", typ)
	}
	v.Set(reflect.ValueOf(ts))
	return nil
}

// NewTimestampCodec returns a TimestampCodec.
func NewTimestampCodec() *TimestampCodec {
	return &TimestampCodec{}
}
