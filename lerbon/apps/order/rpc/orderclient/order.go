// Code generated by goctl. DO NOT EDIT.
// Source: rpc.proto

package orderclient

import (
	"context"

	"go-zeroProject/lerbon/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	OrderItem      = order.OrderItem
	OrdersRequest  = order.OrdersRequest
	OrdersResponse = order.OrdersResponse

	Order interface {
		Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.Orders(ctx, in, opts...)
}
