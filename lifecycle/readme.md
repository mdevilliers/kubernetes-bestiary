

```
cd lifecycle-service-image/src

CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o lifecycle .

cd ..

docker build -t mdevilliers/lifecycle-image:latest .
docker push mdevilliers/lifecycle-image:latest
```

```
kubectl create -f lifecycle-rc.yml 
kubectl delete -f lifecycle-rc.yml 
```

and view the logs...
