package component

import (
	"context"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/khairulharu/ewallet/domain"
)

func GetCacheConnection() domain.CacheRepository {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatalf("error when connect bigCache: %s", err.Error())
	}
	return cache
}
