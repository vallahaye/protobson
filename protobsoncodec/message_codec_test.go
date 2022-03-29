package protobsoncodec

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

func TestStructTagParsers(t *testing.T) {
	for _, params := range []struct {
		name   string
		sf     reflect.StructField
		want   bsoncodec.StructTags
		parser bsoncodec.StructTagParserFunc
	}{
		{
			"JSONPBFallback no bson tag",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag("bar"),
			},
			bsoncodec.StructTags{Name: "bar"},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback empty",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(""),
			},
			bsoncodec.StructTags{Name: "foo"},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback tag only dash",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag("-"),
			},
			bsoncodec.StructTags{Skip: true},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback bson tag only dash",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`bson:"-"`),
			},
			bsoncodec.StructTags{Skip: true},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback all options",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`bar,omitempty,minsize,truncate,inline`),
			},
			bsoncodec.StructTags{Name: "bar", OmitEmpty: true, MinSize: true, Truncate: true, Inline: true},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback all options default name",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`,omitempty,minsize,truncate,inline`),
			},
			bsoncodec.StructTags{Name: "foo", OmitEmpty: true, MinSize: true, Truncate: true, Inline: true},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback bson tag all options",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`bson:"bar,omitempty,minsize,truncate,inline"`),
			},
			bsoncodec.StructTags{Name: "bar", OmitEmpty: true, MinSize: true, Truncate: true, Inline: true},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback bson tag all options default name",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`bson:",omitempty,minsize,truncate,inline"`),
			},
			bsoncodec.StructTags{Name: "foo", OmitEmpty: true, MinSize: true, Truncate: true, Inline: true},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback ignore xml",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`xml:"bar"`),
			},
			bsoncodec.StructTags{Name: "foo"},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback protobuf tag json name",
			reflect.StructField{
				Name: "FooBar",
				Tag:  reflect.StructTag(`protobuf:"bytes,1,opt,name=foo_bar,json=fooBar,proto3"`),
			},
			bsoncodec.StructTags{Name: "fooBar"},
			JSONPBFallbackStructTagParser,
		},
		{
			"JSONPBFallback protobuf tag proto name",
			reflect.StructField{
				Name: "FooBar",
				Tag:  reflect.StructTag(`protobuf:"bytes,1,opt,name=foo_bar,proto3"`),
			},
			bsoncodec.StructTags{Name: "foo_bar"},
			JSONPBFallbackStructTagParser,
		},
		{
			"ProtoNameFallback no bson tag",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag("bar"),
			},
			bsoncodec.StructTags{Name: "bar"},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback empty",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(""),
			},
			bsoncodec.StructTags{Name: "foo"},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback tag only dash",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag("-"),
			},
			bsoncodec.StructTags{Skip: true},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback bson tag only dash",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`bson:"-"`),
			},
			bsoncodec.StructTags{Skip: true},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback all options",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`bar,omitempty,minsize,truncate,inline`),
			},
			bsoncodec.StructTags{Name: "bar", OmitEmpty: true, MinSize: true, Truncate: true, Inline: true},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback all options default name",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`,omitempty,minsize,truncate,inline`),
			},
			bsoncodec.StructTags{Name: "foo", OmitEmpty: true, MinSize: true, Truncate: true, Inline: true},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback bson tag all options",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`bson:"bar,omitempty,minsize,truncate,inline"`),
			},
			bsoncodec.StructTags{Name: "bar", OmitEmpty: true, MinSize: true, Truncate: true, Inline: true},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback bson tag all options default name",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`bson:",omitempty,minsize,truncate,inline"`),
			},
			bsoncodec.StructTags{Name: "foo", OmitEmpty: true, MinSize: true, Truncate: true, Inline: true},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback ignore xml",
			reflect.StructField{
				Name: "Foo",
				Tag:  reflect.StructTag(`xml:"bar"`),
			},
			bsoncodec.StructTags{Name: "foo"},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback protobuf tag json name",
			reflect.StructField{
				Name: "FooBar",
				Tag:  reflect.StructTag(`protobuf:"bytes,1,opt,name=foo_bar,json=fooBar,proto3"`),
			},
			bsoncodec.StructTags{Name: "foo_bar"},
			ProtoNameFallbackStructTagParser,
		},
		{
			"ProtoNameFallback protobuf tag proto name",
			reflect.StructField{
				Name: "FooBar",
				Tag:  reflect.StructTag(`protobuf:"bytes,1,opt,name=foo_bar,proto3"`),
			},
			bsoncodec.StructTags{Name: "foo_bar"},
			ProtoNameFallbackStructTagParser,
		},
	} {
		t.Run(params.name, func(t *testing.T) {
			got, err := params.parser(params.sf)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, params.want) {
				t.Fatalf("Returned struct tags do not match. got %#v; want %#v", got, params.want)
			}
		})
	}
}
