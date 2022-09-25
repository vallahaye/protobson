package protobsoncodec

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func TestStringValueCodec(t *testing.T) {
	t.Run("EncodeToBsontype", func(t *testing.T) {
		for _, params := range []struct {
			val  *wrapperspb.StringValue
			vw   *bsonrwtest.ValueReaderWriter
			want bsonrwtest.Invoked
		}{
			{
				nil,
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Null},
				bsonrwtest.WriteNull,
			},
			{
				wrapperspb.String("Hello, World!"),
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.String},
				bsonrwtest.WriteString,
			},
		} {
			c := NewStringValueCodec()
			v := reflect.ValueOf(params.val)
			err := c.EncodeValue(bsoncodec.EncodeContext{}, params.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, params.want, params.vw.Invoked)
		}
	})
	t.Run("DecodeFromBsontype", func(t *testing.T) {
		for _, params := range []struct {
			vr   *bsonrwtest.ValueReaderWriter
			want *wrapperspb.StringValue
		}{
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.String,
					Return:   "Hello, World!",
				},
				wrapperspb.String("Hello, World!"),
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Binary,
					Return: bsoncore.Value{
						Type: bsontype.Binary,
						Data: bsoncore.AppendBinary(nil, 0x00, []byte("Hello, World!")),
					},
				},
				wrapperspb.String("Hello, World!"),
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Null,
				},
				nil,
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Undefined,
				},
				&wrapperspb.StringValue{},
			},
		} {
			t.Run(params.vr.Type().String(), func(t *testing.T) {
				c := NewStringValueCodec()
				got := reflect.New(reflect.TypeOf(params.want)).Elem()
				err := c.DecodeValue(bsoncodec.DecodeContext{}, params.vr, got)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, got.Interface(), protocmp.Transform())
			})
		}
	})
}
