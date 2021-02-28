package thread_locker

import (
	"sync"
)

// ThreadLocker is a class for counting threads.
//
// Members:
// 	none
//
type ThreadLocker struct {
	threads int
	sync.Mutex
}

// GetThreads gets the number of threads of the ThreadLocker.
//
// Parameters:
// 	none
//
// Returns:
// 	The number of threads of the ThreadLocker.
//
func (threadLock *ThreadLocker) GetThreads() int {
	threadLock.Lock()
	defer threadLock.Unlock()
	return threadLock.threads
}

// AddThread adds one to the the number of threads of the ThreadLocker.
//
// Parameters:
// 	none
//
// Returns:
// 	none
//
func (threadLock *ThreadLocker) AddThread() {
	threadLock.Lock()
	defer threadLock.Unlock()
	threadLock.threads += 1
}

// RemoveThread subtracts one to the the number of threads of the ThreadLocker.
//
// Parameters:
// 	none
//
// Returns:
// 	none
//
func (threadLock *ThreadLocker) RemoveThread() {
	threadLock.Lock()
	defer threadLock.Unlock()
	threadLock.threads -= 1
}

// Init initializes a ThreadLocker.
//
// Parameters:
//  none
//
// Returns:
// 	A ThreadLocker.
//
func Init() *ThreadLocker {
	return &ThreadLocker{threads: 0}
}
