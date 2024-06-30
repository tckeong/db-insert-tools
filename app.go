package main

import (
	"context"
	"fmt"

	DBConn "db-insert-app/internal/dbConnection"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	db  *DBConn.DBConnection
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowCenter(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ConnectToDB(host, port, username, password, dbname string, dbType int) {
	// Create a new DB connection
	a.db = DBConn.New(host, port, username, password, dbname, dbType)
}

func (a *App) shutdown(_ context.Context) {
	a.db.Close()
}

func (a *App) WriteToDB() {
	fmt.Println("Writing to db")
}

func (a *App) Test(host, port, username, password, dbname string) {
}
