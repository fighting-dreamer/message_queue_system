package service

import "sync"

type TransactionLockManager struct {
	Keeper      map[string]*sync.Mutex
	KeeperState map[string]string
}

const (
	Locked = "LOCKED"
)

var mutex sync.Mutex // Global Mutex, all operation of locking and un-locking happen in a serial manner using this.

func (tlm *TransactionLockManager) AcquireLock(entities []string) {
	for _, entity := range entities {
		mutex.Lock()
		tlm.Keeper[entity].Lock() // who ever want to acquire a lock on something already locked, will have to wait.
		tlm.KeeperState[entity] = Locked
		mutex.Unlock()
	}
}

func (tlm *TransactionLockManager) ReleaseLock(entities []string) {
	for _, entity := range entities {
		mutex.Lock()
		if tlm.KeeperState[entity] == Locked {
			tlm.Keeper[entity].Unlock()
			tlm.KeeperState[entity] = ""
		}
		mutex.Unlock()
	}
}
