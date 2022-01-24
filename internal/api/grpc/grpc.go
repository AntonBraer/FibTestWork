package fibgrpc

import (
	"context"
	"fbsTest/internal/api/grpc/proto"
	"fbsTest/internal/service"
	"net"

	"google.golang.org/grpc"
)

type ServerGRPC struct {
	service service.FibService

	fib.UnimplementedFibServer
}

func NewServer(service service.FibService) *ServerGRPC {
	return &ServerGRPC{
		service:                service,
		UnimplementedFibServer: fib.UnimplementedFibServer{},
	}
}

func (s *ServerGRPC) GetFibSeq(ctx context.Context, req *fib.FibRequest) (*fib.FibResponse, error) {
	res, err := s.service.GetFibSeq(ctx, int(req.Start), int(req.End))
	if err != nil {
		return nil, err
	}
	return &fib.FibResponse{Res: res}, nil
}

func (s *ServerGRPC) Run(serv *grpc.Server) error {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}
	fib.RegisterFibServer(serv, s)
	return serv.Serve(listener)
}
