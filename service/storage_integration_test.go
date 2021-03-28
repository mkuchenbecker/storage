//+build integration

package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mkuchenbecker/storage/api"
	testing_model "github.com/mkuchenbecker/storage/testing/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

const port = 50060

func TestService(t *testing.T) {
	service := New(NewMapBackend())
	go func() {
		err := StartService(service, port)
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
	any, err := anypb.New(originalFoo)
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

	fooAny, err := anypb.New(response.Value)
	require.NoError(t, err)
	foo := new(testing_model.Foo)
	err = fooAny.UnmarshalTo(foo)
	require.NoError(t, err)
	assert.Equal(t, originalFoo.Bar, foo.Bar)
}

func TestService_Listen_Failure(t *testing.T) {
	err := StartService(New(NewMapBackend()), -1)
	require.Error(t, err)
	assert.Equal(t,
		"listen tcp: address -1: invalid port",
		errors.Cause(err).Error(),
	)
}
