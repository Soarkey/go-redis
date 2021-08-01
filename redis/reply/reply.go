package reply

import (
	"bytes"
	"strconv"

	"go-redis/interface/redis"
)

var (
	// $-1表示nil响应, 例如get命令查询一个不存在的key时返回$-1
	nullBulkReplyBytes = []byte("$-1")

	// CRLF 是redis序列化协议中行与行的分割符
	CRLF = "\r\n"
)

/* ---- Bulk Reply ---- */

// BulkReply 单条二进制安全字符串响应信息
type BulkReply struct {
	Arg []byte
}

// ToBytes marshal redis.Reply
// 例如
// $4
// a\r\nb
func (r *BulkReply) ToBytes() []byte {
	if len(r.Arg) == 0 {
		return nullBulkReplyBytes
	}
	return []byte("$" + strconv.Itoa(len(r.Arg)) + CRLF + string(r.Arg) + CRLF)
}

// MakeBulkReply 创建BulkReply
func MakeBulkReply(arg []byte) *BulkReply {
	return &BulkReply{Arg: arg}
}

/* ---- Multi Bulk Reply ---- */

// MultiBulkReply 字符串列表响应信息
type MultiBulkReply struct {
	Args [][]byte
}

// ToBytes marshal redis.Reply
// 例如
// *2
// $3
// foo
// $3
// bar
func (r *MultiBulkReply) ToBytes() []byte {
	argLen := len(r.Args)
	var buf bytes.Buffer
	buf.WriteString("*" + strconv.Itoa(argLen) + CRLF)
	for _, arg := range r.Args {
		if arg == nil {
			buf.WriteString("$-1" + CRLF)
		} else {
			buf.WriteString("$" + strconv.Itoa(len(arg)) + CRLF + string(arg) + CRLF)
		}
	}
	return buf.Bytes()
}

// MakeMultiBulkReply 创建 MultiBulkReply
func MakeMultiBulkReply(args [][]byte) *MultiBulkReply {
	return &MultiBulkReply{Args: args}
}

/* ---- Multi Raw Reply ---- */

// MultiRawReply 复杂列表结构响应信息, 如GeoPos命令
type MultiRawReply struct {
	Replies []redis.Reply
}

// ToBytes marshal redis.Reply
func (r *MultiRawReply) ToBytes() []byte {
	argLen := len(r.Replies)
	var buf bytes.Buffer
	buf.WriteString("*" + strconv.Itoa(argLen) + CRLF)
	for _, arg := range r.Replies {
		buf.Write(arg.ToBytes())
	}
	return buf.Bytes()
}

// MakeMultiRawReply 创建 MultiRawReply
func MakeMultiRawReply(replies []redis.Reply) *MultiRawReply {
	return &MultiRawReply{Replies: replies}
}

/* ---- Status Reply ---- */

// StatusReply 单条状态字符串响应信息
type StatusReply struct {
	Status string
}

// ToBytes marshal redis.Reply
// 例如
// +OK\r\n
func (r *StatusReply) ToBytes() []byte {
	return []byte("+" + r.Status + CRLF)
}

// MakeStatusReply 创建 StatusReply
func MakeStatusReply(status string) *StatusReply {
	return &StatusReply{Status: status}
}

/* ---- Int Reply ---- */

// IntReply int64数字响应信息
type IntReply struct {
	Code int64
}

// ToBytes marshal redis.Reply
// 例如
// :1\r\n
func (r *IntReply) ToBytes() []byte {
	return []byte(":" + strconv.FormatInt(r.Code, 10) + CRLF)
}

// MakeIntReply 创建 IntReply
func MakeIntReply(code int64) *IntReply {
	return &IntReply{Code: code}
}

/* ---- Error Reply ---- */

// ErrorReply 错误返回类型
type ErrorReply interface {
	Error() string
	ToBytes() []byte
}

// StandardErrReply 表示服务端错误
type StandardErrReply struct {
	Status string
}

// ToBytes marshal redis.Reply
// 例如
// -ERR Invalid Synatx\r\n
func (r *StandardErrReply) ToBytes() []byte {
	return []byte("-" + r.Status + CRLF)
}

func (r *StandardErrReply) Error() string {
	return r.Status
}

// MakeErrReply 创建 StandardErrReply
func MakeErrReply(status string) *StandardErrReply {
	return &StandardErrReply{Status: status}
}

// IsErrorReply 如果返回值为错误返回类型, 则为true
func IsErrorReply(reply redis.Reply) bool {
	return reply.ToBytes()[0] == '-'
}
