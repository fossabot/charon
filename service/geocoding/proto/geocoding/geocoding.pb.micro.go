// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/geocoding/proto/geocoding/geocoding.proto

package geocoding

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Geocoding service

type GeocodingService interface {
	Forward(ctx context.Context, in *Search, opts ...client.CallOption) (*ForwardResponse, error)
	Reverse(ctx context.Context, in *Coordinates, opts ...client.CallOption) (*ReverseResponse, error)
}

type geocodingService struct {
	c    client.Client
	name string
}

func NewGeocodingService(name string, c client.Client) GeocodingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "geocoding"
	}
	return &geocodingService{
		c:    c,
		name: name,
	}
}

func (c *geocodingService) Forward(ctx context.Context, in *Search, opts ...client.CallOption) (*ForwardResponse, error) {
	req := c.c.NewRequest(c.name, "Geocoding.Forward", in)
	out := new(ForwardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geocodingService) Reverse(ctx context.Context, in *Coordinates, opts ...client.CallOption) (*ReverseResponse, error) {
	req := c.c.NewRequest(c.name, "Geocoding.Reverse", in)
	out := new(ReverseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Geocoding service

type GeocodingHandler interface {
	Forward(context.Context, *Search, *ForwardResponse) error
	Reverse(context.Context, *Coordinates, *ReverseResponse) error
}

func RegisterGeocodingHandler(s server.Server, hdlr GeocodingHandler, opts ...server.HandlerOption) error {
	type geocoding interface {
		Forward(ctx context.Context, in *Search, out *ForwardResponse) error
		Reverse(ctx context.Context, in *Coordinates, out *ReverseResponse) error
	}
	type Geocoding struct {
		geocoding
	}
	h := &geocodingHandler{hdlr}
	return s.Handle(s.NewHandler(&Geocoding{h}, opts...))
}

type geocodingHandler struct {
	GeocodingHandler
}

func (h *geocodingHandler) Forward(ctx context.Context, in *Search, out *ForwardResponse) error {
	return h.GeocodingHandler.Forward(ctx, in, out)
}

func (h *geocodingHandler) Reverse(ctx context.Context, in *Coordinates, out *ReverseResponse) error {
	return h.GeocodingHandler.Reverse(ctx, in, out)
}
