package database

import "database/sql"

// Providers access to a database using an alias which insulates users from the
// underlying specifics of the driver, such as its implementation
//
// Implemenetation must support concurrent calls from goroutines
type DatabaseProvider interface {
	// Attempts to open the database with the specified alias
	Open(alias string) (*sql.DB, error)
}
