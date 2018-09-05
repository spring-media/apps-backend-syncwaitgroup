package syncx

import "sync"

// SyncWaitGroup extends sync.WaitGroup with a mutex provided.
// The main purpose of SyncWaitGroup is to execute and manage the synchronous functions provided
// as async goroutines; additionally it allows to dynamically add more functions - even by the
// function currently managed.
//
// Usage example:
//  var syncWaitGroup = SyncWaitGroup{}
//  syncWaitGroup.AddFunction(...) or syncWaitGroup.AddRunnable(...)
//  syncWaitGroup.Wait()
type SyncWaitGroup struct {
	sync.WaitGroup
	Mutex sync.Mutex
}

// AddRunnable adds a synchronous function to the SyncWaitGroup, executes it as an async goroutine
// and let the WaitGroup wait for it. Runnable may interact with SyncWaitGroup while beeing executed.
func (syncWaitGroup *SyncWaitGroup) AddRunnable(f Runnable) {
	syncWaitGroup.Add(1)
	go func(f Runnable) {
		f.Run(syncWaitGroup)
		defer syncWaitGroup.Done()
	}(f)
}

// AddFunction adds a generic synchronous function to the SyncWaitGroup, executes it as an async goroutine
// and let the WaitGroup wait for it.
func (syncWaitGroup *SyncWaitGroup) AddFunction(f func()) {
	var runnableFunction = runnableFunction{f}
	syncWaitGroup.AddRunnable(&runnableFunction)
}

// Runnable conforming objects are compatible to be managed by SyncWaitGroup.
// Run should provide synchronous code and may dynamically add more goroutines to the
// SyncWaitGroup passed.
type Runnable interface {
	Run(syncWaitGroup *SyncWaitGroup)
}

type runnableFunction struct {
	f func()
}

func (runner *runnableFunction) Run(syncWaitGroup *SyncWaitGroup) {
	runner.f()
}
