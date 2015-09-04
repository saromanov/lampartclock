package lampartclock

import (
	"sync/atomic"
)

type LampartClock struct {
	counter uint64
}

//Inc provides increment counter
func (lc *LampartClock) Inc() {
	atomic.AddUint64(&lc.counter, 1)
}

//Get provides current counter
func (lc *LampartClock) Get() uint64 {
	return atomic.LoadUint64(&lc.counter)
}

func (lc *LampartClock) Set(num uint64) {
	value := lc.Get()
	if num < value {
		return
	} else {
		atomic.StoreUint64(&lc.counter, value+num)
	}
}
