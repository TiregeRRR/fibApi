# fibApi
Сервис реализует HTTP REST и GRPC, возвращает слайс с рядом Фибоначчи от x до y. Данные кэшируются при помощи Redis. 

## Установка 

На компьютере должны быть установлены go и docker 
> docker pull redis
> docker run --rm --name redis-test-instance -p 6379:6379 -d redis
> go mod download
> go mod tidy 

## Запуск 

### HTTP REST
> go run rest/main.go

### GRPC
> go run grpc/main.go