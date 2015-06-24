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
docker exec -i -t 665b4a1e17b6 bash #by ID
```
or 

```
docker exec -i -t loving_heisenberg bash #by Name
```


