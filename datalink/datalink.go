package datalink

import (
	"errors"
	"github.com/sigurn/crc16"
)

const (
	FrameStart   byte = 0x7E
	MaxFrameSize      = 1024
)

type Frame struct {
	Start byte
	Data  []byte
	CRC   uint16
}

func NewFrame(data []byte, table *crc16.Table) *Frame {
	crc := crc16.Checksum(data, table)
	return &Frame{Start: FrameStart, Data: data, CRC: crc}
}

func (f *Frame) Marshal() []byte {
	result := make([]byte, 0, len(f.Data)+4)
	result = append(result, f.Start)
	result = append(result, f.Data...)
	result = append(result, byte(f.CRC>>8), byte(f.CRC&0xFF))
	return result
}

func Unmarshal(data []byte, table *crc16.Table) (*Frame, error) {
	if len(data) < 4 {
		return nil, errors.New("frame too short")
	}
	if data[0] != FrameStart {
		return nil, errors.New("invalid frame start")
	}
	crc := uint16(data[len(data)-2])<<8 | uint16(data[len(data)-1])
	dataContent := data[1 : len(data)-2]
	if crc != crc16.Checksum(dataContent, table) {
		return nil, errors.New("CRC check failed")
	}
	return &Frame{Start: FrameStart, Data: dataContent, CRC: crc}, nil
}
