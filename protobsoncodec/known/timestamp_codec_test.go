package known

import (
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
)

func TestTimestampCodec(t *testing.T) {
	ts := timestamppb.New(time.Date(2022, 5, 30, 11, 43, 26, 0, time.UTC))
	t.Run("EncodeToBsontype", func(t *testing.T) {
		for _, params := range []struct {
			ts   *timestamppb.Timestamp
			vw   *bsonrwtest.ValueReaderWriter
			want bsonrwtest.Invoked
		}{
			{
				nil,
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Null},
				bsonrwtest.WriteNull,
			},
			{
				ts,
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.DateTime},
				bsonrwtest.WriteDateTime,
			},
		} {
			t.Run(params.vw.Type().String(), func(t *testing.T) {
				c := NewTimestampCodec()
				v := reflect.ValueOf(params.ts)
				err := c.EncodeValue(bsoncodec.EncodeContext{}, params.vw, v)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, params.vw.Invoked)
			})
		}
	})
	t.Run("DecodeFromBsontype", func(t *testing.T) {
		for _, params := range []struct {
			vr   *bsonrwtest.ValueReaderWriter
			want *timestamppb.Timestamp
		}{
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.DateTime,
					Return:   int64(1653911006000),
				},
				ts,
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Int64,
					Return:   int64(1653911006000),
				},
				ts,
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.String,
					Return:   "2022-05-30T11:43:26Z",
				},
				ts,
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
				&timestamppb.Timestamp{},
			},
		} {
			t.Run(params.vr.Type().String(), func(t *testing.T) {
				c := NewTimestampCodec()
				got := reflect.New(reflect.TypeOf(params.want)).Elem()
				err := c.DecodeValue(bsoncodec.DecodeContext{}, params.vr, got)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, got.Interface(), protocmp.Transform())
			})
		}
	})
}
