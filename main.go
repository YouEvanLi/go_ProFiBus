package main

import (
	"fmt"
	"go_ProFiBus/application"
	"log"
	"time"
)

func main() {
	polynomial := uint16(0x1021) // CRC-CCITT
	initialValue := uint16(0xFFFF)

	bus, err := application.NewProFiBus("/dev/ttyS0", 115200, polynomial, initialValue)
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	go func() {
		for {
			data := []byte("Hello ProFiBus")
			bus.Send(data)
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		data, err := bus.Receive()
		if err != nil {
			log.Println("Receive error:", err)
			continue
		}
		fmt.Printf("Received: %s\n", data)
	}
}
