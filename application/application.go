package application

import (
	"errors"
	"github.com/sigurn/crc16"
	"go_ProFiBus/datalink"
	"go_ProFiBus/serial"
	"sync"
	"time"
)

type ProFiBus struct {
	physicalLayer *serial.RS485
	crcTable      *crc16.Table
	sendChan      chan []byte
	receiveChan   chan []byte
	closeChan     chan struct{}
	wg            sync.WaitGroup
}

func NewProFiBus(portName string, baudRate int, polynomial uint16, initialValue uint16) (*ProFiBus, error) {
	phy, err := serial.NewRS485(portName, baudRate)
	if err != nil {
		return nil, err
	}
	table := crc16.MakeTable(crc16.Params{Poly: polynomial, Init: initialValue})
	bus := &ProFiBus{
		physicalLayer: phy,
		crcTable:      table,
		sendChan:      make(chan []byte, 10),
		receiveChan:   make(chan []byte, 10),
		closeChan:     make(chan struct{}),
	}
	bus.wg.Add(2)
	go bus.sendLoop()
	go bus.receiveLoop()
	return bus, nil
}

func (bus *ProFiBus) sendLoop() {
	defer bus.wg.Done()
	for {
		select {
		case data := <-bus.sendChan:
			frame := datalink.NewFrame(data, bus.crcTable)
			frameData := frame.Marshal()
			_ = bus.physicalLayer.Write(frameData)
		case <-bus.closeChan:
			return
		}
	}
}

func (bus *ProFiBus) receiveLoop() {
	defer bus.wg.Done()
	buffer := make([]byte, datalink.MaxFrameSize)
	for {
		select {
		case <-bus.closeChan:
			return
		default:
			n, err := bus.physicalLayer.Read(buffer)
			if err != nil {
				continue
			}
			frame, err := datalink.Unmarshal(buffer[:n], bus.crcTable)
			if err != nil {
				continue
			}
			bus.receiveChan <- frame.Data
		}
	}
}

func (bus *ProFiBus) Send(data []byte) {
	bus.sendChan <- data
}

func (bus *ProFiBus) Receive() ([]byte, error) {
	select {
	case data := <-bus.receiveChan:
		return data, nil
	case <-time.After(time.Second * 5):
		return nil, errors.New("timeout")
	}
}

func (bus *ProFiBus) Close() {
	close(bus.closeChan)
	bus.wg.Wait()
	bus.physicalLayer.Close()
}
