package src

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const connectionPattern = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

type DataBaseConfiguration struct {
	Type            string `json:"type"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	DBName          string `json:"dbname"`
	SSL             string `json:"ssl"`
	ConnMaxLifetime int    `json:"connMaxLifetime"`
	MaxIdleConns    int    `json:"maxIdleConns"`
	MaxOpenConns    int    `json:"maxOpenConns"`
}


type PostgresConnector struct {
	db      *sql.DB
}

func NewPostgresConnector(gConfig DataBaseConfiguration) PostgresConnector {
	psqlInfo := fmt.Sprintf(connectionPattern, gConfig.Host, gConfig.Port, gConfig.Username,
		gConfig.Password, gConfig.DBName, gConfig.SSL)
	db, err := sql.Open(gConfig.Type, psqlInfo)
	if err != nil {
		panic("fail to open DB conn")
	}
	db.SetConnMaxLifetime(time.Second * time.Duration(gConfig.ConnMaxLifetime))
	db.SetMaxOpenConns(gConfig.MaxOpenConns)
	db.SetMaxIdleConns(gConfig.MaxIdleConns)

	log.Printf("DBStats:%#v", db.Stats())

	return PostgresConnector{
		db:      db,
	}
}

func (pc PostgresConnector) DBConnection() *sql.DB {
	return pc.db
}