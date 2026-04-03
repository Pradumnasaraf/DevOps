---
sidebar_position: 1
title: GitOps Introduction
description: Learn about GitOps and its benefits.
tags: ["GitOps", "Kubernetes", "Continuous Delivery"]
keywords: ["GitOps", "Kubernetes", "Continuous Delivery"]
slug: "/gitops"
---

GitOps is a way of managing infrastructure and applications by using Git as the source of truth. In a Kubernetes setup, that usually means the desired cluster state lives in a Git repository, and an agent in the cluster applies or reconciles that state.

### Benefits of GitOps

- The history of changes is stored in Git.
- Rolling back to a previous state is easier.
- Changes can be reviewed before they are applied.
- Changes can be tested before they reach the cluster.
- Synchronization can be automated.

### GitOps Tools

- [Argo CD](../argocd/introduction.md) is a GitOps controller that can be used to deploy applications to a Kubernetes cluster.

### What's next?

- [Learning Resources](./learning-resources.md) - Learn more about GitOps with these resources.
- [Argo CD](../argocd/introduction.md) - Learn more about Argo CD.
