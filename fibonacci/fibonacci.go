package fibonacci

import (
	"context"
	"errors"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var client = newPool()

var ctx = context.Background()

// getFibSlice возвращает слайс типа uint64, содержащий в себе необходимые числа ряда, http статус и ошибку
func GetFibSlice(start, end string) ([]string, int, error) {
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
	fibSlice := make([]string, y-x+1)
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
	addr := "redis:" + os.Getenv("redis_port")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("redis_pass"),
		DB:       0,
	})
	if st := rdb.Ping(ctx); st.Err() != nil {
		log.Println("redis connect error: " + st.Err().Error())
	}
	rdb.Set(ctx, "0", 0, 0)
	rdb.Set(ctx, "1", 1, 0)
	return rdb
}

// getFibElementFromCache возвращает элемент ряда Фибоначчи под индексом i
func getFibElementFromCache(i int) (string, error) {
	val := big.NewInt(0)
	s := client.Get(ctx, strconv.Itoa(i)).Val() // Если элемента нет в кэше, то вызываем calculateFibElement и записываем
	if s == "" {                                // полученное в кэш
		val = calculateFibElement(i)
		err1 := client.Set(ctx, strconv.Itoa(int(i)), val.String(), 0).Err()
		if err1 != nil {
			return "", err1
		}
	} else {
		val.SetString(s, 10)
	}
	return val.String(), nil
}

// calculateFibElement считает элемент по индексу
func calculateFibElement(index int) *big.Int {
	el1, el2 := big.NewInt(0), big.NewInt(0)
	s1, _ := getFibElementFromCache(index - 1)
	s2, _ := getFibElementFromCache(index - 2)
	el1.SetString(s1, 10)
	el2.SetString(s2, 10)
	return el1.Add(el1, el2)
}
