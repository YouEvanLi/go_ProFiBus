package serial

// SerialPort 是一个通用的串口通信接口。
type SerialPort interface {
	Open(portName string, options ...Option) error
	Write(data []byte) (int, error)
	Read(buffer []byte) (int, error)
	Close() error
}
