# fibApi
Сервис реализует HTTP REST и GRPC, возвращает слайс с рядом Фибоначчи от x до y. Данные кэшируются при помощи Redis. 

## Установка 

На компьютере должны быть установлены go и docker 
``` sh
$ docker pull redis
$ docker run --rm --name redis-test-instance -p 6379:6379 -d redis
$ go mod download
$ go mod tidy
``` 

## Запуск 

### HTTP REST
``` sh
$ go run rest/main.go
```
### GRPC
``` sh
$ go run grpc/main.go
```

## Использование 

### HTTP REST
Обращение к api происходит через обычный GET-запрос. 
Пример: http://127.0.0.1:8000/api/v1/?x=5&y=10

## GRPC 
Обращение к api происходит через GRPC-клиент. Например: evans, Kreya. Proto-файл находится в fibApi/proto