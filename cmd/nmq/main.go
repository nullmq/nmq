package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/nullmq/nmq"
)

func main() {
	fmt.Println("Welcome to NullMQ")
	cfg := nmq.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
}
