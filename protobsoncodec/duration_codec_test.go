package protobsoncodec

import (
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"
	"gotest.tools/v3/assert"
)

func TestDurationCodec(t *testing.T) {
	elapsed := time.Now().Sub(time.Date(2022, 5, 30, 11, 43, 26, 0, time.UTC))
	t.Run("EncodeToBsontype", func(t *testing.T) {
		for _, params := range []struct {
			dur  *durationpb.Duration
			vw   *bsonrwtest.ValueReaderWriter
			want bsonrwtest.Invoked
		}{
			{
				nil,
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Null},
				bsonrwtest.WriteNull,
			},
			{
				durationpb.New(elapsed),
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Int64},
				bsonrwtest.WriteInt64,
			},
		} {
			t.Run(params.vw.Type().String(), func(t *testing.T) {
				c := NewDurationCodec()
				v := reflect.ValueOf(params.dur)
				err := c.EncodeValue(bsoncodec.EncodeContext{}, params.vw, v)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, params.vw.Invoked)
			})
		}
	})
	t.Run("DecodeFromBsontype", func(t *testing.T) {
		for _, params := range []struct {
			vr   *bsonrwtest.ValueReaderWriter
			want *durationpb.Duration
		}{
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Int64,
					Return:   int64(elapsed),
				},
				durationpb.New(elapsed),
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.String,
					Return:   elapsed.String(),
				},
				durationpb.New(elapsed),
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
				&durationpb.Duration{},
			},
		} {
			t.Run(params.vr.Type().String(), func(t *testing.T) {
				c := NewDurationCodec()
				got := reflect.New(reflect.TypeOf(params.want)).Elem()
				err := c.DecodeValue(bsoncodec.DecodeContext{}, params.vr, got)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, got.Interface(), protocmp.Transform())
			})
		}
	})
}
