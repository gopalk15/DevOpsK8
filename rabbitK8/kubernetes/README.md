# RabbitMQ and Prometheus on Kubernetes 

## Deploying Minikube Cluser - Windows 10 Machine

1. Install Minikube 
    ` choco install minikube -y`
2. Launch cluster with 2 nodes 
    ` minikube start --kubernetes-version=v1.14.8 --extra-config=apiserver.authorization-mode=RBAC --nodes 2 `
3. Set up a namespace `rabbits` in your kubernetes cluster with the following command
    ` kubectl create namespace rabbits`

### Getting the NodePort using kubectl

The minikube VM is exposed to the host system via a host-only IP address, that can be obtained with the `minikube ip` command. Any services of type NodePort can be accessed over that IP address, on the NodePort.

To determine the NodePort for your service, you can use a `kubectl` command like this (note that nodePort begins with lowercase n in JSON output)

```
    kubectl get service <service-name> --output='jsonpath="{.spec.ports[0].nodePort}"
```

### LoadBalancer Access 

A LoadBalancer service is the standard way to expose a service to the internet. With this method, each service gets its own IP address

Services of type `LoadBalancer` can be exposed via the `minikube tunnel` command. It must be run in a separate terminal window to keep the LoadBalancer running. Ctrl-C in the terminal can be used to terminate the process at which time the network routes will be cleaned up.

`minikube tunnel` runs as a process, creating a network route on the host to the service CIDR of the cluster using the cluster’s IP address as a gateway. The tunnel command exposes the external IP directly to any program running on the host operating system.

### DNS Resolution for Services 

#### For "Normal" Services (Not Headless)

```
    my-svc.my-namespace.svc.cluster-domain.example
```

#### For Headless Services (ClusterIP: None)

```
    _my-port-name._my-port-protocol.my-svc.my-namespace.svc.cluster-domain.example

    #For pod backing headless service 

    auto-generated-name.my-svc.my-namespace.svc.cluster-domain.example
```


## Deployments

    ```
        kubectl apply -n rabbits -f .\RabbitCluster\
        kubectl apply -n rabbits -f .\Prometheus\
        kubectl apply -n rabbits -f .\Grafana\
        kubectl apply -n rabbits -f ..\messagingPractice\application\publisher\deployment.yaml
    ```

When using minikube use the following commands to expose pods to localhost.

1. Exposing RabbitMQ management dashboard on `http://localhost:8080' 

```
    kubectl -n rabbits port-forward rabbitmq-0 8080:15672
```

2. Expose Prometheus service using `minikube service prometheus-svc`

3. Expose Grafana service using `minikube service grafana-svc` 
    - Install the Offical RabbitMQ plugin
4. Expose Go publisher app using the following command
```
    kubectl -n rabbits port-forward rabbitm1-publisher<pod-id> 80:80
```
5. Use postman to publish messages to the queue 'http://localhost:80/publish/<message>