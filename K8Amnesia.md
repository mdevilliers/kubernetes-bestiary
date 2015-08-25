
Display everything accross all namespaces

```
kubectl get pods,rc,svc --all-namespaces
```

Exec in a pod

```
kubectl exec {POD NAME} -i -t -- bash -il
```