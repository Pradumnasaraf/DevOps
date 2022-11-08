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

We a file `.dockerignore` which are not required when we copy files into the container.


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

A file or directory on the host machine is mounted into a container, i.e it will match the condition of the file syestem inside a conatiner.

```bash
docker run -v <path to your local sytem>:<conatiner path>
docker run -v /app/content:/usr/share/nginx/html  nginx
docker run -v $(pwd):/user/html nginx
```
 In compose we dont have to give the `pwd`
 
 ```yaml
     volumes:
      - ./:/usr/share/nginx/html:ro
      - ./app:/usr/share/nginx/html/app:ro 
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

```yaml
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

In swarm we don't create create containers directly, instead we create service and that creates container for us. A service can run multiple node on sevral nodes.

![Screenshot from 2022-11-08 13-07-01](https://user-images.githubusercontent.com/51878265/200502631-b574f4fc-8a0c-4e6f-8493-6d666ec1db2e.png)

### Docker Stack

When we have mutiple services and to establish the relation between them we use stack, it is same as compose file.
Here we don't use `build:` object and there is new `deploy:` specfic to swarm to like replicas, secrets.

![Screenshot from 2022-11-04 13-34-28](https://user-images.githubusercontent.com/51878265/199923225-83fe75fc-406a-4d51-b2d4-15fb5ec6b4ee.png)

```yaml
    deploy:
      replicas: 3
```
We deploy stack file with this command

```bash
docker stack deploy -c file.yml <stackname>
```

### Docker Secrets

Docker Swarm supports secrets. We can pass ENV variables like SSH keys, Usernames and passwords with help of that. We can pass secrets from the file or saved Docker secret.

-  We can create Docker secrets though CLI `external:`

```bash
echo "<password text>" | docker secret create psql-pw -
```

or

- Create a file with password and then pass the path in the stack `file:`

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
## Docker Healthcheck

```dockerfile
HEALTHCHECK --interval=30s --timeout=3s \
CMD curl -f http://localhost/ || exit 1
```

## Private Docker Registry

We can create a reg with the offical [Registry image](https://hub.docker.com/_/registry)

 <img src="https://user-images.githubusercontent.com/51878265/200518472-c520103f-11a8-4104-a859-32f5e3c6304e.png">


