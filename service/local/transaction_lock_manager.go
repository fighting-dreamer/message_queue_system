package service

import (
	"nipun.io/message_queue/logger"
	"strings"
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
	mutex.Lock()
	entitiesString := strings.Join(entities, "-")
	if tlm.Keeper[entitiesString] == nil {
		tlm.Keeper[entitiesString] = &sync.Mutex{}
		logger.Logger.Debug().Msgf("Created mutex for entity : %s", entitiesString)
	}
	tlm.Keeper[entitiesString].Lock() // who ever want to acquire a lock on something already locked, will have to wait.
	logger.Logger.Debug().Msgf("Acquired lock on mutex for entity : %s", entitiesString)
	tlm.KeeperState[entitiesString] = Locked
	mutex.Unlock()
}

func (tlm *TransactionLockManager) ReleaseLock(entities []string) {
	mutex.Lock() // to deal with classic deadlock scenario
	entitiesString := strings.Join(entities, "-")

	if tlm.KeeperState[entitiesString] == Locked {
		logger.Logger.Debug().Msgf("released lock on mutex for entity : %s", entitiesString)
		tlm.Keeper[entitiesString].Unlock()
		tlm.KeeperState[entitiesString] = ""
	}
	mutex.Unlock()
}
