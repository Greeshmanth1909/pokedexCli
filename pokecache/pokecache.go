package pokecache

import(
    "time"
    "log"
    "sync"
)

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

// Create new cache for the given time duration
func NewCache(t time.Duration) cache {
    var newCacheStruct cache
    newCacheStruct.duration = t
    return newCacheStruct
}

type cache struct {
    cacheEntries map[string]cacheEntry
    Add(key string, val []byte)
    Get(key string) ([]byte, bool)
    ReapLoop()
    duration time.Duration
    mu sync.Mutex
}

func (c cache) Add(key string, val []byte) {
    // if key exists, do nothing
    c.mu.Lock()
    defer c.mu.Unlock()
    _, ok := c.cacheEntries[key]

    if ok {
        log.Print("Cache exists")
        return
    }
    
    // Cache doesn't exist, create it
    c.cacheEntries[key] = {time.Now(), val}
    log.Print("Added cache")
}

// Return cache value, if exists
func (c cache) Get(key string) ([]byte, bool){
    c.mu.Lock()
    defer c.mu.Unlock()
    cacheStructVal, ok := c.cacheEntries[key]

    if ok {
        return cacheStructVal.val, ok
    }
    return nil, ok
}

// clear the cache of any entry older that duration once that time has elapsed
func (c cache) ReadLoop() {
    ticker := time.NewTicker(c.duration)

    // Start a go routine
    go func () {
        for {
            select {
                case <- t := ticker.C: {
                    clearCache(c)
                }
            }
        }

    }
    return
}

// Removes older entries from the map
func clearCache(c cache) {
    c.mu.Lock()
    defer c.mu.Unlock()
    cacheMap := c.CacheEntries
    currentTime := time.Now()
    duration := c.duration
    for _, cacheStructVal := range cacheMap {
        timeDifference = cacheStructVal.createdAt - currentTime
        if timeDifference >= duration {
            //Clear cache
            log.Print("deleting expired cache..")
            delete(cacheMap, cacheStructVal)
        }
    }
}
