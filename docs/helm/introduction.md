---
sidebar_position: 1
title: Helm Introduction
---

Helm is Package manager for Kubernetes.

### Using a Helm Chart

- Once we install the helm to the system, we have to add a chart repository, like prometheus, Ingress-nginx and [more](https://artifacthub.io/packages/search?kind=0)

Eg:

```bash
helm repo add <name> <repo url>
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
```

- To check list of charts we can install

```bash
helm search repo kubernetes-dashboard
```

- To install a chart 

```
helm install <name> kubernetes-dashboard/kubernetes-dashboard
```

### What's next?

[Learning Resources](./learning-resources.md) - Learn more about Helm with these resources.