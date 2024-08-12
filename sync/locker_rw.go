package sync

import "sync"

// RWLocker is an extracted interface for sync.RWMutex.
// See https://pkg.go.dev/sync#RWMutex for details.
type RWLocker interface {
	// RLock locks rw for reading.
	// See https://pkg.go.dev/sync#RWMutex.Lock for details.
	RLock()
	// TryRLock tries to lock rw for reading.
	// See https://pkg.go.dev/sync#RWMutex.TryLock for details.
	TryRLock() bool
	// RUnlock unlocks rw.
	// See https://pkg.go.dev/sync#RWMutex.Unlock for details.
	RUnlock()
	// Lock locks rw for writing.
	// See https://pkg.go.dev/sync#RWMutex.Lock for details.
	Lock()
	// TryLock tries to lock rw for writing.
	// See https://pkg.go.dev/sync#RWMutex.TryLock for details.
	TryLock() bool
	// Unlock unlocks rw.
	// See https://pkg.go.dev/sync#RWMutex.Unlock for details.
	Unlock()
	// RLocker returns a Locker interface that implements the Lock and Unlock methods by calling rw.RLock and rw.RUnlock.
	// See https://pkg.go.dev/sync#RWMutex.RLocker for details.
	RLocker() sync.Locker
}
