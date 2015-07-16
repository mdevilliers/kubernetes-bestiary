

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