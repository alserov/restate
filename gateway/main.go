package main

import (
	"github.com/alserov/restate/gateway/internal/app"
	"github.com/alserov/restate/gateway/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
