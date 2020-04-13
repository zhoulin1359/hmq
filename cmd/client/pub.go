package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)
//创建全局mqtt publish消息处理 handler
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Pub Client Topic : %s \n", msg.Topic())
	fmt.Printf("Pub Client msg : %s \n", msg.Payload())
}

func main() {
	taskId := 3
	clinetOptions := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetUsername(user).SetPassword(password)
	//设置客户端ID
	clinetOptions.SetClientID(fmt.Sprintf("go Publish client example： %d", time.Now().Unix()))
	//设置handler
	clinetOptions.SetDefaultPublishHandler(messagePubHandler)
	//设置连接超时
	clinetOptions.SetConnectTimeout(1 * time.Second)
	clinetOptions.SetProtocolVersion(4)
	//创建客户端连接
	client := mqtt.NewClient(clinetOptions)

	//客户端连接判断
	if token := client.Connect(); token.WaitTimeout(1*time.Second) && token.Wait() && token.Error() != nil {
		fmt.Printf("[Pub] mqtt connect error, taskId: %d, error: %s \n", taskId, token.Error())
		return
	}

	text := fmt.Sprintf("this is test msg #%d ! from task :%d", 1, taskId)
	//fmt.Printf("start publish msg to mqtt broker, taskId: %d, count: %d \n", taskId, i)
	//发布消息
	token := client.Publish("2", 1, false, text)
	fmt.Printf("[Pub] end publish msg to mqtt broker, taskId: %d, count: %d, token : %s \n", taskId, token)
	token.Wait()

	client.Disconnect(250)
}
