package main

import (
	"context"
	"fmt"
	"runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) (string, error) {
	filename := fmt.Sprintf("build/%s.txt", runtime.GOOS)
	message, err := myFiles.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(message), nil
}
