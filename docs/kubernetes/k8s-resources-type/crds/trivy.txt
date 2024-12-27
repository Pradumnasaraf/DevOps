### Trivy

Trivy helps to detect vulnerabilities in your application's dependencies. It is a simple and comprehensive vulnerability scanner for containers and other artifacts.

#### Installation

To install Trivy, you can use the following Helm chart:

```bash
helm repo add aqua https://aquasecurity.github.io/helm-charts/
helm repo update
helm install trivy-operator aqua/trivy-operator --namespace trivy-system --create-namespace --version 0.23.2
```

#### Uninstallation

To uninstall Trivy, you can use the following Helm command:

```bash
helm uninstall trivy-operator --namespace trivy-system
```
 