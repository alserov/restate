package main

import (
	"github.com/alserov/restate/meetings/internal/app"
	"github.com/alserov/restate/meetings/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
