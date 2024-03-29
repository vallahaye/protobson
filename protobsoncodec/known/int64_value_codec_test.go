package known

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func TestInt64ValueCodec(t *testing.T) {
	t.Run("EncodeToBsontype", func(t *testing.T) {
		for _, params := range []struct {
			val  *wrapperspb.Int64Value
			vw   *bsonrwtest.ValueReaderWriter
			want bsonrwtest.Invoked
		}{
			{
				nil,
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Null},
				bsonrwtest.WriteNull,
			},
			{
				wrapperspb.Int64(42),
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Int64},
				bsonrwtest.WriteInt64,
			},
		} {
			c := NewInt64ValueCodec()
			v := reflect.ValueOf(params.val)
			err := c.EncodeValue(bsoncodec.EncodeContext{}, params.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, params.want, params.vw.Invoked)
		}
	})
	t.Run("DecodeFromBsontype", func(t *testing.T) {
		for _, params := range []struct {
			vr   *bsonrwtest.ValueReaderWriter
			want *wrapperspb.Int64Value
		}{
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Int64,
					Return:   int64(42),
				},
				wrapperspb.Int64(42),
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Int32,
					Return:   int32(42),
				},
				wrapperspb.Int64(42),
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.String,
					Return:   "42",
				},
				wrapperspb.Int64(42),
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
				&wrapperspb.Int64Value{},
			},
		} {
			t.Run(params.vr.Type().String(), func(t *testing.T) {
				c := NewInt64ValueCodec()
				got := reflect.New(reflect.TypeOf(params.want)).Elem()
				err := c.DecodeValue(bsoncodec.DecodeContext{}, params.vr, got)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, got.Interface(), protocmp.Transform())
			})
		}
	})
}
