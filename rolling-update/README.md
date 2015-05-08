Build the docker images and push

```
cd web-image/v1
docker build -t mdevilliers/rolling-update-web-image .
docker tag mdevilliers/rolling-update-web-image:latest mdevilliers/rolling-update-web-image:v1
docker push mdevilliers/rolling-update-web-image:v1

cd web-image/v2
docker build -t mdevilliers/rolling-update-web-image .
docker tag mdevilliers/rolling-update-web-image:latest mdevilliers/rolling-update-web-image:v2
docker push mdevilliers/rolling-update-web-image:v2

````

Run in kubernentes

```
kubectl create -f rolling-update-controller-v1.json
kubectl create -f rolling-update-service.json

kubectl rollingupdate rolling-update-v1 -f rolling-update-controller-v2.json  --update-period=10s

```