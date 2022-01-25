## HTTP & GRPC API для подсчета числа Фибоначи

### Запуск
* docker-compose up
* P.S. В docker-compose используется network: host так как по другому он у меня linux mint не подключается контейнеру. Если будете менять настройки сети, не забудьте поменять MAIN_REDIS_ADDRESS в config/app.env

### Проверка
* по пути http://localhost:8080/swagger/index.html можно проверить через swagger