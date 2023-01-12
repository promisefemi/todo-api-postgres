package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todo/data"

	"github.com/jackc/pgx/v5"
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
	if err != nil {
		errorResponse := data.NewErrorResponse(err, http.StatusInternalServerError, "Internal server error")
		data.Respond(rw, errorResponse)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo data.Todo
		err = rows.Scan(&todo.ID, &todo.Text, &todo.Description, &todo.Completed, &todo.CompletedAt, &todo.CreatedAt)
		if err != nil {
			errorResponse := data.NewErrorResponse(err, http.StatusInternalServerError, "Internal server error")
			data.Respond(rw, errorResponse)
			return
		}
		todos = append(todos, todo)
	}
	if rows.Err() != nil {
		errorResponse := data.NewErrorResponse(err, http.StatusInternalServerError, "Internal server error")
		data.Respond(rw, errorResponse)
		return
	}

	todoPayload := &data.Response{
		Data: todos,
	}
	data.Respond(rw, todoPayload)
}

func (t *TodoHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var todo data.Todo

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		errResponse := data.NewErrorResponse(err, http.StatusInternalServerError, "Please check request body")
		data.Respond(rw, errResponse)
		return
	}
	if _, err := todo.Validate(); err != nil {
		errResponse := data.NewErrorResponse(err, http.StatusBadRequest, err.Error())
		data.Respond(rw, errResponse)
		return
	}

	_, err = t.DBConn.Exec(context.Background(), "INSERT INTO todos (text, description) VALUES ($1,$2)", todo.Text, todo.Description)
	if err != nil {
		errResponse := data.NewErrorResponse(err, http.StatusInternalServerError, "internal server error")
		data.Respond(rw, errResponse)
		return
	}

	todoResponse := &data.Response{
		Data: todo,
	}
	data.Respond(rw, todoResponse)
}

func (t *TodoHandler) MarkAsCompleted(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		errorResponse := data.NewErrorResponse(fmt.Errorf("no ID"), http.StatusBadRequest, "ID not found in query parameter")
		data.Respond(rw, errorResponse)
		return
	}

	var rowID int

	err := t.DBConn.QueryRow(context.Background(), "SELECT id from todos where id = $1", id).Scan(&rowID)
	if err != nil {
		errorResponse := data.NewErrorResponse(err, http.StatusBadRequest, "Cannot find a todo with that id")
		data.Respond(rw, errorResponse)
		return
	}

	_, err = t.DBConn.Exec(context.Background(), "UPDATE todos set completed = $1, completedat = $2 where id = $3", true, time.Now(), id)
	if err != nil {
		errorResponse := data.NewErrorResponse(err, http.StatusInternalServerError, "Unable to complete todo, please contact support")
		data.Respond(rw, errorResponse)
		return
	}

	var todo data.Todo
	err = t.DBConn.QueryRow(context.Background(), "SELECT id, text, description, completed, completedat, createdat from todos where id = $1", id).Scan(&todo.ID, &todo.Text, &todo.Description, &todo.Completed, &todo.CompletedAt, &todo.CreatedAt)
	if err != nil {
		log.Println(err)
		errorResponse := data.NewErrorResponse(err, http.StatusBadRequest, "Something went wrong, please contact support")
		data.Respond(rw, errorResponse)
		return
	}

	todoResponse := &data.Response{
		Message: "Todo completed",
		Data:    todo,
	}
	data.Respond(rw, todoResponse)
}
