package lampartclock

import
(
	"sync"
	"sync/atomic"
)

type LampartClock {
	counter uint64
}

//Inc provides increment counter
func (lc *LampartClock) Inc() {
	atomic.AddUint64(&lc.counter, 1)
}

//Get provides current counter
func (lc *LampartClock) Get() int64 {
	return atomic.LoadUint64(&lc.counter)
}

func (lc *LampartClock) Set(num uint64) {
	value := lc.Get()
	if num < value {
		return 
	} else {
		atomic.StoreUnit64(&lc.counter, num)
	}
}