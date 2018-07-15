package core

import (
	"time"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

type cacheapp struct {
	codec *cache.Codec
}

func (this *cacheapp) connection(config map[string]interface{}) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			config["host"].(string): ":" + config["port"].(string),
		},
		DB: 0,
	})

	codec := &cache.Codec{
		Redis: ring,
	}

	this.codec = codec
}

func (this *cacheapp) Put(key string, content map[string]interface{}, expired time.Duration) {
	this.codec.Set(&cache.Item{
		Key:        key,
		Object:     content,
		Expiration: expired,
	})
}

func (this *cacheapp) Get(key string) map[string]interface{} {
	var data map[string]interface{}

	if err := this.codec.Get(key, &data); err == nil {
		return data
	} else {
		return map[string]interface{}{}
	}
}

func Cache(redisConfig map[string]interface{}) *cacheapp {
	var cache cacheapp

	cache.connection(redisConfig)

	return &cache
}
