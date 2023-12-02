package main

import (
	"context"
	"flag"
	"log"
	"uzum_shop/internal/app"
)

func main() {
	flag.Parse()
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Occured error on creating app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Occured error on running app: %v", err)
	}
}
