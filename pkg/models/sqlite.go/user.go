package sqlite3

import "log"

//CreateUserTable метод для создания таблицы в базе данных
func (db *Database)CreateUserTable(){

	stmt, err := db.SQLDb.Prepare(`CREATE TABLE IF NOT EXISTS "user" (
		"user_id" INTEGER NOT NULL,
		"first_name" VARCHAR(320) NOT NULL, 
		"last_name"	VARCHAR(320) NOT NULL,
		PRIMARY KEY("user_id" AUTOINCREMENT) 
	);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
}
// AddUser функция для добавления юзера в базу данных
func(db *Database)AddUser(){

}