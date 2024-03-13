package service

import (
	"context"
	"fmt"
	"insellisense-mqtt-transport-service/config"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Mqtt struct {
	Mqtt mqtt.Client
}

func NewMqttService(ctx context.Context, conf config.Mqtt) (*Mqtt, error) {
	mqttService := new(Mqtt)

	err := mqttService.connectToMqttServer(ctx, conf.Url, conf.Username, conf.Password)
	if err != nil {
		return nil, err
	}

	return mqttService, nil
}

func (m *Mqtt) connectToMqttServer(ctx context.Context, url string, user string, pass string) error {
	opts := mqtt.NewClientOptions().AddBroker(url).SetUsername(user).SetPassword(pass)
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("MQTT Service Connection lost: %v\n", err)
	})
	opts.SetReconnectingHandler(func(c mqtt.Client, co *mqtt.ClientOptions) {
		log.Printf("MQTT Service Reconnecting...")
	})
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("error when connect to MQTT Broker, url: %s, token: %s", url, token.Error())
	}

	m.Mqtt = client
	return nil
}
