package db

import (
	"errors"
	"sync"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/erfanshekari/url-shortener/config"
)

var lock = &sync.Mutex{}

var dbInstance *BadgerDB

func GetInstance() *BadgerDB {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			conf := config.GetConfigInstance()
			db := BadgerDB{Path: conf.DataDir}
			dbInstance = &db
		}
	}
	return dbInstance
}

type BadgerDB struct {
	Path string
	DB   *badger.DB
}

func (db *BadgerDB) Start() error {
	db_, err := badger.Open(badger.DefaultOptions(db.Path))
	if err != nil {
		return err
	} else {
		db.DB = db_
	}
	return nil
}

func (db *BadgerDB) Set(key []byte, val []byte) error {
	if db.DB == nil {
		return errors.New("BadgerDB is not initialized")
	}
	return db.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}

func (db *BadgerDB) Get(key []byte) (*[]byte, error) {
	if db.DB == nil {
		return nil, errors.New("BadgerDB is not initialized")
	}
	var value *[]byte
	err := db.DB.View(func(txn *badger.Txn) error {
		val, err := txn.Get(key)
		if err != nil {
			return err
		}
		err = val.Value(func(val []byte) error {
			value = &val
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (db *BadgerDB) Close() error {
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}
