package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"todo/data"
)

type TodoHandler struct {
	DBConn *pgx.Conn
}

func (t *TodoHandler) ListAll(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	todos := make([]data.Todo, 0)
	rows, err := t.DBConn.Query(context.Background(), "SELECT id, text, description, completed, completedat, createdat from todos")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		fmt.Println("Count")
		var todo data.Todo
		err = rows.Scan(&todo.ID, &todo.Text, &todo.Description, &todo.Completed, &todo.CompletedAt, &todo.CreatedAt)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
		}
		todos = append(todos, todo)

	}
	if rows.Err() != nil {
		fmt.Println(err)
	}
	byteTodos, err := json.Marshal(todos)
	if err != nil {
		fmt.Printf("error marshalling json: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	_, _ = rw.Write(byteTodos)
}

func (t *TodoHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var todo data.Todo

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte("Please check request body"))
	}
	if _, err := todo.Validate(); err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte("Please check request body"))
	}

	_, err = t.DBConn.Exec(context.Background(), "INSERT INTO todos (text, description) VALUES ($1,$2)", todo.Text, todo.Description)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	byteTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Printf("error marshalling json: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	_, _ = rw.Write(byteTodo)
	return
}

func (t *TodoHandler) MarkAsCompleted(rw http.ResponseWriter, r *http.Request) {

}
