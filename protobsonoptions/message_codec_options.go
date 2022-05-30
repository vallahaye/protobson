package protobsonoptions

import "go.mongodb.org/mongo-driver/bson/bsonoptions"

var defaultUseProtoNames = false

// MessageCodecOptions represents all possible options for proto.Message encoding and decoding.
type MessageCodecOptions struct {
	UseProtoNames *bool // Specifies if field names should be marshaled/unmarshaled using their proto names. Defaults to false.
	*bsonoptions.StructCodecOptions
}

// SetUseProtoNames specifies if field names should be marshaled/unmarshaled using their proto names. Defaults to false.
func (t *MessageCodecOptions) SetUseProtoNames(b bool) *MessageCodecOptions {
	t.UseProtoNames = &b
	return t
}

// MessageCodec creates a new *MessageCodecOptions.
func MessageCodec() *MessageCodecOptions {
	return &MessageCodecOptions{
		StructCodecOptions: bsonoptions.StructCodec(),
	}
}

// MergeMessageCodecOptions combines the given *MessageCodecOptions into a single *MessageCodecOptions in a last one wins fashion.
func MergeMessageCodecOptions(opts ...*MessageCodecOptions) *MessageCodecOptions {
	msgOpts := &MessageCodecOptions{
		UseProtoNames: &defaultUseProtoNames,
	}
	structOpts := make([]*bsonoptions.StructCodecOptions, 0, len(opts))
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.UseProtoNames != nil {
			msgOpts.UseProtoNames = opt.UseProtoNames
		}
		structOpts = append(structOpts, opt.StructCodecOptions)
	}
	msgOpts.StructCodecOptions = bsonoptions.MergeStructCodecOptions(structOpts...)
	return msgOpts
}
