# apps-backend-syncwaitgroup
Go written wrapper for sync.WaitGroup to ease the usage of Goroutines.

## summary
SyncWaitGroup extends sync.WaitGroup with a mutex provided.  
The main purpose of SyncWaitGroup is to execute and manage the synchronous functions provided as async goroutines; additionally it allows to dynamically add more functions - even by the function currently managed.

## usage example

    var syncWaitGroup = SyncWaitGroup{}
    syncWaitGroup.AddFunction(...) or syncWaitGroup.AddRunnable(...)
    syncWaitGroup.Wait()
