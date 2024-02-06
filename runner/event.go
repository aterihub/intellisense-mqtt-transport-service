package runner

import (
	"context"
	"fmt"
	"insellisense-mqtt-transport-service/service"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type RunnerEvent struct {
	totalWorker int
	natsC       *service.Nats
	mqttC       *service.Mqtt
}

func CreateEvent(nats *service.Nats, mqtt *service.Mqtt, totalWorker int) *RunnerEvent {
	return &RunnerEvent{
		totalWorker: totalWorker,
		natsC:       nats,
		mqttC:       mqtt,
	}
}

func (r *RunnerEvent) ListenMessage(ctx context.Context) {
	r.mqttC.Mqtt.Subscribe("AI/#", 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Worker received message on topic %s: %s\n", msg.Topic(), msg.Payload())

		// Forward the message to NATS
		topic := strings.Replace(msg.Topic(), "/", ".", -1)
		err := r.natsC.Nc.Publish(topic, msg.Payload())
		if err != nil {
			fmt.Println("Error forwarding message to NATS:", err)
		}
	})
}
