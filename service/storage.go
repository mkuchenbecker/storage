package service

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/mkuchenbecker/storage/api"
	"github.com/mkuchenbecker/storage/service/datamodel"
	"github.com/mkuchenbecker/storage/service/status"
	"github.com/pkg/errors"
)

type DataBackend interface {
	Save(ctx context.Context, item *datamodel.Item) error
	Get(ctx context.Context, key string) (*datamodel.Item, error)
}

type mapBackend map[string]*any.Any

func NewMapBackend() DataBackend {
	out := make(map[string]*any.Any)
	return mapBackend(out)
}

type service struct {
	primary DataBackend
}

// New creates a new default api.StorageServer.
func New(primary DataBackend) api.StorageServer {
	return &service{
		primary: primary,
	}
}

func (s mapBackend) Save(ctx context.Context, item *datamodel.Item) error {
	s[item.Key] = item.Value
	return nil
}

func (s mapBackend) Get(ctx context.Context, key string) (*datamodel.Item, error) {
	value, ok := s[key]
	if !ok {
		return nil, errors.Wrapf(status.ErrNotFound, "key '%s' was not found", key)
	}
	return &datamodel.Item{Key: key, Value: value}, nil
}

func (s *service) Put(ctx context.Context, req *api.PutRequest) (*api.PutResponse, error) {
	item := datamodel.Item{Key: req.Key.Value, Value: req.Value}
	err := s.primary.Save(ctx, &item)
	return &api.PutResponse{}, err
}

func (s *service) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	item, err := s.primary.Get(ctx, req.Key.Value)
	if err != nil {
		if errors.Cause(err) == status.ErrNotFound {
			return &api.GetResponse{}, grpcStatus.Error(codes.NotFound, err.Error())
		}
		return &api.GetResponse{}, err
	}
	return &api.GetResponse{Value: item.Value}, nil
}

func StartService(server api.StorageServer, port int) error {
	glog.Infof("Starting Service on Port: %d", port)
	defer glog.Flush()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	grpcServer := grpc.NewServer()
	api.RegisterStorageServer(grpcServer, server)
	return grpcServer.Serve(lis)
}
