package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"html/template"

	"chatapp.new.net/internal/database"
	"chatapp.new.net/internal/models"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

type application struct{
	logger *slog.Logger
	users models.UserModelInterface
	rooms models.RoomModelInterface
	messages models.MessageModelInterface
	templateCache map[string]*template.Template

}
func main(){
	addr := flag.String("addr", ":4040", "HTTP network address")
	flag.Parse()
	logger:= slog.New(slog.NewTextHandler(os.Stdout, nil))
	dsn := "chatapp.db"
	db,err := openDB(dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}	
	defer db.Close()

	templateCache, err := newTemplateCache()
     if err != nil {
         logger.Error(err.Error())
         os.Exit(1)
     }
	app :=&application{
		logger: logger,
		users: &models.UserModel{DB: db},
		rooms: &models.RoomModel{DB: db},
		messages: &models.MessageModel{DB: db},
		templateCache: templateCache,
	}
	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
}
func runMigrations(db *sql.DB) error {
	goose.SetBaseFS(database.MigrationsFS)

	if err:= goose.SetDialect("sqlite"); err != nil {
		return err
	}
	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}
	return nil
}
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	err = runMigrations(db)
	if err != nil {
		db.Close()
		return nil, err
	}
	db.SetMaxOpenConns(1)
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil

}