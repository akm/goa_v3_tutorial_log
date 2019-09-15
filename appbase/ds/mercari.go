package ds

import (
	"context"

	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

func NewMercariClient(ctx context.Context) (datastore.Client, error) {
	googleClient, err := NewGoogleClient(ctx)
	if err != nil {
		return nil, err
	}
	return clouddatastore.FromClient(ctx, googleClient)
}

func WithMercariClient(ctx context.Context, f func(datastore.Client) error) error {
	c, err := NewMercariClient(ctx)
	if err != nil {
		return err
	}
	return f(c)

}
