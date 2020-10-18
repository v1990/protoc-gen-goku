// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: order/order.proto

package proto_order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/v1990/protoc-gen-goku/examples/eshop/proto/out_go/common"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	// 创建订单
	Create(ctx context.Context, in *CreateOrderReq, opts ...client.CallOption) (*CreateOrderResp, error)
	// 订单详情
	Detail(ctx context.Context, in *OrderDetailReq, opts ...client.CallOption) (*OrderDetailResp, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) Create(ctx context.Context, in *CreateOrderReq, opts ...client.CallOption) (*CreateOrderResp, error) {
	req := c.c.NewRequest(c.name, "Order.Create", in)
	out := new(CreateOrderResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Detail(ctx context.Context, in *OrderDetailReq, opts ...client.CallOption) (*OrderDetailResp, error) {
	req := c.c.NewRequest(c.name, "Order.Detail", in)
	out := new(OrderDetailResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	// 创建订单
	Create(context.Context, *CreateOrderReq, *CreateOrderResp) error
	// 订单详情
	Detail(context.Context, *OrderDetailReq, *OrderDetailResp) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		Create(ctx context.Context, in *CreateOrderReq, out *CreateOrderResp) error
		Detail(ctx context.Context, in *OrderDetailReq, out *OrderDetailResp) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) Create(ctx context.Context, in *CreateOrderReq, out *CreateOrderResp) error {
	return h.OrderHandler.Create(ctx, in, out)
}

func (h *orderHandler) Detail(ctx context.Context, in *OrderDetailReq, out *OrderDetailResp) error {
	return h.OrderHandler.Detail(ctx, in, out)
}