# Messaging Practice - RabbitMQ 

## RabbitMQ

Running a standalone instance of rabbitMQ
    ```
        #Run stand alone instance 
        docker network create rabbits
        docker run -d --rm --net rabbits --hostname rabbit-1 --name rabbit-1 rabbitmq:3.8

        #Grabbing existing erlang cookie 
        docker exec -it rabbit-1 cat /var/lib/rabbitmq/.erlang.cookie

        # clean up
        docker rm -f rabbit-1
    ```

## Launch Publisher Application

1. Navigate to publisher directory 
    ` cd ~/publisher/ `
2. Build docker image 
    ` docker build . -t aimvector/rabbitmq-publisher:v1.0.0`
3. Spin up publisher container 
    ```
    docker run --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 80:80 aimvector/rabbitmq-publisher:v1.0.0 
    ```
## Launch Consumer Application
1. Navigate to publisher directory 
    ` cd ~/consumer/ `
2. Build docker image 
    ` docker build . -t aimvector/rabbitmq-consumer:v1.0.0`
3. Spin up publisher container 
    ```
    docker run --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 80:80 aimvector/rabbitmq-consumer:v1.0.0 
    ```
