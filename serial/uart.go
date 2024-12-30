package serial

import (
	"fmt"
	"os"
	"syscall"
)

// UART 是 UART 串口通信的实现。
type UART struct {
	file *os.File
}

// Open 打开 UART 设备。
func (u *UART) Open(portName string) error {
	file, err := os.OpenFile(portName, syscall.O_RDWR|syscall.O_NOCTTY, 0666)
	if err != nil {
		return fmt.Errorf("无法打开 UART 设备：%v", err)
	}
	u.file = file
	return nil
}

// Write 写入数据到 UART。
func (u *UART) Write(data []byte) (int, error) {
	if u.file == nil {
		return 0, fmt.Errorf("UART 设备未打开")
	}
	n, err := u.file.Write(data)
	if err != nil {
		return n, fmt.Errorf("写入数据失败：%v", err)
	}
	return n, nil
}

// Read 从 UART 读取数据。
func (u *UART) Read(buffer []byte) (int, error) {
	if u.file == nil {
		return 0, fmt.Errorf("UART 设备未打开")
	}
	n, err := u.file.Read(buffer)
	if err != nil {
		return n, fmt.Errorf("读取数据失败：%v", err)
	}
	return n, nil
}

// Close 关闭 UART 设备。
func (u *UART) Close() error {
	if u.file == nil {
		return fmt.Errorf("UART 设备未打开")
	}
	err := u.file.Close()
	if err != nil {
		return fmt.Errorf("关闭 UART 设备失败：%v", err)
	}
	u.file = nil
	return nil
} 