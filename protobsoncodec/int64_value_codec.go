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

// Int64Value type.
var TypeInt64Value = reflect.TypeOf((*wrapperspb.Int64Value)(nil))

// Int64ValueCodec is the Codec used for *wrapperspb.Int64Value values.
type Int64ValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.Int64Value.
func (vc *Int64ValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeInt64Value {
		return bsoncodec.ValueEncoderError{
			Name:     "Int64ValueCodec.EncodeValue",
			Types:    []reflect.Type{TypeInt64Value},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.Int64Value)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteInt64(val.Value)
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.Int64Value.
func (vc *Int64ValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeInt64Value {
		return bsoncodec.ValueDecoderError{
			Name:     "Int64ValueCodec.DecodeValue",
			Types:    []reflect.Type{TypeInt64Value},
			Received: v,
		}
	}
	val := &wrapperspb.Int64Value{}
	switch bsonTyp := vr.Type(); bsonTyp {
	case bsontype.Int64:
		v, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		val.Value = v
	case bsontype.Int32:
		v, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		val.Value = int64(v)
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseInt(s, 10, 64)
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
		return fmt.Errorf("cannot decode %v into a *wrapperspb.Int64Value", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewInt64ValueCodec returns a Int64ValueCodec.
func NewInt64ValueCodec() *Int64ValueCodec {
	return &Int64ValueCodec{}
}
