package protobson

import (
	"github.com/vallahaye/protobson/protobsoncodec"
	"go.mongodb.org/mongo-driver/bson"
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
