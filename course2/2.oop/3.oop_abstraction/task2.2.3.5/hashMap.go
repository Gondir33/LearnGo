package main

import (
	"hash/crc32"
	"hash/crc64"
	"math"
	"time"

	"github.com/howeyc/crc16"
	"github.com/sigurn/crc8"
)

const BucketSize = 8

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type Buckets struct {
	tophash_key [BucketSize]uint32
	key         [BucketSize]string
	value       [BucketSize]interface{}
}

type HashMap struct {
	sz       uint32
	bucketsC uint8 //log2(sz)
	buckets  []Buckets
	hasher   func(key string) uint32
}

func (hm *HashMap) Set(key string, value interface{}) {
	hashSum := hm.hasher(key)
	m := bucketMask(hm.bucketsC)
	lowbit := hashSum & uint32(m)
	bucket := &hm.buckets[lowbit]
	i := 0
	for ; i < BucketSize; i++ {
		if bucket.tophash_key[i] == 0 {
			break
		}
	}
	if i == BucketSize { // Надо ли нам делать с oldbuckets для большого места???
		return
	}
	bucket.tophash_key[i] = hashSum
	bucket.key[i] = key
	bucket.value[i] = value
}

func bucketMask(b uint8) uint8 {
	return 1<<b - 1
}

func (hm *HashMap) Get(key string) (interface{}, bool) {
	hashSum := hm.hasher(key)
	m := bucketMask(hm.bucketsC)
	lowbit := hashSum & uint32(m)
	bucket := &hm.buckets[lowbit]
	i := 0
	for i < BucketSize && bucket.tophash_key[i] != hashSum {
		i++
	}
	if i == BucketSize {
		return 0, false
	}
	return bucket.value[i], true
}

type HashMapOption func(*HashMap)

func WithHashCRC8() HashMapOption {
	return func(hm *HashMap) {
		hm.hasher = func(key string) uint32 {
			table := crc8.MakeTable(crc8.CRC8_MAXIM)
			crc := crc8.Checksum([]byte(key), table)
			return uint32(crc)
		}
	}
}

func WithHashCRC16() HashMapOption {
	return func(hm *HashMap) {
		hm.hasher = func(key string) uint32 {
			table := crc16.MakeTable(crc16.IBM)
			crc := crc16.Checksum([]byte(key), table)
			return uint32(crc)
		}
	}
}

func WithHashCRC32() HashMapOption {
	return func(hm *HashMap) {
		hm.hasher = func(key string) uint32 {
			table := crc32.MakeTable(crc32.IEEE)
			crc := crc32.Checksum([]byte(key), table)
			return crc
		}
	}
}

func WithHashCRC64() HashMapOption {
	return func(hm *HashMap) {
		hm.hasher = func(key string) uint32 {
			table := crc64.MakeTable(crc64.ECMA)
			crc := crc64.Checksum([]byte(key), table)
			return uint32(crc)
		}
	}
}

func NewHashMap(sz uint32, options ...HashMapOption) *HashMap {
	countBuckets := uint8(math.Log2(float64(sz)))
	hm := &HashMap{
		sz:       sz,
		bucketsC: countBuckets,
		buckets:  make([]Buckets, sz),
	}
	for _, option := range options {
		option(hm)
	}
	return hm
}

func MeassureTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}

/*
func main() {
	m := NewHashMap(16, WithHashCRC64())
	since := MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC32())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC16())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)
	m = NewHashMap(16, WithHashCRC8())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)
}
*/
