package service

import "time"

type UserRepo interface {
}

type ErrorRepo interface {
}

type FileRepo interface {
}

type Cache interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	GetTTL(key string) (time.Duration, error)
}
