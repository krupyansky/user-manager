# user-manager

Тестовое задание
1. Описать proto файл с сервисом из 3 методов: добавить пользователя, удалить пользователя, список пользователей
2. Реализовать gRPC сервис на основе proto файла на PHP (Symfony) или Go
3. Для хранения данных использовать PostgreSQL
4. На запрос получения списка пользователей данные будут кешироваться в Redis на минуту и браться из Redis
5. При добавлении пользователя делать лог в ClickHouse
6. Добавление логов в ClickHouse делать через очередь Kafka

# Поднять postgres:
~~~
docker run --name=user-manager -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
~~~

# Поднять redis
~~~
docker run --name user-manager-redis -p '6377:6379' -d redis
~~~
