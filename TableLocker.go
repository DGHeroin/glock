package glock

import (
    "sync"
)

type TableLocker struct {
    table []*sync.RWMutex
}

func fnv32(key string) uint32 {
    prime32 := uint32(16777619)
    hash := uint32(2166136261)
    for i := 0; i < len(key); i++ {
        hash *= prime32
        hash ^= uint32(key[i])
    }
    return hash
}

func NewTableLocker(tableSize int) Locker {
    table := make([]*sync.RWMutex, tableSize)
    for i := 0; i < tableSize; i++ {
        table[i] = &sync.RWMutex{}
    }
    return &TableLocker{
        table: table,
    }
}

func (l *TableLocker) Lock(key string) {
    index := l.spread(fnv32(key))
    mu := l.table[index]
    mu.Lock()
}

func (l *TableLocker) UnLock(key string) {
    index := l.spread(fnv32(key))
    mu := l.table[index]
    mu.Unlock()
}
func (l *TableLocker) spread(hashCode uint32) uint32 {
    if l == nil {
        panic("dict is nil")
    }
    tableSize := uint32(len(l.table))
    return (tableSize - 1) & uint32(hashCode)
}
