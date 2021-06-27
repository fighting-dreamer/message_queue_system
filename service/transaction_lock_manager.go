package service

type ITransactionLockManager interface {
	AcquireLock([]string)
	ReleaseLock([]string)
}
