---
sidebar_position: 3
title: Docker Команды
description: Коллекция команд Docker для начала работы с Docker.
tags: ["Docker", "Containerization", "DevOps"]
keywords: ["Docker", "Containerization", "DevOps"]
slug: "/docker/commands"
---

### Docker Основы

- Проверить версию Docker

```
docker version
```

- Проверить все доступные образы

```bash
docker images
```

- Скачать/Загрузить образ из Docker реестра на локальную машину.

```bash
docker pull <image name>
//Пример: docker pull nginx
```

- Запустить контейнер (Сначала загрузит образ, если его нет в локальной системе)
  - ПРИМЕЧАНИЕ: Когда мы просто указываем имя образа, он загрузит последний, т.е. `nginx:latest`. Мы также можем указать версию `nginx:1.14`
  - Дополнительно мы можем использовать флаги
  
     - `--name <name> `- Дать имя контейнеру.
     - `-p <Host port:container port>`- Перенаправить порт.
     - `-d` - Запустить в фоновом режиме
     - `-it` - Для интерактивной среды
     - `-e` - Для переменной окружения

```bash
docker run <image-name>
//Пример: docker run nginx
```

- Мы также можем передать полный файл `.env`

```bash
--env-file <path-to-env-file>
Пример: --env-file ./.env
```

### Docker Контейнер

- Остановить запущенный контейнер

```bash
docker stop <container-ID/name>
``` 

- Возобновить остановленный контейнер

```bash
docker start <container-ID/name>
```

- Проверить запущенные процессы внутри контейнера.

```bash
docker top <container-name/id>
```

- Проверить статистику запущенного контейнера.

```bash
docker stats <container-name/id>
```

- Проверить конфигурацию и информацию о контейнере.

```bash
docker inspect <container-name/id>
//Пример: docker inspect mynginx
```

- Проверить все запущенные контейнеры.

```bash
docker ps
или
docker container ls
```

- Запустить интерактивную сессию и подключить терминал к контейнеру.

  - ПРИМЕЧАНИЕ: не каждый образ поддерживает `bash`, поэтому мы должны использовать `sh`

```
docker exec -it <container-ID/name> bash/sh
```

- Проверить, какие порты были открыты и перенаправлены

```bash
docker port <image-name>
```

- Проверить все контейнеры (остановленные, приостановленные и запущенные)

```bash
docker ps -a
```

- Проверить логи контейнера

```bash
docker logs <container-ID/name>
```

- Удалить все остановленные контейнеры

```bash
docker container prune -f
```

- Удалить все контейнеры (остановленные, приостановленные и запущенные)
```bash
docker rm -f $(docker ps -aq)
```
- Удалить все образы 
```bash
docker rmi -f $(docker images -q)
```

- Удалить неиспользуемые образы
```bash
docker image prune -all
```

- Автоматическая очистка при выходе контейнера

```bash
 docker container run —rm <image-name>
```

### Docker Сеть

- Проверить список доступных сетей.

```bash
docker network ls
```

- Проверить компоненты сети, например, какие контейнеры подключены к этой сети.

```bash
docker network inspect <network-name>
```

- Запустить контейнер в определенной сети/собственной созданной сети 

```
docker run --network <network-name> <image-name>
```

```
docker inspect --format "{{.NetworkSettings.IPAddress}}" <container-name>
```

### Docker Образы

- Удалить образ

```bash
docker rmi <image name> -f
```

- Удалить все образы сразу

```bash
docker rmi $(docker images -q)
```

- Проверить слои образа и другую информацию

```bash
docker inspect  <image-name/id>
```

- Проверить формирование слоев образа

```bash 
docker history <image-name/id>
```

- Создать наш собственный образ с существующим образом.

```
docker image tag <image-name:tag> <new-image-name:tag>
docker image tag nginx pradumna/nginx:hello
```

```
docker image tag ubuntu:18.04 pradumna/ubuntu:example
```

### Docker Том

- Создать bind mount
  - Помогает синхронизировать наши локальные файлы с помощью Docker контейнера.
  

- Для синхронизации изменений нашей локальной машины с помощью Docker тома (Bind mount)
    - `- v` используется для определения тома, также мы даем другой флаг `-v` для переопределения изменений, чтобы они не изменились в контейнере.

```bash
docker run -v <path-to-folder-on-local-machine>:<path-to-folder-on-container> -p <host-port>:<container-port> -d --name docker-node docker-node
docker
```

```bash
docker run -v <path-to-folder-on-local-machine>:<path-to-folder-on-container> -v <path-to-file/folder-on-container> -p <local-machine-port>:<container-port> -d --name docker-node docker-node
```
Чтобы сделать его только для чтения, чтобы когда вы добавляете некоторые файлы внутрь, контейнер не создавался на вашей локальной машине, используйте `-v port:port:ro`

- docker volume команда для монтирования docker socket в docker контейнер для доступа к docker daemon хоста для выполнения непрерывной интеграции в jenkins при использовании docker в качестве агента.. 
```bash
docker run -it -v /var/run/docker.sock:/var/run/docker.sock docker-image:version bin/bash
```

### Docker Compose

- Запустить docker compose файл.
  Примечание: По умолчанию он ищет файл с именем `docker-compose.yaml`, чтобы дать файл с другим именем, используйте команду `-f <file-name.yaml>`

```bash
docker compose up -d
```

```bash
docker compose down
```

- Пересобрать новый образ с новыми изменениями

```bash
docker compose up --build
```

- Переопределить существующий compose файл

```bash
docker compose -f docker-compose.yaml  -f docker-compose.dev.yaml
```

###  Docker Swarm и Сервисы

- Инициализировать swarm

```bash
docker swarm init
```

- Проверить все доступные узлы

```bash
docker node ls
```

- Добавить узел как менеджер

```bash
docker node update --role manager <node-name>
```

- Создать overlay сеть

```bash
docker network create -d overlay backend
```

- Создать сервис. Также мы можем добавить флаги для дальнейшей настройки.

    - `--name` - дать имя сервису
    - `--replicas` - определить, сколько запущенных экземпляров одного и того же образа.
    - `-p` - для перенаправления порта

```bash
docker service create -p 8080:80 --name vote --replicas 2 nginx
```

- Получить все задачи контейнеров, запущенные на разных узлах 

```bash
docker service ps <service-name/id>
```

> ОБНОВЛЕНИЕ СЕРВИСА

- Масштабировать сервис вверх (т.е. увеличить количество реплик)

```bash
docker service scale <service name>=<no to scale>
docker service scale mynginx=5
```

- Обновить образ в запущенном сервисе

```bash
docker service update --image <updated image> <service name>
docker service update --image mynginx:1.13.6  web
```

- Обновить порт

Мы не можем напрямую обновить порт. Мы должны добавить и удалить порты

```
docker service update --publish-rm 8080 --publish-add 808180 <service name>
docker service update --publish-rm 8080 --publish-add 808180 mynginx
```

### Docker Stack

- Развернуть stack файл

```bash
docker stack deploy -c <file-name.yaml> <stackname>
```

- Удалить запущенный stack

```
docker stack rm <stack name>
```

- Проверить список запущенных stacks

```
docker stack ls
```

**STACK -> SERVICES -> TASKS -> CONTAINERS**

- Проверить, какие сервисы запущены внутри stack

```
docker stack services <stack name>
```

- Проверить, какие задачи запущены внутри stack

```
docker stack ps <stack name> 
```

> Реестр

```
127.0.0.0:5000/v2/_catalog
```

### Советы и сокращения

- Запустить команду с созданием контейнера

```bash
doc run <image-name> <command>
// Пример: `doc run ubuntu:16.04 echo hey`
```


- Создание нашего собственного образа и контейнера.

```
Шаг 1 - создать Dockerfile
Шаг 2 - docker build -t myimage:1.0 <path-of-dockerfile> (-t для тега)
Шаг 3 - docker run myimage:1.0
```
