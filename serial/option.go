package serial

// Option 是用于配置串口通信参数的选项。
type Option func(config *SerialConfig) error

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
