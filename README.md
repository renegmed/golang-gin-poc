## Golang Gin Framework - Video API (POC)

## Go Module Init
```bash
go mod init gitlab.com/pragmaticreviews/golang-gin-poc
```

## Gin-Gonic library: github.com/gin-gonic/gin

## Run

```bash
go run server.go
```

## Deploy on AWS ElasticBeanstalk from EB CLI

# 1.- Create user (e.g "beanstalk") and allow Programmatic Access
# 2.- Create new group ( e.g. "Beanstalk")
# 3.- Attach AWSElasticBeanstalkFullAccess policy to the group
# 4.- Add the user to the Group
# 5.- Copy user's aws_access_key_id and aws_secret_access_key to .aws/config file
# 6.- From the application directory run:
# 6.a.- eb init
# 6.b.- eb create --single