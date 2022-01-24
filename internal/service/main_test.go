package service

import (
	"fbsTest/internal/repository"
	"os"
	"testing"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var testService *FibService
var brokenService *FibService

func TestMain(m *testing.M) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"redis": "localhost:6379",
		},
	})
	redisCache := cache.New(&cache.Options{
		Redis: ring,
	})
	repo := repository.NewFibRepository(redisCache, ring)
	testService = NewFibService(*repo)

	ringB := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"redis": "localhost:6378", //broken port
		},
	})
	redisCacheB := cache.New(&cache.Options{
		Redis: ringB,
	})
	repoB := repository.NewFibRepository(redisCacheB, ringB)
	brokenService = NewFibService(*repoB)
	os.Exit(m.Run())
}
