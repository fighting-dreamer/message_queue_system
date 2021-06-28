package service

import (
	"nipun.io/message_queue/logger"
	"sync"
)

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
		if tlm.Keeper[entity] == nil {
			tlm.Keeper[entity] = &sync.Mutex{}
			logger.Logger.Debug().Msgf("Created mutex for entity : %s", entity)
		}
		tlm.Keeper[entity].Lock() // who ever want to acquire a lock on something already locked, will have to wait.
		logger.Logger.Debug().Msgf("Acquired lock on mutex for entity : %s", entity)
		tlm.KeeperState[entity] = Locked
		mutex.Unlock()
	}
}

func (tlm *TransactionLockManager) ReleaseLock(entities []string) {
	for _, entity := range entities {
		mutex.Lock()
		if tlm.KeeperState[entity] == Locked {
			logger.Logger.Debug().Msgf("released lock on mutex for entity : %s", entity)
			tlm.Keeper[entity].Unlock()
			tlm.KeeperState[entity] = ""
		}
		mutex.Unlock()
	}
}
