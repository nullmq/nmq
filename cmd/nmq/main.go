package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/nullmq/nmq"
	"github.com/nullmq/nmq/internal/server"
)

func main() {
	fmt.Println("Welcome to NullMQ")
	cfg := nmq.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
	addr := fmt.Sprintf(":%d", cfg.Port)
	srv := server.NewHTTPServer(addr)
	log.Fatal(srv.ListenAndServe())
}
