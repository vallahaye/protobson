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

// UInt64Value type.
var TypeUInt64Value = reflect.TypeOf((*wrapperspb.UInt64Value)(nil))

// UInt64ValueCodec is the Codec used for *wrapperspb.UInt64Value values.
type UInt64ValueCodec struct{}

// EncodeValue is the ValueEncoderFunc for *wrapperspb.UInt64Value.
func (vc *UInt64ValueCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || v.Type() != TypeUInt64Value {
		return bsoncodec.ValueEncoderError{
			Name:     "UInt64ValueCodec.EncodeValue",
			Types:    []reflect.Type{TypeUInt64Value},
			Received: v,
		}
	}
	val := v.Interface().(*wrapperspb.UInt64Value)
	if val == nil {
		return vw.WriteNull()
	}
	return vw.WriteInt64(int64(val.Value))
}

// DecodeValue is the ValueDecoderFunc for *wrapperspb.UInt64Value.
func (vc *UInt64ValueCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || v.Type() != TypeUInt64Value {
		return bsoncodec.ValueDecoderError{
			Name:     "UInt64ValueCodec.DecodeValue",
			Types:    []reflect.Type{TypeUInt64Value},
			Received: v,
		}
	}
	val := &wrapperspb.UInt64Value{}
	switch bsonTyp := vr.Type(); bsonTyp {
	case bsontype.Int64:
		v, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		val.Value = uint64(v)
	case bsontype.Int32:
		v, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		val.Value = uint64(v)
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		v, err := strconv.ParseUint(s, 10, 64)
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
		return fmt.Errorf("cannot decode %v into a *wrapperspb.UInt64Value", bsonTyp)
	}
	v.Set(reflect.ValueOf(val))
	return nil
}

// NewUInt64ValueCodec returns a UInt64ValueCodec.
func NewUInt64ValueCodec() *UInt64ValueCodec {
	return &UInt64ValueCodec{}
}
