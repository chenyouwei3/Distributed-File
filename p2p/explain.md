# 项目结构
    .
    ├── p2p //相关业务逻辑
    |   ├──  encoding.go //解码器
    |   ├──  handshake.go //握手函数 
    |   ├──  message //消息结构
    |   ├──  tcp_transports.go //tcp传输结构以及函数
    |   ├──  tcp_transports_test.go //tcp传输结构以及函数 测试
    |   └──  transports.go //传输方法
    |
    ├── go.mod //依赖
    └── main.go //启动入口
