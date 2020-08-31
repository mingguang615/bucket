package bucket

import (
	"sync"
	"time"
)

type Bucket struct {
	size   int
	ticker *time.Ticker
	ch     chan bool
	mu     *sync.Mutex
}

func NewBucket(size int, interval time.Duration) *Bucket {
	bucket := &Bucket{
		size:   size,
		ticker: time.NewTicker(interval),
		ch:     make(chan bool, size),
		mu:     new(sync.Mutex),
	}
	go bucket.startTicker()
	return bucket
}

func (bucket *Bucket) startTicker() {
	for i := 0; i < bucket.size; i++ {
		bucket.ch <- true
	}
	for {
		select {
		case <-bucket.ticker.C:
			for i := len(bucket.ch); i < bucket.size; i++ {
				bucket.addToken()
			}
		}
	}
}

func (bucket *Bucket) addToken() {
	if len(bucket.ch) < bucket.size {
		bucket.ch <- true
	}
}

func (bucket *Bucket) GetToken() bool {
	select {
	case <-bucket.ch:
		return true
	default:
		return false
	}
}

func (bucket *Bucket) GetBucketSize() int {
	bucket.mu.Lock()
	defer bucket.mu.Unlock()
	return bucket.size
}

func (bucket *Bucket) SetBucketSize(size int) {
	bucket.mu.Lock()
	defer bucket.mu.Unlock()
	if size > 0 && bucket.size != size {
		bucket.size = size
	}
}

func (bucket *Bucket) SetBucketTicker(interval time.Duration) {
	bucket.mu.Lock()
	defer bucket.mu.Unlock()
	bucket.ticker = time.NewTicker(interval)
}
