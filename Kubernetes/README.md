
## Kubernetes Learnings

This repo contains all my learning realted to **Kubernetes**. It will consist resources, code, commands, and configuration files.

#### Playground (Envirnoment to test out Kubernetes)

- [labs.play-with-k8s.com](https://labs.play-with-k8s.com/)
- [killercoda.com/playgrounds](https://killercoda.com/playgrounds)

#### Resources 

- [Docker Mastery: with Kubernetes +Swarm from a Docker Captain](https://www.udemy.com/course/docker-mastery/) Udemy course.
- [BretFisher/udemy-docker-mastery](https://github.com/BretFisher/udemy-docker-mastery) GitHub repo.
- Kubernetes official [docs](https://kubernetes.io/docs/home/)

## Concepts

#### Kubernetes Components - architecture

![Kube-component](https://user-images.githubusercontent.com/51878265/197317939-d7e8ecbb-912c-4223-b64a-1c46cbac255f.png)

<details>
  <summary>Simpler Image</summary>

<img width="872" alt="20200328170549" src="https://user-images.githubusercontent.com/51878265/197317783-ef595279-520d-4354-b995-96bff072485e.png">

</details>

#### Imperative Vs Declarative

- Imperative - When we give command though CLI to run pod/deplyment. For eg: `kubectl run nginx --image=nginx`

- Declrative - Creating deployment though YAML file. 

#### Namespaces

- Isolated environment to the tems. In this we can group resouces seprately like database. Also, great for running different versions of the app.

We can create namespace by

```
kubectl create namespace <name>
kubectl create namespace dev
```

#### Lables and selectors

Lables are for identification


#### Pod Lifecycle

![Pod-Lifecycle](https://user-images.githubusercontent.com/51878265/197347032-cb45f52d-bfae-4ce4-838c-4c3ba9b10fa3.PNG)


#### To see K8s config file

```sh
~/.kube/config
```

- Generally A K8s YAML config file contain 4 properties

```YAML
apiVersion: 
kind:
metadata:
spec:
```
