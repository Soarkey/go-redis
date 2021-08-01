package reply

var pongBytes = []byte("+PONG\r\n")

// PongReply is +PONG
type PongReply struct{}

// ToBytes marshal redis.Reply
func (r *PongReply) ToBytes() []byte {
	return pongBytes
}

var okBytes = []byte("+OK\r\n")

// OkReply is +OK
type OkReply struct{}

// ToBytes marshal redis.Reply
func (r *OkReply) ToBytes() []byte {
	return okBytes
}

var theOkReply = new(OkReply)

func MakeOkReply() *OkReply {
	return theOkReply
}

var nullBulkBytes = []byte("$-1\r\n")

// NullBulkReply is empty string
type NullBulkReply struct{}

// ToBytes marshal redis.Reply
func (r *NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

// MakeNullBulkReply creates a new NullBulkReply
func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

var emptyMultiBulkBytes = []byte("*0\r\n")

// EmptyMultiBulkReply is a empty list
type EmptyMultiBulkReply struct{}

// ToBytes marshal redis.Reply
func (r *EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

// MakeEmptyMultiBulkReply creates EmptyMultiBulkReply
func MakeEmptyMultiBulkReply() *EmptyMultiBulkReply {
	return &EmptyMultiBulkReply{}
}

var noBytes = []byte("")

// NoReply respond nothing, for commands like subscribe
type NoReply struct{}

// ToBytes marshal redis.Reply
func (r *NoReply) ToBytes() []byte {
	return noBytes
}

var queuedBytes = []byte("+QUEUED\r\n")

// QueuedReply is +QUEUED
type QueuedReply struct{}

// ToBytes marshal redis.Reply
func (r *QueuedReply) ToBytes() []byte {
	return queuedBytes
}

var theQueuedReply = new(QueuedReply)

func MakeQueuedReply() *QueuedReply {
	return theQueuedReply
}
