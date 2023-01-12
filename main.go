package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"todo/handler"

	"github.com/jackc/pgx/v5"
)

func main() {

	conn, err := pgx.Connect(context.Background(), "postgres://nefertem:root@localhost:5432/todotest?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	todoBase := &handler.TodoHandler{
		DBConn: conn,
	}

	http.HandleFunc("/todos", todoBase.ListAll)
	http.HandleFunc("/todos/create", todoBase.Create)
	http.HandleFunc("/todos/complete/", todoBase.MarkAsCompleted)

	fmt.Println("Server is running at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
