package database

import (
	"database/sql"
	"errors"
	"fmt"
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

	// The actual database connection
	database *sql.DB

	// If we failed to open a connection to the database then this is why
	openError error
}

type DatabaseManager struct {
	aliases map[string]*aliasDetails
	lock    sync.Mutex
}

// Creates a new database manager
func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{
		aliases: make(map[string]*aliasDetails),
	}
}

// Registers an alias with the database manager.
//
// An application should register all connections during startup and then
// reference this manager via the DatabaseProvider interface
func (self *DatabaseManager) Register(alias string, connectionInfo ConnectionInfo) error {
	if len(alias) == 0 {
		return errors.New("alias is empty")
	}

	if len(connectionInfo.DriverName) == 0 {
		return fmt.Errorf("driver name is empty: %s", alias)
	}

	if len(connectionInfo.ConnectionString) == 0 {
		return fmt.Errorf("connection string is empty: %s", alias)
	}

	self.lock.Lock()
	defer self.lock.Unlock()

	if _, found := self.aliases[alias]; found {
		return fmt.Errorf("alias already exists: %s", alias)
	}

	details := &aliasDetails{
		connectionInfo: connectionInfo,
	}

	self.aliases[alias] = details

	return nil
}

// Attempts to open the database with the specified alias.
// Callers should not close the database, they can hang onto it as long as they like.
// The connections are initialized on demand, and the actual connection will be opened
// the first time someone calls this method with a specific alias
func (self *DatabaseManager) Open(alias string) (*sql.DB, error) {
	details, found := self.grabDetails(alias)
	if !found {
		return nil, fmt.Errorf("alias does not exist: %s", alias)
	}

	details.initializer.Do(func() {
		details.database, details.openError = sql.Open(details.connectionInfo.DriverName, details.connectionInfo.ConnectionString)
	})

	return details.database, details.openError
}

func (self *DatabaseManager) grabDetails(alias string) (*aliasDetails, bool) {
	self.lock.Lock()
	defer self.lock.Unlock()

	details, found := self.aliases[alias]
	return details, found
}
