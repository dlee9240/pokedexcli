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
	cache map[string]cacheEntry
	//unsure if this type is correct or not...
	//made it a pointer???
	cacheMutex *sync.Mutex
}

// **** function to create a new Cache ****
func NewCache(interval time.Duration) Cache {
	//call readLoop() function!!!
	//create and return the new cache.
	c := Cache{
		//creates a new cache map... set to empty values
		cache: make(map[string]cacheEntry),
		//need to understand this better
		//seems to create a new mutex and stores the address of the mutex.
		cacheMutex: &sync.Mutex{},
	}

	//use readloop function in a routine "go"
	go c.readLoop(interval)
	//return the cache after readall executes
	return c
}

// might need to be a pointer?
func (c *Cache) Add(key string, value []byte) {
	//storing the passed in information into the map for the cache it points to...
	//adding by time.Now() at the time it executes
	//lock the mutex when Add function is called...
	c.cacheMutex.Lock()
	//defer to unlock the mutex in the struct when done
	defer c.cacheMutex.Unlock()

	//update the cache via the struct format...
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

// I think I got this one!!!
func (c *Cache) Get(key string) ([]byte, bool) {
	//bool should be true if found and false if not found
	//return the cache record if found and the value if found or not.

	//forgot the c.mux.Lock
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()

	//I think this part was wrong.. Going to comment it out and use the one in the solution
	//result, exists := c.cache[key]
	//if exists {
	//if the key exists then return the cache val ([]bytes) and return true.
	//	return result.val, true
	//}
	//otherwise return an empty result and false not found
	//return nil, false

	//from solution at boot.dev
	val, ok := c.cache[key]

	//return the value of whatever is in cache[key] and the error...
	return val.val, ok

}

func (c *Cache) readLoop(interval time.Duration) {
	//tips said use a goroutine (go) before a function using time.Ticker started by NewCache
	//for range ticker.C

	//creating a new ticker with the interval
	ticker := time.NewTicker(interval)
	//stopping at ticker implementation... step 4. Need to review when picking this shit up again....
	//Reading from ticker.C blocks the goroutine until the specified duration has passed.
	for range ticker.C {
		//passes in the information in ticker.C (channel) and reaps...
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()

	// now = the current time Now().UTC() when the function was called
	// last = interval (5 minutes)

	//k = key
	//v = value
	//loop through the key value pairs in the cache map
	for k, v := range c.cache {
		//checks to see if the value in the cache map createdAt is before the current time -last //interval
		if v.createdAt.Before(now.Add(-last)) {
			//if so, delete the record from the cache with the key k variable from the for loop
			delete(c.cache, k)
		}
	}
}
