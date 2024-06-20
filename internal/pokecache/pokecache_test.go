package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Second)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Second)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "Key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "",
			inputVal: []byte("valEmptyKey"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("%s not found", c.inputKey)
			continue
		}
		if string(actual) != string(c.inputVal) {
			t.Errorf(
				"%s does not match %s",
				string(actual),
				string(c.inputVal),
			)
			continue
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))
	time.Sleep(interval + time.Millisecond*2)

	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("%s should have been reaped", "key1")
	}
}
