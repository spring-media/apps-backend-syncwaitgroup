
[![Build Status](https://travis-ci.com/spring-media/apps-backend-syncwaitgroup.svg?token=yZWuL9Kotp6i5ACnmzh6&branch=master)](https://travis-ci.com/spring-media/apps-backend-syncwaitgroup)

# apps-backend-syncwaitgroup
Go written wrapper for sync.WaitGroup to ease the usage of Goroutines.

## Summary
The main purpose of SyncWaitGroup is to execute and manage the synchronous functions provided as async goroutines; additionally it allows to dynamically add more functions - even by the function currently managed. 
For convenience SyncWaitGroup provides a default mutex.  

## Usage examples

### Simple usage example
    func myFunc() {
        // do something synchronously...
    }

    var syncWaitGroup = SyncWaitGroup{}
    syncWaitGroup.AddFunction(myFunc)
    syncWaitGroup.Wait()

### Advanced usage example
    func myFunc() {
        // do something synchronously...
    }

    type myRunnable struct {}

    func (runner *myRunnable) Run(syncWaitGroup *SyncWaitGroup) {
    	// do something synchronously, maybe synchronized via the Mutex provided...
        syncWaitGroup.Mutex.Lock()
        ...
        syncWaitGroup.Mutex.Unlock()

        // dynamically add more sync code to be executed as Goroutine within this managed Runnable
        syncWaitGroup.AddFunction(myFunc)
    }
