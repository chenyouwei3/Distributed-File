package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn //底层连接
	//如果dial就是true (客户端->服务端)
	//accept就是false (服务端->客户端)
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Close 实现interface的接口
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

// TCPTransportOps TCP传输的配置参数
type TCPTransportOps struct {
	ListenAddr    string // 监听地址
	HandshakeFunc HandshakeFunc
	Decoder       Decoder //解码器
	OnPeer        func(Peer) error
}

// TCPTransport TCP传输的结构体
type TCPTransport struct {
	TCPTransportOps
	Listener net.Listener // TCP 监听器
	rpcch    chan RPC
	mu       sync.RWMutex // 读写锁，用于保护并发访问 prees
}

// NewTCPTransport 创建一个新的 TCPTransport 实例并初始化
func NewTCPTransport(opts TCPTransportOps) *TCPTransport {
	return &TCPTransport{
		TCPTransportOps: opts,
		rpcch:           make(chan RPC),
	}
}

// Consume 返回消费者channel 返回已读内容的唯一通道
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

// ListenAndAccept 开始监听指定地址并接收传入的连接
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddr) // 在指定地址上启动 TCP 监听
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

// startAcceptLoop 运行一个循环，接受传入的连接
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept() // 接受传入的连接
		if err != nil {
			fmt.Println("TCP accept error", err)
		}
		fmt.Printf("new incoming connection %+v\n", conn)
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error
	defer func() {
		fmt.Printf("dropping peer connection:%s", err)
		conn.Close()
	}()
	peer := NewTCPPeer(conn, true) //客户端向我们传输东西
	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error:%s\n", err)
		return
	}
	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}
	rpc := RPC{}
	//buf := make([]byte, 2000)
	//读取连接
	//1.24.34
	for {
		//n, err := conn.Read(buf)
		//if err != nil {
		//	fmt.Printf("TCP error: %s\n", err)
		//}
		//fmt.Printf("message:%+v\n", buf[:n])
		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}
		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc
	}
}
