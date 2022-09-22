# Rabbit MQ

## Build network and instance with docker
1. Create network
```docker network create rabbits```
2. Create instance
```docker run -d --rm --net rabbits -p 8080:15672 --hostname rabbit-1 --name rabbit-1 rabbitmq:3.10-management```

## Run instance
```docker exec -it rabbit-1 bash```</br>
Enable management plugin</br>
```rabbitmq-plugins enable rabbitmq_management```

## Build Publisher
```cd application/publisher```</br>
```docker build . -t demopat/rabbitmq-publisher```

## Run Publisher
```docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 80:80 demopat/rabbitmq-publisher```

## Build Consumer
```cd application/consumer```</br>
```docker build . -t demopat/rabbitmq-consumer```

## Run Consumer
```docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 81:81 demopat/rabbitmq-consumer```


## Note
hostname = Virtual host or logical group of entities
