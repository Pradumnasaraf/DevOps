## Kubernetes Learnings 

## Playground (environment to test out Kubernetes)

- [labs.play-with-k8s.com](https://labs.play-with-k8s.com/)
- [killercoda.com/playgrounds](https://killercoda.com/playgrounds)

## Resources 

- [Docker Mastery: with Kubernetes +Swarm from a Docker Captain](https://www.udemy.com/course/docker-mastery/) Udemy course.
- [BretFisher/udemy-docker-mastery](https://github.com/BretFisher/udemy-docker-mastery) GitHub repo.
- Kubernetes official [docs](https://kubernetes.io/docs/home/)

## Tools arround k8s

- [Validkube](https://validkube.com/) - Kubernetes Manifest file validator (check security, structure).
- [Lens](https://k8slens.dev/) - Lens is the only IDE you need to manage Kubernetes clusters. It's free and open source.

## Kubernetes Components - architecture

![Kube-component](https://user-images.githubusercontent.com/51878265/197317939-d7e8ecbb-912c-4223-b64a-1c46cbac255f.png)

<details>
  <summary>Simpler Image</summary>

<img width="872" alt="20200328170549" src="https://user-images.githubusercontent.com/51878265/197317783-ef595279-520d-4354-b995-96bff072485e.png">

</details>

## Master Node

- **API Server**: 
- **Etcd**: It stores the current state of the cluster. It's like a cluster brain.
- **Scheduler**: Decide which worker node will be best to deploy the next pods, after examining the resources and other paras. It does not schedule it.
- **Controller Manager**: Detect the current state of the cluster and keep the desired state of pods running. Follow requests when some things need to change/added to a worker node


## Worker Node

- **Kubelet**: It is the entry point to the Kubernetes cluster. Help us communicate with different objects in the Cluster
- **Kube Proxy**: Maintains network rules on the node, that allow network communication to your Pods from network sessions inside or outside of your cluster.
- **Container Runtime** - Like Docker, ContainerD, etc. Which runs the container

## Imperative Vs Declarative

- Imperative - When we give a command through CLI to run pod/deployment. For eg: `kubectl run nginx --image=nginx`

- Declarative - Creating deployment through YAML file. 

## Namespaces

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

## Pod Lifecycle

![Pod-Lifecycle](https://user-images.githubusercontent.com/51878265/197347032-cb45f52d-bfae-4ce4-838c-4c3ba9b10fa3.PNG)


## Configuration files

Generally, A K8s YAML config file contains 4 properties

```YAML
apiVersion: # Which version of the API you are using
kind: # What kind of object you are creating
metadata: # Data about the object
spec: # What you want the object to look like
```

#### Labels and selectors

Labels are for identification

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

Services are for internal communication of pods. It also helps give a pop static IP address. Contains routing rules. It also provide loadbalancing.

#### Types of Services

- **ClusterIP**: For inter communication of pods  

- **HeadLess**: It is a direct communication with a pod. No load blancing is required. So in this ClusterIp is none

```yaml
spec:
  clusterIP: None 
```

- **NodePort**: It allow external traffic to a fix port on each worker node.

> By default and for convenience, the `targetPort` is set to the same value as the `port` field.

```yaml
spec:
  type: NodePort
  ports:
    - port: 80
      targetPort: 80
      nodePort: 30007  # By default and for convenience, the Kubernetes control plane will allocate a port from a range (default: 30000-32767)
```

- **LoadBalancer**: Becomes accessile externally through cloud provider LoadBalancer.

General service file.

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

Multi-port service - In this we have to name the ports

```yaml
  ports:
    - name: mongogb
      protocol: TCP
      port: 27017 // Service Port
      targetPort: 27017 // Pod/Container Port
    - name: mongodb-exporter
      protocol: TCP
      port: 9216
      targetPort: 9216 
      
```

<p align="center" ><img width="626" alt="Screenshot 2022-11-28 at 2 15 57 PM" src="https://user-images.githubusercontent.com/51878265/204233963-c7bde7da-f631-49f3-b9db-1750ed55a37f.png"></p>

- Port forwarding

We can forward a port from a pod to our local machine

```bash
kubectl port-forward <pod-name> <localhost-port>:<pod-port>
```

or

Note: In this case pod port is same as localhost port

```bash
kubectl port-forward <pod-name> <localhost-port>
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

We can decode it by

```bash 
echo cHJhZHVtbmE | base64 --decode
```

## StatefulSet

- Any application that stores data to keep it state, like database. In this the name and endpoint stays same when the pods restarted.

## Secret and ConfigMap as volume

We can mount Config and Secret as a volume 

**depployment.yaml**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mosquitto-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mosquitto
  template:
    metadata:
      labels:
        app: mosquitto
    spec:
      containers:
        - name: mosquitto
          image: eclipse-mosquitto:1.6.2
          ports:
          - containerPort: 1883
          volumeMounts:
          - name: mosquitto-config # Volume name which is defined below and need to mounted
            mountPath: /mosquitto/config
            
      volumes: # List of volumes to mount into the container's filesystem.
        - name: mosquitto-config # This is the name of the volume
          configMap: #This is type of volume
            name: mosquitto-config-file
```

**config.yaml**
```Yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mosquitto-config-file
data:
  mosquitto.conf: |
    log_dest stdout
    log_type all
    log_timestamp_format %Y-%m-%dT%H:%M:%S
    listener 9001
```

### Volume VS using it as a ENV.

![Env vs Volume mount](https://user-images.githubusercontent.com/51878265/202616618-c536bbd6-e221-4df9-b57d-8969dc1504a8.png)


## Persistent Volume

First we create the Persistent Volume(PV) and then we claim it by creating Persistent Volume Claim (PVC). Then that claim is mounted to the pod. But when we use cloud provides we can claim the storage directly from the cloud provider.

- Note: Persistent volume is indepent of the namespace, but Persistent Volume Claim is bound to a specfic.

## Cluster Config file

All the Cluster info is stored in the file name `config` with the path:

```bash
~/.kube/config
```

## Networking

Container communication - The container inside a pod communicate via localhost shares the same networking namespace. To test it out, `Curl` the other conatiner by exec into the 1st container.

Steps

1) Create a deploymeny with the config file below 
 
```YAML
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
  labels:
    app: myapp
spec:
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80
      - name: sidecar
        image: curlimages/curl
        command: ["bin/sh"]
        args: ["-c", "echo Hello from the sidecar container! && sleep 3600"]
```

2) Get inside the `sidecar` conatiner in the pod myapp and access the terminal by:

```bash
kubectl exec -it <pod-name> -c sidecar -- /bin/sh
```

3) Curl the localhost with the respective port of other container.

```bash
curl localhost:80
```
## Updating Strategy

Updating means chnaging the image of the pod.

### Rolling Update

The pods are updated one by one, so the service is not down. But the new pods are created with the new image and then the old pods are deleted.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 5
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1 # 1 pod can be created above the desired number of pods. By default it is 25%
      maxUnavailable: 1 # 1 pod can be unavailable during the update. By default it is 25%
  selector:
    matchLabels:
      app: nginx-app
  template:
    metadata:
      labels:
        app: nginx-app
    spec:
      containers:
      - name: myapp
        image: nginx:1.23.2
        ports:
        - containerPort: 80
```


### Recreate

The pods are deleted and then new pods are created. So the service is down for a while.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 5
  strategy:
    type: Recreate # It will delete all the pods and then create new ones
  selector:
    matchLabels:
      app: nginx-app
  template:
    metadata:
      labels:
        app: nginx-app
    spec:
      containers:
      - name: myapp
        image: nginx:1.23.3
        ports:
        - containerPort: 80
```
