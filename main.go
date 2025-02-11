package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Password  string
	Role      string
	IsVerify  bool
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=magic sslmode=disable password=Rick8917")
	if err != nil {
		log.Fatal("can`t open db!")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("can`t pong db!")
	}

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	user, err := getUserById(db, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
	fmt.Println(user)
}

func getUsers(db *sql.DB) ([]User, error) {

	rows, err := db.Query("select * from public.\"User\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Email, &u.Password, &u.Role, &u.IsVerify)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func getUserById(db *sql.DB, id int64) (User, error) {
	var u User
	err := db.QueryRow("select * from public.\"User\" where id = $1", id).
		Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Email, &u.Password, &u.Role, &u.IsVerify)
	
	if err != nil {
		return User{}, err
	}

	return u, nil
}
