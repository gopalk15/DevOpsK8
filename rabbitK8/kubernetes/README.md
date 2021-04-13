# RabbitMQ on Kubernetes 

Set up a namespace `rabbits` in your kubernetes cluster with the following command
    ` kubectl create namespace rabbits`

## Deployments

    ```
        kubectl apply -n rabbits -f .\kubernetes\rabbit-rbac.yaml
        kubectl apply -n rabbits -f .\kubernetes\rabbit-secret.yaml
    ```