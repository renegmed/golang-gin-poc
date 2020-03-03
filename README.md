# Golang Gin Framework - Video API (POC)

## Go Module Init

```bash
go mod init gitlab.com/pragmaticreviews/golang-gin-poc
```

## Gin-Gonic library: github.com/gin-gonic/gin

## Run

```bash
go run server.go
```

# Deploy on AWS ElasticBeanstalk from EB CLI

## 1.- Create user (e.g "beanstalk") and allow Programmatic Access
## 2.- Create new group ( e.g. "Beanstalk")
## 3.- Attach AWSElasticBeanstalkFullAccess policy to the group
## 4.- Add the user to the Group
## 5.- Copy user's aws_access_key_id and aws_secret_access_key to .aws/config file
## 6.- From the application directory run:
## 6.a.- eb init
## 6.b.- eb create --single

# Deploy on AWS ElasticBeanstalk with Docker

## Build the docker image

```bash
docker build --tag pragmaticreviews/golang-gin .
```

## Run the container locally

```bash
docker run -p 5000:5000 pragmaticreviews/golang-gin
```

## Push the image to DockerHub (you need a DockerHub account)

```bash
docker login
```

```bash
docker push pragmaticreviews/golang-gin
```