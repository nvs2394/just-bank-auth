package main

import (
	"github.com/nvs2394/just-bank-auth/app"
	"github.com/nvs2394/just-bank-lib/logger"
)

func main() {
	logger.Info("Starting just bank service")
	app.Start()
}
