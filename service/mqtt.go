package service

import (
	"context"
	"fmt"
	"insellisense-mqtt-transport-service/config"

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
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("error when connect to MQTT Broker, url: %s, token: %s", url, token.Error())
	}

	m.Mqtt = client
	return nil
}
