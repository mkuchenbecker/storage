package service

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/mkuchenbecker/storage/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

type service struct {
	data map[string]*any.Any
}

// New creates a new default api.StorageServer.
func New() *service {
	return &service{
		data: make(map[string]*any.Any),
	}
}

func (s *service) Put(ctx context.Context, req *api.PutRequest) (*api.PutResponse, error) {
	s.data[req.Key.Value] = req.Value
	return &api.PutResponse{}, nil
}

func (s *service) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	data, ok := s.data[req.Key.Value]
	if !ok {
		return &api.GetResponse{}, status.Error(codes.NotFound, "not found")
	}
	return &api.GetResponse{Value: data}, nil
}

func (s *service) Start(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	grpcServer := grpc.NewServer()
	api.RegisterStorageServer(grpcServer, New())
	grpcServer.Serve(lis)
	return nil
}
