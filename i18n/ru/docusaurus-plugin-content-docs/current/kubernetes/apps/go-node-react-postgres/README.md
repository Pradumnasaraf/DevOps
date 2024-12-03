## Multi Service App

This is a simple multi-service app that consists of a React Client, a Node API, and a Golang API. The app uses a Postgresql database for data storage.

### Prerequisites

- Kubernetes Cluster
- Helm v3
- kubectl

## Layout

The following shows the layout of this directory:

```markdown
├── api-golang
│   ├── Deployment.yaml
│   ├── IngressRoute.yaml
│   ├── Secret.yml
│   └── Service.yaml
├── api-node
│   ├── Deployment.yaml
│   ├── IngressRoute.yaml
│   ├── Secret.yaml
│   └── Service.yaml
├── client-react
│   ├── ConfigMap.yaml
│   ├── Deployment.yaml
│   ├── IngressRoute.yaml
│   └── Service.yaml
├── common
│   ├── Middleware.yaml
│   ├── Namespace.yaml
└── postgresql
    ├── Job.db-migrator.yaml
    └── Secret.db-password.yaml
```

### Setting up the APP

Make sure you are in `go-node-react-postgres` directory.

1. Create a Namespace

```bash
kubectl apply -f Namespace.yaml
```

switch to the namespace:

```bash
kubectl config set-context --current --namespace=demo-app
```

2. Install Postgresql Database using Helm

Add the bitnami repo:

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
```

Install the chart:

```bash
helm upgrade --install \
    -n postgres \
    postgres bitnami/postgresql \
    --set auth.postgresPassword=foobarbaz \
    --version 15.3.2 \
    --values ./postgresql/values.yaml \
    --create-namespace
```

3. Database Migration

First we need to create a secret for the database password:

```bash
kubectl apply -f postgresql/Secret.db-password.yaml
```

Then we can run the migration job:

```bash
kubectl apply -f postgresql/Job.db-migrator.yaml
```

5. Install traffic Ingress Controller

Add the traefik repo:

```bash
helm repo add traefik https://traefik.github.io/charts
```

Install the chart:

```bash
helm upgrade --install \
    -n traefik \
    traefik traefik/traefik \
    --version 20.8.0 \
    --create-namespace
```

6. Deploy a Middleware to Strop of the prefix from incoming requests

```bash
kubectl apply -f common/Middleware.yaml
```

7. Deploy the Golang API.

Before deploying, update the `IngressRoute.yaml` file to match the correct host.

```yaml
routes:
    - kind: Rule
      match: Host(`host/domain`) && PathPrefix(`/api/golang`)
```

Then deploy the API:

```bash
kubectl apply -f api-golang
```

8. Deploy the Node API. 

Before deploying, update the `IngressRoute.yaml` file to match the correct host.

```yaml
routes:
    - kind: Rule
      match: Host(`host/domain`) && PathPrefix(`/api/node`)
```

Then deploy the API:

```bash
kubectl apply -f api-node
```

9. Deploy the React Client.

Before deploying, update the `IngressRoute.yaml` file to match the correct host.

```yaml
routes:
    - kind: Rule
      match: Host(`host/domain`) # This will root with no prefix
```

Then deploy the client:

```bash
kubectl apply -f client-react
```

Now you should be able to access the app using the host/domain you have set in the IngressRoute files.

- For Node API, you can access it using `<host>/api/node`.
- For Golang API, you can access it using `<host>/api/golang`.
- For React Client, you can access it using `<host>`.
