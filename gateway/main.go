package main

import (
	"github.com/alserov/restate/gateway/internal/app"
	"github.com/alserov/restate/gateway/internal/config"
)

// @title Real Estate
// @version 1.0
// @description Real Estate app gateway

// @BasePath /v1

func main() {
	app.MustStart(config.MustLoad())
}
