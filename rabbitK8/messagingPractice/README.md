# Messaging Practice - RabbitMQ 

## Deploying Minikube Cluser - Windows 10 Machine

1. Install Minikube 
    ` choco install minikube -y`
2. Launch cluster with 2 nodes 
    ` minikube start --kubernetes-version=v1.14.8 --extra-config=apiserver.authorization-mode=RBAC --nodes 2 `

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
    ```
    cd ~/publisher/ 
    ```
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
## Forming a Manual RabbitMQ cluster

Use the rabbits docker network created with the following command
    ` docker network create rabbits `

1. Spin up three rabbitmq nodes manually using the following commands
    ```
        docker run -d --rm --net rabbits --hostname rabbit-1 --name rabbit-1 -p 8081:15672 -e RABBITMQ_ERLANG_COOKIE=JKJPRRZCSBOUMHFPDZSK rabbitmq:3.8-management
        docker run -d --rm --net rabbits --hostname rabbit-2 --name rabbit-2 -p 8082:15672 -e RABBITMQ_ERLANG_COOKIE=JKJPRRZCSBOUMHFPDZSK rabbitmq:3.8-management
        docker run -d --rm --net rabbits --hostname rabbit-3 --name rabbit-3 -p 8083:15672 -e RABBITMQ_ERLANG_COOKIE=JKJPRRZCSBOUMHFPDZSK rabbitmq:3.8-management
    ```

2. Check clustering status of one of the nodes with the following command
    ` docker exec -t rabbit-1 rabbitmqctl cluster_status `

If any existing node need to join a cluster they will lose all the data they have stored already. Therefore, you will have to join the reset command on every node that will like to join the cluster.
*NOTE: All instances of rabbitmq that are clustered togeter must be spun up with the same erlang cookie*

3. Configure nodes rabbit@rabbit-2 and rabbit@rabbit-3 to cluster with rabbit@rabbit-1
    ```
    #join node 2 to cluster

    docker exec -it rabbit-2 rabbitmqctl stop_app
    docker exec -it rabbit-2 rabbitmqctl reset
    docker exec -it rabbit-2 rabbitmqctl join_cluster rabbit@rabbit-1
    docker exec -it rabbit-2 rabbitmqctl start_app
    docker exec -it rabbit-2 rabbitmqctl cluster_status

    #join node 3 to cluster

    docker exec -it rabbit-3 rabbitmqctl stop_app
    docker exec -it rabbit-3 rabbitmqctl reset
    docker exec -it rabbit-3 rabbitmqctl join_cluster rabbit@rabbit-1
    docker exec -it rabbit-3 rabbitmqctl start_app
    docker exec -it rabbit-3 rabbitmqctl cluster_status
    ```

## Connecting RabbitMQ instance to Prometheus and Grafana

1. Spin up an instance of prometheus and mount a config file 
 `docker run -d --rm --net rabbits --hostname prometheus --name prometheus -p 9090:9090 -v C:/Users/kaylin/git/DevOpsK8/Prometheus/prometheus.yaml:/etc/prometheus/prometheus.yml prom/prometheus --config.file=/etc/prometheus/prometheus.yml`

2. Run Grafana
    `docker run -d --rm --net rabbits --hostname grafana --name grafana -p 3000:3000 grafana/grafana`