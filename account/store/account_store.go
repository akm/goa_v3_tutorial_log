package accountstore

import (
	"context"
	"log"

	"go.mercari.io/datastore"
	"go.mercari.io/datastore/boom"

	"calcsvc/appbase/ds"

	"calcsvc/account/model"
)

type AccountStore struct {
	bm     *boom.Boom
	logger *log.Logger
}

func NewAccountStore(logger *log.Logger) *AccountStore {
	return &AccountStore{logger: logger}
}

func (s *AccountStore) Put(ctx context.Context, m *accountmodel.Account) (res datastore.Key, err error) {
	err = ds.WithBoom(ctx, func(bm *boom.Boom) error {
		key, err := bm.Put(m)
		if err != nil {
			s.logger.Printf("Failed to Put %v because of [%T] %v\n", m, err)
			return err
		}
		res = key
		return nil
	})
	return
}
