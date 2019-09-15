package ds

import (
	"context"

	"cloud.google.com/go/datastore"
)

func NewGoogleClient(ctx context.Context) (*datastore.Client, error) {
	// https://cloud.google.com/datastore/docs/tools/datastore-emulator#setting_environment_variables
	// https://godoc.org/cloud.google.com/go/datastore#NewClient
	return datastore.NewClient(ctx, "")
}

func WithGoogleClient(ctx context.Context, f func(*datastore.Client) error) error {
	c, err := NewGoogleClient(ctx)
	if err != nil {
		return err
	}
	return f(c)
}
