package serial

import (
	"fmt"
	"os"
)

// Modbus 是 Modbus 设备通信的实现。
type Modbus struct {
	file *os.File
}

// Open 打开 Modbus 设备。
func (m *Modbus) Open(devicePath string) error {
	file, err := os.OpenFile(devicePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("无法打开 Modbus 设备：%v", err)
	}
	m.file = file
	return nil
}

// Write 写入数据到 Modbus。
func (m *Modbus) Write(data []byte) (int, error) {
	if m.file == nil {
		return 0, fmt.Errorf("Modbus 设备未打开")
	}
	n, err := m.file.Write(data)
	if err != nil {
		return n, fmt.Errorf("写入数据失败：%v", err)
	}
	return n, nil
}

// Read 从 Modbus 读取数据。
func (m *Modbus) Read(buffer []byte) (int, error) {
	if m.file == nil {
		return 0, fmt.Errorf("Modbus 设备未打开")
	}
	n, err := m.file.Read(buffer)
	if err != nil {
		return n, fmt.Errorf("读取数据失败：%v", err)
	}
	return n, nil
}

// Close 关闭 Modbus 设备。
func (m *Modbus) Close() error {
	if m.file == nil {
		return fmt.Errorf("Modbus 设备未打开")
	}
	err := m.file.Close()
	if err != nil {
		return fmt.Errorf("关闭 Modbus 设备失败：%v", err)
	}
	m.file = nil
	return nil
} 