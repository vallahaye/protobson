package protobsoncodec

import (
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// BytesValue type.
var TypeBytesValue = reflect.TypeOf((*wrapperspb.BytesValue)(nil))

// BytesValueCodec is the Codec used for *wrapperspb.BytesValue values.
type BytesValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.BytesValue.
func (c *BytesValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeBytesValue {
		return bsoncodec.ValueEncoderError{
			Name:     "BytesValueCodec.EncodeValue",
			Types:    []reflect.Type{TypeBytesValue},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.BytesValue)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteBinary(val.Value)
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.BytesValue.
func (c *BytesValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeBytesValue {
		return bsoncodec.ValueDecoderError{
			Name:     "BytesValueCodec.DecodeValue",
			Types:    []reflect.Type{TypeBytesValue},
			Received: v,
		}
	}
	var val *wrapperspb.BytesValue
	switch bsonTyp := vr.Type(); bsonTyp {
	case bsontype.Binary:
		v, _, err := vr.ReadBinary()
		if err != nil {
			return err
		}
		val = wrapperspb.Bytes(v)
	case bsontype.String:
		v, err := vr.ReadString()
		if err != nil {
			return err
		}
		val = wrapperspb.Bytes([]byte(v))
	case bsontype.Null:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bsontype.Undefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		val = &wrapperspb.BytesValue{}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.BytesValue", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewBytesValueCodec returns a BytesValueCodec.
func NewBytesValueCodec() *BytesValueCodec {
	return &BytesValueCodec{}
}
