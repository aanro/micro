package server

import (
	"github.com/micro/go-micro/server"
)

func newOptions(opt ...server.Option) server.Options {
	opts := server.Options{
		Metadata: map[string]string{},
	}

	for _, o := range opt {
		o(&opts)
	}

	if len(opts.Address) == 0 {
		opts.Address = ":10088"
	}

	if len(opts.Name) == 0 {
		opts.Name = "istio-grpc-server"
	}

	return opts
}
