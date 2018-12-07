package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/croomes/counter/pkg/adapter"
	"github.com/croomes/counter/pkg/server"
	"gobot.io/x/gobot/drivers/i2c"
)

func main() {

	// Parse command-line flags
	var device string
	flag.StringVar(&device, "device", "", "i2c device")
	flag.Parse()

	// Initialise i2c Adapter
	adapter, err := adapter.New(device)
	if err != nil {
		log.Fatal(err)
	}
	adapter.Connect()

	// Initialise LCD
	lcd := i2c.NewHT16K33Driver(adapter)
	if err := lcd.Start(); err != nil {
		log.Fatal(err.Error())
	}
	defer lcd.Halt()

	// Start API server
	server := server.New(lcd)
	go func() {
		server.Run()
	}()
	defer server.Shutdown()

	// Wait for Ctrl-C
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

}
