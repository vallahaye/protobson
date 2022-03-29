package protobsonoptions

import "go.mongodb.org/mongo-driver/bson/bsonoptions"

var defaultUseProtoNames = false

type MessageCodecOptions struct {
	UseProtoNames *bool
	*bsonoptions.StructCodecOptions
}

func (t *MessageCodecOptions) SetUseProtoNames(b bool) *MessageCodecOptions {
	t.UseProtoNames = &b
	return t
}

func MessageCodec() *MessageCodecOptions {
	return &MessageCodecOptions{
		StructCodecOptions: bsonoptions.StructCodec(),
	}
}

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
