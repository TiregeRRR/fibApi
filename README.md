# fibApi
Сервис реализует HTTP REST и GRPC, возвращает слайс с рядом Фибоначчи от x до y. Данные кэшируются при помощи Redis. 

## Требования

На компьютере должны быть установлены go, make и docker 

## Запуск 

Первый запуск необходимо осуществлять через
``` sh
$ make build-and-run-docker
```
Последующие можно через
``` sh
$ make run-docker
```
Также можно передать файл конфигурации 
``` sh
$ make run-docker ENVFILE="./config/someConf.env" 
```

## Использование 

### HTTP REST
Обращение к api происходит через POST-запрос, c переданными x, y в json-формате. 
Пример curl-запроса:
``` sh 
$ curl -XPOST -H "Content-type: application/json" -d '{
  "x": 0,
  "y": 1000
}' '172.18.0.2:8000/api/v1/fib'
```

## GRPC 
Обращение к api происходит через GRPC-клиент. Например: evans, Kreya. Proto-файл находится в fibApi/proto

## Получение ip
``` sh
$ docker container ls
$ docker container inspect <container name> | grep IPAddress
```

Были проблемы с подключением к контейнеру из под Windows 10. Помог ответ от Russel Wheeler https://coderedirect.com/questions/355610/ping-docker-container-from-another-machine-in-the-network. На Linux Mint все работает сразу.