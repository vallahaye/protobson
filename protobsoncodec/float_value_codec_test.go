package protobsoncodec

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

func TestFloatValueCodec(t *testing.T) {
	t.Run("EncodeToBsontype", func(t *testing.T) {
		for _, params := range []struct {
			val  *wrapperspb.FloatValue
			vw   *bsonrwtest.ValueReaderWriter
			want bsonrwtest.Invoked
		}{
			{
				nil,
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Null},
				bsonrwtest.WriteNull,
			},
			{
				wrapperspb.Float(3.1415926535),
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Double},
				bsonrwtest.WriteDouble,
			},
		} {
			vc := NewFloatValueCodec()
			v := reflect.ValueOf(params.val)
			err := vc.EncodeValue(bsoncodec.EncodeContext{}, params.vw, v)
			assert.NilError(t, err)
			assert.DeepEqual(t, params.want, params.vw.Invoked)
		}
	})
	t.Run("DecodeFromBsontype", func(t *testing.T) {
		for _, params := range []struct {
			vr   *bsonrwtest.ValueReaderWriter
			want *wrapperspb.FloatValue
		}{
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Null,
				},
				nil,
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Double,
					Return:   float64(3.1415926535),
				},
				wrapperspb.Float(3.1415926535),
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.String,
					Return:   "3.1415926535",
				},
				wrapperspb.Float(3.1415926535),
			},
		} {
			t.Run(params.vr.Type().String(), func(t *testing.T) {
				vc := NewFloatValueCodec()
				got := reflect.New(reflect.TypeOf(params.want)).Elem()
				err := vc.DecodeValue(bsoncodec.DecodeContext{}, params.vr, got)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, got.Interface(), protocmp.Transform())
			})
		}
	})
}
