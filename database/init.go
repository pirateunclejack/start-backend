package database

import (
	"fmt"
	"log"
	"start-backend/helper"

	"github.com/jmoiron/sqlx"
)

var db_host        string
var db_port        string
var db_name        string
var db_user        string
var db_password    string
var redis_addr     string
var redis_port     string
var redis_password string
var redis_database int

func init() {
    vip := helper.GetConfig()
    db_host = vip.GetString("DB_HOST")
    db_port = vip.GetString("DB_PORT")
    db_name = vip.GetString("DB_NAME")
    db_user = vip.GetString("DB_USER")
    db_password = vip.GetString("DB_PASSWORD")
    pg_db := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        db_user, db_password, db_host, db_port, db_name)
    db, err := sqlx.Connect("postgres", pg_db)
    if err != nil {
        log.Fatalf("failed to connect database: %v\n", err)
    }
    db.MustExec(init_schema)
    db.Close()

    redis_addr = vip.GetString("REDIS_ADDR")
    redis_port = vip.GetString("REDIS_PORT")
    redis_password = vip.GetString("REDIS_PASSWORD")
    redis_database = vip.GetInt("REDIS_DATABASE")
}
