---
sidebar_position: 1
title: ArgoCD Introduction
description: A deep dive into ArgoCD, a declarative, GitOps continuous delivery tool for Kubernetes.
tags: ["ArgoCD", "GitOps", "Kubernetes", "Continuous Delivery"]
keywords: ["ArgoCD", "GitOps", "Kubernetes", "Continuous Delivery"]
slug: "/argocd"
---

Argo CD is a GitOps controller for Kubernetes. It watches a Git repository, compares the desired state with the current cluster state, and helps keep them in sync. It is commonly used to deploy applications and manage cluster configuration declaratively.

> One of the main ideas in GitOps is simple: store the desired state in Git.

- For production, you can use tools such as [Autopilot](https://github.com/argoproj-labs/argocd-autopilot), install from the official [manifest directory](https://github.com/argoproj/argo-cd/tree/master/manifests), or use the community [Helm chart](https://github.com/argoproj/argo-helm/tree/master/charts/argo-cd).

<p align="center"><img width="1014" alt="Screenshot 2022-11-29 at 11 44 57 PM" src="https://user-images.githubusercontent.com/51878265/204613004-e5dace25-7502-487d-acea-86d63c70cc2a.png"></img></p>

### Installation

- Install it by applying the manifests. Reference: [Argo CD Installation](https://argoproj.github.io/argo-cd/getting_started/#1-install-argo-cd)

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

Default credentials:

- Username: admin
- Password: You can get it by running the following command

```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

## Sync Policy - Features

<p align="center"><img width="820" alt="Screenshot 2022-11-30 at 4 56 41 PM" src="https://user-images.githubusercontent.com/51878265/204785046-93eb5b20-31b6-4f2b-b2a4-223f88f18718.png"></img></p>

- **Auto Sync** - Argo CD automatically syncs the application with the Git repository.

- **Self Healing** - If someone changes resources directly in the cluster, Argo CD can revert them so the live state matches Git again.

- **Auto Pruning** - Argo CD can automatically remove resources that are no longer defined in Git.

## Progressive Delivery

Progressive Delivery is the practice of deploying an application in a gradual manner allowing for minimum downtime and easy rollbacks. There are several forms of progressive delivery such as blue/green, canary, a/b and feature flags.

- **Blue/Green** - Deploy the new version of the application to a new environment and then switch the traffic to the new environment. This is the most common form of progressive delivery. It is also the most complex to implement. It requires a lot of infrastructure and a lot of testing.

- **Canary** - Deploy the new version of the application to a small subset of users. If the new version is working fine, then deploy it to the rest of the users. This is the most common form of progressive delivery. It is also the most complex to implement. It requires a lot of infrastructure and a lot of testing.

- To keep the whole setup declarative, we can define Argo CD applications just like any other Kubernetes resource and store them in Git as well.

[Docs](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/) for reference.

Sample YAML:

```YAML
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: nginx-app # This is the name of the application
  namespace: argocd 
spec:
    destination:
        namespace: argocd # This is the namespace where the application will be deployed
        server: https://kubernetes.default.svc # This is the default server
    project: default
    source:
        path: ./manifest # This is the path where the manifest is present
        repoURL: https://github.com/sarafpradumna/argo-test # This is the repo where the manifest is present
        targetRevision: HEAD
    syncPolicy:
        automated:
            prune: true
            selfHeal: true
            allowEmpty: false #By default it is false
        syncOptions:
        - CreateNamespace=true # This will create namespace if not present
```

### What's next?

- [Learning Resources](./learning-resources.md) - Learn more about ArgoCD with these resources.
- [Other Resources](./other-resources.md) - Explore more about ArgoCD with these resources.
