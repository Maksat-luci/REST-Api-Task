package sqlite3

import "database/sql"
// Database структура 
type Database struct {
	SQLDb *sql.DB
}
// ConnectDb для открытия базы данных и передачи обьекта для дальнейшего использования
func ConnectDb(driverName string, SQLDbName string) (*Database, error) {
	SQLDb, err := sql.Open(driverName, SQLDbName)
	if err != nil {
		return nil, err
	}
	if err = SQLDb.Ping(); err != nil {
		return nil, err
	}
	return &Database{SQLDb}, nil
}
