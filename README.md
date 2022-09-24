# protobson

[![PkgGoDev](https://pkg.go.dev/badge/go.vallahaye.net/protobson)](https://pkg.go.dev/go.vallahaye.net/protobson) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/vallahaye/protobson) [![GoReportCard](https://goreportcard.com/badge/github.com/vallahaye/protobson)](https://goreportcard.com/badge/github.com/vallahaye/protobson) ![GitHub](https://img.shields.io/github/license/vallahaye/protobson)

BSON codecs for Google's protocol buffers.

This library provides add-ons to [go.mongodb.org/mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver) for first-class protobuf support with similar API design, extensive testing and [high reuse of already existing codecs](https://github.com/vallahaye/protobson/blob/main/protobsoncodec/message_codec.go). The following types are currently mapped:

| Protobuf  | MongoDB    |
|-----------|------------|
| [`message`](https://pkg.go.dev/google.golang.org/protobuf/proto#Message) | [Document](https://www.mongodb.com/docs/manual/core/document/) |
| [`google.protobuf.Timestamp`](https://pkg.go.dev/google.golang.org/protobuf/types/known/timestamppb#Timestamp) | [Timestamp](https://www.mongodb.com/docs/manual/reference/bson-types/#timestamps) |
| [`google.protobuf.Duration`](https://pkg.go.dev/google.golang.org/protobuf/types/known/durationpb#Duration) | [64-bit integer](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |
| [`google.protobuf.BoolValue`](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb#BoolValue) | [Boolean](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |
| [`google.protobuf.BytesValue`](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb#BytesValue) | [Binary](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |
| [`google.protobuf.DoubleValue`](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb#DoubleValue) | [Double](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |
| [`google.protobuf.FloatValue`](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb#FloatValue) | [Double](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |
| [`google.protobuf.Int32Value`](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb#Int32Value) | [32-bit integer](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |
| [`google.protobuf.Int64Value`](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb#Int64Value) | [64-bit integer](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |
| [`google.protobuf.StringValue`](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb#StringValue) | [String](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |
| [`google.protobuf.UInt32Value`](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb#UInt32Value) | [32-bit integer](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |
| [`google.protobuf.UInt64Value`](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb#UInt64Value) | [64-bit integer](https://www.mongodb.com/docs/manual/reference/bson-types/#bson-types) |

This list will grow as we add support for most [Well-Known Types](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf) as well as [Google APIs Common Types](https://github.com/googleapis/api-common-protos).

## Usage

```
$ go get go.vallahaye.net/protobson
```

```go
// Set client options
clientOptions := options.Client().
  SetRegistry(protobson.DefaultRegistry).
  ApplyURI("mongodb://localhost:27017")

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)
if err != nil {
  log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)
if err != nil {
  log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")
```
