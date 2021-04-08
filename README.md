# DevOpsK8
Practice with k8, prometues, and RabbitMQ

## Setting up Postgresql Pod

### Create Postgres config maps resource 
`kubectl create -f configmap\\postgres-configmap.yaml`

### Persistent Storage Volume

To save the data, we will be using Persistent volumes and persistent volume claim resource within Kubernetes to store the data on persistent storages

- Create storage retaled deplyments 
    `kubectl create -f PersistentVolume\\postgres-storage.yaml`


### Deploy Postgresql

- the deployment yaml file found in `/deployment` creates a postgres pod and service 

    `kubectl create -f .\\deployment\\postgres-deployment.yaml`

