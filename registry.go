package protobson

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.vallahaye.net/protobson/protobsoncodec"
)

var (
	defaultMessageCodec   = protobsoncodec.NewMessageCodec()
	defaultTimestampCodec = protobsoncodec.NewTimestampCodec()
)

// DefaultRegistry is the default bsoncodec.Registry with all default protobson
// codecs registered.
var DefaultRegistry = bson.NewRegistryBuilder().
	RegisterCodec(protobsoncodec.TypeTimestamp, defaultTimestampCodec).
	RegisterHookEncoder(protobsoncodec.TypeMessage, defaultMessageCodec).
	RegisterHookDecoder(protobsoncodec.TypeMessage, defaultMessageCodec).
	Build()
