## Kubernetes Learnings 

## Playground (environment to test out Kubernetes)

- [labs.play-with-k8s.com](https://labs.play-with-k8s.com/)
- [killercoda.com/playgrounds](https://killercoda.com/playgrounds)

## Resources 

- [Docker Mastery: with Kubernetes +Swarm from a Docker Captain](https://www.udemy.com/course/docker-mastery/) Udemy course.
- [BretFisher/udemy-docker-mastery](https://github.com/BretFisher/udemy-docker-mastery) GitHub repo.
- Kubernetes official [docs](https://kubernetes.io/docs/home/)

## Concepts

## Kubernetes Components - architecture

![Kube-component](https://user-images.githubusercontent.com/51878265/197317939-d7e8ecbb-912c-4223-b64a-1c46cbac255f.png)

<details>
  <summary>Simpler Image</summary>

<img width="872" alt="20200328170549" src="https://user-images.githubusercontent.com/51878265/197317783-ef595279-520d-4354-b995-96bff072485e.png">

</details>

### Master Node

- **API Server**: 
- **Etcd**: It stores the current state of the cluster. It's like a cluster brain.
- **Scheduler**: Decide which worker node will be best to deploy the next pods, after examining the resources and other paras. It does not schedule it.
- **Controller Manager**: Detect the current state of the cluster and keep the desired state of pods running

Follow requests when some things need to change/added to a worker node

    `Conroller Manager`


### Worker Node

- **Kubelet**: It is the entry point to the Kubernetes cluster. Help us communicate with different objects in the Cluster
- **Kube Proxy**: Maintains network rules on the node, that allow network communication to your Pods from network sessions inside or outside of your cluster.
- **Container Runtime** - Like Docker, ContainerD, etc. Which runs the container

#### Imperative Vs Declarative

- Imperative - When we give a command through CLI to run pod/deployment. For eg: `kubectl run nginx --image=nginx`

- Declarative - Creating deployment through YAML file. 

#### Namespaces

- Isolated environment, we can group resources separately like a database. Also, great for running different versions of the app.

We can add namespace attribute in YAMl file to specify with one it belongs to


```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mongodb-configmap
  namespace: my-namespace
data:
  database_url: mongodb-service
```

We can create a namespace by

```
kubectl create namespace <name>
kubectl create namespace dev
```

#### Labels and selectors

Labels are for identification


#### Pod Lifecycle

![Pod-Lifecycle](https://user-images.githubusercontent.com/51878265/197347032-cb45f52d-bfae-4ce4-838c-4c3ba9b10fa3.PNG)


## Configuration files

Generally, A K8s YAML config file contains 4 properties

```YAML
apiVersion: 
kind:
metadata:
spec:
```

### Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mongo-deployment
  labels:
    app: mongodb
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mongodb
  template:
    metadata:
      labels:
        app: mongodb
    spec:
      containers:
        - name: mongodb
          image: mongo
          ports:
            - containerPort: 27017
          env:
            - name: MONGO_INITDB_ROOT_USERNAME
              valueFrom:
                secretKeyRef:
                  name: mongodb-secrets
                  key: mongo-root-username  
            - name: MONGO_INITDB_ROOT_PASSWORD
              valueFrom: 
                secretKeyRef:
                  name: mongodb-secrets
                  key: mongo-root-password
```

### Services

Services are for internal communication of pods. It also helps give a pop static IP address. Contains routing rules.



```yaml
apiVersion: v1
kind: Service
metadata:
  name: mongodb-service
spec:
  selector:
    app: mongodb //Deployment app label
  ports:
    - protocol: TCP
      port: 27017 // Service Port
      targetPort: 27017 // Pod/Container Port
```

### Ingress

It is use for an external trafic/request, which can be accessed by an URL instaed of `IP-PORT - 17.28.55.44.5:7800`. For that we need an ingress controller to make work of ingress.

![Ingress](https://user-images.githubusercontent.com/51878265/201585224-eca055af-eeb6-473c-bd96-33af9b5f6c55.png)

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: kubernetes-ingress
  namespace: kubernetes-dashboard
spec:
  rules:
  - host: example.com
    http:
      paths:
      - pathType: Prefix
        path: "/"
        backend:
          service:
            name: kubernetes-dashboard
            port: 
              number: 80

```

TLS

<img width="1512" alt="Screenshot 2022-11-14 at 1 17 55 PM" src="https://user-images.githubusercontent.com/51878265/201604299-264768c3-e5b1-48fa-9bc1-3762a3052006.png">



### ConfigMap

Use to store external configurations like database URLs. We put it in simple text format unlike [Secrets](#secrets)

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mongodb-configmap
data:
  database_url: mongodb-service
```

### Secrets

We use secrets to pass environment variables inside the pods.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mongodb-secrets
type: Opaque
data:
  mongo-root-username: cHJhZHVtbmE= //pradumna
  mongo-root-password: c2FyYWYxMjM= //saraf123
```

> Note: the secret value should be `base64` encoded, like `cHJhZHVtbmE` 

```bash
echo -n "value" | base64
```

### Cluster Config file

All the Cluster info is stored in the file name `config` with the path:

```bash
~/.kube/config
```


