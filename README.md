# Todo RestAPI

## Description
Yes i know another todo app, but the purpose of this project is not to create another todo but to combine independent knowledge into use in one coherent project. 

The app is very simple.

### Features include:
- Adding a new todo
- Listing all todos
- Marking a todo as completed



## Technologies in use
- `Language: ` [Golang](https://github.com/golang/go)
- `Database: ` [Postgres](https://github.com/postgres/postgres)
- `Driver: ` [PGX](https://github.com/jackc/pgx)

## Usage
### Prerequisite:
- Golang Installation
- A test Postgres database

Create Datbase table
```sql
    CREATE TABLE IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        text VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        completed boolean NOT NULL DEFAULT FALSE,
        createdat timestamp NOT NULL DEFAULT NOW(),
        completedat time
    );
```

Clone git repo
```bash
git clone git@github.com:promisefemi/todo-api-postgres.git
```

Run program
```bash
go run main.go
```

### Advice and Contributions are apprecited 