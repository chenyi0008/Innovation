package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

const (
	username = "caibin"
	password = "caibin@123"
	port     = 1883
	broker   = "106.52.223.188"
	clientId = "go_mqtt_client"
)

var (
	strMap = make(map[string]string)
)

func MqttMain() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	var topicList []string

	topicList = append(topicList, "topic/test")
	topicList = append(topicList, "topic1")
	topicList = append(topicList, "topic2")
	topicList = append(topicList, "topic3")

	for _, topic := range topicList {
		go listenSub(client, topic)
	}

	for _, topic := range topicList {
		messagePubHandlerInit(client, topic)
	}

}

// 全局MQTT pub消息处理
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	//fmt.Printf("Received message: %s from topic: %s\n", payload, topic)
	if payload != strMap[topic] {
		switch topic {
		case "topic/test":
			fmt.Printf("Received message: %s from topic: %s\n", payload, topic)
		case "topic1":
			fmt.Printf("Received message: %s from topic: %s\n", payload, topic)
		case "topic2":
			fmt.Printf("Received message: %s from topic: %s\n", payload, topic)
		case "topic3":
			fmt.Printf("Received message: %s from topic: %s\n", payload, topic)
		}
		strMap[topic] = payload
	}
}

// 初始化
var messagePubHandlerInit mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	strMap[topic] = payload
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func listenSub(client mqtt.Client, topic string) {
	for {
		time.Sleep(time.Second)
		sub(client, topic)
	}
}

func publish(client mqtt.Client, topic string, text string) {

	token := client.Publish(topic, 0, false, text)
	token.Wait()
	time.Sleep(time.Second)

}

func sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	//fmt.Printf("Subscribed to topic: %s", topic)
}
