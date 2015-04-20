web-image
---------

Will not work in kubernetes until the following are resolved :-

https://github.com/aspnet/KestrelHttpServer/issues/56
https://github.com/aspnet/Hosting/pull/129

```
kubectl create -f aspnext-webimage-pod.yaml
```

```
kubectl delete -f aspnext-webimage-pod.yaml
```