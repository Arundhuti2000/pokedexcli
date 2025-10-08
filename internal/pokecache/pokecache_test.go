package pokecache

import (
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
    // Create a new cache with a 5 second interval
    cache := NewCache(5 * time.Second)
    
    // Add some data to the cache
    key := "test-key"
    value := []byte("test data")
    cache.Add(key, value)
    
    // Try to get the data back
    retrievedValue, ok := cache.Get(key)
    
    // Check if we got the data back
    if !ok {
        t.Errorf("Expected to find key %s in cache", key)
    }
    
    // Check if the value matches
    if string(retrievedValue) != string(value) {
        t.Errorf("Expected %s, got %s", value, retrievedValue)
    }
}

func TestCacheExpiration(t *testing.T) {
    // Create a cache with a very short interval
    cache := NewCache(10 * time.Millisecond)
    
    // Add some data
    key := "test-key"
    value := []byte("test data")
    cache.Add(key, value)
    
    // Wait for the data to expire
    time.Sleep(15 * time.Millisecond)
    
    // Try to get the expired data
    _, ok := cache.Get(key)
    
    // Should return false because the data expired
    if ok {
        t.Error("Expected cache entry to be expired, but it was still found")
    }
}