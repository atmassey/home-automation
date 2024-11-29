package main

import (
	"encoding/json"
	"fmt"
	"home-automation/sensors"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var door sensors.Door
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	if msg.Topic() == "zigbee2mqtt/officedoor" {
		err := json.Unmarshal(msg.Payload(), &door)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Door: %v\n", door)
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	env, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	var broker = env["MQTT_BROKER"]
	port, err := strconv.Atoi(env["MQTT_PORT"])
	if err != nil {
		fmt.Println("Error loading port from .env file")
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(env["MQTT_CLIENT_ID"])
	opts.SetUsername(env["MQTT_USERNAME"])
	opts.SetPassword(env["MQTT_PASSWORD"])
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	select {}
}

func sub(client mqtt.Client) {
	topic := "zigbee2mqtt/#"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
