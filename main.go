package main

import (
	"context"
	"insellisense-mqtt-transport-service/runner"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
)

func main() {
	ctx := context.TODO()

	app := &cli.App{
		Name:  "Intellisense MQTT Transport Service",
		Usage: "Receive data from MQTT and forward to NATS",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "Load configuration from `FILE`",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			runner.Run(ctx, cCtx.Value("config").(string))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("Shutting down...")
}
