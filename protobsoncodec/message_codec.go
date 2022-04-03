package protobsoncodec

import (
	"reflect"
	"strings"

	"github.com/vallahaye/protobson/protobsonoptions"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/proto"
)

// Message type.
var TypeMessage = reflect.TypeOf((*proto.Message)(nil)).Elem()

// MessageCodec is the Codec used for proto.Message values.
type MessageCodec struct {
	*bsoncodec.StructCodec
}

// EncodeValue is the ValueEncoderFunc for proto.Message.
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

// DecodeValue is the ValueDecoderFunc for proto.Message.
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

// NewMessageCodec returns a MessageCodec with options opts.
func NewMessageCodec(opts ...*protobsonoptions.MessageCodecOptions) *MessageCodec {
	mergedOpts := protobsonoptions.MergeMessageCodecOptions(opts...)
	p := JSONPBFallbackStructTagParser
	if mergedOpts.UseProtoNames != nil && *mergedOpts.UseProtoNames {
		p = ProtoNamesFallbackStructTagParser
	}
	sc, _ := bsoncodec.NewStructCodec(p, mergedOpts.StructCodecOptions)
	return &MessageCodec{sc}
}

// JSONPBFallbackStructTagParser is the StructTagParser used by the MessageCodec by default.
// It has the same behavior as bsoncodec.DefaultStructTagParser but will also fallback to
// parsing the protobuf tag on a field where the bson tag isn't available. In this case, the
// key will be taken from the json property, or from the name property if there is none.
//
// An example:
//
//   type T struct {
//     Name   string `protobuf:"bytes,1,opt,name=name,proto3"` // Key is "name"
//     FooBar string `protobuf:"bytes,2,opt,name=foo_bar,json=fooBar,proto3"` // Key is "fooBar"
//     BarFoo string `protobuf:"bytes,3,opt,name=bar_foo,json=barFoo,proto3" bson:"barfoo"` // Key is "barfoo"
//   }
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

// ProtoNamesFallbackStructTagParser has the same behavior as JSONPBFallbackStructTagParser
// except it forces the use of the name property as the key when parsing protobuf tags.
var ProtoNamesFallbackStructTagParser bsoncodec.StructTagParserFunc = func(sf reflect.StructField) (bsoncodec.StructTags, error) {
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
