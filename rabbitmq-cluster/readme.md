

```
docker build -t mdevilliers/rabbitmq-node-image .
docker push mdevilliers/rabbitmq-node-image
```


```
 kubectl create -f rabbitmq-cluster-controller-rc.yaml
 kubectl create -f rabbitmq-cluster-controller-service.yaml
 kubectl create -f rabbitmq-cluster-controller-service-admin.yaml
```