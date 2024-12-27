---
title: K9s
sidebar_position: 5
description: K9s is a terminal-based UI to interact with your Kubernetes clusters. It provides a rich set of features to manage your Kubernetes resources.
tags: [Kubernetes, Tools]
keywords: [Kubernetes, Tools, K9s]
slug: /kubernetes/tools/k9s
---

K9s is a terminal-based UI to interact with your Kubernetes clusters. It provides a rich set of features to manage your Kubernetes resources. It can help you to navigate your Kubernetes clusters, view resource details, and perform actions on resources. Super useful for debugging and troubleshooting your Kubernetes applications.

## Resources

- [K9s GitHub Repository](https://github.com/derailed/k9s)
- [K9s Documentation](https://k9scli.io/)

## Installation 

You can install k9s on MacOS using Homebrew. 

```bash
brew install k9s
```

Once you have K9s installed, you can start it by running the following command

```bash
k9s
```
It will show the show the resources running in your Kubernetes cluster (for the current context and namespace).

![K9s](https://github.com/user-attachments/assets/4d3fefc3-34eb-407d-bae2-3299dd72f88c)

