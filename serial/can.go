package serial

import (
	"fmt"
	"os"
	"syscall"
)

// CAN 是 CAN 串口通信的实现。
type CAN struct {
	file *os.File
}

// Open 打开 CAN 设备。
func (c *CAN) Open(portName string) error {
	file, err := os.OpenFile(portName, syscall.O_RDWR|syscall.O_NOCTTY, 0666)
	if err != nil {
		return fmt.Errorf("无法打开 CAN 设备：%v", err)
	}
	c.file = file
	return nil
}

// Write 写入数据到 CAN。
func (c *CAN) Write(data []byte) (int, error) {
	if c.file == nil {
		return 0, fmt.Errorf("CAN 设备未打开")
	}
	n, err := c.file.Write(data)
	if err != nil {
		return n, fmt.Errorf("写入数据失败：%v", err)
	}
	return n, nil
}

// Read 从 CAN 读取数据。
func (c *CAN) Read(buffer []byte) (int, error) {
	if c.file == nil {
		return 0, fmt.Errorf("CAN 设备未打开")
	}
	n, err := c.file.Read(buffer)
	if err != nil {
		return n, fmt.Errorf("读取数据失败：%v", err)
	}
	return n, nil
}

// Close 关闭 CAN 设备。
func (c *CAN) Close() error {
	if c.file == nil {
		return fmt.Errorf("CAN 设备未打开")
	}
	err := c.file.Close()
	if err != nil {
		return fmt.Errorf("关闭 CAN 设备失败：%v", err)
	}
	c.file = nil
	return nil
} 