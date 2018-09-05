package syncx

import (
	"testing"
	"time"
)

const sleepTime = 2 * time.Second

var mockFunctionRunned = false

func mockFunction() {
	time.Sleep(sleepTime)
	mockFunctionRunned = true
}

type mockRunnable struct {
	runned      bool
	addFunction bool
}

func (mock *mockRunnable) Run(syncWaitGroup *SyncWaitGroup) {
	time.Sleep(sleepTime)
	if mock.addFunction {
		syncWaitGroup.AddFunction(mockFunction)
	}
	mock.runned = true
}

func TestAddFunction(t *testing.T) {
	mockFunctionRunned = false

	var waitGroup = SyncWaitGroup{}
	waitGroup.AddFunction(mockFunction)
	waitGroup.Wait()

	if !mockFunctionRunned {
		t.Errorf("Added function was not executed.")
	}
}

func TestAddRunnable(t *testing.T) {
	var mock = mockRunnable{}
	mock.runned = false
	mock.addFunction = false

	var waitGroup = SyncWaitGroup{}
	waitGroup.AddRunnable(&mock)
	waitGroup.Wait()

	if !mock.runned {
		t.Errorf("Added runnable was not executed.")
	}
}

func TestDynamicAddedFunction(t *testing.T) {
	var mock = mockRunnable{}
	mock.runned = false
	mock.addFunction = true
	mockFunctionRunned = false

	var waitGroup = SyncWaitGroup{}
	waitGroup.AddRunnable(&mock)
	waitGroup.Wait()

	if !mock.runned {
		t.Errorf("Added runnable was not executed.")
	}
	if !mockFunctionRunned {
		t.Errorf("Dynamically added function was not executed.")
	}
}
