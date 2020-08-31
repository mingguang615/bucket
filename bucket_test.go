package bucket

import (
	"fmt"
	"testing"
	"time"
)

func TestNewBucket(t *testing.T) {
	bucket := NewBucket(100, time.Duration(time.Millisecond))

	i := 10000
	for i > 0 {
		token := bucket.GetToken()
		if !token {
			fmt.Println(token)
		}
		i--
	}

	fmt.Println(bucket.GetBucketSize())
	bucket.SetBucketSize(1000)
	fmt.Println(bucket.GetBucketSize())
}
