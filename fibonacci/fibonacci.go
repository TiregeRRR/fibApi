package fibonacci

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var client = newPool()

var ctx = context.Background()

// getFibSlice return slice of fibonacci
func GetFibSlice(start, end string) ([]uint64, error) {
	x, err := strconv.Atoi(start)
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(end)
	if err != nil {
		return nil, err
	}
	if x < 0 {
		return nil, errors.New("invalid input: x < 0")
	} else if y < 0 {
		return nil, errors.New("invalid input: y < 0")
	} else if x > y {
		return nil, errors.New("invalid input: x > y")
	}
	fibSlice := make([]uint64, y-x+1)
	for i := x; i <= y; i++ {
		fibSlice[i-x] = getFibElementFromCache(i)
	}

	return fibSlice, nil
}

func newPool() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rdb.Set(ctx, "0", 0, 0)
	rdb.Set(ctx, "1", 1, 0)
	return rdb
}

func getFibElementFromCache(i int) uint64 {
	val, err := client.Get(ctx, strconv.Itoa(i)).Uint64()
	if err == redis.Nil {
		val = calculateFibElement(i)
		err = client.Set(ctx, strconv.Itoa(int(i)), val, 0).Err()
		if err != nil {
			log.Println(err)
		}
	} else if err != nil {
		log.Println(err)
	}
	return val
}

func calculateFibElement(index int) uint64 {
	return getFibElementFromCache(index-1) + getFibElementFromCache(index-2)
}
