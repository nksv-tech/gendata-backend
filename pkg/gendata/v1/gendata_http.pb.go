// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package gendata

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type GenDataServiceHTTPServer interface {
	Gen(context.Context, *GenRequest) (*GenResponse, error)
	PredefinedSettings(context.Context, *PredefinedSettingsRequest) (*PredefinedSettingsResponse, error)
}

func RegisterGenDataServiceHTTPServer(s *http.Server, srv GenDataServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/gen", _GenDataService_Gen0_HTTP_Handler(srv))
	r.GET("/v1/predefined-settings", _GenDataService_PredefinedSettings0_HTTP_Handler(srv))
}

func _GenDataService_Gen0_HTTP_Handler(srv GenDataServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/gendata.v1.GenDataService/Gen")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Gen(ctx, req.(*GenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GenResponse)
		return ctx.Result(200, reply)
	}
}

func _GenDataService_PredefinedSettings0_HTTP_Handler(srv GenDataServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PredefinedSettingsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/gendata.v1.GenDataService/PredefinedSettings")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PredefinedSettings(ctx, req.(*PredefinedSettingsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PredefinedSettingsResponse)
		return ctx.Result(200, reply)
	}
}

type GenDataServiceHTTPClient interface {
	Gen(ctx context.Context, req *GenRequest, opts ...http.CallOption) (rsp *GenResponse, err error)
	PredefinedSettings(ctx context.Context, req *PredefinedSettingsRequest, opts ...http.CallOption) (rsp *PredefinedSettingsResponse, err error)
}

type GenDataServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewGenDataServiceHTTPClient(client *http.Client) GenDataServiceHTTPClient {
	return &GenDataServiceHTTPClientImpl{client}
}

func (c *GenDataServiceHTTPClientImpl) Gen(ctx context.Context, in *GenRequest, opts ...http.CallOption) (*GenResponse, error) {
	var out GenResponse
	pattern := "/v1/gen"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/gendata.v1.GenDataService/Gen"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GenDataServiceHTTPClientImpl) PredefinedSettings(ctx context.Context, in *PredefinedSettingsRequest, opts ...http.CallOption) (*PredefinedSettingsResponse, error) {
	var out PredefinedSettingsResponse
	pattern := "/v1/predefined-settings"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/gendata.v1.GenDataService/PredefinedSettings"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
