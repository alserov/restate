package main

import (
	"github.com/alserov/restate/estate/internal/app"
	"github.com/alserov/restate/estate/internal/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.MustStart(config.MustLoad())
}
