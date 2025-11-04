package database

// Connection represents a database connection product
type Connection interface {
	Connect() string
}

// QueryBuilder represents a query builder product
type QueryBuilder interface {
	Build() string
}

