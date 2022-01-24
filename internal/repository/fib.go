package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type FibRepository struct {
	redisCache *cache.Cache
	redisRing  *redis.Ring
	key        string
}

func NewFibRepository(redisCache *cache.Cache, redisRing *redis.Ring) *FibRepository {
	return &FibRepository{
		redisCache: redisCache,
		redisRing:  redisRing,
		key:        "Fib",
	}
}

func (f *FibRepository) GetCachedFibSeq(ctx context.Context, last int) ([]int64, error) {
	res, err := f.checkKeys(ctx, last)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FibRepository) SetNewFibSeq(ctx context.Context, fibSeq []int64, last int) error {
	item := cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf("%s_%d", f.key, last),
		Value: fibSeq,
		TTL:   0,
	}

	err := f.redisCache.Set(&item)
	if err != nil {
		return err
	}
	return nil
}

func (f *FibRepository) checkKeys(ctx context.Context, last int) ([]int64, error) {
	var res []int64

	err := f.redisCache.Get(ctx, fmt.Sprintf("%s_%d", f.key, last), &res)
	if err == nil {
		return res, nil
	}

	iter := f.redisRing.Scan(ctx, 0, f.key+"_*", 3).Iterator()
	for iter.Next(ctx) {
		key := strings.Split(iter.Val(), "_")
		keyNum, err := strconv.Atoi(key[1])
		if err != nil {
			return nil, err
		}
		if keyNum > last {
			err = f.redisCache.Get(ctx, fmt.Sprintf("%s_%d", f.key, keyNum), &res)
			if err == nil {
				return res, nil
			}
		}
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("not found keys")
}
