package serial

import (
	"fmt"
	"os"
)

// OneWire 是 1-Wire 设备通信的实现。
type OneWire struct {
	file *os.File
}

// Open 打开 1-Wire 设备。
func (o *OneWire) Open(devicePath string) error {
	file, err := os.OpenFile(devicePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("无法打开 1-Wire 设备：%v", err)
	}
	o.file = file
	return nil
}

// Write 写入数据到 1-Wire。
func (o *OneWire) Write(data []byte) (int, error) {
	if o.file == nil {
		return 0, fmt.Errorf("1-Wire 设备未打开")
	}
	n, err := o.file.Write(data)
	if err != nil {
		return n, fmt.Errorf("写入数据失败：%v", err)
	}
	return n, nil
}

// Read 从 1-Wire 读取数据。
func (o *OneWire) Read(buffer []byte) (int, error) {
	if o.file == nil {
		return 0, fmt.Errorf("1-Wire 设备未打开")
	}
	n, err := o.file.Read(buffer)
	if err != nil {
		return n, fmt.Errorf("读取数据失败：%v", err)
	}
	return n, nil
}

// Close 关闭 1-Wire 设备。
func (o *OneWire) Close() error {
	if o.file == nil {
		return fmt.Errorf("1-Wire 设备未打开")
	}
	err := o.file.Close()
	if err != nil {
		return fmt.Errorf("关闭 1-Wire 设备失败：%v", err)
	}
	o.file = nil
	return nil
} 