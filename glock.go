package glock

type Locker interface {
    Lock(key string)
    UnLock(key string)
}
