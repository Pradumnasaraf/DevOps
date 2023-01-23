## Prometheus

Prometheus is an open-source systems monitoring and alerting toolkit.

### Learning Resources

- [Prometheus Playlist - TechWorld with Nana](https://youtube.com/playlist?list=PLy7NrYWoggjxCF3av5JKwyG7FFF9eLeL4)

## Installation using Helm

We will install Prometheus Operator using Helm. It is collection of Promethus + Grafana + Alertmanager + Node Exporter + Kube State Metrics + Pushgateway + Blackbox Exporter


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
Step 3: Accessing though localhost by Port forwarding

```bash
kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring
```

> The Grafana dashboard will be available at http://localhost:3000. The default username and password are `admin` $ `prom-operator`.

