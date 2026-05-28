package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	//use above struct for the cacheEntry variable
	cacheEntry map[string]cacheEntry
	//unsure if this type is correct or not...
	cacheMutex sync.Mutex
}

// **** function to create a new Cache ****
func NewCache(interval time.Duration) Cache {
	//call readLoop() function!!!
	//create and return the new cache.
}

// might need to be a pointer?
func (c *Cache) Add(key string, val []byte) {
	//storing the passed in information into the map for the cache it points to...
	//adding by time.Now() at the time it executes
	c.cacheEntry[key] = cacheEntry{time.Now(), val}
}

// I think I got this one!!!
func (c *Cache) Get(key string) ([]byte, bool) {
	//bool should be true if found and false if not found
	//return the cache record if found and the value if found or not.
	result, exists := c.cacheEntry[key]
	if exists {
		//if the key exists then return the cache val ([]bytes) and return true.
		return result.val, true
	}
	//otherwise return an empty result and false not found
	return nil, false

}

func (c *Cache) readLoop(interval time.Duration) {
	//tips said use a goroutine (go) before a function using time.Ticker started by NewCache
	//for range ticker.C

	//creating a new ticker with the interval
	ticker := time.NewTicker(interval)
	//stopping at ticker implementation... step 4. Need to review when picking this shit up again....

}
