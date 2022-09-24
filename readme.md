

# Link Image Docker Hub
[AGMC-Task Docker Hub Image](https://hub.docker.com/repository/docker/jiran03/agmc-task)


## Set Env file
Buat file `.env` lalu setting environtmentnya dengan ketentuan variable sebagai berikut:

```
DB=mysql
DB_USERNAME=
DB_PASSWORD=
DB_PORT= 
DB_HOST=
DB_NAME=

JWT_SECRET= 
JWT_EXPIRED=

```


## Docker Command

### Build
`docker build -t <USERNAME_DOCKERHUB>/<REPOSITORY_DOCKERHUB>:<TAG> .`.

Contoh command `docker build -t jiran03/agmc-task:v1 .`

### Run
`docker run -p <OUT_PORT>:<IN_PORT>/tcp --name <CONTAINER_NAME> -it --rm <IMAGE>:<IMAGE_TAG>`.

Contoh command `docker run -p 0.0.0.0:80:8080/tcp --name agmc-container -it --rm jiran03/agmc-task:v1`

### Build & Run using Compose
`docker compose up --build`

### Push
`docker image push <IMAGE>:<IMAGE_TAG>`.

Contoh command `docker image push jiran03/agmc-task:v1`

### Pull
`docker pull <IMAGE>:<IMAGE_TAG>`.

Contoh command `docker pull jiran03/agmc-task:v1`