
Build the docker file

```
cd web-image
docker build -t mdevilliers/test-go-web-image
````

Run in Docker

```
docker run -d -p 85:8080 mdevilliers/test-go-web-image .
```

Push to repository

```
docker push mdevilliers/test-go-web-image
```

Run in kubernentes

```
kubectl create -f cluster-wide-controller.yaml
kubectl create -f cluster-wide-service.yaml
```

Inspect

```
kubectl cluster-info

Kubernetes master is running at http://localhost:8080
cluster-wide-service is running at http://localhost:8080/api/v1beta1/proxy/services/monitoring-c-w-svc/ (note the trailing slash)

```

Version 16.2 Info - service really running at :-

http://172.17.8.101:8080/api/v1beta3/proxy/namespaces/default/services/monitoring-c-w-svc/
