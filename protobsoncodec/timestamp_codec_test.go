package protobsoncodec

import (
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
)

func TestTimestampCodec(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond)
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
				timestamppb.New(now),
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Timestamp},
				bsonrwtest.WriteTimestamp,
			},
		} {
			t.Run(params.vw.Type().String(), func(t *testing.T) {
				tsc := NewTimestampCodec()
				v := reflect.ValueOf(params.ts)
				err := tsc.EncodeValue(bsoncodec.EncodeContext{}, params.vw, v)
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
					BSONType: bsontype.Timestamp,
					Return: bsoncore.Value{
						Type: bsontype.Timestamp,
						Data: bsoncore.AppendTimestamp(nil, uint32(now.Unix()), uint32(now.Nanosecond())),
					},
				},
				timestamppb.New(now),
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.DateTime,
					Return:   now.UnixMilli(),
				},
				timestamppb.New(now),
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Int64,
					Return:   now.Unix(),
				},
				&timestamppb.Timestamp{
					Seconds: now.Unix(),
				},
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Int32,
					Return:   int32(now.Unix()),
				},
				&timestamppb.Timestamp{
					Seconds: now.Unix(),
				},
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.String,
					Return:   now.Format(time.RFC3339Nano),
				},
				timestamppb.New(now),
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
				tsc := NewTimestampCodec()
				got := reflect.New(reflect.TypeOf(params.want)).Elem()
				err := tsc.DecodeValue(bsoncodec.DecodeContext{}, params.vr, got)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, got.Interface(), protocmp.Transform())
			})
		}
	})
}
