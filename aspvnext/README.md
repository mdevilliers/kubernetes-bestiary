web-image
---------

```
cd web-image
docker build -t mdevilliers/aspvnext-image .
docker run -p 5004:5004 mdevilliers/aspvnext-image

cd console-app-image
docker build -t mdevilliers/aspvnext-console-app .
docker run mdevilliers/aspvnext-console-app

```

```
docker push mdevilliers/aspvnext-image
docker push mdevilliers/aspvnext-console-app
```


```
kubectl create -f aspnext-web-demo-rc.yaml
kubectl create -f aspnext-web-demo-service.yaml

```

```
kubectl delete -f aspnext-web-demo-rc.yaml
kubectl delete -f aspnext-web-demo-service.yaml
```