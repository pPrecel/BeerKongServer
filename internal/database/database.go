package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/vrischmann/envconfig"
	"log"
)

type DataConfig struct {
	Driver string `envconfig:"default=mysql"`
	Port string `envconfig:"default=3306"`
	DatabaseName string `envconfig:"default=database"`
	ConnectionProtocole string `envconfig:"default=tcp"`
	Username string `envconfig:"default=root"`
	Password string `envconfig:"default=root"`
	Address string `envconfig:"default=localhost"`
}

const dataSourceNameFormat = "%s:%s@%s(%s:%s)/%s"

type database struct {
	config DataConfig
}

type Database interface {
	Open()
}

func New(config DataConfig) Database {
	return &database{
		config: config,
	}
}

func (s *database) Open() {

	connectionString := fmt.Sprintf(
		dataSourceNameFormat,
		s.config.Username,
		s.config.Password,
		s.config.ConnectionProtocole,
		s.config.Address,
		s.config.Port,
		s.config.DatabaseName,
	)

	db, err := sql.Open(s.config.Driver, connectionString)
	if err != nil {
		log.Fatal("no baza padła kurła")
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Noooo ping nie poszedł kurła: %s", err.Error())
	}

	log.Print("Gites wszystko byyyku")
}