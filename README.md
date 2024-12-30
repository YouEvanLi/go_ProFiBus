# go_ProFiBus

`go_ProFiBus` 是一个用于串口通信的 Go 语言库，支持多种串口协议，包括 UART、CAN、USB、1-Wire、Modbus、RS-232、RS-485 和 SPI。该库提供了统一的接口，便于在不同的串口设备之间进行通信。

## 特性

- 支持多种串口通信协议
- 统一的接口设计
- 简单易用的配置选项
- 错误处理机制
- 模拟串口通信实现

## 安装

确保你已经安装了 Go 语言环境。然后可以通过以下命令克隆项目并安装依赖：

    git clone https://github.com/YouEvanLi/go_ProFiBus.git
    
    cd go_ProFiBus
    
    go mod tidy


## 使用示例

以下是如何使用 `go_ProFiBus` 库的示例：


    package main
    import (
        "fmt"
        "go_ProFiBus/application"
        "log"
    )
    func main() {
        portName := "/dev/ttyUSB0" // 根据实际情况修改
        // 创建 UART 总线
        uartBus, err := application.NewProtocolBus(application.UART, portName)
        if err != nil {
            log.Fatalf("错误: %v", err)
        }
        defer uartBus.Close()
        // 写入数据
        dataToSend := []byte("Hello UART")
        if , err := uartBus.Write(dataToSend); err != nil {
            log.Fatalf("写入错误: %v", err)
        }
        // 读取数据
        buffer := make([]byte, 100)
        n, err := uartBus.Read(buffer)
        if err != nil {
            log.Fatalf("读取错误: %v", err)
        }
        fmt.Printf("接收到数据: %s\n", buffer[:n])
    }


## 支持的协议

- **UART**: 通用异步收发器
- **CAN**: 控制器局域网络
- **USB**: 通用串行总线
- **1-Wire**: 单线通信协议
- **Modbus**: 工业通信协议
- **RS-232**: 串行通信标准
- **RS-485**: 串行通信标准
- **SPI**: 串行外设接口
- **I2C**: 互连集成电路

## 配置选项

可以通过以下选项配置串口通信参数：

- `WithBaudRate(baudRate int)`: 设置波特率
- `WithDataBits(dataBits int)`: 设置数据位
- `WithParity(parity ParityMode)`: 设置校验位
- `WithStopBits(stopBits int)`: 设置停止位
- `WithAddress(address int)`: 设置设备地址（适用于 I2C）

## 贡献

欢迎任何形式的贡献！请提交问题、建议或拉取请求。

## 许可证

该项目使用 MIT 许可证。有关详细信息，请参阅 [LICENSE](LICENSE) 文件。