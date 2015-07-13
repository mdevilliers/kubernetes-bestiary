

```
cd content-provider-image
docker build -t mdevilliers/content-provider-sidecar-image .
cd ..
cd static-web-server
docker build -t mdevilliers/static-web-server-image .
```

```
docker push mdevilliers/content-provider-sidecar-image
docker push mdevilliers/static-web-server-image
```

```
kubectl create -f sidecar-rc.yaml
kubectl create -f sidecar-service.yaml
```

