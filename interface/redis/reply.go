package redis

// Reply redis序列化协议消息通用接口
type Reply interface {
	ToBytes() []byte
}
