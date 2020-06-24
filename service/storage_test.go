package service

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes"
	"github.com/mkuchenbecker/storage/api"
	testing_model "github.com/mkuchenbecker/storage/testing/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "google.golang.org/protobuf/testing/protocmp"
)

func TestStorageService_Put_Get_Success(t *testing.T) {
	service := New()

	key := api.Key{Value: "qux"}
	originalFoo := &testing_model.Foo{Bar: "baz"}
	any, err := ptypes.MarshalAny(originalFoo)
	require.NoError(t, err)

	_, err = service.Put(
		context.Background(),
		&api.PutRequest{
			Key:   &key,
			Value: any,
		},
	)
	require.NoError(t, err)

	response, err := service.Get(
		context.Background(),
		&api.GetRequest{Key: &key},
	)
	require.NoError(t, err)

	foo := &testing_model.Foo{}
	require.NoError(t, ptypes.UnmarshalAny(response.Value, foo))
	assert.Equal(t, originalFoo.Bar, foo.Bar)
}

func TestStorageService_Get_NotFound(t *testing.T) {
	service := New()

	key := api.Key{Value: "qux"}

	_, err := service.Get(
		context.Background(),
		&api.GetRequest{Key: &key},
	)
	require.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))
}
