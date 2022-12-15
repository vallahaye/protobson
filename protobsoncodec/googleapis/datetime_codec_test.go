package googleapis

import (
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"
	"gotest.tools/v3/assert"
)

func TestDateTimeCodec(t *testing.T) {
	dt := &datetime.DateTime{
		Year:    2022,
		Month:   5,
		Day:     30,
		Hours:   11,
		Minutes: 43,
		Seconds: 26,
		TimeOffset: &datetime.DateTime_UtcOffset{
			UtcOffset: &durationpb.Duration{},
		},
	}
	t.Run("EncodeToBsontype", func(t *testing.T) {
		for _, params := range []struct {
			dt   *datetime.DateTime
			vw   *bsonrwtest.ValueReaderWriter
			want bsonrwtest.Invoked
		}{
			{
				nil,
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.Null},
				bsonrwtest.WriteNull,
			},
			{
				dt,
				&bsonrwtest.ValueReaderWriter{BSONType: bsontype.DateTime},
				bsonrwtest.WriteDateTime,
			},
		} {
			t.Run(params.vw.Type().String(), func(t *testing.T) {
				c := NewDateTimeCodec()
				v := reflect.ValueOf(params.dt)
				err := c.EncodeValue(bsoncodec.EncodeContext{}, params.vw, v)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, params.vw.Invoked)
			})
		}
	})
	t.Run("DecodeFromBsontype", func(t *testing.T) {
		for _, params := range []struct {
			vr   *bsonrwtest.ValueReaderWriter
			want *datetime.DateTime
		}{
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.DateTime,
					Return:   int64(1653911006000),
				},
				dt,
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.Int64,
					Return:   int64(1653911006000),
				},
				dt,
			},
			{
				&bsonrwtest.ValueReaderWriter{
					BSONType: bsontype.String,
					Return:   "2022-05-30T11:43:26Z",
				},
				dt,
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
				&datetime.DateTime{},
			},
		} {
			t.Run(params.vr.Type().String(), func(t *testing.T) {
				c := NewDateTimeCodec()
				got := reflect.New(reflect.TypeOf(params.want)).Elem()
				err := c.DecodeValue(bsoncodec.DecodeContext{}, params.vr, got)
				assert.NilError(t, err)
				assert.DeepEqual(t, params.want, got.Interface(), protocmp.Transform())
			})
		}
	})
}

func TestDateTimeToTime(t *testing.T) {
	for _, params := range []struct {
		dt   *datetime.DateTime
		want time.Time
	}{
		{
			dt: &datetime.DateTime{
				Year:    2022,
				Month:   5,
				Day:     30,
				Hours:   11,
				Minutes: 43,
				Seconds: 26,
				TimeOffset: &datetime.DateTime_UtcOffset{
					UtcOffset: &durationpb.Duration{},
				},
			},
			want: time.Date(2022, 5, 30, 11, 43, 26, 0, time.FixedZone("", 0)),
		},
		{
			dt: &datetime.DateTime{
				Year:    2022,
				Month:   5,
				Day:     30,
				Hours:   13,
				Minutes: 43,
				Seconds: 26,
				TimeOffset: &datetime.DateTime_UtcOffset{
					UtcOffset: durationpb.New(2 * time.Hour),
				},
			},
			want: time.Date(2022, 5, 30, 13, 43, 26, 0, time.FixedZone("", 2*60*60)),
		},
		{
			dt: &datetime.DateTime{
				Year:    2022,
				Month:   5,
				Day:     30,
				Hours:   11,
				Minutes: 43,
				Seconds: 26,
				TimeOffset: &datetime.DateTime_TimeZone{
					TimeZone: &datetime.TimeZone{
						Id: "Etc/UTC",
					},
				},
			},
			want: time.Date(2022, 5, 30, 11, 43, 26, 0, time.FixedZone("", 0)),
		},
		{
			dt: &datetime.DateTime{
				Year:    2022,
				Month:   5,
				Day:     30,
				Hours:   13,
				Minutes: 43,
				Seconds: 26,
				TimeOffset: &datetime.DateTime_TimeZone{
					TimeZone: &datetime.TimeZone{
						Id: "Etc/GMT-2",
					},
				},
			},
			want: time.Date(2022, 5, 30, 13, 43, 26, 0, time.FixedZone("", 2*60*60)),
		},
	} {
		got, err := dateTimeToTime(params.dt)
		assert.NilError(t, err)
		assert.DeepEqual(t, params.want, got)
	}
}

func TestTimeToDateTime(t *testing.T) {
	for _, params := range []struct {
		t    time.Time
		want *datetime.DateTime
	}{
		{
			t: time.Date(2022, 5, 30, 11, 43, 26, 0, time.UTC),
			want: &datetime.DateTime{
				Year:    2022,
				Month:   5,
				Day:     30,
				Hours:   11,
				Minutes: 43,
				Seconds: 26,
				TimeOffset: &datetime.DateTime_UtcOffset{
					UtcOffset: &durationpb.Duration{},
				},
			},
		},
		{
			t: time.Date(2022, 5, 30, 13, 43, 26, 0, time.FixedZone("UTC+2", 2*60*60)),
			want: &datetime.DateTime{
				Year:    2022,
				Month:   5,
				Day:     30,
				Hours:   13,
				Minutes: 43,
				Seconds: 26,
				TimeOffset: &datetime.DateTime_UtcOffset{
					UtcOffset: durationpb.New(2 * time.Hour),
				},
			},
		},
	} {
		got := timeToDateTime(params.t)
		assert.DeepEqual(t, params.want, got, protocmp.Transform())
	}
}
