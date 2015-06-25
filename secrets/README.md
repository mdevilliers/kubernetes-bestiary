Build the docker images and push

```
cd web-image
docker build -t mdevilliers/secrets-web-image .
docker push mdevilliers/secrets-web-image

````

Run in kubernentes

```
kubectl create -f secret.yaml
kubectl create -f secret-pod-controller.yaml
kubectl create -f secret-pod-service.yaml

```