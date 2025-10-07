package pokecache

import (
	"sync"
	"time"
)

type  cacheEntry struct{
	createdAt time.Time
	val []byte
}
type Cache struct{
	entries map[string]cacheEntry
	mutex sync.Mutex
	interval time.Duration
}
func NewCache(interval time.Duration) *Cache{
	c:=&Cache{
		entries: make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}
func (c *Cache) Add(key string, val []byte){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key]= cacheEntry{createdAt: time.Now(), val: val}
	
}
func (c *Cache) Get(key string)([]byte, bool){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry,ok:=c.entries[key] 
	
	if !ok{
		return  nil, false
	} 
	if time.Now().After(entry.createdAt.Add(c.interval)) {
		return nil, false 
	}
	return entry.val, true
}
func (c *Cache) reapLoop(){
	ticker := time.NewTicker(c.interval)
	for range ticker.C{
		c.mutex.Lock()
		for key,value := range c.entries{
			if time.Now().After(value.createdAt.Add(c.interval)){
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}