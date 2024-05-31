package serial

import "fmt"

// Option 是用于配置串口通信参数的选项。
type Option func(config *SerialConfig) error

// SimulatedSerialPort 是一个模拟的串口通信实现。
type SimulatedSerialPort struct {
	config *SerialConfig
}

// SerialConfig 是串口通信的配置。
type SerialConfig struct {
	PortName  string     // 端口名称
	BaudRate  int        // 波特率
	DataBits  int        // 数据位
	Parity    ParityMode // 校验位
	StopBits  int        // 停止位
	Simulated bool       // 是否使用模拟模式
	Address   int
}

// ParityMode 是校验位模式。
type ParityMode int

const (
	ParityNone ParityMode = iota
	ParityOdd
	ParityEven
)

// UseSimulatedMode 将串口通信设置为模拟模式。
func UseSimulatedMode() Option {
	return func(config *SerialConfig) error {
		config.Simulated = true
		return nil
	}
}

// WithBaudRate 设置串口通信的波特率。
func WithBaudRate(baudRate int) Option {
	return func(config *SerialConfig) error {
		config.BaudRate = baudRate
		return nil
	}
}

// WithDataBits 设置串口通信的数据位。
func WithDataBits(dataBits int) Option {
	return func(config *SerialConfig) error {
		config.DataBits = dataBits
		return nil
	}
}

// WithParity 设置串口通信的校验位。
func WithParity(parity ParityMode) Option {
	return func(config *SerialConfig) error {
		config.Parity = parity
		return nil
	}
}

// WithStopBits 设置串口通信的停止位。
func WithStopBits(stopBits int) Option {
	return func(config *SerialConfig) error {
		config.StopBits = stopBits
		return nil
	}
}

// NewSerialPort 返回一个串口通信实例。
func NewSerialPort(options ...Option) SerialPort {
	config := &SerialConfig{
		Simulated: true, // 默认使用模拟模式
	}

	// 应用选项
	for _, option := range options {
		option(config)
	}

	if config.Simulated {
		return &SimulatedSerialPort{config: config}
	}
	// 如果使用真实模式，返回相应的实际实现
	return nil // 实际实现未提供，这里返回 nil
}

// Open 打开串口。
func (s *SimulatedSerialPort) Open(portName string, options ...Option) error {
	s.config.PortName = portName
	for _, option := range options {
		option(s.config)
	}
	fmt.Println("打开串口", portName)
	fmt.Printf("波特率: %d, 数据位: %d, 校验位: %v, 停止位: %d\n",
		s.config.BaudRate, s.config.DataBits, s.config.Parity, s.config.StopBits)
	return nil
}

// Write 写入数据到串口。
func (s *SimulatedSerialPort) Write(data []byte) (int, error) {
	fmt.Printf("向串口 %s 写入数据: %s\n", s.config.PortName, data)
	return len(data), nil
}

// Read 从串口读取数据。
func (s *SimulatedSerialPort) Read(buffer []byte) (int, error) {
	fmt.Printf("从串口 %s 读取数据\n", s.config.PortName)
	return 0, nil
}

// Close 关闭串口。
func (s *SimulatedSerialPort) Close() error {
	fmt.Println("关闭串口", s.config.PortName)
	return nil
}
