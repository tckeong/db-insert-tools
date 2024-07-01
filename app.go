package main

import (
	"context"
	"fmt"

	DBConn "db-insert-app/internal/dbConnection"
	models "db-insert-app/internal/models"

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

func (a *App) shutdown(_ context.Context) {
	a.db.Close()
}

func (a *App) ConnectToDB(host, port, username, password, dbname, tableName string, dbType int) {
	// Create a new DB connection
	a.db = DBConn.New(host, port, username, password, dbname, tableName, dbType)
}

func (a *App) WriteToDB(data []models.Pair) error {
	err := a.db.Write(data)

	if err == nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Insert Data",
			Message: "Insert Data Success!",
		})
	} else {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Insert Data",
			Message: "Insert Data Failed!",
		})
	}

	return err
}

func (a *App) Test(host, port, username, password, dbname string) {
}
