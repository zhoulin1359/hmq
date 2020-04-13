package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/fhmq/hmq/broker"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conf := os.Args[1:]
	conf = append(conf, "-config=./conf/hmq.config")
	fmt.Println(conf)

	config, err := broker.ConfigureConfig(conf)
	if err != nil {
		log.Fatal("configure broker config error: ", err)
	}

	b, err := broker.NewBroker(config)
	if err != nil {
		log.Fatal("New Broker error: ", err)
	}
	//b.CheckConnectAuth()
	b.Start()

	s := waitForSignal()
	log.Println("signal received, broker closed.", s)
}

func waitForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
