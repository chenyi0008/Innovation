package mqtt

import (
	"Innovation/model"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

const (
	username = "caibin"
	password = "caibin@123"
	port     = 1883
	broker   = "111.230.194.164"
	clientId = "go_mqtt_client2"
)

var (
	flag      = false
	topicList []string
	client    mqtt.Client
)

func MqttTopicListAppend(topic string) {
	topicList = append(topicList, topic)
}

func MqttInspectListInit() {
	inspectList := model.InspectGetAllMqtt()
	for _, inspect := range inspectList {
		serialNum := inspect.SerialNum
		s := fmt.Sprintf("device/detect/%s/warning", serialNum)
		MqttTopicListAppend(s)
	}

}

func MqttMain() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topicList = append(topicList, "topic1")
	topicList = append(topicList, "topic2")
	topicList = append(topicList, "topic3")

	for _, topic := range topicList {
		listenSubInit(client, topic)
		fmt.Println("subscribe topic:", topic)
	}
	time.Sleep(time.Second)

	flag = true

}

// 全局MQTT pub消息处理
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	//初始化
	if !flag {
		WriteToMap(topic, payload)
		return
	}

	if payload != ReadFromMap(topic) {
		ProcessRequest(topic, payload)
		WriteToMap(topic, payload)
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func listenSub(client mqtt.Client, topic string) {
	//for {
	time.Sleep(time.Second)
	sub(client, topic)
	//}
}

// 初始化
func listenSubInit(client mqtt.Client, topic string) {

	sub(client, topic)

}

func publish(topic string, text string) {
	fmt.Println("publish topic:", topic)
	token := client.Publish(topic, 0, false, text)
	_ = token

}

func sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	//fmt.Printf("Subscribed to topic: %s", topic)
}
