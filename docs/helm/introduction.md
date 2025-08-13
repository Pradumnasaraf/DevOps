---
sidebar_position: 1
title: Введение в Helm
description: Узнайте о Helm и его преимуществах.
tags: ["Helm", "Kubernetes", "Package Manager"]
keywords: ["Helm", "Kubernetes", "Package Manager"]
slug: "/helm"
---

Helm - это менеджер пакетов и движок шаблонов для Kubernetes. Он позволяет вам определять, устанавливать и обновлять даже самые сложные приложения Kubernetes. Это как apt, yum или homebrew для Kubernetes. Основной случай использования - развертывание приложений и управление средой. 

## Архитектура Helm

Новая архитектура Helm 3 очень проста. У нее есть только клиентские компоненты. Серверные компоненты, такие как Tiller, удалены в Helm 3. Helm 3 - это клиентское приложение, которое взаимодействует с API сервером Kubernetes.

### Использование Helm Chart

Для развертывания приложения с помощью Helm вам нужно создать Helm chart. Helm chart - это коллекция файлов, которые описывают связанный набор ресурсов Kubernetes. Один chart может использоваться для развертывания чего-то простого, как pod memcached, или чего-то сложного, как полный стек веб-приложения с HTTP серверами, базами данных, кэшами и так далее.

Сначала нам нужно добавить репозиторий в helm

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
```

- Мы можем обновить репозиторий, используя следующую команду

```bash
helm repo update
```

- Затем мы можем установить chart, используя следующую команду

```bash
helm install postgres bitnami/postgresql
```

мы можем искать доступные charts, используя следующую команду

```bash
helm search repo bitnami
```

Или искать доступные версии chart, используя следующую команду

```bash
helm search repo bitnami/postgresql --versions
```

Мы также можем скачать chart локально, чтобы посмотреть на файлы конфигурации. Флаг `--version` используется для указания версии chart, а не версии приложения

```bash
helm pull bitnami/postgresql --version 16.3.0
```

Мы можем дать значения chart, используя флаг values и передавая значения в yaml файле

```bash
helm install postgresql bitnami/postgresql --values values.yaml
```

Файл values.yaml должен выглядеть так

```yaml
commonAnnotations:
    foo: bar
```

Мы можем обновить chart, используя следующую команду

```bash
helm upgrade --install postgresql bitnami/postgresql --values values.yaml --version=16.4.0 --namespace=my-namespace
```

Мы можем откатить chart, используя следующую команду `postgresql` - это имя релиза, мы можем получить имя релиза, используя `helm list`.

```bash
helm rollback postgresql
```

Чтобы проверить статус helm релиза

```bash
helm list
```

Чтобы проверить значения, предоставленные пользователем. Снова `postgresql` - это имя релиза

```bash
helm get values postgresql
```

Мы также можем получить фактические manifest файлы, которые генерируются helm chart. Снова `postgresql` - это имя релиза

```bash
helm get manifest postgresql
```

Мы можем удалить chart, используя следующую команду. Снова `postgresql` - это имя релиза

### Создание собственного Helm Chart

- Мы можем создать собственный helm chart, используя следующую команду. Он сгенерирует директорию со всем шаблонным кодом.

```bash
helm create mychart
```

- Мы можем установить chart, используя следующую команду. `mychart` - это имя chart, а `./mychart` - это путь к chart

```bash
helm install mychart ./mychart
```

Мы также можем применить с файлом значений

```bash
helm install mychart ./mychart --values values.yaml
```

## Helm Hooks

Helm hooks - это способ взаимодействия с жизненным циклом выпуска. Они позволяют вам вмешиваться в определенные моменты жизненного цикла выпуска. Helm поддерживает следующие хуки:

- `pre-install` - Запускается перед рендерингом шаблонов.
- `post-install` - Запускается после рендеринга шаблонов и создания ресурсов.

### Что дальше?

[Learning Resources](./learning-resources.md) - Узнайте больше о Helm с этими ресурсами.