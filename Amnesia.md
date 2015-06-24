systemd
-------

read logs

https://www.digitalocean.com/community/tutorials/how-to-use-journalctl-to-view-and-manipulate-systemd-logs

```
journalctl -u {{unit name}}
```

docker 
------

attach to running container

```
docker exec -i -t {{id}} bash #by ID
```
or 

```
docker exec -i -t {{tag}} bash #by Name
```

debug a failed build

docker run --rm -it {{id}} bash -il

