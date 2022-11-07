# Concepts


## ⚫ Docker Images

- Images are made up of app binaries, dependencies, and metadata. Don't contain a full OS.
- Images are combination of mutiple layers.
- Each Image has it down unique ID and a tag for different version.

![Screenshot from 2022-11-02 11-57-19](https://user-images.githubusercontent.com/51878265/199414178-d59e8780-c140-4bf1-b27e-7e8f1c723afb.png)

## ⚫ Dockerfile

> Commands:

- `FROM` (base image)
- `COPY` (copy files from local to conatiner)
- `ARG` (pass arrugments)
- `ENV` (environment variable)
- `RUN` (any arbitrary shell command)
- `EXPOSE` (open port from container to virtual network)
- `CMD` (command to run when container starts) 
- `WORKDIR` (Create a dir where all the file will be copy and used.)

To build an image from the **Dockerfile**, use this command

```bash
docker build <path> 
// docker build .
```

**Good Practice**

- Copy the dependencies 1st and then copy rest of the files.

```Dockerfile
COPY package.json ./
RUN npm install
COPY . ./
```

## ⚫ .dockerignore

We a file `..dockerignore` which are not required when we copy files into the container.


## ⚫ Docker Volumes

We need volume to Persist our data, like databases and user info, because conatiner can go up and down, and we need some way preserve our data.

We attach volume during run time

```bash
docker run -v /path/in/container
```

**Named Volume**
We can also name the volume otherwise it will generate the ID and hard to track

```bash
docker run -v <volume name>:</path in container> <image name>
docker run -v myvolume:/src/public nginx
```

### Bind Mounting

```
docker run -v <path to your local sytem>:<conatiner path>
docker run -v /app/content:/usr/share/nginx/html  nginx
docker run -v $(pwd):/user/html nginx
```

## ⚫ Docker Compose

- Compose help us defining and running multi-container Docker applications and configure relationships between containers 
- It also save hassel from entering the commands from the CLI.
- We have to write the configs in the YAML file, by default the file name is `docker-compose.yml`. We can run/stop by `docker compose up/down`

Skeleton of Docker compose

```yaml
services:  # containers. same as docker run
  servicename: # a friendly name. this is also DNS name inside network
    image: # Optional if you use build:
    command: # Optional, replace the default CMD specified by the image
    environment: # Optional, same as -e in docker run
    volumes: # Optional, same as -v in docker run
  servicename2:

volumes: # Optional, same as docker volume create

networks: # Optional, same as docker network create
```

Sample:

```yaml
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

If any container depends on another container

```
depends_on:
  - mysql-primary
```

## ⚫ Docker Swarm

Docker Swarm is an orchestration management tool that runs on Docker applications.Container orchestration automates the deployment, management, scaling, and networking of containers

- Docker Swarm is not enabled by deafult, we have enbaled by

```bash
docker swarm init
```

- In this we create services, insated of creating of container directly

### Docker service

In swarm we don't create create containers directly, instead we create service and that creates container for us.


### Docker Stack

When we have mutiple services and to establish the relation between them we use stack, it is same as compose file.
Here we don't use `build:` object and there is new `deploy:` specfic to swarm to like replicas, secrets.

```yaml
    deploy:
      replicas: 3
```

- It also supports secrets

- To create a secret from the terminal

```bash
echo "<pw>" | docker secret create psql-pw -
```

```
services:
  psql:
    image: postgres
    secrets:
      - psql_user
      - psql_password
    environment:
      POSTGRES_PASSWORD_FILE: /run/secrets/psql_password
      POSTGRES_USER_FILE: /run/secrets/psql_user

secrets:
  psql_user:
    file: ./psql_user.txt
  psql_password:
    file: ./psql_password.txt
```


![Screenshot from 2022-11-04 13-34-28](https://user-images.githubusercontent.com/51878265/199923225-83fe75fc-406a-4d51-b2d4-15fb5ec6b4ee.png)



