package googleapis

import (
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/protobuf/types/known/durationpb"
)

// DateTime type.
var TypeDateTime = reflect.TypeOf((*datetime.DateTime)(nil))

// DateTimeCodec is the Codec used for *datetime.DateTime values.
type DateTimeCodec struct{}

// EncodeValue is the ValueEncoderFunc for *datetime.DateTime.
func (c *DateTimeCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeDateTime {
		return bsoncodec.ValueEncoderError{
			Name:     "DateTimeCodec.EncodeValue",
			Types:    []reflect.Type{TypeDateTime},
			Received: v,
		}
	}
	dt := v.Interface().(*datetime.DateTime)
	if dt == nil {
		return vw.WriteNull()
	}
	t, err := dateTimeToTime(dt)
	if err != nil {
		return err
	}
	return vw.WriteDateTime(t.UnixMilli())
}

// DecodeValue is the ValueDecoderFunc for *datetime.DateTime.
func (c *DateTimeCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeDateTime {
		return bsoncodec.ValueDecoderError{
			Name:     "DateTimeCodec.DecodeValue",
			Types:    []reflect.Type{TypeDateTime},
			Received: v,
		}
	}
	var dt *datetime.DateTime
	switch bsonTyp := vr.Type(); bsonTyp {
	case bsontype.DateTime:
		msec, err := vr.ReadDateTime()
		if err != nil {
			return err
		}
		dt = timeToDateTime(time.UnixMilli(msec).UTC())
	case bsontype.Int64:
		msec, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		dt = timeToDateTime(time.UnixMilli(msec).UTC())
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		t, err := time.Parse(time.RFC3339Nano, s)
		if err != nil {
			return err
		}
		dt = timeToDateTime(t)
	case bsontype.Null:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		dt = nil
	case bsontype.Undefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		dt = &datetime.DateTime{}
	default:
		return fmt.Errorf("cannot decode %v into a *datetime.DateTime", bsonTyp)
	}
	v.Set(reflect.ValueOf(dt))
	return nil
}

// NewDateTimeCodec returns a DateTimeCodec.
func NewDateTimeCodec() *DateTimeCodec {
	return &DateTimeCodec{}
}

func dateTimeToTime(dt *datetime.DateTime) (time.Time, error) {
	loc := time.Local
	if dt.TimeOffset != nil {
		switch timeOffset := dt.TimeOffset.(type) {
		case *datetime.DateTime_UtcOffset:
			loc = time.FixedZone("", int(timeOffset.UtcOffset.Seconds))
		case *datetime.DateTime_TimeZone:
			var err error
			loc, err = time.LoadLocation(timeOffset.TimeZone.Id)
			if err != nil {
				return time.Time{}, err
			}
		}
	}
	return time.Date(int(dt.Year), time.Month(dt.Month), int(dt.Day), int(dt.Hours), int(dt.Minutes), int(dt.Seconds), int(dt.Nanos), loc), nil
}

func timeToDateTime(t time.Time) *datetime.DateTime {
	_, offset := t.Zone()
	return &datetime.DateTime{
		Year:    int32(t.Year()),
		Month:   int32(t.Month()),
		Day:     int32(t.Day()),
		Hours:   int32(t.Hour()),
		Minutes: int32(t.Minute()),
		Seconds: int32(t.Second()),
		Nanos:   int32(t.Nanosecond()),
		TimeOffset: &datetime.DateTime_UtcOffset{
			UtcOffset: durationpb.New(time.Duration(offset) * time.Second),
		},
	}
}
