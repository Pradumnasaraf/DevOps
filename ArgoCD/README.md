## ArgoCD

- [ArgoCD](https://argoproj.github.io/cd)
- [ArgoCD GitHub Repo](https://github.com/argoproj/argo-cd)

> _"One the founding principles of GitOps - everything stored in Git."_

Argo CD is a popular GitOps controller. It is used to deploy applications to Kubernetes clusters. It is also used to manage the configuration of the cluster itself. It can be changed too.

- For production we can use the [Autopilot](https://github.com/argoproj-labs/argocd-autopilot). For a tradational approach we can use the [Manifest directory](https://github.com/argoproj/argo-cd/tree/master/manifests) approach. Community [Helm](https://github.com/argoproj/argo-helm/tree/master/charts/argo-cd) charts are also available 

<p align="center"><img width="1014" alt="Screenshot 2022-11-29 at 11 44 57 PM" src="https://user-images.githubusercontent.com/51878265/204613004-e5dace25-7502-487d-acea-86d63c70cc2a.png"></p>


## Sync Policy - Features

<p align="center"><img width="820" alt="Screenshot 2022-11-30 at 4 56 41 PM" src="https://user-images.githubusercontent.com/51878265/204785046-93eb5b20-31b6-4f2b-b2a4-223f88f18718.png"></p>

- **Auto Sync** - ArgoCD will automatically sync the application with the Git repository. This is done by polling the Git repository at a specified interval.

- **Self Healing** - If the chnages are done directly to deployment/cluster, it will discard those chnages and keep the state as per the Git repository. For example, if someone chnange the replica count to two from 1 from the CLI, it will be changed back to 1.

- **Auto Pruning** - ArgoCD will automatically delete the resources that are not present in the Git repository. By default, it will not.

## Progressive Delivery

Progressive Delivery is the practice of deploying an application in a gradual manner allowing for minimum downtime and easy rollbacks. There are several forms of progressive delivery such as blue/green, canary, a/b and feature flags.

- **Blue/Green** - Deploy the new version of the application to a new environment and then switch the traffic to the new environment. This is the most common form of progressive delivery. It is also the most complex to implement. It requires a lot of infrastructure and a lot of testing.

- **Canary** - Deploy the new version of the application to a small subset of users. If the new version is working fine, then deploy it to the rest of the users. This is the most common form of progressive delivery. It is also the most complex to implement. It requires a lot of infrastructure and a lot of testing.