---
sidebar_position: 1
title: Docker Введение
---

### Обзор Docker

Docker — это платформа с открытым исходным кодом, разработанная для упрощения разработки, развертывания и управления приложениями в контейнеризованной среде, также известной как контейнеры Docker. Контейнеры Docker представляют собой легковесный и портативный способ упаковки и запуска приложений, позволяя разработчикам упаковывать свои приложения со всеми необходимыми зависимостями и конфигурациями в одном пакете, который можно легко перемещать между любыми средами. Контейнеры Docker — это просто работающие экземпляры образа Docker.




### Почему стоит использования Docker

- Docker портативен, что означает, что одно и то же приложение можно легко запускать на разных машинах без каких-либо модификаций, что упрощает перемещение приложений между средами разработки, тестирования и производства.

- Контейнеры Docker изолированы по своей природе, что означает, что каждый контейнер работает в своем собственном изолированном окружении с собственной файловой системой, сетевым протоколом и пространством процессов, обеспечивая уровень безопасности и изоляции, который невозможен с традиционными технологиями виртуализации. Это решает проблему конфликтов с другими приложениями или зависимостями.

- Контейнеры Docker легко масштабируемы, что означает, что можно легко увеличить количество контейнеров, запускающих приложения, добавляя горизонтально больше контейнеров по мере роста спроса.

- Контейнеры Docker эффективны, что означает, что они легковесны и потребляют меньше ресурсов, позволяя запускать больше контейнеров на одном и том же аппаратном обеспечении.

### Образы Docker

- Образы состоят из бинарных файлов приложения, зависимостей и метаданных. Они не содержат полноценную операционную систему.
- Образы представляют собой комбинацию нескольких слоев.
- У каждого образа есть уникальный идентификатор и тег для разных версий.



![Screenshot from 2022-11-02 11-57-19](https://user-images.githubusercontent.com/51878265/199414178-d59e8780-c140-4bf1-b27e-7e8f1c723afb.png)

### Dockerfile

> Commands:

- `FROM` (base image)
- `COPY` (copy files from local to the container)
- `ARG` (pass arguments)
- `ENV` (environment variable)
- `RUN` (any arbitrary shell command)
- `EXPOSE` (open port from container to virtual network)
- `CMD` (command to run when the container starts)
- `WORKDIR` (create a directory where all the files will be copied and used)

### Docker Build Architecture

We can build the image in two ways single architecture or multi-architecture. In the single architecture, we can build the image for a specific architecture, and in multi-architecture, we can build the image for multiple architectures.

#### Single Architecture

```bash
docker build -t <image-name> .
```

#### Multi-Architecture

```bash
docker buildx build --platform linux/amd64,linux/arm64 -t <image-name> .
```

**Good Practice**

- Copy the dependencies first and then copy the rest of the files.

```Dockerfile
COPY package.json ./
RUN npm install
COPY . ./
```

### .dockerignore

The .dockerignore file is used to specify files and directories that are not copied when using the `COPY` command.

### Docker Network

To connect to our created containers, Docker provides several network drivers. The available default drivers are bridge, host, and null.

- Bridge network creates a virtual network that allows containers to communicate with each other using IP addresses. We need to create a custom bridge network to enable DNS resolution between containers. Only containers connected to the same custom bridge network can communicate with each other directly. It doesn't work with the default bridge network.

- Host network uses the host machine's network stack inside the container. We can use this network for applications that require high network performance. We don't need to expose ports here.

- Using Null network driver disables the networking for the container.

![docker network](https://user-images.githubusercontent.com/37767537/223677649-babf850a-a87f-46bd-bb32-425801f05b2e.png)

To create a network, by default the created network will use the bridge network driver:

```bash
docker network create <network-name>
```

### Docker Volumes

We need volumes to persist our data, like databases and user info, because containers can go up and down, and we need some way to preserve our data.

We attach a volume during runtime:

```bash
docker run -v /path/in/container
```

**Named Volume**
We can also name the volume; otherwise, it will generate an ID and be hard to track:

```bash
docker run -v <volume name>:</path in container> <image name>
docker run -v myvolume:/src/public nginx
```

### Bind Mounting

A file or directory on the host machine is mounted into a container, i.e., it will match the condition of the file system inside a container.

```bash
docker run -v <path to your local system>:<container path>
docker run -v /app/content:/usr/share/nginx/html nginx
docker run -v $(pwd):/user/html nginx
```

In Compose, we don't have to give the `pwd`:

```yaml
volumes:
  - ./:/usr/share/nginx/html:ro
  - ./app:/usr/share/nginx/html/app:ro
```

### Docker Compose

- Compose helps us define and run multi-container Docker applications and configure relationships between containers.
- It also saves the hassle of entering the commands from the CLI.
- We have to write the configs in the YAML file, by default the file name is `docker-compose.yml`. We can run/stop by `docker compose up/down`.

The Skeleton of Docker Compose:

```yaml
services:  # containers. same as docker run
  servicename: # a friendly name. this is also the DNS name inside the network
    image: # Optional if you use to build:
    command: # Optional, replace the default CMD specified by the image
    environment: # Optional, same as -e in docker run
    volumes: # Optional, same as -v in docker run
  servicename2:

volumes: # Optional, same as docker volume create

networks: # Optional, same as docker network create
```

Sample:

```yaml
services:
  mongo:
    container_name: mongo
    image: mongo:4.0
    volumes:
      - mongo-db:/data/db
    networks:
      - my-net

volumes:
  mongo-db: # named volume

networks:
  my-net:
    driver: bridge
```

If any container depends on another container:

```yaml
depends_on:
  - mysql-primary
```

### Docker Swarm

Docker Swarm is an orchestration management tool that runs on Docker applications. Container orchestration automates the deployment, management, scaling, and networking of containers.

- Docker Swarm is not enabled by default. We have to enable it by:

```bash
docker swarm init
```

- In Swarm, we create services instead of creating the container directly.

### Docker Service

In Swarm, we don't create containers directly. Instead, we create a service that creates a container for us. A service can run multiple nodes on several nodes.

![Screenshot from 2022-11-08 13-07-01](https://user-images.githubusercontent.com/51878265/200502631-b574f4fc-8a0c-4e6f-8493-6d666ec1db2e.png)

### Docker Stack

When we have multiple services and need to establish the relationship between them, we use the stack. It is the same as the compose file. Here we don't use the `build:` object, and there is a new `deploy:` specific to swarm for things like replicas and secrets.

![Screenshot from 2022-11-04 13-34-28](https://user-images.githubusercontent.com/51878265/199923225-83fe75fc-406a-4d51-b2d4-15fb5ec6b4ee.png)

```yaml
deploy:
  replicas: 3
```

We deploy stack files with this command:

```bash
docker stack deploy -c file.yml <stackname>
```

### Docker Secrets

Docker Swarm supports secrets. We can pass ENV variables like SSH keys, usernames, and passwords with the help of secrets. We can pass secrets from the file or save the Docker secret.

We can create Docker secrets through CLI `external:`:

```bash
echo "<password text>" | docker secret create psql-pw -
```

or

Create a file with a password and then pass the path in the stack `file:`:

```yaml
services:
  postgres:
    image: postgres
    secrets:
      - post-pass
      - post-user
    environment:
      POSTGRES_PASSWORD_FILE: /run/secrets/post-pass
      POSTGRES_USER_FILE: /run/secrets/post-user

secrets:
  post-pass:
    external: true
  post-user:
    file: ./post-user.txt
```

### Docker Healthcheck

```dockerfile
HEALTHCHECK --interval=30s --timeout=3s \
CMD curl -f http://localhost/ || exit 1
```

### Container Registry

A repo - a collection of repositories. Use to store and access container images.

Some popular registries are:

- Docker Hub
- GitHub Container Registry (ghcr.io)
- Google Container Registry (gcr.io)
- Amazon Elastic Container Registry (ECR)
- Azure Container Registry (ACR)

### Private Docker Registry

We can create a registry with the official [Registry image](https://hub.docker.com/_/registry).

![image](https://user-images.githubusercontent.com/51878265/200518472-c520103f-11a8-4104-a859-32f5e3c6304e.png)

### What's next?

- [Docker Commands](./commands.md) - Learn about the most commonly used Docker commands.
- [Learning Resources](./learning-resources.md) - Learn more about Docker with these resources.
- [Other Resources](./other-resources.md) - Explore more about Docker with these resources.
