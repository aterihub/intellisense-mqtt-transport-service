package runner

import (
	"context"
	"encoding/json"
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
		// Forward the message to NATS
		topic := strings.Replace(msg.Topic(), "/", ".", -1)
		arrTopic := strings.Split(topic, ".")

		switch arrTopic[len(arrTopic)-1] {
		case "heartbeat":
			r.PublishMessage(topic, msg.Payload())
			return
		case "data":
			var message MessageData
			err := json.Unmarshal(msg.Payload(), &message)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}

			messageSent := MessageSent{
				MessageId: message.MessageId,
				LoraRssi:  message.LoraRssi,
				Ts:        message.Ts,
				Data:      make(map[string]int),
			}

			for _, modbusData := range message.Data.ModbusData {
				stringCode, err := checkFieldType(message.Data.SlaveId, modbusData.Address)

				if err == nil {
					messageSent.Data[stringCode] = modbusData.Value
				}
			}

			jsonData, err := json.MarshalIndent(messageSent, "", "  ")
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				return
			}
			r.PublishMessage(topic, jsonData)
			return
		}
	})
}

func (r *RunnerEvent) PublishMessage(context string, message []byte) {
	err := r.natsC.Nc.Publish(context, message)
	if err != nil {
		fmt.Println("Error forwarding message to NATS:", err)
	}
}
