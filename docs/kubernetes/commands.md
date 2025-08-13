---
sidebar_position: 3
title: Kubernetes Команды
description: Коллекция команд для работы с Kubernetes.
tags: ["Kubernetes", "Containerization", "DevOps"]
keywords: ["Kubernetes", "Containerization", "DevOps"]
slug: "/kubernetes/commands"
---

> KUBECTL КОМАНДЫ

- Проверить версию:

```bash
kubectl version
kubectl version --output=yaml
```

- Проверить информацию о кластере

```bash
kubectl config view
```

- Запустить Pod

```bash
kubectl run <pod name> --image <image name>
kubectl run myngix --image nginx
```

- Создать deployment

```bash
kubectl create deployment <name> --image <image name>
kubectl create deployment mynginx --image nginx
```

- Масштабировать deployment (увеличить реплики)

```
kubectl scale deployment <deployment name> --replicas <no of replicas>
kubectl scale deployment mynginx --replicas 2
```

- Проверить все запущенные сервисы, поды и т.д.

```bash
kubectl get all
```

- Получить детали из определенного namespace

```bash
kubectl get all -n <namespace name>
```

- Получить запущенные внутренние компоненты

```bash
kubectl get pods -A 
kubectl get pods -A -owide
```

- Проверить все запущенные сервисы

```bash
kubectl get services
```

- Проверить все запущенные поды

```bash
kubectl get pods
```

```bash
// с дополнительными деталями
kubectl get pods -o wide
```

- Проверить все запущенные узлы.

```bash
kubectl get nodes
```

- Проверить все replicaset

```bash
kubectl get replicaset
```

- Проверить все namespace

```bash
kubectl get namespaces
```

- Получить все API ресурсы

```bash
kubectl api-resources
```

- Удалить deployment

```bash
kubectl delete deployment <deployment-name>
```

- Удалить поды 

```bash
kubectl delete pod <pod-name>
```

- Удалить эвакуированные поды

```bash
kubectl delete pod --field-selector="status.phase==Failed"
```

- Получить логи пода

```bash
kubectl logs <pod-name>
```

- Проверить логи или sh/bash контейнера внутри пода. Если поды имеют несколько контейнеров, и мы хотим войти внутрь контейнера

```bash
kube exec -it <pod-name> -c <container-name> -- <bash command>
kube exec -it multi-container -c nginx-container -- curl localhost
kube exec -it multi-container -c nginx-container -- sh
kubectl logs multi-container -c nginx-container
```

- Войти внутрь пода

```
kubectl exec -it <pod name> -- sh
kubectl exec -it nginx -- sh
```

- Получить глубокие детали/изменения состояния о поде 

```bash
kube describe pod <pod -name>
```

- Наблюдать за подами (обновление каждые несколько секунд)
```bash
kubectl get pods -w
```

- Проверить доступные кластеры

```
kube config get-contexts
```

- Мы можем создать namespace с помощью

```bash
kubectl create namespace <name>
kubectl create namespace dev
```

- Сделать dry run и получить вывод в формате YAML

```bash
kubectl create namespace test-name --dry-run=client -oyaml
```

- Редактировать deployment (файл deployment)

```bash
kubectl edit deployment <deployment name>
```

- Удалить все поды
```
kubectl delete pods --all
```

- Применить к определенному namespace

```bash
kubectl apply -f <config file name> --namespace=<namespace name>
```

## Постоянный том

- Получить все PersistentVolume

```bash 
kubectl get pv
```

- Получить все PersistentVolumeClaim (привязанные к namespace)

```bash 
kubectl get pvc
```

- Изменить default/активный namespace

```bash
kubectl config set-context --current --namespace=<namespace name>
```

- Получить детали определенного namespace

```bash 
kubectl get all -n <namespace name>
```
