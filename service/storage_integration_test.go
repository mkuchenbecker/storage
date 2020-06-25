package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mkuchenbecker/storage/api"
	testing_model "github.com/mkuchenbecker/storage/testing/model"
	"google.golang.org/grpc"
)

type SyncFunc func()

const port = 50060

func TestService(t *testing.T) {
	service := New().(*service)
	go func() {
		err := service.Start(port)
		require.NoError(t, err)
	}()
	time.Sleep(1 * time.Second)

	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%d", port),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	require.NoError(t, err)
	defer conn.Close()
	client := api.NewStorageClient(conn)

	key := api.Key{Value: "qux"}
	originalFoo := &testing_model.Foo{Bar: "baz"}
	any, err := ptypes.MarshalAny(originalFoo)
	require.NoError(t, err)

	_, err = client.Put(
		context.Background(),
		&api.PutRequest{
			Key:   &key,
			Value: any,
		},
	)
	require.NoError(t, err)

	response, err := client.Get(
		context.Background(),
		&api.GetRequest{Key: &key},
	)
	require.NoError(t, err)

	foo := &testing_model.Foo{}
	require.NoError(t, ptypes.UnmarshalAny(response.Value, foo))
	assert.Equal(t, originalFoo.Bar, foo.Bar)
}

func TestService_Listen_Failure(t *testing.T) {
	err := New().(*service).Start(-1)
	require.Error(t, err)
	assert.Equal(t,
		"listen tcp: address -1: invalid port",
		errors.Cause(err).Error(),
	)
}
