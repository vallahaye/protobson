package protobsoncodec

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTimestampCodec(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond)
	t.Run("EncodeToBsontype", func(t *testing.T) {
		for _, params := range []struct {
			vw   *bsonrwtest.ValueReaderWriter
			want bsonrwtest.Invoked
		}{
			{
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Timestamp},
				bsonrwtest.WriteTimestamp,
			},
		} {
			t.Run(params.vw.Type().String(), func(t *testing.T) {
				tsc := NewTimestampCodec()
				v := reflect.ValueOf(timestamppb.New(now))
				if err := tsc.EncodeValue(bsoncodec.EncodeContext{}, params.vw, v); err != nil {
					t.Fatal(err)
				}
				if got := params.vw.Invoked; got != params.want {
					t.Fatalf("Invoked method do not match. got %s; want %s", got, params.want)
				}
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
					Return:   nil,
				},
				&timestamppb.Timestamp{},
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Undefined,
					Return:   nil,
				},
				&timestamppb.Timestamp{},
			},
		} {
			t.Run(params.vr.Type().String(), func(t *testing.T) {
				tsc := NewTimestampCodec()
				got := reflect.New(reflect.TypeOf(params.want)).Elem()
				if err := tsc.DecodeValue(bsoncodec.DecodeContext{}, params.vr, got); err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(got.Interface().(*timestamppb.Timestamp), params.want, cmp.Comparer(proto.Equal)) {
					t.Fatalf("Returned timestamp do not match. got %#v; want %#v", got, params.want)
				}
			})
		}
	})
}
