package sync

import (
	"sync"
)

// LockerDummy is a mock implementation of the sync.Locker interface.
type LockerDummy struct{}

// Lock is a dummy implementation of the sync.Locker.Lock method.
func (l LockerDummy) Lock() {}

// Unlock is a dummy implementation of the sync.Locker.Unlock method.
func (l LockerDummy) Unlock() {}

type RWLockerDummy struct{}

// RLock is a dummy implementation of the sync.RWMutex.RLock method.
func (l RWLockerDummy) RLock() {}

// TryRLock is a dummy implementation of the sync.RWMutex.TryRLock method.
func (l RWLockerDummy) TryRLock() bool { return true }

// RUnlock is a dummy implementation of the sync.RWMutex.RUnlock method.
func (l RWLockerDummy) RUnlock() {}

// Lock is a dummy implementation of the sync.RWMutex.Lock method.
func (l RWLockerDummy) Lock() {}

// TryLock is a dummy implementation of the sync.RWMutex.TryLock method.
func (l RWLockerDummy) TryLock() bool { return true }

// Unlock is a dummy implementation of the sync.RWMutex.Unlock method.
func (l RWLockerDummy) Unlock() {}

// RLocker returns a Locker interface that implements the sync.Locker.
func (l RWLockerDummy) RLocker() sync.Locker { return l }
