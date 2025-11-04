package database

// PostgresConnection is a concrete Connection product for Postgres
type PostgresConnection struct{}

func (PostgresConnection) Connect() string {
	return "Postgres connection established"
}

// PostgresQueryBuilder is a concrete QueryBuilder product for Postgres
type PostgresQueryBuilder struct{}

func (PostgresQueryBuilder) Build() string {
	return "Postgres query builder created"
}

// PostgresFactory is a concrete factory for Postgres products
type PostgresFactory struct{}

func (PostgresFactory) CreateConnection() Connection {
	return PostgresConnection{}
}

func (PostgresFactory) CreateQueryBuilder() QueryBuilder {
	return PostgresQueryBuilder{}
}
