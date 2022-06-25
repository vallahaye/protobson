package protobsoncodec

import (
	"fmt"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// DoubleValue type.
var TypeDoubleValue = reflect.TypeOf((*wrapperspb.DoubleValue)(nil))

// DoubleValueCodec is the Codec used for *wrapperspb.DoubleValue values.
type DoubleValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.DoubleValue.
func (vc *DoubleValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeDoubleValue {
		return bsoncodec.ValueEncoderError{
			Name:     "DoubleValueCodec.EncodeValue",
			Types:    []reflect.Type{TypeDoubleValue},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.DoubleValue)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteDouble(val.Value)
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.DoubleValue.
func (vc *DoubleValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeDoubleValue {
		return bsoncodec.ValueDecoderError{
			Name:     "DoubleValueCodec.DecodeValue",
			Types:    []reflect.Type{TypeDoubleValue},
			Received: v,
		}
	}
	val := &wrapperspb.DoubleValue{}
	switch bsonTyp := vr.Type(); bsonTyp {
	case bsontype.Double:
		v, err := vr.ReadDouble()
		if err != nil {
			return err
		}
		val.Value = v
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		val.Value = v
	case bsontype.Null:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bsontype.Undefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.DoubleValue", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewDoubleValueCodec returns a DoubleValueCodec.
func NewDoubleValueCodec() *DoubleValueCodec {
	return &DoubleValueCodec{}
}
