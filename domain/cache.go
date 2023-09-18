package domain

type CacheRepository interface {
	Get(key string) ([]byte, error)
	Set(key string, entry []byte) error
}
