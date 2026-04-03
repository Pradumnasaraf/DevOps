---
sidebar_position: 1
title: Prometheus Introduction
description: A collection of resources to learn more about Prometheus.
tags: ["Prometheus", "Monitoring", "Alerting", "Kubernetes"]
keywords: ["Prometheus", "Monitoring", "Alerting", "Kubernetes"]
slug: "/prometheus"
---

Prometheus is an open-source monitoring and alerting toolkit. It is widely used to collect metrics, query them over time, and trigger alerts when systems behave unexpectedly.

### Installation using Helm

One common way to install Prometheus in Kubernetes is with the `kube-prometheus-stack` Helm chart. It includes Prometheus along with related components such as Grafana, Alertmanager, Node Exporter, and kube-state-metrics.


Step 1: Create a Kubernetes Namespace

```bash
kubectl create namespace monitoring
```
Step 2: Install Prometheus Operator

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack -n monitoring
```
Step 3: Access through localhost using port forwarding

```bash
kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring
```

> The Grafana dashboard will be available at `http://localhost:3000`. The default username is `admin`. The initial password usually comes from the chart configuration or the generated Kubernetes secret, depending on the version you install.


### What's next?

- [Learning Resources](./learning-resources.md) - Learn more about Prometheus with these resources.
