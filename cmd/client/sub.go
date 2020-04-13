package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var (
	num = 0
)
//创建全局mqtt sub消息处理 handler
var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	num++
	if 10 < num {
		brokerLoad <- true
	}

	fmt.Printf("Sub Client Topic : %s \n", msg.Topic())
	fmt.Printf("Sub Client msg : %s \n", msg.Payload())
}


func main() {
	taskId := 1
	//设置连接参数
	clinetOptions := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetUsername(user).SetPassword(password)
	//设置客户端ID
	clinetOptions.SetClientID(fmt.Sprintf("go Subscribe client example： %d-%d", taskId, time.Now().Unix()))
	//设置连接超时
	clinetOptions.SetConnectTimeout(1 * time.Second)
	//创建客户端连接
	client := mqtt.NewClient(clinetOptions)

	//客户端连接判断
	if token := client.Connect(); token.WaitTimeout(1*time.Second) && token.Wait() && token.Error() != nil {
		fmt.Printf("[Sub] mqtt connect error, taskId: %d,  error: %s \n", taskId, token.Error())
		return
	}

	token := client.Subscribe("go-test-topic", 1, messageSubHandler)
	token.WaitTimeout(10 * time.Second)
	fmt.Printf("[Sub] end Subscribe msg to mqtt broker, taskId: %d, token : %s \n", taskId, token)
	token.Wait()

	<-brokerLoad

	client.Disconnect(250)
	fmt.Println("[Sub] task is ok")
}
