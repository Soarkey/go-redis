package tcp

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"go-redis/lib/sync/atomic"
	"go-redis/lib/sync/wait"
)

// EchoHandler 服务端
type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

// Handle 处理请求, 接收client发送的数据行
func (h *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if h.closing.Get() {
		// handler处于关闭中将会拒绝新的连接进入
		_ = conn.Close()
	}

	client := &EchoClient{Conn: conn}
	h.activeConn.Store(client, struct{}{})

	reader := bufio.NewReader(conn)
	for {
		// 可能会碰到的问题: EOF,客户端超时,服务器过早关闭
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
				h.activeConn.Delete(client)
			} else {
				log.Fatal(err)
			}
			return
		}
		log.Printf("[%v] msg: %v", client.Conn.RemoteAddr(), msg)
		client.Waiting.Add(1)
		b := []byte(msg)
		_, _ = conn.Write(b)
		client.Waiting.Done()
	}
}

// Close 关闭EchoHandler
func (h *EchoHandler) Close() error {
	log.Println("handler shutting down...")
	h.closing.Set(true)
	h.activeConn.Range(func(key interface{}, value interface{}) bool {
		client := key.(*EchoClient)
		_ = client.Close()
		return true
	})
	return nil
}

// MakeEchoHandler 创建一个 EchoHandler
func MakeEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

// EchoClient 客户端
type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

// Close 关闭连接
func (c *EchoClient) Close() error {
	c.Waiting.WaitWithTimeout(10 * time.Second)
	c.Conn.Close()
	return nil
}
