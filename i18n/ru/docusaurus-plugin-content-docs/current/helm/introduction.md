---
sidebar_position: 1
title: Helm Introduction
---

Helm — это менеджер пакетов и шаблонизатор для Kubernetes. Он позволяет вам определять, устанавливать и обновлять даже самые сложные приложения Kubernetes. Он подобен apt, yum или homebrew, но предназначен специально для Kubernetes. Основное предназначение — это развертывание приложений и управление средой.

## Helm Архитектура

Новая архитектура Helm 3 очень проста. В ней остались только клиентские компоненты. Серверные компоненты, такие как Tiller, были удалены в Helm 3. Теперь Helm 3 представляет собой клиентское приложение, которое взаимодействует с сервером API Kubernetes. Это упрощение делает использование Helm более безопасным и удобным, так как убирает необходимость в управлении дополнительными серверными компонентами.

### Using a Helm Chart

Чарт Helm — это набор файлов, который описывает связанный набор ресурсов Kubernetes. Один чарт может быть использован для развертывания чего-то простого, как, например, pod с memcached, или чего-то сложного, например, полного стека веб-приложения с HTTP-серверами, базами данных, кешами и так далее.

Сначала нужно добавить репозиторий в Helm.

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
```

- We can update the repository using the following command

```bash
helm repo update
```

- Then we can install the chart using the following command

```bash
helm install postgres bitnami/postgresql
```

Замените название_чарта на нужное вам имя или ключевое слово для поиска. Эта команда позволит вам просмотреть доступные чарты в добавленных репозиториях.

```bash
helm search repo bitnami
```

Замените название_чарта на имя чарта, для которого вы хотите просмотреть доступные версии. Эта команда покажет вам все версии чарта, доступные в добавленных репозиториях.

```bash
helm search repo bitnami/postgresql --versions
```

Мы также можем скачать чарт локально, чтобы ознакомиться с конфигурационными файлами. Для этого используйте следующую команду:

```bash
helm pull bitnami/postgresql --version 16.3.0
```

Мы можем задать значения для чарта, используя флаг --values (или его сокращенную версию -f), предоставив файл YAML с настраиваемыми значениями. Для этого используйте следующую команду:

```bash
helm install postgresql bitnami/postgresql --values values.yaml
```

Файл values.yaml может выглядеть следующим образом, в зависимости от ваших требований и конфигурации приложения:

```yaml
commonAnnotations:
  foo: bar
```

Мы можем обновить чарт, используя следующую команду:

Эта команда обновит релиз под названием my-release до новой версии чарта, используя указанный файл значений. Если вы не укажете флаг --values, Helm будет использовать значения по умолчанию, определенные в чарте.

```bash
helm upgrade --install postgresql bitnami/postgresql --values values.yaml --version=16.4.0 --namespace=my-namespace
```

Мы можем откатить график, используя следующую команду: postgresql - это название релиза, которое мы можем получить, используя helm list.

```bash
helm rollback postgresql
```

Чтобы проверить статус релиза Helm.

```bash
helm list
```

Чтобы проверить значения, предоставленные пользователем, можно использовать следующую команду:

```bash
helm get values postgresql
```

Мы также можем получить фактические манифесты, которые сгенерированы чартом Helm. Снова, postgresql — это имя релиза.

```bash
helm get manifest postgresql
```

Мы можем удалить чарт, используя следующую команду. Снова, postgresql — это имя релиза.

```bash
helm uninstall postgresql
```

### Creating our own Helm Chart

- We can create our own helm chart using the following command. It will generate a dir with all the boilerplate code.

```bash
helm create mychart
```

- We can install the chart using the following command. `mychart` is the chart name and `./mychart` is the path to the chart

```bash
helm install mychart ./mychart
```

Мы также можем применить чарт с файлом значений.

```bash
helm install mychart ./mychart --values values.yaml
```

## Helm Hooks

Helm hooks — это способ взаимодействия с жизненным циклом релиза. Они позволяют вам вмешиваться в определенные моменты жизненного цикла релиза. Helm поддерживает следующие hooks:

- `pre-install` - Runs before any templates are rendered.
- `post-install` - Runs after all templates have been rendered and resources have been created.

### What's next?

[Learning Resources](./learning-resources.md) - Learn more about Helm with these resources.
