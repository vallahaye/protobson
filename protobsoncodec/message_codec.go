package protobsoncodec

import (
	"reflect"
	"strings"

	"github.com/vallahaye/protobson/protobsonoptions"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/proto"
)

var TypeMessage = reflect.TypeOf((*proto.Message)(nil)).Elem()

type MessageCodec struct {
	*bsoncodec.StructCodec
}

func (mc *MessageCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, v reflect.Value) error {
	if !v.IsValid() || (!v.Type().Implements(TypeMessage) && !reflect.PtrTo(v.Type()).Implements(TypeMessage)) {
		return bsoncodec.ValueEncoderError{
			Name:     "MessageCodec.EncodeValue",
			Types:    []reflect.Type{TypeMessage},
			Received: v,
		}
	}
	return mc.StructCodec.EncodeValue(ec, vw, v.Elem())
}

func (mc *MessageCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, v reflect.Value) error {
	if !v.CanSet() || (!v.Type().Implements(TypeMessage) && !reflect.PtrTo(v.Type()).Implements(TypeMessage)) {
		return bsoncodec.ValueDecoderError{
			Name:     "MessageCodec.DecodeValue",
			Types:    []reflect.Type{TypeMessage},
			Received: v,
		}
	}
	return mc.StructCodec.DecodeValue(dc, vr, v.Elem())
}

func NewMessageCodec(opts ...*protobsonoptions.MessageCodecOptions) *MessageCodec {
	mergedOpts := protobsonoptions.MergeMessageCodecOptions(opts...)
	p := JSONPBFallbackStructTagParser
	if mergedOpts.UseProtoNames != nil && *mergedOpts.UseProtoNames {
		p = ProtoNameFallbackStructTagParser
	}
	sc, _ := bsoncodec.NewStructCodec(p, mergedOpts.StructCodecOptions)
	return &MessageCodec{sc}
}

var JSONPBFallbackStructTagParser bsoncodec.StructTagParserFunc = func(sf reflect.StructField) (bsoncodec.StructTags, error) {
	if _, ok := sf.Tag.Lookup("bson"); ok {
		return bsoncodec.DefaultStructTagParser(sf)
	}
	tag, ok := sf.Tag.Lookup("protobuf")
	if !ok {
		return bsoncodec.DefaultStructTagParser(sf)
	}
	return parseTags(tag, false)
}

var ProtoNameFallbackStructTagParser bsoncodec.StructTagParserFunc = func(sf reflect.StructField) (bsoncodec.StructTags, error) {
	if _, ok := sf.Tag.Lookup("bson"); ok {
		return bsoncodec.DefaultStructTagParser(sf)
	}
	tag, ok := sf.Tag.Lookup("protobuf")
	if !ok {
		return bsoncodec.DefaultStructTagParser(sf)
	}
	return parseTags(tag, true)
}

func parseTags(tag string, useProtoNames bool) (bsoncodec.StructTags, error) {
	rawProps := strings.Split(tag, ",")
	props := make(map[string]string, len(rawProps))
	for _, rawProp := range rawProps {
		k, v, _ := strings.Cut(rawProp, "=")
		props[k] = v
	}
	var st bsoncodec.StructTags
	jsonName, hasJSONName := props["json"]
	if !useProtoNames && hasJSONName {
		st.Name = jsonName
	} else {
		st.Name = props["name"]
	}
	return st, nil
}
