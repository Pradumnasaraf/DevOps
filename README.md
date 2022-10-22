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

#### To see K8s config file

```sh
~/.kube/config
```
