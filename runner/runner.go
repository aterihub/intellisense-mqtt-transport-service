package runner

import (
	"context"
	"fmt"
	"insellisense-mqtt-transport-service/config"
	"insellisense-mqtt-transport-service/service"
	"log"
)

const LOGO = `
    _  _____ _____ ____  ___
   / \|_   _| ____|  _ \|_ _|
  / _ \ | | |  _| | |_) || |
 / ___ \| | | |___|  _ < | |
/_/   \_\_| |_____|_| \_\___|
`

const SERVICENAME = "Intellisense MQTT Transport Service"
const VERSION = "v0.1.0"

func Run(ctx context.Context, configPath string) {
	fmt.Print(LOGO + "\n" + SERVICENAME + " " + VERSION + "\n\n")

	conf, err := config.LoadFromFile(configPath)
	if err != nil {
		panic(err)
	}

	log.Printf("Setup NATS Service")
	nats, err := service.NewNatsService(ctx, conf.Nats)
	if err != nil {
		panic(err)
	}

	log.Printf("Setup MQTT Service")
	mqtt, err := service.NewMqttService(ctx, conf.Mqtt)
	if err != nil {
		panic(err)
	}

	log.Printf("Listen mqtt message...")
	event := CreateEvent(nats, mqtt, conf.NumOfWorker)
	event.ListenMessage(ctx)
}
