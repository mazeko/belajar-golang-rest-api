package simple

type Database struct {
	Name string
}

type DatabaseMongoDB Database
type DatabasePostgre Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostgreSQL() *DatabasePostgre {
	return (*DatabasePostgre)(&Database{Name: "PostgreSQL"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgre
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgre *DatabasePostgre, mongodb *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgre, DatabaseMongoDB: mongodb}
}
