package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCaching(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://test1.com",
			val: []byte("The value for test 1 which will be converted into a slice of bytes"),
		},
		{
			key: "https://test2.com/test",
			val: []byte("The value for test 2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Tests case: %v\n", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			existingCache, isPresent := cache.Get(c.key)
			if !isPresent {
				t.Errorf("We expected to find a value for key: %v\n", c.key)
				return
			}
			if string(existingCache) != string(c.val) {
				t.Errorf("The value within cache is different from the cached value. Key: %v", c.key)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	duration := 5 * time.Millisecond
	waitTime := 10 * time.Millisecond
	testUrl := "https://test.com/clearcache"
	cache := NewCache(duration)
	cache.Add(testUrl, []byte("This should test if the cache is cleared"))

	_, isFound := cache.Get(testUrl)
	if !isFound {
		fmt.Errorf("There should have been a value in the cache for the '%v' URL", testUrl)
		return
	}
	time.Sleep(waitTime)
	_, isFound = cache.Get(testUrl)
	if isFound {
		fmt.Errorf("There shouldn't have been any value in the '%v' URL (after waitTime)", testUrl)
		return
	}
}

func TestConcurrentAdd(t *testing.T) {
	duration := 5 * time.Second
	cache := NewCache(duration)
	testCases := []struct {
		key  string
		val1 []byte
		val2 []byte
	}{
		{
			key:  "http://test.com/1",
			val1: []byte("test1 value1"),
			val2: []byte("test1 value2"),
		},
		{
			key:  "http://test.com/2",
			val1: []byte("test2 value1"),
			val2: []byte("test2 value2"),
		},
	}

	for _, c := range testCases {
		go cache.Add("http://test.com/1", c.val1)
		cache.Add("http://test.com/1", c.val2)
		fmt.Println("The execution went well")
		val, _ := cache.Get(c.key)
		fmt.Println(string(val))
	}
}
