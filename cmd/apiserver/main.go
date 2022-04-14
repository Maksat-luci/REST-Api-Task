package main

import (
	"amaymon/internal/app/apiserver"
	sqlite3 "amaymon/pkg/models/sqlite.go"
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/mattn/go-sqlite3"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	// проверка на верных данных в файле apiserver.toml
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlite3.ConnectDb("sqlite3", "TSARKA.sqlite")
	fmt.Println("+DATABASE CREATED")
	if err != nil {
		log.Fatal(err)
	}
	db.CreateUserTable()
	defer db.SQLDb.Close()

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
