package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo-api/database"

	_ "modernc.org/sqlite"
)

// App holds application dependencies.
type App struct {
	queries *database.Queries
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type TodoRequest struct {
	Description string `json:"description"`
}

//go:embed schema.sql
var ddl string

func run() error {
	ctx := context.Background()

	// Initialize database
	db, err := sql.Open("sqlite", ":memory:")
	db.SetMaxOpenConns(1)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	// Check database connection
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("connect database: %w", err)
	}

	// Initialize SQLC queries
	queries := database.New(db)

	app := &App{
		queries: queries,
	}

	// Initialize HTTP routes
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", app.listTodos)
	mux.HandleFunc("POST /todos", app.createTodo)
	mux.HandleFunc("GET /$", app.homeHandler)

	addr := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", addr)

	return http.ListenAndServe(addr, mux)
}

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

// Helper: send JSON response
func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// GET /todos
func (app *App) listTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.queries.ListTodos(r.Context())
	if err != nil {
		// http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, todos)
}

// POST /todos
func (app *App) createTodo(w http.ResponseWriter, r *http.Request) {
	var req TodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Description == "" {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}

	todo, err := app.queries.CreateTodo(
		r.Context(),
		req.Description,
	)
	if err != nil {
		log.Printf("CreateTodo error: %v", err)
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusCreated, todo)
}
