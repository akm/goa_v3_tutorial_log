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

const AccountKind = "Account"

func (s *AccountStore) Kind() string {
	return AccountKind
}

func (s *AccountStore) Put(ctx context.Context, m *accountmodel.Account) (res datastore.Key, err error) {
	err = ds.WithBoom(ctx, func(bm *boom.Boom) error {
		key, err := bm.Put(m)
		if err != nil {
			s.logger.Printf("Failed to Put %v because of [%T] %v\n", m, err, err)
			return err
		}
		res = key
		return nil
	})
	return
}

func (s *AccountStore) ExistByToken(ctx context.Context, token string) (res bool, err error) {
	err = ds.WithBoom(ctx, func(bm *boom.Boom) error {
		q := bm.NewQuery(s.Kind()).Filter("JwtToken =", token).Limit(1).KeysOnly()
		keys, err := bm.GetAll(q, nil)
		if err != nil {
			s.logger.Printf("Failed to Get keys for %s because of [%T] %v\n", token, err, err)
			return err
		}
		res = len(keys) > 0
		return nil
	})
	return
}
