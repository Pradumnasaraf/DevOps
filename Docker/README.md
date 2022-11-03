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

**Sample docker-compose.yaml**

```yaml
version: '3.7'

services:
  devops-website:
    container_name: devops-website
    build: 
      context: .
      dockerfile: Dockerfile
    image: devops-website
    ports:
      - 8080:80
```
- Adding named volume

```yaml
  mongo:
    container_name: mongo
    image: mongo:4.0
    volumes:
      - mongo-db:/data/db
      
volumes:
  mongo-db: # named volume
```

