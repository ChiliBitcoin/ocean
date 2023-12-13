// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package oceanv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	// Notifies about events related to wallet transactions.
	TransactionNotifications(ctx context.Context, in *TransactionNotificationsRequest, opts ...grpc.CallOption) (NotificationService_TransactionNotificationsClient, error)
	// Notifies about events realted to wallet utxos.
	UtxosNotifications(ctx context.Context, in *UtxosNotificationsRequest, opts ...grpc.CallOption) (NotificationService_UtxosNotificationsClient, error)
	// Adds a webhook registered for some kind of event.
	AddWebhook(ctx context.Context, in *AddWebhookRequest, opts ...grpc.CallOption) (*AddWebhookResponse, error)
	// Removes some previously added webhook.
	RemoveWebhook(ctx context.Context, in *RemoveWebhookRequest, opts ...grpc.CallOption) (*RemoveWebhookResponse, error)
	// Returns registered webhooks.
	ListWebhooks(ctx context.Context, in *ListWebhooksRequest, opts ...grpc.CallOption) (*ListWebhooksResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) TransactionNotifications(ctx context.Context, in *TransactionNotificationsRequest, opts ...grpc.CallOption) (NotificationService_TransactionNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationService_ServiceDesc.Streams[0], "/ocean.v1.NotificationService/TransactionNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceTransactionNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_TransactionNotificationsClient interface {
	Recv() (*TransactionNotificationsResponse, error)
	grpc.ClientStream
}

type notificationServiceTransactionNotificationsClient struct {
	grpc.ClientStream
}

func (x *notificationServiceTransactionNotificationsClient) Recv() (*TransactionNotificationsResponse, error) {
	m := new(TransactionNotificationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationServiceClient) UtxosNotifications(ctx context.Context, in *UtxosNotificationsRequest, opts ...grpc.CallOption) (NotificationService_UtxosNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationService_ServiceDesc.Streams[1], "/ocean.v1.NotificationService/UtxosNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceUtxosNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_UtxosNotificationsClient interface {
	Recv() (*UtxosNotificationsResponse, error)
	grpc.ClientStream
}

type notificationServiceUtxosNotificationsClient struct {
	grpc.ClientStream
}

func (x *notificationServiceUtxosNotificationsClient) Recv() (*UtxosNotificationsResponse, error) {
	m := new(UtxosNotificationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationServiceClient) AddWebhook(ctx context.Context, in *AddWebhookRequest, opts ...grpc.CallOption) (*AddWebhookResponse, error) {
	out := new(AddWebhookResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.NotificationService/AddWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) RemoveWebhook(ctx context.Context, in *RemoveWebhookRequest, opts ...grpc.CallOption) (*RemoveWebhookResponse, error) {
	out := new(RemoveWebhookResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.NotificationService/RemoveWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ListWebhooks(ctx context.Context, in *ListWebhooksRequest, opts ...grpc.CallOption) (*ListWebhooksResponse, error) {
	out := new(ListWebhooksResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.NotificationService/ListWebhooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations should embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	// Notifies about events related to wallet transactions.
	TransactionNotifications(*TransactionNotificationsRequest, NotificationService_TransactionNotificationsServer) error
	// Notifies about events realted to wallet utxos.
	UtxosNotifications(*UtxosNotificationsRequest, NotificationService_UtxosNotificationsServer) error
	// Adds a webhook registered for some kind of event.
	AddWebhook(context.Context, *AddWebhookRequest) (*AddWebhookResponse, error)
	// Removes some previously added webhook.
	RemoveWebhook(context.Context, *RemoveWebhookRequest) (*RemoveWebhookResponse, error)
	// Returns registered webhooks.
	ListWebhooks(context.Context, *ListWebhooksRequest) (*ListWebhooksResponse, error)
}

// UnimplementedNotificationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) TransactionNotifications(*TransactionNotificationsRequest, NotificationService_TransactionNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method TransactionNotifications not implemented")
}
func (UnimplementedNotificationServiceServer) UtxosNotifications(*UtxosNotificationsRequest, NotificationService_UtxosNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method UtxosNotifications not implemented")
}
func (UnimplementedNotificationServiceServer) AddWebhook(context.Context, *AddWebhookRequest) (*AddWebhookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWebhook not implemented")
}
func (UnimplementedNotificationServiceServer) RemoveWebhook(context.Context, *RemoveWebhookRequest) (*RemoveWebhookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWebhook not implemented")
}
func (UnimplementedNotificationServiceServer) ListWebhooks(context.Context, *ListWebhooksRequest) (*ListWebhooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWebhooks not implemented")
}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_TransactionNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionNotificationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).TransactionNotifications(m, &notificationServiceTransactionNotificationsServer{stream})
}

type NotificationService_TransactionNotificationsServer interface {
	Send(*TransactionNotificationsResponse) error
	grpc.ServerStream
}

type notificationServiceTransactionNotificationsServer struct {
	grpc.ServerStream
}

func (x *notificationServiceTransactionNotificationsServer) Send(m *TransactionNotificationsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NotificationService_UtxosNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UtxosNotificationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).UtxosNotifications(m, &notificationServiceUtxosNotificationsServer{stream})
}

type NotificationService_UtxosNotificationsServer interface {
	Send(*UtxosNotificationsResponse) error
	grpc.ServerStream
}

type notificationServiceUtxosNotificationsServer struct {
	grpc.ServerStream
}

func (x *notificationServiceUtxosNotificationsServer) Send(m *UtxosNotificationsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NotificationService_AddWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).AddWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.NotificationService/AddWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).AddWebhook(ctx, req.(*AddWebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_RemoveWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveWebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).RemoveWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.NotificationService/RemoveWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).RemoveWebhook(ctx, req.(*RemoveWebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ListWebhooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWebhooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ListWebhooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.NotificationService/ListWebhooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ListWebhooks(ctx, req.(*ListWebhooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocean.v1.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddWebhook",
			Handler:    _NotificationService_AddWebhook_Handler,
		},
		{
			MethodName: "RemoveWebhook",
			Handler:    _NotificationService_RemoveWebhook_Handler,
		},
		{
			MethodName: "ListWebhooks",
			Handler:    _NotificationService_ListWebhooks_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransactionNotifications",
			Handler:       _NotificationService_TransactionNotifications_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UtxosNotifications",
			Handler:       _NotificationService_UtxosNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ocean/v1/notification.proto",
}
