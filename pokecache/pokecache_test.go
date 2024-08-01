package pokecache

import (
    "testing"
    "time"
    "bytes"
)

// TestNewCache tests the NewCache function
func TestNewCache(t *testing.T) {
    duration := 5 * time.Minute
    c := NewCache(duration)

    if c == nil {
        t.Error("Expected non-nil cache")
    }

    if c.interval != duration {
        t.Errorf("Expected duration %v, got %v", duration, c.interval)
    }
}

func TestAdd(t *testing.T) {
    duration := 5 * time.Minute
    c := NewCache(duration)
    expected := []byte("How are you")

    c.Add("hello", expected)
    actual, _ := c.Get("hello")
    if bytes.Equal(actual, expected) {
        t.Errorf("expected value to be :- %v, got %v", expected, actual)
    }
}

func TestReapLoop(t *testing.T) {
    duration := 1 * time.Minute
    c := NewCache(duration)

    // Add to cache
    input := []byte("there")
    c.Add("hello", input)
    
    // Wait for some time for the cache to clear
    time.Sleep(2 * time.Minute)

    // check cache, it must be empty
    val, ok := c.Get("hello")

    if ok {
        t.Errorf("The cache does not clear itself after given time %v, cache still contains %v", duration, val)
    }
}
