package main

import (
	"clearcode/pkg/store/postgres"
	"context"
	"fmt"
)

func main() {
	db := postgres.Postgres{Host: "localhost", Port: 5432, User: "clear", Password: "code", Dbname: "clearcode"}
	conn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	fmt.Printf("DB Connected!")

}
