package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	capacity        int
	numOfTokens     int
	refillFrequency int
	lastRefillStamp time.Time
}

func newTokenBucket(capacity, refillFrequency int) *TokenBucket {
	return &TokenBucket{
		capacity:        capacity,
		numOfTokens:     capacity,
		refillFrequency: refillFrequency,
		lastRefillStamp: time.Now(),
	}
}

func (bucket *TokenBucket) canProcessRequests(numOfRequests int) bool {
	bucket.refill()
	fmt.Println(bucket.capacity)
	fmt.Println(bucket.lastRefillStamp)

	if numOfRequests <= bucket.numOfTokens {
		bucket.numOfTokens -= numOfRequests
		fmt.Println(numOfRequests, "have been consumed")
		fmt.Println(bucket.numOfTokens, "tokens are remaning in the bucket")
		return true
	}

	return false

}

func (bucket *TokenBucket) refill() {

	now := time.Now()
	elapsed := now.Sub(bucket.lastRefillStamp)
	bucket.numOfTokens += int(elapsed.Seconds()) * bucket.refillFrequency
	if bucket.numOfTokens > bucket.capacity {
		bucket.numOfTokens = bucket.capacity
	}
	bucket.lastRefillStamp = time.Now()
	fmt.Println("Time elapsed", elapsed)

}

func tokenBucket() {
	bucket := newTokenBucket(10, 3)
	fmt.Println("token bucket algorithm")

	// Make a call every 200ms aka 1/5th of a second

	for i := 0; i < 15; i++ {
		if bucket.canProcessRequests(2) {
			fmt.Println("Token consumed")
		} else {
			fmt.Println("No token available")
		}
		time.Sleep(time.Millisecond * 200)
	}

}

func main() {

	fmt.Println("rate limit algorithms")
	tokenBucket()

}
