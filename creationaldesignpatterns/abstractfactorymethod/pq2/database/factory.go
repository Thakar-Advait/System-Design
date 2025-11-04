package database

import "fmt"

// DatabaseFactory is the abstract factory interface
// It declares methods to create families of related products
type DatabaseFactory interface {
	CreateConnection() Connection
	CreateQueryBuilder() QueryBuilder
}

func NewDBFactory(dbType string) (DatabaseFactory, error) {
	switch dbType {
	case "mysql":
		return MySQLFactory{}, nil
	case "postgres":
		return PostgresFactory{}, nil
	default:
		return nil, fmt.Errorf("unsupported database: %s", dbType)
	}
}
