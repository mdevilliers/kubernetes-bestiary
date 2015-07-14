

```
cd redis-proxy
docker build -t mdevilliers/redis-proxy-ambassador-image .
```


```
docker push mdevilliers/redis-proxy-ambassador-image
docker push mdevilliers/static-web-server-image
```