package tcp

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"go-redis/interface/tcp"
)

// Config 存储tcp服务器配置信息
type Config struct {
	Address    string        `yaml:"address"`
	MaxConnect uint32        `yaml:"max-connect"`
	Timeout    time.Duration `yaml:"timeout"`
}

// ListenAndServe 监听并提供服务，并在收到 closeChan 发来的关闭通知后关闭
func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) {
	// 监听关闭通知
	go func() {
		<-closeChan
		log.Println("shutting down...")
		// 停止监听，listener.Accept()会立即返回 io.EOF
		_ = listener.Close()
		// 关闭应用层服务器
		_ = handler.Close()
	}()

	// 在异常退出后释放资源
	defer func() {
		_ = listener.Close()
		_ = handler.Close()
	}()

	ctx := context.Background()
	// sync.WaitGroup 类似 java 中 CountDownLatch
	var waitDone sync.WaitGroup

	for {
		// Accept 会一直阻塞直到有新的连接建立或者listen中断才会返回
		conn, err := listener.Accept()
		if err != nil {
			// 通常是由于listener被关闭无法继续监听导致的错误
			log.Fatal(fmt.Sprintf("accept err: %v", err))
			break
		}
		// 开启 goroutine 处理新连接
		log.Printf("accept link from [%v]\n", conn.RemoteAddr())
		waitDone.Add(1)
		go func() {
			defer func() {
				waitDone.Done()
			}()
			handler.Handle(ctx, conn)
		}()
		waitDone.Wait()
	}
}
