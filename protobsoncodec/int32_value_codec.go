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

// Int32Value type.
var TypeInt32Value = reflect.TypeOf((*wrapperspb.Int32Value)(nil))

// Int32ValueCodec is the Codec used for *wrapperspb.Int32Value values.
type Int32ValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.Int32Value.
func (vc *Int32ValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeInt32Value {
		return bsoncodec.ValueEncoderError{
			Name:     "Int32ValueCodec.EncodeValue",
			Types:    []reflect.Type{TypeInt32Value},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.Int32Value)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteInt32(val.Value)
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.Int32Value.
func (vc *Int32ValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeInt32Value {
		return bsoncodec.ValueDecoderError{
			Name:     "Int32ValueCodec.DecodeValue",
			Types:    []reflect.Type{TypeInt32Value},
			Received: v,
		}
	}
	val := &wrapperspb.Int32Value{}
	switch bsonTyp := vr.Type(); bsonTyp {
	case bsontype.Int32:
		v, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		val.Value = v
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return err
		}
		val.Value = int32(v)
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
		return fmt.Errorf("cannot decode %v into a *wrapperspb.Int32Value", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewInt32ValueCodec returns a Int32ValueCodec.
func NewInt32ValueCodec() *Int32ValueCodec {
	return &Int32ValueCodec{}
}
