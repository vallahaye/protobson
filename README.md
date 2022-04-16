# protobson

[![PkgGoDev](https://pkg.go.dev/badge/github.com/vallahaye/protobson)](https://pkg.go.dev/github.com/vallahaye/protobson) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/vallahaye/protobson) [![GoReportCard](https://goreportcard.com/badge/github.com/vallahaye/protobson)](https://goreportcard.com/report/github.com/vallahaye/protobson) ![GitHub](https://img.shields.io/github/license/vallahaye/protobson)

BSON codecs for Google's protocol buffers.

This library provides add-ons to [go.mongodb.org/mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver) for first-class protobuf support with similar API design, extensive testing and [high reuse of already existing codecs](https://github.com/vallahaye/protobson/blob/master/protobsoncodec/message_codec.go). The following types are currently mapped:

| Protobuf  | MongoDB    |
|-----------|------------|
| [`message`](https://pkg.go.dev/google.golang.org/protobuf/proto#Message) | [Document](https://www.mongodb.com/docs/manual/core/document/) |
| [`google.protobuf.Timestamp`](https://pkg.go.dev/google.golang.org/protobuf/types/known/timestamppb) | [Timestamp](https://www.mongodb.com/docs/manual/reference/bson-types/#timestamps) |

This list will grow as we add support for most [Well-Known Types](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf) as well as [Google APIs Common Types](https://github.com/googleapis/api-common-protos).

## Usage

```bash
$ go get -u github.com/vallahaye/protobson
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
