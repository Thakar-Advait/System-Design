package database

// MySQLConnection is a concrete Connection product for MySQL
type MySQLConnection struct{}

func (MySQLConnection) Connect() string {
	return "MySQL connection established"
}

// MySQLQueryBuilder is a concrete QueryBuilder product for MySQL
type MySQLQueryBuilder struct{}

func (MySQLQueryBuilder) Build() string {
	return "MySQL query builder created"
}

// MySQLFactory is a concrete factory for MySQL products
type MySQLFactory struct{}

func (MySQLFactory) CreateConnection() Connection {
	return MySQLConnection{}
}

func (MySQLFactory) CreateQueryBuilder() QueryBuilder {
	return MySQLQueryBuilder{}
}
