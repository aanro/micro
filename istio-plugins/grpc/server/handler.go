package server

import (
	"github.com/micro/go-micro/registry"
	mserver "github.com/micro/go-micro/server"
	"google.golang.org/grpc"
)

type Handler struct {
	Desc   *grpc.ServiceDesc
	Server interface{}
}

func (h *Handler) Name() string {
	return "grpc_handler"
}

func (h *Handler) Handler() interface{} {
	return h.Server
}

func (h *Handler) Endpoints() []*registry.Endpoint {
	return nil
}

func (h *Handler) Options() mserver.HandlerOptions {
	return mserver.HandlerOptions{}
}
