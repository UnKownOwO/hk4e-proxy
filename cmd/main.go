package main

import (
	"context"
	"flag"
	"hk4e-proxy/app"
)

var (
	config = flag.String("config", "application.toml", "config file")
)

func main() {
	flag.Parse()
	err := app.Run(context.TODO(), *config)
	if err != nil {
		panic(err)
	}
}
