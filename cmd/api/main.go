package main

import (
	"os"

	"github.com/PeremyshlevPR/AuthService/internal/app"
)

func main() {
	if len(os.Args) <= 1 {
		panic("config path is required")
	}

	configPath := os.Args[1]
	config, err := app.LoadConfig(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	app := app.NewApp(config)
	if err := app.Run(); err != nil {
		panic("failed to run app: " + err.Error())
	}
}
