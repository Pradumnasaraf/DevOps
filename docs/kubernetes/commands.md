---
sidebar_position: 3
title: Kubernetes Commands
description: A collection of commands to help you with Kubernetes.
tags: ["Kubernetes", "Containerization", "DevOps"]
keywords: ["Kubernetes", "Containerization", "DevOps"]
slug: "/kubernetes/commands"
---

## kubectl Commands

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
kubectl run mynginx --image nginx
```

- To create a deployment

```bash
kubectl create deployment <name> --image <image name>
kubectl create deployment mynginx --image nginx
```

- To scale the deployment (increase replicas)

```
kubectl scale deployment <deployment name> --replicas <no of replicas>
kubectl scale deployment mynginx --replicas 2
```

- To check all running services, pods, etc.

```bash
kubectl get all
```

- To get details from a particular namespace

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

- To check all the running nodes

```bash
kubectl get nodes
```

- To check all ReplicaSets

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

- To open a shell in a specific container or inspect logs in a multi-container Pod

```bash
kubectl exec -it <pod-name> -c <container-name> -- <bash command>
kubectl exec -it multi-container -c nginx-container -- curl localhost
kubectl exec -it multi-container -c nginx-container -- sh
kubectl logs multi-container -c nginx-container
```

- To get inside the pod

```
kubectl exec -it <pod name> -- sh
kubectl exec -it nginx -- sh
```

- Get detailed state and event changes for a Pod

```bash
kubectl describe pod <pod-name>
```

- To watch the pods (watch refresh every few seconds)
```bash
kubectl get pods -w
```

- To check the available cluster contexts

```
kubectl config get-contexts
```

- We can create namespace by

```bash
kubectl create namespace <name>
kubectl create namespace dev
```

- To do a dry run and get the output as YAML

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

- To change the default or active namespace

```bash
kubectl config set-context --current --namespace=<namespace name>
```

- To get the details of a particular namespace

```bash 
kubectl get all -n <namespace name>
```

## Read next

- [Kubernetes Introduction](./introduction.md) - Review the concepts behind Pods, Deployments, and namespaces.
- [Learning Resources](./learning-resources.md) - Continue with courses, docs, and longer tutorials.
- [Playground](./playground.md) - Practice these commands in a browser-based cluster.
