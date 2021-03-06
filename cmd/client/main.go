package main

import (
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
	"time"
)






//连接失败数
var failNums = 0

/***
* 创建客户端连接
 */
func main() {
	clientNum := flag.Uint64("clientNum", 1, "client nums")
	flag.Parse()
	nums := int(*clientNum)
	waitGroup := sync.WaitGroup{}

	for i := 0; i < nums; i++ {
		fmt.Printf("publish client num : %d \n", i)
		waitGroup.Add(1)
		time.Sleep(3 * time.Millisecond)
		//调用连接和发布消息
		go mqttConnPubMsgTask(i, &waitGroup)
		//订阅
		//go mqttConnSubMsgTask(i, &waitGroup)
	}

	waitGroup.Wait()
}

/***
*
* 连接任务和发布消息方法
 */
func mqttConnPubMsgTask(taskId int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	//设置连接参数
	clinetOptions := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetUsername(user).SetPassword(password)
	//设置客户端ID
	clinetOptions.SetClientID(fmt.Sprintf("go Publish client example： %d-%d", taskId, time.Now().Unix()))
	//设置handler
	clinetOptions.SetDefaultPublishHandler(messagePubHandler)
	//设置连接超时
	clinetOptions.SetConnectTimeout(1 * time.Second)
	clinetOptions.SetProtocolVersion(4)
	//创建客户端连接
	client := mqtt.NewClient(clinetOptions)

	//客户端连接判断
	if token := client.Connect(); token.WaitTimeout(1*time.Second) && token.Wait() && token.Error() != nil {
		failNums++
		fmt.Printf("[Pub] mqtt connect error, taskId: %d, fail_nums: %d, error: %s \n", taskId, failNums, token.Error())
		return
	}

	i := 0

	for {
		i++
		time.Sleep(3 * time.Second)
		text := fmt.Sprintf("this is test msg #%d ! from task :%d", i, taskId)
		//fmt.Printf("start publish msg to mqtt broker, taskId: %d, count: %d \n", taskId, i)
		//发布消息
		token := client.Publish("go-test-topic", 1, false, text)
		fmt.Printf("[Pub] end publish msg to mqtt broker, taskId: %d, count: %d, token : %s \n", taskId, i, token)
		token.Wait()
	}

	client.Disconnect(250)
	fmt.Println("[Pub] task is ok")

}

