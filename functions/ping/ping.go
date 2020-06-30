package ping

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/kelseyhightower/envconfig"
	"github.com/mkuchenbecker/storage/api"
	testing_model "github.com/mkuchenbecker/storage/testing/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// var connection grpc.ClientConnInterface
// var client api.StorageClient

// Settings is the setting for the storage service.
type Settings struct {
	DatastoreService string `envconfig:"DATASTORE_SERVICE" default:"storage-bvdpb3qesq-uc.a.run.app:443"`
}

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func getSettings(prefix string) *Settings {
	var s Settings
	err := envconfig.Process(prefix, &s)
	if err != nil {
		log.Fatal(context.Background(), err.Error())
	}
	return &s
}

func init() {
	fmt.Printf("Pinging Datastore")
	flag.Parse()
}

// HelloPubSub consumes a Pub/Sub message.
func PingDatastore(ctx context.Context, m PubSubMessage) error {
	defer glog.Flush()
	glog.Infof("Pinging Datastore")
	fmt.Printf("Pinging Datastore")

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	settings := getSettings("")
	connection, err := grpc.Dial(
		settings.DatastoreService,
		opts...,
	)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	glog.Infof("Connection Set Up")

	client := api.NewStorageClient(connection)
	glog.Infof("Client Set Up")

	key := api.Key{Value: time.Now().String()}
	originalFoo := &testing_model.Foo{Bar: "baz"}
	any, err := ptypes.MarshalAny(originalFoo)
	if err != nil {
		glog.Errorf("encountered an error: %s", err.Error())
		return err
	}

	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	_, err = client.Put(
		ctx,
		&api.PutRequest{
			Key:   &key,
			Value: any,
		},
	)
	if err != nil {
		glog.Errorf("encountered an error: %s", err.Error())
		return err
	}

	response, err := client.Get(
		context.Background(),
		&api.GetRequest{Key: &key},
	)
	if err != nil {
		glog.Errorf("encountered an error: %s", err.Error())
		return err
	}

	foo := &testing_model.Foo{}
	if err != nil {
		glog.Errorf("encountered an error: %s", err.Error())
		return err
	}

	err = ptypes.UnmarshalAny(response.Value, foo)
	if err != nil {
		glog.Errorf("encountered an error: %s", err.Error())
		return err
	}
	glog.Infof("Sent: %v\nRecieved %v\n", originalFoo, foo)
	return nil
}
