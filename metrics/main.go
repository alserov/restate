package main

import (
	"github.com/alserov/restate/metrics/internal/app"
	"github.com/alserov/restate/metrics/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
