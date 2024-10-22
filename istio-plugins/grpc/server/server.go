package server

import (
	"errors"
	"log"
	"net"

	"github.com/micro/go-micro/server"
	"google.golang.org/grpc"
)

type grpcServer struct {
	server  *grpc.Server
	options server.Options
}

func newGRPCServer(mopts []server.Option, opts ...grpc.ServerOption) server.Server {
	return &grpcServer{
		server:  grpc.NewServer(opts...),
		options: newOptions(mopts...),
	}
}

func (*grpcServer) Options() server.Options {
	return server.Options{}
}

func (*grpcServer) Init(...server.Option) error {
	return nil
}

func (s *grpcServer) Handle(h server.Handler) error {
	if h.Handler() == nil {
		return errors.New("handler must not be nil")
	}

	hdl := h.(*Handler)
	s.server.RegisterService(hdl.Desc, hdl.Server)
	return nil
}

func (*grpcServer) NewHandler(interface{}, ...server.HandlerOption) server.Handler {
	return nil
}

func (*grpcServer) NewSubscriber(string, interface{}, ...server.SubscriberOption) server.Subscriber {
	return nil
}

func (*grpcServer) Subscribe(server.Subscriber) error {
	return nil
}

func (*grpcServer) Register() error {
	return nil
}

func (*grpcServer) Deregister() error {
	return nil
}

func (s *grpcServer) Start() error {
	// Run server
	lis, err := net.Listen("tcp", s.options.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
	return nil
}

func (s *grpcServer) Stop() error {
	s.server.Stop()
	return nil
}

func (*grpcServer) String() string {
	return "istio-grpc"
}

func NewServer(mopts []server.Option, opts ...grpc.ServerOption) server.Server {
	return newGRPCServer(mopts, opts...)
}
