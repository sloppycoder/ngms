## Misc helpers for testing with K8S

Create services inside K8S for accessing database outside the cluster, modify hostname first, then run the following

```
kubectl apply -f db.yaml
```

### clean up disk usage in minikube VM
```
cat cleanup | minikube ssh
```

### To run the whole setup include Prometheus and Grafna with minikube
```
# make sure to edit db.yaml first, and populate some seed data, then
./all

# open browser to url displayed below to see locust UI
minikube service locust --url

# open browser to url displayed below to see Grafana
minikube service grafana --url

``` 
