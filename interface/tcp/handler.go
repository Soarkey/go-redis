package tcp

import (
	"context"
	"net"
)

// Handler 是应用层服务器 application server 的抽象
type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}
