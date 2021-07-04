package wait

import (
	"sync"
	"time"
)

// Wait 带timeout超时的WaitGroup
type Wait struct {
	wg sync.WaitGroup
}

// Add 添加delta项任务, delta可能为负数
func (w *Wait) Add(delta int) {
	w.wg.Add(delta)
}

// Done 完成一项任务
func (w *Wait) Done() {
	w.wg.Done()
}

// Wait 此方法会阻塞直到所有任务都完成
func (w *Wait) Wait() {
	w.wg.Wait()
}

// WaitWithTimeout 超时等待, 如果超时未完成返回true, 正常完成返回true
func (w *Wait) WaitWithTimeout(timeout time.Duration) bool {
	c := make(chan bool)
	go func() {
		defer close(c)
		w.wg.Wait()
		c <- true
	}()

	select {
	case <-c: // 正常完成
		return false
	case <-time.After(timeout): // 超时
		return true

	}
}
