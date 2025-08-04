package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bdreece/ephemera"
)

var cfg = ephemera.DefaultConfig

func init() {
	cfg.UnmarshalEnvVars()
	cfg.UnmarshalFlags(flag.CommandLine)
}

func main() {
	defer exit()

	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := ephemera.New(ephemera.WithConfig(&cfg)).Run(ctx); err != nil {
		panic(err)
	}
}

func exit() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "unexpected panic occurred: %v\n", r)
		os.Exit(1)
	}
}
