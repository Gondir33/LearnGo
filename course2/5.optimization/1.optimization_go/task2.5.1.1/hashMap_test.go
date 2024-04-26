package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap_Set_Get(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// В этом блоке мы обрабатываем панику
			t.Errorf("Test panicked with: %v", r)
		}
	}()
	type args struct {
		key   string
		value interface{}
	}
	h := NewHashMap(16, WithHashCRC32())
	tests := []struct {
		name     string
		hashmap  HashMap
		args     args
		keycheck string
		want     interface{}
		want1    bool
	}{
		{
			name:    "1st",
			hashmap: *h,
			args: args{
				key:   "fire",
				value: 1,
			},
			keycheck: "fire",
			want:     1,
			want1:    true,
		}, {
			name:    "2nd",
			hashmap: *h,
			args: args{
				key:   "water",
				value: 2,
			},
			keycheck: "water",
			want:     2,
			want1:    true,
		}, {
			name:    "3nd",
			hashmap: *h,
			args: args{
				key:   "earth",
				value: 3,
			},
			keycheck: "olen",
			want:     nil,
			want1:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hashmap.Set(tt.args.key, tt.args.value)
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.hashmap.Get(tt.keycheck)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashMap.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HashMap.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_bucketMask(t *testing.T) {
	type args struct {
		b uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{{
		name: "bucket",
		args: args{3},
		want: 1<<3 - 1,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bucketMask(tt.args.b); got != tt.want {
				t.Errorf("bucketMask() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestNewHashMap(t *testing.T) {
	// Test HashMap with different hash functions
	hashFunctions := []HashMapOption{
		WithHashCRC32(),
		WithHashCRC64(),
	}

	for _, hashFunc := range hashFunctions {
		hm := NewHashMap(16, hashFunc)

		// Test Set and Get methods
		hm.Set("key", "value")
		value, ok := hm.Get("key")

		if !ok {
			t.Error("Expected value to be present, but it wasn't.")
		}

		if value != "value" {
			t.Errorf("Expected value to be 'value', but got '%v'", value)
		}
	}
}

func BenchmarkHashMapWithCRC32_Set(b *testing.B) {
	hm := NewHashMap(16, WithHashCRC32())
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < 100000; i++ {
		hm.Set("key"+strconv.Itoa(i), i)
	}
}

func BenchmarkHashMapWithCRC32_Get(b *testing.B) {
	hm := NewHashMap(16, WithHashCRC32())
	for i := 0; i < 100000; i++ {
		hm.Set("key"+strconv.Itoa(i), i)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < 100000; i++ {
		hm.Get("key" + strconv.Itoa(i))
	}
}

func BenchmarkHashMapWithCRC64_Set(b *testing.B) {
	hm := NewHashMap(16, WithHashCRC64())
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < 100000; i++ {
		hm.Set("key"+strconv.Itoa(i), i)
	}
}

func BenchmarkHashMapWithCRC64_Get(b *testing.B) {
	hm := NewHashMap(16, WithHashCRC64())
	for i := 0; i < 100000; i++ {
		hm.Set("key"+strconv.Itoa(i), i)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < 100000; i++ {
		hm.Get("key" + strconv.Itoa(i))
	}
}

func TestBucket_push(t *testing.T) {
	bucket := &Bucket{}
	bucket.push(&Data{hash_sum: 123, key: "sad", value: 1})
	bucket.push(&Data{hash_sum: 124, key: "zxc", value: 2})
	bucket.push(&Data{hash_sum: 125, key: "qwe", value: 3})
	bucket.push(&Data{hash_sum: 126, key: "jk;", value: 4})
	got, _ := bucket.GetByKey(123)
	assert.Equal(t, 1, got)
}
