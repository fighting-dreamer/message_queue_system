package service

type ITransactionLockManager interface {
	AcquireLock([]string) map[string]string
	ReleaseLock([]string) map[string]string
}
