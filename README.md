# user-manager

# Поднять postgres:
~~~
docker run --name=user-manager -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
~~~

# Поднять redis
~~~
docker run --name user-manager-redis -p '6377:6379' -d redis
~~~
