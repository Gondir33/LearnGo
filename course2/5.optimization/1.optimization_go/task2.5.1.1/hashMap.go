package main

import (
	"hash/crc32"
	"hash/crc64"
	"math"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type Data struct {
	hash_sum uint32
	key      string
	value    interface{}
}

type Node struct {
	data *Data
	next *Node
}

type Bucket struct {
	head *Node
	curr *Node
}

func (b *Bucket) push(data *Data) {
	if b.head == nil {
		b.curr = &Node{data: data, next: nil}
		b.head = b.curr
		return
	}
	b.curr = b.head
	for b.curr.next != nil {
		b.curr = b.curr.next
	}
	b.curr.next = &Node{data: data, next: nil}
}
func (b *Bucket) GetByKey(hash_sum uint32) (interface{}, bool) {
	b.curr = b.head
	for b.curr != nil {
		if b.curr.data.hash_sum == hash_sum {
			return b.curr.data.value, true
		}
		b.curr = b.curr.next
	}
	return nil, false
}

type HashMap struct {
	sz       uint32
	bucketsC uint8 //log2(sz)
	buckets  []Bucket
	hasher   func(key string) uint32
}

func (hm *HashMap) Set(key string, value interface{}) {
	hashSum := hm.hasher(key)
	m := bucketMask(hm.bucketsC)
	highbit := hashSum & uint32(m)
	bucket := &hm.buckets[highbit]
	bucket.push(&Data{hash_sum: hashSum, key: key, value: value})
}

func bucketMask(b uint8) uint8 {
	return 1<<b - 1
}

func (hm *HashMap) Get(key string) (interface{}, bool) {
	hashSum := hm.hasher(key)
	m := bucketMask(hm.bucketsC)
	lowbit := hashSum & uint32(m)
	bucket := &hm.buckets[lowbit]
	return bucket.GetByKey(hashSum)
}

type HashMapOption func(*HashMap)

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
		buckets:  make([]Bucket, sz),
	}
	for _, option := range options {
		option(hm)
	}
	return hm
}

/*
func main() {
	m := NewHashMap(3, WithHashCRC64())

	m.Set("key1", "value1")
	m.Set("key2", "value2")
	m.Set("key3", "value3")
	m.Set("key4", "value4")
	m.Set("key5", "value5")

	if value, ok := m.Get("key2"); ok {
		fmt.Println("Key2:", value)
	} else {
		fmt.Println("Key2 not found")
	}

	if value, ok := m.Get("key1"); ok {
		fmt.Println("Key1:", value)
	} else {
		fmt.Println("Key1 not found")
	}

	if value, ok := m.Get("key3"); ok {
		fmt.Println("Key3:", value)
	} else {
		fmt.Println("Key3 not found")
	}
}
*/
