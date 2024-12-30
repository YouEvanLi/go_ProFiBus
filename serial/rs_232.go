package serial

import (
	"fmt"
	"os"
	"syscall"
)

// RS232 是 RS-232 串口通信的实现。
type RS232 struct {
	config *SerialConfig
	file   *os.File // 串口文件
}

func (r *RS232) Open(portName string, options ...Option) error {
	// 实现打开 RS-232 串口的逻辑
	file, err := os.OpenFile(portName, syscall.O_RDWR|syscall.O_NOCTTY, 0666)
	if err != nil {
		return fmt.Errorf("无法打开 RS-232 设备：%s", err)
	}

	r.config.PortName = portName
	r.file = file
	return nil
}

// Write 写入数据到 RS-232 设备。
func (r *RS232) Write(data []byte) (int, error) {
	if r.file == nil {
		return 0, fmt.Errorf("RS-232 设备未打开")
	}

	// 写入数据到设备
	n, err := r.file.Write(data)
	if err != nil {
		return n, fmt.Errorf("写入数据失败：%s", err)
	}

	return n, nil
}

// Read 从 RS-232 设备读取数据。
func (r *RS232) Read(buffer []byte) (int, error) {
	if r.file == nil {
		return 0, fmt.Errorf("RS-232 设备未打开")
	}

	// 从设备读取数据
	n, err := r.file.Read(buffer)
	if err != nil {
		return n, fmt.Errorf("读取数据失败：%s", err)
	}

	return n, nil
}
func (r *RS232) Close() error {
	if r.file == nil {
		return fmt.Errorf("RS-232 设备未打开")
	}

	err := r.file.Close()
	if err != nil {
		return fmt.Errorf("关闭 RS-232 设备失败：%s", err)
	}
	r.file = nil // 确保文件指针被清空
	return nil
}
