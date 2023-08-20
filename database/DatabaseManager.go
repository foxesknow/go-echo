package database

import (
	"database/sql"
	"errors"
	"sync"
)

// Information on what we're going to connect to
type ConnectionInfo struct {
	DriverName       string
	ConnectionString string
}

type aliasDetails struct {
	connectionInfo ConnectionInfo
	initializer    sync.Once

	database  *sql.DB
	openError error
}

type DatabaseManager struct {
	aliases map[string]*aliasDetails
}

func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{
		aliases: make(map[string]*aliasDetails),
	}
}

func (self *DatabaseManager) Register(alias string, connectionInfo ConnectionInfo) error {
	if _, found := self.aliases[alias]; found {
		return errors.New("alias already exists")
	}

	details := &aliasDetails{
		connectionInfo: connectionInfo,
	}

	self.aliases[alias] = details

	return nil
}

func (self *DatabaseManager) Open(alias string) (*sql.DB, error) {
	details, found := self.aliases[alias]
	if !found {
		return nil, errors.New("alias does not exist")
	}

	details.initializer.Do(func() {
		details.database, details.openError = sql.Open(details.connectionInfo.DriverName, details.connectionInfo.ConnectionString)
	})

	return details.database, details.openError
}
