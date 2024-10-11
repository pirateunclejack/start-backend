package database

import (
	"fmt"
	"log"
	"start-backend/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


var init_schema = `
CREATE TABLE IF NOT EXISTS users (
    user_id  SERIAL PRIMARY KEY,
    email    TEXT   UNIQUE NOT NULL,
    username TEXT   UNIQUE NOT NULL,
    password TEXT          NOT NULL
);
`

func GetDb() *sqlx.DB {
    pg_db := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        db_user, db_password, db_host, db_port, db_name)
    db, err := sqlx.Connect("postgres", pg_db)
    if err != nil {
        log.Fatalf("failed to connect database: %v\n", err)
    }
    return db
}

func CreateUser(db *sqlx.DB, user model.User) error {
    tx := db.MustBegin()
    res, err := tx.NamedExec(
        "INSERT INTO users (email, username, password) VALUES (:email, :username, :password)", &user)
    if err != nil {
        log.Printf("failed to name user %v creation exec: %v\n", user.Username, err)
        return err
    }
    log.Printf("create user named exec: %v", res)
    err = tx.Commit()
    if err != nil {
        log.Printf("failed to create user %v in database: %v", user.Username, err)
        return err
    }
    return nil
}

func GetUser(db *sqlx.DB, username string) *model.User {
    var user model.User

    err := db.Get(&user, "SELECT * FROM users WHERE username=$1", username)
    if err != nil {
        log.Printf("failed to get user %s, from database with error: %v\n", user.Username, err)
    }
    return &user
}
