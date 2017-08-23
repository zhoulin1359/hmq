package main

import (
	"fhmq/broker"
	"os"
	"os/signal"

	log "github.com/cihub/seelog"
)

const (
	MaxUser               = 1024 * 1024
	MessagePoolNum        = 1024
	MessagePoolUser       = MaxUser / MessagePoolNum
	MessagePoolMessageNum = MaxUser / MessagePoolNum * 4
)

var (
	MSGPool []MessagePool
)

func init() {
	MSGPool = make([]MessagePool, (MessagePoolNum + 2))
	for i := 0; i < (MessagePoolNum + 2); i++ {
		MSGPool[i].Init(MessagePoolUser, MessagePoolMessageNum)
	}
}
func main() {
	broker := broker.NewBroker()
	broker.StartListening()

	s := waitForSignal()
	log.Infof("signal got: %v ,broker closed.", s)
}

func waitForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}