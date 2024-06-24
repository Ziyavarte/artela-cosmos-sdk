package iavl

import (
	dbm "github.com/cosmos/cosmos-db"
	iavldbm "github.com/cosmos/iavl/db"
)

type iavlStoreWrapper struct {
	dbm.DB
}

var _ iavldbm.DB = (*iavlStoreWrapper)(nil)

func (s *iavlStoreWrapper) Iterator(start, end []byte) (iavldbm.Iterator, error) {
	return s.DB.Iterator(start, end)
}

func (s *iavlStoreWrapper) ReverseIterator(start, end []byte) (iavldbm.Iterator, error) {
	return s.DB.ReverseIterator(start, end)
}

func (s *iavlStoreWrapper) NewBatch() iavldbm.Batch {
	return s.DB.NewBatch()
}

func (s *iavlStoreWrapper) NewBatchWithSize(size int) iavldbm.Batch {
	return s.DB.NewBatchWithSize(size)
}

func WrapIAVLDB(db dbm.DB) iavldbm.DB {
	return &iavlStoreWrapper{db}
}
