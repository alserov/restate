package main

import (
	"github.com/alserov/restate/estate/internal/app"
	"github.com/alserov/restate/estate/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
