package serial

import (
	"fmt"
	"os"
)

// USB 是 USB 设备通信的实现。
type USB struct {
	file *os.File
}

// Open 打开 USB 设备。
func (u *USB) Open(devicePath string) error {
	file, err := os.OpenFile(devicePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("无法打开 USB 设备：%v", err)
	}
	u.file = file
	return nil
}

// Write 写入数据到 USB。
func (u *USB) Write(data []byte) (int, error) {
	if u.file == nil {
		return 0, fmt.Errorf("USB 设备未打开")
	}
	n, err := u.file.Write(data)
	if err != nil {
		return n, fmt.Errorf("写入数据失败：%v", err)
	}
	return n, nil
}

// Read 从 USB 读取数据。
func (u *USB) Read(buffer []byte) (int, error) {
	if u.file == nil {
		return 0, fmt.Errorf("USB 设备未打开")
	}
	n, err := u.file.Read(buffer)
	if err != nil {
		return n, fmt.Errorf("读取数据失败：%v", err)
	}
	return n, nil
}

// Close 关闭 USB 设备。
func (u *USB) Close() error {
	if u.file == nil {
		return fmt.Errorf("USB 设备未打开")
	}
	err := u.file.Close()
	if err != nil {
		return fmt.Errorf("关闭 USB 设备失败：%v", err)
	}
	u.file = nil
	return nil
} 