package model

type User struct {
    UserId      int64     `db:"user_id"`
    Email       string    `db:"email"`
    Username    string    `db:"username"`
    Password    string    `db:"password"`
}
