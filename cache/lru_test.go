package cache

import "testing"

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := New(0, nil)
	lru.Add("key1", String("123456"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "123456" {
		t.Fatalf("cache hit key1=123456 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"
	c := len(k1 + k2 + v1 + v2)
	lru := New(int64(c), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("removeOldest key1 failed")
	}
}
