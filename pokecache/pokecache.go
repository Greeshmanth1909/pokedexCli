package pokecache

import(
    "time"
)

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

// Create new cache for the given time duration
func NewCache(t time.Duration) {
    return
}

type cache struct {
    cacheEntries cacheEntry
    Add(key string, val []byte)
    Get(key string) ([]byte, bool)
    ReapLoop()
}
