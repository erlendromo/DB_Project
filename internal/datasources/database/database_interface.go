package database

type Database interface {
	CreateDatabase()
	Start()
	CreateTables()
	Close()
}
