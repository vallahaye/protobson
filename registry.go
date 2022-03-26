package protobson

import (
	"github.com/vallahaye/protobson/protobsoncodec"
	"go.mongodb.org/mongo-driver/bson"
)

var defaultMessageCodec = protobsoncodec.NewMessageCodec()

var DefaultRegistry = bson.NewRegistryBuilder().
	RegisterHookEncoder(protobsoncodec.TypeMessage, defaultMessageCodec).
	RegisterHookDecoder(protobsoncodec.TypeMessage, defaultMessageCodec).
	Build()
