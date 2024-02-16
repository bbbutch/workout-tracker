package app

import (
	"io/fs"
	"log/slog"

	"github.com/alexedwards/scs/v2"
	"github.com/jovandeginste/workouts/pkg/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	echo           *echo.Echo
	log            *slog.Logger
	db             *gorm.DB
	Assets         fs.FS
	sessionManager *scs.SessionManager
	jwtSecret      []byte
}

func (a *App) Connect() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	a.db = db

	return nil
}

func NewApp(l *slog.Logger) *App {
	return &App{
		log:       l,
		jwtSecret: []byte("secret"), // TODO: change to configuration
	}
}
