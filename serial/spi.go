package serial

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

// 定义 SpiIocTransfer 结构体
type SpiIocTransfer struct {
	TxBuf       uintptr
	RxBuf       uintptr
	Len         uint32
	SpeedHz     uint32
	DelayUsecs  uint16
	BitsPerWord uint8
	csChange    uint8
}

// 定义 SPI_IOC_MESSAGE 常量
const SPI_IOC_MESSAGE_1 = 0x40006b00

// SpiDevice 是 SPI 设备。
type SpiDevice struct {
	device *os.File // SPI 设备文件
}

// Open 打开 SPI 设备。
func (s *SpiDevice) Open(devicePath string, options ...Option) error {
	var err error

	// 打开 SPI 设备文件
	s.device, err = os.OpenFile(devicePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("无法打开 SPI 设备：%s", err)
	}

	return nil
}

// Transfer 在 SPI 设备上进行数据传输。
func (s *SpiDevice) Transfer(txData []byte, rxData []byte) error {
	if s.device == nil {
		return fmt.Errorf("SPI 设备未打开")
	}

	// 定义 SpiIocTransfer 结构体
	transfer := SpiIocTransfer{
		TxBuf:   uintptr(unsafe.Pointer(&txData[0])),
		RxBuf:   uintptr(unsafe.Pointer(&rxData[0])),
		Len:     uint32(len(txData)),
		SpeedHz: 1000000, // 设置 SPI 时钟速度
	}

	// 使用 unix.IoctlSetInt 替代 ioctl
	err := unix.IoctlSetInt(int(s.device.Fd()), SPI_IOC_MESSAGE_1, int(uintptr(unsafe.Pointer(&transfer))))
	if err != nil {
		return fmt.Errorf("SPI 数据传输失败：%s", err)
	}

	return nil
}

// Write 向 SPI 设备写入数据。
func (s *SpiDevice) Write(data []byte) (int, error) {
	if s.device == nil {
		return 0, fmt.Errorf("SPI 设备未打开")
	}

	// 执行 SPI 写操作
	_, err := s.device.Write(data)
	if err != nil {
		return 0, fmt.Errorf("SPI 写入数据失败：%s", err)
	}

	return 0, nil
}

// Read 从 SPI 设备读取数据。
func (s *SpiDevice) Read(data []byte) (int, error) {
	if s.device == nil {
		return 0, fmt.Errorf("SPI 设备未打开")
	}

	// 执行 SPI 读操作
	_, err := s.device.Read(data)
	if err != nil {
		return 0, fmt.Errorf("SPI 读取数据失败：%s", err)
	}

	return 0, nil
}

// Close 关闭 SPI 设备。
func (s *SpiDevice) Close() error {
	if s.device == nil {
		return nil // 设备已经关闭，无需重复操作
	}

	err := s.device.Close()
	if err != nil {
		return fmt.Errorf("关闭 SPI 设备失败：%s", err)
	}

	return nil
}
