package atomic

import "sync/atomic"

// Boolean 满足原子操作的布尔值
type Boolean uint32

// Get 原子读
func (b *Boolean) Get() bool {
	return atomic.LoadUint32((*uint32)(b)) != 0
}

// Set 原子写
func (b *Boolean) Set(v bool) {
	if v {
		atomic.StoreUint32((*uint32)(b), 1)
		return
	}
	atomic.StoreUint32((*uint32)(b), 0)
}
