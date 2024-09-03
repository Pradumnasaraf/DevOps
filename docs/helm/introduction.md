---
sidebar_position: 1
title: Helm Introduction
---

Helm is package manager and a templating engine for Kubernetes. It allows you to define, install, and upgrade even the most complex Kubernetes applications. It's like apt, yum, or homebrew for Kubernetes. Primary use case is application deployment and environment management.    

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