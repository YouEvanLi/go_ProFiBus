package serial

import (
	"golang.org/x/sys/unix"
	"os"
	"syscall"
	"unsafe"
)

// RS485 是 RS-485 串口通信的实现。
type RS485 struct {
	// 你可以在这里添加 RS-485 特定的字段或配置
	fd int
}

func (r *RS485) Open(portName string, options ...Option) error {
	// 实现打开 RS-485 串口的逻辑
	file, err := os.OpenFile(portName, syscall.O_RDWR|syscall.O_NOCTTY, 0666)
	if err != nil {
		return err
	}

	fd := int(file.Fd())
	if err := unix.SetNonblock(fd, false); err != nil {
		return err
	}

	var termios unix.Termios
	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(unix.TCGETS), uintptr(unsafe.Pointer(&termios))); errno != 0 {
		return errno
	}

	termios.Cflag = unix.B115200 | unix.CS8 | unix.CLOCAL | unix.CREAD
	termios.Iflag = unix.IGNPAR
	termios.Oflag = 0
	termios.Lflag = 0

	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(unix.TCSETS), uintptr(unsafe.Pointer(&termios))); errno != 0 {
		return errno
	}

	r.fd = fd

	return nil
}

func (r *RS485) Write(data []byte) (int, error) {
	// 实现向 RS-485 串口写入数据的逻辑
	return unix.Write(r.fd, data)
}

func (r *RS485) Read(buffer []byte) (int, error) {
	// 实现从 RS-485 串口读取数据的逻辑
	return unix.Read(r.fd, buffer)
}

func (r *RS485) Close() error {
	// 实现关闭 RS-485 串口的逻辑
	return unix.Close(r.fd)
}
