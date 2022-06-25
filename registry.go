package protobson

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.vallahaye.net/protobson/protobsoncodec"
)

var (
	defaultBoolValueCodec   = protobsoncodec.NewBoolValueCodec()
	defaultBytesValueCodec  = protobsoncodec.NewBytesValueCodec()
	defaultDoubleValueCodec = protobsoncodec.NewDoubleValueCodec()
	defaultFloatValueCodec  = protobsoncodec.NewFloatValueCodec()
	defaultInt32ValueCodec  = protobsoncodec.NewInt32ValueCodec()
	defaultInt64ValueCodec  = protobsoncodec.NewInt64ValueCodec()
	defaultMessageCodec     = protobsoncodec.NewMessageCodec()
	defaultStringValueCodec = protobsoncodec.NewStringValueCodec()
	defaultTimestampCodec   = protobsoncodec.NewTimestampCodec()
	defaultUInt32ValueCodec = protobsoncodec.NewUInt32ValueCodec()
	defaultUInt64ValueCodec = protobsoncodec.NewUInt64ValueCodec()
)

// DefaultRegistry is the default bsoncodec.Registry with all default protobson
// codecs registered.
var DefaultRegistry = bson.NewRegistryBuilder().
	RegisterCodec(protobsoncodec.TypeBoolValue, defaultBoolValueCodec).
	RegisterCodec(protobsoncodec.TypeBytesValue, defaultBytesValueCodec).
	RegisterCodec(protobsoncodec.TypeDoubleValue, defaultDoubleValueCodec).
	RegisterCodec(protobsoncodec.TypeFloatValue, defaultFloatValueCodec).
	RegisterCodec(protobsoncodec.TypeInt32Value, defaultInt32ValueCodec).
	RegisterCodec(protobsoncodec.TypeInt64Value, defaultInt64ValueCodec).
	RegisterHookEncoder(protobsoncodec.TypeMessage, defaultMessageCodec).
	RegisterHookDecoder(protobsoncodec.TypeMessage, defaultMessageCodec).
	RegisterCodec(protobsoncodec.TypeStringValue, defaultStringValueCodec).
	RegisterCodec(protobsoncodec.TypeTimestamp, defaultTimestampCodec).
	RegisterCodec(protobsoncodec.TypeUInt32Value, defaultUInt32ValueCodec).
	RegisterCodec(protobsoncodec.TypeUInt64Value, defaultUInt64ValueCodec).
	Build()
