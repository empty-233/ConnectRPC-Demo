// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: calculator/v1/calculator.proto

package calculatorv1connect

import (
	v1 "ConnectRPC-Demo/gen/calculator/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// CalculatorServiceName is the fully-qualified name of the CalculatorService service.
	CalculatorServiceName = "calculator.v1.CalculatorService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CalculatorServiceAddProcedure is the fully-qualified name of the CalculatorService's Add RPC.
	CalculatorServiceAddProcedure = "/calculator.v1.CalculatorService/Add"
	// CalculatorServiceSubtractProcedure is the fully-qualified name of the CalculatorService's
	// Subtract RPC.
	CalculatorServiceSubtractProcedure = "/calculator.v1.CalculatorService/Subtract"
	// CalculatorServiceMultiplyProcedure is the fully-qualified name of the CalculatorService's
	// Multiply RPC.
	CalculatorServiceMultiplyProcedure = "/calculator.v1.CalculatorService/Multiply"
	// CalculatorServiceDivideProcedure is the fully-qualified name of the CalculatorService's Divide
	// RPC.
	CalculatorServiceDivideProcedure = "/calculator.v1.CalculatorService/Divide"
)

// CalculatorServiceClient is a client for the calculator.v1.CalculatorService service.
type CalculatorServiceClient interface {
	Add(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error)
	Subtract(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error)
	Multiply(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error)
	Divide(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error)
}

// NewCalculatorServiceClient constructs a client for the calculator.v1.CalculatorService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCalculatorServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CalculatorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	calculatorServiceMethods := v1.File_calculator_v1_calculator_proto.Services().ByName("CalculatorService").Methods()
	return &calculatorServiceClient{
		add: connect.NewClient[v1.OperateRequest, v1.OperateResponse](
			httpClient,
			baseURL+CalculatorServiceAddProcedure,
			connect.WithSchema(calculatorServiceMethods.ByName("Add")),
			connect.WithClientOptions(opts...),
		),
		subtract: connect.NewClient[v1.OperateRequest, v1.OperateResponse](
			httpClient,
			baseURL+CalculatorServiceSubtractProcedure,
			connect.WithSchema(calculatorServiceMethods.ByName("Subtract")),
			connect.WithClientOptions(opts...),
		),
		multiply: connect.NewClient[v1.OperateRequest, v1.OperateResponse](
			httpClient,
			baseURL+CalculatorServiceMultiplyProcedure,
			connect.WithSchema(calculatorServiceMethods.ByName("Multiply")),
			connect.WithClientOptions(opts...),
		),
		divide: connect.NewClient[v1.OperateRequest, v1.OperateResponse](
			httpClient,
			baseURL+CalculatorServiceDivideProcedure,
			connect.WithSchema(calculatorServiceMethods.ByName("Divide")),
			connect.WithClientOptions(opts...),
		),
	}
}

// calculatorServiceClient implements CalculatorServiceClient.
type calculatorServiceClient struct {
	add      *connect.Client[v1.OperateRequest, v1.OperateResponse]
	subtract *connect.Client[v1.OperateRequest, v1.OperateResponse]
	multiply *connect.Client[v1.OperateRequest, v1.OperateResponse]
	divide   *connect.Client[v1.OperateRequest, v1.OperateResponse]
}

// Add calls calculator.v1.CalculatorService.Add.
func (c *calculatorServiceClient) Add(ctx context.Context, req *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error) {
	return c.add.CallUnary(ctx, req)
}

// Subtract calls calculator.v1.CalculatorService.Subtract.
func (c *calculatorServiceClient) Subtract(ctx context.Context, req *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error) {
	return c.subtract.CallUnary(ctx, req)
}

// Multiply calls calculator.v1.CalculatorService.Multiply.
func (c *calculatorServiceClient) Multiply(ctx context.Context, req *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error) {
	return c.multiply.CallUnary(ctx, req)
}

// Divide calls calculator.v1.CalculatorService.Divide.
func (c *calculatorServiceClient) Divide(ctx context.Context, req *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error) {
	return c.divide.CallUnary(ctx, req)
}

// CalculatorServiceHandler is an implementation of the calculator.v1.CalculatorService service.
type CalculatorServiceHandler interface {
	Add(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error)
	Subtract(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error)
	Multiply(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error)
	Divide(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error)
}

// NewCalculatorServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCalculatorServiceHandler(svc CalculatorServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	calculatorServiceMethods := v1.File_calculator_v1_calculator_proto.Services().ByName("CalculatorService").Methods()
	calculatorServiceAddHandler := connect.NewUnaryHandler(
		CalculatorServiceAddProcedure,
		svc.Add,
		connect.WithSchema(calculatorServiceMethods.ByName("Add")),
		connect.WithHandlerOptions(opts...),
	)
	calculatorServiceSubtractHandler := connect.NewUnaryHandler(
		CalculatorServiceSubtractProcedure,
		svc.Subtract,
		connect.WithSchema(calculatorServiceMethods.ByName("Subtract")),
		connect.WithHandlerOptions(opts...),
	)
	calculatorServiceMultiplyHandler := connect.NewUnaryHandler(
		CalculatorServiceMultiplyProcedure,
		svc.Multiply,
		connect.WithSchema(calculatorServiceMethods.ByName("Multiply")),
		connect.WithHandlerOptions(opts...),
	)
	calculatorServiceDivideHandler := connect.NewUnaryHandler(
		CalculatorServiceDivideProcedure,
		svc.Divide,
		connect.WithSchema(calculatorServiceMethods.ByName("Divide")),
		connect.WithHandlerOptions(opts...),
	)
	return "/calculator.v1.CalculatorService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CalculatorServiceAddProcedure:
			calculatorServiceAddHandler.ServeHTTP(w, r)
		case CalculatorServiceSubtractProcedure:
			calculatorServiceSubtractHandler.ServeHTTP(w, r)
		case CalculatorServiceMultiplyProcedure:
			calculatorServiceMultiplyHandler.ServeHTTP(w, r)
		case CalculatorServiceDivideProcedure:
			calculatorServiceDivideHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCalculatorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCalculatorServiceHandler struct{}

func (UnimplementedCalculatorServiceHandler) Add(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("calculator.v1.CalculatorService.Add is not implemented"))
}

func (UnimplementedCalculatorServiceHandler) Subtract(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("calculator.v1.CalculatorService.Subtract is not implemented"))
}

func (UnimplementedCalculatorServiceHandler) Multiply(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("calculator.v1.CalculatorService.Multiply is not implemented"))
}

func (UnimplementedCalculatorServiceHandler) Divide(context.Context, *connect.Request[v1.OperateRequest]) (*connect.Response[v1.OperateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("calculator.v1.CalculatorService.Divide is not implemented"))
}
