package pokecache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co",
			val: []byte("kanto-data"),
		},
		{
			key: "https://pokeapi.co",
			val: []byte("johnto-data"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)

			actual, ok := cache.Get(c.key)
			if !ok {
				t.Fatalf("expected to find key %s", c.key)
			}
			if string(actual) != string(c.val) {
				t.Fatalf("expected value %s, got %s", string(c.val), string(actual))
			}
		})
	}
}

func TestReaperLoop(t *testing.T) {
	const interval = 10 * time.Millisecond
	const waitTime = interval + (5 * time.Millisecond)

	cache := NewCache(interval)
	key := "https://pokeapi.co"
	cache.Add(key, []byte("evict-me"))

	time.Sleep(waitTime)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("expected key %s to be reaped and delete", key)
	}
}

func TestReaperKeepsFresh(t *testing.T) {
	const interval = 50 * time.Millisecond
	cache := NewCache(interval)

	key := "fresh-key"
	cache.Add(key, []byte("keep-me"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected key %s to still exist in the cache", key)
	}
}

func TestConcurency(t *testing.T) {
	const interval = 10 * time.Second
	cache := NewCache(interval)

	var wg sync.WaitGroup
	workers := 50

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			val := fmt.Appendf(nil, "value-%d", id)
			cache.Add(key, val)
		}(i)
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			_, _ = cache.Get(key)
		}(i)
	}
	wg.Wait()
}
