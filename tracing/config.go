package main

type Jaeger struct {
	Enabled       bool    `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Host          string  `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port          string  `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	SamplingType  string  `protobuf:"bytes,4,opt,name=samplingType,json=sampling_type,proto3" json:"samplingType,omitempty"`
	SamplerRatio  float32 `protobuf:"fixed32,5,opt,name=samplerRatio,json=sampler_ratio,proto3" json:"samplerRatio,omitempty"`
	SamplingParam float32 `protobuf:"fixed32,6,opt,name=samplingParam,json=sampling_param,proto3" json:"samplingParam,omitempty"`
	LogSpans      bool    `protobuf:"varint,7,opt,name=logSpans,json=log_spans,proto3" json:"logSpans,omitempty"`
	FlushInterval uint32  `protobuf:"varint,8,opt,name=flushInterval,json=flush_interval,proto3" json:"flushInterval,omitempty"`
}

type Otlp struct {
	Enabled      bool    `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Endpoint     string  `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Token        string  `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	SamplerRatio float32 `protobuf:"fixed32,4,opt,name=samplerRatio,json=sampler_ratio,proto3" json:"samplerRatio,omitempty"`
}
