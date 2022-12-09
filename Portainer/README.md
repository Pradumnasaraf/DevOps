### Resources

- [Portainer](https://www.portainer.io/)
- [Portainer Documentation](https://documentation.portainer.io/)
- [Portainer GitHub](https://github.com/portainer/portainer)

### Installation

Via Helm

- Create a namespace for portainer

```bash
kubectl create namespace portainer
```

-  Add the portainer helm repository

```bash
helm repo add portainer https://portainer.github.io/k8s
```

- Install the portainer helm chart

```bash
helm update
helm install -n portainer portainer portainer/portainer 
```

> To install the business edition, use the following command

```bash
helm install --create-namespace -n portainer portainer portainer/portainer \
    --set enterpriseEdition.enabled=true \
    --set tls.force=true
```

- To access the portainer dashboard, run the following command

```bash
kubectl -n portainer port-forward svc/portainer 9000:9000
```

Dashboard will be available at http://localhost:9000. Preview

- For business edition, follow this [guide](https://install.portainer.io/).