package glock

import (
    "math/rand"
    "testing"
)

func TestNewKeyLocker(t *testing.T) {
    l := NewKeyLocker()
    key := "hello"
    l.Lock(key)
    l.UnLock(key)
}

func TestNewTableLocker(t *testing.T) {
    l := NewTableLocker(100)
    key := "hello"
    l.Lock(key)
    l.UnLock(key)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
var keys[] string

func init()  {
    for i := 0; i<1000*1000*10;i++ {
        keys = append(keys, randString(16))
    }
}

func TestBenchmarkKeyLocker(t *testing.T)  {
    l := NewKeyLocker()
    for _, key := range keys {
        l.Lock(key)
        l.UnLock(key)
    }
}

func TestBenchmarkTableLocker(t *testing.T)  {
    l := NewTableLocker(100)
    for _, key := range keys {
        l.Lock(key)
        l.UnLock(key)
    }
}