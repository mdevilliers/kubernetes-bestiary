POC to set up a clustered rabbitmq cluster using kubernetes building blocks.
Clustering the nodes is a manual step however this is envisioned to be completed with a
"sidecar" container e.g. an erlang node calling rabbitmqctl


Services
--------

service name             | ports    | selector         |description
-------------------------|----------|------------------|-----------------------------
rabbitmq-svc             | 5672     | type: rmq-node   | amqp port - applications connect to this
rabbitmq-node-one-svc    | 4369     | name: rmq-node-1 | epmd port for distrubuted erlang
                         | 25672    |                  | rpc port for distrubuted erlang
rabbitmq-node-one-svc-ui | 15672    | name: rmq-node-1 | port for the managment ui
rabbitmq-node-two-svc    | 4369     | name: rmq-node-2 | epmd port for distrubuted erlang
                         | 25672    |                  | rpc port for distrubuted erlang
rabbitmq-node-two-svc-ui | 15672    | name: rmq-node-2 | port for the managment ui                          


Replication Controllers
-----------------------

rc                   | labels
---------------------|----------------------------------
rabbitmq-node-one-rc | name: rmq-node-1,type: rmq-node
rabbitmq-node-two-rc | name: rmq-node-2,type: rmq-node


Example
-------


```
kubectl create -f .
kubectl get pods

NAME               READY     STATUS    RESTARTS   AGE
kube-dns-zcjoc     3/3       Running   0          6h
rmq-node-1-w57q0   1/1       Running   0          59s
rmq-node-2-0bdm1   1/1       Running   0          59s


kubectl exec rmq-node-1-w57q0 -i -t -- bash -il

root@rmq-node-1-svc:/# rabbitmqctl cluster_status

Cluster status of node 'rabbit@rmq-node-1-svc' ...
[{nodes,[{disc,['rabbit@rmq-node-1-svc']}]},
 {running_nodes,['rabbit@rmq-node-1-svc']},
 {cluster_name,<<"rabbit@rmq-node-1-svc.default.svc.cluster.local">>},
 {partitions,[]}]

root@rmq-node-1-svc:/# rabbitmqctl stop_app      

Stopping node 'rabbit@rmq-node-1-svc' ...

root@rmq-node-1-svc:/# rabbitmqctl join_cluster rabbit@rmq-node-2-svc

Clustering node 'rabbit@rmq-node-1-svc' with 'rabbit@rmq-node-2-svc' ...

root@rmq-node-1-svc:/# rabbitmqctl start_app                         

Starting node 'rabbit@rmq-node-1-svc' ...

root@rmq-node-1-svc:/# 

```