package firestore

import (
	"context"
	"fmt"
	"pkg/errors"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/mkuchenbecker/storage/service"
	"github.com/mkuchenbecker/storage/service/datamodel"
	"github.com/mkuchenbecker/storage/service/status"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type firestoreClient struct {
	client     *firestore.Client
	collection string
}

func NewFirestoreClient(client *firestore.Client, collection string) service.DataBackend {
	return &firestoreClient{client: client, collection: collection}
}

func (c *firestoreClient) Save(ctx context.Context, item *datamodel.Item) error {
	_, err := c.client.Collection(c.collection).Doc(item.Key).Set(ctx, map[string]interface{}{"key": item.Key, "value": item.Value})
	return err
}

func (c *firestoreClient) Get(ctx context.Context, key string) (*datamodel.Item, error) {
	iter := c.client.Collection(c.collection).
		Where("key", "==", key).
		Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		keyIface, ok := doc.Data()["key"]
		if !ok {
			return nil, errors.New("malformated data (key)")
		}
		key, ok := keyIface.(string)
		if !ok {
			return nil, errors.New(fmt.Sprintf("key wrong type: %+v", keyIface))
		}

		valueIface, ok := doc.Data()["value"]
		if !ok {
			return nil, errors.New("malformated data (value)")
		}
		value, ok := valueIface.(*any.Any)
		if !ok {
			return nil, errors.New(fmt.Sprintf("value wrong type: %+v", valueIface))
		}
		return &datamodel.Item{Key: key, Value: value}, nil
	}
	return nil, status.ErrNotFound
}
