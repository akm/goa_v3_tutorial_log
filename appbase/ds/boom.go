package ds

import (
	"context"

	"go.mercari.io/datastore/boom"
)

func NewBoom(ctx context.Context) (*boom.Boom, error) {
	mercariClient, err := NewMercariClient(ctx)
	if err != nil {
		return nil, err
	}
	return boom.FromClient(ctx, mercariClient), nil
}

func WithBoom(ctx context.Context, f func(*boom.Boom) error) error {
	bm, err := NewBoom(ctx)
	if err != nil {
		return err
	}
	return f(bm)
}
