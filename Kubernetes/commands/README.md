- `minikube start` - Creating a cluster
- `minikube status` - Running Status of cluster
- `minikube dashboard` - 
- `minikube docker-env` -
- `minikube SSH` -

> KUBECTL


- To check the version:

```bash
kubectl version
kubectl version --output=yaml
```

- To check info about the cluster

```bash
kubectl config view
```

- To run a Pod

```bash
kubectl run <pod name> --image <image name>
kubectl run myngix --image nginx
```

- To create a deployment

```bash
kubectl create deployment <name> --image <image name>
kubectl create deployment mynginx --image nginx
```

- To sacle the deployment (increase replicas)

```
kubectl scale deployment <deployment name> --replicas <no of replicas>
kubectl scale deployment mynginx --replicas 2
```

- To check all running services, pods, etc.

```bash
kubectl get all
```

- To check all the running services

```bash
kubectl get services
```

- To check all the running pods

```bash
kubectl get pods
```

```bash
// with extra details
kubectl get pods -o wide
```

- To check all the running node.

```bash
kubectl get nodes
```

- To check all the replicaset

```bash
kubectl get replicaset
```

- To check all the namespaces

```bash
kubectl get namespaces
```

- To delete the deployment

```bash
kubectl delete deployment <deployment-name>
```

- To delete the pods 

```bash
kubectl delete pod <pod-name>
```

- To get logs of a pod

```bash
kubectl logs <pod-name>
```

- Get a deep details/state chnages about a pod 

```bash
kube describe pod <pod -name>
```

- To watch the pods (watch refresh every few seconds)

```bash
kubectl get pods -w
```

- To check the cluster are avilable

```
kube config get-contexts
```

- `kubectl create deployment <deployment-name> --image=<image-name> ` - cteate a deplyment (pod inside it, you can't directly created pods)
    - `kubectl create deployment ngni-dep --image=ngnix` 
- `kubectl edit deployment ngnix` - Edit the config file.
- `kubectl exec -it <pod-name> -- bin/bash` - open the shell for that pod
- `kubectl get pod -o wide` - get more details about the pods

- `kubectl create namespace <insert-namespace-name-here>` - create a namespace
- `kubectl apply -f <config file name> --namespace=<namespace name>`
    - Eg: `kub apply -f mongo-configmap.yaml --namespace=my-namespace`
- `kub get configmap -n my-namespace`
    - Eg: `kub get configmap -n my-namespace`
    
