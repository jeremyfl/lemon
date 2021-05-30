package internal

import (
	"github.com/naoina/genmai"
	"log"
)

// Database is the high level function contains basic of ORM API
// Following D of SOLID principles
type Database interface {
	Select(output interface{}, args ...interface{}) (err error)
	Insert(obj interface{}) (affected int64, err error)

	CreateTable(table interface{}) error
	Close() error
}

// DatabaseMemImpl implement Database interface
// is the low level - detail class
type DatabaseMemImpl struct {
	Client *genmai.DB
}

// NewMemDatabaseImpl creating the instance of database
func NewMemDatabaseImpl() (Database, error) {
	clientInstance, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		log.Println("Error when creating new database instances from library")

		return nil, err
	}

	return &DatabaseMemImpl{Client: clientInstance}, err
}

func (dbm DatabaseMemImpl) Select(output interface{}, args ...interface{}) (err error) {
	return dbm.Client.Select(output)
}

func (dbm DatabaseMemImpl) Insert(obj interface{}) (affected int64, err error) {
	log.Println("Inserting the data ", obj)

	return dbm.Client.Insert(obj)
}

func (dbm DatabaseMemImpl) CreateTable(table interface{}) error {
	return dbm.Client.CreateTable(table)
}

func (dbm DatabaseMemImpl) Close() error {
	log.Println("Closing the db")

	return dbm.Client.Close()
}
