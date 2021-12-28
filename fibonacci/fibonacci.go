package fibonacci

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/TiregeRRR/fibApi/config"
	"github.com/go-redis/redis/v8"
)

var client = newPool()

var ctx = context.Background()

// getFibSlice возвращает слайс типа uint64, содержащий в себе необходимые числа ряда, http статус и ошибку
func GetFibSlice(start, end string) ([]uint64, int, error) {
	x, err := strconv.Atoi(start)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	y, err := strconv.Atoi(end)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	if x < 0 {
		return nil, http.StatusBadRequest, errors.New("invalid input: x < 0")
	} else if y < 0 {
		return nil, http.StatusBadRequest, errors.New("invalid input: y < 0")
	} else if x > y {
		return nil, http.StatusBadRequest, errors.New("invalid input: x > y")
	}
	fibSlice := make([]uint64, y-x+1)
	for i := x; i <= y; i++ {
		fibSlice[i-x], err = getFibElementFromCache(i)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}

	return fibSlice, http.StatusOK, nil
}

// newPool возвращает указатель на клиент redis'a и заносит первые два элемента в кэш
func newPool() *redis.Client {
	cfg := config.GetConfig()
	addr := cfg.GetString("redis_ip") + ":" + cfg.GetString("redis_port")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.GetString("redis_pass"),
		DB:       0,
	})
	if st := rdb.Ping(ctx); st.Err() != nil {
		log.Fatalln("redis connect error: " + st.Err().Error())
	}
	rdb.Set(ctx, "0", 0, 0)
	rdb.Set(ctx, "1", 1, 0)
	return rdb
}

// getFibElementFromCache возвращает элемент ряда Фибоначчи под индексом i
func getFibElementFromCache(i int) (uint64, error) {
	val, err := client.Get(ctx, strconv.Itoa(i)).Uint64() // Если элемента нет в кэше, то вызываем calculateFibElement и записываем
	if err == redis.Nil {                                 // полученное в кэш
		val = calculateFibElement(i)
		err1 := client.Set(ctx, strconv.Itoa(int(i)), val, 0).Err()
		if err1 != nil {
			return 0, err1
		}
	} else if err != nil {
		return 0, err
	}
	return val, nil
}

// calculateFibElement считает элемент по индексу
func calculateFibElement(index int) uint64 {
	el1, _ := getFibElementFromCache(index - 1)
	el2, _ := getFibElementFromCache(index - 2)
	return el1 + el2
}
