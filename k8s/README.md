## Misc helpers for testing with K8S

Create services inside K8S for accessing database outside the cluster, modify hostname first, then run the following

```
kubectl apply -f db.yaml
```

### clean up disk usage in minikube VM
```
cat cleanup | minikube ssh
```

### use Prometheus operator to collect metrics from API server
```
# stay tuned


``` 


