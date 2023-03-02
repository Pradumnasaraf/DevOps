> KUBECTL COMMANDS

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

- To get the get details from the a particular namespace

```bash
kubectl get all -n <namespace name>
```

- To get the internal components running

```bash
kubectl get pods -A 
kubectl get pods -A -owide
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

- To get all the API resources

```
kubectl api-resources
```

- To delete the deployment

```bash
kubectl delete deployment <deployment-name>
```

- To delete the pods 

```bash
kubectl delete pod <pod-name>
```

- To delete evicted pods

```bash
kubectl delete pod --field-selector="status.phase==Failed"
```

- To get logs of a pod

```bash
kubectl logs <pod-name>
```

- To check logs or sh/bash of a container inside a pod. That if pods have multiple container an we have enter inside a container

```bash
kube exec -it <pod-name> -c <container-name> -- <bash command>
kube exec -it multi-container -c nginx-container -- curl localhost
kube exec -it multi-container -c nginx-container -- sh
kubectl logs multi-container -c nginx-container
```

- To get inside the pod

```
kubectl exec -it <pod name> -- sh
kubectl exec -it nginx -- sh
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

- We can create namespace by

```bash
kubectl create namespace <name>
kubectl create namespace dev
```

- To do a dry nun and get the output as Yaml

```bash
kubectl create namespace test-name --dry-run=client -oyaml
```

- To edit the deployment (deployment file)

```bash
kubectl edit deployment <deployment name>
```

- To delete all the pods
```
kubectl delete pods --all
```

- Apply to a particular namespace

```bash
kubectl apply -f <config file name> --namespace=<namespace name>
```

## Persistent Volume

- Get all the PersistentVolume

```bash 
kubectl get pv
```

- Get all the PersistentVolumeClaim (tied to a namespace)

```bash 
kubectl get pvc
```

- To chnage default/active namespace

```bash
kubectl config set-context --current --namespace=<namespace name>
```

- To get the details of a particular namespace

```bash 
kubectl get all -n <namespace name>
```
