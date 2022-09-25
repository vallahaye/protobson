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

// UInt32Value type.
var TypeUInt32Value = reflect.TypeOf((*wrapperspb.UInt32Value)(nil))

// UInt32ValueCodec is the Codec used for *wrapperspb.UInt32Value values.
type UInt32ValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.UInt32Value.
func (vc *UInt32ValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeUInt32Value {
		return bsoncodec.ValueEncoderError{
			Name:     "UInt32ValueCodec.EncodeValue",
			Types:    []reflect.Type{TypeUInt32Value},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.UInt32Value)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteInt32(int32(val.Value))
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.UInt32Value.
func (vc *UInt32ValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeUInt32Value {
		return bsoncodec.ValueDecoderError{
			Name:     "UInt32ValueCodec.DecodeValue",
			Types:    []reflect.Type{TypeUInt32Value},
			Received: v,
		}
	}
	var val *wrapperspb.UInt32Value
	switch bsonTyp := vr.Type(); bsonTyp {
	case bsontype.Int32:
		v, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		val = wrapperspb.UInt32(uint32(v))
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return err
		}
		val = wrapperspb.UInt32(uint32(v))
	case bsontype.Null:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		val = nil
	case bsontype.Undefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
		val = &wrapperspb.UInt32Value{}
	default:
		return fmt.Errorf("cannot decode %v into a *wrapperspb.UInt32Value", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewUInt32ValueCodec returns a UInt32ValueCodec.
func NewUInt32ValueCodec() *UInt32ValueCodec {
	return &UInt32ValueCodec{}
}
