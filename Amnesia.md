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

run with ability to detach

docker run -i -t  {{tag}}

Detach with Shift-P, Shift-Q

clean all untagged images

docker rmi -f $(docker images | grep "<none>" | awk "{print \$3}")


kubernentes
-----------

cadvisor port 4194

kube-proxy
----------

to view logs

journalctl --unit kube-proxy | head

to view args
 
cat /etc/sysconfig/kube-proxy

