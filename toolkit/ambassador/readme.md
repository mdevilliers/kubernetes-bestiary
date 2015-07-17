

```
cd redis-proxy
docker build -t mdevilliers/redis-proxy-ambassador-image .
cd ../web-app
docker build -t mdevilliers/web-app-ambassador-image .
```


```
docker push mdevilliers/redis-proxy-ambassador-image
docker push mdevilliers/web-app-ambassador-image

```

```
kubectl create -f sharded-redis-rc.yml
kubectl create -f ambassador-rc.yml
kubectl create -f ambassador-service.yml
```

```
kubectl describe services ambassador-service
```

The service is exported as a NodePort - using the information from 


```
kubectl describe services ambassador-service
```

you can then craft some web requests -

```
curl 'http://172.17.8.103:30565/api/store?key=hello&value=gthhfghfghgfhgf'
curl 'http://172.17.8.103:30565/'
```