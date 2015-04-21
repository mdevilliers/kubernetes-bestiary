
Build the docker file

```
cd web-image
docker build -t mdevilliers/test-go-web-image
````

Run in Docker

```
docker run -d -p 85:8080 mdevilliers/test-go-web-image
```

Push to repository

```
docker push mdevilliers/test-go-web-image
```

