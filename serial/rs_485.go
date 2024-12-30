package serial

import (
	"fmt"
	"os"
	"syscall"
	"golang.org/x/sys/unix"
	"unsafe"
)

// RS485 是 RS-485 串口通信的实现。
type RS485 struct {
	fd int
}

// NewRS485 创建一个新的 RS485 实例并打开指定的端口。
func NewRS485(portName string, baudRate int) (*RS485, error) {
	r := &RS485{}
	if err := r.Open(portName); err != nil {
		return nil, fmt.Errorf("failed to open RS485 port: %v", err)
	}
	return r, nil
}

func (r *RS485) Open(portName string, options ...Option) error {
	file, err := os.OpenFile(portName, syscall.O_RDWR|syscall.O_NOCTTY, 0666)
	if err != nil {
		return fmt.Errorf("无法打开 RS-485 设备：%v", err)
	}
	defer func() {
		if err != nil {
			file.Close()
		}
	}()

	fd := int(file.Fd())
	if err := unix.SetNonblock(fd, false); err != nil {
		return fmt.Errorf("设置非阻塞模式失败：%v", err)
	}

	var termios unix.Termios
	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(unix.TCGETS), uintptr(unsafe.Pointer(&termios))); errno != 0 {
		return fmt.Errorf("获取终端属性失败：%v", errno)
	}

	termios.Cflag = unix.B115200 | unix.CS8 | unix.CLOCAL | unix.CREAD
	termios.Iflag = unix.IGNPAR
	termios.Oflag = 0
	termios.Lflag = 0

	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(unix.TCSETS), uintptr(unsafe.Pointer(&termios))); errno != 0 {
		return fmt.Errorf("设置终端属性失败：%v", errno)
	}

	r.fd = fd
	return nil
}

func (r *RS485) Write(data []byte) (int, error) {
	return unix.Write(r.fd, data)
}

func (r *RS485) Read(buffer []byte) (int, error) {
	return unix.Read(r.fd, buffer)
}

func (r *RS485) Close() error {
	if r.fd == 0 {
		return fmt.Errorf("RS-485 设备未打开")
	}
	if err := unix.Close(r.fd); err != nil {
		return fmt.Errorf("关闭 RS-485 设备失败：%v", err)
	}
	r.fd = 0 // 确保文件描述符被清空
	return nil
}
