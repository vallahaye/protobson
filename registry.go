package protobson

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.vallahaye.net/protobson/protobsoncodec"
	googleapiscodec "go.vallahaye.net/protobson/protobsoncodec/googleapis"
	knowncodec "go.vallahaye.net/protobson/protobsoncodec/known"
)

var (
	defaultBoolValueCodec    = knowncodec.NewBoolValueCodec()
	defaultBytesValueCodec   = knowncodec.NewBytesValueCodec()
	defaultDoubleValueCodec  = knowncodec.NewDoubleValueCodec()
	defaultDurationCodec     = knowncodec.NewDurationCodec()
	defaultFloatValueCodec   = knowncodec.NewFloatValueCodec()
	defaultInt32ValueCodec   = knowncodec.NewInt32ValueCodec()
	defaultInt64ValueCodec   = knowncodec.NewInt64ValueCodec()
	defaultMessageCodec      = protobsoncodec.NewMessageCodec()
	defaultStringValueCodec  = knowncodec.NewStringValueCodec()
	defaultTimestampCodec    = knowncodec.NewTimestampCodec()
	defaultUInt32ValueCodec  = knowncodec.NewUInt32ValueCodec()
	defaultUInt64ValueCodec  = knowncodec.NewUInt64ValueCodec()
	defaultGAPIDateTimeCodec = googleapiscodec.NewDateTimeCodec()
)

// DefaultRegistry is the default bsoncodec.Registry with all default protobson
// codecs registered.
var DefaultRegistry = bson.NewRegistryBuilder().
	RegisterCodec(knowncodec.TypeBoolValue, defaultBoolValueCodec).
	RegisterCodec(knowncodec.TypeBytesValue, defaultBytesValueCodec).
	RegisterCodec(knowncodec.TypeDoubleValue, defaultDoubleValueCodec).
	RegisterCodec(knowncodec.TypeDuration, defaultDurationCodec).
	RegisterCodec(knowncodec.TypeFloatValue, defaultFloatValueCodec).
	RegisterCodec(knowncodec.TypeInt32Value, defaultInt32ValueCodec).
	RegisterCodec(knowncodec.TypeInt64Value, defaultInt64ValueCodec).
	RegisterHookEncoder(protobsoncodec.TypeMessage, defaultMessageCodec).
	RegisterHookDecoder(protobsoncodec.TypeMessage, defaultMessageCodec).
	RegisterCodec(knowncodec.TypeStringValue, defaultStringValueCodec).
	RegisterCodec(knowncodec.TypeTimestamp, defaultTimestampCodec).
	RegisterCodec(knowncodec.TypeUInt32Value, defaultUInt32ValueCodec).
	RegisterCodec(knowncodec.TypeUInt64Value, defaultUInt64ValueCodec).
	RegisterCodec(googleapiscodec.TypeDateTime, defaultGAPIDateTimeCodec).
	Build()
