// Code generated by goa v3.3.1, DO NOT EDIT.
//
// codeset gRPC client
//
// Command:
// $ goa gen github.com/fuseml/fuseml-core/design

package client

import (
	"context"

	codesetpb "github.com/fuseml/fuseml-core/gen/grpc/codeset/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli codesetpb.CodesetClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the codeset service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: codesetpb.NewCodesetClient(cc),
		opts:    opts,
	}
}

// List calls the "List" function in codesetpb.CodesetClient interface.
func (c *Client) List() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildListFunc(c.grpccli, c.opts...),
			EncodeListRequest,
			DecodeListResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Register calls the "Register" function in codesetpb.CodesetClient interface.
func (c *Client) Register() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildRegisterFunc(c.grpccli, c.opts...),
			EncodeRegisterRequest,
			DecodeRegisterResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Get calls the "Get" function in codesetpb.CodesetClient interface.
func (c *Client) Get() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildGetFunc(c.grpccli, c.opts...),
			EncodeGetRequest,
			DecodeGetResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}
