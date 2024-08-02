package pokecache

import (
    "time"
    "log"
    "sync"
)

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

type cache struct {
    cacheMap map[string]cacheEntry
    interval time.Duration
    mu  sync.Mutex
}

// This function returns a pointer to a new cache that clears after the given duration
func NewCache(t time.Duration) *cache {
    cacheMap := make(map[string]cacheEntry)
    currentCache := cache{cacheMap, t, sync.Mutex{}}
    // Setup reap loop
    ticker := time.NewTicker(t)
    go func() {
        for {
            select {
                case <- ticker.C:
                    currentCache.ReapLoop()
            }
        }
    }()
    return &currentCache
}

// The add method adds the given value to the cache, returns an error
func (c *cache) Add(key string, val []byte) error {
    // The data structure looks like this key --> {createdat, value}
    // If the entry already exists, do nothing and return
    c.mu.Lock()
    defer c.mu.Unlock()
    _, ok := c.cacheMap[key]
    if ok {
        log.Print("Attempted to cache existing entry, ignoring")
        return nil
    }

    // Add a new entry to the map
    entry := cacheEntry{
        createdAt: time.Now(),
        val: val,
    }
    c.cacheMap[key] = entry

    log.Print("Cached entry")
    return nil
}

// The get method returns the corresponding value to the given key and a boolean value indication the success of the operation
func (c *cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    val, ok := c.cacheMap[key]

    if !ok {
        log.Printf("current key cannot be accessed: %v", key)
        return nil, ok
    }
    cachedval := val.val
    return cachedval, ok
}

// Reap Loop method clears the cache of entries that are older than the given duration, Note this method must be setup with time.Ticker
func (c *cache) ReapLoop() {
    duration := c.interval
    // lock the map
    c.mu.Lock()
    defer c.mu.Unlock()
    now := time.Now()
    for key, cacheEntry := range c.cacheMap {
        createdAt := cacheEntry.createdAt
        life := now.Sub(createdAt)
        if life >= duration {
            // delete the entry
            delete(c.cacheMap, key)
            log.Print("Clearing old cache")
        }
    }
}
